package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// RegisterRequest представляет тело запроса регистрации.
type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// CORSMiddleware добавляет необходимые CORS-заголовки
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// registerHandler обрабатывает запрос на регистрацию пользователя.
func registerHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Шифруем пароль с помощью bcrypt.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка шифрования пароля"})
		return
	}

	// Вставляем данные в таблицу users.
	_, err = db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", req.Login, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка регистрации пользователя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно зарегистрирован"})
}

func main() {
	var err error
	db, err = connectDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Создаем таблицу, если её нет.
	err = initDB(db)
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}

	router := gin.Default()
	router.Use(CORSMiddleware()) // Подключаем CORS middleware

	router.POST("/register", registerHandler)
	router.Run("0.0.0.0:8080")

}
