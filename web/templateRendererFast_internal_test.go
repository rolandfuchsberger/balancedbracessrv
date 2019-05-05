package web

import (
	"html/template"
	"testing"
)

func TestInitErrorFast(t *testing.T) {

	wrongTempName := "wrongTemplateName.wrong"

	_, err := newTemplateRendererFastName(wrongTempName)

	if err == nil {
		t.Errorf("No error received while loading %q", wrongTempName)
	}

}

func TestHandleErrorInTemplateFast(t *testing.T) {
	tr, _ := newTemplateRendererFast()

	//Intercept template
	templateTestString := "{{.wrongVarName}}"
	tr.template, _ = template.New("balancedbraces").Parse(templateTestString)

	testHandleErrorInTemplate(tr, t)
}

func TestHandleFast(t *testing.T) {

	tr, _ := newTemplateRendererFast()

	testHandle(tr, t)
}
