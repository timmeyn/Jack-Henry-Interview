
## Usage

1. Run the server:
	
	```
	go run cmd/main.go
	```

	The server will start on http://localhost:8080.

2. Send a request
	
	You can use cURL or Postman to send a POST request to:
	```
	http://localhost:8080/getWeatherData
	```
	Request body (JSON):
	```
	{
	"Latitude": 40.7128,
	"Longitude": -74.0060
	}
	```
3. Example with cURL
	```
	curl -X POST http://localhost:8080/getWeatherData \
	-H "Content-Type: application/json" \
	-d '{"Latitude":40.7128,"Longitude":-74.0060}'
	```
## Response Example
```
{
"ShortForecast": "Sunny",
"TemperatureRating": "It will be hot today, make sure to drink water!"
}
```
Where:
- `ShortForecast` — short text describing current/forecast conditions from the NWS API.
- `TemperatureRating` — one of "Cold", "Moderate", or "Hot" based on temperature thresholds.
## Notes
- Testing

	With more time, I would add basic unit tests and integration tests to verify functionality across edge cases and to detect changes to the NWS API specification early.

- Production Considerations

	In a production environment, I would:
		- Utilize more fields from the NWS API response.
		- Create a dedicated API client wrapper for cleaner and more maintainable requests.
		- Add proper logging, metrics, and error tracking.
		- Configure request timeouts and retry logic for robustness.
