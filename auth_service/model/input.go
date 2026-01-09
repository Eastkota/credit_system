package model

import (
    "github.com/google/uuid"
)


type SignupInput struct {
    Name       string `json:"name"`
    Password   string `json:"password"`
    Status     string `json:"status"`
    PhoneNumber string `json:"phone_number"`
    StoreName  string `json:"store_name"`
    Address    string `json:"address"`
}

type UpdatePasswordInput struct {
    CurrentPassword string `json:"current_password"`
    Password        string `json:"password"`
    ConfirmPassword string `json:"confirm_password"`
    OwnerId          uuid.UUID `json:"owner_id"`
}


