package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"github.com/scythe504/skribbler-backend/internals"
)

func ReadCsvFile(filePath string) []internals.Word {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)

	}
	
	var words []internals.Word

	for _, record := range records {
		if len(record) < 2 {
			log.Println("Skipping invalid record: ", record)
			continue
		}
		count, err := strconv.Atoi(record[1])

		if err != nil {
			log.Println("Invalid count value:", record[1], "in record", record)
			continue
		}

		word := internals.Word {
			Word: record[0],
			Count: count,
		}

		words = append(words, word)
	}

	return words

}