package main

import(
    "fmt"
    "strings"
    )

func OutputWeatherData(weatherData *OpenWeatherMapWeatherResponse, requestedData string){
        switch strings.ToLower(requestedData) {
        case "longitude", "lon":
            fmt.Println(requestedData, ":", weatherData.Coord.Lon)
        case "latitude", "lat":
            fmt.Println(requestedData, ":", weatherData.Coord.Lat)
        case "weatherid", "wid":
            fmt.Println(requestedData, ":", weatherData.Weather[0].Id)
        case "weatherdesc", "wd":
            fmt.Println(requestedData, ":", weatherData.Weather[0].Description)
        case "weathericon", "wi":
            fmt.Println(requestedData, ":", weatherData.Weather[0].Icon)
        case "base":
            fmt.Println(requestedData, ":", weatherData.Base)
        case "tempc", "tc":
            fmt.Println(requestedData, ":", KelvinToCelsius(weatherData.Main.Temp))
        case "tempf", "tf":
            fmt.Println(requestedData, ":", KelvinToFahrenheit(weatherData.Main.Temp))
        case "tempk", "tk":
            fmt.Println(requestedData, ":", weatherData.Main.Temp)
        case "pressure", "p":
            fmt.Println(requestedData, ":", weatherData.Main.Pressure)
        case "humidity", "h":
            fmt.Println(requestedData, ":", weatherData.Main.Humidity)
        case "tempminc", "tminc":
            fmt.Println(requestedData, ":", KelvinToCelsius(weatherData.Main.Temp))
        case "tempminf", "tminf":
            fmt.Println(requestedData, ":", KelvinToFahrenheit(weatherData.Main.Temp))
        case "tempmink", "tmink":
            fmt.Println(requestedData, ":", weatherData.Main.Temp)
        case "tempmaxc", "tmaxc":
            fmt.Println(requestedData, ":", KelvinToCelsius(weatherData.Main.Temp))
        case "tempmaxf", "tmaxf":
            fmt.Println(requestedData, ":", KelvinToFahrenheit(weatherData.Main.Temp))
        case "tempmaxk", "tmaxk":
            fmt.Println(requestedData, ":", weatherData.Main.Temp)
        case "pressuresealevel", "psl":
            fmt.Println(requestedData, ":", weatherData.Main.SeaLevel)
        case "pressuregroundlevel", "pgl":
            fmt.Println(requestedData, ":", weatherData.Main.GroundLevel)
        case "windspeed", "wispd":
            fmt.Println(requestedData, ":", weatherData.Wind.Speed)
        case "winddeg", "wideg":
            fmt.Println(requestedData, ":", weatherData.Wind.Deg)
        case "clouds", "c":
            fmt.Println(requestedData, ":", weatherData.Clouds.All)
        case "rain", "r":
            fmt.Println(requestedData, ":", weatherData.Rain.ThreeH)
        case "snow", "s":
            fmt.Println(requestedData, ":", weatherData.Snow.ThreeH)
        case "dt":
            fmt.Println(requestedData, ":", weatherData.Dt)
        case "country", "cou":
            fmt.Println(requestedData, ":", weatherData.Sys.Country)
        case "sunrise", "sr":
            fmt.Println(requestedData, ":", weatherData.Sys.Sunrise)
        case "sunset", "ss":
            fmt.Println(requestedData, ":", weatherData.Sys.Sunset)
        case "id":
            fmt.Println(requestedData, ":", weatherData.Id)
        case "name", "n":
            fmt.Println(requestedData, ":", weatherData.Name)
        default:
            fmt.Println(requestedData, "could not be found.")
        }
}


func  KelvinToCelsius(kelvin float32) float32 {
    return float32(kelvin - 273.15)
}

func KelvinToFahrenheit(kelvin float32) float32 {
        return float32((9/5*(kelvin-273.15))+32)
}
