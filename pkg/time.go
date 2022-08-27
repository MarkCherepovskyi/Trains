package pkg

import "fmt"

func findTheBestTime(i int) {

	for _, indexesByTime[i] = range distanceByTime[indexes[i]][indexes[i+1]].trs {

		if i < len(indexesByTime)-1 {
			findTheBestTime(i + 1)
		}

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

		for k := 0; k < len(list)-1; k++ {

			if k < len(list)-2 {
				pt.time += GetTimeInWay(indexesByTime[k]) + GetTimeInStation(indexesByTime[k], indexesByTime[k+1])
			} else {
				pt.time += GetTimeInWay(indexesByTime[k])
			}

		}

		if pt.time < minTime {
			minTime = pt.time

			for k := 0; k < len(list)-1; k++ {
				pt.price += indexesByTime[k].Price
			}

			minCounter = counter
			bestPath = pt.path
			bestPrice = pt.price
			roads = append(roads, pt)
		}
		counter++
	}

}
