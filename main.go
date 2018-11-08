package main

import (
	"log"
	"net/http"
	
	"github.com/janaonline/icmyc/config"
	"github.com/janaonline/icmyc/routers"
	"github.com/urfave/negroni"
	"github.com/rs/cors"

	_ "github.com/sirupsen/logrus"
)

func main()  {
	config.Init()

	c := cors.New(cors.Options{})

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(config.Get().ServerPort, n))

}