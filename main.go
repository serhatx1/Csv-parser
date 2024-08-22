package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Tech struct {
	CompanyName  string
	Founded      uint16
	HeadQuarters string
	Ceo          string
	Industry     string
}

const (
	CompanyName = iota
	Founded
	HeadQuarters
	Ceo
	Industry
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		FoundedInt, err := strconv.ParseUint(row[Founded], 10, 64)
		if err != nil {
			return
		}
		FoundedUInt := uint16(FoundedInt)
		Json, err := json.Marshal(Tech{
			CompanyName:  row[CompanyName],
			Founded:      FoundedUInt,
			HeadQuarters: row[HeadQuarters],
			Ceo:          row[Ceo],
			Industry:     row[Industry],
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", Json)
	}
}
