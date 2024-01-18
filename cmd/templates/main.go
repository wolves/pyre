package templates

import _ "embed"

//go:embed files/index.ts.tmpl
var indexTemplate []byte

//go:embed files/ng-package.json.tmpl
var ngPackageTemplate []byte

//go:embed files/public-api.ts.tmpl
var publicApiTemplate []byte

//go:embed files/src/feature.component.ts.tmpl
var featureComponentTemplate []byte

func IndexTemplate() []byte {
	return indexTemplate
}

func NgPackageTemplate() []byte {
	return ngPackageTemplate
}

func PublicApiTemplate() []byte {
	return publicApiTemplate
}

func FeatureComponentTemplate() []byte {
	return featureComponentTemplate
}

//go:embed files/state/feature-state.module.ts.tmpl
var stateModuleTemplate []byte

//go:embed files/state/state.ts.tmpl
var stateTemplate []byte

//go:embed files/state/feature.actions.ts.tmpl
var actionsTemplate []byte

//go:embed files/state/feature.effects.ts.tmpl
var effectsTemplate []byte

//go:embed files/state/feature.facade.ts.tmpl
var facadeTemplate []byte

//go:embed files/state/feature.reducer.ts.tmpl
var reducerTemplate []byte

//go:embed files/state/feature.selector.ts.tmpl
var selectorTemplate []byte

func CreateStateTemplate(key string) []byte {
	switch key {
	case "stateModule":
		return stateModuleTemplate
	case "actions":
		return actionsTemplate
	case "effects":
		return effectsTemplate
	case "facade":
		return facadeTemplate
	case "reducer":
		return reducerTemplate
	case "selector":
		return selectorTemplate
	case "state":
		return stateTemplate
	}

	return []byte{}
}

//go:embed files/translations/translations.ts.tmpl
var translationsTemplate []byte

//go:embed files/translations/feature.i18n.ts.tmpl
var i18nTemplate []byte

//go:embed files/translations/index.ts.tmpl
var i18nIndexTemplate []byte

func TranslationsTemplate() []byte {
	return translationsTemplate
}

func I18nTemplate() []byte {
	return i18nTemplate
}

func I18nIndexTemplate() []byte {
	return i18nIndexTemplate
}
