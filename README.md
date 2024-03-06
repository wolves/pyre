# Pyre

_A CLI helper for common Sunbird development tasks_

## Table of Contents

- [Description](#desc)
- [Install](#install)
- [Usage Examples](#usage)

<a id="desc"></a>

## Description

Pyre is a CLI used to simplify some of the Sunbird development teams common tasks.

In these early stages, it is currently capable of automating the boilerplate file creation necessary for implimenting new features in the different Sunbird Angular projects.

The intention is that as more pain points are indentified Pyre can be expanded to help alleviate those tasks as well.

<a id="install"></a>

## Installation

### Via Prebuilt Binaries

Prebuilt binaries are available for a variety of operating systems and architectures. Visit the [latest release](https://github.com/wolves/pyre/releases/latest) page, and scroll down to the Assets section.

1. Download the archive for the desired edition, operating system, and architecture
2. Extract the archive
3. Move the executable to the desired directory
4. Add this directory to the PATH environment variable
5. Verify that you have execute permission on the file

Please consult your operating system documentation if you need help setting file permissions or modifying your PATH environment variable.

### Via `go install`

```sh
go install github.com/wolves/pyre@latest
```

This installs a go binary that will automatically bind to your $GOPATH

<a id="usage"></a>

## Usage

### `feature create`

This command will generate the necessary Angular feature boilerplate files including specs and state management architecture. The component naming will be camel cased within the generated files (eg. `YourFeatureName`) but should be provided as kebab-case like in the example.

The `--path/-p` flag is required in order to determine the location that the feature directory should be created.

`--no-tests/-x` can be provided as a flag if you wish to generate the feature files but exclude the test specs

_Example:_

```sh
pyre feature create your-feature-name -p ~/code/project-name/features-folder
```

_Results in:_

```sh
~/code/project-name/features-folder/your-feature-name
├── src
│  ├── +state
│  │  ├── state.ts
│  │  ├── your-feature-name-state.module.ts
│  │  ├── your-feature-name.actions.spec.ts
│  │  ├── your-feature-name.actions.ts
│  │  ├── your-feature-name.effects.spec.ts
│  │  ├── your-feature-name.effects.ts
│  │  ├── your-feature-name.facade.spec.ts
│  │  ├── your-feature-name.facade.ts
│  │  ├── your-feature-name.reducer.spec.ts
│  │  ├── your-feature-name.reducer.ts
│  │  └── your-feature-name.selector.ts
│  ├── assets
│  │  └── translations
│  │     ├── de
│  │     │  ├── index.ts
│  │     │  └── your-feature-name.i18n.ts
│  │     ├── en
│  │     │  ├── index.ts
│  │     │  └── your-feature-name.i18n.ts
│  │     ├── fr
│  │     │  ├── index.ts
│  │     │  └── your-feature-name.i18n.ts
│  │     ├── ja
│  │     │  ├── index.ts
│  │     │  └── your-feature-name.i18n.ts
│  │     ├── ru
│  │     │  ├── index.ts
│  │     │  └── your-feature-name.i18n.ts
│  │     ├── tr
│  │     │  ├── index.ts
│  │     │  └── your-feature-name.i18n.ts
│  │     ├── zh
│  │     │  ├── index.ts
│  │     │  └── your-feature-name.i18n.ts
│  │     └── translations.ts
│  ├── models
│  │  ├── your-feature-name-response.model.ts
│  │  └── your-feature-name.model.ts
│  ├── services
│  │  ├── your-feature-name.adapter.ts
│  │  ├── your-feature-name.service.spec.ts
│  │  └── your-feature-name.service.ts
│  ├── your-feature-name-routing.module.ts
│  ├── your-feature-name.component.html
│  ├── your-feature-name.component.scss
│  ├── your-feature-name.component.spec.ts
│  ├── your-feature-name.component.ts
│  └── your-feature-name.module.ts
├── index.ts
├── ng-package.json
└── public-api.json
```
