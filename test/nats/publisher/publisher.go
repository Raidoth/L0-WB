package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

const (
	ClusterID  = "test-cluster"
	PublicFile = "model.json"
)

func main() {
	sc, err := stan.Connect(ClusterID, "test")
	if err != nil {
		log.Fatalf("Can't connect: %v.\n", err)
	}
	defer sc.Close()
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			data, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				log.Fatalf("Cannot read file %v", err)
			}
			err = sc.Publish("WB", data)
			if err != nil {
				fmt.Println("File no publish")
			} else {
				fmt.Println("File publish")
			}
		}

	} else {
		data, err := ioutil.ReadFile(PublicFile)
		if err != nil {
			log.Fatalf("Cannot read file %v", err)
		}
		err = sc.Publish("WB", data)
		if err != nil {
			fmt.Println("File no publish")
		} else {
			fmt.Println("File publish")
		}
	}

}
