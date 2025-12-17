package utilities

type T_Files struct {
	Files []string
}

func NewTemplateFiles(file string) *T_Files {
	return &T_Files{Files: []string{file}}
}

func NewTemplateFilesLayout(file string) *T_Files {
	layoutFile := "root/layout.tmpl"
	tFiles := &T_Files{Files: []string{layoutFile}}
	tFiles.Files = append(tFiles.Files, file)
	return tFiles
}
func (tdata *T_Files) AddTemplate(file string) {
	tdata.Files = append(tdata.Files, file)
}

func (tdata *T_Files) GetAllTemplates() []string {
	return tdata.Files
}
