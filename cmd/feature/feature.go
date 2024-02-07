package feature

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/wolves/pyre/cmd/templates"
)

type Component struct {
	SunbirdDir string
	Filename   string
	Name       string

	srcPath      string
	statePath    string
	i18nPath     string
	modelsPath   string
	servicesPath string
}

func (c *Component) Create(noTests bool) error {
	// TODO: Add actual way to either select or provide a real project dir
	projPath := filepath.Join(c.SunbirdDir, "exp")
	if _, err := os.Stat(projPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Project directory %s does not exist.", projPath)
		os.Exit(1)
	}

	featPath, err := createFeatureDir(projPath, c.Filename)
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

	srcPath, err := createSrcDir(featPath)
	cobra.CheckErr(err)
	c.srcPath = srcPath

	statePath, err := c.createStateDir()
	cobra.CheckErr(err)
	c.statePath = statePath

	err = c.createStateFiles()
	cobra.CheckErr(err)

	langs := []string{"de", "en", "fr", "ja", "ru", "tr", "zh"}

	i18nPath, err := c.createI18nDirs(langs)
	cobra.CheckErr(err)
	c.i18nPath = i18nPath

	err = c.createI18nFiles(langs)
	cobra.CheckErr(err)

	modelsPath, err := c.createModelsDir()
	cobra.CheckErr(err)
	c.modelsPath = modelsPath

	err = c.createModelFiles()
	cobra.CheckErr(err)

	servicesPath, err := c.createServicesDir()
	cobra.CheckErr(err)
	c.servicesPath = servicesPath

	err = c.createServiceFiles()
	cobra.CheckErr(err)

	err = c.createComponentFiles()
	cobra.CheckErr(err)

	if !noTests {
		err = c.createTestFiles()
		cobra.CheckErr(err)
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

func (c Component) createStateDir() (string, error) {
	statePath := filepath.Join(c.srcPath, "+state")
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		err = os.Mkdir(statePath, 0o751)
		if err != nil {
			log.Printf("Error creating +state directory: %v\n", err)
			return "", err
		}
	} else {
		log.Printf("+state directory with path '%s' already exists", statePath)
	}

	return statePath, nil
}

func (c Component) createStateFiles() error {
	stateMap := map[string]string{
		"stateModule": c.Filename + "-state.module.ts",
		"actions":     c.Filename + ".actions.ts",
		"effects":     c.Filename + ".effects.ts",
		"facade":      c.Filename + ".facade.ts",
		"reducer":     c.Filename + ".reducer.ts",
		"selector":    c.Filename + ".selector.ts",
		"state":       "state.ts",
	}

	for key, file := range stateMap {
		createdFile, err := os.Create(filepath.Join(c.statePath, file))
		if err != nil {
			return fmt.Errorf("Error: State file creation error: %v", err)
		}
		defer createdFile.Close()

		createdTmpl := template.Must(template.New(key).Parse(string(templates.CreateStateTemplate(key))))
		err = createdTmpl.Execute(createdFile, c)
		if err != nil {
			return fmt.Errorf("Error: State template execution error: %v", err)
		}
	}

	return nil
}

func (c Component) createI18nDirs(langs []string) (string, error) {
	i18nPath := filepath.Join(c.srcPath, "assets/translations")
	if _, err := os.Stat(i18nPath); os.IsNotExist(err) {
		err = os.MkdirAll(i18nPath, 0o751)
		if err != nil {
			log.Printf("Error creating i18n directory: %v\n", err)
			return "", err
		}
	} else {
		log.Printf("i18n directory with path '%s' already exists", i18nPath)
	}

	for _, lang := range langs {
		langPath := filepath.Join(i18nPath, lang)
		if _, err := os.Stat(langPath); os.IsNotExist(err) {
			err = os.Mkdir(langPath, 0o751)
			if err != nil {
				log.Printf("Error creating i18n %s directory: %v\n", lang, err)
				return "", err
			}
		} else {
			log.Printf("i18n language directory with path '%s' already exists", langPath)
		}
	}

	return i18nPath, nil
}

func (c Component) createI18nFiles(langs []string) error {
	data := struct {
		I18nKey  string
		LangList []string
		Lang     string
		Filename string
		Name     string
	}{
		I18nKey:  strings.ToUpper(strings.ReplaceAll(c.Filename, "-", "_")),
		LangList: langs,
		Filename: c.Filename,
		Name:     c.Name,
	}

	for _, lang := range langs {
		createdFile, err := os.Create(filepath.Join(c.i18nPath, lang, c.Filename+".i18n.ts"))
		if err != nil {
			return fmt.Errorf("Error: i18n file creation error: %v", err)
		}
		defer createdFile.Close()

		createdIdx, err := os.Create(filepath.Join(c.i18nPath, lang, "index.ts"))
		if err != nil {
			return fmt.Errorf("Error: i18n file creation error: %v", err)
		}
		defer createdIdx.Close()

		data.Lang = lang

		createdTmpl := template.Must(template.New(lang).Parse(string(templates.I18nTemplate())))
		err = createdTmpl.Execute(createdFile, data)
		if err != nil {
			return fmt.Errorf("Error: i18n template execution error: %v", err)
		}

		createdIdxTmpl := template.Must(template.New(lang + "Idx").Parse(string(templates.I18nIndexTemplate())))
		err = createdIdxTmpl.Execute(createdIdx, data)
		if err != nil {
			return fmt.Errorf("Error: i18n index template execution error: %v", err)
		}
	}

	createdTranslationsFile, err := os.Create(filepath.Join(c.i18nPath, "translations.ts"))
	if err != nil {
		return fmt.Errorf("Error: i18n translations file creation error: %v", err)
	}
	defer createdTranslationsFile.Close()

	createdTranslationsTmpl := template.Must(template.New("translations").Parse(string(templates.TranslationsTemplate())))
	err = createdTranslationsTmpl.Execute(createdTranslationsFile, data)
	if err != nil {
		return fmt.Errorf("Error: i18n translations template execution error: %v", err)
	}

	return nil
}

func (c Component) createModelsDir() (string, error) {
	modelsPath := filepath.Join(c.srcPath, "models")
	if _, err := os.Stat(modelsPath); os.IsNotExist(err) {
		err = os.Mkdir(modelsPath, 0o751)
		if err != nil {
			log.Printf("Error creating models directory: %v\n", err)
			return "", err
		}
	} else {
		log.Printf("models directory with path '%s' already exists", modelsPath)
	}

	return modelsPath, nil
}

func (c Component) createModelFiles() error {
	modelFile, err := os.Create(filepath.Join(c.modelsPath, c.Filename+".model.ts"))
	if err != nil {
		return fmt.Errorf("Error: model file creation error: %v", err)
	}
	defer modelFile.Close()

	responseModelFile, err := os.Create(filepath.Join(c.modelsPath, c.Filename+"-response.model.ts"))
	if err != nil {
		return fmt.Errorf("Error: response model file creation error: %v", err)
	}
	defer responseModelFile.Close()

	modelTmpl := template.Must(template.New("model").Parse(string(templates.ModelTemplate())))
	err = modelTmpl.Execute(modelFile, c)
	if err != nil {
		return fmt.Errorf("Error: model template execution error: %v", err)
	}

	responseModelTmpl := template.Must(template.New("response-model").Parse(string(templates.ResponseModelTemplate())))
	err = responseModelTmpl.Execute(responseModelFile, c)
	if err != nil {
		return fmt.Errorf("Error: response model template execution error: %v", err)
	}

	return nil
}

func (c Component) createServicesDir() (string, error) {
	servicesPath := filepath.Join(c.srcPath, "services")
	if _, err := os.Stat(servicesPath); os.IsNotExist(err) {
		err = os.Mkdir(servicesPath, 0o751)
		if err != nil {
			log.Printf("Error creating services directory: %v\n", err)
			return "", err
		}
	} else {
		log.Printf("services directory with path '%s' already exists", servicesPath)
	}

	return servicesPath, nil
}

func (c Component) createServiceFiles() error {
	serviceFile, err := os.Create(filepath.Join(c.servicesPath, c.Filename+".service.ts"))
	if err != nil {
		return fmt.Errorf("Error: service file creation error: %v", err)
	}
	defer serviceFile.Close()

	adapterFile, err := os.Create(filepath.Join(c.servicesPath, c.Filename+".adapter.ts"))
	if err != nil {
		return fmt.Errorf("Error: service adapter file creation error: %v", err)
	}
	defer adapterFile.Close()

	serviceTmpl := template.Must(template.New("service").Parse(string(templates.ServiceTemplate())))
	err = serviceTmpl.Execute(serviceFile, c)
	if err != nil {
		return fmt.Errorf("Error: service template execution error: %v", err)
	}

	adapterTmpl := template.Must(template.New("adapter").Parse(string(templates.AdapterTemplate())))
	err = adapterTmpl.Execute(adapterFile, c)
	if err != nil {
		return fmt.Errorf("Error: service adapter template execution error: %v", err)
	}

	return nil
}

func (c Component) createComponentFiles() error {
	moduleFile, err := os.Create(filepath.Join(c.srcPath, c.Filename+".module.ts"))
	if err != nil {
		return fmt.Errorf("Error: module file creation error: %v", err)
	}
	defer moduleFile.Close()

	componentFile, err := os.Create(filepath.Join(c.srcPath, c.Filename+".component.ts"))
	if err != nil {
		return fmt.Errorf("Error: component file creation error: %v", err)
	}
	defer componentFile.Close()

	htmlFile, err := os.Create(filepath.Join(c.srcPath, c.Filename+".component.html"))
	if err != nil {
		return fmt.Errorf("Error: component html template creation error: %v", err)
	}
	defer htmlFile.Close()

	scssFile, err := os.Create(filepath.Join(c.srcPath, c.Filename+".component.scss"))
	if err != nil {
		return fmt.Errorf("Error: component scss template creation error: %v", err)
	}
	defer scssFile.Close()

	routingFile, err := os.Create(filepath.Join(c.srcPath, c.Filename+"-routing.module.ts"))
	if err != nil {
		return fmt.Errorf("Error: routing file creation error: %v", err)
	}
	defer routingFile.Close()

	moduleTemplate := template.Must(template.New("module").Parse(string(templates.ModuleTemplate())))
	err = moduleTemplate.Execute(moduleFile, c)
	if err != nil {
		return fmt.Errorf("Error: module template execution error: %v", err)
	}

	componentTemplate := template.Must(template.New("component").Parse(string(templates.ComponentTemplate())))
	err = componentTemplate.Execute(componentFile, c)
	if err != nil {
		return fmt.Errorf("Error: component template execution error: %v", err)
	}

	htmlTemplate := template.Must(template.New("html").Parse(string(templates.HtmlTemplate())))
	err = htmlTemplate.Execute(htmlFile, c)
	if err != nil {
		return fmt.Errorf("Error: html template execution error: %v", err)
	}

	scssTemplate := template.Must(template.New("scss").Parse(string(templates.ScssTemplate())))
	err = scssTemplate.Execute(scssFile, c)
	if err != nil {
		return fmt.Errorf("Error: scss template execution error: %v", err)
	}

	routingTemplate := template.Must(template.New("routing").Parse(string(templates.RoutingModuleTemplate())))
	err = routingTemplate.Execute(routingFile, c)
	if err != nil {
		return fmt.Errorf("Error: routing template execution error: %v", err)
	}

	return nil
}

func (c Component) createTestFiles() error {
	fmt.Println("CREATING SPECS")
	stateMap := map[string]string{
		"actions":  c.Filename + ".actions.spec.ts",
		"effects":  c.Filename + ".effects.spec.ts",
		"facade":   c.Filename + ".facade.spec.ts",
		"reducer":  c.Filename + ".reducer.spec.ts",
		"selector": c.Filename + ".selector.spec.ts",
	}

	for key, file := range stateMap {
		createdFile, err := os.Create(filepath.Join(c.statePath, file))
		if err != nil {
			return fmt.Errorf("Error: State file creation error: %v", err)
		}
		defer createdFile.Close()

		createdTmpl := template.Must(template.New(key).Parse(string(templates.CreateStateTestTemplate(key))))
		err = createdTmpl.Execute(createdFile, c)
		if err != nil {
			return fmt.Errorf("Error: State template execution error: %v", err)
		}
	}

	return nil
}
