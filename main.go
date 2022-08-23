package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"test/database"
	"test/service"
	"test/service/config"
	"test/service/subscriber"
)

func main() {

	var wg sync.WaitGroup
	config.InfoConfigParams()
	wg.Add(1)
	database.ConnectionDB(&wg)
	wg.Wait()
	// go database.CheckDB()

	sc, sub := subscriber.Subscribe()
	go service.StartServer()
	// go func() {
	// 	for {
	// 		time.Sleep(3 * time.Second)
	// 		cache.CheckCacheKey()
	// 	}
	// }()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Println("Server close")
	subscriber.NatsClose(*sc, *sub)
	database.DisconnectionDB()

}
