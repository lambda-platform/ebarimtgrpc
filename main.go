package main

import (
	"fmt"
	"github.com/lambda-platform/ebarimtgrpc/config"
	"github.com/lambda-platform/ebarimtgrpc/server"
	"github.com/lambda-platform/ebarimtgrpc/upload"
	"net/http"
)

func main() {

	/*GRPC SERVER*/
	server.StartGRPC()
	http.HandleFunc("/upload", upload.UploadFile)
	fmt.Println("Server is listening on port " + config.HTTPPort)
	http.ListenAndServe(":"+config.HTTPPort, nil)

}
