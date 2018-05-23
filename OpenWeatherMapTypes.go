package main

type OpenWeatherMapWeatherResponse struct {
    Coord OpenWeatherMapCoord `json:"coord"`
    Weather []OpenWeatherMapWeather`json:"weather"`
    Base string `json:"base"`
    Main OpenWeatherMapMain `json:"main"`
    Visibility int `json:"visibility"`
    Wind OpenWeatherMapWind `json:"wind"`
    Clouds OpenWeatherMapClouds `json:"clouds"`
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
    Pressure int `json:"pressure"`
    TempMin float32 `json:"temp-min"`
    TempMax float32 `json:"temp-min"`
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
    Deg int `json:"deg"`
}

