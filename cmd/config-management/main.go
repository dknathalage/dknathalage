package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	Name   string `yaml:"name"`
	Gcloud struct {
		Project          string `yaml:"project"`
		Region           string `yaml:"region"`
		IdentityProvider string `yaml:"identity_provider"`
	} `yaml:"gcloud"`
	Images []struct {
		Name         string `yaml:"name"`
		RegistryUrl  string `yaml:"registry_url"`
		RegistryName string `yaml:"registry_name"`
		Context      string `yaml:"context"`
		Dockerfile   string `yaml:"dockerfile"`
	} `yaml:"images"`
	Source struct {
		Paths []string `yaml:"paths"`
	} `yaml:"source"`
}

func main() {
	fileName := os.Args[1]

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	config := ServiceConfig{}
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		panic(err)
	}

	templ, err := template.New("BuildWorkflow").Parse(BuildWorkflow)
	if err != nil {
		panic(err)
	}

	outputDir := filepath.Join(".github", "workflows")
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directories '%s': %v", outputDir, err)
	}

	file, err := os.Create(outputDir + "/gen_build_" + config.Name + ".yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = templ.Execute(file, config)
	if err != nil {
		panic(err)
	}
}

const BuildWorkflow = `# Auto generated. Do not edit.

name: Build {{ .Name }} 🏗️

on:
	push:
		paths:
			{{- range .Source.Paths }}
			- {{ . }}
			{{- end }}

jobs:
	build:
		runs-on: ubuntu-latest
		steps:
		- name: Checkout
			uses: actions/checkout@v4
`
