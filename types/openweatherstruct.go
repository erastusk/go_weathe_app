package types

type Clouds struct {
	All int `json:"all"`
}
type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type Main struct {
	FeelsLike float64 `json:"feels_like"`
	Humidity  int     `json:"humidity"`
	Pressure  int     `json:"pressure"`
	Temp      float64 `json:"temp"`
	TempMax   float64 `json:"temp_max"`
	TempMin   float64 `json:"temp_min"`
}
type Resp struct {
	City        string  `json:"city"`
	Description string  `json:"description"`
	FeelsLike   float64 `json:"feels_like"`
	Humidity    int     `json:"humidity"`
	Pressure    int     `json:"pressure"`
	Temp        float64 `json:"temp"`
	TempMax     float64 `json:"temp_max"`
	TempMin     float64 `json:"temp_min"`
}
type Sys struct {
	Country string `json:"country"`
	ID      int    `json:"id"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
	Type    int    `json:"type"`
}
type Weather []struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
	ID          int    `json:"id"`
	Main        string `json:"main"`
}
type Wind struct {
	Deg   int     `json:"deg"`
	Speed float64 `json:"speed"`
}
type OpenWeatherStruct struct {
	Base       string  `json:"base"`
	Clouds     Clouds  `json:"clouds"`
	Cod        int     `json:"cod"`
	Coord      Coord   `json:"coord"`
	Dt         int     `json:"dt"`
	ID         int     `json:"id"`
	Main       Main    `json:"main"`
	Name       string  `json:"name"`
	Sys        Sys     `json:"sys"`
	Timezone   int     `json:"timezone"`
	Visibility int     `json:"visibility"`
	Weather    Weather `json:"weather"`
	Wind       Wind    `json:"wind"`
}
