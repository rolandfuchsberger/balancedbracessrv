package web

import (
	"testing"
)

func TestInitPlush(t *testing.T) {

	tr, err := newTemplateRendererPlush()
	if err != nil || tr.template == "" {
		t.Errorf("Template was not loaded: %v", err)
	}
}

func TestInitErrorPlush(t *testing.T) {

	wrongTempName := "wrongTemplateName.wrong"

	_, err := newTemplateRendererPlushName(wrongTempName)

	if err == nil {
		t.Errorf("No error received while loading %q", wrongTempName)
	}

}

func TestHandleErrorInTemplatePlush(t *testing.T) {
	tr, _ := newTemplateRendererPlush()

	//Intercept template
	tr.template = "<%= wrong.var.name %>"

	testHandleErrorInTemplate(tr, t)
}

func TestHandlePlush(t *testing.T) {

	tr, _ := newTemplateRendererPlush()

	testHandle(tr, t)
}
