package infra

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/domain/entity"
)

var (
	Db *gorm.DB
)

func Init() {
	dbSource := getDataSource()
	initDb(dbSource)
	migrateDb()
}

// Read yaml file
func readYml(ymlName string) Yml {
	yml, err := ioutil.ReadFile(ymlName)
	if err != nil {
		panic(err)
	}

	var ymlData Yml
	err = yaml.Unmarshal(yml, &ymlData)
	if err != nil {
		panic(err)
	}

	return ymlData
}

// Get data source
func getDataSource() string {
	var dbSource string

	switch os.Getenv("ENV") {
	case "test":
		yml := readYml("../app-test.yaml")
		dbSource = yml.GetDataSourceUrl()
	default:
		yml := readYml("app.yaml")
		dbSource = yml.GetDataSourceUrl()
	}

	return dbSource
}

// Init database
func initDb(dbSource string) {
	db, err := gorm.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}

	db.DB()
	Db = db
}

// Database migration
func migrateDb() {

	// users
	if !Db.HasTable(entity.User{}.TableName()) {
		Db.CreateTable(&entity.User{})
		hash, _ := bcrypt.GenerateFromPassword([] byte("admin"), bcrypt.DefaultCost)
		user := entity.User{
			Username: "admin",
			Email:    "admin@gmail.com",
			Password: string(hash),
			Uuid:     uuid.NewV4(),
		}
		Db.Create(&user)
	}

	// groups
	if !Db.HasTable(entity.Group{}.GetTableName()) {
		Db.CreateTable(&entity.Group{})
		group := entity.Group{
			Domain: "admin.com",
		}
		Db.Create(&group)
	}

	// principals
	if !Db.HasTable(entity.Principal{}.TableName()) {
		Db.CreateTable(&entity.Principal{})
		principal := entity.Principal{
			UserId:  1,
			GroupId: 1,
		}
		Db.Create(&principal)
	}

	// roles
	if !Db.HasTable(entity.Role{}.TableName()) {
		Db.CreateTable(&entity.Role{})
		role := entity.Role{
			Permission: "admin",
		}
		Db.Create(&role)
	}

	// members
	if !Db.HasTable(entity.Member{}.TableName()) {
		Db.CreateTable(&entity.Member{})
		member := entity.Member{
			UserId: 1,
			RoleId: 1,
		}
		Db.Create(&member)
	}
}
