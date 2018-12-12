package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func InCluster() bool {
	host, port := os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("KUBERNETES_SERVICE_PORT")
	if len(host) == 0 || len(port) == 0 {
		return false
	}
	return true
}

const EmailServerInCluster = "http://mail-sender.monitoring.svc:8803/mails/v1"
const EmailServer = "http://10.78.26.19:30883/mails/v1"

type Mail struct {
	Receivers []string `json:"receivers"`
	Subject string `json:"subject"`
	Text string `json:"test"`
}

func SendMail(mail Mail) string {
	log.Println("Mailing ...")
	var uri string
	if InCluster() {
		uri = EmailServerInCluster
	} else {
		uri = EmailServer
	}
	body, err := json.Marshal(mail)
	if err != nil {
		fmt.Println("Marshal ", err)
	}

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err.Error()
	} else {
		return resp.Status
	}
}

func GetLogPath() string {
	return os.Getenv("LOG_PATH")
}