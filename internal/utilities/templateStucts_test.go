package utilities

import "testing"

func TestNewTemplateStructures(t *testing.T) {
	name0 := "layout"
	fn0 := "layout.tmpl"
	name1 := "top"
	fn1 := "authentication/welcome.tmpl"
	numberofTemplatesExpected := 2

	templates := NewTemplateStructures(name1, fn1)
	if len(templates.Data) != numberofTemplatesExpected {
		t.Errorf("Expected %d templates, got %d", numberofTemplatesExpected, len(templates.Data))
	}
	if templates.Data[0].Name != name0 {
		t.Errorf("First name incorrect, got: %s, want: %s.", templates.Data[0].Name, name0)
	}
	if templates.Data[0].Fn != fn0 {
		t.Errorf("First file name incorrect, got: %s, want: %s.", templates.Data[0].Fn, fn0)
	}
	if templates.Data[1].Name != name1 {
		t.Errorf("Second name incorrect, got: %s, want: %s.", templates.Data[1].Name, name1)
	}
	if templates.Data[1].Fn != fn1 {
		t.Errorf("Second file name incorrect, got: %s, want: %s.", templates.Data[1].Fn, fn1)
	}

}
func TestAddTemplateStructure(t *testing.T) {
	name0 := "layout"
	fn0 := "layout.tmpl"
	name1 := "top"
	fn1 := "authentication/welcome.tmpl"
	name2 := "bottom"
	fn2 := "people/cc.tmpl"
	numberofTemplatesExpected := 3

	templates := NewTemplateStructures(name1, fn1)
	templates.AddData(name2, fn2)
	if len(templates.Data) != numberofTemplatesExpected {
		t.Errorf("Expected %d templates, got %d", numberofTemplatesExpected, len(templates.Data))
	}
	if templates.Data[0].Name != name0 {
		t.Errorf("First name incorrect, got: %s, want: %s.", templates.Data[0].Name, name0)
	}
	if templates.Data[0].Fn != fn0 {
		t.Errorf("First file name incorrect, got: %s, want: %s.", templates.Data[0].Fn, fn0)
	}
	if templates.Data[1].Name != name1 {
		t.Errorf("Second name incorrect, got: %s, want: %s.", templates.Data[1].Name, name1)
	}
	if templates.Data[1].Fn != fn1 {
		t.Errorf("Second file name incorrect, got: %s, want: %s.", templates.Data[1].Fn, fn1)
	}
	if templates.Data[2].Name != name2 {
		t.Errorf("Third name incorrect, got: %s, want: %s.", templates.Data[2].Name, name2)
	}
	if templates.Data[2].Fn != fn2 {
		t.Errorf("Third file name incorrect, got: %s, want: %s.", templates.Data[2].Fn, fn2)
	}
}
func TestGetAllTemplateStructures(t *testing.T) {
	name0 := "layout"
	fn0 := "layout.tmpl"
	name1 := "top"
	fn1 := "authentication/welcome.tmpl"
	name2 := "bottom"
	fn2 := "people/cc.tmpl"
	numberofTemplatesExpected := 3

	templates := NewTemplateStructures(name1, fn1)
	templates.AddData(name2, fn2)
	templatesSlice := templates.GetAllTemplateStructures()
	if len(templatesSlice) != numberofTemplatesExpected {
		t.Errorf("Expected %d templates, got %d", numberofTemplatesExpected, len(templatesSlice))
	}
	if templatesSlice[0].Name != name0 {
		t.Errorf("First name incorrect, got: %s, want: %s.", templatesSlice[0].Name, name0)
	}
	if templatesSlice[0].Fn != fn0 {
		t.Errorf("First file name incorrect, got: %s, want: %s.", templatesSlice[0].Fn, fn0)
	}
	if templatesSlice[1].Name != name1 {
		t.Errorf("Second name incorrect, got: %s, want: %s.", templatesSlice[1].Name, name1)
	}
	if templatesSlice[1].Fn != fn1 {
		t.Errorf("Second file name incorrect, got: %s, want: %s.", templatesSlice[1].Fn, fn1)
	}
	if templatesSlice[2].Name != name2 {
		t.Errorf("Third name incorrect, got: %s, want: %s.", templatesSlice[2].Name, name2)
	}
	if templatesSlice[2].Fn != fn2 {
		t.Errorf("Third file name incorrect, got: %s, want: %s.", templatesSlice[2].Fn, fn2)
	}
}
