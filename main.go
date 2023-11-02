package main

import (
	"github.com/Bobby-P-dev/testingcyclic/database"
	"github.com/Bobby-P-dev/testingcyclic/initiallizers"
	"github.com/Bobby-P-dev/testingcyclic/routers"
)

func init() {
	initiallizers.LoadEnvVar()
	database.ConnectToDB()
}

func main() {
	database.ConnectToDB()

	r := routers.RouterA()
	r.Run()
}
