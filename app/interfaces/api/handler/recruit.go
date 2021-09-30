package handler

import (
	"net/http"
	"time"

	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"

	"guild-hack-api/app/interfaces/api/httputil"
	"guild-hack-api/app/usecase"
	"guild-hack-api/swagger/generated_swagger"
)

type RecruitHandler struct {
	recruitUseCase *usecase.RecruitUseCase
}

func NewRecruitHandler(recruitUseCase *usecase.RecruitUseCase) *RecruitHandler {
	return &RecruitHandler{
		recruitUseCase: recruitUseCase,
	}
}

func (h *RecruitHandler) Create(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var req generated_swagger.CreateRecruitRequest
	if err := parseModelRequest(r, &req); err != nil {
		return http.StatusBadRequest, nil, err
	}

	recruit, err := h.recruitUseCase.CreateAndFindRecruit(
		swag.StringValue(req.Title),
		user.Id,
		swag.StringValue(req.Organizer),
		swag.StringValue(req.CommitFrequency),
		swag.StringValue(req.Message),
		swag.Int64Value(req.JoinNumber),
		func(b bool) int64 {
			if b {
				return 1
			}
			return 0
		}(swag.BoolValue(req.IsBeginner)),
		func(b bool) int64 {
			if b {
				return 1
			}
			return 0
		}(swag.BoolValue(req.IsAward)),
		req.SlackURL.String(),
		time.Time(conv.DateValue(req.StartDate)),
		time.Time(conv.DateValue(req.EndDate)),
	)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusCreated, recruit.CreateResponseSwaggerModel(), nil
}
