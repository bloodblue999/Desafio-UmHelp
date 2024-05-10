package mysql

import (
	"github.com/bloodblue999/umhelp/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	cli *sqlx.DB

	UserAccount *UserAccount
}

func New(cfg *config.Config) (*Repo, error) {
	url := cfg.MySQLConfig.Username + ":" + cfg.MySQLConfig.Password + "@tcp(" + cfg.MySQLConfig.Host + ":" + cfg.MySQLConfig.Port + ")/" + cfg.MySQLConfig.Database + "?parseTime=true"

	cli, err := sqlx.Connect("mysql", url)
	if err != nil {
		return nil, err
	}

	cli.DB.SetConnMaxLifetime(time.Minute * 5)
	cli.DB.SetMaxIdleConns(5)
	cli.DB.SetMaxOpenConns(100)

	if err := cli.Ping(); err != nil {
		return nil, err
	}

	return &Repo{
		cli:         cli,
		UserAccount: &UserAccount{cli: cli},
	}, nil
}
