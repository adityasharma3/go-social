package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/adityasharma3/go-social/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgresql! ")
	return db
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert string into int %v", err)
	}

	stock, err = getStock(int(id))
	if err != nil {
		log.Fatalf("could not find stock %v", err)
	}

	json.NewEncoder(w).Encode(&stock)
}

func getStock(stockId int) (models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stock models.Stock
	dbQuery := `SELECT * FROM STOCKS WHERE stockid=$1`
	err := db.QueryRow(dbQuery, stockId).Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

	if err != nil {
		log.Fatal("Could not query stock with stock_id %v %v", stockId, err)

		switch err {
		case sql.ErrNoRows:
			log.Fatal("No rows were returned")

		default:
			log.Fatal("error occurred while running query %v", err)
		}
	}

	return stock, err
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	// decode req body into json & store it into stock variable
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatal("Failed to decode req body, %v", err)
	}

	insertedStockId, err := insertStock(stock)
	if err != nil {
		log.Fatal("error ocurred while inserting stock %v", err)
	}

	res := response{
		ID:      insertedStockId,
		Message: "Created new stock successfully!",
	}

	json.NewEncoder(w).Encode(res)
}

func insertStock(stock models.Stock) (int64, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO stocks(name, price, company) VALUES ($1, $2, %3) RETURNING stockid`
	var id int64

	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to get query row data %v", err)
	}

	return id, nil
}

func getStocks() ([]models.Stock, error) {
	db := createConnection()
	defer db.Close()

	var stocks []models.Stock
	dbQuery := `SELECT * FROM STOCKS`

	rows, err := db.Query(dbQuery)
	if err != nil {
		log.Fatal("Unable to execute the query %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var stock models.Stock
		err = rows.Scan(&stock.StockId, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatal("Unable to scan row %v", err)
		}

		stocks = append(stocks, stock)
	}

	return stocks, err
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	var stocks []models.Stock
	stocks, err := getStocks()

	if err != nil {
		log.Fatal("unable to fetch all stocks err occurred , %v", err)
	}

	json.NewEncoder(w).Encode(&stocks)
}
