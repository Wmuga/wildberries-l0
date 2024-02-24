package entity

import (
	"math/rand"
	"time"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	idRunes     = []rune("0123456789ABCDEF")
	numbers     = []rune("0123456789")
)

func GenerateOrder() *Order {
	uid := generateID(8)
	trackNumber := generateID(10)
	return &Order{
		OrderUID:          uid,
		TrackNumber:       trackNumber,
		Entry:             generateString(5),
		Delivery:          generateDelivery(),
		Payment:           generatePayment(uid),
		Items:             generateItems(2, trackNumber),
		Locale:            "EN",
		InternalSignature: generateString(10),
		CustomerID:        generateID(5),
		DeliveryService:   "WB",
		ShardKey:          generateString(20),
		SmID:              5,
		DateCreated:       time.Now(),
		OofShard:          "",
	}
}

func generateDelivery() Delivery {
	return Delivery{
		Name:    generateString(10),
		Phone:   "+" + generateFromSlice(numbers, 10),
		Zip:     generateFromSlice(numbers, 6),
		Address: generateString(10),
		Region:  "GB",
		Email:   generateString(10) + "@" + generateString(4) + ".com",
	}
}

func generatePayment(transcation string) Payment {
	return Payment{
		Transaction:  transcation,
		RequestID:    generateID(6),
		Currency:     "USD",
		Provider:     generateString(5),
		Amount:       rand.Intn(60) + 10,
		Bank:         generateString(6),
		DeliveryCost: rand.Intn(10),
		GoodsTotal:   rand.Intn(300),
		CustomFee:    rand.Intn(10),
	}
}

func generateItems(count int, trackNumber string) []Item {
	items := make([]Item, count)
	for i := range items {
		items[i] = Item{
			ChrtID:      rand.Int(),
			TrackNumber: trackNumber,
			Price:       rand.Intn(50),
			RID:         generateID(5),
			Name:        generateString(20),
			Sale:        rand.Intn(10),
			Size:        "",
			NmID:        rand.Int(),
			Brand:       generateString(10),
			Status:      0,
		}
		items[i].TotalPrice = items[i].Price - items[i].Sale
	}
	return items
}

func generateString(n int) string {
	return generateFromSlice(letterRunes, n)
}

func generateID(n int) string {
	return generateFromSlice(idRunes, n)
}

func generateFromSlice(slice []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = slice[rand.Intn(len(slice))]
	}
	return string(b)
}
