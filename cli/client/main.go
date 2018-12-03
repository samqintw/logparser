package main

import (
	"flag"
	"fmt"
	"github.com/samqintw/logparser/pkg/client"
	"os"
	"strings"
)

const IP = "localhost"
const PORT = "2148"

//var (
//	file = flag.String("file", "", "attachment")
//	ip = flag.String( "ip", IP, "ip of the log parser service")
//	port =flag.String( "addr", PORT, "port of the log parser service")
//)

func main()  {
	var ip, port, file, mailAddrs string
	flag.StringVar(&ip, "ip", IP, "ip of the log parser service")
	flag.StringVar(&port, "addr", PORT, "port of the log parser service")
	flag.StringVar(&file, "file", "", "attachment")
	flag.StringVar(&mailAddrs, "mail_addr", "", "mail addresses, ex: mail01,mail01 ...")
	flag.Parse()

	receivers := strings.Split(mailAddrs,",")
	if len(mailAddrs) > 0 && len(receivers) == 0 {
		flag.Usage()
		fmt.Println(`wrong mail address format, --mail_addr="mail01,mail02"`)
		return
	}

	if fi, err := os.Stat(file); !os.IsNotExist(err) && !fi.IsDir() {
		fmt.Println(ip, port, file)
		c := client.CreateClient(fmt.Sprintf("%v:%v", ip, port))
		defer c.Close()
		reply := client.SendLog(c, file, receivers)
		fmt.Println(reply.Message)
	} else {
		flag.Usage()
		fmt.Println(`ex: client --ip="" --port="" --file="" --mail_addr=""`)
		fmt.Println(os.IsNotExist(err))
		fmt.Println(fi.IsDir())
	}
}