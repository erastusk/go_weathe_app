package api

import (
	"encoding/json"
	"github/erastusk/go-weather-map/inits"
	"github/erastusk/go-weather-map/types"
	"net/http"
	"reflect"
)

func HandleResponse(l *inits.Logging, resp *http.Response) []byte {
	out := types.OpenWeatherStruct{}
	mainresp := types.Resp{}
	defer resp.Body.Close()
	// jsonFile, err := ioutil.ReadFile("data/tampa.json")
	// if err != nil {
	// 	l.Logger.Printf("Unable to read Json: %+v", err)
	// }
	// err := json.Unmarshal(resp.Body, &out)
	// if err != nil {
	// 	l.Logger.Printf("Unable to Unmarshal json file %+v", err)
	// }
	err := json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		l.Logger.Printf("Unable to read Json: %+v", err)
	}
	refltype := reflect.TypeOf(out)
	refval := reflect.ValueOf(out)
	for i := 0; i < refltype.NumField(); i++ {
		//l.Logger.Println(refltype.Field(i).Type.Kind())
		switch refltype.Field(i).Type.Kind() {
		case reflect.Struct:
			for a := 0; a < refval.Field(i).NumField(); a++ {
				if refltype.Field(i).Type.Name() == "Main" {
					//l.Logger.Printf("Struct Name--> %v ----type %v---  FieldName--> %v ----Field Value--> %v", refltype.Field(i).Type.Name(), refltype.Field(i).Type.Field(a).Type, refval.Field(i).Type().Field(a).Name, refval.Field(i).Field(a).Interface())
					switch val := refval.Field(i).Type().Field(a).Name; val {
					case "FeelsLike":
						mainresp.FeelsLike = refval.Field(i).Field(a).Float()
					case "Humidity":
						mainresp.Humidity = int(refval.Field(i).Field(a).Int())
					case "Pressure":
						mainresp.Pressure = int(refval.Field(i).Field(a).Int())
					case "Temp":
						mainresp.FeelsLike = refval.Field(i).Field(a).Float()
					case "TempMax":
						mainresp.FeelsLike = refval.Field(i).Field(a).Float()
					case "TempMin":
						mainresp.FeelsLike = refval.Field(i).Field(a).Float()
					}
				}
			}
		case reflect.Slice:
			//l.Logger.Println(refltype.Field(i).Type.Kind(), refval.Field(i).Len())
			for b := 0; b < refval.Field(i).Len(); b++ {
				//l.Logger.Println(refval.Field(i).Index(b).Interface())
				if refval.Field(i).Index(b).Kind() == reflect.Struct {
					for c := 0; c < refval.Field(i).Index(b).NumField(); c++ {
						//l.Logger.Println(refltype.Field(i).Type.Name(), refval.Field(i).Index(b).Type().Field(c).Name, refval.Field(i).Index(b).Field(c))
						if refval.Field(i).Index(b).Type().Field(c).Name == "Description" {
							//l.Logger.Println(refval.Field(i).Index(b).Field(c).String())
							mainresp.Description = refval.Field(i).Index(b).Field(c).String()
						}

					}
				}
			}
		case reflect.String:
			if refltype.Field(i).Name == "Name" {
				mainresp.City = refval.Field(i).String()
			}
		}

	}
	r, err := json.Marshal(mainresp)
	if err != nil {
		l.Logger.Println(err)
	}

	return r
}
