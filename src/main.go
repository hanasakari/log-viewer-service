package main

import (
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

func main() {
	start()
}

func start() {
	router := gin.New()
	log.Println("Serving at localhost:8000...")
	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)

	}
}
