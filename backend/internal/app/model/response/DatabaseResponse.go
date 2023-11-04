package response

type Database struct {
	Id       int
	Host     *string
	Port     *string
	Username *string
	Password *string
	DbName   *string
	Schema   *string
	Title    *string
	//Organization
}
