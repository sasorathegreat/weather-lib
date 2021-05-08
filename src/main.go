package main

// Importing packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// Finding the Max of two numbers
	//var result string
	fmt.Println("Hello, world!")
	fmt.Println(Get("28027"))
	//result = Get("28027")
}

func Get(zipcode string) string {
	var value string
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?zip=28027&lang=fr&appid=809f761fd91b3990cdc45262b01aa174")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	return value
}
