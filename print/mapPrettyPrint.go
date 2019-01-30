package print

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(m interface{}) string {
	recordString, error := json.MarshalIndent(m, "", "\t")
	if error == nil {
		return fmt.Sprintf("%s", recordString)
	} else {
		return fmt.Sprintf("%+v", m)
	}
}
