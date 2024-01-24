package templates

import _ "embed"

//go:embed files/index.ts.tmpl
var indexTemplate []byte

//go:embed files/ng-package.json.tmpl
var ngPackageTemplate []byte

//go:embed files/public-api.ts.tmpl
var publicApiTemplate []byte

func IndexTemplate() []byte {
	return indexTemplate
}

func NgPackageTemplate() []byte {
	return ngPackageTemplate
}

func PublicApiTemplate() []byte {
	return publicApiTemplate
}

//go:embed files/src/feature.component.ts.tmpl
var componentTemplate []byte

func ComponentTemplate() []byte {
	return componentTemplate
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

//go:embed files/models/feature.model.ts.tmpl
var modelTemplate []byte

//go:embed files/models/feature-response.model.ts.tmpl
var responseModelTemplate []byte

func ModelTemplate() []byte {
	return modelTemplate
}

func ResponseModelTemplate() []byte {
	return responseModelTemplate
}

//go:embed files/services/feature.service.ts.tmpl
var serviceTemplate []byte

//go:embed files/services/feature.adapter.ts.tmpl
var adapterTemplate []byte

func ServiceTemplate() []byte {
	return serviceTemplate
}

func AdapterTemplate() []byte {
	return adapterTemplate
}
