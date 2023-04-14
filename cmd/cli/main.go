package main

import (
	"log"
	"os"

	"github.com/nacho692/live-free-or-die-jugging/pkg/app"
	"github.com/nacho692/live-free-or-die-jugging/pkg/iterative"
)

func main() {

	application, err := app.New(app.Configuration{
		Output: os.Stdout,
		Input:  os.Stdin,
		Solver: app.SolverFun(iterative.Solve),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
