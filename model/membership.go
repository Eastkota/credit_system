package model

import (
    "time"

    "github.com/google/uuid"
)

type AuthUserMembership struct {
    ID                   uuid.UUID `json:"id"`
    MembershipJoinDate   time.Time `json:"membership_join_date"`
    MembershipEndDate    time.Time `json:"membership_end_date"`
    UserId               uuid.UUID `json:"user_id"`
    MembershipDurationId uuid.UUID `json:"membership_duration_id"`
    
    MembershipDuration   *MembershipDuration `gorm:"foreignKey:MembershipDurationID" json:"membership_duration"`
}

type Package struct {
	ID          uuid.UUID            `gorm:"type: uuid; not null" json:"id"`
	Name        string               `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string               `gorm:"type: text" json:"description"`
	PackageType string               `gorm:"type:string" json:"package_type"`
	Durations   []MembershipDuration `gorm:"foreignKey:PackageId;references:ID" json:"durations"`
	UserLimit   int                  `gorm:"type:int; not null" json:"user_limit"`
	CreatedAt   time.Time            `gorm:"type:timestamptz; not null" json:"created_at"`
	UpdatedAt   time.Time            `gorm:"type:timestamptz; not null" json:"updated_at"`
}

type MembershipDuration struct {
	ID           uuid.UUID `gorm:"type:uuid" json:"id,omitempty"`
	DurationType string    `gorm:"type:varchar(255);unique;not null" json:"duration_type"`
	Price        float32   `gorm:"type:float; not null" json:"price"`
	Description  string    `gorm:"type:string" json:"description"`
	UpdatedAt    time.Time `gorm:"type:timestamptz; not null" json:"updated_at"`
	CreatedAt    time.Time `gorm:"type:timestamptz; not null" json:"created_at"`
	PackageId    uuid.UUID `gorm:"type: uuid; not null" json:"package_id"`
	NumberOfDays string       `gorm:"type: varchar; not null" json:"number_of_days"`
	PriceUSD     float32   `gorm:"type: float; not null" json:"price_usd"`
	Package      *Package  `gorm:"foreignKey:PackageId;references:ID" json:"package"`
}