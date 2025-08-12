package objects

//Response items for overall weather forecast from base level endpoint
type WeatherResponse struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Geometry   WeatherGeometry   `json:"geometry"`
	Properties WeatherProperties `json:"properties"`
}

type WeatherGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type WeatherProperties struct {
	ID                  string           `json:"@id"`
	Type                string           `json:"@type"`
	Cwa                 string           `json:"cwa"`
	ForecastOffice      string           `json:"forecastOffice"`
	GridID              string           `json:"gridId"`
	GridX               int              `json:"gridX"`
	GridY               int              `json:"gridY"`
	Forecast            string           `json:"forecast"`
	ForecastHourly      string           `json:"forecastHourly"`
	ForecastGridData    string           `json:"forecastGridData"`
	ObservationStations string           `json:"observationStations"`
	RelativeLocation    RelativeLocation `json:"relativeLocation"`
	ForecastZone        string           `json:"forecastZone"`
	County              string           `json:"county"`
	FireWeatherZone     string           `json:"fireWeatherZone"`
	TimeZone            string           `json:"timeZone"`
	RadarStation        string           `json:"radarStation"`
}

type RelativeLocation struct {
	Type       string             `json:"type"`
	Geometry   LocationGeometry   `json:"geometry"`
	Properties LocationProperties `json:"properties"`
}

type LocationGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type LocationProperties struct {
	City     string `json:"city"`
	State    string `json:"state"`
	Distance struct {
		UnitCode string  `json:"unitCode"`
		Value    float64 `json:"value"`
	} `json:"distance"`
	Bearing struct {
		UnitCode string `json:"unitCode"`
		Value    int    `json:"value"`
	} `json:"bearing"`
}

//Response items for forecast endpoint given by base endpoint
type Forecast struct {
	Type       string             `json:"type"`
	Properties ForecastProperties `json:"properties"`
}

type ForecastProperties struct {
	Units             string            `json:"units"`
	ForecastGenerator string            `json:"forecastGenerator"`
	GeneratedAt       string            `json:"generatedAt"`
	UpdateTime        string            `json:"updateTime"`
	ValidTimes        string            `json:"validTimes"`
	Elevation         ForecastElevation `json:"elevation"`
	Periods           []ForecastPeriod  `json:"periods"`
}

type ForecastElevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type ForecastPeriod struct {
	Number                     int                       `json:"number"`
	Name                       string                    `json:"name"`
	StartTime                  string                    `json:"startTime"`
	EndTime                    string                    `json:"endTime"`
	IsDaytime                  bool                      `json:"isDaytime"`
	Temperature                int                       `json:"temperature"`
	TemperatureUnit            string                    `json:"temperatureUnit"`
	TemperatureTrend           string                    `json:"temperatureTrend"`
	ProbabilityOfPrecipitation ForecastPrecipProbability `json:"probabilityOfPrecipitation"`
	WindSpeed                  string                    `json:"windSpeed"`
	WindDirection              string                    `json:"windDirection"`
	Icon                       string                    `json:"icon"`
	ShortForecast              string                    `json:"shortForecast"`
	DetailedForecast           string                    `json:"detailedForecast"`
}

type ForecastPrecipProbability struct {
	UnitCode string `json:"unitCode"`
	Value    int    `json:"value"`
}
