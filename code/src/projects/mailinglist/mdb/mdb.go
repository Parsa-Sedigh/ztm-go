package mdb

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"log"
	"time"
)

// email records that are in the DB
type EmailEntry struct {
	Id          int64
	Email       string
	ConfirmedAt *time.Time
	OptOut      bool // whether or not they opted out of getting email
}

// a function to make a DB. It will use an existing DB or create a new one if it doesn't exist.
func TryCreate(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE emails (
	    id INTEGER PRIMARY KEY,
	    email TEXT UNIQUE,
	    confirmed_at INTEGER,
	    opt_out INTEGER
	);
`)
	if err != nil {
		/* There are many failure conditions that can happen when we try to run the query above, but we're only interested in one and that's if the failure occurs
		because the table already exists. That would indicate that we already have a DB and we don't need to continue.

		Since this function is supposed to be running during program initialization, we just completely abort the program using log.Fatal() and passing the error mesaage.*/
		if sqlError, ok := err.(sqlite3.Error); ok {
			// code 1 means table already exists
			if sqlError.Code != 1 {
				log.Fatal(sqlError)
			}
		} else {
			log.Fatal(err)
		}

	}
}

// a function that converts query data into a go data structure.
func emailEntryFromRow(row *sql.Rows) (*EmailEntry, error) {
	// create variables for each field in the structure:
	var id int64
	var email string
	var confirmedAt int64
	var optOut bool

	// use pointers to each variable in order to read data into those variables
	// The order of args to Scan should be the same as order of columns in DB
	err := row.Scan(&id, &email, &confirmedAt, &optOut)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	/* We're storing the times in the DB as unix times, which are just integers. So we're gonna convert it into a time.Time struct using Unix() function. */
	t := time.Unix(confirmedAt, 0)

	return &EmailEntry{Id: id, Email: email, ConfirmedAt: &t, OptOut: optOut}, nil
}

func CreateEmail(db *sql.DB, email string) error {
	/* The only column we're concerned with is the email column, so we use ? for it in the query and the passed email value will be substituted for the question mark.
	The value 0 for `confirmed_at` will indicate that the email has not been confirmed and for right now, we default the `opt_out` to false.*/
	_, err := db.Exec(`INSERT INTO 
		emails(email, confirmed_at, opt_out)
		VALUES (?, 0, false)`, email)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil // return nil as error indicating it was ok
}

func GetEmail(db *sql.DB, email string) (*EmailEntry, error) {
	rows, err := db.Query(`
		SELECT id, email,confirmed_at, opt_out
		FROM emails
		WHERE email = ?`, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	/* Since we supplied a unique constraint(email column is unique), so we should ever only get 1 row back from GetEmail function and we return that data
	immediately in this for loop. However, if we don't find anything in the rows(rows would be empty), this for loop won't get executed and we return nil, nil , which means
	we have no data and no error.*/
	for rows.Next() {
		return emailEntryFromRow(rows)
	}

	return nil, nil
}

func UpdateEmail(db *sql.DB, entry EmailEntry) error {
	// our db is storing times as unix times, which is just an integer
	t := entry.ConfirmedAt.Unix()

	// this query is an upsert operation:
	_, err := db.Exec(`
		INSERT INTO emails(email, confirmed_at, opt_out)
		VALUES (?, ?, ?) 
		ON CONFLICT(email) DO UPDATE SET
			confirmed_at=?,
			opt_out=?`, entry.Email, t, entry.OptOut, t, entry.OptOut)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteEmail(db *sql.DB, email string) error {
	/* Normally, when you have a DELETE operation, you would ACTUALLY delete the data out of the DB. However, with this specific app, since we're dealing with a mailing list,
	we want to just set the opt_out to be true as we do here and the reason for this is the user could opt out of emails and if we were to just delete instead and then the email
	would not be in the DB, someone could later add their email again, either on purpose or by accident and then we would continue to send emails to that address and if we continue
	sending emails after they've opted out, that would be considered spamming and then our server could endup on spam lists which is bad!*/

	_, err := db.Exec(`
		UPDATE emails
		SET opt_out=true
		WHERE email=?`, email)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type GetEmailBatchQueryParams struct {
	Page  int
	Count int
}

func GetEmailBatch(db *sql.DB, params GetEmailBatchQueryParams) ([]EmailEntry, error) {
	var empty []EmailEntry

	rows, err := db.Query(`
		SELECT id, email, confirmed_at, opt_out
		FROM emails
		WHERE opt_out = false
		ORDER BY id ASC -- so we don't have one email appear in different pages
		LIMIT ? OFFSET ?`, params.Count, (params.Page-1)*params.Count)

	if err != nil {
		return empty, err
	}

	/* We only need to do the Close operation after we've actually read some data out of the rows. So it's OK to return our error after the checking for error and possibly
	return that error.*/
	defer rows.Close()

	emails := make([]EmailEntry, 0, params.Count)

	for rows.Next() {
		email, err := emailEntryFromRow(rows)
		if err != nil {
			// this function is gonna be all or nothing but error. So we'll never be able to get a partial list of data.
			return nil, err
		}

		emails = append(emails, *email)
	}

	return emails, nil
}
