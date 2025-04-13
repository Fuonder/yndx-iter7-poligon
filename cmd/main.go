package main

import (
	_ "github.com/Fuonder/yndx-iter7-poligon.git/internal/builder"
	"github.com/Fuonder/yndx-iter7-poligon.git/internal/creatingpatterns/abstractfactory"
	_ "github.com/Fuonder/yndx-iter7-poligon.git/internal/factory"
	_ "github.com/Fuonder/yndx-iter7-poligon.git/internal/prototype"
	"github.com/Fuonder/yndx-iter7-poligon.git/internal/runner"
	_ "github.com/Fuonder/yndx-iter7-poligon.git/internal/singleton"
)

func main() {
	runner.Run(abstractfactory.Example1)

}
