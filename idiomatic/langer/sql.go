package langer

var Sql *langer

func init() {
	Sql = new("sql")
}
