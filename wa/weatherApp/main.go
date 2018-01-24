package main

import (
	"encoding/json"
	"fmt"
	"github.com/mivan/weatherApp/api"
	"github.com/mivan/weatherApp/entities"
	//	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	//"google.golang.org/grpc"
	//"net"
	//"github.com/mivan/weatherApp/store"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	//"github.com/go-sql-driver/mysql"
	"github.com/mivan/weatherApp/server"
)

//var M map[string]entities.DBStruct

func main() {

	var country, city string
	var myJSONstruct = entities.JsonStruct{}
	var structMap = entities.DBStruct{}

	m := make(map[string]entities.DBStruct)

	fmt.Println("\n\nThis is weather app, follow instructions")
	for i := 0; i < 3; i++ {

		fmt.Println("\nEnter country code , us for usa, uk for united kingdom etc.")
		fmt.Scanf("%s", &country)

		for len(country) > 3 {

			fmt.Println("\nEnter country code , us for usa, uk for united kingdom etc.")
			fmt.Scanf("%s", &country)
		}

		fmt.Println("\nEnter city name")
		fmt.Scanf("%s", &city)

		key := city + "," + country
		key = strings.ToLower(key)

		if _, ok := m[key]; ok {

			fmt.Println("\n\n***It is in local map***")
			fmt.Println(m[key])

		} else {

			url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "," + country + "&appid=fad7dc4ab0d7660a024b3f7a2534665e"

			resp, err := http.Get(url)

			if err != nil {

				fmt.Println("Neka greška")

			} else if resp.StatusCode >= 200 && resp.StatusCode < 300 {

				body, _ := ioutil.ReadAll(resp.Body)

				err1 := json.Unmarshal(body, &myJSONstruct)
				if err1 != nil {
					fmt.Println(err1)
				}

				str := "\nWeather Today : " + string(myJSONstruct.Weather[0].Description) + "\nTemp min : " + strconv.FormatFloat(myJSONstruct.Main.TempMin, 'f', -1, 64) + "\nTemp max : " + strconv.FormatFloat(myJSONstruct.Main.TempMax, 'f', -1, 64) + "\nWind : " + strconv.FormatFloat(myJSONstruct.Wind.Speed, 'f', -1, 64)

				//structMap := entities.DBStruct{}
				structMap.CityName = city
				structMap.CountryCode = country
				structMap.Text = str

				//m := make(map[string]entities.DBStruct)
				m[key] = structMap

				for k, v := range m {

					fmt.Println("\n\n", k, v)

				}

			}
		}
	}

	// handler nekakav

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatal(err)
	}

	//type Server struct{}
	/*
		func (t *Server) GetWeatherForecast(ctx context.Context, in * api.CityNameCityCodeRequest) (*api.ForecastResponse, error){


			// pogledaj za taj unos da li postoji prognoza i vrati ju van

			return &api.ForecastResponse{Text:"Neki tekst prognoze"}, nil
		}
	*/
	s := server.Server{}
	grpcServer := grpc.NewServer()
	server.MyMap = m

	//api.RegisterWeatherAppServiceServer(grpcServer, &s)

	// neki grpcServer.Serve() ...

	api.RegisterWeatherAppServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)

	}

	for i := 0; i < 5; i++ {
		fmt.Println("PRINT ----------", i)
	}

}
