package main

import (
	"github/erastusk/go-weather-map/api"
	"github/erastusk/go-weather-map/inits"
)

func init() {
	inits.LoadVars()
}

func main() {
	logger := inits.NewLogger()
	//app, err := api.CallWeatherApp(logger)
	p := api.NewGetWeather("tampa", 10)
	r, c, err := p.CallWeatherApp(logger)
	if err != nil {
		logger.Logger.Println(err)
	}
	if c != 200 {
		logger.Logger.Fatalf("Non 200 response code, wait and try again.....")
	}
	resp := api.HandleResponse(logger, r)
	logger.Logger.Printf("%v", string(resp))
}
