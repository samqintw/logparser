    docker build -t samqintw/logparser:latest .

    docker stop logparser; docker rm logparser
    docker run -d -p 2148:2148 --name logparser samqintw/logparser
    docker run -d -v /Users/somin_chin/Documents/log:/log -e "LOG_PATH=/log" -p 2148:2148 --name logparser samqintw/logparser

    go build ./cli/client
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cli/client

    go run ./cli/client/main.go -file="./testlog" -mail_addr="somin_chin@asus.com"