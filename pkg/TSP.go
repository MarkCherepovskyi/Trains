package pkg

import (
	"fmt"
)

var (
	firstStation    int
	list            []int
	numOfStation    int
	info            *Info
	roads           []path
	minPath         float32 = 10000
	counter         int     = 0
	minCounter      int
	bestPath        string = "undefined path"
	bestPrice       float32
	distanceByPrice [][]float32
	cheapTrains     []Train

	switcher = 2

	indexes       []int
	indexesByTime []Train
)

const (
	inf = 1000000
)

var (
	distanceByTime [][]Trains
	minTime        int = inf
)

type Trains struct {
	trs []Train
}

type path struct {
	price float32
	time  int
	path  string
}

func InitTSP(inf *Info) {

	info = inf
	list = info.GetListOfCity()
	numOfStation = len(list)

	distanceByPrice = make([][]float32, numOfStation)
	for i := range distanceByPrice {
		distanceByPrice[i] = make([]float32, numOfStation)
	}
	for i := 0; i < numOfStation; i++ {
		for j := 0; j < numOfStation; j++ {
			tr := Train{}
			distanceByPrice[i][j], tr = info.GetMinPrice(list[i], list[j])
			cheapTrains = append(cheapTrains, tr)
		}
	}

	distanceByTime = make([][]Trains, numOfStation)
	for i := range distanceByTime {
		distanceByTime[i] = make([]Trains, numOfStation)
	}
	for i := 0; i < numOfStation; i++ {
		for j := 0; j < numOfStation; j++ {
			trs := Trains{}
			trs.trs = info.GetTrains(list[i], list[j])
			distanceByTime[i][j] = trs
		}
	}
}

func Do() {
	info.ShowListOfCity()
	i := 0
	j := 0
	indexes = make([]int, len(list))
	indexesByTime = make([]Train, len(list)-1)

	recursion(i, j)
	fmt.Println("best path = ", bestPath)
	fmt.Println("price of the best path = ", bestPrice)
	fmt.Println("time in way = ", info.minInDays(minTime))
	fmt.Println("attempt number = ", minCounter)

}

func recursion(i, j int) {
	for indexes[i] = 0; indexes[i] < len(list); indexes[i]++ {

		if i < numOfStation-1 {
			recursion(i+1, j)
		}

		for index := range indexes {
			if indexes[index] >= len(list) {
				indexes[index]--
			}
		}

		if indexes[i] >= numOfStation {
			indexes[i]--
		}
		bufferAccess := true
		for y := 0; y < len(list)-1; y++ {
			for x := y; x < len(list)-1; x++ {
				if indexes[y] == indexes[x+1] {
					bufferAccess = false
					break
				}
			}
			if !bufferAccess {
				break
			}
		}

		if bufferAccess {
			if switcher == 1 {
				findTheBestPrice()
			} else if switcher == 2 {
				findTheBestTime(j)
			}
		}
	}
}
