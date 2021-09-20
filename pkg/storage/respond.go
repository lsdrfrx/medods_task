package storage

//* Структура, необходимая для получения информации
//* из базы данных при использовании метода db.Get()
type Respond struct {
	UserId       string `bson:"_id"`
	RefreshToken string `bson:"Refresh-Token"`
}
