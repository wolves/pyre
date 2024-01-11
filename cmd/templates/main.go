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
