package service

import (
	"textnet/http"

	"app/internal/domain"
	"app/internal/pkg"

	"github.com/labstack/echo/v4"
)

// LandingPage: It returns the landing HTML page
func LandingPage(c echo.Context) error {
	return c.Render(http.StatusOK, pkg.TemplateIndex, "")
}

// DocsModule: It parses the content of the entire docs directory and build in memory the sidebar index
// It also gets the content of each file and parses it to HTML
func DocsModule(c echo.Context) error {
	var currentIndex domain.Index
	var test pkg.File

	pathVariable := c.Param(pkg.ParamFile)

	if pathVariable == "" {
		pathVariable = pkg.TemplateIntroduction
	}

	fail := test.ParseFile(c, pathVariable)

	if fail != nil {
		return c.String(http.StatusNotFound, "Document not found")
	}

	data, err := pkg.ParseMarkdown(c, pathVariable, pkg.ContentPath)
	if err != nil {
		return c.String(http.StatusNotFound, "Document not found")
	}

	indexesBar := pkg.SidebarIndexes()

	for i := 0; i < len(indexesBar); i++ {
		if pathVariable == indexesBar[i].Name {
			indexesBar[i].Selected = true
			currentIndex = indexesBar[i]
		}
	}

	current := domain.NewCurrent(data.ModTime, data.Content, currentIndex)
	metafield := domain.CreateMetafield("docs", indexesBar)
	bundle := domain.NewData(current, metafield)

	if c.Request().Header.Get("HX-request") == "true" {
		return c.Render(http.StatusOK, pkg.TemplatePortal, bundle)
	}

	return c.Render(http.StatusOK, pkg.TemplateDocuments, bundle)
}

// ReferencesModule: This gets the file content and pareses it to HTML
func ReferencesModule(c echo.Context) error {
	data, _ := pkg.ParseMarkdown(c, "/references", "assets/references")

	if c.Request().Header.Get("HX-request") == "true" {
		return c.Render(http.StatusOK, pkg.TemplateWindow, data)
	}

	return c.Render(http.StatusOK, pkg.TemplateFrame, data)
}

// HistoriesModule: This gets the file content and pareses it to HTML
func HistoriesModule(c echo.Context) error {
	data, _ := pkg.ParseMarkdown(c, "/histories", "assets/histories")

	if c.Request().Header.Get("HX-request") == "true" {
		return c.Render(http.StatusOK, pkg.TemplateWindow, data)
	}

	return c.Render(http.StatusOK, pkg.TemplateFrame, data)
}

// BacklogModule: This gets the file content and pareses it to HTML
func BacklogModule(c echo.Context) error {
	data, _ := pkg.ParseMarkdown(c, "/backlog", "assets/backlog")

	if c.Request().Header.Get("HX-request") == "true" {
		return c.Render(http.StatusOK, pkg.TemplateWindow, data)
	}

	return c.Render(http.StatusOK, pkg.TemplateFrame, data)
}
