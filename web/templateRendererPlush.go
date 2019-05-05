package web

import (
	"log"
	"net/http"

	"github.com/gobuffalo/plush"

	"fuchsberger.email/balancedbracessrv/balancedbraces"
)

const (
	templateNameBBPlush = "balancedbraces.plush.html"
)

//TemplateRendererPlush asdf
type templateRendererPlush struct {
	templateName string
	template     string
}

//NewTemplateRendererPlush asdf
func newTemplateRendererPlush() (*templateRendererPlush, error) {
	return newTemplateRendererPlushName(templateNameBBPlush)
}

func newTemplateRendererPlushName(templateName string) (*templateRendererPlush, error) {

	tr := &templateRendererPlush{}

	tr.templateName = templateName

	// Get the string representation of a file, or an error if it doesn't exist:
	tmpl, err := box.FindString(templateName)

	if err != nil {
		return nil, err
	}

	tr.template = tmpl

	return tr, nil

}

//Handle s web requests for balanced braces
func (t *templateRendererPlush) Handle(w http.ResponseWriter, r *http.Request) {

	//get expression get parameter
	exprs := r.URL.Query()["expression"]
	expr := ""
	if exprs != nil && len(exprs) > 0 {
		expr = exprs[0]
	}

	//set context
	ctx := plush.NewContext()
	ctx.Set("DisplayBalanced", expr != "")
	ctx.Set("IsBalanced", balancedbraces.BalancedBraces(expr))
	ctx.Set("Input", expr)

	s, err := plush.Render(t.template, ctx)
	if err != nil {
		http.Error(w, "500 Internal Server Error", 500)
		log.Printf("error parsing %q : %v\n", t.templateName, err)
	}

	w.Write([]byte(s))

}
