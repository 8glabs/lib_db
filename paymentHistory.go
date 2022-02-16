package lib_db

import (
	"github.com/8glabs/lib_db/models"
)

func (repo *Repository) AddPaymentHistory(paymentHistory *models.PaymentHistory) (uint64, error) {
	err := repo.DB.Create(paymentHistory).Error
	if err != nil {
		return 0, err
	}
	return paymentHistory.Id, nil
}
