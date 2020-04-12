package main

// Authors: Sebastian Roy,...
// Date: 19. March 2020


// Run local: go build && ./coronaapp_server
// Run server: GOOS=linux GOARCH=amd64 go build
// scp coronaapp_server SERVERIP:/home/USER/corozone

import (
	"database/sql"
    "database/sql/driver"
	"fmt"
    "time"
    "errors"
	"github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"net/http"
    jwt "github.com/dgrijalva/jwt-go"
//    "crypto/sha1"
    "strings"
    _ "crypto/sha256"
    "github.com/google/uuid"
    gonfig "github.com/eduardbcom/gonfig"
    "encoding/json"
 //   "github.com/lib/pq"

)


//// Consts, Vars, Structs/////
var db *sql.DB
type JwtClaims struct {
    Name string `json:"name"`
    jwt.StandardClaims
}

type user struct {
        UserID   string    `json:"user_id"`
        FirstName string `json:"firstName"`
        LastName string `json:"lastName"`
        Email string `json:"email"`
      //  Password string `json: "password"`
        Mobile string `json:"mobile"`
    }

type location struct {
    Street    string  `json:"street"`
    District  string  `json:"district"`
    City      string  `json:"city"`
    Country   string  `json:"country"`
    Lat       float64   `json:"lat"`
    Long      float64   `json:"long"`
}

type Items struct {
    Name string `json:"name"`
    // Add quantity and Unit maybe later
}


type ItemsSlice []Items

func (i ItemsSlice) Value() (driver.Value, error) {
    return json.Marshal(i)
}

func (i *ItemsSlice) Scan(src interface{}) error {
    switch v := src.(type) {
    case []byte:
        return json.Unmarshal(v, i)
    case string:
        return json.Unmarshal([]byte(v), i)
    }
    return errors.New("type assertion failed")

}



type groceryRequest struct {
    OrderID        string `json:"order_id"`
    CreatedBy      string   `json:"createdBy"`
    Location       location `json:"location"`
    Budget         float32   `json:"budget"`
    ForSomeoneElse bool `json:"forSomeoneElse"`
    InQuarantine   bool `json:"inQuarantine"`
    MinimumSupply  bool `json:"minimumSupply"`
    Elderly        bool `json:"elderly"`
    RequestedItems ItemsSlice `json:"requestedItems"`    // This could be a string array []string
}

func (l location) Value() (driver.Value, error) {
    return json.Marshal(l)
}

func (l *location) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type asserting to []byte failed")
    }
    return json.Unmarshal(b, &l)
}

// TODO: Hospital struct

type DBConfig struct {
    Name string `json:"name"`
    Host string `json:"host"`
    Port int `json:"port"`
    User string`json:"user"`
    Password string`json:"password"`
    DBName string`json:"dbname"`
    SSL string`json:"ssl"`
}

type apikeys struct {
    GFirebase string
    Mapbox string
    GMmaps string
}



var seq = 1
var users = map[int]*user{}

// TODO: Delete this
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Test!")
}


func main() {

dbconf := &DBConfig{}
if rawData, err := gonfig.Read(); err != nil {
    panic(err)
} else {
    if err := json.Unmarshal(rawData, dbconf); err != nil {
        panic(err)
    } else {
		fmt.Printf(
            "{\"name\": \"%s\", \"dbConfig\": {\"host\": \"%s\", port: \"%d\"}}\n",
            dbconf.Name,
            dbconf.Host,
            dbconf.Port ) // {"name": "new-awesome-name", "dbConfig": {"host": "prod-db-server", port: "1"}}
    }
}

    var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbconf.Host, dbconf.Port, dbconf.User, dbconf.Password, dbconf.DBName, dbconf.SSL)
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
    // TODO: pr group should be deleted
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
    userGroup.GET("/:id", getUser)
    userGroup.POST("/createprofile", createUserProfile)
    //userGroup.GET("/login", loginUser)

    groceryGroup := e.Group("/groceries")
    groceryGroup.POST("/create", createGroceryRequest)
    groceryGroup.GET("/getgroceries", getGroceries)

	e.Logger.Fatal(e.Start(":1323"))
}

func createGroceryRequest(c echo.Context) error {
    g := &groceryRequest{
    }
    if err:= c.Bind(g); err != nil {
        return err
    }

    DBCreateGroceryRequest(g)
    return c.JSON(http.StatusCreated, g)
}

func DBCreateGroceryRequest(g *groceryRequest) {
     sqlInsertStatement := fmt.Sprintf(`
    INSERT INTO delivery_order (order_id, request_user_id, budget, for_someone_else, in_quarantine, elderly, requested_Items, geom, location, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, ST_GeomFromText('POINT(%v %v)', 4326), $8, $9) returning order_id`, g.Location.Long, g.Location.Lat)
    var order_id string

    err := db.QueryRow(sqlInsertStatement,
    uuid.New(), g.CreatedBy, g.Budget, g.ForSomeoneElse, g.InQuarantine, g.Elderly, g.RequestedItems, g.Location, time.Now()).Scan(&order_id)
    //_, err := db.Exec(sqlInsertStatement,
    //uuid.New(), g.Budget, g.ForSomeoneElse, g.InQuarantine, g.Elderly, g.RequestedItems, g.Location, time.Now())
    //_, err := db.Exec(sqlInsertStatement, u.FirstName, u.LastName, u.Email, u.Mobile, time.Now())
	if err != nil {
		panic(err)
	}
}

func getGroceries(c echo.Context) error {
   // queryid := c.QueryParam("9ffbc79a-77b2-46a8-b3c4-a0dbed9ffa96")
   // queryid := "9ffbc79a-77b2-46a8-b3c4-a0dbed9ffa96"
   // select order_id, location, first_name, last_name from delivery_order inner join user_profile on delivery_order.request_user_id = user_profile.user_id
	sqlStatement := `SELECT order_id, first_name, location, budget, elderly, for_someone_else, in_quarantine, requested_Items FROM delivery_order INNER JOIN user_profile ON delivery_order.request_user_id = user_profile.user_id ORDER BY delivery_order.created_at DESC limit 30;`
    //var lastname string
	//var user_id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
    rows, err := db.Query(sqlStatement)
    groceryRequests := make([]*groceryRequest, 0)
    //groceryRequest := groceryRequest{}
if err != nil {
    // handle this error better than this
    panic(err)
    //return
}
    // Defer needs to be after error checking, see https://stackoverflow.com/questions/16280176/go-panic-runtime-error-invalid-memory-address-or-nil-pointer-dereference
    defer rows.Close()
    for rows.Next() {

    g := new(groceryRequest)
	switch err := rows.Scan(&g.OrderID, &g.CreatedBy, &g.Location, &g.Budget, &g.Elderly, &g.ForSomeoneElse, &g.InQuarantine, &g.RequestedItems ); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
        return c.String(http.StatusOK, fmt.Sprintf("No user found"))
	case nil:
		//fmt.Println(g.Location)
        //return c.String(http.StatusOK,
        //fmt.Sprintf("ID: %d \n User: %s \n Last Name: %s", u.UserID, u.FirstName, u.LastName))
        groceryRequests = append(groceryRequests, g)
    //    return c.JSON(http.StatusOK, g)
	default:
		panic(err)
        return c.String(http.StatusOK, fmt.Sprintf("Error!"))
	}


}
    return c.JSON(http.StatusOK, groceryRequests)
}

func createUser(c echo.Context) error {
	// User ID from path `users/:id`
    u := &user{
	//    UserID: seq,
	}
	if err := c.Bind(u); err != nil {
	    return err
	}
	//id := c.Param("id")
    DBCreateUser(u)
    return c.JSON(http.StatusCreated, u)
    //return c.String(http.StatusCreated, fmt.Sprintf("User %s %s", u.FirstName, u.LastName))
}

func createUserProfile(c echo.Context) error {
  u := &user{
    }
    if err:= c.Bind(u); err != nil {
        return err
    }

    DBCreateUserProfile(u)
    return c.JSON(http.StatusCreated, u)
}

func DBCreateUserProfile(u *user) {
  var id string
    sqlInsertStatement := `
    INSERT INTO user_profile (user_id, first_name, last_name, email, phone_nr, created_at)
    VALUES ($1, $2, $3, $4, $5, $6) returning user_id`

    err := db.QueryRow(sqlInsertStatement, u.UserID, u.FirstName, u.LastName, u.Email, u.Mobile, time.Now()).Scan(&id)
    //_, err := db.Exec(sqlInsertStatement, u.FirstName, u.LastName, u.Email, u.Mobile, time.Now())
	if err != nil {
		panic(err)
	}


}

func DBCreateUser(u *user) {
    // Creates a user in the DB in table user_profile. With the returned user_id it
    // adds the credentials to and user_auth with a hashed password
    // TODO: Make one out of the two SQL statements
    // TODO: Salt the hashes
    var id string
    sqlInsertStatement := `
    INSERT INTO user_profile (first_name, last_name, email, phone_nr, created_at)
    VALUES ($1, $2, $3, $4, $5) returning user_id`

    err := db.QueryRow(sqlInsertStatement, u.FirstName, u.LastName, u.Email, u.Mobile, time.Now()).Scan(&id)
    //_, err := db.Exec(sqlInsertStatement, u.FirstName, u.LastName, u.Email, u.Mobile, time.Now())
	if err != nil {
		panic(err)
	}

  //  h := sha256.New()
  //  h.Write([]byte(u.Password))
  //  hashsum := h.Sum(nil)

  // sqlPassStatement := `
  //  INSERT INTO user_auth (user_id, pw_hash, pw_salt, hash_algorithm, created_at)
  //  VALUES ($1, $2, $3, $4, $5)`
  //  _, err2 := db.Exec(sqlPassStatement, id, fmt.Sprintf("%x", hashsum), "no salt", "sha256, no salt", time.Now())
//	if err2 != nil {
//		panic(err)
//	}

    // Some server-side reponse when a user is created. Shoud go to logfiles and contain useful information
    fmt.Println(u.FirstName)

}


func getUser(c echo.Context) error {
	// User ID from path `users/:id`
    queryid := c.QueryParam("id")
	sqlStatement := `SELECT user_id, first_name, last_name, email FROM user_profile WHERE user_id=$1;`
	var u user
    //var lastname string
	//var user_id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, queryid)
	switch err := row.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
        return c.String(http.StatusOK, fmt.Sprintf("No user found"))
	case nil:
		fmt.Println(u.UserID,u.FirstName)
        //return c.String(http.StatusOK,
        //fmt.Sprintf("ID: %d \n User: %s \n Last Name: %s", u.UserID, u.FirstName, u.LastName))
        return c.JSON(http.StatusOK, u)
	default:
		panic(err)
        return c.String(http.StatusOK, fmt.Sprintf("Error!"))
	}
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
