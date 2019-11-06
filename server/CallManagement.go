package server

import (
	"../database"
	"log"
	"time"
)

func AssignToEmployee()  {
	var phoneNumber string
	for true {
		time.Sleep(30 *time.Second)

		phoneNumber = database.SelectFirstCall()
		log.Print(phoneNumber)

		if phoneNumber != "" {
			employee := database.SelectFreeUsers("respondent")
			if employee != "" {
				log.Print(employee, " answered ", phoneNumber)
			} else {
				employee = database.SelectFreeUsers("manager")
				if employee != "" {
					log.Print(employee, " answered ", phoneNumber)
				} else {
					employee = database.SelectFreeUsers("director")
					if employee != "" {
						log.Print(employee, " answered ", phoneNumber)
					} else {
						log.Print("nobody is free, please free one user")
						database.InsertNewCall(phoneNumber)
					}

				}

			}

		} else {
			log.Print("we have no incoming call")
		}
	}

}
