package employeewb

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tphume/seatlect-services/internal/commonErr"
	"github.com/tphume/seatlect-services/internal/database/typedb"
	"github.com/tphume/seatlect-services/internal/gen_openapi/employee_api"
	"net/http"
)

type Server struct {
	Repo Repo
}

func (s *Server) GetEmployeeBusinessId(ctx echo.Context, businessId string) error {
	employees, err := s.Repo.ListEmployee(ctx.Request().Context(), businessId)
	if err != nil {
		if err == commonErr.INVALID {
			return ctx.String(http.StatusBadRequest, "Business id is incorrect")
		} else if err == commonErr.NOTFOUND {
			return ctx.String(http.StatusNotFound, "Business not found with given id")
		}

		return ctx.String(http.StatusInternalServerError, "Database error")
	}

	res := employee_api.ListBusinessResponse{Employees: typedbListToOapi(employees)}
	return ctx.JSONPretty(http.StatusOK, res, "  ")
}

func (s *Server) PostEmployeeBusinessId(ctx echo.Context, businessId string) error {
	panic("implement me")
}

func (s *Server) DeleteEmployeeBusinessIdUsername(ctx echo.Context, businessId string, username string) error {
	panic("implement me")
}

type Repo interface {
	ListEmployee(ctx context.Context, businessId string) ([]typedb.Employee, error)
	CreateEmployee(ctx context.Context, businessId string, employee typedb.Employee) error
	DeleteEmployee(ctx context.Context, businessId string, username string) error
}

// Helper function
func typedbListToOapi(employees []typedb.Employee) *[]employee_api.Employee {
	res := make([]employee_api.Employee, len(employees))
	for i, e := range employees {
		res[i] = typedbToOapi(e)
	}

	return &res
}

func typedbToOapi(e typedb.Employee) employee_api.Employee {
	return employee_api.Employee{Username: e.Username, Password: e.Password}
}
