package lib_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/8glabs/lib_db/models"
	_ "github.com/lib/pq"
)

//*models.User
func (repo *Repository) CreateUserAndWallet(data interface{}) error {
	// if user.ChainWallet == nil {
	// 	return errors.New("invalid wallet")
	// }
	result := repo.DB.Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//*models.NftCollection
func (repo *Repository) CreateMomentAndNfts(data interface{}) error {
	result := repo.DB.Create(data)
	if result.Error != nil {
		fmt.Println("Failed to insert nft collection")
		return result.Error
	}
	return nil
}

func CheckEmailExist(emailAddress string, db *sql.DB) (uint64, bool) {
	var id uint64
	isExist := true

	err := db.QueryRow("SELECT id FROM users WHERE email_address = $1", emailAddress).Scan(&id)
	if err != nil {
		log.Printf("Can't find user by email_address %v. %v", emailAddress, err)
		return id, !isExist
	}

	return id, isExist
}

//*[]models.User
func (repo *Repository) GetAllCreators(from int, to int, data interface{}) error {
	err := repo.DB.Where("user_type = ?", models.UserType(models.USER_TYPE_CREATOR)).Offset(from).Limit(to).Find(data).Error
	if err != nil {
		return err
	}
	return nil
}
