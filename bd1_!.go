package main

    import (
        "database/sql"
        "fmt"
        _ "github.com/lib/pq"
		"time"
		
    )

    const (
        DB_USER     = "vtwlajng"
        DB_PASSWORD = "Sty59HjeuNLpFjjhRA5HNok1gHc58lVs"
        DB_NAME     = "vtwlajng"
    )

    func main() {
        dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        checkErr(err)
        defer db.Close()

        fmt.Println("# Inserting values")

        var lastInsertId int
        err = db.QueryRow("INSERT INTO admin(nombre,apellido,cmail) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
        //checkErr(err)
        fmt.Println("last inserted id =", lastInsertId)

        fmt.Println("# Updating")
        stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
        /*checkErr(err)*/

        res, err := stmt.Exec("astaxieupdate", lastInsertId)
        checkErr(err)

        affect, err := res.RowsAffected()
        checkErr(err)

        fmt.Println(affect, "rows changed")

        fmt.Println("# Querying")
        rows, err := db.Query("SELECT * FROM userinfo")
        checkErr(err)

        for rows.Next() {
            var nombre string
            var apellido string
            var mail string
            var created time.Time
            err = rows.Scan( &nombre, &apellido, &mail)
            checkErr(err)
            fmt.Println(" nombre | apellido | mail ")
            fmt.Printf("% %8v | %69v | %30v\n", nombre, apellido, mail, created)
        }

        fmt.Println("# Deleting")
        stmt, err = db.Prepare("delete from userinfo where uid=$1")
        checkErr(err)

        res, err = stmt.Exec(lastInsertId)
        checkErr(err)

        affect, err = res.RowsAffected()
        checkErr(err)

        fmt.Println(affect, "rows changed")
    }

    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }
	