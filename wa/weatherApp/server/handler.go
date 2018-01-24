package server

import (
	"fmt"
	"github.com/mivan/weatherApp/api"
	"github.com/mivan/weatherApp/entities"
	"golang.org/x/net/context"
	//"log"
	"strings"
)

var MyMap map[string]entities.DBStruct

type Server struct {
}

func (t *Server) GetWeatherForecast(ctx context.Context, in *api.CityNameCountryCodeRequest) (*api.ForecastResponse, error) {

	// look up for forecast and give it back

	key := in.CityName + "," + in.CountryCode
	strings.ToLower(key)

	fmt.Println("\n\nPRINT FROM SERVER")

	var text string
	if _, ok := MyMap[key]; ok {

		fmt.Println("FOUND")
		fmt.Println(MyMap[key])
		text = MyMap[key].Text

	} else {

		fmt.Println("NOT FOUND")
		text = "NOT FOUND"

	}

	return &api.ForecastResponse{Text: text}, nil
}
