package common

import (
	"log"
)

// CheckError To check the error information
func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
}