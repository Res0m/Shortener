package stat

import "GolangAdvanced/pkg/db"

type StatRepository struct {
	*db.Db
}


func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}



func (repo *StatRepository) AddClick(linkId uint){
	
}