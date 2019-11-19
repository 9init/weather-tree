package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Jacket struct {
	state bool
}

type WindyS struct {
	Windy    Jacket
	NotWindy Jacket
}

type Humidity struct {
	High   WindyS
	Normal WindyS
}

type Tempreture struct {
	Cool Humidity
	Mild Humidity
	Hot  Humidity
}

type Outlook struct {
	Sunny    Tempreture
	Overcast Tempreture
	Rainy    Tempreture
}

var jacket Jacket
var windy WindyS
var humidity Humidity
var tempreture_Sunny, tempreture_Overcast, tempreture_Rainy Tempreture
var lastTemp string

//Train starts here
func (m *Outlook) Train(data []string) {

	switch data[4] {
	case "Yes":
		jacket.state = true
	case "No":
		jacket.state = false
	}

	switch data[3] {
	case "FALSE":
		windy.NotWindy = jacket
	case "TRUE":
		windy.Windy = jacket
	}

	switch data[2] {
	case "High":
		humidity.High = windy
	case "Normal":
		humidity.Normal = windy
	}

	switch data[0] {
	case "Sunny":
		switch data[1] {
		case "Hot":
			tempreture_Sunny.Hot = humidity
		case "Cool":
			tempreture_Sunny.Cool = humidity
		case "Mild":
			tempreture_Sunny.Mild = humidity
		}
		m.Sunny = tempreture_Sunny
	case "Overcast":
		switch data[1] {
		case "Hot":
			tempreture_Overcast.Hot = humidity
		case "Cool":
			tempreture_Overcast.Cool = humidity
		case "Mild":
			tempreture_Overcast.Mild = humidity
		}
		m.Overcast = tempreture_Overcast
	case "Rainy":
		switch data[1] {
		case "Hot":
			tempreture_Rainy.Hot = humidity
		case "Cool":
			tempreture_Rainy.Cool = humidity
		case "Mild":
			tempreture_Rainy.Mild = humidity
		}
		m.Rainy = tempreture_Rainy
	}
}

func main() {

	var tree Outlook
	//read data from csv
	file, err := os.Open(`D:\Ahmed\assignment\Serin one\assignment.csv`)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataList := strings.Split(scanner.Text(), ",")
		tree.Train(dataList)
	}
	fmt.Printf("state of wearing jacket is : %v", tree.Overcast.Hot.Normal.NotWindy.state)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
