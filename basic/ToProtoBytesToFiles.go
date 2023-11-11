package basic

import (
	"grpc-go/protogen/basic"
	"log"
	"os"
	"google.golang.org/protobuf/proto"
)

func ToProtoBinary() {
	user := &basic.User {
		Username: "Bob",
		Id: 1,
		Gender: basic.Gender_GENDER_MALE,
	}

	bytes, err := proto.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	if err := os.WriteFile("files/user.bin", bytes, 0666); err != nil{
		log.Fatal(err)
	}
}

func FromProtoBinary() {
	bytes, err := os.ReadFile("files/user.bin")
	if err != nil {
		log.Fatalln(err)
	}

	message := &basic.User{}
	if err := proto.Unmarshal(bytes, message); err != nil{
		log.Fatalln(err)
	}

	log.Println(message)
}
