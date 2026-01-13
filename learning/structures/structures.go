package main

import (
	"fmt"
	"time"
)

type car struct {

	id string
	liters uint8
	kmpl uint16
	name string
	status string
	onProgress bool
	// buildAt time.Time
	buildAt string
}

func (s *car) changeStatus(status string) {
	s.status = status
}


func main() {

	// var myCar car = car{liters: 100, kmpl: 20}
	// // myEngine.kmpl = 100

	// fmt.Println(myCar.kmpl, myCar.liters)
	// var myCar car = car{}
	// fmt.Println(myCar)
	now := time.Now()

	myCar2 := car{id: "1", liters: 200, kmpl: 100, name: "BMW", status: "Good", onProgress: true, buildAt: now.Format("2006-01-02")}
	fmt.Println(myCar2)
	myCar2.changeStatus("changedStatus")

	fmt.Println(myCar2)


}