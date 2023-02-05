package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/arfanxn/cli-tool-golang/helpers"
	"github.com/arfanxn/cli-tool-golang/services"
	"github.com/spf13/cobra"
)

type CMDHandler struct {
}

func NewCMDHandler() *CMDHandler {
	return new(CMDHandler)
}

func (handler *CMDHandler) RootHandler(cmd *cobra.Command, args []string) {
	// prepare cmd validator service
	validator := services.NewCMDValidatorService()

	// validate the main argument
	helpers.PrintErrorAndExit(validator.ArgumentMustBeOne(args))

	// Prepare the arguments
	source := args[0]
	sourceExtension := filepath.Ext(source)
	typ := strings.ToLower(cmd.Flag("type").Value.String())
	destination := cmd.Flag("output").Value.String()
	destinationExtension := filepath.Ext(destination)

	// Check if the source file is a log file
	if sourceExtension != ".log" {
		err := errors.New("Invalid log file, log file should be end with \".log\" extension.")
		helpers.PrintErrorAndExit(err)
	}

	// Read the source file and return error if file not found
	bytes, err := ioutil.ReadFile(source)
	if err != nil {
		helpers.PrintErrorAndExit(errors.New("File not found: " + source))
	}

	// check if the type is valid
	helpers.PrintErrorAndExit(validator.TypeMustJsonOrTxt(typ))

	// check if destination extension and type are equals
	switch true {
	case (typ == "txt" || typ == "text") && destinationExtension == "":
		destinationExtension = ".txt"
		break
	case (typ == "json") && destinationExtension == "":
		destinationExtension = ".json"
		break
	case
		((typ == "txt" || typ == "text") && destinationExtension != ".txt") ||
			((typ == "json") && destinationExtension != ".json"):
		err := errors.New(`Type: "` + typ + `" is not equals with output file extension: "` + destinationExtension + `".`)
		helpers.PrintErrorAndExit(err)
		break
	}

	// convert file to json if the requested type is json
	if typ == "json" {
		jsonService := services.NewJSONService()
		bytes, err = jsonService.LogFileToJson(bytes)
		helpers.PrintErrorAndExit(err)
	}

	// Guess the destination
	if destination == "./" || destination == "" {
		currentWorkingDirectory, err := os.Getwd()
		helpers.PrintErrorAndExit(err)
		fileNameWithoutExt := strings.TrimSuffix(filepath.Base(source), filepath.Ext(source))
		destination = currentWorkingDirectory + "/" + fileNameWithoutExt + destinationExtension
	} else {
		destination = filepath.Dir(destination) + "/" + filepath.Base(destination)
	}

	// Save to the destination
	err = helpers.CreateFile(bytes, destination)
	helpers.PrintErrorAndExit(err)

	// Indicate that operation successfully then end the operation with exit code 0
	message := `Successfully retrieved log file from "` + source + `" and saved to "` + destination + `" as "` + typ + `".`
	fmt.Println(message)
	os.Exit(0)

}
