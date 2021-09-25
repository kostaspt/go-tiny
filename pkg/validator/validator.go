package validator

import (
	"net/http"

	gpValidator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *gpValidator.Validate
}

func New() *Validator {
	return &Validator{validator: gpValidator.New()}
}

type Response struct {
	Errors map[string]string `json:"errors"`
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errs := err.(gpValidator.ValidationErrors)

		resp := Response{
			Errors: make(map[string]string, len(errs)),
		}

		for _, e := range errs {
			resp.Errors[e.Field()] = e.Tag()
		}

		return echo.NewHTTPError(http.StatusUnprocessableEntity, resp)
	}
	return nil
}
