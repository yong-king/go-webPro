package log

import (
	"log"
	"net/http"
	"os"
)

func SetupLogger() {
	logFile, _ := os.OpenFile("./testNorm.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFile)
}

func SimpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("http.Get(url) failed, err:%v", err)
		return
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}
