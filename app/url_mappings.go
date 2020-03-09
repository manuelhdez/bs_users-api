package app

import (
	"github.com/manuelhdez/bs_users-api/controllers/ping"
	"github.com/manuelhdez/bs_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search/:term", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
