package legacy

import (
	"log"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}