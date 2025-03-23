package internal

import (
	"fmt"
	"os"
	"sort"
	"strings"
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
	sections := GetRegisteredSections()

	sort.Slice(sections, func(i, j int) bool {
		return sections[i].Order() < sections[j].Order()
	})

	var content strings.Builder
	content.WriteString("# Don Athalage\n\n")

	for _, section := range sections {
		sectionContent, err := section.Generate()
		if err != nil {
			return "", fmt.Errorf("error generating section '%s': %w", section.Name(), err)
		}

		content.WriteString("## " + section.Name() + "\n\n")
		content.WriteString(sectionContent + "\n\n")
	}

	return content.String(), nil
}
