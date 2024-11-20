package service_config

import (
	"bytes"
	"fmt"
	"text/template"
)

// DockerConfig represents the Docker-specific build configuration.
type DockerConfig struct {
	Name       string `yaml:"name"`
	Image      string `yaml:"image"`
	Dockerfile string `yaml:"dockerfile"`
	Context    string `yaml:"context"`

	GoogleServiceAccount   string `yaml:"google_service_account"`
	GoogleArtifactRegistry string `yaml:"google_artifact_registry"`
	GoogleWorkloadIDP      string `yaml:"google_workload_idp"`
}

// ServiceConfig represents the overall service configuration.
type ServiceConfig struct {
	Name  string `yaml:"name"`
	Build struct {
		DockerConfigs []DockerConfig `yaml:"docker_configs"`
	} `yaml:"build"`
}

// WorkflowData holds the data passed to the template.
type WorkflowData struct {
	ServiceName  string
	DockerConfig DockerConfig
}

// GetDockerBuildWorkflows generates GitHub Actions workflow YAML strings for each Docker configuration.
func (c *ServiceConfig) GetDockerBuildWorkflows() ([]string, error) {
	// Define the workflow template with custom delimiters to avoid conflicts.
	buildWorkflowTemplate := `
name: [[ .ServiceName ]].[[ .DockerConfig.Name ]] - Docker Build Workflow 🐳
on:
  workflow_dispatch:
    inputs:
      image_tag:
        type: string
        required: true
  workflow_call:
    inputs:
      image_tag:
        type: string
        required: true

permissions:
  contents: read
  id-token: write

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout 📥
        uses: actions/checkout@v4

      - name: Set up Docker Buildx 🏗️
        uses: docker/setup-buildx-action@v2

      - name: Login to Google Artifact Registry 🛅
        uses: google-github-actions/auth@v2
        with:
          token_format: 'access_token'
          workload_identity_provider: [[ .DockerConfig.GoogleWorkloadIDP ]]
          service_account: [[ .DockerConfig.GoogleServiceAccount ]]      
      
      - name: Configure Docker to use Google Artifact Registry
        run: |
          echo ${{ secrets.GOOGLE_CREDENTIALS }} | docker login -u oauth2accesstoken --password-stdin https://[[ .DockerConfig.GoogleArtifactRegistry ]]

      - name: Build Docker Image 🐳  
        shell: bash
        run: |
          docker build -t [[ .DockerConfig.Image ]]:${{ github.event.inputs.image_tag }} \
                       -t [[ .DockerConfig.Image ]]:${{ github.sha }} \
                       -f [[ .DockerConfig.Dockerfile ]] \
                       [[ .DockerConfig.Context ]]
              
      - name: Push Docker Image 🚀
        run: |
          docker push [[ .DockerConfig.Image ]]:${{ github.event.inputs.image_tag }}
          docker push [[ .DockerConfig.Image ]]:${{ github.sha }}
`

	// Parse the template with custom delimiters.
	tmpl, err := template.New("BuildWorkflow").Delims("[[", "]]").Parse(buildWorkflowTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	var workflows []string

	// Iterate over each Docker configuration to generate workflows.
	for _, dockerConfig := range c.Build.DockerConfigs {
		data := WorkflowData{
			ServiceName:  c.Name,
			DockerConfig: dockerConfig,
		}

		var buf bytes.Buffer
		err := tmpl.Execute(&buf, data)
		if err != nil {
			return nil, fmt.Errorf("failed to execute template for build '%s': %w", dockerConfig.Name, err)
		}

		workflows = append(workflows, buf.String())
	}

	return workflows, nil
}
