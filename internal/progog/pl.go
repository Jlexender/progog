package progog

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func exportBlockPoolToProlog() string {
	var output []string

	output = append(output, "% Bitcoin Blockchain Knowledge Base")
	output = append(output, "% Auto-generated using progog")
	output = append(output, "")

	for _, block := range BlockPool {
		val := reflect.ValueOf(block)
		typ := reflect.TypeOf(block)

		var facts []string
		for i := 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Name
			fieldValue := val.Field(i).Interface()
			// Convert field name to lowercase for Prolog style
			prologFact := fmt.Sprintf("%s(%v, %v).", strings.ToLower(typ.Name()), strings.ToLower(fieldName), fieldValue)
			facts = append(facts, prologFact)
		}
		output = append(output, strings.Join(facts, "\n"))
		output = append(output, "")
	}

	return strings.Join(output, "\n")
}

func ExportToKB(filename string) error {
	data := exportBlockPoolToProlog()
	return os.WriteFile(filename, []byte(data), 0644)
}
