package query

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/toshim45/demo-gqlgen/graph/model"
	"github.com/toshim45/demo-gqlgen/tools"
)

type Thing struct {
	ID      uuid.UUID `gorm:"primaryKey;column:id"`
	Name    string    `gorm:"column:name"`
	Amount  int       `gorm:"column:amount"`
	Created time.Time `gorm:"column:created"`
}

func ResolveThing(limit, offset *int, where *model.ThingsBoolExp) ([]*model.Thing, error) {
	var thingList []Thing
	var err error

	db := tools.DB().Debug()

	if limit != nil {
		db = db.Limit(*limit)
	}

	if where != nil {
		if where.Name != nil {
			if where.Name.Eq != nil {
				db = db.Where("name = ?", *where.Name.Eq)
			} else if where.Name.Like != nil {
				db = db.Where("name like ?", *where.Name.Like)
			} else if where.Name.Ilike != nil {
				db = db.Where("name ilike ?", *where.Name.Ilike)
			}

		}

		if where.Amount != nil {
			if where.Amount.Eq != nil {
				db = db.Where("amount = ?", *where.Amount.Eq)
			} else if where.Amount.Neq != nil {
				db = db.Where("amount <> ?", *where.Amount.Neq)
			} else if where.Amount.Gt != nil {
				db = db.Where("amount > ?", *where.Amount.Gt)
			}

		}

	}

	err = db.Find(&thingList).Error

	var thingResultList []*model.Thing

	for _, t := range thingList {
		amount := strconv.Itoa(t.Amount)
		tr := model.Thing{
			ID:      t.ID.String(),
			Name:    &t.Name,
			Amount:  &amount,
			Created: t.Created.String(),
		}
		thingResultList = append(thingResultList, &tr)
	}
	return thingResultList, err
}
