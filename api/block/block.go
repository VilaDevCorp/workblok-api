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
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	userId := c.Request.Context().Value(utils.UserIdKey).(uuid.UUID)
	form.UserId = userId
	svc := svc.Get()
	result, err, resultCode := svc.Block.Create(c.Request.Context(), form)
	if resultCode == http.StatusInternalServerError {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	if resultCode == http.StatusConflict {
		res := utils.ErrorResponse(http.StatusConflict, nil, nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}

	res := utils.SuccessResponse(result)
	c.JSON(res.Status, res.Body)
}

func ApplyPenalty(c *gin.Context) {
	var form block.UpdateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	svc := svc.Get()
	block, err := svc.Block.Get(c, form.BlockId)
	if err != nil {
		res := utils.NotFound("Block", string(form.BlockId.String()))
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	userId := c.Request.Context().Value(utils.UserIdKey).(uuid.UUID)
	if block.Edges.User.ID != userId {
		res := utils.ErrorResponse(
			http.StatusForbidden,
			nil,
			nil,
		)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	result, err := svc.Block.Update(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	res := utils.SuccessResponse(result)
	c.JSON(res.Status, res.Body)
}

func Finish(c *gin.Context) {
	unparsedId, _ := c.Params.Get("id")
	parsedId, err := uuid.Parse(unparsedId)
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
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
		res := utils.NotFound("Block", string(parsedId.String()))
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	userId := c.Request.Context().Value(utils.UserIdKey).(uuid.UUID)
	if block.Edges.User.ID != userId {
		res := utils.ErrorResponse(
			http.StatusForbidden,
			nil,
			nil,
		)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	result, err := svc.Block.Finish(c.Request.Context(), parsedId, isAuto)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	res := utils.SuccessResponse(result)
	c.JSON(res.Status, res.Body)
}

func Get(c *gin.Context) {
	unparsedId, _ := c.Params.Get("id")
	parsedId, err := uuid.Parse(unparsedId)
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	svc := svc.Get()
	result, err := svc.Block.Get(c.Request.Context(), parsedId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	res := utils.SuccessResponse(result)
	c.JSON(res.Status, res.Body)
}

func GetActive(c *gin.Context) {
	userId := c.Request.Context().Value(utils.UserIdKey).(uuid.UUID)
	svc := svc.Get()
	result, err := svc.Block.GetActive(c.Request.Context(), userId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	if result == nil {
		res := utils.NotFound("Active block", "")
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	res := utils.SuccessResponse(result)
	c.JSON(res.Status, res.Body)
}

func Search(c *gin.Context) {
	var form block.SearchForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	userId := c.Request.Context().Value(utils.UserIdKey).(uuid.UUID)
	form.UserId = userId

	svc := svc.Get()
	result, err := svc.Block.Search(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	res := utils.SuccessResponse(result)
	c.JSON(res.Status, res.Body)
}

func Delete(c *gin.Context) {
	var form block.DeleteForm
	err := c.ShouldBind(&form)

	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	blockIds := form.BlockIds

	svc := svc.Get()
	err = svc.Block.Delete(c.Request.Context(), blockIds)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	res := utils.SuccessResponse(nil)
	c.JSON(res.Status, res.Body)
}

func Stats(c *gin.Context) {
	var form block.StatsForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	userId := c.Request.Context().Value(utils.UserIdKey).(uuid.UUID)
	form.UserId = &userId
	svc := svc.Get()
	result, err := svc.Block.Stats(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	res := utils.SuccessResponse(result)
	c.JSON(res.Status, res.Body)
}
