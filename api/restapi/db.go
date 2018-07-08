package restapi

import (
	"github.com/opus47/cloud/api/models"

	"database/sql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"log"
)

const connStr string = "host=db user=postgres dbname=opus47 sslmode=disable"

func connect() (*sql.DB, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("pg-connect error: %v", err)
		return nil, err
	}

	return db, nil

}

func getKey(name string) (*models.Key, errort) {

	// connect to the database
	db, err := connect()
	if err != nil {
		return nil, SysError("%v", err)
	}
	defer db.Close()

	k := &models.Key{}

	q := `
		SELECT id, name
		FROM keys
		WHERE $1 ILIKE name
	`

	stmt, err := db.Prepare(q)
	if err != nil {
		return nil, SysError("getKey prepare select error: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&k.ID, &k.Name)
	if err != nil {
		return nil, SysError("getKey scan error: %v", err)
	}

	return k, nil

}

func getComposer(tx *sql.Tx, first, last string, create bool) (*models.Composer, errort) {

	c := &models.Composer{}

	// first try to find the composer if it already exists

	q := `
		SELECT id, first, last
		FROM composers
		WHERE
			first ILIKE $1 AND last ILIKE $2
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return nil, SysError("getComposer prepare select error: %v", err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(first, last).Scan(&c.ID, &c.First, &c.Last)
	switch {
	case err == sql.ErrNoRows:
		if create {

			q := `
				INSERT INTO composers (first, last)
				values ($1, $2)
				RETURNING id
			`

			stmt, err := tx.Prepare(q)
			if err != nil {
				return nil, SysError("getCompser prepare insert error: %v", err)
			}
			defer stmt.Close()

			err = stmt.QueryRow(first, last).Scan(&c.ID)
			if err != nil {
				return nil, SysError("getComposer exec insert error: %v", err)
			}

			c.First = first
			c.Last = last

			return c, nil
		} else {
			return nil, nil
		}
	case err != nil:
		return nil, SysError("getComposer exec select error: %v", err)
	default:
		return c, nil
	}

}

func addPiece(piece *models.Piece) errort {

	// connect to the database
	db, err := connect()
	if err != nil {
		return SysError("addPiece: db connect error: %v", err)
	}
	defer db.Close()

	// start up a transaction
	tx, err := db.Begin()
	if err != nil {
		return SysError("addPiece: failed to create transaction: %v", err)
	}

	// Get the composer information, adding them if they do not exist
	composer, errt := getComposer(tx, piece.Composer.First, piece.Composer.Last, true)
	if err != nil {
		return errt
	}
	piece.Composer = composer

	// Get the key information
	key, errt := getKey(piece.Key.Name)
	if err != nil {
		return errt
	}
	piece.Key = key
	if key == nil {
		return UserError("unkown key")
	}

	// add the piece data

	q := `
		INSERT INTO pieces
		(composer, title, key, number, catalog)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	stmt, err := tx.Prepare(q)
	if err != nil {
		return SysError("addPiece prepare piece insert error: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		piece.Composer.ID,
		piece.Title,
		piece.Key.ID,
		piece.Number,
		piece.Catalog,
	).Scan(&piece.ID)
	if err != nil {
		pq_err := err.(*pq.Error)
		if pq_err.Code == "23505" {
			return UserError("duplicate piece")
		} else {
			return SysError("addPiece exec piece insert error: %#v", err)
		}
	}

	// add the movements

	for _, x := range piece.Movements {

		q := `
			INSERT INTO movements
			(piece, title, number)
			VALUES ($1, $2, $3)
		`

		stmt, err := tx.Prepare(q)
		if err != nil {
			return SysError("addPiece prepare insert movement error: %v", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(piece.ID, x.Title, x.Number)

		if err != nil {
			return SysError("add piece exec insert movement error: %v", err)
		}

	}

	// add the parts

	for _, x := range piece.Parts {

		// resolve the part id
		q := `
			SELECT id FROM parts
			WHERE name ILIKE $1
		`

		stmt, err := tx.Prepare(q)
		if err != nil {
			return SysError("addPiece prepare part query error: %v", err)
		}
		defer stmt.Close()

		err = stmt.QueryRow(x.Name).Scan(&x.ID)
		if err != nil {
			return SysError("add piece exec part query error: %v", err)
		}

		// insert the piece part
		q = `
			INSERT INTO piece_parts
			(piece, part)
			VALUES ($1, $2)
		`

		stmt, err = tx.Prepare(q)
		if err != nil {
			return SysError("addPiece prepare part insert error: %v", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(piece.ID, x.ID)

		if err != nil {
			return SysError("addPiece exec part insert error: %v", err)
		}

	}

	_, err = tx.Exec("REFRESH MATERIALIZED VIEW mv_pieces")
	if err != nil {
		return SysError("addPiece error refreshing mv_pieces: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return SysError("addPiece tx fail: %v", err)
	}

	return nil

}

func updatePiece(piece *models.Piece) errort {

	// connect to the database
	db, err := connect()
	if err != nil {
		return SysError("updatePiece: db connect error: %v", err)
	}
	defer db.Close()

	// start up a transaction
	tx, err := db.Begin()
	if err != nil {
		return SysError("addPiece: failed to create transaction: %v", err)
	}

	// Get the composer information, adding them if they do not exist
	composer, errt := getComposer(tx, piece.Composer.First, piece.Composer.Last, true)
	if err != nil {
		return errt
	}
	piece.Composer = composer

	// Get the key information
	key, errt := getKey(piece.Key.Name)
	if err != nil {
		return errt
	}
	piece.Key = key
	if key == nil {
		return UserError("unkown key")
	}

	// update the basic piece information
	q := `
		UPDATE pieces SET
		composer = $1,
		title = $2,
		key = $3,
		number = $4,
		catalog = $5
		WHERE id = $6
	`

	stmt, err := tx.Prepare(q)
	if err != nil {
		return SysError("update piece prepare piece insert error: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		piece.Composer.ID,
		piece.Title,
		piece.Key.ID,
		piece.Number,
		piece.Catalog,
		piece.ID,
	)
	if err != nil {
		return SysError("updatePiece exec piece error: %#v", err)
	}

	// update the movements

	for _, x := range piece.Movements {

		/*
			q := `
				UPDATE movements SET
				title = $1
				WHERE piece = $2 AND number = $3 AND title != $1
			`
		*/

		q := `
			INSERT INTO movements
			(piece, title, number)
			VALUES ($1, $2, $3)
			ON CONFLICT (piece, number) DO UPDATE SET title = $2
		`

		log.Printf("update %s %d", x.Title, x.Number)

		stmt, err := tx.Prepare(q)
		if err != nil {
			return SysError("updatePiece prepare update movement error: %v", err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(piece.ID, x.Title, x.Number)
		if err != nil {
			return SysError("updatePiece: update movement error: %v", err)
		}

	}

	// update the parts

	// clear out existing parts
	q = `DELETE FROM piece_parts WHERE piece = $1`
	stmt, err = tx.Prepare(q)
	if err != nil {
		return SysError("updatePiece: prepare update part failed: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(piece.ID)
	if err != nil {
		return SysError("updatePiece: error clearing out old parts: %v", err)
	}
	stmt.Close()

	for _, x := range piece.Parts {

		// get part ids
		q = `SELECT id FROM parts WHERE name ILIKE $1`
		stmt, err = tx.Prepare(q)
		if err != nil {
			return SysError("updatePiece: prepare update part failed: %v", err)
		}
		defer stmt.Close()
		err = stmt.QueryRow(x.Name).Scan(&x.ID)
		if err != nil {
			return SysError("updatePiece: exec part query error: %v", err)
		}
		stmt.Close()

		// insert the piece part
		q = `INSERT INTO piece_parts (piece, part) VALUES ($1, $2)`
		stmt, err = tx.Prepare(q)
		if err != nil {
			return SysError("updatePiece: prepare part insert error: %v", err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(piece.ID, x.ID)
		if err != nil {
			return SysError("updatePiece: exec part insert error: %v", err)
		}

	}

	_, err = tx.Exec("REFRESH MATERIALIZED VIEW mv_pieces")
	if err != nil {
		return SysError("updatePiece: error refreshing mv_pieces: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return SysError("updatePiece tx fail: %v", err)
	}

	return nil
}
