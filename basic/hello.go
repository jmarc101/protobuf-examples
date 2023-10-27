package basic

import (
    "grpc-go/protogen/basic"
    "log"
)

func BasicHello(){
    h := basic.Hello{
        Name: "Jean-Marc",
    }

    log.Println("Hello,", h.Name)
}

