package db

import(
        _ "github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
        "log"
)

func test() string {
        cnn, err := sql.Open("mysql", "order_system:wei_zhangr@tcp(db:3306)/order_system")
        if err != nil {
                log.Fatal(err)
        }

        id := 1
        var lat string

        if err := cnn.QueryRow("SELECT origin_lat FROM orders WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
                log.Fatal(err)
        }

        fmt.Println(id, lat)

	return lat
}

