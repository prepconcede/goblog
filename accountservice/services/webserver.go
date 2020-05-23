package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)
	log.Println("starting http services at " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Println("an error occurred starting http listener at port " + port)
		log.Println("error: " + err.Error())
	}
}
