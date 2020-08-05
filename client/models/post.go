// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Post Post represents a contribution that a user has made to the blog.
//
// A Post can be created, edited or deleted by its author.
//
// swagger:model Post
type Post struct {

	// ID of the author user
	// Required: true
	AuthorID *uint32 `json:"author_id"`

	// Content holds the main information that the use wants to share
	// Required: true
	Content *string `json:"content"`

	// createdAt is the time that the post was created in the db
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// ID that belongs to the post, assigned by the db
	// Minimum: 1
	ID uint64 `json:"id,omitempty"`

	// The post's title
	// Required: true
	Title *string `json:"title"`

	// updatedAt is the time that the post is updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this post
func (m *Post) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Post) validateAuthorID(formats strfmt.Registry) error {

	if err := validate.Required("author_id", "body", m.AuthorID); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", int64(m.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Post) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Post) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Post) UnmarshalBinary(b []byte) error {
	var res Post
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
