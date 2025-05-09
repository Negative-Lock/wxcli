# wxcli

wxcli is a simple command-line application written in Go that fetches and displays current weather data using the [OpenWeatherMap API](https://openweathermap.org/api).

## 🌦 What It Does

- Retrieves current weather information based on your specified location.
- Displays current temperature and a brief weather description in your terminal.
- Plans to add more functionality and learn more about Go!

## 📚 Purpose

This project is primarily a **learning tool**. I’m using it to explore and deepen my understanding of:

- Making HTTP requests in Go
- Parsing JSON data into Go structs
- Working with Go modules
- Building terminal-based applications
- Go in general

## 🔧 How It Works

1. Sends a GET request to the OpenWeatherMap API.
2. Unmarshals the JSON response into Go structs.
3. Outputs the relevant weather information to the terminal.
