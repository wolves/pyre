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

	featPath, err := createFeatureDir(projPath, c.Filename)
	cobra.CheckErr(err)

	srcPath, err := createSrcDir(featPath)
	cobra.CheckErr(err)

	indexFile, err := os.Create(filepath.Join(featPath, "index.ts"))
	cobra.CheckErr(err)
	defer indexFile.Close()

	indexTemplate := template.Must(template.New("index").Parse(string(templates.IndexTemplate())))
	err = indexTemplate.Execute(indexFile, c)
	if err != nil {
		return err
	}

	ngPackageFile, err := os.Create(filepath.Join(featPath, "ng-package.json"))
	cobra.CheckErr(err)
	defer ngPackageFile.Close()

	ngPackageTemplate := template.Must(template.New("ngPackage").Parse(string(templates.NgPackageTemplate())))
	err = ngPackageTemplate.Execute(ngPackageFile, c)
	if err != nil {
		return err
	}

	publicApiFile, err := os.Create(filepath.Join(featPath, "public-api.json"))
	cobra.CheckErr(err)
	defer publicApiFile.Close()

	publicApiTemplate := template.Must(template.New("publicApi").Parse(string(templates.PublicApiTemplate())))
	err = publicApiTemplate.Execute(publicApiFile, c)
	if err != nil {
		return err
	}

	componentFile, err := os.Create(filepath.Join(srcPath, c.Filename+".component.ts"))
	cobra.CheckErr(err)
	defer componentFile.Close()

	componentTemplate := template.Must(template.New("component").Parse(string(templates.FeatureComponentTemplate())))
	err = componentTemplate.Execute(componentFile, c)
	if err != nil {
		return err
	}

	return nil
}

func createFeatureDir(projPath, feat string) (string, error) {
	featPath := filepath.Join(projPath, feat)
	if _, err := os.Stat(featPath); os.IsNotExist(err) {
		err = os.Mkdir(featPath, 0o751)
		if err != nil {
			log.Printf("Error creating feature directory: %v\n", err)
			return "", err
		}
	} else {
		log.Printf("Feature directory with name '%s' already exists", feat)
	}

	return featPath, nil
}

func createSrcDir(featPath string) (string, error) {
	srcPath := filepath.Join(featPath, "src")
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		err = os.Mkdir(srcPath, 0o751)
		if err != nil {
			log.Printf("Error creating src directory: %v\n", err)
			return "", err
		}
	} else {
		log.Printf("Src directory with path '%s' already exists", srcPath)
	}

	return srcPath, nil
}
