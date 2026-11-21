package currency

import (
	"encoding/json"
	"os"
)

func SaveToJSON(results []CurrencyResult, path string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
