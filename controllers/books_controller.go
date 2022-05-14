package controllers

import (
	"strconv"

	"github.com/alefejonatha/Http/restapi/database"
	"github.com/alefejonatha/Http/restapi/models"
	"github.com/gin-gonic/gin"
)

func ShowAllBooks(c *gin.Context) {
	db := database.GetDatabase()

	var b []models.Book
	err := db.Find(&b).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books:" + err.Error(),
		})
		return
	}
	c.JSON(200, b)
}

func ShowBook(c *gin.Context) {
	id := c.Param("id") //vem como string

	newid, err := strconv.Atoi(id) //converter para int (ver como funciona)
	if err != nil {
		c.JSON(400, gin.H{ //passando status 400 com um objeto gin.H
			"Error": "Id has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error //Essa lista vai retornar um objeto tipo gorm.DB com um .ERROR
	if err != nil {
		c.JSON(400, gin.H{
			//Parecido com map[string] interface{} {"":""} //CHAVE tipo string || VALOR tipo interface{}: Pronunciado 'empty interface', é a interface que não especifica nenhum método ou valor!
			//Uma variável declarada como interface{} pode conter um valor de string, um inteiro, qualquer tipo de estrutura, um ponteiro para um os.Fileou qualquer coisa que você possa imaginar.
			"error": "Cannot find book:" + err.Error(),
		})
		return
	}
	c.JSON(200, book) //se deu certo
}

func CreateBooK(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book) //parseando o dado do body para json
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON:" + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book:" + err.Error(),
		})
		return
	}
	c.JSON(200, book) //se deu certo
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book) //parseando o dado do body para json
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON:" + err.Error(),
		})
		return
	}

	err = db.Omit("created_at").Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update book:" + err.Error(),
		})
		return
	}
	c.JSON(200, book) //se deu certo
	//Está retornando book com created_at zerado (Verificar)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id") //vem como string

	newid, err := strconv.Atoi(id) //converter para int (ver como funciona)
	if err != nil {
		c.JSON(400, gin.H{ //passando status 400 com um objeto gin.H
			"Error": "Id has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{ //passando status 400 com um objeto gin.H
			"Error": "cannot delete book",
		})
		return
	}
	c.Status(204)
}
