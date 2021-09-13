package util

import (
	"encoding/json"
	"os"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/raismaulana/assetsP/application/apperror"
)

var Trans ut.Translator

// MustJSON is converter from interface{} to string
// Warning! this function will always assume the convertion is success
// if you are not sure the convertion is always succeed then use ToJSON
func MustJSON(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}

// GetValidationErrorMessage is extractor error message from binding http request data
func GetValidationErrorMessage(err error) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return apperror.FailUnmarshalResponseBodyError
	}
	var errorMessages []string
	for i, e := range errs {
		errorMessages[i] = e.Translate(Trans)
	}
	errorMessage := strings.Join(errorMessages, " \n")
	return apperror.ERR400.Var(errorMessage)
}

// A new folder is created at the root of the project.
func CreateDirectoryIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return apperror.ERR500.Var(err)
		}
	}
	return nil
}
