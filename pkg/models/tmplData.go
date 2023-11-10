package models

//tmplData holds any kind of Data for the template
type TmplData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	DataMap map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}
