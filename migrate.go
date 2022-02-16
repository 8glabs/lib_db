package lib_db

import (
	"github.com/8glabs/lib_db/models"
)

func MigrateDb(repo *Repository) error {
	migrator := repo.DB.Migrator()
	repo.DB.Migrator().DropTable(&models.AccountToken{})
	repo.DB.Migrator().DropTable(&models.GiftingCampaign{})
	migrator.DropTable(&models.PrimarySaleCampaign{})
	repo.DB.Migrator().DropTable(&models.NftTxnHistory{})
	repo.DB.Migrator().DropTable(&models.PaymentHistory{})
	migrator.DropTable(&models.ChainWallet{})
	migrator.DropTable(&models.FiatWallet{})
	repo.DB.Migrator().DropTable(&models.Nft{})
	migrator.DropTable(&models.NftCollection{})
	migrator.DropTable(&models.User{})

	repo.DB.AutoMigrate(
		&models.ChainWallet{},
		&models.FiatWallet{},
		&models.User{},
		&models.Nft{},
		&models.NftCollection{},
		&models.NftTxnHistory{},
		&models.PaymentHistory{},
		&models.PrimarySaleCampaign{},
		&models.GiftingCampaign{},
		&models.AccountToken{},
	)
	return nil
}
