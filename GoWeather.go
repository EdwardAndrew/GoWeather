package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    //"os"
    "encoding/json"
    "time"
    "flag"
    )

func main() {

    apiKey := flag.String("key", "REQUIRED", "OpenWeatherMap Current Weather API key")
    location := flag.String("l", "London", "Location")
    //TODO System for parsing tokens in tail
    flag.Parse()

    if *apiKey == "" {
        fmt.Println("API Key needs to be provided.")
        return
    }

    URL := "http://api.openweathermap.org/data/2.5/weather?q="+*location+"&APPID="+*apiKey

    var myClient = &http.Client{Timeout: 10 * time.Second}
    response, httpErr := myClient.Get(URL)
    if httpErr != nil {
        fmt.Println(httpErr.Error())
    }
    defer response.Body.Close()

    bodyBytes, ioErr := ioutil.ReadAll(response.Body)
    if ioErr != nil{
        fmt.Println(ioErr.Error())
    }

    weatherData := new(OpenWeatherMapWeatherResponse)

    jsonErr := json.Unmarshal(bodyBytes, weatherData)
    if jsonErr != nil {
        fmt.Println(jsonErr.Error())
    }
    //TODO If no data requested return Temperature
    if len(flag.Args()) <= 0 {
            fmt.Println(KelvinToCelsius(weatherData.Main.Temp))
    }
}
