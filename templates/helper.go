package templates

import (
	"fmt"
	"os"
	"path/filepath"
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
func CreateFile(dir, name, tpl string, data *Data) {
	file, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	t := template.Must(template.New(name).Parse(tpl))
	t.Execute(file, data)
	file.Sync()
}
