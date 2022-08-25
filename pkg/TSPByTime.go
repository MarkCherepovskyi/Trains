package pkg

import (
	"fmt"
	"sync"
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

func InitTSP2(inf *Info) {
	info = inf
	list = info.GetListOfCity()
	numOfStation = len(list)

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

func DoTSPbyTime(wt *sync.WaitGroup) {

	for i1 := 0; i1 < 6; i1++ {
		for i2 := 0; i2 < 6; i2++ {
			for i3 := 0; i3 < 6; i3++ {
				for i4 := 0; i4 < 6; i4++ {
					for i5 := 0; i5 < 6; i5++ {
						for i6 := 0; i6 < 6; i6++ {
							if i1 != i2 && i1 != i3 && i1 != i4 && i1 != i5 && i1 != i6 && i2 != i3 && i2 != i4 && i2 != i5 && i2 != i6 && i3 != i4 && i3 != i5 && i3 != i6 && i4 != i5 && i4 != i6 && i5 != i6 {

								for _, tr1 := range distanceByTime[i1][i2].trs {
									for _, tr2 := range distanceByTime[i2][i3].trs {
										for _, tr3 := range distanceByTime[i3][i4].trs {
											for _, tr4 := range distanceByTime[i4][i5].trs {
												for _, tr5 := range distanceByTime[i5][i6].trs {
													time := GetTimeInWay(tr1.DeparturesTime, tr1.ArrivalTime) + GetTimeInStation(tr1, tr2) + GetTimeInWay(tr2.DeparturesTime, tr2.ArrivalTime) + GetTimeInStation(tr2, tr3) + GetTimeInWay(tr3.DeparturesTime, tr3.ArrivalTime) + GetTimeInStation(tr3, tr4) + GetTimeInWay(tr4.DeparturesTime, tr4.ArrivalTime) + GetTimeInStation(tr4, tr5) + GetTimeInWay(tr5.DeparturesTime, tr5.ArrivalTime)
													pt := path{}
													pt.path = fmt.Sprint(list[i1], " -", tr1.NumOfTrain, "> ", list[i2], " -", tr2.NumOfTrain, "> ", list[i3], " -", tr3.NumOfTrain, "> ", list[i4], " -", tr4.NumOfTrain, "> ", list[i5], " -", tr5.NumOfTrain, "> ", list[i6])
													if findNumberOfTrain(list[i1], list[i2]) == inf || findNumberOfTrain(list[i2], list[i3]) == inf || findNumberOfTrain(list[i3], list[i4]) == inf || findNumberOfTrain(list[i4], list[i5]) == inf || findNumberOfTrain(list[i5], list[i6]) == inf {
														continue
													}

													if time < minTime {
														minTime = time
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
	fmt.Println(bestPath)
	fmt.Println(info.minInDays(minTime))
	wt.Done()
}
