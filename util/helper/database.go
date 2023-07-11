package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover() // overlaping
	if err != nil {
		errorRollback := tx.Rollback()
		PanicError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicError(errorCommit)
	}
}
