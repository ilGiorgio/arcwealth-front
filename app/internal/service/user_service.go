package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"app/internal/pkg"

	"app/internal/domain"

	"github.com/labstack/echo/v4"
)

// LandingPage: It returns the landing HTML page
func SignupPage(c echo.Context) error {
	uri := pkg.Currencies
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Get(uri)
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusOK, pkg.TemplateSignup, "")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return c.Render(http.StatusOK, pkg.TemplateSignup, "")
	}
	var currencies domain.ApiResponse[[]domain.Currency]

	err = json.NewDecoder(response.Body).Decode(&currencies)

	fmt.Println(currencies.Data)

	return c.Render(http.StatusOK, pkg.TemplateSignup, currencies.Data)
}

// RegisterUser: It returns the register HTML page
func RegisterUser(c echo.Context) error {
	var form domain.RegisterUserRequest

	if err := c.Bind(&form); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	fmt.Printf("User: %+v\n", form)

	if err := c.Validate(form); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	jsonData, err := json.Marshal(form)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error while parsing json data")
	}

	uri := pkg.SignupURI
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Post(uri, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	fmt.Println("status: ", response.StatusCode)

	return c.Render(http.StatusOK, pkg.TemplateModalSuccess, "")
}
