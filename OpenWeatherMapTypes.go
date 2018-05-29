package main

type OpenWeatherMapWeatherResponse struct {
    Coord OpenWeatherMapCoord `json:"coord"`
    Weather []OpenWeatherMapWeather`json:"weather"`
    Base string `json:"base"`
    Main OpenWeatherMapMain `json:"main"`
    Visibility int `json:"visibility"`
    Wind OpenWeatherMapWind `json:"wind"`
    Clouds OpenWeatherMapClouds `json:"clouds"`
    Rain OpenWeatherMapRain `json:"rain"`
    Snow OpenWeatherMapSnow `json:"snow"`
    Dt int64 `json:"dt"`
    Sys OpenWeatherMapSys `json:"sys"`
    Id int `json:"id"`
    Name string `json:"name"`
    Cod int `json:"cod"`
}

type OpenWeatherMapCoord struct {
    Lon float32 `json:"lon"`
    Lat float32 `json:"lat"`
}

type OpenWeatherMapWeather struct {
    Id int `json:"id"`
    Main string `json:"main"`
    Description string `json:"description"`
    Icon string `json:"icon"`
}

type OpenWeatherMapMain struct {
    Temp float32 `json:"temp"`
    Pressure float32 `json:"pressure"`
    Humidity float32 `json:"humidity"`
    TempMin float32 `json:"temp-min"`
    TempMax float32 `json:"temp-min"`
    SeaLevel float32 `json:"sea_level"`
    GroundLevel float32 `json:"grnd_level"`
}

type OpenWeatherMapSys struct {
    Type int `json:"type"`
    Id int`json:"id"`
    Message float32 `json:"message"`
    Country string `json:"country"`
    Sunrise int64 `json:"sunrise"`
    Sunset int64 `json:"sunset"`
}

type OpenWeatherMapClouds struct {
    All int `json:"all"`
}

type OpenWeatherMapWind struct {
    Speed float32 `json:"speed"`
    Deg float32 `json:"deg"`
}

type OpenWeatherMapRain struct {
        ThreeH float32 `json:"3h"`
}

type OpenWeatherMapSnow struct {
        ThreeH float32 `json:"3h"`
}

