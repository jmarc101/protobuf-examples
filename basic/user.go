package basic

import (
    "grpc-go/protogen/basic"
    "log"
)

func BasicUser(){
    user := &basic.User{
        Id: 1,
        Username: "Clark kent",
        IsActive: true,
        Password: []byte("123456"),
    }  

    log.Println(user)
}


