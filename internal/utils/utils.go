package utils

import (
	"fmt"
	"strings"

	"os"
)

type HeadersArray [][]string

func (ha *HeadersArray) String() string {
	return fmt.Sprintf("%v", *ha)
}
func (ha *HeadersArray) Set(value string) error {
	headerParts := strings.SplitN(value, ": ", 2)

	if len(headerParts) == 2 {
		key := strings.TrimSpace(headerParts[0])
		value := strings.TrimSpace(headerParts[1])

		*ha = append(*ha, []string{key, value})
	} else {
		return fmt.Errorf("invalid header format: %v", value)
	}

	return nil
}

func FileWriter(
	filePath string,
	contentType string,
	data []byte,
) string {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Sprintf("Error writing to file: %v", err)
	}
	return fmt.Sprintf("Response body written to file: %v", filePath)
}
