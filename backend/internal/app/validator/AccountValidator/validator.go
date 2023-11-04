package AccountValidator

import (
	"fmt"
	"hackathon-tg-bot/internal/app/model/entity"
	"hackathon-tg-bot/internal/app/validator"
	"hackathon-tg-bot/pkg/errorHandler"
	"net/http"
	"net/mail"
)

func ValidateAccountRegistration(account *entity.Account) *errorHandler.HttpErr {
	if validator.IsStringEmpty(account.FirstName) {
		return errorHandler.New("firstName is empty", http.StatusBadRequest)
	}

	if validator.IsStringEmpty(account.LastName) {
		return errorHandler.New("lastName is empty", http.StatusBadRequest)
	}

	if validator.IsStringEmpty(account.Email) {
		return errorHandler.New("email is empty", http.StatusBadRequest)
	}
	_, err := mail.ParseAddress(account.Email)
	if err != nil {
		return errorHandler.New("email is invalid", http.StatusBadRequest)
	}

	if validator.IsStringEmpty(account.Password) {
		return errorHandler.New("password is empty", http.StatusBadRequest)
	}
	return nil
}

func ValidateAccount(account *entity.Account) *errorHandler.HttpErr {
	httpErr := ValidateAccountRegistration(account)
	if httpErr != nil {
		return httpErr
	}

	if account.Role != entity.AdminRole && account.Role != entity.ChipperRole && account.Role != entity.UserRole {
		return errorHandler.New(fmt.Sprintf("role must be in [%s, %s, %s]", entity.AdminRole, entity.ChipperRole, entity.UserRole), http.StatusBadRequest)
	}

	return nil
}
