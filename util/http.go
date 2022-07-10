package util

import (
	"fmt"
	"net/http"
)

type AllowCROS struct {
}

func (AllowCROS) SetupCORS(w *http.ResponseWriter) {
	if w == nil {
		return
	}
	fmt.Println("set CORS succeed")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
	(*w).Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
	(*w).Header().Set("Access-Control-Max-Age", "172800")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("content-type", "application/json")
}
