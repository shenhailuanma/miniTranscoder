package dao

func PrepareJobsTable() error {

	db, err := DatabaseOpen()
	if err != nil {
		return err
	}

	db.AutoMigrate()

	return nil
}
