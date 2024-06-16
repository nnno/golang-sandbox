/**
 * -race オプションを付けてテストを実行すると、データ競合が発生していないかを検証できる
 * $ go test 0312.go 0312_test.go
 * $ go test -race 0312.go 0312_test.go
 */
package main

import "testing"

func TestMain(m *testing.M) {
	main()
}
