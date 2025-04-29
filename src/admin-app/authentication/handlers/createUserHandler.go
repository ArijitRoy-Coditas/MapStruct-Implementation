package handlers

import (
	genericConstants "ABCD/src/constants"
	genericModels "ABCD/src/models"
	mapStruct "ABCD/src/utils/mapStruct"
	"ABCD/src/utils/validation"
	"authentication/business"
	"authentication/models"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

type createUserController struct {
	service *business.CreateUserService
}

func NewGetCreateUserController(service *business.CreateUserService) *createUserController {
	return &createUserController{
		service: service,
	}
}

func (controller *createUserController) HandleCreateUser(ctx *gin.Context) {
	var bffCreateUserRequest models.BFFCreateUserRequest

	if err := ctx.ShouldBindJSON(&bffCreateUserRequest); err != nil {
		errorMsgs := genericModels.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: genericConstants.JsonBindingFieldError}
		ctx.IndentedJSON(400, errorMsgs)
	}

	if err := validation.GetBFFValidation().Struct(&bffCreateUserRequest); err != nil {
		ValidationErros := validation.FormatValidation(err)
		ctx.IndentedJSON(400, ValidationErros)
	}

	var response models.BFFCreateUserResponse
	fmt.Println("---- Request Model Field Types & Values ----")
	vReq := reflect.ValueOf(bffCreateUserRequest)
	tReq := reflect.TypeOf(bffCreateUserRequest)
	for i := 0; i < tReq.NumField(); i++ {
		fmt.Printf("Field: %-12s Type: %-10T Value: %v\n", tReq.Field(i).Name, vReq.Field(i).Interface(), vReq.Field(i).Interface())
	}

	mapStruct.MapStruct(bffCreateUserRequest, &response)

	fmt.Println("---- Response Model Field Types & Values ----")
	vResp := reflect.ValueOf(response)
	tResp := reflect.TypeOf(response)
	for i := 0; i < tResp.NumField(); i++ {
		fmt.Printf("Field: %-12s Type: %-10T Value: %v\n", tResp.Field(i).Name, vResp.Field(i).Interface(), vResp.Field(i).Interface())
	}

	ctx.IndentedJSON(200, response)

}
