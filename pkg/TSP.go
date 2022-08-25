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
	bestPath        string
	bestPrice       float32
	distanceByPrice [][]float32
	cheapTrains     []Train

	switcher = 2
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

	for i1 := 0; i1 < 6; i1++ {
		for i2 := 0; i2 < 6; i2++ {
			for i3 := 0; i3 < 6; i3++ {
				for i4 := 0; i4 < 6; i4++ {
					for i5 := 0; i5 < 6; i5++ {
						for i6 := 0; i6 < 6; i6++ {
							if i1 != i2 && i1 != i3 && i1 != i4 && i1 != i5 && i1 != i6 && i2 != i3 && i2 != i4 && i2 != i5 && i2 != i6 && i3 != i4 && i3 != i5 && i3 != i6 && i4 != i5 && i4 != i6 && i5 != i6 {
								if switcher == 1 {
									pt := path{}
									pt.path = fmt.Sprint(list[i1], " -", findNumberOfTrain(list[i1], list[i2]), "> ", list[i2], " -", findNumberOfTrain(list[i2], list[i3]), "> ", list[i3], " -", findNumberOfTrain(list[i3], list[i4]), "> ", list[i4], " -", findNumberOfTrain(list[i4], list[i5]), "> ", list[i5], " -", findNumberOfTrain(list[i5], list[i6]), "> ", list[i6])
									if findNumberOfTrain(list[i1], list[i2]) == inf || findNumberOfTrain(list[i2], list[i3]) == inf || findNumberOfTrain(list[i3], list[i4]) == inf || findNumberOfTrain(list[i4], list[i5]) == inf || findNumberOfTrain(list[i5], list[i6]) == inf {
										continue
									}
									if distanceByPrice[i1][i2]+distanceByPrice[i2][i3]+distanceByPrice[i3][i4]+distanceByPrice[i4][i5]+distanceByPrice[i5][i6] < minPath {
										minPath = distanceByPrice[i1][i2] + distanceByPrice[i2][i3] + distanceByPrice[i3][i4] + distanceByPrice[i4][i5] + distanceByPrice[i5][i6]
										minCounter = counter
										bestPath = pt.path
										pt.price = minPath
										minTime = GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[i1][i2])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[i1][i2]), findTrainByPrice(cheapTrains, distanceByPrice[i2][i3])) + GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[i2][i3])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[i2][i3]), findTrainByPrice(cheapTrains, distanceByPrice[i3][i4])) + GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[i3][i4])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[i3][i4]), findTrainByPrice(cheapTrains, distanceByPrice[i4][i5])) + GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[i4][i5])) + GetTimeInStation(findTrainByPrice(cheapTrains, distanceByPrice[i4][i5]), findTrainByPrice(cheapTrains, distanceByPrice[i5][i6])) + GetTimeInWay(findTrainByPrice(cheapTrains, distanceByPrice[i5][i6]))
										bestPrice = minPath
										roads = append(roads, pt)
									}
									counter++
								} else if switcher == 2 {
									for _, tr1 := range distanceByTime[i1][i2].trs {
										for _, tr2 := range distanceByTime[i2][i3].trs {
											for _, tr3 := range distanceByTime[i3][i4].trs {
												for _, tr4 := range distanceByTime[i4][i5].trs {
													for _, tr5 := range distanceByTime[i5][i6].trs {
														time := GetTimeInWay(tr1) + GetTimeInStation(tr1, tr2) + GetTimeInWay(tr2) + GetTimeInStation(tr2, tr3) + GetTimeInWay(tr3) + GetTimeInStation(tr3, tr4) + GetTimeInWay(tr4) + GetTimeInStation(tr4, tr5) + GetTimeInWay(tr5)
														pt := path{}
														pt.path = fmt.Sprint(list[i1], " -", tr1.NumOfTrain, "> ", list[i2], " -", tr2.NumOfTrain, "> ", list[i3], " -", tr3.NumOfTrain, "> ", list[i4], " -", tr4.NumOfTrain, "> ", list[i5], " -", tr5.NumOfTrain, "> ", list[i6])
														if findNumberOfTrain(list[i1], list[i2]) == inf || findNumberOfTrain(list[i2], list[i3]) == inf || findNumberOfTrain(list[i3], list[i4]) == inf || findNumberOfTrain(list[i4], list[i5]) == inf || findNumberOfTrain(list[i5], list[i6]) == inf {
															continue
														}

														if time < minTime {
															minTime = time
															bestPrice = tr1.Price + tr2.Price + tr3.Price + tr4.Price + tr5.Price
															minCounter = counter
															bestPath = pt.path
															pt.time = time

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
						}
					}
				}
			}
		}
	}

	fmt.Println(bestPath)
	fmt.Println(bestPrice)
	fmt.Println(info.minInDays(minTime))
	fmt.Println(minCounter)
}

func findNumberOfTrain(departures, arrival int) int {
	for _, st := range cheapTrains {
		if st.ArrivalStation == arrival && st.DeparturesStation == departures {
			return st.NumOfTrain
		}
	}
	return inf
}

func findTrainByPrice(list []Train, price float32) Train {
	for _, tr := range list {
		if tr.Price == price {
			return tr
		}
	}
	return Train{}
}
