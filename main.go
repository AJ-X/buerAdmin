package main

import (
	"buerAdmin/model"
	"buerAdmin/router"
	"log"
)

func main() {
	model.InitDb()
	router.InitRouter()
	defer func() {
		if model.Conn!=nil {
			err := model.Conn.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
}
