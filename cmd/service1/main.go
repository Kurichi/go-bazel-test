package main

import (
	"fmt"
	"log"

	"github.com/Kurichi/go-bazel-test/pkg/uuid"
)

func main() {
	uuid, err := uuid.Generate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uuid)
}
