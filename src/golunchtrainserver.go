package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"lunchtrainapi"
	"net/http"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(rest.Get("/lunchtrain", lunchtrainapi.GetLunchTrain),
		rest.Post("/lunchtrain/:place", lunchtrainapi.AddPlace),
		rest.Post("/lunchtrain/:place/:person", lunchtrainapi.AddPersonToPlace),
		rest.Get("/lunchtrainhistory", lunchtrainapi.GetLunchTrainHistory),
		rest.Get("/lunchtrainhistory/:date", lunchtrainapi.GetLunchTrainForDate))

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
