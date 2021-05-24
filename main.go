package main

import (
	"log"
	"os"

	"github.com/jammy-dodgers/gophenslide/openslide"
)

func main() {
	slide, err := openslide.Open("./test/hello.svs")

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	image, error := slide.ReadRegion(0, 0, 0, 300, 300)

	if error != nil {
		log.Fatal(err.Error())
		panic(error)
	}

	f, err := os.Create("region.png")
	f.Write(image)
}
