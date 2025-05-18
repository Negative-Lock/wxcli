# wxcli

**wxcli** is a simple command-line tool built in Go that fetches current weather data from the [OpenWeatherMap API](https://openweathermap.org/api). The project was created primarily as a learning exercise to explore the Go programming language, building CLI applications, and integrating with external APIs.

## Features

- Fetch current weather by latitude and longitude
- Output currently includes the weather (in Fahrenheit for now) and a summary of conditions
- Includes a setup utility to configure your OpenWeatherMap API key and location
- Written in Go

## Usage

Before using the tool, run the setup command to store your OpenWeatherMap API key and coordinate information:

```bash
wxcli setup
