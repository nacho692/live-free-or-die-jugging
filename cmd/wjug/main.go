package main

import (
	"flag"
	"log"
	"os"

	"github.com/nacho692/live-free-or-die-jugging/pkg/app"
	"github.com/nacho692/live-free-or-die-jugging/pkg/iterative"
)

func main() {

	silent := flag.Bool("s", false, "silences most output so only the solution is printed")
	flag.Parse()

	application, err := app.New(app.Configuration{
		Output: os.Stdout,
		Silent: *silent,
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
