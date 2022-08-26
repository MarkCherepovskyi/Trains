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
	indexesByTime = make([]Train, len(list)-2)
	//for i1 := 0; i1 < 6; i1++ {
	//	for i2 := 0; i2 < 6; i2++ {
	//		for i3 := 0; i3 < 6; i3++ {
	//			for i4 := 0; i4 < 6; i4++ {
	//				for i5 := 0; i5 < 6; i5++ {
	//					for i6 := 0; i6 < 6; i6++ {
	//						if i1 != i2 && i1 != i3 && i1 != i4 && i1 != i5 && i1 != i6 && i2 != i3 && i2 != i4 && i2 != i5 && i2 != i6 && i3 != i4 && i3 != i5 && i3 != i6 && i4 != i5 && i4 != i6 && i5 != i6 {
	//
	//
	//
	//						}
	//					}
	//				}
	//			}
	//		}
	//	}
	//}
	recursion(i, j)
	fmt.Println(bestPath)
	fmt.Println(bestPrice)
	fmt.Println(info.minInDays(minTime))
	fmt.Println(minCounter)
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
				//pt.path = fmt.Sprint(list[indexes[0]], " --", findNumberOfTrain(list[indexes[0]], list[indexes[1]]), "-> ", list[indexes[1]], " --", findNumberOfTrain(list[indexes[1]], list[indexes[2]]), "-> ", list[indexes[2]], " --", findNumberOfTrain(list[indexes[2]], list[indexes[3]]), "-> ", list[indexes[3]], " --", findNumberOfTrain(list[indexes[3]], list[indexes[4]]), "-> ", list[indexes[4]], " --", findNumberOfTrain(list[indexes[4]], list[indexes[5]]), "-> ", list[indexes[5]])
				if findNumberOfTrain(list[indexes[0]], list[indexes[1]]) == inf || findNumberOfTrain(list[indexes[1]], list[indexes[2]]) == inf || findNumberOfTrain(list[indexes[2]], list[indexes[3]]) == inf || findNumberOfTrain(list[indexes[3]], list[indexes[4]]) == inf || findNumberOfTrain(list[indexes[4]], list[indexes[5]]) == inf {
					continue
				}

				for k := 0; k < len(list)-1; k++ {
					if k != len(list)-2 {
						pt.path += fmt.Sprint(list[indexes[k]], " --", findNumberOfTrain(list[indexes[k]], list[indexes[k+1]]), "-> ")
					} else {
						pt.path += fmt.Sprint(list[indexes[k]])
					}
				}

				if distanceByPrice[indexes[0]][indexes[1]]+distanceByPrice[indexes[1]][indexes[2]]+distanceByPrice[indexes[2]][indexes[3]]+distanceByPrice[indexes[3]][indexes[4]]+distanceByPrice[indexes[4]][indexes[5]] < minPath {
					minPath = distanceByPrice[indexes[0]][indexes[1]] + distanceByPrice[indexes[1]][indexes[2]] + distanceByPrice[indexes[2]][indexes[3]] + distanceByPrice[indexes[3]][indexes[4]] + distanceByPrice[indexes[4]][indexes[5]]
					minCounter = counter
					bestPath = pt.path
					pt.price = minPath

					for k := 0; k < len(list)-1; k++ {
						if k != len(list)-2 {
							pt.time += GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[k]][indexes[k+1]])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[indexes[k]][k+1]), findTrainByPrice(cheapTrains, distanceByPrice[k+1][k+2]))
						} else {
							pt.time += GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[indexes[k]][indexes[k+1]]))
						}
					}
					minTime = pt.time
					bestPrice = minPath
					roads = append(roads, pt)
				}
				counter++
			} else if switcher == 2 {
				bestTime(j + 1)
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

/////////////////
//вот здесь ошибка
func bestTime(i int) {
	if i < len(indexesByTime)-3 {
		for _, indexesByTime[i] = range distanceByTime[indexes[i]][indexes[i+1]].trs {
			time := GetTimeInWay(indexesByTime[0]) + GetTimeInStation(indexesByTime[0], indexesByTime[1]) + GetTimeInWay(indexesByTime[1]) + GetTimeInStation(indexesByTime[1], indexesByTime[2]) + GetTimeInWay(indexesByTime[2]) + GetTimeInStation(indexesByTime[2], indexesByTime[3]) + GetTimeInWay(indexesByTime[3]) + GetTimeInStation(indexesByTime[3], indexesByTime[4]) + GetTimeInWay(indexesByTime[4])
			pt := path{}
			pt.path = fmt.Sprint(list[indexes[0]], " --", indexesByTime[0].NumOfTrain, "-> ", list[indexes[1]], " --", indexesByTime[1].NumOfTrain, "-> ", list[indexes[2]], " --", indexesByTime[2].NumOfTrain, "-> ", list[indexes[3]], " --", indexesByTime[3].NumOfTrain, "-> ", list[indexes[4]], " --", indexesByTime[4].NumOfTrain, "-> ", list[indexes[5]])
			if findNumberOfTrain(list[indexes[0]], list[indexes[1]]) == inf || findNumberOfTrain(list[indexes[1]], list[indexes[2]]) == inf || findNumberOfTrain(list[indexes[2]], list[3]) == inf || findNumberOfTrain(list[indexes[3]], list[indexes[4]]) == inf || findNumberOfTrain(list[indexes[4]], list[indexes[5]]) == inf {
				continue
			}

			if time < minTime {
				minTime = time
				bestPrice = indexesByTime[0].Price + indexesByTime[1].Price + indexesByTime[2].Price + indexesByTime[3].Price + indexesByTime[4].Price
				minCounter = counter
				bestPath = pt.path
				pt.time = time
				roads = append(roads, pt)
			}
			counter++
		}
	}
}

func findTrainByPrice(list []Train, price float32) Train {
	for _, tr := range list {
		if tr.Price == price {
			return tr
		}
	}
	return Train{}
}
