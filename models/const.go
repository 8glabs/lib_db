package models

type TxnStatus string

const (
	TXN_STATUS_NOT_LISTED   TxnStatus = "NOT_LISTED"
	TXN_STATUS_LISTED       TxnStatus = "LISTED"
	TXN_STATUS_PENDING      TxnStatus = "PENDING"
	TXN_STATUS_PAYMENT_MADE TxnStatus = "PAYMENT_MADE"
	TXN_STATUS_FAILED       TxnStatus = "FAILED"
)

type NftCollectionStatus string

const (
	GIFTING        NftCollectionStatus = "GIFTING"
	PRIMARY_SALE   NftCollectionStatus = "PRIMARY_SALE"
	SECONDARY_SALE NftCollectionStatus = "SECONDARY_SALE"
)

type TxnAuthority string

const (
	ADMIN   TxnAuthority = "ADMIN"
	PROGRAM TxnAuthority = "PROGRAM"
	CREATOR TxnAuthority = "CREATOR"
	OWNER   TxnAuthority = "OWNER"
)

type UserType string

const (
	USER_TYPE_NORMAL  UserType = "USER_TYPE_NORMAL"
	USER_TYPE_CREATOR UserType = "USER_TYPE_CREATOR"
)
