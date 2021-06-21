package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
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

// type Job struct {
// 	job     int
// 	entries []Entry
// }

func readEntries(dataChan chan []string, entryChan chan Entry) {
	wg := sync.WaitGroup{}
	defer close(entryChan)

	for line := range dataChan {
		wg.Add(1)

		go func(line []string) {
			defer wg.Done()

			entry := Entry{
				Icao24:              line[0],
				Registration:        line[1],
				Manufacturericao:    line[2],
				Manufacturername:    line[3],
				Model:               line[4],
				Typecode:            line[5],
				Serialnumber:        line[6],
				Linenumber:          line[7],
				Icaoaircrafttype:    line[8],
				Operator:            line[9],
				Operatorcallsign:    line[10],
				Operatoricao:        line[11],
				Operatoriata:        line[12],
				Owner:               line[13],
				Testreg:             line[14],
				Registered:          line[15],
				Reguntil:            line[16],
				Status:              line[17],
				Built:               line[18],
				Firstflightdate:     line[19],
				Seatconfiguration:   line[20],
				Engines:             line[21],
				Modes:               line[22],
				Adsb:                line[23],
				Acars:               line[24],
				Notes:               line[25],
				CategoryDescription: line[26]}

			entryChan <- entry
		}(line)
	}

	wg.Wait()

	fmt.Println("Finished Reading")
}

func addToMaps(entry Entry, tailNoMap *map[string]Entry, modeSMap *map[string]Entry, wg *sync.WaitGroup, tnMutex *sync.Mutex, mSMutex *sync.Mutex) {
	defer wg.Done()

	if entry.Registration != "" {
		tnMutex.Lock()
		(*tailNoMap)[entry.Registration] = entry
		tnMutex.Unlock()
	}

	if entry.Icao24 != "" {
		mSMutex.Lock()
		(*modeSMap)[entry.Icao24] = entry
		mSMutex.Unlock()
	}
}

func makeMaps(dataChan chan []string) (map[string]Entry, map[string]Entry) {
	tailNoMap := make(map[string]Entry)
	modeSMap := make(map[string]Entry)
	entryChan := make(chan Entry)
	wg := sync.WaitGroup{}
	tnMutex := sync.Mutex{}
	mSMutex := sync.Mutex{}

	fmt.Println("Starting Jobs")

	go readEntries(dataChan, entryChan)

	fmt.Println("Ready to Recieve")

	for entry := range entryChan {
		wg.Add(1)
		go addToMaps(entry, &tailNoMap, &modeSMap, &wg, &tnMutex, &mSMutex)
	}

	wg.Wait()

	fmt.Println("Receive Complete")

	return tailNoMap, modeSMap
}

func readData(reader *csv.Reader, dataChan chan []string) {
	defer close(dataChan)

	for {
		line, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			dataChan <- line
		}
	}
}

func main() {
	dataChan := make(chan []string)
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

	fmt.Println("Read Data")

	go readData(reader, dataChan)

	fmt.Println("Making Maps")

	makeMaps(dataChan)
	// tailNoMap, modeSMap := makeMaps(dataChan)

	fmt.Println("Maps Complete")

	// for {
	// 	var result Entry
	// 	var option string

	// 	fmt.Println("Search by Tail Number (1) or Mode S ID (2) or (q) to quit: ")

	// 	_, err := fmt.Scan(&option)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if option == "1" {
	// 		var searchTerm string

	// 		fmt.Println("Enter Tail No: ")
	// 		_, err := fmt.Scan(&searchTerm)

	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		result = tailNoMap[strings.ToUpper(searchTerm)]

	// 	} else if option == "2" {
	// 		var searchTerm string

	// 		fmt.Println("Enter Mode S ID in hex without the 0x: ")
	// 		_, err := fmt.Scan(&searchTerm)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		result = modeSMap[strings.ToLower(searchTerm)]

	// 	} else if option == "Q" || option == "q" {
	// 		break
	// 	} else {
	// 		continue
	// 	}

	// 	if result == (Entry{}) {
	// 		fmt.Println("Item not in DB")
	// 		continue
	// 	} else {
	// 		fmt.Println()
	// 		valueOfResult := reflect.ValueOf(result)
	// 		typeOfResult := valueOfResult.Type()

	// 		for i := 0; i < valueOfResult.NumField(); i++ {
	// 			fmt.Printf("%20s: %s\n", typeOfResult.Field(i).Name, valueOfResult.Field(i).Interface())
	// 		}

	// 		fmt.Println("--------------------------------")
	// 		fmt.Println()
	// 	}
	// }
}
