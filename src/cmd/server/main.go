package main

import (
	"github.com/DrDowellsHead/Go_Bootcamp_4/src/di"
	"go.uber.org/fx"
)

func main() {
	fx.New(di.Module).Run()
}
