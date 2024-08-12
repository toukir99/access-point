package db

func InitDB() {
	ConnectDB()
	InitQueryBuilder(GetWriteDB())
	//MigrateDB()
}