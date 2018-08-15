package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "encoding/json"
    "time"
    "flag"
    )

func main() {

    file, _ := os.Open("conf.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    confErr := decoder.Decode(&configuration)
    if confErr != nil {
        fmt.Println(confErr.Error())
    }

    location := flag.String("l", "London", "Location")
    flag.Parse()

    if configuration.Apikey == "" {
        fmt.Println("API Key needs to be provided in conf.json")
        return
    }

    URL := "http://api.openweathermap.org/data/2.5/weather?q="+*location+"&APPID="+configuration.Apikey

    var myClient = &http.Client{Timeout: 10 * time.Second}
    response, httpErr := myClient.Get(URL)
    if httpErr != nil {
        fmt.Println(httpErr.Error())
    }
    defer response.Body.Close()

    if response.StatusCode == 401 {
        fmt.Println("API key is not valid")
        return
    } else if response.StatusCode == 404 {
        fmt.Println(*location, "is not a valid location") 
        return
    } else if (int(response.StatusCode)/100) != 2  {
        fmt.Println("Something went wrong! Is the service running?")
        return
    }

    bodyBytes, ioErr := ioutil.ReadAll(response.Body)
    if ioErr != nil{
        fmt.Println(ioErr.Error())
    }

    weatherData := new(OpenWeatherMapWeatherResponse)

    jsonErr := json.Unmarshal(bodyBytes, weatherData)
    if jsonErr != nil {
        fmt.Println(jsonErr.Error())
    }

    // If no data is explicitly requested, output the temperature in Celsius.
    if len(flag.Args()) <= 0 {
            fmt.Println(KelvinToCelsius(weatherData.Main.Temp))
    } else {
    // Output requested data in the order it was requested.
            var tail []string = flag.Args()
            for i := 0; i < len(tail); i++ {
                    OutputWeatherData(weatherData, tail[i])
            }
    }
}
