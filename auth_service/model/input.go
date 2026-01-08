package model


type SignupInput struct {
    Name       string `json:"name"`
    Password   string `json:"password"`
    Status     string `json:"status"`
    PhoneNumber string `json:"phone_number"`
}


