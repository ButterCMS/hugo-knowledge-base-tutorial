package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func mainn() {
	// Read the JSON file
	data, err := ioutil.ReadFile("./data/kbdata.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the JSON data into a map[string]interface{}
	var d map[string]interface{}
	err = json.Unmarshal(data, &d)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the data array
	dataArray, ok := d["data"].([]interface{})
	if !ok {
		fmt.Println("Invalid data format")
		return
	}

	// Define a template for the Markdown files
	const templateString = `# {{ .Title }}

{{ range $key, $value := .Data }}
{{ $key }}: {{ $value }}
{{ end }}
`

	// Create the "content" folder if it doesn't exist
	err = os.MkdirAll("content", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Loop through the data objects and create a Markdown file for each one
	for i, dataObj := range dataArray {
		// Convert the data object to a map[string]interface{}
		dataMap, ok := dataObj.(map[string]interface{})
		if !ok {
			fmt.Printf("Invalid data object format: %v\n", dataObj)
			return
		}

		// Create the Markdown file
		fileName := fmt.Sprintf("data%d.md", i+1)
		filePath := filepath.Join("content", fileName)
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Render the Markdown template for the data object
		title := fmt.Sprintf("Data %d", i+1)
		tmpl, err := template.New("markdown").Parse(templateString)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = tmpl.Execute(file, map[string]interface{}{
			"Title": title,
			"Data":  dataMap,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
