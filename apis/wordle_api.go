package apis

import (
	"WordleServer/db"
	"WordleServer/util"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type WordleApi struct {
	*util.AllowCROS
}

var date int = 0
var dayIndex int = 0

func (api WordleApi) GetAllWords(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	result := db.QueryAllWords()
	api.SetupCORS(&w)
	res, _ := json.Marshal(result)
	resStr := string(res)
	fmt.Println(w.Header().Get("Access-Control-Allow-Origin"))
	fmt.Fprint(w, resStr)
}

func (api WordleApi) GetIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	api.SetupCORS(&w)

	now := time.Now()
	day := now.Day()
	if date != day {
		result := db.QueryAllWords()
		rand.Seed(int64(day))
		index := rand.Intn(len(result))
		dayIndex = index
		date = day
	}

	fmt.Fprint(w, dayIndex)
}
