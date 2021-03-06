package db

import (
	"github.com/itxor/service/model"
)

// AddRooms saves a list of rooms in Postgres
func (db *PgDB) AddRooms(rooms []*model.Room) error {
	_, err := db.Model(&rooms).
		OnConflict("(room_number) DO UPDATE").
		Insert()

	return err
}
