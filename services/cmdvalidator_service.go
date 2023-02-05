package services

import (
	"errors"
	"path/filepath"
	"strings"
)

type CMDValidatorService struct {
}

func NewCMDValidatorService() *CMDValidatorService {
	return new(CMDValidatorService)
}

func (v *CMDValidatorService) ArgumentMustBeOne(args []string) error {
	if len(args) == 1 {
		return nil
	} else {
		return errors.New("Argument must be one")
	}
}

func (v *CMDValidatorService) MustLogFile(source string) error {
	if filepath.Ext(source) == ".log" {
		return nil
	} else {
		return errors.New("File is not a valid log")
	}
}

func (v *CMDValidatorService) TypeMustJsonOrTxt(typ string) error {
	if strings.Contains(typ, "json") || strings.Contains(typ, "txt") || strings.Contains(typ, "text") {
		return nil
	} else {
		return errors.New("Type must be a json or txt (text)")
	}
}
