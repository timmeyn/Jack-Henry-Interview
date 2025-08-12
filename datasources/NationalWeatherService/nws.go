package datasources

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"weatherHTTP/objects"
)

const NWSURL = "https://api.weather.gov"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// Exported function to do everything with one call
func GetWeatherData(ctx context.Context, lat, lon float32) (objects.WeatherOutput, error) {

	var coordData objects.WeatherResponse
	if err := getWithContext(ctx, fmt.Sprintf("%s/points/%v,%v", NWSURL, lat, lon), &coordData); err != nil {
		return objects.WeatherOutput{}, fmt.Errorf("get coord weather: %w", err)
	}

	var forecast objects.Forecast
	if err := getWithContext(ctx, coordData.Properties.Forecast, &forecast); err != nil {
		return objects.WeatherOutput{}, fmt.Errorf("get forecast: %w", err)
	}

	if len(forecast.Properties.Periods) == 0 {
		return objects.WeatherOutput{}, fmt.Errorf("empty forecast periods")
	}

	return convertToReadable(forecast)
}

// Simple get request wrapper to simplify caller
func getWithContext(ctx context.Context, url string, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("status %d: %s", resp.StatusCode, body)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}

// Converting temperature to an enum for ease of understandability
func convertTempToEnum(temp int) objects.TemperatureType {
	switch {
	case temp < 45:
		return objects.TempTypeCold
	case temp < 80:
		return objects.TempTypeModerate
	default:
		return objects.TempTypeHot
	}
}

// Conversion of received large forecast object to the readable response format specified in problem
func convertToReadable(forecast objects.Forecast) (objects.WeatherOutput, error) {
	shortForecast := forecast.Properties.Periods[0].ShortForecast
	temp := convertTempToEnum(forecast.Properties.Periods[0].Temperature)

	output := objects.WeatherOutput{
		ShortForecast:     shortForecast,
		TemperatureRating: temp,
	}

	return output, nil
}
