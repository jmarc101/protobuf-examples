// Package basic demonstrates the usage of packages to prevent name collisions.
package basic

import (
	"grpc-go/protogen/basic"     // Import the main package for the application
	"grpc-go/protogen/basic/first" // Import the first package for job-related details
	"grpc-go/protogen/basic/second" // Import the second package for software-related details
	"log"
)

// Application is a function that demonstrates the usage of packages to prevent name collisions.
func Application() {
	// Create an instance of the main Application from the 'basic' package
	app := &basic.Application{
		// Initialize the JobApplication field with an instance from the 'first' package
		JobApplication: &first.Application{
			ApplicantName: "Jean-Marc P.", // Set the applicant's name
		},
		// Initialize the SoftwareApplication field with an instance from the 'second' package
		SoftwareApplication: &second.Application{
			SoftwareTitle: "Emacs", // Set the software title
			Os:            "it's own", // Set the operating system
		},
	}

	// Log the created application instance
	log.Println(app)
}
