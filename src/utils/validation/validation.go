package validation

import (
	"ABCD/src/constants"
	"ABCD/src/models"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var bffValidation *validator.Validate

func GetBFFValidation() *validator.Validate {
	return bffValidation
}

func FormatValidation(err error) []models.ErrorMessage {
	validationErrorMsgs := err.(validator.ValidationErrors)
	errorMsgs := make([]models.ErrorMessage, len(validationErrorMsgs))
	for i, fieldError := range validationErrorMsgs {
		errorMsgs[i] = models.ErrorMessage{
			Key:          strings.ToLower(fieldError.StructField()),
			ErrorMessage: strings.ToLower(fieldError.Field() + " required"),
		}
	}
	return errorMsgs
}

func panCardValidation(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(constants.PanCardRegex, fl.Field().String())
	return matched
}
func init() {
	bffValidation = validator.New()
	bffValidation.RegisterValidation("panCard", panCardValidation)
}
