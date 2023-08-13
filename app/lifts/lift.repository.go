package lifts

import (
	"database/sql"
	"go-gin-crud-auth/utils/db"
)

type liftRepository struct {
}

func liftMapper(rows *sql.Rows, lift *Lift) error {
	return rows.Scan(&lift.Id, &lift.UserId, &lift.Name, &lift.LiftDate, &lift.Weight, &lift.Reps)
}

func (this *liftRepository) findAll(tx *sql.Tx, userId int) ([]*Lift, error) {
	return db.SelectMultiple[Lift](
		tx,
		liftMapper,
		"SELECT id, user_id, name, lift_date, weight, reps FROM lifts WHERE user_id = ?",
		userId,
	)
}

func (this *liftRepository) findById(tx *sql.Tx, id int, userId int) (*Lift, error) {
	return db.SelectSingle[Lift](
		tx,
		liftMapper,
		"SELECT id, user_id, name, lift_date, weight, reps FROM lifts WHERE id = ? AND user_id =?",
		id,
		userId,
	)
}

func (this *liftRepository) create(tx *sql.Tx, lift *Lift, userId int) error {
	id, error := db.Insert(
		tx,
		"INSERT INTO lifts (user_id, name, lift_date, weight, reps) VALUES (?, ?, ?, ?, ?)",
		userId, lift.Name, lift.LiftDate, lift.Weight, lift.Reps,
	)
	lift.Id = id
	lift.UserId = userId

	return error
}

func (this *liftRepository) update(tx *sql.Tx, lift *Lift, userId int) error {
	error := db.Update(
		tx,
		"UPDATE lifts SET name = ?, lift_date = ?, weight = ?, reps = ? WHERE id = ? AND user_id = ?",
		lift.Name, lift.LiftDate, lift.Weight, lift.Reps, lift.Id, userId,
	)

	return error
}

func (this *liftRepository) delete(tx *sql.Tx, liftId int, userId int) error {
	error := db.Delete(
		tx,
		"DELETE FROM lifts WHERE id = ? AND user_id = ?",
		liftId,
		userId,
	)

	return error
}

var LiftRepository = liftRepository{}
