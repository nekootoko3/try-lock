package util

import (
	"bufio"
	"fmt"
	"os"
)

func CallExternalAPI() {
	// 一定時間かかる外部 API 呼び出しを擬似的に行うため標準入力待ちしている
	fmt.Println("Enter を押すと支払いを実行します")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
