package routers

import (
	"WordleServer/apis"
	"github.com/julienschmidt/httprouter"
)

func InitRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/words/getAll", apis.WordleApi{}.GetAllWords)
	return router
}
