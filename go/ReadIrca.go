package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("openskies/aircraftDatabase.csv")

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f)

	header, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	data, err := reader.ReadAll()

	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	var tailNoMap = make(map[string]map[string]string)
	var modeSMap = make(map[string]map[string]string)

	for i := 0; i < len(data); i++ {
		var entryMap = make(map[string]string)

		for j := 0; j < len(header); j++ {
			entryMap[header[j]] = data[i][j]
		}

		if entryMap["registration"] != "" {
			tailNoMap[entryMap["registration"]] = entryMap
		}

		if entryMap["icao24"] != "" {
			modeSMap[entryMap["icao24"]] = entryMap
		}
	}

	for {
		var result map[string]string
		var option string

		fmt.Println("Search by Tail Number (1) or Mode S ID (2) or (q) to quit: ")

		_, err := fmt.Scan(&option)

		if err != nil {
			log.Fatal(err)
		}

		if option == "1" {
			var searchTerm string

			fmt.Println("Enter Tail No: ")
			_, err := fmt.Scan(&searchTerm)

			if err != nil {
				log.Fatal(err)
			}

			result = tailNoMap[strings.ToUpper(searchTerm)]

		} else if option == "2" {
			var searchTerm string

			fmt.Println("Enter Mode S ID in hex without the 0x: ")
			_, err := fmt.Scan(&searchTerm)
			if err != nil {
				log.Fatal(err)
			}

			result = modeSMap[strings.ToLower(searchTerm)]

		} else if option == "Q" || option == "q" {
			break
		} else {
			continue
		}

		if result == nil {
			fmt.Println("Item not in DB")
			continue
		} else {
			fmt.Println()
			for key, value := range result {
				fmt.Printf("%20s: %s\n", key, value)
			}
			fmt.Println("--------------------------------")
			fmt.Println()
		}
	}
}
