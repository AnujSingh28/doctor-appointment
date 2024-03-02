package main

import (
	"doc-appointment/controller"
	"fmt"
	"log"
)

func main(){
	log.Println("Welcome to App")
	var err error
	for true{
		var command1 string
		log.Println("are you doctor or patient? type exit to end")
		fmt.Scanln(&command1)
		if command1 == "exit"{
			break
		}
		if command1 == "doctor" {
			log.Println("register, search appointment")
			if command1 == "register" {
				var name, specialization string
				var rating int
				fmt.Scanln(&name, &specialization, &rating)
				err = controller.RegisterDoctor(name, specialization, rating)
				if err != nil {
					log.Println(err.Error())
				} else{
					log.Println("Doctor registered")
				}
				continue
			} else if command1 == "search" {

			}
		} else {

		}
	}
}
