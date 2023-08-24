//go:build !release

package main

import _ "embed"

//go:embed config/debug.json
var configBytes []byte
