package httpErrors

import (
	"MarketEye/config"
	"MarketEye/pkg/logger"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var showUnknownErrors bool
var lgr logger.Logger

func Init(c *config.Config, lgrIn logger.Logger) {
	if c.Server.ShowUnknownErrorsInResponse {
		showUnknownErrors = true
	}
	lgr = lgrIn
}

type responseMsg struct {
	Message string `json:"message"`
}

func Handler(c *fiber.Ctx, err error) error {
	//defer sentry.Flush(2 * time.Second)
	var response responseMsg
	var statusCode int
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"app_user_username_uindex\"") {
		response.Message = "username unique constraint"
		statusCode = fiber.StatusBadRequest
	} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"app_user_phone_uindex\"") {
		response.Message = "phone unique constraint"
		statusCode = fiber.StatusBadRequest
	} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"app_user_email_uindex\"") {
		response.Message = "email unique constraint"
		statusCode = fiber.StatusBadRequest
	} else if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"qiwi_vault_name_uindex\"") {
		response.Message = "wallet with this name already exists"
		statusCode = fiber.StatusConflict
	} else if strings.Contains(err.Error(), "Field validation for") {
		response.Message = "Неверные данные"
		if showUnknownErrors {
			response.Message = fmt.Sprintf("%s:\n\n %s (на проде этого сообщения не будет)", response.Message, err.Error())
		}
		statusCode = fiber.StatusBadRequest
	} else if err.Error() == "not found confirmation code" {
		response.Message = err.Error()
		statusCode = fiber.StatusBadRequest
	} else if err.Error() == "wrong current password hash" {
		response.Message = err.Error()
		statusCode = fiber.StatusBadRequest
	} else if err.Error() == "session doesnt found" {
		response.Message = err.Error()
		statusCode = fiber.StatusUnauthorized
	} else if err.Error() == "already checked in" {
		response.Message = err.Error()
		statusCode = fiber.StatusBadRequest
	} else if err.Error() == "registration record not found" {
		response.Message = err.Error()
		statusCode = fiber.StatusBadRequest
	} else if err.Error() == "no Device-UserID" {
		//response.Message = err.Error()
		statusCode = fiber.StatusBadRequest
	}

	if statusCode == 0 {
		statusCode = fiber.StatusInternalServerError
	}
	if response.Message == "" && showUnknownErrors {
		response.Message = fmt.Sprintf("на проде этого сообщения не будет:\n\n %s", err.Error())
	} else if response.Message == "" {
		lgr.Error(err)
		response.Message = "unknown error"
	}

	//sentry.CaptureException(err)
	return c.Status(statusCode).JSON(response)
}
