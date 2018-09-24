package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	t := time.Now().UTC()

	if len(os.Args) > 1 {
		d, err := time.ParseDuration(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		t = t.Add(d)
	}

	fmt.Println(t.UTC().Format(time.RFC3339))
}
