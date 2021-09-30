package usecase

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"guild-hack-api/app/db/dbutil"
	"guild-hack-api/app/domain/model"
	"guild-hack-api/app/domain/repository"
)

type RecruitUseCase struct {
	db *sqlx.DB
}

func NewRecruitUseCase(db *sqlx.DB) *RecruitUseCase {
	return &RecruitUseCase{
		db: db,
	}
}

func (r *RecruitUseCase) CreateAndFindRecruit(title string, userId int64, organizer string, commitFrequency string, message string, joinNumber int64, isBeginner int64, isAward int64, slackUrl string, startDate time.Time, endDate time.Time) (*model.Recruit, error) {
	newRecruit := &model.Recruit{
		UserId:          userId,
		Title:           title,
		Organizer:       organizer,
		CommitFrequency: commitFrequency,
		Message:         message,
		JoinNumber:      joinNumber,
		IsBeginner:      isBeginner,
		IsAward:         isAward,
		SlackUrl:        slackUrl,
		StartDate:       startDate,
		EndDate:         endDate,
	}

	var createId int64
	if err := dbutil.TXHandler(r.db, func(tx *sqlx.Tx) error {
		id, err := repository.InsertRecruit(tx, newRecruit)
		if err != nil {
			return err
		}
		createId = id
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return nil, fmt.Errorf("failed recruit insert transaction: %w", err)
	}

	return repository.FindRecruitById(r.db, createId)
}
