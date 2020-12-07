package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
		type person struct {
			personId int
			firstName string
			lastName string
		}

		newPerson := person{1, "Fred", "Flintstone"}

		fmt.Printf("Our Person is %s %s with an id of %d\n", newPerson.firstName, newPerson.lastName, newPerson.personId)
		fmt.Printf("The type is %v \n", reflect.TypeOf(newPerson))
		fmt.Printf("The value is %v \n", reflect.ValueOf(newPerson))
		fmt.Printf("The kind is %v \n", reflect.ValueOf(newPerson).Kind())
	*/
	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	type customer struct {
		customerId int
		firstName  string
		lastName   string
		company    string
	}

	newEmployee := employee{0, "Fred", "Flintstone"}

	newCustomer := customer{234, "Barney", "Rubble", "Slate Rock and Gravel"}

	addPerson(newEmployee)
	addPerson(newCustomer)

}

func addPerson(p interface{}) bool{

	if reflect.ValueOf(p).Kind() == reflect.Struct {

		v := reflect.ValueOf(p)

		switch reflect.TypeOf(p).Name() {
		case "employee":

			empSqlString := "INSERT INTO employees (personId, firstName, lastName) VALUES (?,?,?)"
			fmt.Printf("SQL: %s\n", empSqlString)
			fmt.Printf("Added %v\n", v.Field(1))
		case "customer":
			custSqlString := "INSERT INTO customers (customerId, firstName,lastName, company) VALUES (?,?,?,?)"
			fmt.Printf("SQL: %s\n", custSqlString)
			fmt.Printf("Added %v\n", v.Field(1))
		}
		return true
	}else {
		return false
	}
}