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

	switcher = 1

	indexes       []int
	indexesByTime []Train

	localTime int
	max       int
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
	fmt.Println(bestPath)
	fmt.Println(bestPrice)
	fmt.Println(info.minInDays(minTime))
	fmt.Println(minCounter)
	fmt.Println(info.minInDays(localTime))
	fmt.Println(max)
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
				pt := path{}
				bufferInnerAccess := true
				for k := 0; k < len(list)-1; k++ {
					if findNumberOfTrain(list[indexes[k]], list[indexes[k+1]]) == inf {
						bufferInnerAccess = false
						break
					}

				}
				if !bufferInnerAccess {
					continue
				}

				for k := 0; k < len(list); k++ {
					if k < len(list)-1 {
						pt.path += fmt.Sprint(list[indexes[k]], " --", findNumberOfTrain(list[indexes[k]], list[indexes[k+1]]), "-> ")
					} else {
						pt.path += fmt.Sprint(list[indexes[k]])
					}
				}

				for k := 0; k < len(list)-1; k++ {
					pt.price += distanceByPrice[indexes[k]][indexes[k+1]]
				}

				if pt.price < minPath {
					minPath = pt.price
					minCounter = counter
					bestPath = pt.path
					pt.price = minPath
					//				localTime = GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[0]][indexes[1]], list[0], list[1])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[indexes[0]][indexes[1]], list[0], list[1]), findTrainByPrice(cheapTrains, distanceByPrice[indexes[1]][indexes[2]], list[1], list[2])) + GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[1]][indexes[2]], list[1], list[2])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[indexes[1]][indexes[2]], list[1], list[2]), findTrainByPrice(cheapTrains, distanceByPrice[indexes[2]][indexes[3]], list[2], list[3])) + GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[2]][indexes[3]], list[2], list[3])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[indexes[2]][indexes[3]], list[2], list[3]), findTrainByPrice(cheapTrains, distanceByPrice[indexes[3]][indexes[4]], list[3], list[4])) + GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[3]][indexes[4]], list[3], list[4])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[indexes[3]][indexes[4]], list[3], list[4]), findTrainByPrice(cheapTrains, distanceByPrice[indexes[4]][indexes[5]], list[4], list[5])) + GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[4]][indexes[5]], list[4], list[5]))

					for k := 0; k < len(list)-1; k++ {

						if k < len(list)-2 {
							pt.time += GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[k]][indexes[k+1]], list[indexes[k]], list[indexes[k+1]])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[indexes[k]][indexes[k+1]], list[indexes[k]], list[indexes[k+1]]), findTrainByPrice(cheapTrains, distanceByPrice[indexes[k+1]][indexes[k+2]], list[indexes[k+1]], list[indexes[k+2]]))
						} else {
							pt.time += GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[k]][indexes[k+1]], list[indexes[k]], list[indexes[k+1]]))
						}

					}

					minTime = pt.time
					bestPrice = minPath
					roads = append(roads, pt)
				}
				counter++
			} else if switcher == 2 {
				bestTime(j)

			}
		}
	}
}

func findNumberOfTrain(departures, arrival int) int {
	for _, st := range cheapTrains {
		if st.ArrivalStation == arrival && st.DeparturesStation == departures {
			return st.NumOfTrain
		}
	}
	return inf
}

func bestTime(i int) {

	for _, indexesByTime[i] = range distanceByTime[indexes[i]][indexes[i+1]].trs {

		if i < len(indexesByTime)-1 {
			bestTime(i + 1)
		}

		max++
		bufferInnerAccess := true
		for k := 0; k < len(list)-1; k++ {
			if findNumberOfTrain(list[indexes[k]], list[indexes[k+1]]) == inf {
				bufferInnerAccess = false
				break
			}

		}
		if !bufferInnerAccess {
			continue
		}

		pt := path{}
		checkTrain := true
		for k := 0; k < len(list); k++ {
			if k < len(list)-1 {
				if indexesByTime[k].NumOfTrain == 0 {
					checkTrain = false
					break
				}
				pt.path += fmt.Sprint(list[indexes[k]], " --", indexesByTime[k].NumOfTrain, "-> ")
			} else {
				pt.path += fmt.Sprint(list[indexes[k]])
			}
		}

		if !checkTrain {
			continue
		}

		//pt.time = GetTimeInWay(indexesByTime[0]) + GetTimeInStation(indexesByTime[0], indexesByTime[1]) + GetTimeInWay(indexesByTime[1]) + GetTimeInStation(indexesByTime[1], indexesByTime[2]) + GetTimeInWay(indexesByTime[2]) + GetTimeInStation(indexesByTime[2], indexesByTime[3]) + GetTimeInWay(indexesByTime[3]) + GetTimeInStation(indexesByTime[3], indexesByTime[4]) + GetTimeInWay(indexesByTime[4])

		//Исправить этот цикл
		//for k := 0; k < len(list)-1; k++ {
		//	pt.time += GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[k]][indexes[k+1]]))
		//	if k != len(list)-2 {
		//		pt.time += GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[indexes[k]][indexes[k+1]]), findTrainByPrice(cheapTrains, distanceByPrice[k+1][indexes[k+2]]))
		//	}
		//}

		//pt.path = fmt.Sprint(list[indexes[0]], " --", indexesByTime[0].NumOfTrain, "-> ", list[indexes[1]], " --", indexesByTime[1].NumOfTrain, "-> ", list[indexes[2]], " --", indexesByTime[2].NumOfTrain, "-> ", list[indexes[3]], " --", indexesByTime[3].NumOfTrain, "-> ", list[indexes[4]], " --", indexesByTime[4].NumOfTrain, "-> ", list[indexes[5]])

		if pt.time < minTime {
			minTime = pt.time
			bestPrice = indexesByTime[0].Price + indexesByTime[1].Price + indexesByTime[2].Price + indexesByTime[3].Price + indexesByTime[4].Price
			minCounter = counter
			bestPath = pt.path

			roads = append(roads, pt)
		}
		counter++
	}

}

func findTrainByPrice(list []Train, price float32, departures, arrival int) Train {
	for _, tr := range list {
		if tr.Price == price && tr.ArrivalStation == arrival && tr.DeparturesStation == departures {
			return tr
		}
	}
	return Train{}
}
