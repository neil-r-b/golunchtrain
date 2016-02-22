package lunchtrainapi

import (
	"github.com/ant0ine/go-json-rest/rest"
	"strings"
	"time"
)

var dailyLunchTrains = make(map[string]*LunchTrain)

type FoodPlace struct {
	Name   string   `json:"name"`
	People []string `json:"people,omitempty"`
}

type LunchTrain struct {
	Places []*FoodPlace `json:"places,omitempty"`
	Date   string       `json:"date"`
}

func GetLunchTrain(w rest.ResponseWriter, req *rest.Request) {
	date := getTodaysDate()
	lunchTrain, ok := dailyLunchTrains[date]

	if !ok {
		lunchTrain = NewLunchTrain(date)
		dailyLunchTrains[date] = lunchTrain
	}

	w.WriteJson(lunchTrain)
}

func GetLunchTrainHistory(w rest.ResponseWriter, req *rest.Request) {
	w.WriteJson(dailyLunchTrains)
}

func GetLunchTrainForDate(w rest.ResponseWriter, req *rest.Request) {
	date := req.PathParam("date")
	lunchTrain, ok := dailyLunchTrains[date]

	if !ok {
		lunchTrain = NewLunchTrain(date)
		dailyLunchTrains[date] = lunchTrain
	}

	w.WriteJson(lunchTrain)
}

func AddPlace(w rest.ResponseWriter, req *rest.Request) {
	placeName := req.PathParam("place")
	date := getTodaysDate()

	lunchTrain, ok := dailyLunchTrains[date]
	if !ok {
		lunchTrain = NewLunchTrain(date)
		dailyLunchTrains[date] = lunchTrain
	}

	var foodPlace *FoodPlace
	var foundPlace bool
	for _, foodPlace = range lunchTrain.Places {
		if strings.ToLower(foodPlace.Name) == strings.ToLower(placeName) {
			foundPlace = true
			break
		}
	}

	if !foundPlace {
		foodPlace = NewFoodPlace(placeName)
		lunchTrain.Places = append(lunchTrain.Places, foodPlace)
		dailyLunchTrains[date] = lunchTrain
	}

	w.WriteJson(lunchTrain)
}

func AddPersonToPlace(w rest.ResponseWriter, req *rest.Request) {
	placeName := req.PathParam("place")
	person := req.PathParam("person")
	date := getTodaysDate()

	lunchTrain, ok := dailyLunchTrains[date]
	if !ok {
		lunchTrain = NewLunchTrain(date)
		dailyLunchTrains[date] = lunchTrain
	}

	var foodPlace *FoodPlace
	var foundPlace bool
	for _, foodPlace = range lunchTrain.Places {
		if strings.ToLower(foodPlace.Name) == strings.ToLower(placeName) {
			foundPlace = true
			break
		}
	}

	if !foundPlace {
		foodPlace = NewFoodPlace(placeName)
		lunchTrain.Places = append(lunchTrain.Places, foodPlace)
	}

	var personAlreadyOnTrain bool
	for _, existingPerson := range foodPlace.People {
		if strings.ToLower(existingPerson) == strings.ToLower(person) {
			personAlreadyOnTrain = true
			break
		}
	}

	if !personAlreadyOnTrain {
		foodPlace.People = append(foodPlace.People, person)
	}

	dailyLunchTrains[date] = lunchTrain
	w.WriteJson(lunchTrain)
}

func NewFoodPlace(name string) *FoodPlace {
	foodPlace := FoodPlace{Name: name}

	return &foodPlace
}

func NewLunchTrain(date string) *LunchTrain {
	lunchTrain := LunchTrain{Date: date}

	return &lunchTrain
}

func getTodaysDate() string {
	return time.Now().Format("2006-01-02")
}
