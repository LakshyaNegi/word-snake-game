package utils

import (
	"encoding/json"
	"os"
)

func Log(data interface{}, name string) {
	file, _ := json.MarshalIndent(data, "", " ")

	_ = os.WriteFile(name+".json", file, 0644)
}
