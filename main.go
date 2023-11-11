package main

import (
	"fmt"
	"grpc-go/basic"
	// "grpc-go/protogen/basic"
	"log"
	"time"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// basic.BasicHello()
	// basic.BasicUser()
	// basic.BasicUserGroup()
	// basic.AnyCommunicationChannel()
	// basic.ToProtoBinary()
	// basic.FromProtoBinary()
	// basic.ToProtoJsonFile()
	// basic.FromJsonFileToProto()
	basic.DownloadedGoogleProtoTypes()
}
