# Ad Auto Translate

## Overview

Ad Auto Translate is a tool built to automate the addition of new translations. This document provides a comprehensive guide on how to install, configure, and use Ad Auto Translate.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Live Development](#live-development)
- [Building](#building)

## Installation

### Prerequisites
- [Go](https://golang.org/doc/install) version 1.2+
- [Wails](https://wails.io/) version 2.8+

Step 2: Install dependencies -> go mod download </br>
Step 3: Install FE dependencies -> cd /frontend - npm install
## Configuration

### General Config
To configure your project, simply open the `wails.json` file in a text editor and modify the values as needed. After saving your changes, they will be applied the next time you run or build your application.
For a more detailed explanation of the fields in `wails.json` and how to use them, please refer to the [Wails project configuration documentation](https://wails.io/docs/reference/project-config).

### Project Config
Within the `functions/configuration.go` file, there is a variable named `initialConfig`. This variable holds all the necessary information that is generated during the initial load of the project.

## Live Development
`wails dev`.

## Building
`wails build`.
