// Michał Pomirski 06.06.2024
package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"github.com/gocolly/colly"
)

type EiffelTowerHeaders struct {
	Name string
	Country string
	Location string
	Height string
	Notes string
}

type EiffelTower struct {
	Name  string
	Country string
	Location string
	Height string
	Notes string
}

type EiffelTowerList struct {
	EiffelTowers []EiffelTower
}

func getTower(el *colly.HTMLElement) EiffelTower {
	tower := EiffelTower{}
	tower.Name = el.ChildText("td:nth-child(1)")
	tower.Country = el.ChildText("td:nth-child(2)")
	tower.Location = el.ChildText("td:nth-child(3)")
	tower.Height = el.ChildText("td:nth-child(4)")
	tower.Notes = el.ChildText("td:nth-child(7)")
	return tower
}


func writeCSVFile(headers EiffelTowerHeaders, towers EiffelTowerList) {
	log.Println("Opening CSV file")
	csvFile, err := os.Create("towers.csv")
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	csvWriter := csv.NewWriter(csvFile)
	defer func(){
		log.Println("Flushing")
		csvWriter.Flush()
		log.Println("Closing")
		csvFile.Close()
	}()

	csvWriter.Write([]string{headers.Name, headers.Country, headers.Location, headers.Height, headers.Notes})

	for _, tower := range towers.EiffelTowers {
		row := []string{tower.Name, tower.Country, tower.Location, tower.Height, tower.Notes}
		log.Println("Writing row", row)
		csvWriter.Write(row)
	}
	log.Println("Done writing")
}

func main() {
	c := colly.NewCollector()

	towerList := EiffelTowerList{}
	EiffelTowerHeaders := EiffelTowerHeaders{"Name", "Country", "Location", "Height", "Notes"}

	c.OnHTML("table > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			tower := getTower(el)
			if tower.Name != "" && tower.Country != "" && tower.Location != "" && tower.Height != ""{
				towerList.EiffelTowers = append(towerList.EiffelTowers, tower)
			}
		})
	})

	c.Visit("https://en.wikipedia.org/wiki/Eiffel_Tower_replicas_and_derivatives")

	log.Println("Finished scraping")

	writeCSVFile(EiffelTowerHeaders, towerList)

}