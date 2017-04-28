package main

import(
	"fmt"
	log "gopkg.in/inconshreveable/log15.v2"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/jmoiron/sqlx"
)
type DataBase struct{

	Host string
	Port int
	User string
	Password string
	DBname string

	Db *sqlx.DB
}

func NewDataBaseInstance (host string ,port int, user string,  password string,  dbname string) *DataBase{
	return &DataBase{
		Host :host,
		Port:port,
		User:user,
		Password:password,
		DBname:dbname,
		Db: nil,
	}
}

func (dataBase *DataBase) Init () error{

		dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dataBase.Host, dataBase.Port, dataBase.User, dataBase.Password, dataBase.DBname)
		db, err := sqlx.Connect("postgres", dbinfo)
		if(err!=nil){
			log.Error("err", "", err)
			return err
		}else{
			dataBase.Db = db
			log.Info("init", "", dataBase.Db)
			return nil
		}


}

func (dataBase *DataBase) Query (sql string)(*sql.Rows, error){
	if dataBase.Db != nil {
		return dataBase.Db.Query(sql)
	}else{
		return nil, nil
	}
}

func (dataBase *DataBase) QueryStruct (sql string,object interface{}) error{
	if dataBase.Db != nil {
		rows,error:= dataBase.Db.Queryx(sql)
		if(error!=nil){
			return error
		}else{
			for rows.Next() {
				err := rows.StructScan(object)
				if err != nil {
					log.Error("", "", err)
				}
				fmt.Printf("%#v\n", object)

			}
			return nil
		}

	}else{
		return nil
	}
}



