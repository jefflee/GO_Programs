package main

import (
	"net/http"

	"github.com/jeff/myservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	// u := models.User{
	// 	ID:        2,
	// 	FirstName: "Tricia",
	// 	LastName:  "McMillan",
	// }
	// fmt.Println(u)
}
