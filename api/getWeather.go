package api

import (
	"context"
	"github/erastusk/go-weather-map/inits"
	"net/http"
	"os"
	"time"
)

// url = "https://api.openweathermap.org/data/2.5/weather?q={city_name}&units=imperial&appid=abf5cde6ff725334c4c2dab475a4ba51"
type GetWeather struct {
	ctx     context.Context
	city    string
	timeout int
	url     string
}

func NewGetWeather(c string, t int) *GetWeather {
	str1 := "https://api.openweathermap.org/data/2.5/weather?q="
	str2 := "&units=imperial&appid="
	url := str1 + c + str2 + os.Getenv("API_KEY")
	ctx := context.Background()
	return &GetWeather{
		ctx:     ctx,
		city:    c,
		timeout: t,
		url:     url,
	}
}

func (g *GetWeather) CallWeatherApp(l *inits.Logging) (*http.Response, int, error) {
	timeout := time.Duration(time.Second * time.Duration(g.timeout))
	ctx, cancel := context.WithTimeout(g.ctx, timeout)
	defer cancel()
	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, g.url, nil)
	if err != nil {
		l.Logger.Println(err)
		return nil, 0, err
	}
	client := inits.InitHttpClient()

	res, err := client.Do(resp)
	if err != nil {
		return nil, 0, err
	}

	l.Logger.Println(res.StatusCode)
	return res, res.StatusCode, nil

}
