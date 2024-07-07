package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateQuery(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../recitewords/internal/query", // default “./dao/query”
		ModelPkgPath:  "./models",
		OutFile:       "query.go", //default：gen.go
		FieldNullable: true,
		WithUnitTest:  true,
	})

	g.UseDB(db)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}

func main() {
	dsn := "root:123456@tcp(localhost:3306)/recite_english?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("Start...")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	GenerateQuery(db)
}
