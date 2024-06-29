package BackengHelpersGo

import (
	"encoding/json"
)

func StringFromObj(obj any) string{
	msgToPrint, _ := json.MarshalIndent(obj, "", "    ")
	return string(msgToPrint)
}