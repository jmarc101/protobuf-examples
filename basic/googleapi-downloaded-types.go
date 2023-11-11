package basic

import (
	"grpc-go/protogen/basic"
	"log"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DownloadedGoogleProtoTypes() {
	birthday := &date.Date{
		Year: 2000,
		Month: 01,
		Day: 01,
	}

	position := &latlng.LatLng{
		Latitude: 10.22233,
		Longitude: 100.222,
	}

	user := &basic.User{
		Username: "Googleapi user",
		LastLoginAttempt: timestamppb.Now(),
		Birthday: birthday,
		Position: position,
	}

	log.Println(user)
}
