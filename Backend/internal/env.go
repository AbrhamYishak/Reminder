package internal
import (
	"os"
	"github.com/joho/godotenv"
	"log"
	"strconv"
)
type env struct{
   JwtKey string
   APIKey string
   Dbname string
   Dbhost string
   Dbpassword string
   Dbport string
   Dbusername string
   Port string
   Host string
   BackupEmail string
   BackupEmailAppPassword string
   EmailPort1 int
   EmailPort2 int
}
var Env env
func Init(){
   if err := godotenv.Load(".env"); err != nil{
	   log.Fatal("could not load the env")
   }
   Env.JwtKey = os.Getenv("secret_key")
   Env.Dbusername = os.Getenv("dbusername")
   Env.Dbpassword = os.Getenv("dbpassword")
   Env.APIKey = os.Getenv("api_key")
   Env.Dbhost = os.Getenv("dbhost")
   Env.Dbport = os.Getenv("dbport")
   Env.Dbname = os.Getenv("dbname") 
   Env.Port = os.Getenv("port")
   Env.Host = os.Getenv("host")
   Env.BackupEmail = os.Getenv("backupemail")
   Env.BackupEmailAppPassword = os.Getenv("backupemailpassword") 
   EmailPort1,err := strconv.Atoi(os.Getenv("emailport1"))
   if err != nil{
	   log.Fatal("could not load port")
   }
   Env.EmailPort1 = EmailPort1
   EmailPort2,err := strconv.Atoi(os.Getenv("emailport2"))
   if err != nil{
	   log.Fatal("could not load port")
   }
   Env.EmailPort2 = EmailPort2
}
