package routers

import (
	"WordleServer/apis"
	"WordleServer/util"

	"github.com/julienschmidt/httprouter"
)

func InitRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/words/getAll", apis.WordleApi{AllowCROS: &util.AllowCROS{}}.GetAllWords)
	router.GET("/words/random-index", apis.WordleApi{AllowCROS: &util.AllowCROS{}}.GetIndex)
	return router
}
