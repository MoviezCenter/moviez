package main

import (
	"fmt"

	"github.com/MoviezCenter/moviez/cmd"
	"github.com/MoviezCenter/moviez/config"
)

func main() {
	cmd.Execute()
	fmt.Printf("%+v", config.AppConfigInstance)
}
