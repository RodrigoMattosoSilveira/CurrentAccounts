package utilities

import "testing"

func TestNew(t *testing.T) {
	templateFiles := NewTemplateFiles("layout.html")
	if len(templateFiles.Files) != 1 {
		t.Errorf("Expected 1 file, got %d", len(templateFiles.Files))
	}
	if templateFiles.Files[0] != "layout.html" {
		t.Errorf("First file incorrect, got: %s, want: %s.", templateFiles.Files[0], "layout.html")
	}
}

func TestNewTemplateFilesLayout(t *testing.T) {
	templateFiles := NewTemplateFilesLayout("header.html")
	if len(templateFiles.Files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(templateFiles.Files))
	}
	if templateFiles.Files[0] != "root/layout.tmpl" {
		t.Errorf("First file incorrect, got: %s, want: %s.", templateFiles.Files[0], "root/layout.tmpl")
	}
	if templateFiles.Files[1] != "header.html" {
		t.Errorf("Second file incorrect, got: %s, want: %s.", templateFiles.Files[1], "header.html")
	}
}
func TestAdd(t *testing.T) {
	templateFiles := NewTemplateFiles("layout.html")
	templateFiles.AddTemplate("header.html")
	templateFiles.AddTemplate("footer.html")
	templateFileLen := len(templateFiles.Files)
	if templateFileLen != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", templateFileLen, 3)
	}
	if templateFiles.Files[1] != "header.html" {
		t.Errorf("Second file incorrect, got: %s, want: %s.", templateFiles.Files[1], "header.html")
	}
	if templateFiles.Files[2] != "footer.html" {
		t.Errorf("Third file incorrect, got: %s, want: %s.", templateFiles.Files[2], "footer.html")
	}
}

func TestGetAll(t *testing.T) {
	templateFiles := NewTemplateFiles("layout.html")
	templateFiles.AddTemplate("header.html")
	templateFiles.AddTemplate("footer.html")
	templateFilesAll := templateFiles.GetAllTemplates()
	if len(templateFilesAll) != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", len(templateFilesAll), 3)
	}
	expectedFiles := []string{"layout.html", "header.html", "footer.html"}
	for i, file := range expectedFiles {
		if templateFilesAll[i] != file {
			t.Errorf("File at index %d incorrect, got: %s, want: %s.", i, templateFilesAll[i], file)
		}
	}
}
