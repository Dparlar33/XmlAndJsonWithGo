package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	NewFile, err := os.Open("breakfast.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer NewFile.Close()
	data_file, _ := ioutil.ReadAll(NewFile)
	var y Breakfst
	xml.Unmarshal(data_file, &y)
	fmt.Println(y.Food)
}

type Breakfst struct {
	XmlName xml.Name    `xml:"breakfast_menu"`
	Food    []food_menu `xml:"food"`
}

type food_menu struct {
	XmlName     xml.Name `xml:"food"`
	Food_name   string   `xml:"name"`
	Price_food  string   `xml:"price"`
	Description string   `xml:"description"`
	Calories    int      `xml:"calories"`
}

func (p food_menu) String() string {
	return fmt.Sprintf("Name of food: %s,Price of food: %s,Description of food: %s,Calori of food: %d \n", p.Food_name, p.Price_food, p.Description, p.Calories)
}
