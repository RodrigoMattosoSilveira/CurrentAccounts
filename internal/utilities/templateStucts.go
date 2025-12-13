package utilities

type TemplateStruct struct {
	Name     string
	Fn       string
	FullName string
}

type TemplateStructs struct {
	Data []TemplateStruct
}

func NewTemplateStructures(name string, fn string) *TemplateStructs {
	templateStruct := TemplateStruct{"layout", "layout.tmpl", ""}

	templateStructs := TemplateStructs{
		Data: []TemplateStruct{templateStruct},
	}
	templateStructs.Data = append(templateStructs.Data, TemplateStruct{name, fn, ""})
	return &templateStructs
}

func (tdata *TemplateStructs) AddData(name string, fn string) {
	tdata.Data = append(tdata.Data, TemplateStruct{name, fn, ""})
}

func (tdata *TemplateStructs) GetAllTemplateStructures() []TemplateStruct {
	return tdata.Data
}
