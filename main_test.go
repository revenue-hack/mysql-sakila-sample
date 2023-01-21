package main_test

import (
	"database/sql"
	"fmt"
	"runtime"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func genSetting(username, password, host, port, dbName string) string {
	address := host + ":" + port
	setting := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName)
	return setting
}

func BenchmarkUseIdleConn(b *testing.B) {
	conn, err := sql.Open("mysql", genSetting("root", "root", "localhost", "3306", "employees"))
	if err != nil {
		b.Error(err)
		return
	}
	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(30)
	conn.SetConnMaxLifetime(time.Hour)

	b.Run("test", func(b *testing.B) {
		sem := make(chan struct{}, runtime.NumCPU())
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sem <- struct{}{}
			go func(db *sql.DB) {
				defer func() {
					<-sem
				}()
				if _, err := db.Exec("select sleep(0.01)"); err != nil {
					panic(err)
				}
			}(conn)
		}

	})
	conn.Close()
}

func BenchmarkAppend(b *testing.B) {
	base := []int{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base = append(base, i)
	}
}

func BenchmarkMake(b *testing.B) {
	base := make([]int, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		base[i] = i
	}
}
