package progog

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"
)

func exportBlockPoolToProlog() string {
	var output []string

	output = append(output, "% Bitcoin Blockchain Knowledge Base")
	output = append(output, "% Auto-generated using progog")
	output = append(output, "")

	var facts []string
	for _, block := range BlockPool {
		if block == nil {
			continue
		}

		val := reflect.ValueOf(*block)
		typ := reflect.TypeOf(*block)

		var blockHash string
		for i := 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Name
			if strings.ToLower(fieldName) == "hash" {
				blockHash = fmt.Sprintf("'%v'", val.Field(i).Interface())
				break
			}
		}

		if blockHash == "" {
			continue
		}

		
		for i := 0; i < val.NumField(); i++ {
			fieldName := typ.Field(i).Name
			fieldValue := val.Field(i).Interface()

			prologFieldName := strings.ToLower(fieldName)

			var formattedValue string
			switch v := fieldValue.(type) {
			case string:
				if prologFieldName == "hash" {
					facts = append(facts, fmt.Sprintf("block_hash(%s).", blockHash))
					continue
				} else {
					formattedValue = fmt.Sprintf("'%s'", v)
				}
			case []string:
				if len(v) == 0 {
					formattedValue = "[]"
				} else {
					var quotedItems []string
					for _, item := range v {
						quotedItems = append(quotedItems, fmt.Sprintf("'%s'", item))
					}
					formattedValue = fmt.Sprintf("[%s]", strings.Join(quotedItems, ", "))
				}
			default:
				formattedValue = fmt.Sprintf("%v", v)
			}

			// Create facts in the format: field_name(BlockHash, Value)
			prologFact := fmt.Sprintf("block_%s(%s, %s).", prologFieldName, blockHash, formattedValue)
			facts = append(facts, prologFact)
		}
	}

	slices.Sort(facts)
	output = append(output, facts...)
	output = append(output, "")

	return strings.Join(output, "\n")
}

func ExportToKB(filename string) error {
	data := exportBlockPoolToProlog()
	return os.WriteFile(filename, []byte(data), 0644)
}
