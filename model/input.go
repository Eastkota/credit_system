package model

import "github.com/google/uuid"

type SignupInput struct {
    Email    string `json:"email"`
    MobileNo string `json:"mobile_no"`
    Name     string `json:"name"`
    Gender   string `json:"gender"`
    Password string `json:"password"`
}

type UpdatePasswordInput struct {
    CurrentPassword string `json:"current_password"`
    Password        string `json:"password"`
    ConfirmPassword string `json:"confirm_password"`
    UserId          uuid.UUID `json:"user_id"`
}

type UserActivityInput struct {
    Activity string `json:"activity"`
    UserID        uuid.UUID `json:"user_id"`
}
