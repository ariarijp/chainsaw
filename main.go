package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

func createTempTable(fn string, db *sql.DB, tblName string, colName string) {
	sql_ := fmt.Sprintf("CREATE TABLE %s (%s JSON);", tblName, colName)
	_, err := db.Exec(sql_)
	if err != nil {
		log.Fatal(err)
	}

	var fp *os.File
	if fn == "" {
		fp = os.Stdin
	} else {
		fp, err = os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()
	}

	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		l := sc.Text()
		sql_ = fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s');", tblName, colName, l)
		_, err = db.Exec(sql_)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func fetch(rows *sql.Rows, cols *[]string) []interface{} {
	dataPtrs := make([]interface{}, len(*cols))
	data := make([]interface{}, len(*cols))
	for i := range data {
		dataPtrs[i] = &data[i]
	}

	err := rows.Scan(dataPtrs...)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func rowToStrings(data []interface{}) []string {
	result := []string{}

	for _, v := range data {
		switch v2 := v.(type) {
		case int64:
			result = append(result, fmt.Sprintf("%d", v2))
		case float64:
			result = append(result, fmt.Sprintf("%f", v2))
		case []uint8:
			result = append(result, fmt.Sprintf("%s", v2))
		default:
			result = append(result, fmt.Sprint(v2))
		}
	}

	return result
}

func main() {
	var fn, sql_ string

	tblName := flag.String("table", "_", "Table name")
	colName := flag.String("column", "json", "JSON column name")

	flag.Parse()
	if (flag.NArg() == 1) {
		sql_ = flag.Arg(0)
	} else if (flag.NArg() == 2) {
		fn = flag.Arg(0)
		sql_ = flag.Arg(1)
	} else {
		log.Fatal("Argument error")
	}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTempTable(fn, db, *tblName, *colName)

	rows, err := db.Query(sql_)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	tbl := tablewriter.NewWriter(os.Stdout)
	tbl.SetHeader(cols)
	tbl.SetColWidth(80)

	for rows.Next() {
		data := fetch(rows, &cols)
		tbl.Append(rowToStrings(data))
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	tbl.Render()
}
