package basic

import (
	"grpc-go/protogen/basic"
	"log"
)

func MapsProto() {
	skillsMap := make(map[string]uint32)
	skillsMap["Go"] = 10
	skillsMap["protobuf"] = 6

	user := &basic.User{
		Username: "Bob",
		Skills: skillsMap,
	}

	log.Println(user)
}
