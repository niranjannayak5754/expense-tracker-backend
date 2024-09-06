package handlers

import (
	"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	"expense-tracker/database"
	"log"
)

func CreateBudget(c *gin.Context) {
	var budget struct {
		UserID uint    `json:"user_id"`
		Name   string  `json:"name"`
		Amount float64 `json:"amount"`
	}
	if err := c.BindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO budgets (user_id, name, amount) VALUES ($1, $2, $3) RETURNING id`
	var id uint
	err := database.DB.QueryRow(query, budget.UserID, budget.Name, budget.Amount).Scan(&id)
	if err != nil {
		log.Println("Error inserting budget:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create budget"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetBudget(c *gin.Context) {
	id := c.Param("id")

	query := "SELECT user_id, name, amount FROM budgets WHERE id = $1"
	row := database.DB.QueryRow(query, id)

	var userID uint
	var name string
	var amount float64

	err := row.Scan(&userID, &name, &amount)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
		} else {
			log.Println("Error retrieving budget:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve budget"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID, "name": name, "amount": amount})
}
