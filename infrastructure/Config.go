package infrastructure

type Config struct {
	DB struct {
		Production struct {
			Host string
			Username string
			Password string
			DBName string
		}
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.DB.Production.Host = "localhost"
	c.DB.Production.Username = "root"
	c.DB.Production.Password = "nuoun1012"
	c.DB.Production.DBName = "sample"

	return c
}