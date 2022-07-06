package main

import (
	"WordleServer/routers"
	"fmt"
	"log"
	"net/http"
)

// @title wordle服务端
// @version 1.0
// @description 提供wordle小游戏后端支持
// @termsOfService http://swagger.io/terms/

// @contact.name 771447612@qq.com
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /
func main() {
	router := routers.InitRouter()
	fmt.Println("开始监听 7767端口")
	log.Fatal(http.ListenAndServe(":7767", router))
}
