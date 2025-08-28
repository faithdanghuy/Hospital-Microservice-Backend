package mapper

import (
	"github.com/Hospital-Microservice/prescription-service/entity"
	"github.com/Hospital-Microservice/prescription-service/model/res"
)

func TransformMedicationEntityToRes(med *entity.MedicationEntity) *res.MedicationRes {
	if med == nil {
		return nil
	}
	return &res.MedicationRes{
		ID:          *med.ID,
		DrugName:    med.DrugName,
		Stock:       med.Stock,
		Unit:        med.Unit,
		Description: med.Description,
	}
}

func TransformMedicationEntitiesToResList(meds []*entity.MedicationEntity) []*res.MedicationRes {
	list := make([]*res.MedicationRes, 0, len(meds))
	for _, m := range meds {
		list = append(list, TransformMedicationEntityToRes(m))
	}
	return list
}
