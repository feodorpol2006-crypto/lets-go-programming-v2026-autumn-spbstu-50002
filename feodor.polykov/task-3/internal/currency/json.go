package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveToJSON(filePath string, data []Currency) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	if err := os.WriteFile(filePath, jsonData, 0o644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}
