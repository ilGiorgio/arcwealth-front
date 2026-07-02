package api

import (
	"html/template"
	"io"

	"app/internal/service"

	"github.com/labstack/echo/v4"
)

// Templates wraps html/template and implements echo.Renderer
type Templates struct {
	templates *template.Template
}

// Render implements echo.Renderer
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseFiles(
			"web/index.html",
			"web/pages/frame.html",
			"web/components/signup.html",
			"web/templates/modals/success.html",
			"web/snippets/introduction.html",
			"web/snippets/portal.html",
			"web/snippets/window.html",
			"web/templates/banner.html",
			"web/templates/footer.html",
			"web/templates/navbar.html",
			"web/templates/sidebar.html",
		)),
	}
}

func RegisterRoutes(e *echo.Echo) {
	e.Renderer = newTemplate()

	e.GET("/", service.LandingPage)

	e.GET("/signup", service.SignupPage)

	e.POST("/register", service.RegisterUser)

	e.GET("/docs/:file", service.DocsModule)

	e.GET("/references", service.ReferencesModule)

	e.GET("/backlog", service.BacklogModule)

	e.GET("/histories", service.HistoriesModule)
}
