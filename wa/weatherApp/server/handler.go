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

	// pogledaj za taj unos da li postoji prognoza i vrati ju van

	key := in.CityName + "," + in.CountryCode
	strings.ToLower(key)

	fmt.Println("\n\nIspis iz server metode")

	var text string
	if _, ok := MyMap[key]; ok {

		fmt.Println("PRONAĐENO")
		fmt.Println(MyMap[key])
		text = MyMap[key].Text

	} else {

		fmt.Println("NOT FOUND")
		text = "NOT FOUND"

	}
	//log.Printf("Receive message %s", in.Text)

	return &api.ForecastResponse{Text: text}, nil
}
