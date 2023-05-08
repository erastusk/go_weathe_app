package api

import (
	"bytes"
	"encoding/json"
	"github/erastusk/go-weather-map/types"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleResponse(t *testing.T) {
	out := types.OpenWeatherStruct{}
	jsonFile, err := ioutil.ReadFile("../data/tampa.json")
	if err != nil {
		t.Fatalf("Unable to read Json: %+v", err)
	}
	reader := bytes.NewReader(jsonFile)
	err = json.NewDecoder(reader).Decode(&out)
	if err != nil {
		t.Fatalf("Unable to read Json: %+v", err)
	}
	refltype := reflect.TypeOf(out)
	refval := reflect.ValueOf(out)
	for i := 0; i < refltype.NumField(); i++ {
		switch refltype.Field(i).Type.Kind() {
		case reflect.Struct:
			for a := 0; a < refval.Field(i).NumField(); a++ {
				if refltype.Field(i).Type.Name() == "Main" {
					switch val := refval.Field(i).Type().Field(a).Name; val {
					case "FeelsLike":
						assert.EqualValues(t, "FeelsLike", val, "Expected FeelsLike")
					case "Humidity":
						assert.EqualValues(t, "Humidity", val, "Expected Humidity")
					case "Pressure":
						assert.EqualValues(t, "Pressure", val, "Expected Pressure")
					case "Temp":
						assert.EqualValues(t, "Temp", val, "Expected Temp")
					case "TempMax":
						assert.EqualValues(t, "TempMax", val, "Expected TempMax")
					case "TempMin":
						assert.EqualValues(t, "TempMin", val, "Expected TempMin")
					}
				}
			}
		case reflect.Slice:
			for b := 0; b < refval.Field(i).Len(); b++ {
				if refval.Field(i).Index(b).Kind() == reflect.Struct {
					for c := 0; c < refval.Field(i).Index(b).NumField(); c++ {
						if refval.Field(i).Index(b).Type().Field(c).Name == "Description" {
							assert.EqualValues(t, "Description", refval.Field(i).Index(b).Type().Field(c).Name, "Expected Description")
						}

					}
				}
			}
		case reflect.String:
			if refltype.Field(i).Name == "Name" {
				assert.EqualValues(t, "Name", refltype.Field(i).Name, "Expected Name")
			}
		}

	}
}
