package resolvers

import (
	"credit_system/auth_service/helpers"
	"credit_system/auth_service/model"
	"credit_system/auth_service/services"

	"encoding/json"

	"github.com/graphql-go/graphql"
	"github.com/google/uuid"
)

type AuthResolver struct {
	Services services.Services // Inject Services
}

func NewResolver(service services.Services) *AuthResolver {
	return &AuthResolver{Services: service}
}

func (ar *AuthResolver) CheckForExistingUser(p graphql.ResolveParams) *model.GenericAuthResponse {
	
	field := p.Args["field"].(string)
	value := p.Args["value"].(string)
	result, err := ar.Services.CheckForExistingUser(field, value)
	if err != nil {
		return helpers.FormatError(err)
	}

	if result == nil {
        return &model.GenericAuthResponse{
            Data: map[string]interface{}{
                "exist_user": false,
                "owner_id":    nil,
            },
            Error: nil,
        }
    }
	
	return &model.GenericAuthResponse{
		Data: map[string]interface{}{
			"exist_user": result != nil,
			"owner_id":    result.ID,
		},
		Error: nil,
	}

}

func (ar *AuthResolver) Signup(p graphql.ResolveParams) *model.GenericAuthResponse {

	var signupInput model.SignupInput
	inputData := p.Args["signup_input"].(map[string]interface{})

	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return helpers.FormatError(err)
	}
	err = json.Unmarshal(jsonData, &signupInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	owner, token, store, err := ar.Services.Signup(signupInput)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericAuthResponse{
		Data: &model.LoginSuccessData{
			Owner:   owner,
			Token:  token,
			Store:  store,
		},
		Error: nil,
	}
}

func (ar *AuthResolver) FetchStore(p graphql.ResolveParams) *model.GenericAuthResponse {
	storeID := p.Args["store_id"].(uuid.UUID)	
	result, err := ar.Services.FetchStore(storeID)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericAuthResponse{
		Data: &model.StoreResult{
			Store: result,
		},
		Error: nil,
	}
}

func (ar *AuthResolver) FetchStoreByOwnerID(p graphql.ResolveParams) *model.GenericAuthResponse {
	ownerID := p.Args["owner_id"].(uuid.UUID)	
	result, err := ar.Services.FetchStoreByOwnerID(ownerID)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericAuthResponse{
		Data: &model.StoreResult{
			Store: result,
		},
		Error: nil,
	}
}


func (ar *AuthResolver) Login(p graphql.ResolveParams) *model.GenericAuthResponse {
	phoneNumber := p.Args["phone_number"].(string)
	password := p.Args["password"].(string)
	owner, token, store, err := ar.Services.Login(phoneNumber, password)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericAuthResponse{
		Data: &model.LoginSuccessData{
			Owner:       owner,
			Store:    store,
			Token:      token,
		},
		Error: nil,
	}
}

func (ar *AuthResolver) ValidateToken(p graphql.ResolveParams) *model.GenericAuthResponse {
	tokenString := p.Args["token"].(string)
	owner, err := ar.Services.ValidateToken(tokenString)
	if err != nil {
		return helpers.FormatError(err)
	}
	
	return &model.GenericAuthResponse{
		Data: &model.ValidateTokenSuccessData{
			Owner:  owner,
		},
		Error: nil,
	}
}

func (ar *AuthResolver) RefreshToken(p graphql.ResolveParams) *model.GenericAuthResponse {
	tokenString := p.Args["token"].(string)
	owner, token, store, err := ar.Services.RefreshToken(tokenString)
	if err != nil {
		return helpers.FormatError(err)
	}
	
	return &model.GenericAuthResponse{
		Data: &model.LoginSuccessData{
			Owner:       owner,
			Store:    store,
			Token:      token,
		},
		Error: nil,
	}
}	

// func (ar *AuthResolver) Logout(p graphql.ResolveParams) *model.GenericAuthResponse {
// 	req, ok := p.Context.Value(model.RequestKey).(*http.Request)
// 	if !ok {
// 		return &model.GenericAuthResponse{
// 			Data: nil,
// 			Error: &model.AuthError{
// 				Message: "invalid_token",
// 			},
// 		}
// 	}

// 	tokenString := req.Header.Get("Authorization")
// 	tokenId,ok := ar.Services.IsValidToken(tokenString)
// 	if !ok {
// 		return &model.GenericAuthResponse{
// 			Data: nil,
// 			Error: &model.AuthError{
// 				Message: "invalid_token",
// 			},
// 		}
// 	}
// 	err := ar.Services.Logout(tokenId)

// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}

// 	return &model.GenericAuthResponse{
// 		Data: &model.GenericAuthSuccessData{
// 			Message: "Successfully logged out",
// 		},
// 		Error: nil,
// 	}
// }

// func (ar *AuthResolver) UpdatePassword(p graphql.ResolveParams) *model.GenericAuthResponse {
// 	var updatePasswordInput model.UpdatePasswordInput
// 	inputData := p.Args["input"].(map[string]interface{})

// 	jsonData, err := json.Marshal(inputData)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	err = json.Unmarshal(jsonData, &updatePasswordInput)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	err = ar.Services.UpdatePassword(updatePasswordInput)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	return &model.GenericAuthResponse{
// 		Data: &model.GenericAuthSuccessData{
// 			Message: "Password updated successfully",
// 		},
// 		Error: nil,
// 	}
// }
// func (ar *AuthResolver) UpdateSingleDataByID(p graphql.ResolveParams) *model.GenericAuthResponse {
// 	ctx := p.Context
//     user := ctx.Value(model.UserKey).(*model.User)
// 	if user == nil {
// 		return helpers.FormatError(fmt.Errorf("invalid_token"))
// 	}
		
// 	field := p.Args["field"].(string)
// 	value := p.Args["value"].(string)
// 	password := p.Args["password"].(string)

// 	userID := user.ID 
// 	//Implement  the login to fetch user form Authorization
// 	result, err := ar.Services.UpdateSingleDataByID(userID, field, value, password)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	return &model.GenericAuthResponse{
// 		Data: &model.UserResult{
// 			User: result,
// 		},
// 		Error: nil,
// 	}
// }

func (ar *AuthResolver) FetchOwnerByID(p graphql.ResolveParams) *model.GenericAuthResponse {
	ownerID := p.Args["owner_id"].(uuid.UUID)
	result, err := ar.Services.FetchOwnerByID(ownerID)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericAuthResponse{
		Data: &model.OwnerResult{
			Owner: result,
		},
		Error: nil,
	}
}

// func (ar *AuthResolver) ResetPassword(p graphql.ResolveParams) *model.GenericAuthResponse {
// 	userID := p.Args["user_id"].(uuid.UUID)
// 	password := p.Args["password"].(string)
// 	confirmPassword := p.Args["confirm_password"].(string)
// 	err := ar.Services.ResetPassword(userID, password, confirmPassword)
// 	if err != nil {
// 		return helpers.FormatError(err)
// 	}
// 	return &model.GenericAuthResponse{
// 		Data: &model.GenericAuthSuccessData{
// 			Message: "Password updated successfully",
// 		},
// 		Error: nil,
// 	}
// }

// func (ar *AuthResolver) DeleteUser(p graphql.ResolveParams) *model.GenericAuthResponse {
//     userID, ok := p.Args["userID"].(uuid.UUID)
//     if !ok || userID == uuid.Nil {
//         return helpers.FormatError(fmt.Errorf("userID is required"))
//     }

//     password, ok := p.Args["password"].(string)
//     if !ok || password == "" {
//         return helpers.FormatError(fmt.Errorf("Password is required and must be a string"))
//     }

//     // The 'status' variable from the args should be a string, not a boolean as your code previously suggested.
//     status, ok := p.Args["status"].(string)
//     if !ok {
//         return helpers.FormatError(fmt.Errorf("User status is required and must be a string"))
//     }

//     // Pass the password to the service function
//     result, err := ar.Services.DeleteUser(p.Context, userID, status, password)
//     if err != nil {
//         return helpers.FormatError(err)
//     }

//     return &model.GenericAuthResponse{
//         Data: &model.DeleteUserResult{
//             User: result,
//         },
//         Error: nil,
//     }
// }

// func (ar *AuthResolver) SaveUserActivity(p graphql.ResolveParams) (interface{}, error) {
//     input, ok := p.Args["input"].(map[string]interface{})
//     if !ok {
//         return nil, fmt.Errorf("invalid input format")
//     }

//     activity, _ := input["activity"].(string)
//     user_id, _ := input["user_id"].(uuid.UUID)

//     newActivity, err := ar.Services.SaveUserActivity(p.Context, &model.UserActivityInput{
//         Activity: activity,
//         UserID:   user_id,
//     })
//     if err != nil {
//         return nil, fmt.Errorf("failed to save user activity: %v", err)
//     }

//     return &model.GenericAuthResponse{
//         Data: &model.UserActivityResult{
//             UserActivity: newActivity, 
//         },
//     }, nil
// }

// func (ar *AuthResolver) SaveGameActivity(p graphql.ResolveParams) (interface{}, error) {
//     userID, _ := p.Args["user_id"].(uuid.UUID)

//     err := ar.Services.SaveGameActivity(p.Context, userID)
    
//     if err != nil {
//         return nil, fmt.Errorf("failed to save game activity: %v", err)
//     }

//     return &model.GenericAuthResponse{
//         Data: map[string]interface{}{
//             "message": "success",
//             "user_activity": nil,
//         },
//     }, nil
// }
