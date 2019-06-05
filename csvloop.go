package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func getRawReader() io.Reader {
	in := `first_namee,last_name,username
"Rob","Pike2",rob
Ken,Thompson,keeen
"Robert","Griesemer","gri"
`
	return strings.NewReader(in)
}

func getCsvReader(r io.Reader) *csv.Reader {
	return csv.NewReader(r)
}

func mainLoop(r *csv.Reader, action func([]string)) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		action(record)
	}
}

func main() {

	r := getCsvReader(getRawReader())
	mainLoop(r, func(record []string) { fmt.Println(record) })

}
