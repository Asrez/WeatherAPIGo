# WeatherAPIGo

WeatherAPIGo is a RESTful API service built in Go that provides weather information based on user requests. It utilizes external weather data providers to fetch accurate and up-to-date weather data.

## Features

- Fetch current weather information for a specific location.
- Retrieve weather forecast for a specific location.
- Support for multiple weather data providers.

## Prerequisites

To run WeatherAPIGo, make sure you have the following prerequisites installed on your system:

- Go programming language (version 1.16 or higher)
- Git (optional)

## Getting Started

Follow the steps below to get WeatherAPIGo up and running on your local machine:

Clone the repository using Git:

```bash
git clone https://github.com/Asrez/WeatherAPIGo
```

Alternatively, you can download the source code as a ZIP file and extract it.

Change to the project directory:

```bash
cd WeatherAPIGo
```

Build the project:

```bash
go build
```

Run the executable:

```bash
./weatherapigo
```

By default, the server will start on localhost at port 8080. You can access the API using the base URL http://localhost:8080.

## API Endpoints
The following API endpoints are available:

### Get Current Weather

Endpoint: /weather/current

Method: GET

Parameters:
- location (required): The location for which to fetch the weather information (e.g., city name, ZIP code, latitude/longitude).

Response:

```json
{
  "location": "New York",
  "temperature": 25,
  "humidity": 62,
  "description": "Partly cloudy"
  ...
}
```

### Get Weather Forecast

Endpoint: /weather/forecast

Method: GET

Parameters:
- location (required): The location for which to fetch the weather forecast (e.g., city name, ZIP code, latitude/longitude).

Response:

```json
{
  "location": "New York",
  "forecast": [
    {
      "date": "2023-07-08",
      "temperature": 25,
      "humidity": 62,
      "description": "Partly cloudy"
      ...
    },
    ...
  ]
}
```

## Configuration

WeatherAPIGo uses a configuration file (config.json) to manage the API settings. You can modify the configuration according to your needs. The default configuration file looks like this:

```json
{
  "port": 8080,
  "providers": {
    "provider1": {
      "name": "Provider 1",
      "api_key": "your_api_key"
    },
    "provider2": {
      "name": "Provider 2",
      "api_key": "your_api_key"
    }
  }
}
```

- port: The port on which the API server will listen.
- providers: Weather data providers configuration. Add or modify providers as needed, providing the name and API key for each provider.

## Extending Providers

WeatherAPIGo is designed to be extensible with additional weather data providers. To add a new provider, follow these steps:

- Implement a new provider struct that satisfies the weather.Provider interface defined in provider.go.
- Add the necessary fields to the configuration file (config.json) for the new provider.
- Update the GetWeatherProvider function in provider.go to return an instance of the new provider based on the configuration.
- Implement the required methods for the new provider, such as fetching current weather or forecast data.
- Build and run the application.

## Conclusion

WeatherAPIGo provides a simple and flexible RESTful API service to retrieve weather information. It can be easily extended to support multiple weather data providers, allowing you to choose the one that best fits your needs. Feel free to contribute, report issues, or add new features to this project!

Copyright 2023, Asrez group
