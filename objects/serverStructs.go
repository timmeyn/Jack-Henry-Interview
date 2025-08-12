package objects

//expected request form
type WeatherInput struct {
	Latitude  float32
	Longitude float32
}

//response form
type WeatherOutput struct {
	ShortForecast     string
	TemperatureRating TemperatureType
}

//setting up an enum for general temperature ratings
type TemperatureType string

var (
	TempTypeHot      TemperatureType = "It will be hot today, make sure to drink water!"
	TempTypeModerate TemperatureType = "The temperature will be moderate today."
	TempTypeCold     TemperatureType = "It will be cold today, bundle up!"
)
