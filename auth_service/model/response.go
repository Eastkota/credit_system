package model

type LoginSuccessData struct {
	Owner *StoreOwner `json:"owner"`
	Store *Store      `json:"store"`
	Token *Token      `json:"token"`
}

type ValidateTokenSuccessData struct {
	Owner *StoreOwner `json:"owner"`
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
	Owner *StoreOwner `json:"owner"`
}

type OwnerResult struct {
	Owner *StoreOwner `json:"owner"`
}

type StoreResult struct {
	Store *Store      `json:"store"`
	Owner *StoreOwner `json:"owner"`
}
