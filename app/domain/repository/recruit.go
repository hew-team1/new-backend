package repository

import (
	"github.com/jmoiron/sqlx"

	"guild-hack-api/app/domain/model"
)

func InsertRecruit(db *sqlx.Tx, recruit *model.Recruit) (int64, error) {
	stmt, err := db.Preparex(
		`INSERT INTO recruits (
			user_id,
			title,
			organizer,
			commit_frequency,
			message,
			join_number,
			is_beginner,
			is_award,
			slack_url,
			start_date,
			end_date
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var id int64
	result, err := stmt.Exec(
		recruit.UserId,
		recruit.Title,
		recruit.Organizer,
		recruit.CommitFrequency,
		recruit.Message,
		recruit.JoinNumber,
		recruit.IsBeginner,
		recruit.IsAward,
		recruit.SlackUrl,
		recruit.StartDate,
		recruit.EndDate,
	)
	// この辺りでDBのエラーを
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FindRecruitById(db *sqlx.DB, id int64) (*model.Recruit, error) {
	var recruit model.Recruit
	err := db.Get(&recruit, "SELECT * FROM recruits WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &recruit, nil
}
