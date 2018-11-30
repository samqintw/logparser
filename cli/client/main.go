package main

import (
	"flag"
	"fmt"
	"github.com/samqintw/logparser/pkg/client"
	"os"
)

const IP = "localhost"
const PORT = "2148"

//var (
//	file = flag.String("file", "", "attachment")
//	ip = flag.String( "ip", IP, "ip of the log parser service")
//	port =flag.String( "addr", PORT, "port of the log parser service")
//)

func main()  {
	var ip, port, file string
	flag.StringVar(&ip, "ip", IP, "ip of the log parser service")
	flag.StringVar(&port, "addr", PORT, "port of the log parser service")
	flag.StringVar(&file, "file", "", "attachment")
	flag.Parse()

	if fi, err := os.Stat(file); !os.IsNotExist(err) && !fi.IsDir() {
		fmt.Println(ip, port, file)
		c := client.CreateClient(fmt.Sprintf("%v:%v", ip, port))
		defer c.Close()
		reply := client.SendLog(c, file)
		fmt.Println(reply.Message)
	} else {
		flag.Usage()
		fmt.Println(`ex: client --ip="" --port="" --file=""`)
		fmt.Println(os.IsNotExist(err))
		fmt.Println(fi.IsDir())
	}
}