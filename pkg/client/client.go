package client

import (
	"github.com/samqintw/logparser/pkg/contract"
	"io/ioutil"
	"log"
	"net/rpc"
	"path/filepath"
)

func CreateClient(address string) *rpc.Client {
	log.Println("dialing to", address)
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func SendLog(c *rpc.Client, file string, mailAddr []string) contract.HealthCheckResponse {
	log.Println("Sending log")
	path, _  := filepath.Abs(file)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("error:", err)
	}

	args := &contract.HealthCheckRequest{Name: "World", Log: data, MailAddress: mailAddr}
	var reply contract.HealthCheckResponse
	err = c.Call("ParserLog.Exec", args, &reply)

	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}