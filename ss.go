package main

// SystemState is model
type SystemState struct {
	ID int64 `gorm:"primary_key;not null"`
}

// TableName returns name of table
func (ss *SystemState) TableName() string {
	return "system_states"
}

// GetAllSystemStatesIDs is retrieving all system states ids
func GetAllSystemStatesIDs() ([]int64, error) {
	states := new([]SystemState)
	if err := DBConn.Find(&states).Order("id").Error; err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(*states))
	for _, s := range *states {
		ids = append(ids, s.ID)
	}
	return ids, nil
}
