package block

import (
	"net/http"
	"workblok/svc"
	"workblok/svc/block"
	"workblok/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(c *gin.Context) {
	var form block.CreateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	jwt, _ := c.Cookie("JWT_TOKEN")
	tokenClaims, _ := utils.ValidateToken(jwt)
	form.UserId = tokenClaims.Id
	svc := svc.Get()
	result, err, resultCode := svc.Block.Create(c.Request.Context(), form)
	if resultCode == http.StatusInternalServerError {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	if resultCode == http.StatusConflict {
		res := utils.BlockNotFinishedExists()
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}

	res := utils.OkCreated(result)
	c.JSON(res.Status, res.Result)
}

func ApplyPenalty(c *gin.Context) {
	var form block.UpdateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	block, err := svc.Block.Get(c, form.BlockId)
	if err != nil {
		res := utils.NotFoundEntity(string(form.BlockId.String()))
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	jwt, _ := c.Cookie("JWT_TOKEN")
	tokenClaims, _ := utils.ValidateToken(jwt)
	if block.Edges.User.ID != tokenClaims.Id {
		res := utils.Unauthorized("You can't apply a penalty to a block that is not yours", "")
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	result, err := svc.Block.Update(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkUpdated(result)
	c.JSON(res.Status, res.Result)
}

func Finish(c *gin.Context) {
	unparsedId, _ := c.Params.Get("id")
	parsedId, err := uuid.Parse(unparsedId)
	if err != nil {
		res := utils.BadRequest(unparsedId, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	isAutoParam := c.Query("auto")
	isAuto := false
	if isAutoParam == "true" {
		isAuto = true
	}

	svc := svc.Get()
	block, err := svc.Block.Get(c, parsedId)
	if err != nil {
		res := utils.NotFoundEntity(string(parsedId.String()))
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	jwt, _ := c.Cookie("JWT_TOKEN")
	tokenClaims, _ := utils.ValidateToken(jwt)
	if block.Edges.User.ID != tokenClaims.Id {
		res := utils.Unauthorized("You can't finish a block that is not yours", "")
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	result, err := svc.Block.Finish(c.Request.Context(), parsedId, isAuto)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkUpdated(result)
	c.JSON(res.Status, res.Result)
}

func Get(c *gin.Context) {
	unparsedId, _ := c.Params.Get("id")
	parsedId, err := uuid.Parse(unparsedId)
	if err != nil {
		res := utils.BadRequest(unparsedId, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Block.Get(c.Request.Context(), parsedId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func GetActive(c *gin.Context) {
	jwt, _ := c.Cookie("JWT_TOKEN")
	tokenClaims, _ := utils.ValidateToken(jwt)
	svc := svc.Get()
	result, err := svc.Block.GetActive(c.Request.Context(), tokenClaims.Id)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	if result == nil {
		res := utils.NotFoundEntity("Active block")
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func Search(c *gin.Context) {
	var form block.SearchForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	jwt, _ := c.Cookie("JWT_TOKEN")
	tokenClaims, _ := utils.ValidateToken(jwt)
	form.UserId = tokenClaims.Id

	svc := svc.Get()
	result, err := svc.Block.Search(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func Delete(c *gin.Context) {
	var form block.DeleteForm
	err := c.ShouldBind(&form)

	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	blockIds := form.BlockIds

	svc := svc.Get()
	err = svc.Block.Delete(c.Request.Context(), blockIds)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkDeleted()
	c.JSON(res.Status, res.Result)
}

func Stats(c *gin.Context) {
	var form block.StatsForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	jwt, _ := c.Cookie("JWT_TOKEN")
	tokenClaims, _ := utils.ValidateToken(jwt)
	form.UserId = &tokenClaims.Id
	svc := svc.Get()
	result, err := svc.Block.Stats(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkOperation(result)
	c.JSON(res.Status, res.Result)
}
