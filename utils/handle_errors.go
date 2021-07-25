package utils

import (
	"fmt"
	"log"
)

func HandleErrors(err error, logMessage string) {
	fmt.Println(err)
	if err != nil {
		log.Fatalln(logMessage)
	}
}


