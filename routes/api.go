package routes

import (
	controllers "bpjs-mono/controllers"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var prefix = "/api/v1"
var Auth = middleware.JWT([]byte(os.Getenv("JWT_SECRET")))

func implContains(sl []string, name string) bool {
	// iterate over the array and compare given string to each element
	for _, value := range sl {
		if strings.Contains(name, value) {
			return true
		}
	}
	return false
}

// Build returns the application routes
func Build(e *echo.Echo) {
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte(os.Getenv("JWT_SECRET")),
	// 	Skipper: func(c echo.Context) bool {
	// 		// Skip middleware if path is equal 'login'
	// 		log.Println(c.Request().URL.Path)
	// 		arrPath := []string{
	// 			prefix + "/qoincash/callback/status-payment",
	// 			prefix + "/qoincash/callback/sso-token",

	// 			/* API Unused JWT */
	// 			prefix + "/access/registerUser",
	// 			prefix + "/access/login",
	// 		}

	// 		var skip bool = true
	// 		skip = implContains(arrPath, strings.ToLower(c.Request().URL.Path))
	// 		return skip
	// 	},
	// }))


	//e.Use(echoMidleware.JWTWithConfig(DefaultJWTConfig))
	RouteGeneralApi(e)
}

func RouteGeneralApi(e *echo.Echo) {
	r := e.Group(prefix)

	/* Transaction */
	r.POST("/bpjs/transaction", controllers.CreateTransaction)
	r.GET("/bpjs/transaction", controllers.GetListTransactions)
	r.GET("/bpjs/transaction/:request_id", controllers.GetDetailTransaction)
	r.POST("/bpjs/transactionJson", controllers.CreateTransactionJSON)
	r.GET("/bpjs/transactionJson", controllers.GetListTransactionsJSON)
	r.GET("/bpjs/transactionJson/:request_id", controllers.GetDetailTransactionJson)
}