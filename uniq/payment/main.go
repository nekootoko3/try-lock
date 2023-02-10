package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"try-lock/util"

	_ "github.com/lib/pq"
)

func main() {
	db, billID := prepare()

	fmt.Printf("bill %d に対する支払いトランザクションを開始します\n", billID)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("payments の作成を開始します。")

	// 他トランザクションによって同一 bill_id のレコード作成が行われており、そのトランザクションがコミットされていない場合には他トランザクションがコミットかアボートされるまで処理が中断される
	if _, err := tx.Exec("INSERT INTO payments (bill_id) VALUES ($1)", billID); err != nil {
		log.Fatal(err)
	}

	fmt.Println("payments の作成を完了しました。")

	// 一定時間かかる外部 API 呼び出しを擬似的に行うため標準入力待ちしている
	util.CallExternalAPI()

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("支払いが完了しました。")
}

func prepare() (*sql.DB, int) {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("bill id を指定してください")
	}

	billID, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	db, err := util.ConnectUniqDB()
	if err != nil {
		log.Fatal(err)
	}

	return db, billID
}
