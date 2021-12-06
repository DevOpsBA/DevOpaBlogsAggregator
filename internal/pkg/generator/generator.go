package generator

import (
	"html/template"
	"log"
	"os"
)

// Основная функция которая должна производить генерацию
func Generate(tmpl string, outfilePath string, fields interface{}) error {
	t, err := template.New("Article").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(outfilePath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	errEx := t.Execute(outFile, fields)
	if errEx != nil {
		log.Println("executing template:", errEx)
	}

	return nil
}
