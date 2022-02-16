package lib_db

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/8glabs/lib_db/models"
	"github.com/8glabs/lib_db/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

var (
	NumUsers   int = 20
	NumMoments int = 10
	// User 0 created 6 collections
	NumMomentsOfFirstUser int = 6
	// First 2 of them are gifting
	NumGiftingOfFirstUser int = 2
	// Next 2 of them are primary sale
	NumPrimarySaleOfFirstUser int = 2
	// User 1 created 2 collections
	NumMomentsOfSecondUser int = 2
	// User 2 created 2 collections
	NumMomentsOfThirdUser int = 2

	TXN_STATUS_LIST []models.TxnStatus = []models.TxnStatus{
		models.TXN_STATUS_NOT_LISTED,
		models.TXN_STATUS_LISTED,
		models.TXN_STATUS_PENDING,
		models.TXN_STATUS_PAYMENT_MADE,
		models.TXN_STATUS_FAILED,
	}
)

func SeedDb(repo *Repository, momentVideoUrls *[]string, userImgUrls *[]string, videoCoverImgUrls *[]string) error {
	err := MigrateDb(repo)
	if err != nil {
		return err
	}
	// Initialize data
	// Users.
	// 10 in total. 0,1,2 are creators.
	gofakeit.Seed(19260817) // +1s
	// users := make([]models.User, NumUsers)
	for i := 0; i < NumUsers; i++ {
		var user models.User
		if i == 1 { // Generate meishmore
			user.DisplayName = "meishmoe"
			user.UserName = "meishmoe1"
			user.UserType = models.UserType(models.USER_TYPE_CREATOR)
			user.AvatarUrl = "https://storage.googleapis.com/creatornfts-bucket-dev/test-files/meish-moe-1.jpeg"
			user.EmailAddress = "admin-exec@8glabs.com"
			user.HashedPassword = "324a7abd7a4d86895eabf8a8fb879fdc" // meishmoe123#
			user.Bio = `We all want to be heard; this is my way of making sure you listen.
			Thank you for recognizing me and my work.
			Soon enough these sketches will be on bigger screens.
			This channel will begin my career on the bigger screen.
			You will never get what you get here, anywhere else.
			If there's anyone who can take me to the next level reading this... Just know I've been waiting for you`
			user.Tags = append(user.Tags, "comedian")
			user.InstagramUrl = "https://www.instagram.com/meishmoe/?hl=en"
			user.TwitterUrl = "https://twitter.com/meishmoe"
			user.RegisterChannel = models.RegisterChannel(models.EMAIL_REGISTER)
			user.YoutubeUrl = "https://www.youtube.com/channel/UC-ycMMkHVRjIGh65rsrdamA"
			user.IntroYoutubeEmbedUrl = "https://www.youtube.com/embed/OV250xp44xA"
		} else if i == 2 {
			user.DisplayName = "Hope For Paws"
			user.UserName = "HopeForPaws"
			user.UserType = models.UserType(models.USER_TYPE_CREATOR)
			user.AvatarUrl = "https://storage.googleapis.com/creatornfts-bucket-dev/test-files/hopeforpaw.png"
			user.EmailAddress = "eldad75@gmail.com"
			user.HashedPassword = "535203e3991cbc0e939f9f42bd8d0b4c" // hopeforpaw12345@
			user.Bio = `Hope For Paws is a 501 C-3 non-profit animal rescue organization (E.I.N: 26-2869386). We rescue dogs, cats and wildlife suffering on the streets or are seriously injured in shelters. Through our rescue missions, Hope For Paws works to raise awareness for abandoned animals. Every NFT sold will help us save another precious life ❤️`
			user.Tags = []string{"HopeForPaws", "Hope For Paws", "Animal Rescue", "Dog Rescue", "Cat Rescue"}
			user.YoutubeUrl = "https://www.youtube.com/channel/UCdu8QrpJd6rdHU9fHl8J01A"
			user.RegisterChannel = models.RegisterChannel(models.EMAIL_REGISTER)
		} else {
			gofakeit.Struct(&user)
			imgUrlIndex := i % len(*userImgUrls)
			user.AvatarUrl = (*userImgUrls)[imgUrlIndex]
			if i == 5 { // User 5 has no avatar url
				user.AvatarUrl = ""
			} else if i == 6 { // User 6 has no youtube url
				user.YoutubeUrl = ""
			} else if i == 7 {
				user.InstagramUrl = ""
			} else if i == 8 { // User 8 has no bio
				user.Bio = ""
			}
		}
		// We have the first three as creators
		if i < 3 {
			user.UserType = models.UserType(models.USER_TYPE_CREATOR)
		}
		// Create wallets for them
		wallet := utils.GenerateChainWallet()
		wallet.OwnerId = user.Id
		user.ChainWallet = wallet
		// Write users with their wallets to DB.
		err := repo.CreateUserAndWallet(&user) // pass pointer of data to Create
		if err != nil {
			return err
		}
		fiatWallet := models.FiatWallet{
			OwnerId:         user.Id,
			BalanceCurrency: "usd",
			Balance:         float64(rand.Intn(100)),
		}
		err = repo.DB.Create(&fiatWallet).Error
		if err != nil {
			fmt.Println("Failed to create fiat wallet for user")
		}
	}
	// Get all users after insertion. Limit to 50
	users, err := repo.GetUsers(50)
	if err != nil {
		return err
	}
	// Moments
	// Creators: 0,1,2. They have moments
	// Distribution: User 0 has 0~5. User 1 has 6,7,8. User 2 has 9
	for i := 0; i < NumMoments; i++ {
		creatorId := uint64(0)
		if i < NumMomentsOfFirstUser {
			creatorId = (*users)[0].Id
		} else if i < NumMomentsOfFirstUser+NumMomentsOfSecondUser {
			creatorId = (*users)[1].Id
		} else {
			creatorId = (*users)[2].Id
		}
		// Create associated NFT collection
		var nftCollection models.NftCollection
		gofakeit.Struct(&nftCollection)
		// nftCollection 0,1 are in gifting; 2,3 are primary sale, all others are secondary sale
		if i < NumGiftingOfFirstUser {
			nftCollection.Status = models.GIFTING
		} else if i < NumGiftingOfFirstUser+NumPrimarySaleOfFirstUser {
			nftCollection.Status = models.PRIMARY_SALE
		} else {
			nftCollection.Status = models.SECONDARY_SALE
		}
		// nftCollection.SolanaMetadataAccountSecretKey = ""
		primarySalePrice := float64(0)
		if i == 0 {
			nftCollection.NftsAmount = 100
			primarySalePrice = 20
			nftCollection.Status = models.PRIMARY_SALE
			nftCollection.Description = ` It can be kinda awkaward ya know?
			Leave a like & Sub if you smiled`
			nftCollection.CollectionName = "Being black in history class"
		}
		numNftsInCollection := nftCollection.NftsAmount
		if numNftsInCollection > 30 {
			nftCollection.Rarity = models.Rarity(models.Common)
		} else if numNftsInCollection > 15 {
			nftCollection.Rarity = models.Rarity(models.Uncommon)
		} else if numNftsInCollection > 10 {
			nftCollection.Rarity = models.Rarity(models.Rare)
		} else if numNftsInCollection > 5 {
			nftCollection.Rarity = models.Rarity(models.Heroic)
		} else {
			nftCollection.Rarity = models.Rarity(models.Mythic)
		}
		nftsInCollection := make([]models.Nft, numNftsInCollection)
		for j := 0; j < numNftsInCollection; j++ {
			var nft models.Nft
			gofakeit.Struct(&nft)
			nft.TxnStatus = models.TxnStatus(models.TXN_STATUS_NOT_LISTED)
			nft.SerialId = j + 1
			nft.TxnType = models.TxnType(models.TXN_TYPE_FIXED_PRICE)
			nft.FixedPriceCurrency = "usd"
			nft.InitialNftCollectionStatus = nftCollection.Status
			nft.ChainType = models.ChainType(models.CHAIN_TYPE_SOLANA)
			nft.CreatorId = creatorId

			if i == 1 {
				nft.MediaType = models.MediaType(models.MEDIA_TYPE_VIDEO)
				nft.MediaUrl = "https://storage.googleapis.com/creatornfts-bucket-dev/test-files/output_960x540.mp4"
				nft.CoverImageUrl = "https://storage.googleapis.com/creatornfts-bucket-dev/test-files/meishmoe_cover.png"
			} else {
				nft.MediaType = models.MediaType(models.MEDIA_TYPE_VIDEO)
				nft.MediaUrl = (*momentVideoUrls)[i%len(*momentVideoUrls)]
				nft.CoverImageUrl = (*videoCoverImgUrls)[i%len(*videoCoverImgUrls)]
			}
			// Assign owner
			if i < NumGiftingOfFirstUser {
				// NFTs in gifting can be owned by anyone. Set first 66.7% to creator
				// Others allocate randomly
				if (j+1)*3 < numNftsInCollection*2 {
					nft.OwnerId = creatorId
				} else {
					nft.OwnerId = (*users)[rand.Intn(NumUsers)].Id
				}
				// txn_authority should be PROGRAM
				nft.TxnAuthority = models.TxnAuthority(models.PROGRAM)
				// Set MSRP
				nft.FixedPrice = primarySalePrice

			} else if i < NumGiftingOfFirstUser+NumPrimarySaleOfFirstUser {
				// NFTs in primary sale can be owned by anyone. Set first 66.7% to creator
				// Others allocate randomly
				if (j+1)*3 < numNftsInCollection*2 {
					nft.OwnerId = creatorId
				} else {
					nft.OwnerId = (*users)[rand.Intn(NumUsers)].Id
				}
				// txn_authority should be CREATOR
				nft.TxnAuthority = models.TxnAuthority(models.CREATOR)
				// In primary sale, when owner is not the creator, txn_status must be NOT_LISTED
				// Because only a few NFTs is owned by the creator, we
				// mark all the nfts lowned by creator with LISTED
				if nft.OwnerId == creatorId {
					nft.TxnStatus = models.TxnStatus(models.TXN_STATUS_LISTED)
				} else {
					nft.TxnStatus = models.TxnStatus(models.TXN_STATUS_NOT_LISTED)
				}
				nft.FixedPrice = primarySalePrice
				// for those that are in txn, buyer can be anyone except owner. They should also have txn id
				if nft.TxnStatus != models.TXN_STATUS_LISTED && nft.TxnStatus != models.TXN_STATUS_NOT_LISTED {
					nft.BuyerId, err = PickUserIdExcept(nft.OwnerId)
					if err != nil {
						return err
					}
					nft.TxnId, err = uuid.NewRandom()
					if err != nil {
						return err
					}
				}
			} else {
				// NFTs in secondary sale can be owned by anyone
				nft.OwnerId = (*users)[rand.Intn(NumUsers)].Id
				// txn_authority should be OWNER
				nft.TxnAuthority = models.TxnAuthority(models.OWNER)
				// txn_status can be anything
				nft.TxnStatus = models.TxnStatus(TXN_STATUS_LIST[rand.Intn(len(TXN_STATUS_LIST))])
				// for those that are in txn, buyer can be anyone except owner. They should also have txn id
				if nft.TxnStatus != models.TXN_STATUS_LISTED && nft.TxnStatus != models.TXN_STATUS_NOT_LISTED {
					nft.BuyerId, err = PickUserIdExcept(nft.OwnerId)
					if err != nil {
						return err
					}
					nft.TxnId, err = uuid.NewRandom()
					if err != nil {
						return err
					}
				}
			}
			if i == 0 {
				nft.OwnerId = (*users)[0].Id
				nft.TxnAuthority = models.TxnAuthority(models.CREATOR)
				nft.TxnStatus = models.TxnStatus(models.TXN_STATUS_LISTED)
			}
			nftsInCollection[j] = nft
		}
		nftCollection.Nfts = &nftsInCollection
		err := repo.CreateMomentAndNfts(&nftCollection)
		if err != nil {
			return err
		}
		if nftCollection.Status == models.GIFTING {
			var giftingCampaign models.GiftingCampaign
			giftingCampaign.NftCollectionId = nftCollection.Id
			giftingCampaign.CreatorId = creatorId
			giftingCampaign.PrimarySalePrice = primarySalePrice
			giftingCampaign.NextAvailableSerialId = 1
			// By default, started yesterday and ends tomorrow
			startTime := time.Now().Add(-24 * time.Hour)
			giftingCampaign.GiftingStartTime = &startTime
			endTime := time.Now().Add(24 * time.Hour)
			giftingCampaign.GiftingEndTime = &endTime
			err := repo.CreateGiftingCampaign(giftingCampaign)
			if err != nil {
				return err
			}
		} else if nftCollection.Status == models.PRIMARY_SALE {
			var primarySaleCampaign models.PrimarySaleCampaign
			primarySaleCampaign.NftCollectionId = nftCollection.Id
			primarySaleCampaign.CreatorId = creatorId
			primarySaleCampaign.PrimarySalePrice = primarySalePrice
			// By default, started yesterday and ends tomorrow
			startTime := time.Now().Add(-24 * time.Hour)
			primarySaleCampaign.PrimarySaleStartTime = &startTime
			endTime := time.Now().Add(24 * time.Hour)
			primarySaleCampaign.PrimarySaleEndTime = &endTime
			err := repo.CreatePrimarySaleCampaign(primarySaleCampaign)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func PickUserIdExcept(expectId uint64) (uint64, error) {
	pickedId := rand.Intn(NumUsers-1) + 1
	numTries := 0
	for pickedId == int(expectId) {
		pickedId = rand.Intn(NumUsers-1) + 1
		numTries++
		if numTries >= 20 {
			return 0, errors.New("failed to pick user id")
		}
	}
	return uint64(pickedId), nil
}
