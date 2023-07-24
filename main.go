package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "serli2007"
	dbname   = "enigma_laundry"
)

var db *sql.DB

func initDB() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sucessfully Connected")
}

func main() {
	initDB()
	defer db.Close()

	var choice int
	for {
		fmt.Println("===== ENIGMA LAUNDRY =====")
		fmt.Println("===== MENU CUSTOMERS =====")
		fmt.Println("1. View Customers")
		fmt.Println("2. Add New Customer")
		fmt.Println("3. Edit Customer")
		fmt.Println("4. Delete Customer")
		fmt.Println("===== MENU SERVICES =====")
		fmt.Println("5. View Services")
		fmt.Println("6. Add New Service")
		fmt.Println("7. Edit Service")
		fmt.Println("8. Delete Service")
		fmt.Println("===== MENU TRANSACTIONS =====")
		fmt.Println("9. View Transactions")
		fmt.Println("10. Add New Transaction")
		fmt.Println("11. Cetak Invoice")
		fmt.Println("0. Exit")
		fmt.Print("Select Menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			fmt.Println("Exit the application.")
			os.Exit(0)
		case 1:
			viewCustomers()
		case 2:
			addNewCustomer()
		case 3:
			editCustomer()
		case 4:
			deleteCustomer()
		case 5:
			viewServices()
		case 6:
			addNewService()
		case 7:
			editService()
		case 8:
			deleteService()
		case 9:
			viewTransactions()
		case 10:
			addNewTransaction()
		case 11:
			cetakInvoice()
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func viewCustomers() {
	rows, err := db.Query("SELECT customer_id, customer_name, phone_number FROM Customer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("===== Customer List =====")
	fmt.Printf("%-10s %-30s %-15s\n", "ID", "Customer Name", "Phone Number")
	fmt.Println("----------------------------------------")
	for rows.Next() {
		var customerID int
		var customerName, phoneNumber string
		err := rows.Scan(&customerID, &customerName, &phoneNumber)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-10d %-30s %-15s\n", customerID, customerName, phoneNumber)
	}
	fmt.Println()
}

func addNewCustomer() {
	var customerName, phoneNumber string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Customer Name: ")
	customerName, _ = reader.ReadString('\n')
	customerName = strings.TrimSpace(customerName)

	fmt.Print("Enter Phone Number: ")
	phoneNumber, _ = reader.ReadString('\n')
	phoneNumber = strings.TrimSpace(phoneNumber)

	//Validasi dasar pada input data
	if len(customerName) == 0 || len(phoneNumber) == 0 {
		fmt.Println("Invalid input. Customer Name and Phone Number cannot be empty.")
		return
	}

	// Validasi nomor telepon untuk memastikan hanya berisi angka (0-9)
	if !isValidPhoneNumber(phoneNumber) {
		fmt.Println("Invalid input. Phone Number must only contain digits (0-9).")
		return
	}

	// Dapatkan ID customer terakhir dari database
	lastCustomerID, err := getLastCustomerID()
	if err != nil {
		log.Fatal(err)
	}

	// Hitung ID customer berikutnya
	nextCustomerID := lastCustomerID + 1

	// Masukkan customer baru ke dalam database dengan nextCustomerID yang dihitung
	result, err := db.Exec("INSERT INTO Customer (customer_id, customer_name, phone_number) VALUES ($1, $2, $3)", nextCustomerID, customerName, phoneNumber)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("%d customer added successfully.\n", rowsAffected)
}

// Jika tidak ada catatan dalam tabel, kembalikan ID customer 0 (nilai awal)
func getLastCustomerID() (int, error) {
	var lastCustomerID int
	err := db.QueryRow("SELECT MAX(customer_id) FROM Customer").Scan(&lastCustomerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return lastCustomerID, nil
}

// Nomor telepon hanya berisi angka (0-9)
func isValidPhoneNumber(phoneNumber string) bool {
	for _, char := range phoneNumber {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func editCustomer() {
	// Menampilkan daftar customer untuk dipilih pengguna
	viewCustomers()

	var customerID int
	fmt.Print("Enter Customer ID to edit: ")
	fmt.Scanln(&customerID)

	// Periksa ID customer valid
	if !isCustomerIDValid(customerID) {
		fmt.Println("Invalid Customer ID.")
		return
	}

	var newCustomerName, newPhoneNumber string

	fmt.Print("Enter New Customer Name: ")
	fmt.Scanln(&newCustomerName)
	fmt.Print("Enter New Phone Number: ")
	fmt.Scanln(&newPhoneNumber)

	// validasi dasar pada input data
	if len(newCustomerName) == 0 || len(newPhoneNumber) == 0 {
		fmt.Println("Invalid input. New Customer Name and Phone Number cannot be empty.")
		return
	}

	// Validasi nomor telepon untuk memastikan hanya berisi angka (0-9)
	if !isValidPhoneNumber(newPhoneNumber) {
		fmt.Println("Invalid input. New Phone Number must only contain digits (0-9).")
		return
	}

	// Perbarui customer di database
	_, err := db.Exec("UPDATE Customer SET customer_name = $1, phone_number = $2 WHERE customer_id = $3", newCustomerName, newPhoneNumber, customerID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Customer updated successfully.")
}

func deleteCustomer() {
	// Menampilkan daftar customer untuk dipilih pengguna
	viewCustomers()

	var customerID int
	fmt.Print("Enter Customer ID to delete: ")
	fmt.Scanln(&customerID)

	// Periksa apakah ID customer valid
	if !isCustomerIDValid(customerID) {
		fmt.Println("Invalid Customer ID.")
		return
	}

	// Periksa apakah customer memiliki transaksi di database
	if hasTransactions(customerID) {
		fmt.Println("Cannot delete customer. This customer has transactions.")
		return
	}

	// Hapus customer dari database
	_, err := db.Exec("DELETE FROM Customer WHERE customer_id = $1", customerID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Customer deleted successfully.")
}

func hasTransactions(customerID int) bool {
	// Memeriksa apakah customer memiliki transaksi
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM Transaction WHERE customer_id = $1", customerID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func viewServices() {
	rows, err := db.Query("SELECT service_id, service_name, service_unit, price_per_unit FROM Service")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("===== Service List =====")
	fmt.Printf("%-10s %-30s %-15s %-10s\n", "Service ID", "Service Name", "Unit", "Price Per Unit")
	fmt.Println("----------------------------------------------------")
	for rows.Next() {
		var serviceID, pricePerUnit int
		var serviceName, serviceUnit string
		err := rows.Scan(&serviceID, &serviceName, &serviceUnit, &pricePerUnit)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-10d %-30s %-15s %-10d\n", serviceID, serviceName, serviceUnit, pricePerUnit)
	}
	fmt.Println()
}

func addNewService() {
	var serviceName, serviceUnit string
	var pricePerUnit int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Service Name: ")
	serviceName, _ = reader.ReadString('\n')
	serviceName = strings.TrimSpace(serviceName)

	fmt.Print("Enter Unit: ")
	serviceUnit, _ = reader.ReadString('\n')
	serviceUnit = strings.TrimSpace(serviceUnit)

	fmt.Print("Enter Price Per Unit: ")
	fmt.Scanln(&pricePerUnit)

	// validasi dasar pada input data
	if len(serviceName) == 0 || len(serviceUnit) == 0 || pricePerUnit <= 0 {
		fmt.Println("Invalid input. Service Name, Unit, and Price Per Unit must be provided and Price Per Unit must be greater than 0.")
		return
	}

	// Periksa apakah layanan dengan nama yang sama sudah ada di database
	var existingServiceID int
	err := db.QueryRow("SELECT service_id FROM Service WHERE service_name = $1", serviceName).Scan(&existingServiceID)
	if err == nil {
		fmt.Println("Service with the same name already exists. Please choose a different name.")
		return
	}

	// Dapatkan ID layanan terakhir dari database
	lastServiceID, err := getLastServiceID()
	if err != nil {
		log.Fatal(err)
	}

	// Hitung ID Service selanjutnya
	nextServiceID := lastServiceID + 1

	// Masukkan layanan baru ke dalam database dengan nextServiceID yang dihitung
	result, err := db.Exec("INSERT INTO Service (service_id, service_name, service_unit, price_per_unit) VALUES ($1, $2, $3, $4)", nextServiceID, serviceName, serviceUnit, pricePerUnit)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("%d service added successfully.\n", rowsAffected)
}

func getLastServiceID() (int, error) {
	var lastServiceID int
	err := db.QueryRow("SELECT MAX(service_id) FROM Service").Scan(&lastServiceID)
	if err != nil {
		// If there are no records in the table, return service ID 0 (starting value)
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return lastServiceID, nil
}

func editService() {
	// Menampilkan daftar service untuk dipilih pengguna
	viewServices()

	var serviceID, pricePerUnit int
	var serviceName, serviceUnit string

	fmt.Print("Enter Service ID to edit: ")
	fmt.Scanln(&serviceID)

	// Periksa apakah ID Service valid
	if !isServiceIDValid(serviceID) {
		fmt.Println("Invalid Service ID.")
		return
	}

	// Dapatkan detail layanan saat ini dari database
	row := db.QueryRow("SELECT service_name, service_unit, price_per_unit FROM Service WHERE service_id = $1", serviceID)
	err := row.Scan(&serviceName, &serviceUnit, &pricePerUnit)
	if err != nil {
		log.Fatal(err)
	}

	// Dapatkan detail layanan terbaru dari pengguna
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Current Service Name: %s\n", serviceName)
	fmt.Printf("Current Unit: %s\n", serviceUnit)
	fmt.Printf("Current Price Per Unit: %d\n", pricePerUnit)

	var updatedServiceName, updatedServiceUnit string
	var updatedPricePerUnitStr string

	fmt.Print("Enter Updated Service Name (0 to keep current): ")
	updatedServiceName, _ = reader.ReadString('\n')
	updatedServiceName = strings.TrimSpace(updatedServiceName)

	fmt.Print("Enter Updated Unit (0 to keep current): ")
	updatedServiceUnit, _ = reader.ReadString('\n')
	updatedServiceUnit = strings.TrimSpace(updatedServiceUnit)

	fmt.Print("Enter Updated Price Per Unit (0 to keep current): ")
	updatedPricePerUnitStr, _ = reader.ReadString('\n')
	updatedPricePerUnitStr = strings.TrimSpace(updatedPricePerUnitStr)

	// Konversikan updatedPricePerUnitStr menjadi bilangan bulat
	updatedPricePerUnit, err := strconv.Atoi(updatedPricePerUnitStr)
	if err != nil {
		fmt.Println("Invalid input. Updated Price Per Unit must be a valid integer.")
		return
	}

	// validasi dasar pada input data
	if updatedServiceName == "0" {
		updatedServiceName = serviceName
	}
	if updatedServiceUnit == "0" {
		updatedServiceUnit = serviceUnit
	}
	if updatedPricePerUnit <= 0 {
		updatedPricePerUnit = pricePerUnit
	}

	// Periksa apakah Service dengan nama yang sama sudah ada di database
	var existingServiceID int
	err = db.QueryRow("SELECT service_id FROM Service WHERE service_name = $1 AND service_id <> $2", updatedServiceName, serviceID).Scan(&existingServiceID)
	if err == nil {
		fmt.Println("Service with the same name already exists. Please choose a different name.")
		return
	}

	// Perbarui service di database
	_, err = db.Exec("UPDATE Service SET service_name = $1, service_unit = $2, price_per_unit = $3 WHERE service_id = $4",
		updatedServiceName, updatedServiceUnit, updatedPricePerUnit, serviceID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service updated successfully.")
}

func deleteService() {
	// Menampilkan daftar service untuk dipilih pengguna
	viewServices()

	var serviceID int
	fmt.Print("Enter Service ID to delete: ")
	fmt.Scanln(&serviceID)

	// Periksa apakah ID service valid
	if !isServiceIDValid(serviceID) {
		fmt.Println("Invalid Service ID.")
		return
	}

	// Hapus service dari database
	_, err := db.Exec("DELETE FROM Service WHERE service_id = $1", serviceID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service deleted successfully.")
}

func viewTransactions() {
	rows, err := db.Query("SELECT t.transaction_id, c.customer_name, s.service_name, t.quantity, t.total_price, t.entry_date, t.completion_date, t.received_by FROM Transaction t JOIN Customer c ON t.customer_id = c.customer_id JOIN Service s ON t.service_id = s.service_id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("===== Transaction List =====")
	fmt.Printf("%-10s %-20s %-20s %-8s %-10s %-12s %-12s %-10s\n", "Trans. ID", "Customer Name", "Service Name", "Quantity", "Total Price", "Entry Date", "Completion Date", "Received By")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	for rows.Next() {
		var transactionID, quantity, totalPrice int
		var customerName, serviceName, entryDate, completionDate, receivedBy string
		err := rows.Scan(&transactionID, &customerName, &serviceName, &quantity, &totalPrice, &entryDate, &completionDate, &receivedBy)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-10d %-20s %-20s %-8d %-10d %-12s %-12s %-10s\n", transactionID, customerName, serviceName, quantity, totalPrice, entryDate, completionDate, receivedBy)
	}
	fmt.Println()
}

func getLastTransactionID() (int, error) {
	var lastTransactionID int
	err := db.QueryRow("SELECT MAX(transaction_id) FROM Transaction").Scan(&lastTransactionID)
	if err != nil {
		// Jika tidak ada transaction dalam tabel atau nilainya NULL, kembalikan ID transactin 0 (nilai awal)
		if err == sql.ErrNoRows {
			return 0, nil
		}
		// Saat nilainya NULL dengan menyetelnya ke 0
		if strings.Contains(err.Error(), "converting NULL to int") {
			return 0, nil
		}
		return 0, err
	}
	return lastTransactionID, nil
}

func isTransactionDataValid(quantity, totalPrice int, entryDate, completionDate, receivedBy string) bool {
	// Validasi quantitas (harus lebih besar dari 0)
	if quantity <= 0 {
		fmt.Println("Invalid input. Quantity must be greater than 0.")
		return false
	}

	// Validasi tanggal (harus dalam format YYYY-MM-DD)
	_, err := time.Parse("2006-01-02", entryDate)
	if err != nil {
		fmt.Println("Invalid input. Entry Date must be in the format YYYY-MM-DD.")
		return false
	}

	// Validasi tanggal (harus dalam format YYYY-MM-DD)
	_, err = time.Parse("2006-01-02", completionDate)
	if err != nil {
		fmt.Println("Invalid input. Completion Date must be in the format YYYY-MM-DD.")
		return false
	}

	// Validasi yang diterima oleh (tidak boleh kosong)
	if len(receivedBy) == 0 {

		fmt.Println("Invalid input. Received By cannot be empty.")
		return false
	}

	return true
}

func calculateTotalPrice(serviceID, quantity int) int {
	// Dapatkan harga per unit service yang dipilih dari database
	var pricePerUnit int
	err := db.QueryRow("SELECT price_per_unit FROM Service WHERE service_id = $1", serviceID).Scan(&pricePerUnit)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate total price
	totalPrice := pricePerUnit * quantity

	return totalPrice
}

func addNewTransaction() {
	var customerID, serviceID, quantity, totalPrice int
	var entryDate, completionDate, receivedBy string

	// Menampilkan daftar customer yang tersedia
	viewCustomers()

	// Dapatkan ID customer dari pengguna
	fmt.Print("Enter Customer ID: ")
	fmt.Scanln(&customerID)

	// Periksa apakah ID customer valid
	if !isCustomerIDValid(customerID) {
		fmt.Println("Invalid Customer ID.")
		return
	}

	// Menampilkan daftar service yang tersedia
	viewServices()

	// Dapatkan ID service dari pengguna
	fmt.Print("Enter Service ID: ")
	fmt.Scanln(&serviceID)

	// Periksa apakah ID service valid
	if !isServiceIDValid(serviceID) {
		fmt.Println("Invalid Service ID.")
		return
	}

	// Masukan quantity transaction dari perguna
	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&quantity)

	// Hitung harga total berdasarkan harga per unit service dan quantitas yang dipilih
	totalPrice = calculateTotalPrice(serviceID, quantity)

	fmt.Print("Enter Entry Date (YYYY-MM-DD): ")
	fmt.Scanln(&entryDate)
	fmt.Print("Enter Completion Date (YYYY-MM-DD): ")
	fmt.Scanln(&completionDate)

	// AcceptBy tidak boleh kosong
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Received By: ")
		receivedBy, _ = reader.ReadString('\n')
		receivedBy = strings.TrimSpace(receivedBy)

		if len(receivedBy) > 0 {
			break
		} else {
			fmt.Println("Received By cannot be empty.")
		}
	}

	// Validasi dasar pada input data
	if !isTransactionDataValid(quantity, totalPrice, entryDate, completionDate, receivedBy) {
		fmt.Println("Invalid input data.")
		return
	}

	// Dapatkan ID transaction terakhir dari database
	lastTransactionID, err := getLastTransactionID()
	if err != nil {
		log.Fatal(err)
	}

	// Hitung ID transaction berikutnya
	nextTransactionID := lastTransactionID + 1

	// Masukkan transaction baru ke dalam database
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO Transaction (transaction_id, customer_id, service_id, quantity, total_price, entry_date, completion_date, received_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		nextTransactionID, customerID, serviceID, quantity, totalPrice, entryDate, completionDate, receivedBy)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction added successfully.")
}

func isCustomerIDValid(customerID int) bool {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM Customer WHERE customer_id = $1", customerID)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func isServiceIDValid(serviceID int) bool {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM Service WHERE service_id = $1", serviceID)
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func cetakInvoice() {
	// Tampilkan daftar customer dengan ID dan nama customer
	viewCustomersWithIDName()

	// Dapatkan ID customer dari pengguna
	var customerID int
	fmt.Print("Enter Customer ID to generate invoice: ")
	fmt.Scanln(&customerID)

	// Periksa apakah ID customer valid
	if !isCustomerIDValid(customerID) {
		fmt.Println("Invalid Customer ID.")
		return
	}

	// Hasilkan dan cetak semua transaction yang terkait dengan customer yang dipilih
	err := printTransactionsByCustomerID(customerID)
	if err != nil {
		log.Fatal(err)
	}
}

// Menampilkan daftar customer beserta ID dan namanya
func viewCustomersWithIDName() {
	rows, err := db.Query("SELECT customer_id, customer_name FROM Customer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Customer ID\tCustomer Name")
	fmt.Println("-----------\t-------------")

	for rows.Next() {
		var customerID int
		var customerName string
		err := rows.Scan(&customerID, &customerName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\t\t%s\n", customerID, customerName)
	}
}

func printTransactionsByCustomerID(customerID int) error {
	// Ambil detail customer dari database
	var customerName, phoneNumber string
	err := db.QueryRow("SELECT customer_name, phone_number FROM Customer WHERE customer_id = $1", customerID).Scan(&customerName, &phoneNumber)
	if err != nil {
		return err
	}

	// Ambil detail transaction dan service untuk customer yang dipilih dari database menggunakan JOIN
	rows, err := db.Query(`
		SELECT t.transaction_id, s.service_name, t.quantity, s.service_unit, s.price_per_unit, t.entry_date, t.completion_date, t.received_by, t.quantity * s.price_per_unit AS total
		FROM Transaction t
		JOIN Service s ON t.service_id = s.service_id
		WHERE t.customer_id = $1
	`, customerID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var transactions []TransactionData

	// Loop yang digabungkan dan simpan data dalam transaction
	for rows.Next() {
		var transactionID, quantity, pricePerUnit, total int
		var entryDate, completionDate time.Time
		var serviceName, serviceUnit, receivedBy string
		err := rows.Scan(&transactionID, &serviceName, &quantity, &serviceUnit, &pricePerUnit, &entryDate, &completionDate, &receivedBy, &total)
		if err != nil {
			return err
		}

		transaction := TransactionData{
			TransactionID:  transactionID,
			ServiceName:    serviceName,
			Quantity:       quantity,
			ServiceUnit:    serviceUnit,
			PricePerUnit:   pricePerUnit,
			EntryDate:      entryDate,
			CompletionDate: completionDate,
			ReceivedBy:     receivedBy,
			TotalPrice:     total,
		}

		transactions = append(transactions, transaction)
	}

	// Cetak nota
	fmt.Printf("ENIGMA LAUNDRY\t\t\t\t\t\n")
	fmt.Printf("NO\t%d\t\t\t\t\t\t\n", customerID)
	fmt.Printf("TANGGAL MASUK:\t%s\t\tNama Customer\t%s\t\n", transactions[0].EntryDate.Format("2/1/2006"), customerName)
	fmt.Printf("TANGGAL SELESAI:\t%s\t\tNo HP\t\t%s\t\n", transactions[0].CompletionDate.Format("2/1/2006"), phoneNumber)
	fmt.Printf("DITERIMA OLEH:\t%s\t\t\t\t\t\n\n", transactions[0].ReceivedBy)

	fmt.Println("NO\tPELAYANAN\t \tJUMLAH\tSATUAN\tHARGA\tTOTAL")
	fmt.Println("--\t----------\t \t------\t------\t-----\t-----")

	var totalPrice int

	// Cetak setiap data transaction
	for i, transaction := range transactions {
		fmt.Printf("%d\t%s\t%d\t%s\t%d\t%d\n", i+1, transaction.ServiceName, transaction.Quantity, transaction.ServiceUnit, transaction.PricePerUnit, transaction.TotalPrice)
		totalPrice += transaction.TotalPrice
	}

	// Cetak total harga untuk semua transaction
	fmt.Printf("\t\t\t\t\tTOTAL HARGA\t%d\n", totalPrice)

	return nil
}

// Struktur data untuk menyimpan informasi transaction
type TransactionData struct {
	TransactionID  int
	ServiceName    string
	Quantity       int
	ServiceUnit    string
	PricePerUnit   int
	EntryDate      time.Time
	CompletionDate time.Time
	ReceivedBy     string
	TotalPrice     int
}
