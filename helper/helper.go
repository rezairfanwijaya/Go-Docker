package helper

import (
	"go-docker/user"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type userFormat struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// mengambil env
func GetEnv(path string) (map[string]string, error) {
	env, err := godotenv.Read(path)
	if err != nil {
		return env, err
	}

	return env, nil
}

func ErrorBindingFormatter(errs error) []string {
	var myErr []string
	for _, err := range errs.(validator.ValidationErrors) {
		myErr = append(myErr, err.Error())
	}

	return myErr
}

func ResponseAPIFormatter(status, message string, code int, data interface{}) responseAPI {
	meta := meta{
		Status:  status,
		Code:    code,
		Message: message,
	}

	return responseAPI{
		Meta: meta,
		Data: data,
	}
}

func UserFormatter(user user.User) userFormat {
	return userFormat{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
