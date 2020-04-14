package main

import (
	"fmt"
	ac "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/collections"
	apt "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"os"
)
func main() {
	f, err := os.Open("./lazyIterator/data/example.json")
	if err != nil {
		log.Fatal(err)
	}
	var pb apt.TimeSeries
	if err :=jsonpb.Unmarshal(f, &pb); err != nil {
		log.Fatal(err)
	}

	itr, err := ac.NewEagerIterator(&pb)
	if err != nil {
		log.Fatal(err)
	}
	var sum float32
	for itr.Next() {
		rec := itr.Float32Record()

		if !rec.Missing {
			sum += rec.Value
		}
	}

	fmt.Printf("the sum is %f", sum)
}
