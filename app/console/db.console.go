package console

import (
	"fmt"
	"go-fiber-minimal/database/seeder"
	"go-fiber-minimal/lib"
	"os"

	"gorm.io/gorm"
)

var DB dbManager

type dbManager struct{}

func (dbManager) Seed(g *gorm.DB) {
	if err := seeder.Seed(g); err != nil {
		fmt.Printf("Error seeder: %v \n\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Seeder successfully")
}

func (dbManager) Migrate(g *gorm.DB) {
	if err := g.AutoMigrate(lib.DB.Entity()...); err != nil {
		fmt.Printf("Error migration: %v \n\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Migrate successfully")
}

func (dbManager) Drop(g *gorm.DB) {
	if err := g.Migrator().DropTable(lib.DB.Entity()...); err != nil {
		fmt.Printf("Error drop all table: %v \n\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Deleted all table successfully.")
}
