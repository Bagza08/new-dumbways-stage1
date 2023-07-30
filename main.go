package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"personal-web/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type projects struct {
	Id      int
	StrNama string
	//StrDurasi int64
	StrSdate time.Time
	StrEdate time.Time
	StrDesc  string
	Fix      []string
	StrImg   string
	//StrAuthorId int
}

type user struct {
	Id           int
	UsNama       string
	UsEmail      string
	UsExperience []string
	UsYear       []string
}

// type hayu [4]string {
// 	Satu  string
// 	Dua   string
// 	Tiga  string
// 	Empat string
// }

// type tech struct {
// 	TechCBrejs string
// 	TechBnojs  string
// 	TechCBbs   string
// 	TechCBlar  string
// }

var dataprojects = []projects{
	{
		//Id:        0,
		StrNama: "bagza",
		//StrDurasi: 30,
		StrSdate: time.Now(),
		StrEdate: time.Now(),
		StrDesc:  "hallo world",
		// StrCBrejs: "fa-brands fa-react fa-2xl ",
		// StrCBnojs: "fa-brands fa-node fa-2xl ",
		// StrCBbs:   "fa-brands fa-bootstrap fa-2xl ",
		// StrCBlar:  "fa-brands fa-laravel fa-2xl ",
	},
	{
		//Id:        1,
		StrNama: "jaya",
		//StrDurasi: 30,
		StrSdate: time.Now(),
		StrEdate: time.Now(),
		StrDesc:  "hallo GUYSSSSSSSSS",
		// StrCBrejs: "fa-brands fa-react fa-2xl ",
		// StrCBnojs: "fa-brands fa-node fa-2xl ",
		// StrCBbs:   "fa-brands fa-bootstrap fa-2xl ",
		// StrCBlar:  "fa-brands fa-laravel fa-2xl ",
	},
}

func main() {
	// fmt.Println(dataprojects[0].Nama)
	e := echo.New()

	connection.DatabaseConn()

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

	query_projects, err_query_projects := connection.Conn.Query(context.Background(), "SELECT id, name,start_date,end_date,description,technologies,image FROM tb_project")

	if err_query_projects != nil {
		return c.JSON(http.StatusInternalServerError, err_query_projects.Error())
	}

	var result_projects []projects

	for query_projects.Next() {
		var each = projects{}

		err := query_projects.Scan(&each.Id, &each.StrNama, &each.StrSdate, &each.StrEdate, &each.StrDesc, &each.Fix, &each.StrImg)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// tech_fix := []string{"fa-brands fa-react fa-2xl", "fa-brands fa-node fa-2xl", "fa-brands fa-bootstrap fa-2xl", "fa-brands fa-laravel fa-2xl"}

		// each.Fix = tech_fix

		result_projects = append(result_projects, each)

	}

	alldataproject := map[string]interface{}{
		"project": result_projects,
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

	userid := 3

	var dataUser = user{}

	connection.Conn.QueryRow(context.Background(), "SELECT id,name,email,experience, year FROM tb_user WHERE id=$1", userid).Scan(&dataUser.Id, &dataUser.UsNama, &dataUser.UsEmail, &dataUser.UsExperience, &dataUser.UsYear)

	var datarespon map[string]interface{}

	if dataUser.Id == 0 {
		dataUser = user{
			Id:           0,
			UsNama:       "UNKNOWN",
			UsEmail:      "",
			UsExperience: []string{"NO EXPERIENCE"},
			UsYear:       []string{"NO YEAR"},
		}
		datarespon = map[string]interface{}{
			"user": dataUser,
		}
	} else {
		datarespon = map[string]interface{}{
			"user": dataUser,
		}
	}

	fmt.Println("data user", datarespon)

	return tmpl.Execute(c.Response(), datarespon)

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

	//ambil 1 data menggunakan query
	errqueryrow := connection.Conn.QueryRow(context.Background(), "SELECT id,name,start_date,end_date,description,technologies,image FROM tb_project WHERE id=$1", idconvert).Scan(&projectdetailawal.Id, &projectdetailawal.StrNama, &projectdetailawal.StrSdate, &projectdetailawal.StrEdate, &projectdetailawal.StrDesc, &projectdetailawal.Fix, &projectdetailawal.StrImg)

	// fmt.Println("data awal", errqueryrow)

	if errqueryrow != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// for index, data := range dataprojects {
	// 	if index == idconvert {
	// 		projectdetailawal = projects{
	// 			//Id:        data.Id,
	// 			StrNama: data.StrNama,
	// 			//StrDurasi: data.StrDurasi,
	// 			StrSdate: data.StrSdate,
	// 			StrEdate: data.StrEdate,
	// 			StrDesc:  data.StrDesc,
	// 			// StrCBrejs: data.StrCBrejs,
	// 			// StrCBnojs: data.StrCBnojs,
	// 			// StrCBbs:   data.StrCBbs,
	// 			// StrCBlar:  data.StrCBlar,
	// 		}
	// 	}
	// }

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
	irejs := "fa-brands fa-react fa-2xl"
	inojs := "fa-brands fa-node fa-2xl"
	ibs := "fa-brands fa-bootstrap fa-2xl"
	ilar := "fa-brands fa-laravel fa-2xl"
	tech_fix := []string{}

	if rejs == "reactjs" {
		rejs = irejs
		tech_fix = []string{rejs, "", "", ""}
		//fmt.Println("HAHAHHA BERHASIL", tech_fix)
	}

	if nojs == "nodejs" && rejs == irejs {
		nojs = inojs
		tech_fix = []string{rejs, nojs, "", ""}
		//tech_fix := []string{rejs, nojs, "", ""}
	} else if nojs == "nodejs" && rejs == "" {
		nojs = inojs
		tech_fix = []string{"", nojs, "", ""}
	}

	if bs == "bootstrap" && rejs == irejs && nojs == inojs {
		bs = ibs
		tech_fix = []string{rejs, nojs, bs, ""}
		//tech_fix := []string{rejs, nojs, "", ""}
	} else if bs == "bootstrap" && rejs == "" && nojs == inojs {
		bs = ibs
		tech_fix = []string{"", nojs, bs, ""}
	} else if bs == "bootstrap" && rejs == irejs && nojs == "" {
		bs = ibs
		tech_fix = []string{rejs, "", bs, ""}
	} else if bs == "bootstrap" && rejs == "" && nojs == "" {
		bs = ibs
		tech_fix = []string{"", "", bs, ""}
	}

	if lar == "laravel" && bs == ibs && rejs == irejs && nojs == inojs {
		lar = ilar
		tech_fix = []string{rejs, nojs, bs, lar}
	} else if lar == "laravel" && bs == "" && rejs == irejs && nojs == inojs {
		lar = ilar
		tech_fix = []string{rejs, nojs, "", lar}
	} else if lar == "laravel" && bs == ibs && rejs == "" && nojs == inojs {
		lar = ilar
		tech_fix = []string{"", nojs, bs, lar}
	} else if lar == "laravel" && bs == ibs && rejs == irejs && nojs == "" {
		lar = ilar
		tech_fix = []string{rejs, "", bs, lar}
	} else if lar == "laravel" && bs == "" && rejs == "" && nojs == "" {
		lar = ilar
		tech_fix = []string{"", "", "", lar}
	} else if lar == "laravel" && bs == "" && rejs == "" && nojs == inojs {
		lar = ilar
		tech_fix = []string{"", nojs, "", lar}
	}

	//append

	// newProject := projects{
	// 	StrNama: name,
	// 	//StrDurasi: hari,
	// 	StrSdate: time.Now(),
	// 	StrEdate: time.Now(),
	// 	StrDesc:  desc,
	// 	// StrCBrejs: rejs,
	// 	// StrCBnojs: nojs,
	// 	// StrCBbs:   bs,
	// 	// StrCBlar:  lar,
	// }

	_, errex := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (name,start_date,end_date,description,technologies,image) VALUES ($1, $2, $3, $4, $5, $6)", name, sdate, edate, desc, tech_fix, "contoh.jpg")

	if errex != nil {
		return c.JSON(http.StatusInternalServerError, errex.Error())
	}

	//dataprojects = append(dataprojects) // timpa dataprojets yang ada di atas

	// cekk masuk variabel gak ????

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

func fromProjectUpdate(c echo.Context) error {
	id := c.Param("id")

	idconvert, _ := strconv.Atoi(id)

	//projectupdateawal := projects{}

	tmpl, err := template.ParseFiles("views/updateproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	projectUpdate := map[string]interface{}{ //interface in anything
		"Id": idconvert,
	}

	return tmpl.Execute(c.Response(), projectUpdate)

	// for index, data := range dataprojects {
	// 	if index == idconvert {
	// 		projectupdateawal = projects{
	// 			//Id:        data.Id,
	// 			StrNama: data.StrNama,
	// 			//StrDurasi: data.StrDurasi,
	// 			StrSdate: data.StrSdate,
	// 			StrEdate: data.StrEdate,
	// 			StrDesc:  data.StrDesc,
	// 			// StrCBrejs: data.StrCBrejs,
	// 			// StrCBnojs: data.StrCBnojs,
	// 			// StrCBbs:   data.StrCBbs,
	// 			// StrCBlar:  data.StrCBlar,
	// 		}
	// 	}
	// }

	// // test := projects{
	// // 	StrNama:   "bagza",
	// // 	StrDurasi: "30 Hari",
	// // 	StrDesc:   "hahahahahahah",
	// // }
}

func updateProject(c echo.Context) error {
	//id := c.Param("id")

	//idconvert, _ := strconv.Atoi(id)

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

	irejs := "fa-brands fa-react fa-2xl"
	inojs := "fa-brands fa-node fa-2xl"
	ibs := "fa-brands fa-bootstrap fa-2xl"
	ilar := "fa-brands fa-laravel fa-2xl"

	tech_fix := []string{}

	if rejs == "reactjs" {
		rejs = irejs
		tech_fix = []string{rejs, "", "", ""}
		//fmt.Println("HAHAHHA BERHASIL", tech_fix)
	}

	if nojs == "nodejs" && rejs == irejs {
		nojs = inojs
		tech_fix = []string{rejs, nojs, "", ""}
		//tech_fix := []string{rejs, nojs, "", ""}
	} else if nojs == "nodejs" && rejs == "" {
		nojs = inojs
		tech_fix = []string{"", nojs, "", ""}
	}

	if bs == "bootstrap" && rejs == irejs && nojs == inojs {
		bs = ibs
		tech_fix = []string{rejs, nojs, bs, ""}
		//tech_fix := []string{rejs, nojs, "", ""}
	} else if bs == "bootstrap" && rejs == "" && nojs == inojs {
		bs = ibs
		tech_fix = []string{"", nojs, bs, ""}
	} else if bs == "bootstrap" && rejs == irejs && nojs == "" {
		bs = ibs
		tech_fix = []string{rejs, "", bs, ""}
	} else if bs == "bootstrap" && rejs == "" && nojs == "" {
		bs = ibs
		tech_fix = []string{"", "", bs, ""}
	}

	if lar == "laravel" && bs == ibs && rejs == irejs && nojs == inojs {
		lar = ilar
		tech_fix = []string{rejs, nojs, bs, lar}
	} else if lar == "laravel" && bs == "" && rejs == irejs && nojs == inojs {
		lar = ilar
		tech_fix = []string{rejs, nojs, "", lar}
	} else if lar == "laravel" && bs == ibs && rejs == "" && nojs == inojs {
		lar = ilar
		tech_fix = []string{"", nojs, bs, lar}
	} else if lar == "laravel" && bs == ibs && rejs == irejs && nojs == "" {
		lar = ilar
		tech_fix = []string{rejs, "", bs, lar}
	} else if lar == "laravel" && bs == "" && rejs == "" && nojs == "" {
		lar = ilar
		tech_fix = []string{"", "", "", lar}
	} else if lar == "laravel" && bs == "" && rejs == "" && nojs == inojs {
		lar = ilar
		tech_fix = []string{"", nojs, "", lar}
	}

	// fmt.Println("hayu id", idconvert)

	// if idedit == 0 {
	// 	idedit = idconvert
	// }

	// fmt.Println(idedit)

	//dataprojects[idedit].Id = idconvert
	// dataprojects[idedit].StrNama = name
	// //dataprojects[idedit].StrDurasi = hari
	// dataprojects[idedit].StrSdate = time.Now()
	// dataprojects[idedit].StrEdate = time.Now()
	// dataprojects[idedit].StrDesc = desc
	// dataprojects[idedit].StrCBrejs = rejs
	// dataprojects[idedit].StrCBnojs = nojs
	// dataprojects[idedit].StrCBbs = bs
	// dataprojects[idedit].StrCBlar = lar

	fmt.Println("ini adalah name :", name)
	fmt.Println("ini adalah sdate :", sdate)
	fmt.Println("ini adalah edate :", edate)
	fmt.Println("durasi ini hari: ", hari)
	fmt.Println("ini adalah desc :", desc)
	fmt.Println("ini adalah rejs :", rejs)
	fmt.Println("ini adalah nojs :", nojs)
	fmt.Println("ini adalah bs :", bs)
	fmt.Println("ini adalah lar :", lar)

	_, errex := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET name=$1 ,start_date=$2 ,end_date = $3,description=$4,technologies=$5,image=$6 WHERE id=$7", name, sdate, edate, desc, tech_fix, "contoh.jpg", idedit)

	if errex != nil {
		return c.JSON(http.StatusInternalServerError, errex.Error())
	}

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

	return c.Redirect(http.StatusMovedPermanently, "/project")

}

func deleteProject(c echo.Context) error {
	id := c.Param("id")

	idconvertchar, _ := strconv.Atoi(id)

	// APPEND

	//dataprojects = append(dataprojects[:idconvertchar], dataprojects[idconvertchar+1:]...) // timpa dataprojets yang ada di atas

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", idconvertchar)

	fmt.Println("ini adalah id delete :", id)

	return c.Redirect(http.StatusMovedPermanently, "/project")

}
