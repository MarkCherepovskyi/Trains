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
	MinInDay        = 1440
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
	data := make([]byte, 64*250)
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

func (info *Info) ShowListOfCity() {
	fmt.Println(info.GetListOfCity())
}

func (info *Info) GetListOfCity() []int {
	list := make([]int, 0)
	for i := 0; i < len(info.Trains); i++ {
		var buffer bool
		for _, station := range list {
			if info.Trains[i].ArrivalStation == station {
				buffer = true
				break
			}
		}
		if !buffer {
			list = append(list, info.Trains[i].ArrivalStation)
		}
	}
	return list
}

func GetTimeInWay(departures, arrival int) int {
	timeInWay := arrival - departures
	if timeInWay <= 0 {
		timeInWay = arrival + MinInDay - departures
	}
	return timeInWay
}

func (info *Info) GetMinPrice(departures, arrival int) (float32, Train) {
	var min float32 = 10000000
	var trainWithMinPrice Train
	for _, train := range info.Trains {
		if train.ArrivalStation == arrival && train.DeparturesStation == departures {
			if train.Price < min {
				min = train.Price
				trainWithMinPrice = train
			}
		}
	}

	return min, trainWithMinPrice
}

func GetTimeInStation(tr1, tr2 Train) int {

	timeInStation := tr2.DeparturesTime - tr1.ArrivalTime
	if timeInStation <= 0 {
		timeInStation = timeInStation + MinInDay
	}
	return timeInStation

}

func (info *Info) GetTrains(departures, arrival int) []Train {
	listOfTrains := make([]Train, 0)
	for _, tr := range info.Trains {
		if tr.DeparturesStation == departures && tr.ArrivalStation == arrival {
			listOfTrains = append(listOfTrains, tr)
		}
	}
	return listOfTrains
}

func (info *Info) minInDays(time int) string {
	days := time / MinInDay
	hours := (time % MinInDay) / 60
	min := time - (days*MinInDay + hours*60)
	return fmt.Sprint(days, " days ", hours, " hours ", min, " min")
}

func ModelInit() *Info {
	info := Info{}
	info.Trains = make([]Train, 0)
	parse(&info, pathToLocalList)
	//info.ShowAll()

	info.GetListOfCity()

	return &info
}
