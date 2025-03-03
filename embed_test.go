package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed sakura.jpeg

var sakura []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("sakura_copy.jpeg", sakura, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Println("success copy file")
}
