package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Eulenwald/GoWeb/pkg/config"
	"github.com/Eulenwald/GoWeb/pkg/models"
)

func AddDefaultData(tmplData *models.TmplData) *models.TmplData {	
	return tmplData 
}

var pApp *config.AppConfig

func NewTemplate(p *config.AppConfig) {
	pApp = p
}

func RenderATemplate(w http.ResponseWriter, tmpl string,  pTmplData *models.TmplData) {

	var err error
	var mpCache map[string]*template.Template

	if pApp.UseCache {
		// get the cache from AppConfig
		mpCache = pApp.TemplateCache
	} else {
		mpCache, _ = CreateTemplateCache()
 }
	// get the caching template
	actTempl, ok := mpCache[tmpl]
	if !ok {
		log.Fatalf("Error %s not cached.\n", tmpl)
	}

	// store the result in a buffer to check its valid
	buf := new(bytes.Buffer)
	
	pTmplData = AddDefaultData(pTmplData)
	
	err = actTempl.Execute(buf, pTmplData)
	if err != nil {
		log.Panicln(err)
	}

	// Now we render the actTemplate
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Panicln(err)
	}	
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//var err error
	//var pages []string
	mpCacheVal := map[string]*template.Template{}

	// slTemplate => slice of strings
	slTempl, err := filepath.Glob("./templates/*-page.tmpl")
	if err != nil {
		return mpCacheVal, err
	}

	// iterate over a sliceOfString
	for _, sTempl := range slTempl {
		// get the filename of the path
		name := filepath.Base(sTempl)
		// create a new template
		pTempl, err := template.New(name).ParseFiles(sTempl)
		if err != nil {
			return mpCacheVal, err
		}

		slLayoutTempl, err := filepath.Glob("./templates/*-layout.tmpl")
		if err != nil {
			return mpCacheVal, err
		}

		if len(slLayoutTempl) > 0 {
			// layout founded. actuall page and founded layouts concat
			pTempl, err = pTempl.ParseGlob("./templates/*-layout.tmpl")
			if err != nil {
				return mpCacheVal, err
			}
		}
		mpCacheVal[name] = pTempl
	}

	return mpCacheVal, nil
}
