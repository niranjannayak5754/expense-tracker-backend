// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "expense-tracker/database"
    "expense-tracker/handlers"
)

func main() {
    router := gin.Default()
    database.InitDB()
    router.POST("/budgets", handlers.CreateBudget)
    router.GET("/budgets/:id", handlers.GetBudget)
    // router.PUT("/budgets/:id", handlers.UpdateBudget)
    // router.DELETE("/budgets/:id", handlers.DeleteBudget)
    router.Run(":8080")
}
