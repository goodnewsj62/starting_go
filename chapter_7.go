package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct{
	plate_no string
	color string
	model int
}


type Truck struct{
	Car
	wheels string
}


func main(){
	//structs 
	var car Car

	car.color = "red"
	car.model = 1195
	car.plate_no = "190245FRT"

	fmt.Println("car_color:",car.color,"car_model:",car.model,"car_plate_no:",car.plate_no)
	c := Car{"1dft568KJB","red",1997}

	c.prtCarInfo()
	x := Truck{Car{plate_no:"jiG3458"},"8"}
	reWorkVehicle(&c,&x)
	c.prtCarInfo()
	x.prtCarInfo()

}

type Vehicle interface{
	upgrade() string
	repaint() string
}

// this way you can create a method in go quite easily by just 
// using a variable and pointer before the function name(reciever)
func (c *Car) prtCarInfo() {
	fmt.Println("color:",c.color,"model:",c.model,"plate_no:",c.plate_no)
}

func (c *Car) upgrade() string{
	rand.Seed(time.Now().UnixNano())
	c.model += rand.Intn(200)
	return "vehicle model now: " + string(c.model)
}
func (c *Car) repaint() string{
	c.color  = "orange"
	return "vehicle color now: " + c.color
}


func reWorkVehicle(vehicles ...Vehicle){
	for _ , value := range vehicles{
		value.repaint() 
		value.upgrade()
	}	
}
