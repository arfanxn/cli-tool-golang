package services

import (
	"encoding/json"
	"strings"
)

type JSONService struct {
}

func NewJSONService() *JSONService {
	return new(JSONService)
}

func (service *JSONService) LogFileToJson(bytes []byte) ([]byte, error) {
	lines := strings.Split(string(bytes), "\n")
	return json.Marshal(lines)
}
