package main

import (
  "time"
  "fmt"
  "github.com/vodolaz095/goweather"
)

func main(){
//asyncronous call
  var c chan goweather.Weather = make(chan goweather.Weather)
  cities:= [3]string{
    "London",
    "Moscow",
    "Klin",
  }

  for i:=0; i<len(cities); i++ {
    goweather.GetWeatherAsync(cities[i], c)
    time.Sleep(100 * time.Millisecond)
    goweather.PrintWeather(<- c)
  }

//syncronous call
  w, err:=goweather.GetWeather("London")
  if err != nil {
    fmt.Println("Some error", err)
  }
  goweather.PrintWeather(w)
}