package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type Entry struct {
	Icao24              string
	Registration        string
	Manufacturericao    string
	Manufacturername    string
	Model               string
	Typecode            string
	Serialnumber        string
	Linenumber          string
	Icaoaircrafttype    string
	Operator            string
	Operatorcallsign    string
	Operatoricao        string
	Operatoriata        string
	Owner               string
	Testreg             string
	Registered          string
	Reguntil            string
	Status              string
	Built               string
	Firstflightdate     string
	Seatconfiguration   string
	Engines             string
	Modes               string
	Adsb                string
	Acars               string
	Notes               string
	CategoryDescription string
}

type Job struct {
	job     int
	entries []Entry
}

func makeMaps(data [][]string) (map[string]Entry, map[string]Entry) {
	tailNoMap := make(map[string]Entry)
	modeSMap := make(map[string]Entry)
	ch := make(chan Entry, len(data))

	fmt.Println("Starting Jobs")

	for i := range data {
		go func(line int) {
			entry := Entry{
				Icao24:              data[line][0],
				Registration:        data[line][1],
				Manufacturericao:    data[line][2],
				Manufacturername:    data[line][3],
				Model:               data[line][4],
				Typecode:            data[line][5],
				Serialnumber:        data[line][6],
				Linenumber:          data[line][7],
				Icaoaircrafttype:    data[line][8],
				Operator:            data[line][9],
				Operatorcallsign:    data[line][10],
				Operatoricao:        data[line][11],
				Operatoriata:        data[line][12],
				Owner:               data[line][13],
				Testreg:             data[line][14],
				Registered:          data[line][15],
				Reguntil:            data[line][16],
				Status:              data[line][17],
				Built:               data[line][18],
				Firstflightdate:     data[line][19],
				Seatconfiguration:   data[line][20],
				Engines:             data[line][21],
				Modes:               data[line][22],
				Adsb:                data[line][23],
				Acars:               data[line][24],
				Notes:               data[line][25],
				CategoryDescription: data[line][26]}

			ch <- entry
		}(i)
	}

	for range data {
		result := <-ch

		if result.Registration != "" {
			tailNoMap[result.Registration] = result
		}

		if result.Icao24 != "" {
			modeSMap[result.Icao24] = result
		}
	}

	return tailNoMap, modeSMap
}

func main() {
	f, err := os.Open("openskies/aircraftDatabase.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Println("Reading Data")

	reader := csv.NewReader(f)

	_, err2 := reader.Read()

	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Println("Read All")

	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Making Maps")

	tailNoMap, modeSMap := makeMaps(data)

	fmt.Println("Maps Complete")

	for {
		var result Entry
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

		if result == (Entry{}) {
			fmt.Println("Item not in DB")
			continue
		} else {
			fmt.Println()
			valueOfResult := reflect.ValueOf(result)
			typeOfResult := valueOfResult.Type()

			for i := 0; i < valueOfResult.NumField(); i++ {
				fmt.Printf("%20s: %s\n", typeOfResult.Field(i).Name, valueOfResult.Field(i).Interface())
			}

			fmt.Println("--------------------------------")
			fmt.Println()
		}
	}
}
