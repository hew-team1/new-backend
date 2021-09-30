package model

import (
	"time"

	"github.com/go-openapi/strfmt"

	"guild-hack-api/swagger/generated_swagger"
)

type Recruit struct {
	Id              int64     `db:"id"`
	UserId          int64     `db:"user_id"`
	Title           string    `db:"title"`
	Organizer       string    `db:"organizer"`
	CommitFrequency string    `db:"commit_frequency"`
	Message         string    `db:"message"`
	JoinNumber      int64     `db:"join_number"`
	IsBeginner      int64     `db:"is_beginner"`
	IsAward         int64     `db:"is_award"`
	SlackUrl        string    `db:"slack_url"`
	StartDate       time.Time `db:"start_date"`
	EndDate         time.Time `db:"end_date"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func (r *Recruit) CreateResponseSwaggerModel() *generated_swagger.CreateRecruitResponse {
	return &generated_swagger.CreateRecruitResponse{
		ID:              r.Id,
		UserID:          r.UserId,
		Title:           r.Title,
		Organizer:       r.Organizer,
		CommitFrequency: r.CommitFrequency,
		Message:         r.Message,
		JoinNumber:      r.JoinNumber,
		IsBeginner: func(i int64) bool {
			if i == 0 {
				return false
			}
			return true
		}(r.IsBeginner),
		IsAward: func(i int64) bool {
			if i == 0 {
				return false
			}
			return true
		}(r.IsAward),
		SlackURL:  strfmt.URI(r.SlackUrl),
		StartDate: strfmt.Date(r.StartDate),
		EndDate:   strfmt.Date(r.EndDate),
		CreatedAt: r.CreatedAt.Format("2006-01-02T15:04"),
		UpdatedAt: r.UpdatedAt.Format("2006-01-02T15:04"),
	}
}
