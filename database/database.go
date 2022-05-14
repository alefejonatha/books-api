package database

import (
	"log"
	"time"

	"github.com/alefejonatha/Http/restapi/database/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB //vai ser minúsculo porque não vai ser exportável

func StartDB() {

	// const DB_USERNAME = "root" Trocar comentários
	// const DB_PASSWORD = "root"
	// const DB_NAME = "books"
	// const DB_HOST = "localhost"
	// const DB_PORT = "3306"

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	DB_USERNAME,
	// 	DB_PASSWORD,
	// 	DB_HOST,
	// 	DB_PORT,
	// 	DB_NAME)

	dsn := "root:root@tcp(localhost:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:  ", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)           //Definir máximo de conexões ociosas ||Define o número máximo de conexões no pool de conexões inativas
	config.SetMaxOpenConns(100)          //Definir Máximo de conexões abertas ||Define o número máximo de conexões abertas com o banco de dados.
	config.SetConnMaxLifetime(time.Hour) //Vida útil máxima da conexão ||
	migrations.RunMigrations(db)         //Executar as migrações IMPORTANTE
}

//Fechando Conexão
func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB { // vai retornar o mesmo banco, para não precisar abrir várias conexões
	return db
}
