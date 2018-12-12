package main

import (
	"flag"
	"github.com/samqintw/logparser/pkg/server"
)

const PORT = "2148"

var (
	port = flag.String("addr", PORT, "endpoint of the log parser service")
)

func main() {
	server.StartServer(*port)
}