package basic

import (
    "google.golang.org/protobuf/encoding/protojson"
    "grpc-go/protogen/basic"
    "log"
)

func BasicUser(){
    user := &basic.User{
        Id: 1,
        Username: "Clark kent",
        IsActive: true,
        Password: []byte("123456"),
        // Super original emails!!
        Emails: []string{"foo@gmail.com", "bar@icloud.com"},
        // Some comment here
        Gender: basic.Gender_GENDER_MALE,
        Address: &basic.Address{Street: "Foo street", City: "Bar", Country: "Some Country"},
    }  

    log.Println(user)
}

func ProtoToJsonUser(){
    user := &basic.User{
        Id: 99,
        Username: "Cat women",
        IsActive: true,
        Password: []byte("654321"),
        Emails: []string{"Email@email.com"},
        Gender: basic.Gender_GENDER_FEMALE,
    }

    bytesProto, _ := protojson.Marshal(user)
    jsonUser := string(bytesProto)

    log.Println(jsonUser)
}

func JsontoProto(){
    user := `{
        "id":99,
        "username":"Cat women",
        "is_active":true,
        "password":"NjU0MzIx",
        "gender":"GENDER_FEMALE",
        "emails":["Email@email.com"]
    }`

    bytes := []byte(user)
    protoUser := &basic.User{}

    err :=protojson.Unmarshal(bytes, protoUser)
    if err != nil {
        log.Println(err)
    }

    log.Println(protoUser)
}
