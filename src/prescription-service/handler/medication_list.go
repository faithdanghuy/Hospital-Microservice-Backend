package handler

import (
	"net/http"

	"github.com/Hospital-Microservice/hospital-core/record"
	"github.com/Hospital-Microservice/hospital-core/transport/http/response"
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/mapper"
	"github.com/Hospital-Microservice/prescription-service/model/req"
	"github.com/labstack/echo/v4"
)

// ListMedications godoc
// @Summary      List Medications with pagination & filter
// @Description  List all medication records with pagination & filter
// @Tags         medication
// @Accept       json
// @Produce      json
// @Param        limit  query int    false "Limit per page"
// @Param        page   query int    false "Page number"
// @Param        sort   query string false "Sort (e.g. drug_name asc)"
// @Param        drug_name query string false "Filter by drug name"
// @Param        unit   query string false "Filter by unit"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  response.ResErr
// @Failure      500  {object}  response.ResErr
// @Router       /medication/filter [get]
func (h *prescriptionHandlerImpl) HandleListMedications(c echo.Context) error {
	p := new(record.Pagination)
	if err := c.Bind(p); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	filter := new(req.MedicationFilterReq)
	if err := c.Bind(filter); err != nil {
		return response.Error(c, http.StatusBadRequest, err.Error())
	}

	result, err := h.listMedicationUseCase.Execute(c.Request().Context(), p, filter)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err.Error())
	}

	// Transform entities to response DTOs
	meds, ok := result.Rows.([]*entity.MedicationEntity)
	if !ok {
		return response.Error(c, http.StatusInternalServerError, "invalid result type")
	}
	resMeds := mapper.TransformMedicationEntitiesToResList(meds)
	result.Rows = resMeds

	return response.OK(c, http.StatusOK, "Success", result)
}
