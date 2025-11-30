package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type ClubData struct {
	Name       string  `json:"name"`
	City       string  `json:"city"`
	Titles     int     `json:"titles"`
	AverageAge float64 `json:"average_age"`
}

func main() {
	// Подключение к БД
	connStr := "user=andrejkozura dbname=sports_clubs host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic("Не удалось подключиться к БД: " + err.Error())
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.StaticFile("/", "./index.html")

	// Эндпоинт для поиска
	r.GET("/search", func(c *gin.Context) {
		query := "%" + c.Query("query") + "%"

		rows, err := db.Query(`
            SELECT c.name, ci.city, club_info.titles, club_info.average_age
            FROM club_info
            JOIN clubs c ON club_info.club_id = c.id
            JOIN cities ci ON club_info.city_id = ci.id
            WHERE c.name ILIKE $1 OR ci.city ILIKE $1
        `, query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var results []ClubData
		for rows.Next() {
			var data ClubData
			if err := rows.Scan(&data.Name, &data.City, &data.Titles, &data.AverageAge); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			results = append(results, data)
		}

		// Если ничего не найдено, вернуть пустой массив
		if results == nil {
			results = []ClubData{}
		}

		c.JSON(http.StatusOK, results)
	})

	// Запуск сервера
	println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
