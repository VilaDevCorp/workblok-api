package template

import (
	"sensei/svc"
	"sensei/svc/template"
	"sensei/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "github.com/google/uuid"
)

func Create(c *gin.Context) {
	var form template.CreateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Template.Create(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkCreated(result)
	c.JSON(res.Status, res.Result)
}

func Update(c *gin.Context) {
	var form template.UpdateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Template.Update(c.Request.Context(), form)
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
	result, err := svc.Template.Get(c.Request.Context(), parsedId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func Search(c *gin.Context) {
	var form template.SearchForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Template.Search(c.Request.Context(), form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(result)
	c.JSON(res.Status, res.Result)
}

func Delete(c *gin.Context) {
	var form template.DeleteForm
	err := c.ShouldBind(&form)

	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	templateIds := form.TemplateIds

	svc := svc.Get()
	err = svc.Template.Delete(c.Request.Context(), templateIds)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkDeleted()
	c.JSON(res.Status, res.Result)
}

func CreateTask(c *gin.Context) {
	var form template.CreateTaskForm
	unparsedId, _ := c.Params.Get("id")
	parsedId, err := uuid.Parse(unparsedId)
	if err != nil {
		res := utils.BadRequest(unparsedId, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}

	err = c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	result, err := svc.Template.CreateTask(c.Request.Context(), parsedId, form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkCreated(result)
	c.JSON(res.Status, res.Result)
}

func DeleteTasks(c *gin.Context) {
	var form template.DeleteTasksForm
	err := c.ShouldBind(&form)

	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	taskIds := form.TaskIds

	svc := svc.Get()
	err = svc.Template.DeleteTasks(c.Request.Context(), taskIds)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkDeleted()
	c.JSON(res.Status, res.Result)
}

func ApplyTemplate(c *gin.Context) {
	var form template.ApplyTemplateForm
	unparsedId, _ := c.Params.Get("id")
	parsedId, err := uuid.Parse(unparsedId)
	if err != nil {
		res := utils.BadRequest(unparsedId, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}

	err = c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	err = svc.Template.ApplyTemplate(c.Request.Context(), parsedId, form)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkOperation(nil)
	c.JSON(res.Status, res.Result)
}
