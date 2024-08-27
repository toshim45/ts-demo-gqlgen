package query

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/toshim45/demo-gqlgen/graph/model"
	"github.com/toshim45/demo-gqlgen/tools"

	"gorm.io/gorm"
)

type Thing struct {
	ID      uuid.UUID `gorm:"primaryKey;column:id"`
	Name    string    `gorm:"column:name"`
	Amount  int       `gorm:"column:amount"`
	Created time.Time `gorm:"column:created"`
}

func whereUUIDExp(db *gorm.DB, w *model.UUIDComparisonExp, col string) *gorm.DB {
	if w == nil {
		return db
	}

	if w.Eq != nil {
		db = db.Where(col+" = ?", *w.Eq)
	}
	return db
}

func whereStringExp(db *gorm.DB, w *model.StringComparisonExp, col string) *gorm.DB {
	if w == nil {
		return db
	}

	if w.Eq != nil {
		db = db.Where(col+" = ?", *w.Eq)
	} else if w.Like != nil {
		db = db.Where(col+" like ?", *w.Like)
	} else if w.Ilike != nil {
		db = db.Where("LOWER("+col+") LIKE LOWER(?)", *w.Ilike)
	}

	return db
}

func whereSmallIntExp(db *gorm.DB, w *model.SmallintComparisonExp, col string) *gorm.DB {
	if w == nil {
		return db
	}

	if w.Eq != nil {
		if n, err := strconv.Atoi(*w.Eq); err == nil {
			db = db.Where(col+" = ?", n)
		}
	} else if w.Neq != nil {
		if n, err := strconv.Atoi(*w.Neq); err == nil {
			db = db.Where(col+" <> ?", n)
		}
	} else if w.Gt != nil {
		if n, err := strconv.Atoi(*w.Gt); err == nil {
			db = db.Where(col+" > ?", n)
		}
	}

	return db
}

func ResolveThing(limit, offset *int, where *model.ThingsBoolExp) ([]*model.Thing, error) {
	var thingList []Thing
	var err error

	db := tools.DB().Debug()

	if limit != nil {
		db = db.Limit(*limit)
	}

	if where != nil {
		db = whereUUIDExp(db, where.ID, "id")
		db = whereStringExp(db, where.Name, "name")
		db = whereSmallIntExp(db, where.Amount, "amount")
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
