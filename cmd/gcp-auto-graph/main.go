package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/infra-tools/gcp-auto-graph/pkg/google/compute"
)

func main() {

	// TODO(mccallry,juliosueiras) Use Viper for config
	templateVMs, err := compute.GetVMs("random-project", "northamerica-northeast1-a")
	// TODO Testing TodoBot
	if err != nil {
		fmt.Println(err)
	}

	tmpl, err := template.New("result").Parse(`
@startuml
top to bottom direction
skinparam linetype ortho
{{range .VMs.Items}}
node "{{.Name}}" [[https://console.cloud.google.com/compute/instancesDetail/zones/northamerica-northeast1-a/instances/{{.Name}}?project=random-project]]
{{end}}

@enduml
  `)

	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(os.Stdout, templateVMs)

	if err != nil {
		fmt.Println(err)
	}
}
