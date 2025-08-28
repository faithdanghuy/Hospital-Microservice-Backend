package req

type MedicationUpdateReq struct {
	DrugName    *string `json:"drug_name,omitempty"`
	Stock       *int    `json:"stock,omitempty"`
	Unit        *string `json:"unit" validate:"required,oneof=tablet capsule syrup injection ointment drop inhaler patch suppository other"`
	Description *string `json:"description,omitempty"`
}
