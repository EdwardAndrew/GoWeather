package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
    "time"
    )

func main() {

    arguments := os.Args[1:]

    if len(os.Args) < 2 {
        fmt.Println("API and Location key need to be provided.")
        return
    }else if len(os.Args) < 3 {
        fmt.Println("Location key needs to be provided.")
        return
    }

    dataRequested := len(os.Args) > 3

    apiKey := arguments[0]
    location := arguments[1]

    URL := "http://api.openweathermap.org/data/2.5/weather?q="+location+"&APPID="+apiKey

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
    if !dataRequested {
        // If no weather token is given just provide the temperature.
        fmt.Println(KelvinToCelsius(weatherData.Main.Temp))
    }
}
