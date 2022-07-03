package apis

import (
	"WordleServer/db"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type WordleApi struct {
}

func (_ WordleApi) GetAllWords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	result := db.QueryAllWords()
	fmt.Fprint(w, result)
}
