package helpers

import (
	"auth_service/config"
	"auth_service/model"

	"context"
	"fmt"
	"strings"
	"time"
	"os"

	"github.com/google/uuid"
	"github.com/machinebox/graphql"
)

func GetMembership(userID uuid.UUID) (*model.AuthUserMembership, error) {
	// Prepare GraphQL request
	membershipServiceClient := graphql.NewClient(config.MembershipApi())
	req := graphql.NewRequest(`
		query FetchMembershipByUserId($user_id: UUID) {
			fetchMembershipByUserId(user_id: $user_id) {
				data {
					user_membership {
						id
						membership_duration_id
						membership_end_date
						membership_join_date
						user_id
						membership_duration {
							description
							duration_type
							id
							number_of_days
							package_id
							price
							price_usd
							package {
								created_at
								description
								id
								name
								package_type
								updated_at
								user_limit
							}
						}
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
		FetchMembershipByUserId struct {
			Data struct {
				UserMembership struct {
					ID                   uuid.UUID `json:"id"`
					MembershipDurationID uuid.UUID `json:"membership_duration_id"`
					MembershipJoinDate   time.Time `json:"membership_join_date"`
					MembershipEndDate    time.Time `json:"membership_end_date"`
					UserId               uuid.UUID `json:"user_id"`
					MembershipDuration *struct {
						ID           uuid.UUID `json:"id"`
						DurationType string    `json:"duration_type"`
						PackageID    uuid.UUID `json:"package_id"`
						NumberOfDays string       `json:"number_of_days"`
						Description  string    `json:"description"`
						Package *struct {
							ID          uuid.UUID `json:"id"`
							Name        string    `json:"name"`
							Description string    `json:"description"`
							PackageType string    `json:"package_type"`
							UserLimit   int       `json:"user_limit"`
						} `json:"package"`
					} `json:"membership_duration"`
				} `json:"user_membership"`
			} `json:"data"`
			Error struct {
				Code    string `json:"code"`
				Field   string `json:"field"`
				Message string `json:"message"`
			} `json:"error"`
		} `json:"fetchMembershipByUserId"`
	}

	// Execute the request
	err := membershipServiceClient.Run(context.Background(), req, &tempResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user profile : %v", err)
	}
	if tempResponse.FetchMembershipByUserId.Error.Message != "" {
		if strings.Contains(tempResponse.FetchMembershipByUserId.Error.Message, "no documents in result") {
			return nil, nil
		}
		return nil, fmt.Errorf("error from user service: %s", tempResponse.FetchMembershipByUserId.Error.Message)
	}

	var membershipDuration *model.MembershipDuration = nil

    if md := tempResponse.FetchMembershipByUserId.Data.UserMembership.MembershipDuration; md != nil {
        
        var pkg *model.Package = nil
        
        if p := md.Package; p != nil {
            pkg = &model.Package{
                ID:          p.ID,
                Name:        p.Name,
                Description: p.Description,
                PackageType: p.PackageType,
				UserLimit:   p.UserLimit,
            }
        }
        
        membershipDuration = &model.MembershipDuration{
            ID:           md.ID,
            DurationType: md.DurationType,
            PackageId:    md.PackageID,
            NumberOfDays: md.NumberOfDays,
            Description:  md.Description, 
            Package:      pkg,            
        }
    }

    formattedMembership := &model.AuthUserMembership{
        ID:                   tempResponse.FetchMembershipByUserId.Data.UserMembership.ID,
        MembershipJoinDate:   tempResponse.FetchMembershipByUserId.Data.UserMembership.MembershipJoinDate,
        MembershipEndDate:    tempResponse.FetchMembershipByUserId.Data.UserMembership.MembershipEndDate,
        UserId:               tempResponse.FetchMembershipByUserId.Data.UserMembership.UserId,
        MembershipDurationId: tempResponse.FetchMembershipByUserId.Data.UserMembership.MembershipDurationID,
        MembershipDuration:   membershipDuration, 
    }

    return formattedMembership, nil
}
