package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type weatherResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`

	Current struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`

	Daily []struct {
		Dt        int     `json:"dt"`
		Sunrise   int     `json:"sunrise"`
		Sunset    int     `json:"sunset"`
		Moonrise  int     `json:"moonrise"`
		Moonset   int     `json:"moonset"`
		MoonPhase float64 `json:"moon_phase"`
		Summary   string  `json:"summary"`
		Temp      struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float64 `json:"dew_point"`
		WindSpeed float64 `json:"wind_speed"`
		WindDeg   int     `json:"wind_deg"`
		WindGust  float64 `json:"wind_gust"`
		Weather   []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds int     `json:"clouds"`
		Pop    float64 `json:"pop"`
		Rain   float64 `json:"rain,omitempty"`
		Uvi    float64 `json:"uvi"`
	} `json:"daily"`
}

func printHelp() {

	// Title/Header
	fmt.Println("\nUsage: wcli <command>")
	fmt.Println("A simple CLI weather app.")
	fmt.Println("USAGE:")

	/*
		Defining the anonymous commands struct allows us to create structured data that will
		be easily accessible later on in for loop.

		In the second set of curly braces after the commands declaration we are assigning our
		actual values for our command and description. New commands and their description will
		be added here.
	*/

	commands := []struct {
		Name        string
		Description string
	}{
		{"setup", "Configure your location and API key"},
		{"current", "Show current weather info"},
		{"help", "Display this help message"},
	}

	/*
		Ignore index returned by range with _ in for loop. Iterate over commands struct with cmd
		and access structs components (Name, Description) with cmd.Name and cmd.Description and
		display it to the screen follow by newline.

		Format: %-10s = left-aligned, 10-character column
	*/

	for _, cmd := range commands {
		fmt.Printf("  %-10s %s\n", cmd.Name, cmd.Description)
	}

	fmt.Println()
}

func setup() {

	// Attempt to create (or overwrite) a file named ".env"
	file, err := os.Create(".env")
	if err != nil {
		// If there's an error creating the file, log it and stop execution
		log.Fatalf("Could not create file: %v", err)
	}
	// Ensure the file is closed after this function ends
	defer file.Close()

	// Create a buffered reader to read input from standard input (the terminal)
	reader := bufio.NewReader(os.Stdin)

	// A list of environment variable keys we want the user to input
	prompts := []string{"LATITUDE", "LONGITUDE", "API_KEY"}

	// Loop over each key and prompt the user for a value
	for _, key := range prompts {
		// Prompt the user
		fmt.Printf("Enter %s: ", key)
		// Read the user input until a newline is encountered
		input, _ := reader.ReadString('\n')
		// Remove the newline character from the input
		input = strings.TrimSpace(input)
		// Write the key-value pair to the .env file in "KEY=value" format
		_, err := fmt.Fprintf(file, "%s=%s\n", key, input)
		if err != nil {
			// If there's an error writing to the file, log it and stop execution
			log.Fatalf("Failed to write to file: %v", err)
		}
	}
	// Let the user know the .env file has been created
	fmt.Println(".env file created successfully.")

}

func getWeather() weatherResponse {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File. Run wcli help for setup command.")
	}

	lat := os.Getenv("LATITUDE")
	lon := os.Getenv("LONGITUDE")
	appid := os.Getenv("API_KEY")

	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&units=imperial&exclude=minutely,hourly,alerts&appid=%s", lat, lon, appid)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var weatherData weatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Fatalln(err)
	}

	return weatherData
}

func weatherToday(weatherData weatherResponse) {

	fmt.Printf("Current Temp: %.1f°F\n", weatherData.Current.Temp)
	fmt.Printf("Weather: %s\n", weatherData.Daily[0].Summary)

}

func weatherDaily(weatherData weatherResponse) {

	days := len(weatherData.Daily)

	for day := range days {

		t := time.Unix(int64(weatherData.Daily[day].Dt), 0)

		fmt.Printf("%s ", t.Weekday())
		fmt.Printf("%.1f°F\n", weatherData.Daily[day].Temp.Day)
	}

}

func main() {

	if len(os.Args) < 2 {

		printHelp()
		log.Fatal("No arguments provided.")
	}

	switch os.Args[1] {

	case "setup":
		setup()

	case "current":
		weatherToday(getWeather())

	case "daily":
		weatherDaily(getWeather())

	case "help":
		printHelp()

	default:
		fmt.Println("Unknown argument. Use 'wcli help' for a list of commands.")
		os.Exit(1)
	}
}
