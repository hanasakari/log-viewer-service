package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
	"net/http"
)

// up grader http to websocket
var upGrader = websocket.Upgrader{
	// R/W buffer size
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartConnectHandler() {

}

func ErrorConnectHandler() {

}

func DestroyConnectHandler() {

}

func TailHandler() {

}

func GetMessage(c *gin.Context) {
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
		config := tail.Config{MustExist: false, Follow: true}
		// filepathin message
		//t, _ := tail.TailFile(string(message), config)
		t, _ := tail.TailFile("../log-viewer-service/test/resource/test.txt", config)

		if string(message) == "on" {
			for line := range t.Lines {
				//fmt.Println(time.Now())
				ws.WriteMessage(mt, []byte(line.Text))
				//fmt.Println(line.Text)
			}
			err = ws.WriteMessage(mt, message)
			if err != nil {
				break
			}
		} else {
			//t.Stop()
		}

	}
}
