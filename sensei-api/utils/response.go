package utils

import (
	"fmt"
	"net/http"
)

type ResponseResult struct {
	Message string      `json:"message"`
	Obj     interface{} `json:"obj"`
	Err     error       `json:"err"`
}

type HttpResponse struct {
	Status int
	Result ResponseResult
}

func OkOperation(result interface{}) HttpResponse {
	return HttpResponse{Status: http.StatusOK, Result: ResponseResult{Message: "Succesful operation", Obj: result}}
}

func OkCreated(result interface{}) HttpResponse {
	return HttpResponse{Status: http.StatusOK, Result: ResponseResult{Message: "Entity succesfully created", Obj: result}}
}

func OkUpdated(result interface{}) HttpResponse {
	return HttpResponse{Status: http.StatusOK, Result: ResponseResult{Message: "Entity succesfully updated", Obj: result}}
}

func OkGet(result interface{}) HttpResponse {
	return HttpResponse{Status: http.StatusOK, Result: ResponseResult{Message: "Here is your entity", Obj: result}}
}

func OkDeleted() HttpResponse {
	return HttpResponse{Status: http.StatusOK, Result: ResponseResult{Message: "Entity succesfully deleted"}}
}

func OkLogged(result LoginResult) HttpResponse {
	return HttpResponse{Status: http.StatusOK, Result: ResponseResult{Message: "Logged succesfully!", Obj: result}}
}

func BadRequest(result interface{}, err error) HttpResponse {
	return HttpResponse{Status: http.StatusBadRequest, Result: ResponseResult{Message: "Form is not correct", Obj: result, Err: err}}
}

func NotFoundEntity(id string) HttpResponse {
	return HttpResponse{Status: http.StatusBadRequest, Result: ResponseResult{Message: "Entity wasnt found", Obj: id, Err: nil}}
}

func InternalError(err error) HttpResponse {
	fmt.Print(err)
	return HttpResponse{Status: http.StatusInternalServerError, Result: ResponseResult{Message: "Internal error occurred", Err: err}}
}

func Forbidden(cause string, err error) HttpResponse {
	return HttpResponse{Status: http.StatusForbidden, Result: ResponseResult{Message: fmt.Sprintf("You dont have permission for this operation: %s", cause), Err: err}}
}

func Unauthorized(cause string) HttpResponse {
	return HttpResponse{Status: http.StatusUnauthorized, Result: ResponseResult{Message: fmt.Sprintf("Unauthorized user: %s", cause)}}
}

func TaskAlreadyCompleted() HttpResponse {
	return HttpResponse{Status: http.StatusConflict, Result: ResponseResult{Message: "Some tasks has already been completed/uncompleted"}}
}
