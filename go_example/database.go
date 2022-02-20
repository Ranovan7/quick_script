package main

import (
    "fmt"

    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type User struct {
    gorm.Model
    Username  string
    Role      string
}

func main() {
    fmt.Println("GORM")
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&Product{})

    // Create
    db.Create(&Product{Username: "Kamaboko", Role: "Not-Hashira"})

    // Read
    var product Product
    db.First(&product, 1) // find product with integer primary key
    db.First(&product, "username = ?", "Kamaboko") // find product with code D42

    // Update - update product's price to 200
    db.Model(&product).Update("Role", "Hashira")
    // Update - update multiple fields
    // db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
    // db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

    // Delete - delete product
    db.Delete(&product, 1)
}
