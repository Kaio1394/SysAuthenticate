package config

type Config struct {
	App App
	Server
	DataBase
}
type App struct {
	Name string
}

type Server struct {
	Host string
	Port int
}

type DataBase struct {
	TypeDatabase     string
	StringConnection string
}
