package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbPath = "./tmp/blocks"
)

func main() {
	// Өгөгдлийн сангийн холболтыг нээнэ үү
	db, err := sql.Open("mysql", "root:8853d4E!@tcp(localhost:3306)/blockchain")
	if err != nil {
		fmt.Println("Өгөгдлийн санд холбогдож чадсангүй:", err)
		return
	}
	defer db.Close()

	// Холболтыг шалгана уу
	err = db.Ping()
	if err != nil {
		fmt.Println("Өгөгдлийн санг ping хийж чадсангүй:", err)
		return
	}

	fmt.Println("MySQL мэдээллийн санд холбогдсон")

	// Өгөгдлийн сангийн үйлдлийг энд хийнэ.

	//блокчэйнийг эхлүүлж,блок нэмж,блок бүрийн талаарх мэдээллийг хэвлэнэ.
	chain := InitBlockChain()

	chain.AddBlock("2dahi block")
	chain.AddBlock("3dahi block")
	chain.AddBlock("4dahi block")
	chain.AddBlock("5dahi block")
	chain.AddBlock("6dahi block")

	for _, block := range chain.blocks {

		fmt.Printf("Өмнөх блок: %x\n", block.PrevHash)
		fmt.Printf("Блоконд байгаа өгөгдөл: %s\n", block.Data)
		fmt.Printf("Хашлагдсан утга: %x\n", block.Hash)

	}
}
