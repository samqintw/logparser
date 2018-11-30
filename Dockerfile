# build stage for golang
FROM golang:1.10 AS golang_env
MAINTAINER sam.qin.tw@gmail.com

ARG app_name=logparser
ARG app_package=github.com/samqintw/logparser
ADD pkg $GOPATH/src/$app_package/pkg
ADD cli $GOPATH/src/$app_package/cli

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $GOPATH/bin/$app_name $app_package/cli/server && \
    chmod +x $GOPATH/bin/$app_name

# final stage
FROM scratch
COPY --from=golang_env /go/bin/logparser /usr/local/bin/
ENTRYPOINT ["logparser"]