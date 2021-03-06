// Code generated by go-swagger; DO NOT EDIT.

package generated_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecruitList RecruitList
//
// 募集情報のリスト
//
// swagger:model RecruitList
type RecruitList struct {

	// 募集の一覧
	Recruits []*Recruit `json:"recruits"`
}

// Validate validates this recruit list
func (m *RecruitList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecruits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecruitList) validateRecruits(formats strfmt.Registry) error {
	if swag.IsZero(m.Recruits) { // not required
		return nil
	}

	for i := 0; i < len(m.Recruits); i++ {
		if swag.IsZero(m.Recruits[i]) { // not required
			continue
		}

		if m.Recruits[i] != nil {
			if err := m.Recruits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recruits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this recruit list based on the context it is used
func (m *RecruitList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecruits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecruitList) contextValidateRecruits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Recruits); i++ {

		if m.Recruits[i] != nil {
			if err := m.Recruits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recruits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecruitList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecruitList) UnmarshalBinary(b []byte) error {
	var res RecruitList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
