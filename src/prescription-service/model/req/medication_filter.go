package req

type MedicationFilterReq struct {
	DrugName string `json:"drug_name,omitempty"`
	Unit     string `json:"dosage,omitempty"`
	Page     int    `json:"page,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Sort     string `json:"sort,omitempty"`
}
