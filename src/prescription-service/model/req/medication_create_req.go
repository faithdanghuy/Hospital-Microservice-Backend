package req

type MedicationCreateReq struct {
	DrugName    *string `json:"drug_name" validate:"required"`
	Stock       *int    `json:"stock" validate:"required"`
	Unit        *string `json:"unit" validate:"required,oneof=tablet capsule syrup injection ointment drop inhaler patch suppository other"`
	Description *string `json:"description,omitempty"`
}
