package repositories

import (
    "auth_service/config"
    "auth_service/helpers"
    "auth_service/model"

    "errors"
    "fmt"

    "gorm.io/gorm"
	"github.com/google/uuid"	
)

func (repo *AuthRepository) CheckForExistingUser(field, value string) (*model.StoreOwner, error) {
    var store_owner model.StoreOwner
    err := repo.DB.Where(fmt.Sprintf("%s = ? AND status != ?", field), value, "Deleted").First(&store_owner).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("failed to find user with %v %v", field, value)
    }
    return &store_owner, nil
}

func (repo *AuthRepository) FetchUserByID(field, value string) (*model.StoreOwner, error) {
    var store_owner model.StoreOwner
    err := repo.DB.Where(fmt.Sprintf("%s = ?", field), value).First(&store_owner).Error
    if err != nil {
        return nil, fmt.Errorf("failed to find user with %v %v", field, value)
    }
    return &store_owner, nil
}

func (repo *AuthRepository) RegisterUser(signupInput *model.SignupInput) (*model.StoreOwner, *model.Token, *model.Store, error) {
    var storeOwner model.StoreOwner
    var tokenResult *model.Token
    var storeResult model.Store

    err := repo.DB.Where("phone_number = ?", signupInput.PhoneNumber).First(&storeOwner).Error
    if err == nil && storeOwner.Status == "Active" {
        return nil, nil, nil, fmt.Errorf("user already exists with provided credentials")
    }

    hashedPassword, err := helpers.EncryptPassword(signupInput.Password)
    if err != nil {
        return nil, nil, nil, fmt.Errorf("failed to hash password: %v", err)
    }

    err = repo.DB.Transaction(func(tx *gorm.DB) error {
        
        storeOwner = model.StoreOwner{
            ID:          uuid.New(),
            Name:        signupInput.Name,
            PhoneNumber: signupInput.PhoneNumber,
            Password:    hashedPassword,
            Status:      "Active",
        }

        if err := tx.Create(&storeOwner).Error; err != nil {
            return fmt.Errorf("failed to create store owner: %v", err)
        }

        storeResult = model.Store{
            ID:          uuid.New(),
            Name:        signupInput.Name,
            OwnerID:     storeOwner.ID,
        }

        if err := tx.Create(&storeResult).Error; err != nil {
            return fmt.Errorf("failed to create store: %v", err)
        }

        tokenResult, err = repo.CreateToken(&storeOwner, config.ClientId(), config.ClientSecret(), tx)
        if err != nil {
            return fmt.Errorf("token generation failed: %v", err)
        }

        return nil
    })

    if err != nil {
        return nil, nil, nil, err
    }

    return &storeOwner, tokenResult, &storeResult, nil
}

func (repo *AuthRepository) FetchStore(storeID uuid.UUID) (*model.Store, error) {
    var store model.Store
    if err := repo.DB.First(&store, "id = ?", storeID).Error; err != nil {
        return nil, fmt.Errorf("store not found: %v", err)
    }
    return &store, nil
}

// func (repo *AuthRepository) Login(loginId, password string) (*model.User, *model.Token, error) {
//     var user *model.User
//     var err error
    
//     emailRegx := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
//     if emailRegx.MatchString(loginId) {
//         user, err = repo.FetchUserByLoginID("email", loginId)
//     } else {
//         user, err = repo.FetchUserByLoginID("mobile_no", loginId)
//     }
//     if err != nil {
//         return nil, nil, err
//     }

//     // if user.Status != "Active" {
//     //     updateErr := repo.DB.Model(&user).Update("status", "Active").Error
//     //     if updateErr != nil {
//     //         return nil, nil, fmt.Errorf("failed to update user status: %w", updateErr)
//     //     }
//     //     user.Status = "Active"
//     // }

//     if user.Status == "Deleted" {
//         return nil, nil, fmt.Errorf("user account is deleted")
//     }

//     if helpers.IsValidPassword(password, user.Password) {
//         tokenResult, err := repo.CreateToken(user, config.ClientId(), config.ClientSecret(), nil)
//         if err != nil {
//             return nil, nil, err
//         }
//         return user, tokenResult, nil
//     }

//     return nil, nil, fmt.Errorf("invalid credentials")
// }

// func (repo *AuthRepository) UpdateSingleDataByID(userID uuid.UUID, field, value string) (*model.User, error) {
//     updateData := map[string]interface{}{
//         field:   value,
//         "updated_at": time.Now(),
//     }
    
//     result := repo.DB.Model(&model.User{}).Where("id = ?", userID).Updates(updateData)

//     if result.Error != nil {
//         return nil, fmt.Errorf("failed to update user: %v", result.Error)
//     }

//     if result.RowsAffected == 0 {
//         return nil, fmt.Errorf("user not found with ID: %v", userID)
//     }

//     var updatedUser model.User
//     if err := repo.DB.First(&updatedUser, "id = ?", userID).Error; err != nil {
//         return nil, fmt.Errorf("failed to retrieve updated user: %v", err)
//     }

//     return &updatedUser, nil
// }

func (repo *AuthRepository) FetchUser(ownerID uuid.UUID) (*model.StoreOwner, error) {
    var owner model.StoreOwner
    if err := repo.DB.First(&owner, "id = ?", ownerID).Error; err != nil {
        return nil, fmt.Errorf("owner not found: %v", err)
    }
    return &owner, nil
}

// func (repo *AuthRepository) DeleteUser(ctx context.Context, userID uuid.UUID, status, password string) (*model.User, error) {
//     var user model.User
//     if err := repo.DB.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
//         return nil, fmt.Errorf("user not found with ID: %s", userID)
//     }

//     // Use the existing IsValidPassword helper
//     if !helpers.IsValidPassword(password, user.Password) {
//         return nil, fmt.Errorf("invalid password")
//     }

//     updates := map[string]interface{}{
//         "status": status,
//         "updated_at":   time.Now(),
//     }

//     if err := repo.DB.WithContext(ctx).Model(&user).Updates(updates).Error; err != nil {
//         return nil, fmt.Errorf("failed to update user status: %w", err)
//     }

//     var updatedUser model.User
//     if err := repo.DB.WithContext(ctx).Where("id = ?", userID).First(&updatedUser).Error; err != nil {
//         return nil, fmt.Errorf("failed to fetch the updated USER: %w", err)
//     }

//     return &updatedUser, nil
// }

// func (repo *AuthRepository) CreateUserActivity(ctx context.Context, inputActivityType string, userID uuid.UUID) (*model.UserActivity, error) {
//     now := time.Now()
//     currentMonth := now.Month()
//     currentYear := now.Year()

//     var canonicalActivityName string
//     if strings.Contains(inputActivityType, "video") || strings.Contains(inputActivityType, "Video") {
//         canonicalActivityName = "video_watched"
//     } else {
//         canonicalActivityName = "others"       
//     }

//     var userActivity model.UserActivity
    
//     result := repo.DB.WithContext(ctx).
//         Where("user_id = ?", userID).
//         Where("activity = ?", canonicalActivityName).
//         Where("month = ?", currentMonth).
//         Where("year = ?", currentYear).
//         First(&userActivity)
    
//     if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
//         return nil, result.Error
//     }

//     if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//         userActivity = model.UserActivity{
//             ID:     uuid.New(),
//             UserID: userID,
//             Month:  int(currentMonth),
//             Year:   currentYear,
//             Activity: canonicalActivityName,
//             Count:  1,
//         }

//         if err := repo.DB.WithContext(ctx).Create(&userActivity).Error; err != nil {
//             return nil, err
//         }

//     } else {
//         if err := repo.DB.WithContext(ctx).
//             Model(&userActivity).
//             Where("id = ?", userActivity.ID).
//             Update("count", gorm.Expr("count + ?", 1)). 
//             Error; err != nil {
//             return nil, err
//         }
        
//         if err := repo.DB.WithContext(ctx).
//             Preload("User").
//             First(&userActivity, "id = ?", userActivity.ID).
//             Error; err != nil {
//             return nil, err
//         }
//     }

//     if userActivity.User == nil {
//         repo.DB.WithContext(ctx).Preload("User").First(&userActivity, "id = ?", userActivity.ID)
//     }

//     return &userActivity, nil
// }

// func (repo *AuthRepository) CreateGameActivity(ctx context.Context, userID uuid.UUID) error {
//     now := time.Now()
//     currentMonth := now.Month()
//     currentYear := now.Year()
//     var canonicalActivityName = "CheyCheyActivity"

//     var userActivity model.UserActivity
    
//     result := repo.DB.WithContext(ctx).
//         Where("user_id = ? AND activity = ? AND month = ? AND year = ?", 
//             userID, canonicalActivityName, int(currentMonth), currentYear).
//         First(&userActivity)
    
//     if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
//         return result.Error
//     }

//     if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//         newActivity := model.UserActivity{
//             ID:       uuid.New(),
//             UserID:   userID,
//             Month:    int(currentMonth),
//             Year:     currentYear,
//             Activity: canonicalActivityName,
//             Count:    1,
//         }
//         return repo.DB.WithContext(ctx).Create(&newActivity).Error
//     } 

//     return repo.DB.WithContext(ctx).
//         Model(&userActivity).
//         Update("count", gorm.Expr("count + ?", 1)). 
//         Error
// }