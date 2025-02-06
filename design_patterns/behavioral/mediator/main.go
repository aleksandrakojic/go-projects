/* Mediator is a behavioral design pattern that reduces coupling between components of a program by making them communicate indirectly, through a special mediator object. */

package main

func main() {
    stationManager := newStationManger()

    passengerTrain := &PassengerTrain{
        mediator: stationManager,
    }
    freightTrain := &FreightTrain{
        mediator: stationManager,
    }

    passengerTrain.arrive()
    freightTrain.arrive()
    passengerTrain.depart()
}