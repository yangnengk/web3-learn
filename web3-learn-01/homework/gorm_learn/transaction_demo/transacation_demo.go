package transactiondemo

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string    `gorm:"type:varchar(20);not null;"`
	Balance float64   `gorm:"type:decimal(10,2);not null;"`
	Date    time.Time `gorm:"type:datetime;"`
}

type Transaction struct {
	gorm.Model
	FromAccountID uint    `gorm:"not null"`
	ToAccountID   uint    `gorm:"not null"`
	Amount        float64 `gorm:"type:decimal(10,2);not null;"`
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	// db.Create(&Account{Name: "A", Balance: 90})
	// db.Create(&Account{Name: "B", Balance: 10})
	// db.Create(&Account{Name: "B", Balance: 10, Date: time.Now()})

	err := db.Transaction(func(tx *gorm.DB) error {
		//tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).Where("balance > ? and name = ?", 100, "A").Find(&[]Account{}) // 相当于for update
		var amount float64 = 100
		update := tx.Model(&Account{}).Where("name = ? and balance >= ?", "A", amount).Update("balance", gorm.Expr("balance - ?", amount))
		if update.RowsAffected > 0 {
			tx.Model(&Account{}).Where("name = ?", "B").Update("balance", gorm.Expr("balance + ?", amount))
			return nil
		} else {
			log.Println("余额不足")
			return errors.New("余额不足")
		}
	})
	if err != nil {
		log.Println("开启事务失败， 原因：", err)
	}

}
