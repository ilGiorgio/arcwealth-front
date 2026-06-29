package domain

import (
	"html/template"
	// "errors"
)

type Module struct {
	Name string
	Data template.HTML
}

func NewModule(name string, template template.HTML) Module {
	return Module {
		Name: name,
		Data: template,
	}
}

type Page struct {
	Module Module
}

func NewPage(module Module) Page {
	return Page{
		Module: module,
	}
}

type Index struct {
	Name string
	Url string
	Title string
	Selected bool
}

func CreateIndex(a string, b string, c string) Index {
	return Index {
		Name: a,
		Url: b,
		Title: c,
		Selected: false,
	}
}

type Metafield struct {
	Module string
	Indexes []Index
}

func CreateMetafield(a string, b []Index) Metafield {
	return Metafield {
		Module: a,
		Indexes: b,
	}
}

type Current struct {
	UpdatedAt string
	Content template.HTML
	Selected Index
}

func NewCurrent(a string, b template.HTML, c Index) Current {
	return Current {
		UpdatedAt: a,
		Content: b,
		Selected: c,
	}
}

type Data struct {
	Current Current
	Metafield Metafield
}

func NewData(current Current, metafield Metafield) Data {
	return Data {
		Current: current,
		Metafield: metafield,
	}
}

type FileInfo struct {
	Content template.HTML
	ModTime string
}


