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
        fmt.Println("Location and API key need to be provided.")
        return
    }else if len(os.Args) < 3 {
        fmt.Println("API key needs to be provided.")
        return
    }

    location := arguments[0]
    apiKey := arguments[1]

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

    fmt.Println(time.Unix(weatherData.Sys.Sunset,0).String())
}
