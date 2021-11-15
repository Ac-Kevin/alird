package main

import (
	"app/aliclient"
	"app/config"
	l "app/log"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// 异常捕获
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	l.Init()
	runServer()
}

func runServer() {

	for {
		sleeptime := os.Getenv("INTERVAL_TIME")
		sleep, _ := strconv.Atoi(sleeptime)
		if sleep < 30 {
			sleep = 30
		}
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Error : %v", err)
				}
			}()
			cfg, err := config.NewConfig()
			if err != nil {

				panic(err)
			}
			client, err := aliclient.NewAliClient()
			if err != nil {
				panic(err)
			}
			run(cfg, client)
		}()
		time.Sleep(time.Duration(sleep) * time.Second)
	}

}
