package ssdeep_test

import (
	"fmt"
	"github.com/DanielPels/ssdeep"
	"log"
	"math/rand"
	"os"
)

func ExampleFuzzyFilename() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h, err := ssdeep.FuzzyFile(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h)
}

func ExampleFuzzyBytes() {
	buffer := make([]byte, 4097)
	rand.Read(buffer)
	h, err := ssdeep.FuzzyBytes(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(h)
}
