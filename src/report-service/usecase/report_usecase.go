package usecase

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/Hospital-Microservice/report-service/provider"
)

// ChartData response shape for bar chart
type ChartData struct {
	Labels []string `json:"labels"`
	Values []int    `json:"values"`
}

// PatientsReportUseCase
type PatientsReportUseCase struct {
	UserClient provider.UserService
}

func NewPatientsReportUseCase(u provider.UserService) *PatientsReportUseCase {
	return &PatientsReportUseCase{UserClient: u}
}

// month: 1..12
func (r *PatientsReportUseCase) Execute(ctx context.Context, month int, year int, auth string) (ChartData, error) {
	loc := time.UTC
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	end := start.AddDate(0, 1, 0)

	// request params to get users created in range (services may ignore unused params)
	params := url.Values{}
	params.Set("from", start.Format(time.RFC3339))
	params.Set("to", end.Format(time.RFC3339))
	params.Set("limit", "10000")

	items, err := r.UserClient.GetUsers(ctx, params, auth)
	if err != nil {
		return ChartData{}, err
	}

	// number of days in this month
	daysInMonth := end.AddDate(0, 0, -1).Day()
	labels := make([]string, 0, daysInMonth)
	values := make([]int, daysInMonth)

	for d := 1; d <= daysInMonth; d++ {
		labels = append(labels, strconv.Itoa(d))
	}

	for _, it := range items {
		// try to get created_at
		if v, ok := it["created_at"]; ok {
			if s, ok := v.(string); ok {
				if t, err := time.Parse(time.RFC3339, s); err == nil {
					if !t.Before(start) && t.Before(end) {
						day := t.Day()
						values[day-1]++
					}
				}
			}
		}
	}

	return ChartData{Labels: labels, Values: values}, nil
}

// AppointmentsReportUseCase
type AppointmentsReportUseCase struct {
	ApptClient provider.AppointmentService
}

func NewAppointmentsReportUseCase(a provider.AppointmentService) *AppointmentsReportUseCase {
	return &AppointmentsReportUseCase{ApptClient: a}
}

// doctorID optional (empty => all doctors)
func (r *AppointmentsReportUseCase) Execute(ctx context.Context, doctorID string, month int, year int, auth string) (ChartData, error) {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	params := url.Values{}
	params.Set("from", start.Format(time.RFC3339))
	params.Set("to", end.Format(time.RFC3339))
	if doctorID != "" {
		params.Set("doctor_id", doctorID)
	}
	items, err := r.ApptClient.GetAppointments(ctx, params, auth)
	if err != nil {
		return ChartData{}, err
	}

	daysInMonth := end.AddDate(0, 0, -1).Day()
	labels := make([]string, 0, daysInMonth)
	values := make([]int, daysInMonth)

	for d := 1; d <= daysInMonth; d++ {
		labels = append(labels, strconv.Itoa(d))
	}

	for _, it := range items {
		if v, ok := it["scheduled_at"]; ok {
			if s, ok := v.(string); ok {
				if t, err := time.Parse(time.RFC3339, s); err == nil {
					if !t.Before(start) && t.Before(end) {
						day := t.Day()
						values[day-1]++
					}
				}
			}
		}
	}
	return ChartData{Labels: labels, Values: values}, nil
}

// PrescriptionsReportUseCase
type PrescriptionsReportUseCase struct {
	PresClient provider.PrescriptionService
}

func NewPrescriptionsReportUseCase(p provider.PrescriptionService) *PrescriptionsReportUseCase {
	return &PrescriptionsReportUseCase{PresClient: p}
}

func (r *PrescriptionsReportUseCase) Execute(ctx context.Context, month int, year int, auth string) (ChartData, error) {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	params := url.Values{}
	params.Set("from", start.Format(time.RFC3339))
	params.Set("to", end.Format(time.RFC3339))
	items, err := r.PresClient.GetPrescriptions(ctx, params, auth)
	if err != nil {
		return ChartData{}, err
	}

	daysInMonth := end.AddDate(0, 0, -1).Day()
	labels := make([]string, 0, daysInMonth)
	values := make([]int, daysInMonth)

	for d := 1; d <= daysInMonth; d++ {
		labels = append(labels, strconv.Itoa(d))
	}

	for _, it := range items {
		if v, ok := it["created_at"]; ok {
			if s, ok := v.(string); ok {
				if t, err := time.Parse(time.RFC3339, s); err == nil {
					if !t.Before(start) && t.Before(end) {
						day := t.Day()
						values[day-1]++
					}
				}
			}
		}
	}
	return ChartData{Labels: labels, Values: values}, nil
}
