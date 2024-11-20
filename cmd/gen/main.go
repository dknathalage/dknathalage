package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dknathalage/internal/gen/service_config"
	"gopkg.in/yaml.v3"
)

// LoadServiceConfig reads a YAML file and unmarshals it into a ServiceConfig struct.
func LoadServiceConfig(fileName string) (*service_config.ServiceConfig, error) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file '%s': %w", fileName, err)
	}
	log.Println("Successfully read file")

	config := service_config.ServiceConfig{}
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %w", err)
	}
	log.Println("Successfully unmarshalled YAML")

	return &config, nil
}

func main() {
	// Define a command-line flag for the config file path.
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "", "Path to the service configuration YAML file")
	flag.Parse()

	// Validate the input.
	if configFilePath == "" {
		fmt.Printf("Usage: %s -config=<config-file.yaml>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	// Load the service configuration.
	config, err := LoadServiceConfig(configFilePath)
	if err != nil {
		log.Fatalf("Error loading service config: %v", err)
	}

	// Generate Docker build workflows.
	workflows, err := config.GetDockerBuildWorkflows()
	if err != nil {
		log.Fatalf("Error generating workflows: %v", err)
	}

	// Define the output directory.
	outputDir := filepath.Join(".github", "workflows")
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory '%s': %v", outputDir, err)
	}

	// Iterate over the generated workflows and write them to files.
	for i, workflow := range workflows {
		// Ensure there is a corresponding DockerConfig.
		if i >= len(config.Build.DockerConfigs) {
			log.Printf("Warning: More workflows generated than DockerConfigs defined.")
			break
		}

		buildConfig := config.Build.DockerConfigs[i]
		buildName := buildConfig.Name

		// Create a unique filename using the build name.
		outputFileName := fmt.Sprintf("gen_build_%s_%s.yml", config.Name, buildName)
		outputFilePath := filepath.Join(outputDir, outputFileName)

		// Write the workflow content to the file.
		err := os.WriteFile(outputFilePath, []byte(workflow), 0644)
		if err != nil {
			log.Fatalf("Error writing workflow file '%s': %v", outputFilePath, err)
		}

		log.Printf("Workflow file generated successfully at '%s'.", outputFilePath)
	}
}
