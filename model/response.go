package model

type LoginSuccessData struct {
    User       *User              `json:"user"`
    Profile    *AuthUserProfile   `json:"profile"`
    Token      *Token             `json:"token"`
    Membership *AuthUserMembership `json:"membership"`
}

type ValidateTokenSuccessData struct {
    User  *User  `json:"user"`
}

type GenericAuthSuccessData struct {
    Message string `json:"message"`
    Code    string `json:"code"`
}

type GenericAuthResponse struct {
    Data  interface{} `json:"data,omitempty"`
    Error *AuthError  `json:"error,omitempty"`
}

type DeleteUserResult struct {
    User       *User              `json:"user"`
}

type UserActivityResult struct {
    UserActivity    *UserActivity   `json:"user_activity"`
}