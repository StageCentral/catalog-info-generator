package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Name        string
	Description string
}

type Spec struct {
	Type      string
	Lifecycle string
	Owner     string
}

type Entity struct {
	apiVersion string
	Kind       string
	Metadata   Metadata
	Spec       Spec
}

var qs = []*survey.Question{
	{
		Name: "Kind",
		Prompt: &survey.Select{
			Message: "Enter your entity kind:",
			Options: []string{"Component", "API", "Resource", "System"},
			Default: "Component",
		},
	},
	{
		Name:   "Name",
		Prompt: &survey.Input{Message: "What is your entity name?"},
	},
	{
		Name:   "Description",
		Prompt: &survey.Input{Message: "What is your entity description?"},
	},
	{
		Name: "Type",
		Prompt: &survey.Select{
			Message: "Enter your entity type:",
			Options: []string{"Service", "Database", "Website"},
			Default: "Service",
		},
	},
	{
		Name:   "Owner",
		Prompt: &survey.Input{Message: "Who is the entity owner?"},
	},
	{
		Name: "Lifecycle",
		Prompt: &survey.Select{
			Message: "Enter your entity lifecycle:",
			Options: []string{"production", "experimental", "deprecated"},
			Default: "production",
		},
	},
}

func main() {

	answers := struct {
		Kind        string
		Description string
		Name        string
		Type        string
		Lifecycle   string
		Owner       string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cataloginfo := Entity{}

	cataloginfo.apiVersion = "backstage.io/v1alpha1"
	cataloginfo.Kind = answers.Kind
	cataloginfo.Metadata.Name = answers.Name
	cataloginfo.Metadata.Description = answers.Description
	cataloginfo.Spec.Owner = answers.Owner
	cataloginfo.Spec.Type = answers.Type

	yamlData, err := yaml.Marshal(&cataloginfo)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fmt.Println(" --- YAML ---")
	fmt.Println(string(yamlData)) // yamlData will be in bytes. So converting it to string.
}
