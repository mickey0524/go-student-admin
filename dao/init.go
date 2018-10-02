package dao

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql-driver
)

var DBSession *DBHandler
var dbConfDir = "conf/db.conf"

type DBHandler struct {
	db        *gorm.DB
	optional  *DBOptional
	connected bool
	sync.Mutex
}

func (handler *DBHandler) connectDB() error {
	if handler.connected {
		return nil
	}

	config := handler.optional.generateConfig()
	handler.Lock()
	defer handler.Unlock()

	db, err := gorm.Open(handler.optional.DriverName, config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	handler.connected = true
	handler.db = db
	handler.db.DB().SetMaxIdleConns(handler.optional.MaxIdleConns)
	handler.db.DB().SetMaxOpenConns(handler.optional.MaxOpenConns)
	handler.db.SingularTable(true)
	return nil
}

func (handler *DBHandler) GetConnection() (*gorm.DB, error) {
	err := handler.connectDB()

	if err != nil {
		return nil, err
	}

	return handler.db, nil
}

func (handler *DBHandler) CloseDBConnect() error {
    handler.Lock()
    defer handler.Unlock()

    if handler.connected {
        handler.connected = false
        err :=  handler.db.Close()
        return err
    }

    return nil
}

type DBOptional struct {
	DriverName   string
	Timeout      string
	ReadTimeout  string
	WriteTimeout string
	User         string
	Password     string
	DBName       string
	DBCharset    string
	DBHostname   string
	DBPort       string
	MaxIdleConns int
	MaxOpenConns int
}

/**
 * 构造访问数据库配置，schema：[user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
 */
func (optional *DBOptional) generateConfig() string {
	if optional.DBCharset == "" {
		optional.DBCharset = "utf8"
	}

	format := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s&readTimeout=%s&writeTimeout=%s"
	config := fmt.Sprintf(format, optional.User, optional.Password, optional.DBHostname, optional.DBPort,
		optional.DBName, optional.DBCharset, optional.Timeout, optional.ReadTimeout, optional.WriteTimeout)
	return config
}

func getDefaultDBOptional() DBOptional {
	return DBOptional{
		DriverName:   "mysql",
		Timeout:      "100ms",
		ReadTimeout:  "2.0s",
		WriteTimeout: "5.0s",
		DBHostname:   "localhost",
		DBPort:       "3306",
		DBCharset:    "utf8", // use utf8 as default
		MaxIdleConns: 10,
		MaxOpenConns: 100,
	}
}

func newDBHandlerWithOptional(optional *DBOptional) *DBHandler {
	return &DBHandler{
		connected: false,
		optional:  optional,
	}
}

func getDBConf() (map[string]string, error) {
	confMap := make(map[string]string)

	confStr, err := ioutil.ReadFile(dbConfDir)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	conf := string(confStr)

	for _, line := range strings.Split(conf, "\n") {
		line = strings.Trim(line, "")
		kvs := strings.Fields(line)
		confMap[kvs[0]] = kvs[1]
	}

	return confMap, nil
}

// Init 函数初始化db配置
func Init() {
	confMap, _ := getDBConf()
	DBOpt := getDefaultDBOptional()
	DBOpt.DBHostname = confMap["DBHostname"]
	DBOpt.DBPort = confMap["DBPort"]
	DBOpt.User = confMap["User"]
	DBOpt.Password = confMap["Password"]
	DBOpt.DBName = confMap["DBName"]

	DBSession = newDBHandlerWithOptional(&DBOpt)
	if err := DBSession.connectDB(); err != nil {
		fmt.Println(err)
	}
}
