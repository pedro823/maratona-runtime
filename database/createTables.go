package database

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/pedro823/maratona-runtime/model"
	"log"
)

func CreateSchema(db *pg.DB, logger *log.Logger) error {
	models := model.GetAll()
	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	logger.Println("Asserted database tables exist")
	return nil
}
