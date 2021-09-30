// Code generated by go-swagger; DO NOT EDIT.

package generated_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateRecruitResponse CreateRecruitResponse
//
// 募集作成時のレスポンス
//
// swagger:model CreateRecruitResponse
type CreateRecruitResponse struct {

	// 募集に対して参加できる頻度
	// Example: 3
	// Enum: [1 2 3 4 5 6 7]
	CommitFrequency string `json:"commit_frequency,omitempty"`

	// 募集情報の作成日時
	CreatedAt string `json:"created_at,omitempty"`

	// 募集の参加ハッカソンの開発の終了日
	// Example: 2006-01-02
	// Format: date
	EndDate strfmt.Date `json:"end_date,omitempty"`

	// 募集のid
	// Example: 9999
	ID int64 `json:"id,omitempty"`

	// 募集の参加ハッカソンに賞があるかどうか
	IsAward bool `json:"is_award,omitempty"`

	// 募集が初心者歓迎かどうか
	IsBeginner bool `json:"is_beginner,omitempty"`

	// 募集の参加者数
	JoinNumber int64 `json:"join_number,omitempty"`

	// 募集のメッセージ
	// Example: みんなでわいわい話しながら開発がしたいです！
	Message string `json:"message,omitempty"`

	// 募集の参加ハッカソンの開催者
	// Example: Google
	Organizer string `json:"organizer,omitempty"`

	// 募集の参加者がやりとりをするSlackのURL
	// Format: uri
	SlackURL strfmt.URI `json:"slack_url,omitempty"`

	// 募集の参加ハッカソンの開発の開始日
	// Example: 2006-01-02
	// Format: date
	StartDate strfmt.Date `json:"start_date,omitempty"`

	// 募集のタイトル
	// Example: ハッカソンに参加しよう！
	Title string `json:"title,omitempty"`

	// 募集情報の更新日時
	UpdatedAt string `json:"updated_at,omitempty"`

	// 募集をしたユーザのid
	// Example: 99999
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this create recruit response
func (m *CreateRecruitResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommitFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createRecruitResponseTypeCommitFrequencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","2","3","4","5","6","7"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createRecruitResponseTypeCommitFrequencyPropEnum = append(createRecruitResponseTypeCommitFrequencyPropEnum, v)
	}
}

const (

	// CreateRecruitResponseCommitFrequencyNr1 captures enum value "1"
	CreateRecruitResponseCommitFrequencyNr1 string = "1"

	// CreateRecruitResponseCommitFrequencyNr2 captures enum value "2"
	CreateRecruitResponseCommitFrequencyNr2 string = "2"

	// CreateRecruitResponseCommitFrequencyNr3 captures enum value "3"
	CreateRecruitResponseCommitFrequencyNr3 string = "3"

	// CreateRecruitResponseCommitFrequencyNr4 captures enum value "4"
	CreateRecruitResponseCommitFrequencyNr4 string = "4"

	// CreateRecruitResponseCommitFrequencyNr5 captures enum value "5"
	CreateRecruitResponseCommitFrequencyNr5 string = "5"

	// CreateRecruitResponseCommitFrequencyNr6 captures enum value "6"
	CreateRecruitResponseCommitFrequencyNr6 string = "6"

	// CreateRecruitResponseCommitFrequencyNr7 captures enum value "7"
	CreateRecruitResponseCommitFrequencyNr7 string = "7"
)

// prop value enum
func (m *CreateRecruitResponse) validateCommitFrequencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createRecruitResponseTypeCommitFrequencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateRecruitResponse) validateCommitFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.CommitFrequency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCommitFrequencyEnum("commit_frequency", "body", m.CommitFrequency); err != nil {
		return err
	}

	return nil
}

func (m *CreateRecruitResponse) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("end_date", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateRecruitResponse) validateSlackURL(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackURL) { // not required
		return nil
	}

	if err := validate.FormatOf("slack_url", "body", "uri", m.SlackURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateRecruitResponse) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create recruit response based on context it is used
func (m *CreateRecruitResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateRecruitResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateRecruitResponse) UnmarshalBinary(b []byte) error {
	var res CreateRecruitResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
