package templates

const EmbedGo = `// +build embed

package main

//go:generate go run -tags embed ./embed
func main() {

}
`

const EmbeddedGo = `package embed

import "github.com/gocomu/comu/embed"

//func FS() {

//}
`
