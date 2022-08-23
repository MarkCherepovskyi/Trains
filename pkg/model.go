package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	pathToLocalList = "config/test.csv"
	secInDay        = 86400
)

type Train struct {
	NumOfTrain        int
	DeparturesStation int
	ArrivalStation    int
	Price             float32
	DeparturesTime    int
	ArrivalTime       int
}

type Info struct {
	Trains []Train
}

func (info Info) GetInfoByArrivalStation(numOfStation int) {

}
func (info Info) GetInfoByDeparturesStation(numOfStation int) {

}

func parse(info *Info, pathToList string) {
	file, err := os.Open(pathToList)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	data := make([]byte, 50*250)
	tr := Train{}
	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}

	}

	var str string
	var j int
	for _, char := range string(data) {

		if char == ';' {
			if strings.Contains(str, ".") {
				parseFloat, err := strconv.ParseFloat(string(str[:j]), 64)
				if err != nil {
					log.Println(err)
				}
				tr.Price = float32(parseFloat)
			} else if strings.Contains(str, ":") {
				fmt.Println(str)
				if tr.DeparturesTime == 0 {
					tr.DeparturesTime = strToTime(string(str[:]))
				}

			} else if !strings.Contains(str, ".") && !strings.Contains(str, ";") {
				parseInt, err := strconv.ParseInt(string(str[:j]), 10, 64)
				if err != nil {
					log.Println(err)
				}
				if tr.NumOfTrain == 0 {
					tr.NumOfTrain = int(parseInt)
				} else if tr.DeparturesStation == 0 {
					tr.DeparturesStation = int(parseInt)
				} else if tr.ArrivalStation == 0 {
					tr.ArrivalStation = int(parseInt)
				}

			}
			str = ""
			j = 0
		} else if char == '\n' {
			if tr.ArrivalTime == 0 {

				tr.ArrivalTime = strToTime(string(str[:]))
			}
			info.Trains = append(info.Trains, tr)
			tr = Train{}
			str = ""
			j = 0
		} else if char != ';' && char != '\n' {
			j++
			str += string(char)
		}

	}

}

func strToTime(str string) int {
	var (
		timeInInt [3]int64
		err       error
	)
	timeInInt[0], err = strconv.ParseInt(str[:2], 10, 64)
	if err != nil {
		log.Panic(err)
	}
	timeInInt[1], err = strconv.ParseInt(str[3:5], 10, 64)
	if err != nil {
		log.Panic(err)
	}
	timeInInt[2], err = strconv.ParseInt(str[6:8], 10, 64)
	if err != nil {
		log.Panic(err)
	}
	return int(timeInInt[0]*60 + timeInInt[1])
}

func (info *Info) ShowAll() {
	fmt.Println(len(info.Trains))
	for _, train := range info.Trains {
		text, _ := json.Marshal(train)
		fmt.Print(string(text))
		fmt.Println()
	}
}

func ModelInit() *Info {
	info := Info{}
	info.Trains = make([]Train, 0)
	parse(&info, pathToLocalList)
	info.ShowAll()

	return &info
}
