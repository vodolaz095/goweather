package goweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherCoordSt struct {
	Lon float64
	Lat float64
}

type WeatherMainSt struct {
	Temp     float64
	Temp_min float64
	Tem_max  float64
	Pressure int
	Humidity int
}

//"sys": {"type":1,"id":5091,"message":0.1975,"country":"GB","sunrise":1407040087,"sunset":1407095111}
type WeatherSysSt struct {
	Type    int
	Id      int
	Message float64
	Country string
	Sunrise int
	Sunset  int
}

type WeatherWindSt struct {
	Speed float64
	Deg   int
}

type Weather struct {
	Name  string
	Cod   int
	Dt    int
	Id    int
	Coord WeatherCoordSt
	Main  WeatherMainSt
	Sys   WeatherSysSt
	Wind  WeatherWindSt
}

func GetWeather(city string) (w Weather, err error) {

	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("Error getting data", err)
		return Weather{}, err
	} else {
		rawWeather, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading data", err)
			return Weather{}, err
		} else {
			//var rawWeather = []byte(`{"coord":{"lon":-0.13,"lat":51.51},"sys": {"type":1,"id":5091,"message":0.1975,"country":"GB","sunrise":1407040087,"sunset":1407095111},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"base":"cmc stations","main": {"temp":290.52,"pressure":1012,"humidity":55,"temp_min":289.15,"temp_max":292.15},"wind": {"speed":4.6,"deg":230},"clouds":{ "all":36 },"dt":1407100800,"id":2643743,"name":"London","cod":200}`)
			err = json.Unmarshal(rawWeather, &w)
			//      fmt.Println("Parsing", w)
			return w, err
		}
	}
}

func GetWeatherAsync(city string, c chan Weather) {
	go func() {
		w, err := GetWeather(city)
		if err != nil {
			c <- Weather{}
		} else {
			c <- w
		}
		return
	}()
}

func PrintWeather(w Weather) {
	//  fmt.Printf("%+v\n", w)

	fmt.Println(
		"The weather in city of ",
		w.Name,
		" is following: \nTemperature: ",
		int(w.Main.Temp-273.0),
		"C' \nPressure: ",
		int(w.Main.Pressure),
		"mm Hg\nHumidity: ",
		int(w.Main.Humidity),
		"%\nHave a nice day!\n\n")
}
