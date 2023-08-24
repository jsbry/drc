//go:build release

package main

import _ "embed"

//go:embed config/release.json
var configBytes []byte
