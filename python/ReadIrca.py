from pathlib import Path
import csv
import timeit
from dataclasses import dataclass
from typing import List

@dataclass
class Entry():
	Icao24             : str = ''
	Registration       : str = ''
	Manufacturericao   : str = ''
	Manufacturername   : str = ''
	Model              : str = ''
	Typecode           : str = ''
	Serialnumber       : str = ''
	Linenumber         : str = ''
	Icaoaircrafttype   : str = ''
	Operator           : str = ''
	Operatorcallsign   : str = ''
	Operatoricao       : str = ''
	Operatoriata       : str = ''
	Owner              : str = ''
	Testreg            : str = ''
	Registered         : str = ''
	Reguntil           : str = ''
	Status             : str = ''
	Built              : str = ''
	Firstflightdate    : str = ''
	Seatconfiguration  : str = ''
	Engines            : str = ''
	Modes              : str = ''
	Adsb               : str = ''
	Acars              : str = ''
	Notes              : str = ''
	CategoryDescription: str = ''

def main() -> None:
    with Timer('Reading Data'):
        with open(Path('openskies/aircraftDatabase.csv'), 'r', encoding='utf-8', newline='') as inputFile:
            reader = csv.DictReader(inputFile)
            inputList: List[Entry] = []

            for row in reader:
                entry = Entry(
                    row['icao24'],
                    row['registration'],
                    row['manufacturericao'],
                    row['manufacturername'],
                    row['model'],
                    row['typecode'],
                    row['serialnumber'],
                    row['linenumber'],
                    row['icaoaircrafttype'],
                    row['operator'],
                    row['operatorcallsign'],
                    row['operatoricao'],
                    row['operatoriata'],
                    row['owner'],
                    row['testreg'],
                    row['registered'],
                    row['reguntil'],
                    row['status'],
                    row['built'],
                    row['firstflightdate'],
                    row['seatconfiguration'],
                    row['engines'],
                    row['modes'],
                    row['adsb'],
                    row['acars'],
                    row['notes'],
                    row['categoryDescription'],
                )

                inputList.append(entry)

    with Timer('Parsing by Tail Number'):
        tailNoDict = {entry.Registration: entry for entry in inputList if entry.Registration not in ['', None]}

    with Timer('Parsing by Mode S ID'):
        modeSDict = {entry.Icao24: entry for entry in inputList if entry.Icao24 not in ['', None]}

    while(True):
        option = input('Search by Tail Number (1) or Mode S ID (2) or (q) to quit: ').upper()

        if option == '1':
            searchTerm = input('Enter Tail No: ').upper()
            result = tailNoDict.get(searchTerm)
        elif option == '2':
            searchTerm = input('Enter Mode S ID in hex without the 0x: ').lower()
            result = modeSDict.get(searchTerm)
        elif option ==  'Q':
            break
        else:
            continue

        if result == None:
            print('Item not in DB')
            continue
        else:
            print()
            for key, value in result.__dict__.items(): print(f'{key:20}: {value}')
            print('--------------------------------')
            print()

class Timer():
    def __init__(self, message: str = 'Operation in Progress') -> None:
        self.message = message
        self.startTime: float = 0.0

    def __enter__(self):
        print(f'{self.message}...')
        self.startTime = timeit.default_timer()

    def __exit__(self, exc_type, exc_value, exc_tb):
        if exc_type == ValueError:
            print(exc_value)
        print(f'{self.message}: {timeit.default_timer() - self.startTime:.2f} seconds')

main()
