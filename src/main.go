package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log-viewer/src/websocket"
)

// setting port default 8000
var port = flag.Int("p", 8000, "port is number please use int parameter")

func main() {
	flag.Parse()
	fmt.Println("-port:", *port)
	start(*port)
}

func start(port int) {
	address := fmt.Sprintf("%s:%d", "localhost", port)
	router := gin.New()
	router.GET("/", websocket.GetMessage)
	router.Run(fmt.Sprintf(":%d", port))
	log.Println("Serving at " + address)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
