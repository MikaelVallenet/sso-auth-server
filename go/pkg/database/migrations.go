package database

import "gorm.io/gorm"

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(All()...)
	if err != nil {
		return err
	}
	return nil
}

func AllMap() map[string]interface{} {
	return map[string]interface{}{
		"User": User{},
	}
}

func All() []interface{} {
	out := []interface{}{}
	for _, v := range AllMap() {
		out = append(out, v)
	}
	return out
}
