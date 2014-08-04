Package to parse weather API
=====================================

We simply get JSON data from this API endpoint and parse them
[http://api.openweathermap.org/data/2.5/weather?q=Moscow](http://api.openweathermap.org/data/2.5/weather?q=Moscow)



Usage
=====================================

```go

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

```


Testing
=====================================

```shell

    $ go test

```


Time taken
====================================

53 minutes