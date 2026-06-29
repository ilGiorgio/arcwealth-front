package pkg

import (
	"github.com/gomarkdown/markdown/parser"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown"
	"os"
	"html/template"
	"github.com/labstack/echo/v4"
	"app/internal/domain"
	"log"
	"time"
	"fmt"
	"bytes"
	"path/filepath"
	"io/fs"
	"strings"
	"errors"
)

type File struct {
	Title string
	Content template.HTML
	ModTime string
}

// getFile: This get the document specified in the path and retruns it
func getFile(path string) (ast.Node, error) {
	var badRead ast.Node
	input, err := os.ReadFile(path)
	if err != nil {
		return badRead, err
	}
	
	// Markdown parsing
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	doc := markdown.Parse(input, parser.NewWithExtensions(extensions))

	return doc, nil
}

// getFileTitle
// it returns the documents title
func getFileTitle(doc ast.Node) string {
	for _, node := range doc.GetChildren() {
		if heading, ok := node.(*ast.Heading); ok && heading.Level == 2 {
			var buf bytes.Buffer
			for _, child := range heading.GetChildren() {
				if textNode, ok := child.(*ast.Text); ok {
					buf.Write(textNode.Literal)
				}
			}
			return buf.String()
		}
	}
	return "Untitled"
}

const (
	ASSETS = "assets"
	MD = ".md"
	DOCS = "docs/"
	WHOLE_PATH_DOCS = "/docs/"
)

// loadIndexes: this function should instantiate a new Index and return it
func LoadIndexes(element string) domain.Index {
	urlExtention, foundURLExtention := strings.CutPrefix(element, ASSETS)
	if foundURLExtention != true {
		err := errors.New("cutting from this pattern " + ASSETS)
		fmt.Println("Error: ", err)
	}

	url, foundURL := strings.CutSuffix(urlExtention, MD)
	if foundURL != true {
		fmt.Println("ERROR: Getting this pattern", MD)
	}

	path, _ := strings.CutPrefix(element, DOCS)

	doc, err := getFile(path)
	if err != nil {
		fmt.Println("ERROR: There was an error while getting the file", DOCS)
	}

	title := getFileTitle(doc)

	name, foundPattern := strings.CutPrefix(url, WHOLE_PATH_DOCS)
	if foundPattern	!= true {
		fmt.Println("ERROR: There was an error while getting the prefix", WHOLE_PATH_DOCS)
	}

	return domain.Index{
		Name: name,
		Url: url,
		Title: title,
		Selected: false,
	}
}

// directoryContent
// it should traverse all the file inside a directory
// it retruns an Index type
func DirectoryContent(root string, loader func(element string) domain.Index) []domain.Index {
	content := make([]domain.Index, 0, 100)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // stop on error
		}

		if !d.IsDir() {
			_, found := strings.CutSuffix(path, ".md")
			if found {
				content = append(content, loader(path))
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking directory:", err)
	}

	return content
}

/* File's method: It returns the file struct populated */
func (element *File) ParseFile(c echo.Context, fetchFile string) error {
	var filename string

	if fetchFile != "" {
		filename = fetchFile + ".md"
	} else {
		filename = c.Param(ParamFile) + ".md"
	}

	mdPath := ContentPath + filename

	input, err := os.ReadFile(mdPath)
	if err != nil {
		return fmt.Errorf("File was not read: %w", err)
	}

	fileparsed, error := os.Stat(mdPath)

	if error != nil {
		return fmt.Errorf("File was not read: %w", err)
	}

	// Getting the last modification time
	modTime := fileparsed.ModTime()
	modTimeStr := modTime.Format(time.DateOnly)

	// Markdown parsing
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	doc := markdown.Parse(input, parser.NewWithExtensions(extensions))

	// Preparing HTML
	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})
	htmlOutput := markdown.Render(doc, renderer)

	// Populating fields
	element.Content = template.HTML(htmlOutput)
	element.ModTime = modTimeStr
	element.Title = getFileTitle(doc)

	// No errors, returns nil
	return nil
}

func ParseMarkdown(c echo.Context, fetchFile string, path string) (domain.FileInfo, error) {
	var fileInfo domain.FileInfo
	var filename string

	if fetchFile != "" {
		filename = fetchFile + ".md"
	} else {
		filename = c.Param(ParamFile) + ".md"
	}

	mdPath := path + filename

	// getting the file
	doc, err := getFile(mdPath)
	if err != nil {
		return fileInfo, err
	}

	fileparsed, error := os.Stat(mdPath)
	if error != nil {
		log.Fatal(error)
	}

	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})
	htmlOutput := markdown.Render(doc, renderer)

	// Getting the last modification time
	modTime := fileparsed.ModTime()

	fileInfo.Content = template.HTML(htmlOutput)
	modTimeStr := modTime.Format(time.DateOnly)
	fileInfo.ModTime = modTimeStr

	return fileInfo, nil
}
