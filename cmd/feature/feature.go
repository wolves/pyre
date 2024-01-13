package feature

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wolves/pyre/cmd/templates"
)

type Component struct {
	SunbirdDir string
	Filename   string
	Name       string
}

func (c *Component) Create() error {
	fmt.Printf("Component %+v", c)

	// TODO: Add actual way to either select or provide a real project dir
	projPath := filepath.Join(c.SunbirdDir, "exp")
	if _, err := os.Stat(projPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Project directory %s does not exist.", projPath)
		os.Exit(1)
	}

	err := createFeatureDir(projPath, c.Filename)
	cobra.CheckErr(err)

	// err := createSrcDir()

	publicApiFile, err := os.Create(filepath.Join(projPath, "public-api.json"))
	cobra.CheckErr(err)
	defer publicApiFile.Close()

	publicApiTemplate := template.Must(template.New("publicApi").Parse(string(templates.PublicApiTemplate())))
	err = publicApiTemplate.Execute(publicApiFile, c)
	if err != nil {
		return err
	}

	return nil
}

func createFeatureDir(projPath, feat string) error {
	featPath := filepath.Join(projPath, feat)
	if _, err := os.Stat(featPath); os.IsNotExist(err) {
		err = os.Mkdir(featPath, 0751)
		if err != nil {
			log.Printf("Error creating feature directory: %v\n", err)
			return err
		}
	} else {
		log.Printf("Feature directory with name '%s' already exists", feat)
	}

	return nil
}
