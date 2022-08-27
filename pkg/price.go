package pkg

import "fmt"

func findTheBestPrice() {
	pt := path{}
	bufferInnerAccess := true
	for k := 0; k < len(list)-1; k++ {
		if findNumberOfTrain(list[indexes[k]], list[indexes[k+1]]) == inf {
			bufferInnerAccess = false
			break
		}

	}
	if !bufferInnerAccess {
		return
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
}

func findNumberOfTrain(departures, arrival int) int {
	for _, st := range cheapTrains {
		if st.ArrivalStation == arrival && st.DeparturesStation == departures {
			return st.NumOfTrain
		}
	}
	return inf
}

func findTrainByPrice(list []Train, price float32, departures, arrival int) Train {
	for _, tr := range list {
		if tr.Price == price && tr.ArrivalStation == arrival && tr.DeparturesStation == departures {
			return tr
		}
	}
	return Train{}
}
