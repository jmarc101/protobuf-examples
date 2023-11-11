package basic

import (
	"grpc-go/protogen/basic"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
)

func ToProtoJsonFile(){
	user := &basic.User{
		Id: 9000,
		Username: "Json User",
		Gender: basic.Gender_GENDER_MALE,
	}

	jsonProto, err := protojson.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	os.WriteFile("files/user.json", []byte(jsonProto), 0666)
}

func FromJsonFileToProto(){
	bytes, err := os.ReadFile("files/user.json")
	if err != nil {
		log.Fatalln(err)
	}

	user := &basic.User{}
	if err := protojson.Unmarshal(bytes, user); err != nil {
		log.Fatalln(err)
	}

	log.Println(user)
}
