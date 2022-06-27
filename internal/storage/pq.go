package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/GoSeoTaxi/yandex_go_05/internal/etc"
	"log"
	"strconv"
	"time"
)

var StingConnectPQ string

type IDelPQALL interface {
	IDelPQALLs() (string, error)
}

type TdelPQALL struct {
	CALLDElStr string
}

func (s TdelPQALL) IDelPQALLs() (string, error) {

	duration := time.Second * 2
	time.Sleep(duration)

	if s.CALLDElStr == "Y" {
		fmt.Println(`ALL KILL`)
	}

	rows, err := etc.DB1.Query(`SELECT ID FROM public.shortyp10;`)
	if err != nil {
		fmt.Println(err)
	}

	var id1 int
	var batchSQL string

	for rows.Next() {
		err := rows.Scan(&id1)
		if err != nil {
			log.Fatal(err)
		}
		batchSQL = batchSQL + "UPDATE public.shortyp10 SET is_del = true WHERE  ID=" + strconv.Itoa(id1) + ";"
	}
	if rows.Err() != nil {
		fmt.Println(`ERR rows`)
	}

	_, err = etc.DB1.Exec(batchSQL)
	if err != nil {
		fmt.Println(err)
	}

	Status := "OK"
	return Status, err
}

func DelPQ(link, login string) {

	db, err := sql.Open("postgres", StringConnect)
	if err != nil {
		fmt.Println(`err sql open`)
	}
	defer db.Close()

	//  ;
	fmt.Println(`ВЫПОЛНЯЮ УДАЛЕНИЕ`)
	fmt.Println(link)
	fmt.Println(login)
	fmt.Println(`+++++++++++++++`)

	sqlStatement := `UPDATE public.shortyp10 SET is_del = true WHERE  link=$1 and login=$2;`
	_, err = db.Exec(sqlStatement, link, login)
	if err != nil {
		fmt.Println(err)
	}
}

func PutPQ(link, login, stringConnect string) (int, error) {

	db, err := sql.Open("postgres", StringConnect)
	if err != nil {
		fmt.Println(`err open sql`)
	}
	defer db.Close()

	ctx := context.Background()
	var idLinkOld int
	var idLinkLast int
	var idLink int
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(`err start ctx`)
	}

	if err = tx.QueryRowContext(ctx, "select  last_value(id) over (order by id desc) from shortyp10 limit 1").Scan(&idLinkOld); err != nil {
		if err.Error() == "sql: no rows in result set" {
			fmt.Println(`first query`)

			_, err = db.Exec("INSERT INTO shortyp10 (link, login, is_del) 	VALUES ('localhost', 'anonimus', DEFAULT);")
			if err != nil {
				fmt.Println(err)
			}

			idLinkOld = 1

		} else {
			tx.Rollback()
			fmt.Println("Transaction rollback1!")
		}

	}

	//	ansIns, err = tx.QueryRowContext(ctx, "INSERT INTO shortyp10 (link, login) values ($1, $2) ON CONFLICT DO NOTHING RETURNING id", link, login)
	//	if err = tx.QueryRowContext(ctx, "INSERT INTO shortyp10(link, login) VALUES ($1, $2) ON CONFLICT (link) DO UPDATE SET link=$1 RETURNING id", link, login).Scan(&idLinkLast); err != nil {

	if err = tx.QueryRowContext(ctx, "INSERT INTO shortyp10(link, login, is_del ) VALUES ($1, $2, default) ON CONFLICT (link) DO UPDATE SET link=$1 RETURNING id", link, login).Scan(&idLinkLast); err != nil {
		//	if err = tx.QueryRowContext(ctx, "INSERT INTO shortyp10(link, login) VALUES ($1, $2) ON CONFLICT (link) DO nothing RETURNING id", link, login).Scan(&idLinkLast); err != nil {
		tx.Rollback()
		fmt.Println((err))
		fmt.Println("Transaction rollback12")
		return idLinkLast, err
	}

	//	if err = tx.QueryRowContext(ctx, "select  last_value(id) over () from shortyp10 limit 1").Scan(&idLink); err != nil {

	if err = tx.QueryRowContext(ctx, "select  last_value(id) over (order by id desc) from shortyp10 limit 1").Scan(&idLink); err != nil {
		tx.Rollback()
		fmt.Println((err))
		fmt.Println("Transaction rollback3!")
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println((err))
		fmt.Println("Transaction rollback4!")
	}

	if idLinkLast != idLink || idLinkOld == idLink {
		//	fmt.Println(`КОнфликт`)
		return idLinkLast, fmt.Errorf(etc.ErrNameConlict)
	}

	return idLink, err
}
