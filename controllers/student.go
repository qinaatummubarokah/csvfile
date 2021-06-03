package controllers

import (
	"crudgolang/models"
	"encoding/csv"
	"log"
	"net/http"
	"os"

	// "strconv"

	"github.com/dnlo/struct2csv"
	"github.com/labstack/echo"
)

func GetStudent(c echo.Context) error {
	idStr := c.Param("id")
	var student models.Student
	status := models.GetStudent(&student, idStr)
	log.Println(status)
	if status == http.StatusOK {
		return c.JSON(status, student)
	} else {
		return echo.NewHTTPError(status)
	}
}

func GetData(c echo.Context) error {
	idStr := c.Param("id")
	var data models.Data
	status := models.GetData(&data, idStr)
	log.Println(status)
	if status == http.StatusOK {
		return c.JSON(status, data)
	} else {
		return echo.NewHTTPError(status)
	}
}

func GetAllStudent(c echo.Context) error {
	var err error

	var students []models.Student
	total, err := models.GetAllStudent(&students, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if total == 0 {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	// result := lib.Paginate(c, categories, total, limit, page, offset, pagination)

	return c.JSON(http.StatusOK, students)
}

func GetAllData(c echo.Context) error {
	var err error

	var data []models.Data
	err = models.GetAllData(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	// result := lib.Paginate(c, categories, total, limit, page, offset, pagination)

	return c.JSON(http.StatusOK, data)
}

func CreateStudent(c echo.Context) error {
	params := make(map[string]string)
	// Get parameter id
	id := c.FormValue("id")
	if id != "" {
		params["id"] = id
	} else {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Get parameter name
	name := c.FormValue("name")
	if name != "" {
		params["name"] = name
	} else {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Get parameter age
	age := c.FormValue("age")
	if age != "" {
		params["age"] = age
	} else {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Get parameter grade
	grade := c.FormValue("grade")
	if grade != "" {
		params["grade"] = grade
	} else {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	status := models.CreateStudent(params)
	return echo.NewHTTPError(status)
}

func UpdateStudent(c echo.Context) error {
	params := make(map[string]string)
	// Get parameter id
	id := c.Param("id")
	if id != "" {
		params["id"] = id
	} else {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Get parameter name
	name := c.FormValue("name")
	if name != "" {
		params["name"] = name
	}

	// Get parameter age
	age := c.FormValue("age")
	if age != "" {
		params["age"] = age
	}

	// Get parameter grade
	grade := c.FormValue("grade")
	if grade != "" {
		params["grade"] = grade
	}

	status := models.UpdateStudent(params)
	return echo.NewHTTPError(status)
}

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		status := models.DeleteStudent(id)
		return echo.NewHTTPError(status)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
}

func Reportcsv(c echo.Context) error {
	var datas []models.Data
	err := models.GetAllData(&datas)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	csvfile, err := os.Create("report.csv")
	w := csv.NewWriter(csvfile)
	enc := struct2csv.New()
	colhdrs, err := enc.GetColNames(datas[0])
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	w.Write(colhdrs)
	// get the data from each struct
	for _, v := range datas {
		row, err := enc.GetRow(v)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		w.Write(row)
	}
	w.Flush()

	csvfile.Close()
	return c.Attachment("report.csv", "report.csv")
}
