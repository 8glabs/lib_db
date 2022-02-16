package lib_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/8glabs/lib_db/models"
	"github.com/google/uuid"
)

func GetUserById(userId uint64, db *sql.DB) (*models.User, error) {
	sqlStatement := `SELECT id, first_name, last_name, display_name, email_address, avatar_url, bio 
					FROM users 
					WHERE id = $1`

	var user models.User

	err := db.QueryRow(sqlStatement, userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.DisplayName, &user.EmailAddress, &user.AvatarUrl, &user.Bio)

	// Returning ErrNoRows is expected behavior when no rows found, and it's not problemic. Simplely return an empty map.
	if err == sql.ErrNoRows {
		// TODO(hanhu): Return self-defined error, to not expose sql package outside
		return &user, err
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Query user by id failed %v", err)
	}

	return &user, err
}

func (repo *Repository) GetUserById(userId uint64) (*models.User, error) {
	var user models.User
	if err := repo.DB.Preload("ChainWallet").Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) GetUsers(limit int) (*[]models.User, error) {
	var users []models.User
	if err := repo.DB.Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (repo *Repository) GetUserByEmailAddress(emailAddress string) (*models.User, error) {
	var user models.User
	queryResult := repo.DB.Where("email_address = ?", emailAddress).First(&user)
	fmt.Println("Result is", queryResult)
	if queryResult.Error != nil {
		fmt.Print("Error getting user with email address")
		return nil, queryResult.Error
	}
	return &user, nil
}

func (repo *Repository) UpdateUser(user *models.User) error {
	if err := repo.DB.Model(user).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetCreatedNftCollectionSupporters(creatorId uint64) (*[]models.User, error) {
	var nfts []models.Nft
	if err := repo.DB.Where("creator_id = ?", creatorId).Find(&nfts).Error; err != nil {
		return nil, err
	}

	idMap := make(map[uuid.UUID]int)
	for _, nft := range nfts {
		if _, ok := idMap[nft.Id]; !ok {
			idMap[nft.Id] = 1
		}
	}

	idArray := make([]uuid.UUID, len(idMap))
	for k, _ := range idMap {
		idArray = append(idArray, k)
	}

	var owners []models.User
	err := repo.DB.Joins(
		"JOIN nfts on nfts.owner_id=users.id").Where(
		"nfts.creator_id in (?)", creatorId).Not("users.id = ?", creatorId).Group("users.id").Find(
		&owners).Error

	if err != nil {
		log.Printf("Query created nft collections by a user failed %v", err)
		return nil, err
	}
	return &owners, nil
}
