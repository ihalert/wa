package main

import (
	"fmt"
	"github.com/mivan/weatherApp/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	//var conn *grpc.ClientConn
	var cityName, countryCode string

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("First ---- Server needs to fill its cache, after that enter name and code ----")
	fmt.Println("\nEnter city name AND Enter country code - us for usa, uk for united kingdom etc. AND")
	fmt.Scanf("%s%s", &cityName, &countryCode)

	c := api.NewWeatherAppServiceClient(conn)

	response, err := c.GetWeatherForecast(context.Background(), &api.CityNameCountryCodeRequest{CityName: cityName, CountryCode: countryCode})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("--- Response from server: ---\n*******\n", response.Text)
}
