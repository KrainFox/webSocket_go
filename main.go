package main

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	// "time"

	"golang.org/x/net/websocket"
)

func getDate(ws *websocket.Conn) {
	// currentTime := time.Now()
	// ws.Write([]byte(currentTime.String()))
	// for i := 0; i < 10; i++ {
	// 	ws.Write([]byte(fmt.Sprintf("%d", i)))
	// 	time.Sleep(time.Duration(1000))
	// }
	lenBuf := make([]byte, 5)
	for {
		_, err := ws.Read(lenBuf)
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}

		length, _ := strconv.Atoi(strings.TrimSpace(string(lenBuf)))
		if length > 65536 {
			log.Println("Error: too big length: ", length)
			return
		}

		if length <= 0 {
			log.Println("Empty length: ", length)
			return
		}

		buf := make([]byte, length)
		_, err = ws.Read(buf)

		if err != nil {
			log.Println("Could not read ", length, " bytes: ", err.Error())
			return
		} else {
			ws.Write(buf)
		}

	}
}

func main() {
	log.Println("Server start...")
	http.Handle("/getDateNow", websocket.Handler(getDate))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Crash server! Reason: ", err)
	}
}
