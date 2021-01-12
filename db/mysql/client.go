package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/kashifkhan0771/HostService/config"
	"github.com/kashifkhan0771/HostService/db"
	domainErr "github.com/kashifkhan0771/HostService/errors"
	"github.com/kashifkhan0771/HostService/models"
)

const (
	hostTableName = "hosts"
)

func init() {
	db.Register("mysql", NewClient)
}

// The first implementation.
type client struct {
	db *sqlx.DB
}

func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

// NewClient initializes a mysql database connection.
func NewClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection: " + formatDSN())
	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}
	return &client{db: cli}, nil
}

func (c *client) AddHost(host *models.Host) (string, error) {
	if host.ID != "" {
		return "", errors.New("id is not empty")
	}

	host.ID = uuid.NewV4().String()
	//jsonStr, err := json.Marshal(host.Metadata)
	//if err == nil{
	//	fmt.Println(string(jsonStr))
	//}
	names := host.Names()
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, hostTableName,
		strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
			return prefix + name
		}), ","))
	if _, err := c.db.NamedExec(query, host); err != nil {
		return "", errors.Wrap(err, "failed to add Host")
	}

	return "", nil
}

func (c *client) GetHost(id string) (*models.Host, error) {
	var stu models.Host
	if err := c.db.Get(&stu, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, hostTableName, id)); err != nil {
		if err == sql.ErrNoRows {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("Host: %s not found", id))
		}
		return nil, err
	}
	return &stu, nil
}

func (c *client) UpdateHost(Host *models.Host) error {
	names := Host.Names()
	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id= '%s'`, hostTableName,
		strings.Join(mkPlaceHolder(names, "=:", func(name, prefix string) string {
			return name + prefix + name
		}), ","), Host.ID)
	log().Info(query)
	if _, err := c.db.NamedExec(query, Host); err != nil {
		return errors.Wrap(err, "failed to update Host")
	}

	return nil
}

func (c *client) DeleteHost(id string) error {
	if _, err := c.db.Query(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, hostTableName, id)); err != nil {
		return errors.Wrap(err, "failed to delete Host")
	}

	return nil
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}

	return ph
}
