package main

import (
	"encoding/csv"
	"fmt"
	"io"
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

	var tailNoMap = make(map[string]Entry)
	var modeSMap = make(map[string]Entry)

	fmt.Println("Making Maps")

	for {
		data, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		entry := Entry{
			Icao24:              data[0],
			Registration:        data[1],
			Manufacturericao:    data[2],
			Manufacturername:    data[3],
			Model:               data[4],
			Typecode:            data[5],
			Serialnumber:        data[6],
			Linenumber:          data[7],
			Icaoaircrafttype:    data[8],
			Operator:            data[9],
			Operatorcallsign:    data[10],
			Operatoricao:        data[11],
			Operatoriata:        data[12],
			Owner:               data[13],
			Testreg:             data[14],
			Registered:          data[15],
			Reguntil:            data[16],
			Status:              data[17],
			Built:               data[18],
			Firstflightdate:     data[19],
			Seatconfiguration:   data[20],
			Engines:             data[21],
			Modes:               data[22],
			Adsb:                data[23],
			Acars:               data[24],
			Notes:               data[25],
			CategoryDescription: data[26]}

		if entry.Registration != "" {
			tailNoMap[entry.Registration] = entry
		}

		if entry.Icao24 != "" {
			modeSMap[entry.Icao24] = entry
		}
	}

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
