package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Transatcion struct to hold each transaction's details
type Transaction struct {
	ID 			int
	Amount 		float64
	Category 	string
	Date 		time.Time
	Type 		string
}

type BudgetTracker struct {
	transactions	[]Transaction
	nextID			int
}

// Interface for common behavior, as a key to achieve polymorphic
type FinancialRecord interface {
	GetAmount() float64
	GetType()	string
}

// Implement interface methods for Transaction struct
func (t Transaction) GetAmount() float64 {
	return t.Amount
}

func (t Transaction) GetType() string {
	return t.Type
}

// Add new Transaction

func (bt *BudgetTracker) AddTransaction(amount float64, category, tType string) {
	newTransaction := Transaction {
		ID: bt.nextID,
		Amount: amount,
		Category: category,
		Date: time.Now(),
		Type: tType,
	}
	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextID++
}

// Creating DisplayTransaction method
func (bt BudgetTracker) DisplayTransaction() {
	fmt.Println("ID\tAmount\tCategory\tDate\tType")
	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n", 
		transaction.ID, transaction.Amount,
		transaction.Category, transaction.Date.Format("2006-01-02"), transaction.Type)
	}
}

func (bt BudgetTracker) calculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.transactions{
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

func (bt BudgetTracker) SaveToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file) // creating a new CSV file
	defer writer.Flush() // flush is very important to make sure that data is written before the file is closed
	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"})

	// Write Data
	for _, t := range bt.transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("2006-01-02"),
			t.Type,
		}
		writer.Write(record)
	}
	fmt.Println("Transactions saved to", filename)
	return nil
}

func main() {
	bt := BudgetTracker{}
	for {
		fmt.Println("\n--- Personal Budget Tracker ---")
		fmt.Println("1. Add Transactions")
		fmt.Println("2. Display Transations")
		fmt.Println("3. Show Total Income")
		fmt.Println("4. Show Total Expenses")
		fmt.Println("5. Save Transatcions to CSV")
		fmt.Println("6. Exit")
		fmt.Println("Choose an option: ")
	
	var choice int
	fmt.Scanln(&choice)

	switch choice {
		case 1:
            fmt.Print("Enter Amount: ")
			var amount float64
			fmt.Scanln(&amount)

            fmt.Print("Enter Category: ")
			var category string
			fmt.Scanln(&category)

            fmt.Print("Enter Income/Expense: ")
			var tType string
			fmt.Scanln(&tType)

            bt.AddTransaction(amount, category, tType)
			fmt.Print("Transaction Added")
            break

        case 2:
            bt.DisplayTransaction()
            break
        case 3:
            fmt.Printf("Total Income: %.2f\n", bt.calculateTotal("Income"))
            break
        case 4:
			fmt.Printf("Total Expenses: %.2f\n", bt.calculateTotal("expense"))
            break
        case 5:
			fmt.Printf("Enter filename (e.g., transactions.csv): ")
			var filename string
			fmt.Scanln(&filename)
			if err:= bt.SaveToCSV(filename); err!= nil {
				fmt.Println("Error saving transactions to CSV:", err)
			}
            break
        case 6:
            fmt.Println("Exiting...")
            os.Exit(0)
			return
        default:
            fmt.Println("Invalid option!")
		}
	}
    
}