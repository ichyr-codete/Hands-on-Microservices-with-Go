package repositories

import (
	"database/sql"
	"log"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-7/video-3/managers-service/src/entities"
	_ "github.com/go-sql-driver/mysql"
)

type MariaDBManagersRepository struct {
	db *sql.DB
}

func NewMariaDBManagersRepository() *MariaDBManagersRepository {

	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "packt:packt@/managers?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	repo := &MariaDBManagersRepository{
		db,
	}

	return repo
}

func (repo *MariaDBManagersRepository) Close() {
	repo.db.Close()
}

func (repo *MariaDBManagersRepository) InsertManagerPlayer(managerID uint32, playerID uint32) error {
	rows, err := repo.db.Query("Insert into manager_player(manager_id, player_id) values(?, ?);", managerID, playerID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (repo *MariaDBManagersRepository) GetManagerByID(managerID uint32) (*entities.Manager, error) {
	m := &entities.Manager{}
	row := repo.db.QueryRow("Select id, manager, account from manager where username=?", managerID)
	err := row.Scan(&m.ID, &m.Name, &m.Account)
	if err != nil {
		return nil, err
	}
	return m, nil
}