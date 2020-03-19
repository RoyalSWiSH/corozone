package main

// Tutorial https://www.youtube.com/watch?v=9VEJyPFz7WY&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY&index=3
// https://www.sohamkamani.com/blog/golang/sql-transactions/
// https://www.restapiexample.com/golang-tutorial/creating-golang-api-echo-framework-postgresql/

// Authors: Sebastian Roy,...
// Date: 19. March 2020

// TODO: Reafactor code, separate files...
// TODO: CI
// TODO: Docker

import (
	"database/sql"
	"fmt"
    "time"
	"github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"net/http"
    jwt "github.com/dgrijalva/jwt-go"
//    "crypto/sha1"
    "strings"
    "crypto/sha256"
)


//// Consts, Vars, Structs/////
var db *sql.DB

type JwtClaims struct {
    Name string `json:"name"`
    jwt.StandardClaims
}
type user struct {
        ID   int    `json:"id"`
        FirstName string `json:"firstName"`
        LastName string `json:"lastName"`
        Email string `json:"email"`
        Password string `json: "password"`
        Mobile string `json:"mobile"`
    }


// TODO: Hospital struct

// TODO: Groceriy struct
// TODO Have a readonly backup SQL server available and one source of truth
// TODO: Have a test system with test data  and one live system
// TODO: Migrate to a mangaged PostreSQL server from digitalocean
const (
	host     = "kandula.db.elephantsql.com"
	port     = 5432
	dbuser     = "lwnxzsyj"
	dbpassword = "kST7gElVkHJA6IWrRuLx9wEDhmGnS0tF"
	dbname   = "lwnxzsyj"
)

var seq = 1
var users = map[int]*user{}

// TODO: Delete this
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Test!")
}
func main() {

    var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbuser, dbpassword, dbname)
    db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
		sqlStatement2 := `UPDATE projects SET name = $2, date = $3 WHERE id = $1;`
	res2, err := db.Exec(sqlStatement2, 2, "NewFirst", "01.01.2020")
	if err != nil {
		panic(err)
	}
	count, err := res2.RowsAffected()
	fmt.Printf("%d rows.\n", count)
e := echo.New()
	e.GET("/", hello)
    pr := e.Group("/projects")
    pr.Use(ServerHeader)
    pr.Use(middleware.Logger())
    pr.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool,error) {
        if username == "jack" && password == "1234" {
            return true, nil;
        }
        return false, nil;
    }))
	pr.GET("/new", newProject)
    pr.GET("/show", showProject)
    pr.GET("/del", delProject)
    cookieGroup := e.Group("/cookie")
    pr.Use(checkCookie)
    cookieGroup.GET("/set", cookieLogin)
    cookieGroup.GET("/main", mainCookie)
    cookieGroup.Use(middleware.Logger())

    userGroup := e.Group("/users")
    userGroup.POST("/create", createUser)
    userGroup.GET("/get", getUser)
    userGroup.GET("/login", loginUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func createUser(c echo.Context) error {
	// User ID from path `users/:id`
    u := &user{
	    ID: seq,
	}
	if err := c.Bind(u); err != nil {
	    return err
	}
	//id := c.Param("id")
    DBCreateUser(u)
    return c.JSON(http.StatusCreated, u)
    //return c.String(http.StatusCreated, fmt.Sprintf("User %s %s", u.FirstName, u.LastName))
}

func DBCreateUser(u *user) {
    // Creates a user in the DB in table user_profile. With the returned user_id it
    // adds the credentials to and user_auth with a hashed password
    // TODO: Make one out of the two SQL statements
    // TODO: Salt the hashes
    var id int
    sqlInsertStatement := `
    INSERT INTO user_profile (first_name, last_name, email, phone_nr, created_at)
    VALUES ($1, $2, $3, $4, $5) returning user_id`

    err := db.QueryRow(sqlInsertStatement, u.FirstName, u.LastName, u.Email, u.Mobile, time.Now()).Scan(&id)
    //_, err := db.Exec(sqlInsertStatement, u.FirstName, u.LastName, u.Email, u.Mobile, time.Now())
	if err != nil {
		panic(err)
	}

    h := sha256.New()
    h.Write([]byte(u.Password))
    hashsum := h.Sum(nil)

   sqlPassStatement := `
    INSERT INTO user_auth (user_id, pw_hash, pw_salt, hash_algorithm, created_at)
    VALUES ($1, $2, $3, $4, $5)`
    _, err2 := db.Exec(sqlPassStatement, id, fmt.Sprintf("%x", hashsum), "no salt", "sha256, no salt", time.Now())
	if err2 != nil {
		panic(err)
	}

    // Some server-side reponse when a user is created. Shoud go to logfiles and contain useful information
    fmt.Println(u.FirstName)

}


func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}


func loginUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// TODO: Remove everything related to projects
func newProject(c echo.Context) error {
projectName := c.QueryParam("name")
sqlStatement := `
    INSERT INTO projects (name, date, author)
    VALUES ($1, $2, $3)`
    _, err := db.Exec(sqlStatement, projectName, "11.02.2020", "Seb")
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, fmt.Sprintf("Neues Projekt %s erstellt", projectName))
}
func showProject(c echo.Context) error {
    queryid := c.QueryParam("id")
	sqlStatement4 := `SELECT id, name FROM projects WHERE id=$1;`
	var name string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement4, queryid)
	switch err := row.Scan(&id, &name); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
        return c.String(http.StatusOK, fmt.Sprintf("Keine Projekte gefunden"))
	case nil:
		fmt.Println(id, name)
        return c.String(http.StatusOK, fmt.Sprintf("ID: %d Projektname: %s ", id, name))
	default:
		panic(err)
        return c.String(http.StatusOK, fmt.Sprintf("Fehler!"))
	}
}

func delProject(c echo.Context) error {
    queryid := c.QueryParam("id")
    var id int
    var name string
    sqlStatement3 := `DELETE FROM projects WHERE id = $1 returning id, name`
    err := db.QueryRow(sqlStatement3, queryid).Scan(&id, &name)
	if err != nil {
        return c.String(http.StatusOK, fmt.Sprintf("Project with id %s could not be deleted.", queryid))
		panic(err)
	}

    return c.String(http.StatusOK, fmt.Sprintf("Deleted Project id %d mit name %s", id, name))
}
/////////// Authentication, Cookies, Login  /////////
func cookieLogin(c echo.Context) error {
    username := c.QueryParam("username")
    password := c.QueryParam("password")

    //h := sha256.New()
    //h.Write([]byte(password))
    //hashsum := h.Sum(nil)
    if username == "seb" && password== "pass" {
        cookie := &http.Cookie{}
        cookie.Name = "sessionID"
        cookie.Value = "hashstring"
        cookie.Expires = time.Now().Add(48 * time.Hour)
        c.SetCookie(cookie)
        token,err := createJwtToken()
        if err != nil {
            fmt.Printf("Error Creating JWT token", err)
            return c.String(http.StatusInternalServerError, "something went wrong")
        }
        return c.JSON(http.StatusOK, map[string]string{
            "message": "you were logged in!",
            "token": token,
        })
       // return c.String(http.StatusOK, fmt.Sprintf("You were logged in! Password is %s, Hash is %x",
        //password, hashsum))
    }
    return c.String(http.StatusOK, "Your username or password were wrong!")
}
func mainCookie(c echo.Context) error {
    return c.String(http.StatusOK, fmt.Sprintf("Cookie"))
}
//////// Middleware

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        c.Response().Header().Set(echo.HeaderServer, "SebServ/1.0")
        c.Response().Header().Set("notreallyheader", "lalal")
        return next(c)
    }
}
    func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            cookie, err := c.Cookie("sessionID")
            if err != nil {
                if strings.Contains(err.Error(), "named cookie not present") {
                    return c.String(http.StatusOK, fmt.Sprintf("Not any Cookie"))
                }
                fmt.Println(err)

                return err
            }
            if cookie.Value == "hashstring" {
                return next(c)
            }
            return c.String(http.StatusOK, fmt.Sprintf("No Cookie"))


        }
    }

// JWT Tokens can be used for Researcher to obtain private data. It can also be used
// for login, but maybe use sessions. i
func createJwtToken() (string, error) {
    claims := JwtClaims{
        "jack",
        jwt.StandardClaims{
            Id: "main_user_id",
            ExpiresAt: time.Now().Add(24*time.Hour).Unix(),
        },
    }
    rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
    token, err := rawToken.SignedString([]byte("mySecret"))
    if err!= nil {
        return "", err
    }
    return token, nil
}
