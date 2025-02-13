package models

// Weather API Result model, used https://mholt.github.io/json-to-go/ to convert json to go struct
type WeatherAPIResult struct {
	Address           string             `json:"address"`
	Alerts            []Alerts           `json:"alerts"`
	CurrentConditions CurrentConditions  `json:"currentConditions"`
	Days              []Days             `json:"days"`
	Description       string             `json:"description"`
	Latitude          float64            `json:"latitude"`
	Longitude         float64            `json:"longitude"`
	QueryCost         float64            `json:"queryCost"`
	ResolvedAddress   string             `json:"resolvedAddress"`
	Stations          map[string]Station `json:"stations"`
	Timezone          string             `json:"timezone"`
	Tzoffset          float64            `json:"tzoffset"`
}
type Alerts struct {
	Description string  `json:"description"`
	Ends        string  `json:"ends"`
	EndsEpoch   float64 `json:"endsEpoch"`
	Event       string  `json:"event"`
	Headline    string  `json:"headline"`
	ID          string  `json:"id"`
	Language    string  `json:"language"`
	Link        any     `json:"link"`
	Onset       string  `json:"onset"`
	OnsetEpoch  float64 `json:"onsetEpoch"`
}
type CurrentConditions struct {
	Cloudcover     float64  `json:"cloudcover"`
	Conditions     string   `json:"conditions"`
	Datetime       string   `json:"datetime"`
	DatetimeEpoch  float64  `json:"datetimeEpoch"`
	Dew            float64  `json:"dew"`
	Feelslike      float64  `json:"feelslike"`
	Humidity       float64  `json:"humidity"`
	Icon           string   `json:"icon"`
	Moonphase      float64  `json:"moonphase"`
	Precip         float64  `json:"precip"`
	Precipprob     float64  `json:"precipprob"`
	Preciptype     any      `json:"preciptype"`
	Pressure       float64  `json:"pressure"`
	Snow           float64  `json:"snow"`
	Snowdepth      float64  `json:"snowdepth"`
	Solarenergy    float64  `json:"solarenergy"`
	Solarradiation float64  `json:"solarradiation"`
	Source         string   `json:"source"`
	Stations       []string `json:"stations"`
	Sunrise        string   `json:"sunrise"`
	SunriseEpoch   float64  `json:"sunriseEpoch"`
	Sunset         string   `json:"sunset"`
	SunsetEpoch    float64  `json:"sunsetEpoch"`
	Temp           float64  `json:"temp"`
	Uvindex        float64  `json:"uvindex"`
	Visibility     float64  `json:"visibility"`
	Winddir        float64  `json:"winddir"`
	Windgust       float64  `json:"windgust"`
	Windspeed      float64  `json:"windspeed"`
}
type Hours struct {
	Cloudcover     float64  `json:"cloudcover"`
	Conditions     string   `json:"conditions"`
	Datetime       string   `json:"datetime"`
	DatetimeEpoch  float64  `json:"datetimeEpoch"`
	Dew            float64  `json:"dew"`
	Feelslike      float64  `json:"feelslike"`
	Humidity       float64  `json:"humidity"`
	Icon           string   `json:"icon"`
	Precip         float64  `json:"precip"`
	Precipprob     float64  `json:"precipprob"`
	Preciptype     any      `json:"preciptype"`
	Pressure       float64  `json:"pressure"`
	Severerisk     float64  `json:"severerisk"`
	Snow           float64  `json:"snow"`
	Snowdepth      float64  `json:"snowdepth"`
	Solarenergy    float64  `json:"solarenergy"`
	Solarradiation float64  `json:"solarradiation"`
	Source         string   `json:"source"`
	Stations       []string `json:"stations"`
	Temp           float64  `json:"temp"`
	Uvindex        float64  `json:"uvindex"`
	Visibility     float64  `json:"visibility"`
	Winddir        float64  `json:"winddir"`
	Windgust       float64  `json:"windgust"`
	Windspeed      float64  `json:"windspeed"`
}
type Days struct {
	Cloudcover     float64  `json:"cloudcover"`
	Conditions     string   `json:"conditions"`
	Datetime       string   `json:"datetime"`
	DatetimeEpoch  float64  `json:"datetimeEpoch"`
	Description    string   `json:"description"`
	Dew            float64  `json:"dew"`
	Feelslike      float64  `json:"feelslike"`
	Feelslikemax   float64  `json:"feelslikemax"`
	Feelslikemin   float64  `json:"feelslikemin"`
	Hours          []Hours  `json:"hours"`
	Humidity       float64  `json:"humidity"`
	Icon           string   `json:"icon"`
	Moonphase      float64  `json:"moonphase"`
	Precip         float64  `json:"precip"`
	Precipcover    float64  `json:"precipcover"`
	Precipprob     float64  `json:"precipprob"`
	Preciptype     any      `json:"preciptype"`
	Pressure       float64  `json:"pressure"`
	Severerisk     float64  `json:"severerisk"`
	Snow           float64  `json:"snow"`
	Snowdepth      float64  `json:"snowdepth"`
	Solarenergy    float64  `json:"solarenergy"`
	Solarradiation float64  `json:"solarradiation"`
	Source         string   `json:"source"`
	Stations       []string `json:"stations"`
	Sunrise        string   `json:"sunrise"`
	SunriseEpoch   float64  `json:"sunriseEpoch"`
	Sunset         string   `json:"sunset"`
	SunsetEpoch    float64  `json:"sunsetEpoch"`
	Temp           float64  `json:"temp"`
	Tempmax        float64  `json:"tempmax"`
	Tempmin        float64  `json:"tempmin"`
	Uvindex        float64  `json:"uvindex"`
	Visibility     float64  `json:"visibility"`
	Winddir        float64  `json:"winddir"`
	Windgust       float64  `json:"windgust"`
	Windspeed      float64  `json:"windspeed"`
}
type Station struct {
	Contribution float64 `json:"contribution"`
	Distance     float64 `json:"distance"`
	ID           string  `json:"id"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	Quality      float64 `json:"quality"`
	UseCount     float64 `json:"useCount"`
}
