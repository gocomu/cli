package templates

import (
	"os"
	"text/template"
)

// Data holds values to be substituted
// when creating/executing biolerplate templates
type Data struct {
	ProjectName string
	//ProjectType int
	//Output      int
}

func CreateFile(path, name, tpl string, data *Data) {
	file, _ := os.Create(path + name)
	defer file.Close()
	t := template.Must(template.New(name).Parse(tpl))
	t.Execute(file, data)
	file.Sync()
}
