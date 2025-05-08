package main

import (
	"io"
	"log"
	"net/http"
	"flag"
	"fmt"
)

func main(){
	
	lat := flag.String("lat","","latitude")
	lon := flag.String("lon","","longitude")
	appid := flag.String("appid","","api key from Open Weather Map")
	flag.Parse()

	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&appid=%s", *lat, *lon, *appid)
	
	resp,err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}
