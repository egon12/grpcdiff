package bodygenerator

import (
	"log"
	"os"
)

func Main() {
	inputSource := os.Args[1]
	inputTmpl := os.Args[2]

	s := &SeperatedValueSource{}
	s.Parse(inputSource)

	err := Generate(s, inputTmpl, os.Stdout)
	if err != nil {
		log.Println(err)
	}
}
