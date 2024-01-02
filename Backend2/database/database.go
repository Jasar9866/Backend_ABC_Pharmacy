package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Item struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	UnitPrice    int    `json:"unit_price"`
	ItemCategory string `json:"item_category"`
}

type Invoice struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MobileNo    string `json:"mobile_no"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	BillingType string `json:"billing_type"`
}

var db *sql.DB

func init() {
	connStr := "user=abc_pharmacy_user password=Nafras@9898 dbname=ABC_Pharmacy sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func GetItems() ([]Item, error) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name, &item.UnitPrice, &item.ItemCategory)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// AddNewItem adds a new item to the database
func AddNewItem(item Item) error {
	_, err := db.Exec("INSERT INTO items (name, unit_price, item_category) VALUES ($1, $2, $3)",
		item.Name, item.UnitPrice, item.ItemCategory)
	if err != nil {
		return err
	}

	fmt.Println("Item saved successfully")

	return nil
}

// GetInvoices retrieves all invoices from the database
func GetInvoices() ([]Invoice, error) {
	rows, err := db.Query("SELECT * FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []Invoice
	for rows.Next() {
		var invoice Invoice
		err := rows.Scan(&invoice.ID, &invoice.Name, &invoice.MobileNo, &invoice.Email, &invoice.Address, &invoice.BillingType)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

// AddNewInvoice adds a new invoice to the database
func AddNewInvoice(invoice Invoice) error {
	_, err := db.Exec("INSERT INTO invoices (name, mobile_no, email, address, billing_type) VALUES ($1, $2, $3, $4, $5)",
		invoice.Name, invoice.MobileNo, invoice.Email, invoice.Address, invoice.BillingType)
	return err
}

// UpdateItem updates an existing item in the database
func UpdateItem(itemID int, updatedItem Item) error {
	_, err := db.Exec("UPDATE items SET name=$1, unit_price=$2, item_category=$3 WHERE id=$4",
		updatedItem.Name, updatedItem.UnitPrice, updatedItem.ItemCategory, itemID)
	return err
}

// DeleteItem deletes an item from the database
func DeleteItem(itemID int) error {
	_, err := db.Exec("DELETE FROM items WHERE id=$1", itemID)
	return err
}
