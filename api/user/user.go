package user

import (
	"workblok/svc"
	"workblok/svc/user"
	"workblok/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(c *gin.Context) {
	var form user.CreateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.User.Create(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkCreated(result)
	c.JSON(res.Status, res.Result)
}

func ConfUpdate(c *gin.Context) {
	var form user.ConfigForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	jwt, _ := c.Cookie("JWT_TOKEN")
	tokenClaims, _ := utils.ValidateToken(jwt)
	form.Id = &tokenClaims.Id
	result, err := svc.User.Update(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkUpdated(result)
	c.JSON(res.Status, res.Result)
}

func CompleteTutorial(c *gin.Context) {
	unparsedId, _ := c.Params.Get("id")
	parsedId, err := uuid.Parse(unparsedId)
	if err != nil {
		res := utils.BadRequest(unparsedId, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	err = svc.User.CompleteTutorial(c.Request.Context(), parsedId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkOperation(nil)
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
	result, err := svc.User.Get(c.Request.Context(), parsedId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func Search(c *gin.Context) {
	var form user.SearchForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.User.Search(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func Delete(c *gin.Context) {
	unparsedId, _ := c.Params.Get("id")
	parsedId, err := uuid.Parse(unparsedId)
	if err != nil {
		res := utils.BadRequest(unparsedId, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	err = svc.User.Delete(c.Request.Context(), parsedId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkDeleted()
	c.JSON(res.Status, res.Result)
}
