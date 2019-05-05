package web

import (
	"html/template"
	"testing"
)

func TestInitError(t *testing.T) {

	wrongTempName := "wrongTemplateName.wrong"

	_, err := newTemplateRenderer(wrongTempName)

	if err == nil {
		t.Errorf("No error received while loading %q", wrongTempName)
	}

}

func TestHandleErrorInTemplate(t *testing.T) {
	tr, _ := NewTemplateRenderer()

	//Intercept template
	templateTestString := "{{.wrongVarName}}"
	tr.template, _ = template.New("balancedbraces").Parse(templateTestString)

	testHandleErrorInTemplate(tr, t)
}

func TestHandle(t *testing.T) {

	tr, _ := NewTemplateRenderer()

	testHandle(tr, t)
}
