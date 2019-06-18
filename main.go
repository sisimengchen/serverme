package main

import (
	"github.com/sisimengchen/serverme/app"
	"github.com/sisimengchen/serverme/configs"
	"github.com/sisimengchen/serverme/models"
	// "fmt"
	// "github.com/casbin/casbin"
)

func main() {
	// e := casbin.NewEnforcer("configs/rbac_model.conf", "configs/rbac_policy.csv")
	// allSubjects := e.GetAllSubjects()
	// allObjects := e.GetAllObjects()
	// allNamedObjects := e.GetAllNamedObjects("p")
	// allActions := e.GetAllActions()
	// allRoles := e.GetAllRoles()
	// sub := "admin" // the user that wants to access a resource.
	// obj := "/" // the resource that is going to be accessed.
	// act := ""  // the operation that the user performs on the resource.
	// fmt.Printf("%v", allSubjects)
	// fmt.Printf("%v", allObjects)
	// fmt.Printf("%v", allNamedObjects)
	// fmt.Printf("%v", allActions)
	// fmt.Printf("%v", allRoles)
	// if e.Enforce(sub, obj, act) == true {
	// 	fmt.Printf("1")
	// } else {
	// 	fmt.Printf("2")
	// }
	configs.Init()
	models.Init()
	router := app.Init()
	router.Run(configs.Viper.GetString("app.addr"))
}
