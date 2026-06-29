package pkg

import (
	"app/internal/domain"
)

// Available Templates
const (
	TemplateDocuments    = "documents"
	TemplateIndex        = "index"
	TemplateSignup       = "signup"
	TemplateIntroduction = "introduction"
	TemplatePortal       = "portal"
	TemplateFrame        = "frame"
	TemplateWindow       = "window"
)

// URL Params
const (
	ParamFile = "file" // DocsModule
)

// General
const (
	ContentPath = "assets/docs/"
)

// gettingDocsIndexes: gets the indexes inside the doc directory
func gettingDocsIndexes() []domain.Index {
	root := "./assets/docs" // your base directory
	return DirectoryContent(root, LoadIndexes)
}

func SidebarIndexes() []domain.Index {
	return gettingDocsIndexes()
}
