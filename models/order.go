//TODO:
//1. Even Distance should return in float64 type
package model

import (
	"log"
	db "github.com/OrderSystem_WeiZhang/db"
)

type OrderRequest struct {
	Origin      []string `json:"origin"`
	Destination []string `json:"destination"`
}

type OrderRespond struct {
	Id       int    `json:"id"`
	Distance string `json:"distance"`
	Status   string `json:"status"`
}
func (o *OrderRequest) Add( distance string) (id int64, err error){
	rs, err := db.CONN.Exec("INSERT INTO `orders`(`origin_lat`,`origin_long`,`destination_lat`,`destination_long`,`distance`)VALUES(?,?,?,?,?);", o.Origin[0], o.Origin[1],o.Destination[0],o.Destination[1],distance)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

//DB Select
func (o *OrderRespond) GetAll() (orders []OrderRespond, err error) {
	orders = make([]OrderRespond, 0)
	rows, err := db.CONN.Query("SELECT id, distance, `status` FROM orders")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var order OrderRespond
		rows.Scan(&order.Id, &order.Distance, &order.Status)
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//DB Update
func (o *OrderRespond) Update() (num int64, err error) {

	//NOTE: MYSQL default setting is repeatable read
	//NOTE: TRANCTION can keep data consistent
	tx, err := db.CONN.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	rs,err := tx.Exec("Update orders Set `status` = 'TAKEN' WHERE id = ?", o.Id)
	num, _ = rs.RowsAffected()

	if err != nil {
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback() Error:" + err.Error())
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback() Error:" + err.Error())
			return
		}
	}


	return 

}
