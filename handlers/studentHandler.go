package handlers

import (
	db "hserver/dblayer"
	"hserver/jtypes"
	"hserver/mappers"
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct{}

func (h *Handler) GetStudentDetails(c echo.Context) error {
	studentDetails, err := db.GetStudentDetails()

	if err != nil {
		c.Error(err)
	}

	return c.JSON(http.StatusOK, studentDetails)
}

func (h *Handler) PostGrievance(c echo.Context) error {
	studentGrievance := &jtypes.Grievance{}

	if err := c.Bind(studentGrievance); err != nil {
		c.Error(err)
	}

	_, err := db.GetStudentDetails()

	if err != nil {
		c.Error(err)
		return err
	}

	pgGrievance := mappers.MapJGrievanceToPgGrievance(studentGrievance)

	err = db.CreateStudentGrievance(pgGrievance)

	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, "Grievance Successfully submitted")
}

func (h *Handler) GetStudentGrievances(c echo.Context) error {
	_, err := db.GetStudentDetails()

	if err != nil {
		return err
	}

	res, err := db.GetStudentGrievances(1)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
