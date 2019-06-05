package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

type Record []string
type Transformer func(Record) Record
type Validator func(Record) bool

func getRawReader() io.Reader {
	in := `first_namee,last_name,username
"Rob","Pike2",robdd
Ken,Thompson,keeen
"Robert","Griesemer","gri"
`
	return strings.NewReader(in)
}

func getCsvReader(r io.Reader) *csv.Reader {
	return csv.NewReader(r)
}

func mainLoop(r *csv.Reader, t Transformer, v Validator) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if !v(record) {
			log.Fatal("failed to validate ", record)
		}

		fmt.Println(t(record))
	}
}

func main() {
	optimisticValidator := func(Record) bool { return true }
	identity := func(record Record) Record { return record }
	r := getCsvReader(getRawReader())

	mainLoop(r, identity, optimisticValidator)

}
