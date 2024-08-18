package db

func InitDB() {
	ConnectDB()
	InitRedis()
	InitUserRepo()
	InitQueryBuilder(GetWriteDB())
	//MigrateDB()
}