package orm

import (
	"fmt"
	"gcode/orm/migration"
	log "gcode/utils/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin123"
	dbname   = "gblog"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

func init() {
	// dialect = utils.MustGet("GORM_DIALECT")
	// dsn = utils.MustGet("GORM_CONNECTION_DSN")
	// seedDB = utils.MustGetBool("GORM_SEED_DB")
	// logMode = utils.MustGetBool("GORM_LOGMODE")
	// autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")

	//sslmode 安全验证模式
	pi := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dialect = "postgres"
	dsn = pi
	logMode = true
	autoMigrate = true
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(logMode)
	// Automigrate tables
	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Info("[ORM] Database connection initialized.")
	return orm, err
}
