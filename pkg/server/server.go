package server

import (
	"fmt"
	"github.com/samqintw/logparser/pkg/contract"
	"github.com/samqintw/logparser/pkg/utils"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

type ParserLog struct{}

func (h *ParserLog) Exec(args *contract.HealthCheckRequest, reply *contract.HealthCheckResponse) error {
	log.Println("Exec")
	log.Println("file size", len(args.Log), "bytes")
	fmt.Println("path:", utils.GetLogPath())
	if path := utils.GetLogPath(); len(path) != 0 {
		if fi, err := os.Stat(path); !os.IsNotExist(err) && fi.IsDir() {
			file := fmt.Sprintf("%s/%s", path, time.Now().Format("2006-01-02_15:04:05"))
			log.Println("Saving at", file)
			err := ioutil.WriteFile(file, args.Log, 0644)
			if err != nil {
				log.Fatal("Writing files", err)
			}
		} else {
			log.Println("File didn't be saved", os.IsNotExist(err), fi.IsDir(), fi.Mode())
		}
	}

	mail := utils.Mail{Receivers: []string{"somin_chin@asus.com"}, Subject: "Health Checks", Text: "Health Checks"}
	reply.Message = utils.SendMail(mail)
	return nil
}

func StartServer(port string) {
	parserLog := &ParserLog{}
	rpc.Register(parserLog)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	log.Printf("Server starting on port %v\n", port)

	http.Serve(l, nil)
}