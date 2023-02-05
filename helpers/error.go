package helpers

import (
	"fmt"
	"os"
)

func PrintErrorAndExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
