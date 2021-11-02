package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//XML part
	NewFile, err := os.Open("breakfast.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer NewFile.Close()
	data_file, _ := ioutil.ReadAll(NewFile)
	var y Breakfst
	xml.Unmarshal(data_file, &y)
	fmt.Println(y.Food)

	//TO JSON
	var food jsonfood_menu
	var foods []jsonfood_menu

	for _, value := range y.Food {
		food.Food_name = value.Food_name
		food.Description = value.Description
		food.Price_food = value.Price_food
		food.Calories = value.Calories

		foods = append(foods, food)
	}

	jsonData, err := json.Marshal(foods)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//Output
	fmt.Println(string(jsonData))

	//Write into a file
	jsonFile, err := os.Create("./Breakfast_menu.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)

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
type jsonfood_menu struct {
	Food_name   string
	Price_food  string
	Description string
	Calories    int
}

func (p food_menu) String() string {
	return fmt.Sprintf("Name of food: %s,Price of food: %s,Description of food: %s,Calori of food: %d \n", p.Food_name, p.Price_food, p.Description, p.Calories)
}
