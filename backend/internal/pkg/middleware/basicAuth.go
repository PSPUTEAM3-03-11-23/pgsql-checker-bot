package middleware

import (
	"github.com/gin-gonic/gin"
)

func DecodeCredentials(c *gin.Context) (string, string, bool) {
	r := c.Request
	return r.BasicAuth()
}

//
//func GetAccountByCreds(c *gin.Context) (*entity.Account, error) {
//	login, password, ok := DecodeCredentials(c)
//	if !ok {
//		return nil, errors.New("")
//	}
//
//	account := &entity.Account{
//		Email:    login,
//		Password: password,
//	}
//
//	accountRepo := repository.NewAccountRepository(helpers.GetConnectionOrCreateAndGet())
//	accountService := service.NewAccountService(accountRepo)
//
//	return accountService.GetByCreds(account), nil
//
//}

// BasicAuth middleware для basic auth
func BasicAuth(c *gin.Context) {
	//acc, err := GetAccountByCreds(c)
	//
	//if err != nil || acc.Id == 0 {
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//	c.Next()
	//	return
	//}
	//
	//c.Set("account", acc)
	c.Next()
}
