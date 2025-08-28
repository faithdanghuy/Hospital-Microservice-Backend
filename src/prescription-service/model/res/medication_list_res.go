package res

type MedicationListRes struct {
	Total int              `json:"total"`
	List  []*MedicationRes `json:"list"`
}
