# Persisting Data

## Database Setup Docker Container

```
docker run --rm -it -d^
   -v %CD%/data:/var/lib/mysql ^
   -p 3306:3306 ^
   -e MYSQL_ROOT_PASSWORD=password123 ^
   mysql:latest
```


## Demo: Connecting to a Database

```
package database

import (
	"database/sql"
	"log"
	"time"
)

var DbConn *sql.DB

// SetupDatabase
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Demo: Querying Data

```
func getProduct(productID int) (*Product, error) {	
	row := database.DbConn.QueryRow( `SELECT 
	productId, 
	manufacturer, 
	sku, 
	upc, 
	pricePerUnit, 
	quantityOnHand, 
	productName 
	FROM products 
	WHERE productId = ?`, productID)

	product := &Product{}
	err := row.Scan(
		&product.ProductID,
		&product.Manufacturer,
		&product.Sku,
		&product.Upc,
		&product.PricePerUnit,
		&product.QuantityOnHand,
		&product.ProductName,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}

func getProductList() ([]Product, error) {
	results, err := database.DbConn.Query(`SELECT 
	productId, 
	manufacturer, 
	sku, 
	upc, 
	pricePerUnit, 
	quantityOnHand, 
	productName 
	FROM products`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	products := make([]Product, 0)
	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)

		products = append(products, product)
	}
	return products, nil
}
```

#### Executing SQL Statements

```
func updateProduct(product Product) error {
	if product.ProductID == nil || *product.ProductID == 0 {
		return errors.New("product has invalid ID")
	}
	_, err := database.DbConn.Exec(`UPDATE products SET 
		manufacturer=?, 
		sku=?, 
		upc=?, 
		pricePerUnit=CAST(? AS DECIMAL(13,2)), 
		quantityOnHand=?, 
		productName=?
		WHERE productId=?`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func insertProduct(product Product) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO products  
	(manufacturer, 
	sku, 
	upc, 
	pricePerUnit, 
	quantityOnHand, 
	productName) VALUES (?, ?, ?, ?, ?, ?)`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertID), nil
}

func removeProduct(productID int) error {
	_, err := database.DbConn.Exec(`DELETE FROM products where productId = ?`, productID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
```