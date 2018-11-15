package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func main() {

	n := howMany()

	for i := 0; i < n; i++ {
		printUUID()
	}
}

func howMany() int {

	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("%s", err)
		}

		return i
	}

	return 1
}

func printUUID() {
	v, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println(v.String())
}
