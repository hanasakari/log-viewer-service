package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 升级http为websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 定义端口号参数形式
var port = flag.Int("p", 8000, "请使用int类型参数")

func main() {
	flag.Parse()
	fmt.Println("-port:", *port)
	start(*port)
}

func start(port int) {
	address := fmt.Sprintf("%s:%d", "localhost", port)
	router := gin.New()
	router.GET("/", getMessage)
	router.Run(fmt.Sprintf(":%d", port))
	log.Println("Serving at " + address)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal("failed run app: ", err)

	}
}

func getMessage(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(message))
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}

	}
}
