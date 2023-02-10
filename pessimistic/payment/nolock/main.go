package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"try-lock/util"

	_ "github.com/lib/pq"
)

type Bill struct {
	ID     int
	PaidAt *time.Time
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("bill id を指定してください")
	}

	billID, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	db, err := util.ConnectPessimisticDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bill %d に対する支払いトランザクションを開始します\n", billID)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("bill %d の読み込みを開始します\n", billID)
	bill := &Bill{}
	// 排他的ロックを取得していないので、他トランザクションによって読み書きされる可能性がある
	if err := tx.QueryRow("SELECT id, paid_at FROM bills WHERE id = $1", billID).Scan(&bill.ID, &bill.PaidAt); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bill %d の読み込みが完了しました\n", billID)

	// 支払い済の場合には決済を行わない
	if bill.PaidAt != nil {
		log.Fatal("すでに支払い済みです。")
	}

	// 決済を行ったことを記録するため、bills.paid_at を現在時刻で更新する
	// レコードへの更新を行えることを確認するために外部 API 呼び出し前に更新を行っている
	paidAt := time.Now()
	if _, err := tx.Exec("UPDATE bills SET paid_at = $1 WHERE id = $2", &paidAt, billID); err != nil {
		log.Fatal(err)
	}

	// 一定時間かかる外部 API 呼び出しを擬似的に行うため標準入力待ちしている
	util.CallExternalAPI()

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("支払いが完了しました。支払い時刻は %s です。\n", paidAt.Format("2006-01-02 15:04:05"))
}
