package server

import (
	"net/http"
	"strconv"

	"github.com/sunderee/password-service/service"
)

func ServerHandler(writer http.ResponseWriter, request *http.Request) {
	passwordOptions := new(service.Options)

	queryParams := request.URL.Query()
	length := queryParams.Get("length")
	if len(length) == 0 {
		passwordOptions.PasswordLength = 8
	} else {
		convertedLength, conversionError := strconv.ParseInt(length, 10, 0)
		if conversionError != nil {
			writer.WriteHeader(http.StatusNotAcceptable)
			writer.Write([]byte("Error with the length parameter; are you sure it's an integer?"))
			return
		}
		passwordOptions.PasswordLength = int(convertedLength)
	}

	hasUppercase := queryParams.Get("upper")
	if len(hasUppercase) == 0 {
		passwordOptions.HasUppercase = false
	} else {
		boolean, conversionError := strconv.ParseBool(hasUppercase)
		if conversionError != nil {
			writer.WriteHeader(http.StatusNotAcceptable)
			writer.Write([]byte("Error with the upper parameter; are you sure it's a boolean?"))
			return
		}
		passwordOptions.HasUppercase = boolean
	}

	hasNumbers := queryParams.Get("numbers")
	if len(hasNumbers) == 0 {
		passwordOptions.HasNumbers = false
	} else {
		boolean, conversionError := strconv.ParseBool(hasNumbers)
		if conversionError != nil {
			writer.WriteHeader(http.StatusNotAcceptable)
			writer.Write([]byte("Error with the numbers parameter; are you sure it's a boolean?"))
			return
		}
		passwordOptions.HasNumbers = boolean
	}

	hasSpecials := queryParams.Get("specials")
	if len(hasSpecials) == 0 {
		passwordOptions.HasSpecials = false
	} else {
		boolean, conversionError := strconv.ParseBool(hasSpecials)
		if conversionError != nil {
			writer.WriteHeader(http.StatusNotAcceptable)
			writer.Write([]byte("Error with the specials parameter; are you sure it's a boolean?"))
			return
		}
		passwordOptions.HasSpecials = boolean
	}

	generatedPassword := service.GeneratePassword(*passwordOptions)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(generatedPassword))
}
