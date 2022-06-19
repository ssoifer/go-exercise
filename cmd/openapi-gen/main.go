package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-exercise/internal/rest"
	"log"
	"os"
	"path"

	"github.com/ghodss/yaml"
)

func main() {
	var output string

	flag.StringVar(&output, "path", "./internal/rest/generated/swagger", "Path to use for generating OpenAPI 3 files")
	flag.Parse()

	if output == "" {
		log.Fatalln("path is required")
	}

	swagger := rest.NewOpenAPI3()

	// openapi3.json
	data, err := json.Marshal(&swagger)
	if err != nil {
		log.Fatalf("Couldn't marshal json: %s", err)
	}

	if err := os.WriteFile(path.Join(output, "openapi3.json"), data, 0644); err != nil {
		log.Fatalf("Couldn't write json: %s", err)
	}

	// openapi3.yaml
	data, err = yaml.Marshal(&swagger)
	if err != nil {
		log.Fatalf("Couldn't marshal json: %s", err)
	}

	if err := os.WriteFile(path.Join(output, "openapi3.yaml"), data, 0644); err != nil {
		log.Fatalf("Couldn't write json: %s", err)
	}

	fmt.Println("all generated")
}
