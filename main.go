package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type projects struct {
	Id        int
	StrNama   string
	StrDurasi int64
	StrSdate  string
	StrEdate  string
	StrDesc   string
	StrCBrejs string
	StrCBnojs string
	StrCBbs   string
	StrCBlar  string
}

var dataprojects = []projects{
	{
		Id:        0,
		StrNama:   "bagza",
		StrDurasi: 30,
		StrSdate:  "2015-09-01",
		StrEdate:  "2015-09-02",
		StrDesc:   "hallo world",
		StrCBrejs: "fa-brands fa-react fa-2xl ",
		StrCBnojs: "fa-brands fa-node fa-2xl ",
		StrCBbs:   "fa-brands fa-bootstrap fa-2xl ",
		StrCBlar:  "fa-brands fa-laravel fa-2xl ",
	},
	{
		Id:        1,
		StrNama:   "jaya",
		StrDurasi: 30,
		StrSdate:  "2015-09-01",
		StrEdate:  "2015-09-02",
		StrDesc:   "hallo GUYSSSSSSSSS",
		StrCBrejs: "fa-brands fa-react fa-2xl ",
		StrCBnojs: "fa-brands fa-node fa-2xl ",
		StrCBbs:   "fa-brands fa-bootstrap fa-2xl ",
		StrCBlar:  "fa-brands fa-laravel fa-2xl ",
	},
}

func main() {
	// fmt.Println(dataprojects[0].Nama)
	e := echo.New()

	e.Static("/assets", "assets") // kegunaan static lock lokasi folder assets untuk bisa di gunakan di semua akses

	//e.GET("/hallo", halloWorld) // rout & handler func
	//e.GET("/json", aboutMe)

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/from-project", formproject) // form add rpoject
	e.GET("/project", project)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/project-update/:id", fromProjectUpdate) //form update project
	e.GET("/testimonial", testimonial)
	e.POST("/addproject", addProject)       // proses add project
	e.POST("/updateproject", updateProject) // proses update project
	e.POST("/delete-project/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func project(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	alldataproject := map[string]interface{}{
		"project": dataprojects,
	}

	return tmpl.Execute(c.Response(), alldataproject)

}

// handler / controler (di php)
// func halloWorld(c echo.Context) error {
// 	return c.String(http.StatusOK, "hallo world")
// }

// func aboutMe(c echo.Context) error {
// 	return c.JSON(http.StatusOK, map[string]string{
// 		"name":    "Bagza",
// 		"address": "Bandung",
// 	})
// }

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

func formproject(c echo.Context) error {
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

	idconvert, _ := strconv.Atoi(id)

	projectdetailawal := projects{}

	for index, data := range dataprojects {
		if index == idconvert {
			projectdetailawal = projects{
				Id:        data.Id,
				StrNama:   data.StrNama,
				StrDurasi: data.StrDurasi,
				StrSdate:  data.StrSdate,
				StrEdate:  data.StrEdate,
				StrDesc:   data.StrDesc,
				StrCBrejs: data.StrCBrejs,
				StrCBnojs: data.StrCBnojs,
				StrCBbs:   data.StrCBbs,
				StrCBlar:  data.StrCBlar,
			}
		}
	}

	// test := projects{
	// 	StrNama:   "bagza",
	// 	StrDurasi: "30 Hari",
	// 	StrDesc:   "hahahahahahah",
	// }

	projectDetail := map[string]interface{}{ //interface in anything
		"Id":      id,
		"project": projectdetailawal,
	}

	return tmpl.Execute(c.Response(), projectDetail)

}

func fromProjectUpdate(c echo.Context) error {
	id := c.Param("id")

	idconvert, _ := strconv.Atoi(id)

	projectupdateawal := projects{}

	for index, data := range dataprojects {
		if index == idconvert {
			projectupdateawal = projects{
				Id:        data.Id,
				StrNama:   data.StrNama,
				StrDurasi: data.StrDurasi,
				StrSdate:  data.StrSdate,
				StrEdate:  data.StrEdate,
				StrDesc:   data.StrDesc,
				StrCBrejs: data.StrCBrejs,
				StrCBnojs: data.StrCBnojs,
				StrCBbs:   data.StrCBbs,
				StrCBlar:  data.StrCBlar,
			}
		}
	}

	// test := projects{
	// 	StrNama:   "bagza",
	// 	StrDurasi: "30 Hari",
	// 	StrDesc:   "hahahahahahah",
	// }

	projectUpdate := map[string]interface{}{ //interface in anything
		"Id":      idconvert,
		"project": projectupdateawal,
	}

	tmpl, err := template.ParseFiles("views/updateproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), projectUpdate)

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

	// now := time.Time(sdate)

	// fmt.Println(now.Second())

	// stime := sdate
	// etime := edate

	waktuStart, _ := time.Parse("2006-01-02", sdate)
	waktuEnd, _ := time.Parse("2006-01-02", edate)

	var difference time.Duration

	if waktuStart.Month() > waktuEnd.Month() || waktuStart.Day() > waktuEnd.Day() {
		difference = waktuStart.Sub(waktuEnd)
	} else if waktuStart.Month() < waktuEnd.Month() || waktuStart.Day() < waktuEnd.Day() {
		difference = waktuEnd.Sub(waktuStart)
	}

	// difference := waktuEnd.Sub(waktuStart)

	hari := int64(difference.Hours() / 24)

	// fmt.Println("durasi ini hari: ", hari)
	//fmt.Println(waktuEnd)

	if rejs == "reactjs" {
		irejs := "fa-brands fa-react fa-2xl"
		rejs = irejs
		// fmt.Println("HAHAHHA BERHASIL")
	} else {
		rejs = ""
	}

	if nojs == "nodejs" {
		inojs := "fa-brands fa-node fa-2xl"
		nojs = inojs
	} else {
		rejs = ""
	}

	if bs == "bootstrap" {
		ibs := "fa-brands fa-bootstrap fa-2xl"
		bs = ibs
	} else {
		rejs = ""
	}

	if lar == "laravel" {
		ilar := "fa-brands fa-laravel fa-2xl"
		lar = ilar
	} else {
		rejs = ""
	}

	//append

	newProject := projects{
		StrNama:   name,
		StrDurasi: hari,
		StrSdate:  sdate,
		StrEdate:  edate,
		StrDesc:   desc,
		StrCBrejs: rejs,
		StrCBnojs: nojs,
		StrCBbs:   bs,
		StrCBlar:  lar,
	}

	dataprojects = append(dataprojects, newProject) // timpa dataprojets yang ada di atas

	fmt.Println("ini adalah name :", name)
	fmt.Println("ini adalah sdate :", sdate)
	fmt.Println("ini adalah edate :", edate)
	fmt.Println("durasi ini hari: ", hari)
	fmt.Println("ini adalah desc :", desc)
	fmt.Println("ini adalah rejs :", rejs)
	fmt.Println("ini adalah nojs :", nojs)
	fmt.Println("ini adalah bs :", bs)
	fmt.Println("ini adalah lar :", lar)

	return c.Redirect(http.StatusMovedPermanently, "/project")

}

func updateProject(c echo.Context) error {
	id := c.Param("id")

	idconvert, _ := strconv.Atoi(id)

	idedit, _ := strconv.Atoi(c.FormValue("id"))
	name := c.FormValue("name")
	sdate := c.FormValue("sdate")
	edate := c.FormValue("edate")
	desc := c.FormValue("desc")
	rejs := c.FormValue("rejs")
	nojs := c.FormValue("nojs")
	bs := c.FormValue("bs")
	lar := c.FormValue("lar")

	// now := time.Time(sdate)

	// fmt.Println(now.Second())

	// stime := sdate
	// etime := edate

	waktuStart, _ := time.Parse("2006-01-02", sdate)
	waktuEnd, _ := time.Parse("2006-01-02", edate)

	var difference time.Duration

	if waktuStart.Month() > waktuEnd.Month() || waktuStart.Day() > waktuEnd.Day() {
		difference = waktuStart.Sub(waktuEnd)
	} else if waktuStart.Month() < waktuEnd.Month() || waktuStart.Day() < waktuEnd.Day() {
		difference = waktuEnd.Sub(waktuStart)
	}

	// difference := waktuEnd.Sub(waktuStart)

	hari := int64(difference.Hours() / 24)

	// fmt.Println("durasi ini hari: ", hari)
	//fmt.Println(waktuEnd)

	if rejs == "reactjs" {
		irejs := "fa-brands fa-react fa-2xl"
		rejs = irejs
		// fmt.Println("HAHAHHA BERHASIL")
	} else {
		rejs = ""
	}

	if nojs == "nodejs" {
		inojs := "fa-brands fa-node fa-2xl"
		nojs = inojs
	} else {
		rejs = ""
	}

	if bs == "bootstrap" {
		ibs := "fa-brands fa-bootstrap fa-2xl"
		bs = ibs
	} else {
		rejs = ""
	}

	if lar == "laravel" {
		ilar := "fa-brands fa-laravel fa-2xl"
		lar = ilar
	} else {
		rejs = ""
	}

	// fmt.Println("hayu id", idconvert)

	// if idedit == 0 {
	// 	idedit = idconvert
	// }

	// fmt.Println(idedit)

	dataprojects[idedit].Id = idconvert
	dataprojects[idedit].StrNama = name
	dataprojects[idedit].StrDurasi = hari
	dataprojects[idedit].StrSdate = sdate
	dataprojects[idedit].StrEdate = edate
	dataprojects[idedit].StrDesc = desc
	dataprojects[idedit].StrCBrejs = rejs
	dataprojects[idedit].StrCBnojs = nojs
	dataprojects[idedit].StrCBbs = bs
	dataprojects[idedit].StrCBlar = lar

	//append

	// newProject := projects{
	// 	StrNama:   name,
	// 	StrDurasi: hari,
	// 	StrSdate:  sdate,
	// 	StrEdate:  edate,
	// 	StrDesc:   desc,
	// 	StrCBrejs: rejs,
	// 	StrCBnojs: nojs,
	// 	StrCBbs:   bs,
	// 	StrCBlar:  lar,
	// }

	//dataprojects = append(dataprojects, newProject) // timpa dataprojets yang ada di atas

	fmt.Println("ini adalah name :", name)
	fmt.Println("ini adalah sdate :", sdate)
	fmt.Println("ini adalah edate :", edate)
	fmt.Println("durasi ini hari: ", hari)
	fmt.Println("ini adalah desc :", desc)
	fmt.Println("ini adalah rejs :", rejs)
	fmt.Println("ini adalah nojs :", nojs)
	fmt.Println("ini adalah bs :", bs)
	fmt.Println("ini adalah lar :", lar)

	return c.Redirect(http.StatusMovedPermanently, "/project")

}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idconvertchar, _ := strconv.Atoi(id)

	// APPEND

	dataprojects = append(dataprojects[:idconvertchar], dataprojects[idconvertchar+1:]...) // timpa dataprojets yang ada di atas

	fmt.Println("ini adalah id delete :", id)

	return c.Redirect(http.StatusMovedPermanently, "/project")

}
