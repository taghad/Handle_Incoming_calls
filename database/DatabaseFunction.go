package database

//have some changes after November 7
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type manager struct {
	userName string
	password string
	connection *sql.DB
}



//changing database to mysql : TODO
func (e *manager)pathConnection() *sql.DB {
	if e.connection != nil{
		return e.connection
	}
	var err0 error
	e.connection,err0 = e.ConnectDB("manager","Manager@123456")
	if err0 != nil {
		log.Print(err0.Error())
	}
	return e.connection

}

func (e *manager) ConnectDB(user string, password string) (*sql.DB, error) {
	db, err0 := sql.Open("mysql", user+":"+password+"@tcp(localhost)/callDB")
	if err0 != nil {
		return nil,err0
	}

	//create tables
	createUsersTable(db)
	createcallsTable(db)
	return db, nil
}

func createUsersTable(db *sql.DB) {
	st, err0:= db.Prepare("CREATE TABLE IF NOT EXISTS users(ID INTEGER PRIMARY KEY AUTOINCREMENT, userName           TEXT      NOT NULL,role TEXT NOT NULL, state TEXT NOT NULL);")
	//st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER NOT NULL AUTOINCREMENT,userName varchar(255),role varchar(20),state varchar(20),PRIMARY KEY (id))")
	if err0 != nil {
		log.Print(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		log.Print(err1.Error())
	}

}
func InsertNewUser(db *sql.DB,userName string, role string) {

	st, err0 := db.Prepare("insert into users (userName,role , state) values (?,?,?)")
	if err0 != nil {
		log.Print(err0.Error())
	}

	_, err1 := st.Exec(userName, role, "free")
	if err1 != nil {
		log.Print(err1.Error())
	}
	db.Close()

}


func UpdateUserState(db *sql.DB,userName string,state string) {

	st, err0 := db.Prepare("update users SET (state) where (userName) values (?,?)")
	if err0 != nil {
		log.Print(err0.Error())
	}
	_, err1 := st.Exec(state,userName)
	if err1 != nil {
		log.Print(err1.Error())
	}

}

func SelectFreeUsers(db *sql.DB, role string) (userName string) {

	results, err0 := db.Query("select userName from users where state = ? and role = ?","free",role)
	if err0 != nil {
		log.Print(err0.Error()) // proper error handling instead of panic in your app
	}
	hasNext := results.Next()
	if hasNext == true {
		results.Scan(&userName)
		UpdateUserState(db,userName,"busy")
		db.Close()
		return userName
	}
	db.Close()
	return ""


}

func createcallsTable(db *sql.DB)  {
	st, err0:= db.Prepare("CREATE TABLE IF NOT EXISTS calls(ID INTEGER PRIMARY KEY AUTOINCREMENT, phoneNumber           TEXT      NOT NULL);")
	//st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS calls(id INTEGER NOT NULL AUTOINCREMENT,phoneNumber varchar(255),PRIMARY KEY (id))")
	if err0 != nil {
		log.Print(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		log.Print(err1.Error())
	}

}

func InsertNewCall(db *sql.DB,phoneNumber string)  {
	st, err0 := db.Prepare("insert into calls (phoneNumber) values (?)")
	if err0 != nil {
		log.Print(err0.Error())
	}

	_, err1 := st.Exec(phoneNumber)
	if err1 != nil {
		log.Print(err1.Error())
	}

	db.Close()

}

func SelectFirstCall(db *sql.DB) (phoneNumber string) {


	results, err0 := db.Query("select phoneNumber from calls")
	if err0 != nil {
		log.Print(err0.Error()) // proper error handling instead of panic in your app

	}
	hasNext := results.Next()
	if hasNext == true {
		results.Scan(&phoneNumber)
		deleteCall(phoneNumber)
		db.Close()
		return phoneNumber
	}
		db.Close()
		return ""
}

func deleteCall(db *sql.DB,phoneNumber string)  {


	st, err0 := db.Prepare("delete from calls where  phoneNumber = ?")
	if err0 != nil {
		log.Print(err0.Error())
	}

	_, err1 := st.Exec(phoneNumber)
	if err1 != nil {
		log.Print(err1.Error())
	}

}



