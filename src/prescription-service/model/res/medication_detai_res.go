package res

type MedicationRes struct {
	ID          string  `json:"id"`
	DrugName    *string `json:"drug_name"`
	Stock       *int    `json:"stock"`
	Unit        *string `json:"unit"`
	Description *string `json:"description,omitempty"`
}
