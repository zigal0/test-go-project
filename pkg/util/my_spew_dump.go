package util

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func MySpew(i interface{}) {
	fmt.Println("My spew dump: ")

	spew.Dump(i)
}
