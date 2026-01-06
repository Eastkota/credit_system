package helpers

import (
	"auth_service/config"
	"auth_service/model"

	"context"
	"fmt"
	"os"
	"time"

	"github.com/machinebox/graphql"
	"github.com/google/uuid"
)


func CreateUserProfile(profileInputData map[string]interface{}) (*model.AuthUserProfile, error) {
	userServiceClient := graphql.NewClient(config.UserServiceApi())
	req := graphql.NewRequest(`
		mutation CreateUserProfile($input:  UserProfileInput){
			createUserProfile(input: $input) {
				data {
					user_profile {
						gender
						id
						name
						profile_picture
					}
				}
				error {
					code
					field
					message
				}
			}
		}
	`)
	req.Var("input", profileInputData)
	req.Header.Set("Cache-Control", "no-cache")

	// If a public access token is configured, use it to authenticate internal service-to-service calls
	if t := os.Getenv("PUBLIC_ACCESS_TOKEN"); t != "" {
		if len(t) > 6 && (t[:7] == "Bearer " || t[:7] == "bearer ") {
			req.Header.Set("Authorization", t)
		} else {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t))
		}
	}

	var response struct {
		CreateUserProfile struct {
			Data struct {
				UserProfile model.AuthUserProfile `json:"user_profile"`
			} `json:"data"`
			Error struct {
				Code    string `json:"code"`
				Field   string `json:"field"`
				Message string `json:"message"`
			} `json:"error"`
		} `json:"createUserProfile"`
	}

	err := userServiceClient.Run(context.Background(), req, &response)

	return &response.CreateUserProfile.Data.UserProfile, err
}

func GetUserProfile(userID uuid.UUID) (*model.AuthUserProfile, error) {

	// Prepare GraphQL request
	userServiceClient := graphql.NewClient(config.UserServiceApi())
	req := graphql.NewRequest(`
		query FetchProfileByUserId($user_id: UUID) {
			fetchProfileByUserId(user_id: $user_id) {
				data {
					user_profile {
						created_at
						gender
						id
						name
						profile_picture
						updated_at
					}
				}
				error {
					code
					field
					message
				}
			}
		}
	`)

	// Set the variable
	req.Var("user_id", userID)
	req.Header.Set("Cache-Control", "no-cache")

	// If a public access token is configured, use it to authenticate internal service-to-service calls
	if t := os.Getenv("PUBLIC_ACCESS_TOKEN"); t != "" {
		if len(t) > 6 && (t[:7] == "Bearer " || t[:7] == "bearer ") {
			req.Header.Set("Authorization", t)
		} else {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t))
		}
	}

	// Define response struct
	var tempResponse struct {
        FetchProfileByUserId struct {
            Data struct {
                UserProfile struct {
                    CreatedAt               time.Time `json:"created_at"`
                    Gender                  string `json:"gender"`
                    ID                      uuid.UUID `json:"id"`
                    Name                    string `json:"name"`
                    ProfilePicture          string `json:"profile_picture"`
                    UpdatedAt               time.Time `json:"updated_at"`
                } `json:"user_profile"`
            } `json:"data"`
            Error struct {
                Code    string `json:"code"`
                Field   string `json:"field"`
                Message string `json:"message"`
            } `json:"error"`
        } `json:"fetchProfileByUserId"`
    }


	err := userServiceClient.Run(context.Background(), req, &tempResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user profile : %v", err)
	}
	// Check for errors in response
	if tempResponse.FetchProfileByUserId.Error.Message != "" {
		return nil, fmt.Errorf("error from user service: %s", tempResponse.FetchProfileByUserId.Error.Message)
	}

	formattedProfile := &model.AuthUserProfile{
        ID:                      tempResponse.FetchProfileByUserId.Data.UserProfile.ID,
        Name:                    tempResponse.FetchProfileByUserId.Data.UserProfile.Name,
        ProfilePicture:          tempResponse.FetchProfileByUserId.Data.UserProfile.ProfilePicture,
        Gender:                  tempResponse.FetchProfileByUserId.Data.UserProfile.Gender,
        CreatedAt:               tempResponse.FetchProfileByUserId.Data.UserProfile.CreatedAt,
        UpdatedAt:               tempResponse.FetchProfileByUserId.Data.UserProfile.UpdatedAt,
    }


    return formattedProfile, nil
}
