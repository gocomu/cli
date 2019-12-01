package templates

// const GocomuYaml = `---
// name: {{ .ProjectName }}
// type: {{ .ProjectType }}
// serve: {{ .RTout }}
// version: 0.0.0
// `

type GocomuYaml struct {
	F int `yaml:"a,omitempty"`
	B int
}
