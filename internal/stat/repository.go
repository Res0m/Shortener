package stat

import (
	"GolangAdvanced/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	// если нет статистики за сегодня по ссылке - создаем
	//если есть - увеличиваем на 1
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id = ? and date = ?", linkId, datatypes.Date(time.Now()))
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks:  1,
			Date:   currentDate,
		})
	}else{
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}
