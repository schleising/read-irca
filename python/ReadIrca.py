from pathlib import Path
import csv
import timeit

def main() -> None:
    with Timer('Reading Data'):
        with open(Path('openskies/aircraftDatabase.csv'), 'r', encoding='utf-8') as inputFile:
            reader = csv.DictReader(inputFile)
            inputList = [row for row in reader]

    with Timer('Parsing by Tail Number'):
        tailNoDict = {entry['registration']: entry for entry in inputList if entry['registration'] not in ['', None]}

    with Timer('Parsing by Mode S ID'):
        modeSDict = {entry['icao24']: entry for entry in inputList if entry['icao24'] not in ['', None]}

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
            for key, value in result.items(): print(f'{key:20}: {value}')
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
