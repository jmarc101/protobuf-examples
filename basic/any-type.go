package basic

import (
	"grpc-go/protogen/basic"
	"log"
	"google.golang.org/protobuf/types/known/anypb"
)

func logAnyUnderlyingType(user *basic.User) {
	m, err := user.CommunicationChannel.UnmarshalNew()
	if err != nil {
		log.Fatal(err)
	}

	switch m.(type) {
	case *basic.InstantMessaging:
		log.Printf("Username: %s has a communication channel of type InstantMessaging \n%s", user.Username, m)
	case *basic.PaperMail:
		log.Printf("Username: %s has a communication channel of type PaperMail \n%s", user.Username, m)
	case *basic.SocialMedia:
		log.Printf("Username: %s has a communication channel of type SocialMedia \n%s", user.Username, m)
	}
}

func AnyCommunicationChannel() {
	im := &basic.InstantMessaging{
		PhoneNumber: "555-5555",
	}
	sm := &basic.SocialMedia{
		Platform: "X",
	}
	pm := &basic.PaperMail{
		Adress: "Cool place",
	}

	anyIm, err := anypb.New(im)
	if err != nil {
		log.Fatal(err)
	}

	anySm, err := anypb.New(sm)
	if err != nil {
		log.Fatal(err)
	}

	anyPm, err := anypb.New(pm)
	if err != nil {
		log.Fatal(err)
	}

	imUser := &basic.User{
		Username: "im user",
		CommunicationChannel: anyIm,
	}

	smUser := &basic.User{
		Username: "sm user",
		CommunicationChannel: anySm,
	}

	pmUser := &basic.User{
		Username: "pm user",
		CommunicationChannel: anyPm,
	}

	logAnyUnderlyingType(imUser)
	logAnyUnderlyingType(smUser)
	logAnyUnderlyingType(pmUser)
}
