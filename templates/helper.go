package templates

import (
	"fmt"
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

// CreateFile is an helper for creating new files from templates
func CreateFile(path, name, tpl string, data *Data) {
	dir, _ := os.Getwd()
	file, err := os.Create(dir + "/" + path + name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	t := template.Must(template.New(name).Parse(tpl))
	t.Execute(file, data)
	file.Sync()
}
