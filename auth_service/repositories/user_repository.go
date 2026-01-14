package repositories

import (
    "credit_system/config"
    "credit_system/auth_service/helpers"
    "credit_system/auth_service/model"

    "errors"
    "fmt"
    "time"


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

func (repo *AuthRepository) FetchOwner(field, value string) (*model.StoreOwner, error) {
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

    var existingAccount model.StoreOwner
    err = repo.DB.Where("account_number = ?", signupInput.AccountNumber).First(&existingAccount).Error
    if err == nil {
        return nil, nil, nil, fmt.Errorf("user already exists with provided account number")
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
            AccountNumber: signupInput.AccountNumber,
        }

        if err := tx.Create(&storeOwner).Error; err != nil {
            return fmt.Errorf("failed to create store owner: %v", err)
        }

        storeResult = model.Store{
            ID:          uuid.New(),
            Name:        signupInput.StoreName,
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

func (repo *AuthRepository) FetchStoreByOwnerID(ownerID uuid.UUID) (*model.Store, error) {
    var store model.Store
    if err := repo.DB.
        Preload("Owner").
        Where("owner_id = ?", ownerID).First(&store).Error; err != nil {
        return nil, fmt.Errorf("store not found: %v", err)
    }
    return &store, nil
}

func (repo *AuthRepository) Login(PhoneNumber, password string) (*model.StoreOwner, *model.Token, *model.Store, error) {
    var owner *model.StoreOwner
    var store *model.Store
    
    owner, err := repo.FetchOwner("phone_number", PhoneNumber)
    if err != nil {
        return nil, nil, nil, err
    }
    if owner == nil {
        return nil, nil, nil, fmt.Errorf("user not found")
    }

    if owner.Status == "Deleted" {
        return nil, nil, nil, fmt.Errorf("user account is deleted")
    }

    if !helpers.IsValidPassword(password, owner.Password) {
        return nil, nil, nil, fmt.Errorf("invalid credentials")
    }

    store, err = repo.FetchStoreByOwnerID(owner.ID) 
    if err != nil {
        return nil, nil, nil, fmt.Errorf("could not find associated store: %v", err)
    }

    // 5. Generate Token
    tokenResult, err := repo.CreateToken(owner, config.ClientId(), config.ClientSecret(), nil)
    if err != nil {
        return nil, nil, nil, err
    }

    return owner, tokenResult, store, nil
}

func (repo *AuthRepository) UpdateSingleDataByID(ownerID uuid.UUID, field, value string) (*model.StoreOwner, error) {
    updateData := map[string]interface{}{
        field:   value,
        "updated_at": time.Now(),
    }
    
    result := repo.DB.Model(&model.StoreOwner{}).Where("id = ?", ownerID).Updates(updateData)

    if result.Error != nil {
        return nil, fmt.Errorf("failed to update user: %v", result.Error)
    }

    if result.RowsAffected == 0 {
        return nil, fmt.Errorf("user not found with ID: %v", ownerID)
    }

    var updatedOwner model.StoreOwner
    if err := repo.DB.First(&updatedOwner, "id = ?", ownerID).Error; err != nil {
        return nil, fmt.Errorf("failed to retrieve updated user: %v", err)
    }

    return &updatedOwner, nil
}

func (repo *AuthRepository) FetchOwnerByID(ownerID uuid.UUID) (*model.StoreOwner, error) {
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

