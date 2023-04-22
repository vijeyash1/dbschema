package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Roles       []Role    `gorm:"foreignKey:ServiceID"`
	Actions     []Action  `gorm:"foreignKey:ServiceID"`
	Features    []Feature `gorm:"foreignKey:ServiceID"`
}

func (u *Service) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type Role struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Name          string         `gorm:"size:255;not null" json:"name"`
	DisplayName   string         `gorm:"size:255;not null" json:"displayname"`
	Organizations []Organization `gorm:"many2many:organization_user_roles;"`
	Users         []User         `gorm:"many2many:organization_user_roles;"`
	Actions       []Action       `gorm:"many2many:role_actions;"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}

type Organization struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	DisplayName string    `gorm:"size:255;not null" json:"displayname"`
	Description string    `gorm:"size:255" json:"description"`
	OryId       string    `gorm:"size:255" json:"oryid"`
	Users       []User    `gorm:"many2many:organization_user_roles;"`
	Roles       []Role    `gorm:"many2many:organization_user_roles;"`
	Features    []Feature `gorm:"many2many:organization_features;"`
}

func (o *Organization) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}

type Feature struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Name          string         `gorm:"size:255"`
	ServiceID     uuid.UUID      `gorm:"type:binary(16)"`
	Organizations []Organization `gorm:"many2many:organization_features;"`
}

func (a *Feature) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}

type User struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Username      string         `gorm:"size:255;not null" json:"username"`
	Email         string         `gorm:"size:255;not null;unique" json:"email"`
	OryID         string         `json:"ory_id"`
	Organizations []Organization `gorm:"many2many:organization_user_roles;"`
	Roles         []Role         `gorm:"many2many:organization_user_roles;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type Action struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	DisplayName string    `gorm:"size:255;not null" json:"displayname"`
	Roles       []Role    `gorm:"many2many:role_actions;"`
}

func (a *Action) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
