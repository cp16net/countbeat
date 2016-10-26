package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/cp16net/countbeat/beater"
)

func main() {
	err := beat.Run("countbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
