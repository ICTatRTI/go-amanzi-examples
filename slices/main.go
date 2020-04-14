package main

import (
	"fmt"
	"github.com/ICTatRTI/amanzi-timeseries/go-amanzi/collections"
	"github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Open("./slices/data/example.json")
	if err != nil {
		log.Fatal(err)
	}
	var pb ptypes.TimeSeries
	if err :=jsonpb.Unmarshal(f, &pb); err != nil {
		log.Fatal(err)
	}

	data := collections.Float32TsData{
		Times:   make([]time.Time, 0, len(pb.Data)),
		Values:  make([]float32, 0, len(pb.Data)),
		Missing: make(map[int]struct{}),
	}

	if err := collections.Float32Slices(&pb, &data); err != nil {
		log.Fatal(err)
	}
	var sum float32
	for i := 0; i< len(pb.Data); i++ {
		if _, ok := data.Missing[i]; !ok {
			sum += data.Values[i]
		}
	}
	fmt.Printf("the sum is %f", sum)
}
