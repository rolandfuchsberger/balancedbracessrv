package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"fuchsberger.email/balancedbracessrv/balancedbraces"
)

//TemplateRendererFast asdf
type TemplateRendererFast struct {
	templateName string
	template     *template.Template
}

//NewTemplateRendererFast asdf
func NewTemplateRendererFast() (*TemplateRendererFast, error) {
	return newTemplateRendererFast(templateNameBB)
}

//for testing: able to inject wrong template name for error testing
func newTemplateRendererFast(templateName string) (*TemplateRendererFast, error) {

	tr := &TemplateRendererFast{}

	// Get the string representation of a file, or an error if it doesn't exist:
	bbString, err := box.FindString(templateName)

	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("balancedbraces").Parse(bbString)

	if err != nil {
		return nil, fmt.Errorf("Cannot parse %q: %v", templateName, err)
	}

	//load templates
	tr.template = tmpl

	return tr, nil

}

//Handle s web requests for balanced braces
func (t *TemplateRendererFast) Handle(w http.ResponseWriter, r *http.Request) {

	//get expression get parameter
	exprs := r.URL.Query()["expression"]
	expr := ""
	if exprs != nil && len(exprs) > 0 {
		expr = exprs[0]
	}

	//build context for template
	ctx := struct {
		DisplayBalanced, IsBalanced bool
		Input                       string
	}{
		expr != "", balancedbraces.BalancedBraces(expr),
		expr,
	}

	err := t.template.Execute(w, ctx)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Printf("error parsing %q : %v\n", t.templateName, err)
		return
	}
}
