package lib_db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/8glabs/lib_db/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (repo *Repository) CreateUserAndWallet(user *models.User) error {
	if user.ChainWallet == nil {
		return errors.New("invalid wallet")
	}
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *Repository) CreateMomentAndNfts(nftCollection *models.NftCollection) error {
	result := repo.DB.Create(&nftCollection)
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

func (repo *Repository) GetAllCreators(from int, to int) (*[]models.User, error) {
	var creators []models.User
	err := repo.DB.Where("user_type = ?", models.UserType(models.USER_TYPE_CREATOR)).Offset(from).Limit(to).Find(&creators).Error
	if err != nil {
		return nil, err
	}
	return &creators, nil
}

func GetNftInfo(nftId uuid.UUID, db *sql.DB) (*models.Nft, error) {
	sqlStatement := `SELECT t1.id AS nft_id, t2.collection_name AS nft_collection_name, t4.id AS creator_id, t4.display_name AS creator_name, t1.owner_id AS owner_id,
						t7.display_name AS owner_name, t2.description AS nft_collection_description, t1.serial_id AS serial_id, t2.nfts_amount AS nfts_amount,
						t2.lowest_ask_price AS lowest_ask_price, t2.highest_sale_price AS highest_sale_price, t1.media_url AS media_url, t1.media_type AS media_type,
						t4.avatar_url AS creator_avatar_url,  t7.avatar_url AS owner_avatar_url, t1.txn_status AS txn_status, t1.txn_type AS txn_type, t1.fixed_price AS ask_price, t2.id AS nft_collection_id
					FROM nfts AS t1, nft_collections AS t2, users AS t4,
						(SELECT t5.display_name, t5.avatar_url
						FROM users AS t5, nfts AS t6
						WHERE t6.id = $1
						AND t5.id = t6.owner_id) AS t7
					WHERE t1.id = $1
					AND t1.nft_collection_id = t2.id
					AND t1.creator_id = t4.id`

	var nft models.Nft

	err := db.QueryRow(sqlStatement, nftId).Scan(&nft.Id, &nft.NftCollection.CollectionName, &nft.CreatorId, &nft.Creator.DisplayName, &nft.OwnerId, &nft.Owner.DisplayName, &nft.NftCollection.Description, &nft.SerialId, &nft.NftCollection.NftsAmount, &nft.NftCollection.HighestSalePrice, &nft.MediaUrl, &nft.MediaType, &nft.Creator.AvatarUrl, &nft.Owner.AvatarUrl, &nft.TxnStatus, &nft.TxnType, &nft.FixedPrice, &nft.NftCollectionId)

	// Returning ErrNoRows is expected behavior when no rows found, and it's not problemic. Simplely return an empty map.
	if err == sql.ErrNoRows {
		// TODO(hanhu): Return self-defined error, to not expose sql package outside
		return &nft, err
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Query nft info failed %v", err)
	}

	return &nft, err
}

func GetChainWalletByUserId(userId uint64, db *sql.DB) (*models.ChainWallet, error) {
	sqlStatement := `SELECT t2.id, t2.custodial_starkex_public_key, t2.custodial_starkex_private_key, t2.custodial_starkex_next_vault_id, t2.custodial_ethereum_wallet_address, t2.custodial_ethereum_private_key
					FROM users AS t1, chain_wallets AS t2
					WHERE t1.id = $1
					AND t2.owner_id = t1.id`

	var chainWallets models.ChainWallet

	err := db.QueryRow(sqlStatement, userId).Scan(
		&chainWallets.Id,
		&chainWallets.CustodialStarkexPublicKey,
		&chainWallets.CustodialStarkexPrivateKey,
		&chainWallets.CustodialStarkexNextVaultId,
		&chainWallets.CustodialEthereumWalletAddress,
		&chainWallets.CustodialEthereumPrivateKey)

	// Returning ErrNoRows is expected behavior when no rows found, and it's not problemic. Simplely return an empty map.
	if err == sql.ErrNoRows {
		// TODO(hanhu): Return self-defined error, to not expose sql package outside
		return &chainWallets, err
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Query user by id failed %v", err)
	}
	return &chainWallets, err
}
