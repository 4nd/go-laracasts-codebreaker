package views

import (
	"encoding/json"
	"fmt"
	"os"
)

const manifestFile = "assets/dist/.vite/manifest.json"

type Manifest struct {
	File    string   `json:"file"`
	Name    string   `json:"name"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	Css     []string `json:"css"`
}

func ParseManifest() (styles []string, scripts []string) {
	if _, err := os.Stat(manifestFile); os.IsNotExist(err) {
		fmt.Printf("manifest file does not exist")
		return styles, scripts
	}
	data, err := os.ReadFile(manifestFile)
	if err != nil {
		fmt.Printf("Error reading manifest file: %v", err)
		return
	}

	var manifests map[string]Manifest
	err = json.Unmarshal(data, &manifests)
	if err != nil {
		fmt.Printf("Error json decoding manifest: %v", err)
		return
	}

	for _, manifest := range manifests {
		scripts = append(scripts, manifest.File)
		if len(manifest.Css) > 0 {
			for _, css := range manifest.Css {
				styles = append(styles, css)
			}
		}
	}

	return styles, scripts
}
