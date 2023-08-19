package task

import (
	"net/http"
	"sensei/svc"
	"sensei/svc/task"
	"sensei/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(c *gin.Context) {
	var form task.CreateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Task.Create(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkCreated(result)
	c.JSON(res.Status, res.Result)
}

func Update(c *gin.Context) {
	var form task.UpdateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Task.Update(c.Request.Context(), form)
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
	result, err := svc.Task.Get(c.Request.Context(), parsedId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func Search(c *gin.Context) {
	var form task.SearchForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Task.Search(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func Delete(c *gin.Context) {
	var form task.DeleteForm
	err := c.ShouldBind(&form)

	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	taskIds := form.TaskIds

	svc := svc.Get()
	err = svc.Task.Delete(c.Request.Context(), taskIds)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkDeleted()
	c.JSON(res.Status, res.Result)
}

func Complete(c *gin.Context) {
	var form task.CompleteForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	status, err := svc.Task.Complete(c.Request.Context(), form.TaskIds, form.IsCompleted)
	if status == http.StatusInternalServerError {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	if status == http.StatusConflict {
		res := utils.TaskAlreadyCompleted()
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	result := struct {
		id         []uuid.UUID
		isComplete bool
	}{
		id:         form.TaskIds,
		isComplete: form.IsCompleted,
	}
	res := utils.OkOperation(result)
	c.JSON(res.Status, res.Result)
}

func Stats(c *gin.Context) {
	var form task.StatsForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Task.Stats(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkOperation(result)
	c.JSON(res.Status, res.Result)
}
