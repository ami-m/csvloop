package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const bufferSize = 100
const workerCount = 5

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

func processRecord(records <-chan Record, wg *sync.WaitGroup, t *Transformer, v *Validator, w *csv.Writer) {
	defer (*wg).Done()
	for record := range records {
		if !(*v)(record) {
			log.Fatal("failed to validate ", record)
		}
		//output <- (*t)(record)
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
}

func mainLoop(r *csv.Reader, w *csv.Writer, t Transformer, v Validator) {
	var wg sync.WaitGroup
	records := make(chan Record, bufferSize)

	// increment the WaitGroup before starting the worker
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go processRecord(records, &wg, &t, &v, w)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		records <- record
	}

	// to stop the worker, first close the job channel
	close(records)

	// then wait using the WaitGroup
	wg.Wait()
	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	optimisticValidator := func(Record) bool { return true }
	identity := func(record Record) Record { return record }
	r := getCsvReader(getRawReader())
	w := csv.NewWriter(os.Stdout)

	mainLoop(r, w, identity, optimisticValidator)

}
