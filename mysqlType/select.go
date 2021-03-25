package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/127_tms?charset=utf8"
	//链接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库链接失败", err.Error())
	} else {
		fmt.Printf("链接成功\n")

		//DbSelectAll(db)

		//DbSelectFirst(db)

		DbSelectStmt(db)
	}
	db.Close()
}

//多数据查询
func DbSelectAll(db *sql.DB){
	sql1 := "select * from oc_lm_shipment limit ?"
	rows, _ := db.Query(sql1,1)
	for rows.Next() {
		var id int
		var shipment_name string
		var contract_no string
		var country_end_id int
		var pay_time string
		var supplier_id string
		var founder string
		var remarks string
		var status int
		var supplier_founder string
		var supplier_address string
		var created_at string
		var updated_at string
		var delete_status string
		x,_ := rows.Columns()  //获取该表的字段名
		err := rows.Scan(&id,&shipment_name,&contract_no,&country_end_id,&pay_time,&supplier_id,&founder,&remarks,&status,&created_at,&updated_at,&delete_status,&supplier_founder,&supplier_address)
		if err != nil {
			panic("数据异常")
			break
		}
		fmt.Println("x:",x)
		fmt.Println("\nid:",id)
		fmt.Println("shipment_name:",shipment_name)
		fmt.Println("contract_no:",contract_no)
		fmt.Println("country_end_id:",country_end_id)
		fmt.Println("pay_time:",pay_time)
		fmt.Println("supplier_id:",supplier_id)
		fmt.Println("founder:",founder)
		fmt.Println("remarks:",remarks)
		fmt.Println("status:",status)
		fmt.Println("supplier_founder:",supplier_founder)
		fmt.Println("supplier_address:",supplier_address)
	}
}

//单条数据查询
func DbSelectFirst(db *sql.DB){
    sql := "select * from oc_lm_shipment limit ?"
    rows := db.QueryRow(sql,10)
	var id int
	var shipment_name string
	var contract_no string
	var country_end_id int
	var pay_time string
	var supplier_id string
	var founder string
	var remarks string
	var status int
	var supplier_founder string
	var supplier_address string
	var created_at string
	var updated_at string
	var delete_status string
	err := rows.Scan(&id,&shipment_name,&contract_no,&country_end_id,&pay_time,&supplier_id,&founder,&remarks,&status,&created_at,&updated_at,&delete_status,&supplier_founder,&supplier_address)
	if err != nil {
		panic("数据异常")
	}
	fmt.Println("id:",id)
	fmt.Println("shipment_name:",shipment_name)
	fmt.Println("contract_no:",contract_no)
	fmt.Println("country_end_id:",country_end_id)
	fmt.Println("pay_time:",pay_time)
	fmt.Println("supplier_id:",supplier_id)
	fmt.Println("founder:",founder)
	fmt.Println("remarks:",remarks)
	fmt.Println("status:",status)
	fmt.Println("supplier_founder:",supplier_founder)
	fmt.Println("supplier_address:",supplier_address)
}

//stmt查询
func DbSelectStmt(db *sql.DB){
	sql := "select * from oc_lm_shipment limit ?"
	stmt, _ := db.Prepare(sql)
	rows, _ := stmt.Query(10)
	x,_ := rows.Columns()
	for rows.Next() {
		var id int
		var shipment_name string
		var contract_no string
		var country_end_id int
		var pay_time string
		var supplier_id string
		var founder string
		var remarks string
		var status int
		var supplier_founder string
		var supplier_address string
		var created_at string
		var updated_at string
		var delete_status string
		err := rows.Scan(&id,&shipment_name,&contract_no,&country_end_id,&pay_time,&supplier_id,&founder,&remarks,&status,&created_at,&updated_at,&delete_status,&supplier_founder,&supplier_address)
		if err != nil {
			panic("数据异常")
			continue
		}
		fmt.Printf("x:",x)
		fmt.Println("\nid:",id)
		fmt.Println("shipment_name:",shipment_name)
		fmt.Println("contract_no:",contract_no)
		fmt.Println("country_end_id:",country_end_id)
		fmt.Println("pay_time:",pay_time)
		fmt.Println("supplier_id:",supplier_id)
		fmt.Println("founder:",founder)
		fmt.Println("remarks:",remarks)
		fmt.Println("status:",status)
		fmt.Println("supplier_founder:",supplier_founder)
		fmt.Println("supplier_address:",supplier_address)
	}



}