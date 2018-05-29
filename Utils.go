package main

import(
    "fmt"
    "strings"
    )

func OutputWeatherData(weatherData *OpenWeatherMapWeatherResponse, requestedData string){
        switch strings.ToLower(requestedData) {
        case "longitude":
            fmt.Println(requestedData, ":", weatherData.Coord.Lon)
        case "latitude":
            fmt.Println(requestedData, ":", weatherData.Coord.Lat)
        case "weatherid":
            fmt.Println(requestedData, ":", weatherData.Weather[0].Id)
        case "weatherdesc":
            fmt.Println(requestedData, ":", weatherData.Weather[0].Description)
        case "weathericon":
            fmt.Println(requestedData, ":", weatherData.Weather[0].Icon)
        case "base":
            fmt.Println(requestedData, ":", weatherData.Base)
        case "tempc":
            fmt.Println(requestedData, ":", KelvinToCelsius(weatherData.Main.Temp))
        case "tempf":
            fmt.Println(requestedData, ":", KelvinToFahrenheit(weatherData.Main.Temp))
        case "tempk":
            fmt.Println(requestedData, ":", weatherData.Main.Temp)
        case "pressure":
            fmt.Println(requestedData, ":", weatherData.Main.Pressure)
        case "humidity":
            fmt.Println(requestedData, ":", weatherData.Main.Humidity)
        case "tempminc":
            fmt.Println(requestedData, ":", KelvinToCelsius(weatherData.Main.Temp))
        case "tempminf":
            fmt.Println(requestedData, ":", KelvinToFahrenheit(weatherData.Main.Temp))
        case "tempmink":
            fmt.Println(requestedData, ":", weatherData.Main.Temp)
        case "tempmaxc":
            fmt.Println(requestedData, ":", KelvinToCelsius(weatherData.Main.Temp))
        case "tempmaxf":
            fmt.Println(requestedData, ":", KelvinToFahrenheit(weatherData.Main.Temp))
        case "tempmaxk":
            fmt.Println(requestedData, ":", weatherData.Main.Temp)
        case "sealevel":
            fmt.Println(requestedData, ":", weatherData.Main.SeaLevel)
        case "groundlevel":
            fmt.Println(requestedData, ":", weatherData.Main.GroundLevel)
        case "windspeed":
            fmt.Println(requestedData, ":", weatherData.Wind.Speed)
        case "winddeg":
            fmt.Println(requestedData, ":", weatherData.Wind.Deg)
        case "clouds":
            fmt.Println(requestedData, ":", weatherData.Clouds.All)
        case "rain":
            fmt.Println(requestedData, ":", weatherData.Rain.ThreeH)
        case "snow":
            fmt.Println(requestedData, ":", weatherData.Snow.ThreeH)
        case "dt":
            fmt.Println(requestedData, ":", weatherData.Dt)
        case "country":
            fmt.Println(requestedData, ":", weatherData.Sys.Country)
        case "sunrise":
            fmt.Println(requestedData, ":", weatherData.Sys.Sunrise)
        case "sunset":
            fmt.Println(requestedData, ":", weatherData.Sys.Sunset)
        case "id":
            fmt.Println(requestedData, ":", weatherData.Id)
        case "name":
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
