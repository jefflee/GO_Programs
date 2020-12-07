package main

func main() {

	/*
	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	newEmployee := employee{0, "Lloyd", "Christmas"}

	fmt.Printf("The type is %v\n", reflect.TypeOf(newEmployee))
	fmt.Printf("The kind is %v\n", reflect.TypeOf(newEmployee).Kind())
	fmt.Printf("The value is %v\n", reflect.ValueOf(newEmployee))
	*/

	/*
	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	employees := make([]employee, 3)

	employees = append(employees, employee{0, "Lloyd", "Christmas"})
	employees = append(employees, employee{1, "Harry", "Dunne"})
	employees = append(employees, employee{0, "Sea", "Bass"})

	fmt.Printf("The name of the type is %v\n", reflect.TypeOf(employees).Name())
	fmt.Printf("The type is %v\n", reflect.TypeOf(employees))
	fmt.Printf("The kind is %v\n", reflect.TypeOf(employees).Kind())
	fmt.Printf("The value is %v\n", reflect.ValueOf(employees))
	*/

	/*
	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	employees := make([]employee, 3)

	employees = append(employees, employee{0, "Lloyd", "Christmas"})
	employees = append(employees, employee{1, "Harry", "Dunne"})
	employees = append(employees, employee{2, "Sea", "Bass"})

	eType := reflect.TypeOf(employees)

	newEmployeeList := reflect.MakeSlice(eType, )0, 0)
	newEmployeeList = reflect.Append(newEmployeeList, reflect.ValueOf(employee{1, "test", "test"}))

	fmt.Printf("First list of employees: %v\n\n", employees)
	fmt.Printf("List created by reflection: %v\n", newEmployeeList)
	*/




}
