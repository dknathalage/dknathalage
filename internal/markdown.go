package internal

import (
	"fmt"
	"os"
)

func SaveMarkdown(config *Config, markdown string) error {
	filename := config.Filename

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(markdown)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Printf("Markdown saved to '%s'\n", filename)
	return nil
}

func CheckFileSize(config *Config) error {
	filename := config.Filename

	fileInfo, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("error checking file size: %w", err)
	}

	size := fileInfo.Size()
	fmt.Printf("File '%s' size: %d bytes\n", filename, size)

	const maxSize = 512 * 1024
	if size > maxSize {
		return fmt.Errorf("file '%s' size (%d bytes) exceeds maximum allowed size of %d bytes", filename, size, maxSize)
	}

	return nil
}

func GenerateMarkdown() (string, error) {
	return "# Don Athalage", nil
}
