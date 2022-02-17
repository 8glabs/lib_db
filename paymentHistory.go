package lib_db

//*models.PaymentHistory
func (repo *Repository) AddPaymentHistory(paymentHistory interface{}) error {
	err := repo.DB.Create(paymentHistory).Error
	if err != nil {
		return err
	}
	// return paymentHistory.Id, nil
	return nil
}
