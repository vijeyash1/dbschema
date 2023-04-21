package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
User and Organization have a many-to-many relationship through the OrganizationUserRole model.

Role and Action have a many-to-many relationship through the RoleActionService model.

Organization and Features have a many-to-many relationship through the OrganizationFeatureService model.

Service has a one-to-many relationship with Role, Action, and Features.
*/

type User struct {
	ID       string
	Username string
	Email    string
	OryID    string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

type Organization struct {
	ID          string
	Name        string
	Description string
	OryId       string
}

func (o *Organization) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}

type Service struct {
	ID          string
	Name        string
	Description string
}

func (o *Service) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}

type Role struct {
	ID        string
	Name      string
	ServiceID string
}

func (o *Role) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}

type Action struct {
	ID        string
	Name      string
	ServiceID string
}

func (a *Action) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return
}

type Features struct {
	ID        string
	Name      string
	ServiceID string
}

func (a *Features) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	return
}

type OrganizationUserRole struct {
	ID             string
	OrganizationID string
	UserID         string
	RoleID         string
}

func (o *OrganizationUserRole) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}

type RoleActionService struct {
	ID        string
	RoleID    string
	ActionID  string
	ServiceID string
}

func (r *RoleActionService) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	return
}

type OrganizationFeatureService struct {
	ID             string
	OrganizationID string
	FeatureID      string
	ServiceID      string
}

func (o *OrganizationFeatureService) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New().String()
	return
}
