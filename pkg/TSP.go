package pkg

import (
	"fmt"
	"sync"
)

var (
	firstStation int
	list         []int
	numOfStation int
	info         *Info
	roads        []path
	minPath      float32 = 10000
	counter      int     = 0
	minCounter   int
	bestPath     string
	bestPrice    float32
	distance     [][]float32
	cheapTrains  []Train
)

type path struct {
	price float32
	time  int
	path  string
}

func InitTSP(inf *Info) {
	info = inf
	list = info.GetListOfCity()
	numOfStation = len(list)

	distance = make([][]float32, numOfStation)
	for i := range distance {
		distance[i] = make([]float32, numOfStation)
	}
	for i := 0; i < numOfStation; i++ {
		for j := 0; j < numOfStation; j++ {
			tr := Train{}
			distance[i][j], tr = info.GetMinPrice(list[i], list[j])
			cheapTrains = append(cheapTrains, tr)
		}
	}
}

func Do(wt *sync.WaitGroup) {

	for i1 := 0; i1 < 6; i1++ {
		for i2 := 0; i2 < 6; i2++ {
			for i3 := 0; i3 < 6; i3++ {
				for i4 := 0; i4 < 6; i4++ {
					for i5 := 0; i5 < 6; i5++ {
						for i6 := 0; i6 < 6; i6++ {
							if i1 != i2 && i1 != i3 && i1 != i4 && i1 != i5 && i1 != i6 && i2 != i3 && i2 != i4 && i2 != i5 && i2 != i6 && i3 != i4 && i3 != i5 && i3 != i6 && i4 != i5 && i4 != i6 && i5 != i6 {
								pt := path{}
								pt.path = fmt.Sprint(list[i1], " -", findNumberOfTrain(list[i1], list[i2]), "> ", list[i2], " -", findNumberOfTrain(list[i2], list[i3]), "> ", list[i3], " -", findNumberOfTrain(list[i3], list[i4]), "> ", list[i4], " -", findNumberOfTrain(list[i4], list[i5]), "> ", list[i5], " -", findNumberOfTrain(list[i5], list[i6]), "> ", list[i6])
								if findNumberOfTrain(list[i1], list[i2]) == inf || findNumberOfTrain(list[i2], list[i3]) == inf || findNumberOfTrain(list[i3], list[i4]) == inf || findNumberOfTrain(list[i4], list[i5]) == inf || findNumberOfTrain(list[i5], list[i6]) == inf {
									continue
								}
								if distance[i1][i2]+distance[i2][i3]+distance[i3][i4]+distance[i4][i5]+distance[i5][i6] < minPath {
									minPath = distance[i1][i2] + distance[i2][i3] + distance[i3][i4] + distance[i4][i5] + distance[i5][i6]
									minCounter = counter
									bestPath = pt.path
									pt.price = minPath
									bestPrice = minPath
									roads = append(roads, pt)
								}
								counter++
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(bestPath)
	fmt.Println(bestPrice)
	wt.Done()
}

func findNumberOfTrain(departures, arrival int) int {
	for _, st := range cheapTrains {
		if st.ArrivalStation == arrival && st.DeparturesStation == departures {
			return st.NumOfTrain
		}
	}
	return inf
}
