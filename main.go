package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets") // kegunaan static lock lokasi folder assets untuk bisa di gunakan di semua akses

	e.GET("/hallo", halloWorld) // rout & handler func
	e.GET("/json", aboutMe)

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/project", project)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/testimonial", testimonial)
	e.POST("/addproject", addProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// handler / controler (di php)
func halloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "hallo world")
}

func aboutMe(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"name":    "Bagza",
		"address": "Bandung",
	})
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)

}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/form.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)

}

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/addproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)

}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)

}

func projectDetail(c echo.Context) error {
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	projectDetail := map[string]interface{}{ //interface in anything
		"Id":          id,
		"Title":       "Percobaan pertama",
		"Description": "disini adalah descripsi pertama, Lorem ipsum dolor sit amet consectetur adipisicing elit. Quia id et ad velit rerum molestiae, vero eum enim esse labore, deserunt nostrum modi consectetur commodi libero illo sapiente. Ipsam, tempore.",
	}

	return tmpl.Execute(c.Response(), projectDetail)

}

func addProject(c echo.Context) error {
	name := c.FormValue("name")
	sdate := c.FormValue("sdate")
	edate := c.FormValue("edate")
	desc := c.FormValue("desc")
	rejs := c.FormValue("rejs")
	nojs := c.FormValue("nojs")
	bs := c.FormValue("bs")
	lar := c.FormValue("lar")

	if rejs == "reactjs" {
		irejs := `<i class="fa-brands fa-react fa-2xl marlef"></i>`
		rejs = irejs
		// fmt.Println("HAHAHHA BERHASIL")
	}

	if nojs == "nodejs" {
		inojs := `<i class="fa-brands fa-node fa-2xl marlef"></i>`
		nojs = inojs
	}

	if bs == "bootstrap" {
		ibs := `<i class="fa-brands fa-bootstrap fa-2xl marlef"></i>`
		bs = ibs
	}

	if lar == "laravel" {
		ilar := `<i class="fa-brands fa-laravel fa-2xl marlef"></i>`
		lar = ilar
	}

	fmt.Println("ini adalah name :", name)
	fmt.Println("ini adalah sdate :", sdate)
	fmt.Println("ini adalah edate :", edate)
	fmt.Println("ini adalah desc :", desc)
	fmt.Println("ini adalah rejs :", rejs)
	fmt.Println("ini adalah nojs :", nojs)
	fmt.Println("ini adalah bs :", bs)
	fmt.Println("ini adalah lar :", lar)

	return c.Redirect(http.StatusMovedPermanently, "/project")

}
