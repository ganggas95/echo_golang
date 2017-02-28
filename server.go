package main

import (
	"Echo/models"
	"Echo/routes"
	templ "Echo/template"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	GORM      *gorm.DB
	DB_USER   string
	DB_PASS   string
	DB_NAME   string
	DB_HOST   string
	DB_DRIVER string
	Secret    string
)

func init() {
	ConfigDb()
	InitDB()
	OpenRsa()

}

func main() {
	t := &templ.Template{
		Templates: template.Must(template.ParseGlob("views/*")),
	}
	e := echo.New()
	s := &http.Server{
		Addr:         ":210",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Renderer = t
	routes.InitRoute(e, GORM, Secret)
	e.Debug = true
	e.Logger.Fatal(e.StartServer(s))
}
func InitDB() {
	var err error

	GORM, err = gorm.Open(DB_DRIVER, fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME))
	if err != nil {
		fmt.Println(err)
	}
	GORM.SingularTable(true)
	GORM.AutoMigrate(&models.User{})
	GORM.LogMode(true)
	fmt.Println("dbCOnnected")
}

func ConfigDb() {
	viper.SetConfigName("app") // no need to include file extension
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {
		DB_DRIVER = viper.GetString("database.driver")
		DB_HOST = viper.GetString("database.host")
		DB_USER = viper.GetString("database.user")
		DB_PASS = viper.GetString("database.pass")
		DB_NAME = viper.GetString("database.name")
	}
}

func OpenRsa() {
	viper.SetConfigName("app") // no need to include file extension
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	} else {

		Secret = viper.GetString("jwt_config.secret")
	}
}

/*
func readKeyFile(keyPath string) []byte {
	file, err := ioutil.ReadFile(keyPath)
	if err != nil {
		panic(err)
	}
	return file
}

func CreatePem(key []byte) *pem.Block {
	keyData, _ := pem.Decode(key)
	return keyData
}

func OpenRsa() {
	privateKey, err := x509.ParsePKCS1PrivateKey(CreatePem(readKeyFile("app.rsa")).Bytes)
	if err != nil {
		log.Println(err)
	}
	prublicKey, err := x509.ParsePKIXPublicKey(CreatePem(readKeyFile("app.rsa.pub")).Bytes)
	if err != nil {
		log.Println(err)
	}
	rsaPublicKey, es := prublicKey.(*rsa.PublicKey)
	if !es {
		log.Println(err)
	}
	PrivateKey = privateKey
	PublicKey = rsaPublicKey
}

*/
