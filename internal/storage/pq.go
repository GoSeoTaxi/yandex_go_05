package storage

import (
	"context"
	"database/sql"
	"fmt"
)

func PutPQ(link, login, stringConnect string) (int, error) {
	db, err := sql.Open("postgres", StringConnect)
	defer db.Close()

	ctx := context.Background()
	var idLinkLast int
	var idLink int
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(`err start ctx`)
	}
	//	ansIns, err = tx.QueryRowContext(ctx, "INSERT INTO shortyp10 (link, login) values ($1, $2) ON CONFLICT DO NOTHING RETURNING id", link, login)

	if err = tx.QueryRowContext(ctx, "INSERT INTO shortyp10(link, login) VALUES ($1, $2) ON CONFLICT (link) DO UPDATE SET link=$1 RETURNING id", link, login).Scan(&idLinkLast); err != nil {
		//	if err = tx.QueryRowContext(ctx, "INSERT INTO shortyp10(link, login) VALUES ($1, $2) ON CONFLICT (link) DO nothing RETURNING id", link, login).Scan(&idLinkLast); err != nil {
		tx.Rollback()
		fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
		return 0, err
	}

	if err = tx.QueryRowContext(ctx, "select  last_value(id) over (order by id desc) from shortyp10 limit 1").Scan(&idLink); err != nil {
		tx.Rollback()
		fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("\n", (err), "\n ....Transaction err!\n")
	} //else {
	//fmt.Println("....Transaction committed\n")
	//}

	if idLinkLast != idLink {
		return idLinkLast, fmt.Errorf("Conflict")
	}

	return idLink, err
}
