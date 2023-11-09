package basic

import (
	"grpc-go/protogen/basic"
	"log"
)

func BasicUserGroup(){
    user1 := &basic.User{
        Id: 1,
        Username: "Clark kent",
        IsActive: true,
        Password: []byte("123456"),
        Emails: []string{"foo@gmail.com", "bar@icloud.com"},
        Gender: basic.Gender_GENDER_MALE,
        Address: &basic.Address{Street: "Foo street", City: "Bar", Country: "Some Country"},
    }

	user2 := &basic.User{
        Id: 2,
        Username: "Clark bent",
        IsActive: false,
        Password: []byte("654321"),
        Emails: []string{"NO_EMAIL@icloud.com"},
        Gender: basic.Gender_GENDER_MALE,
        Address: &basic.Address{Street: "Foo street", City: "Bar", Country: "Some Country"},
    }

	userGroup := &basic.UserGroup{
		GroupId: 1,
		GroupName: "Cool name",
		Roles: []string{"admin", "manager"},
		Users: []*basic.User{
			user1,
			user2,
		},
		Description: "Some rare description!",
	}

	log.Println(userGroup)
}
