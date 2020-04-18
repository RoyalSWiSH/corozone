package main

// Authors: Sebastian Roy,...
// Date: 19. March 2020

// Run local: go build && ./coronaapp_server
// Run server: GOOS=linux GOARCH=amd64 go build
// scp coronaapp_server SERVERIP:/home/USER/corozone

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	//    "crypto/sha1"
	"bytes"
	_ "crypto/sha256"
	"encoding/json"
	"io/ioutil"
	_ "net/http/httptrace"
	"net/http/httputil"
	"strings"

	gonfig "github.com/eduardbcom/gonfig"
	"github.com/google/uuid"
	//   "github.com/lib/pq"
)

//// Consts, Vars, Structs/////
var db *sql.DB

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type pushtoken struct {
	FirebasePushtoken string `json:"firebasePushtoken"`
}

type user struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	//  Password string `json: "password"`
	Mobile string `json:"mobile"`
}

type location struct {
	Street   string  `json:"street"`
	District string  `json:"district"`
	City     string  `json:"city"`
	Country  string  `json:"country"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
}

type Items struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	//  Market string `json:"status"`
	//Brand string `json:"status"`
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

type groceryStatus struct {
	OrderID       uuid.UUID `json:"orderID"`
	HelperID      string    `json:"helperID"`
	Status        string    `json:"status"`
	ReceiptAmount float32   `json:"receiptAmount"`
}

type groceryRequest struct {
	OrderID        string     `json:"order_id"`
	CreatedBy      string     `json:"createdBy"`
	Status         string     `json:"status"`
	Location       location   `json:"location"`
	Budget         float32    `json:"budget"`
	ForSomeoneElse bool       `json:"forSomeoneElse"`
	InQuarantine   bool       `json:"inQuarantine"`
	MinimumSupply  bool       `json:"minimumSupply"`
	Elderly        bool       `json:"elderly"`
	RequestedItems ItemsSlice `json:"requestedItems"` // This could be a string array []string
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
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSL      string `json:"ssl"`
}

type apikeys struct {
	GFirebase string
	Mapbox    string
	GMmaps    string
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
				dbconf.Port) // {"name": "new-awesome-name", "dbConfig": {"host": "prod-db-server", port: "1"}}
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
	pr.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "jack" && password == "1234" {
			return true, nil
		}
		return false, nil
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

	apiGroup := e.Group("/api")
	apiv1Group := apiGroup.Group("/v1")

	userGroup := apiv1Group.Group("/users")
	userGroup.POST("/create", createUser)
	userGroup.GET("/:id", getUser)
	userGroup.POST("/createprofile", createUserProfile)
	userGroup.POST("/:user_id/pushtoken", updatePushtoken)
	//userGroup.GET("/login", loginUser)

	groceryGroup := apiv1Group.Group("/groceries")
	groceryGroup.POST("/create", createGroceryRequest)
	groceryGroup.GET("/", getGroceries)
	//Routes authenticate for helper
	groceryGroup.POST("/:id/accept", acceptGroceries)
	groceryGroup.POST("/:id/paid", paidGroceries)
	groceryGroup.POST("/:id/delivered", deliveredGroceries)
	groceryGroup.POST("/:id/repaid", repaidGroceries)
	//Routes authenticate for requester

	// token := DBGetPushtoken("c299a902-9589-42cc-9bfc-5f7221f81364")
	// fmt.Printf(token.FirebasePushtoken)

	e.Logger.Fatal(e.Start(":1323"))
}

func createGroceryRequest(c echo.Context) error {
	g := &groceryRequest{}
	if err := c.Bind(g); err != nil {
		return err
	}

	orderID := dbCreateGroceryRequest(g)

	gstatus := &groceryStatus{
		OrderID: orderID,
		Status:  "open"}

	DBOpenStatusGroceryRequest(gstatus, orderID)

	return c.JSON(http.StatusCreated, g)
}

func dbCreateGroceryRequest(g *groceryRequest) uuid.UUID {
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
	return uuid.MustParse(order_id)
}

func acceptGroceries(c echo.Context) error {
	orderID := c.Param("id")
	guid := uuid.MustParse(orderID)
	g := &groceryStatus{
		OrderID: guid}

	if err := c.Bind(g); err != nil {
		return err
	}
	DBAcceptStatusGroceryRequest(g, guid)

	helperFirstName := DBGetFirstname(g.HelperID)
	NotificationAccepted(guid, g, fmt.Sprintf("%v hat ihren Einkauf angenommen", helperFirstName))

	// TODO: Just return when DBAcceptStatusGroceryRequest has been executed without errors
	return c.String(http.StatusOK, fmt.Sprintf("Accepted grocery request %s", orderID))

}

// NotificationAccepted send a notification via Google Firebase when
// a shopper accepts a request
func NotificationAccepted(guid uuid.UUID, groceryStatus *groceryStatus, message string) {
	// Get token for device of requester based on grocery orderID
	requesterToken := DBGetPushtoken(guid)
	// Get name of helper from his ID

	client := &http.Client{}
	requestBody, err := json.Marshal(map[string]interface{}{
		"notification": map[string]string{
			"title":        "Corozone",
			"body":         message,
			"sound":        "default",
			"click_action": "FCM_PLUGIN_ACTIVITY",
			"icon":         "fcm_push_icon"},
		//  "data":map[string]string{
		//    "from":"codewithabdul.com" },
		"to":       requesterToken.FirebasePushtoken,
		"priority": "high"})

	req, err3 := http.NewRequest("POST",
		"https://fcm.googleapis.com/fcm/send",
		bytes.NewBuffer(requestBody))
	req.Header.Set("Content-type", "application/json")
	// Server token from Google Firebase Console
	// TODO: Put this in config file
	req.Header.Set("Authorization", "key=AAAA81wLwEw:APA91bE_mRbaG79hv7FC8S36ie2YUN3O_Krzayg86uRrFrbWVMuqLgqUg14dHGSziYUqnfwt8kahVNgf11bXK5IV7RFewV8otCoqthFwovJgMNAOnq7ftCv5hFHNKP8QAyyDgBcs4oCt")

	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	} else {
		fmt.Print(string(requestDump))
	}

	if err3 != nil {
		panic(err)
	}
	defer req.Body.Close()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to reach the server.")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("body=", string(body))
		// TODO check if success:1 in response, else return error
	}
	resp.Body.Close()
}

// DBGetPushtoken retreives the token of a users device
// to send messages via GCM
func DBGetPushtoken(guid uuid.UUID) pushtoken {
	sqlStatement := `SELECT user_profile.firebase_pushtoken
    FROM delivery_order
    INNER JOIN user_profile ON delivery_order.request_user_id = user_profile.user_id
    WHERE delivery_order.order_id = $1`

	var token pushtoken
	//var lastname string
	//var user_id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, guid)
	switch err := row.Scan(&token.FirebasePushtoken); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		panic(err)
		// return c.String(http.StatusOK, fmt.Sprintf("No user found"))
	case nil:
		//panic(err)
		//fmt.Print(token.FirebasePushtoken)
		//fmt.Println(u.UserID,u.FirstName)
		//return c.String(http.StatusOK,
		//fmt.Sprintf("ID: %d \n User: %s \n Last Name: %s", u.UserID, u.FirstName, u.LastName))
		//    return c.JSON(http.StatusOK, "T")
	default:
		panic(err)
		//    return c.String(http.StatusOK, fmt.Sprintf("Error!"))
	}
	return token
}
func DBGetFirstname(uid string) string {
	sqlStatement := `SELECT user_profile.first_name
    FROM user_profile
    WHERE user_profile.user_id = $1`

	var firstName string
	//var lastname string
	//var user_id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, uid)
	switch err := row.Scan(&firstName); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		panic(err)
		// return c.String(http.StatusOK, fmt.Sprintf("No user found"))
	case nil:
		//panic(err)
		//fmt.Print(token.FirebasePushtoken)
		//fmt.Println(u.UserID,u.FirstName)
		//return c.String(http.StatusOK,
		//fmt.Sprintf("ID: %d \n User: %s \n Last Name: %s", u.UserID, u.FirstName, u.LastName))
		//    return c.JSON(http.StatusOK, "T")
	default:
		panic(err)
		//    return c.String(http.StatusOK, fmt.Sprintf("Error!"))
	}
	return firstName
}

func paidGroceries(c echo.Context) error {
	orderID := c.Param("id")
	guid := uuid.MustParse(orderID)
	g := &groceryStatus{}

	if err := c.Bind(g); err != nil {
		return err
	}
	DBPaidStatusGroceryRequest(g, guid)

	helperFirstName := DBGetFirstname(g.HelperID)
	NotificationAccepted(guid, g, fmt.Sprintf("%v hat %v EUR für deinen Einkauf bezahlt.", helperFirstName, g.ReceiptAmount))
	// TODO: Just return when DBAcceptStatusGroceryRequest has been executed without errors
	return c.String(http.StatusOK, fmt.Sprintf("Paid groceries %s", orderID))

}

func deliveredGroceries(c echo.Context) error {
	order_id := c.Param("id")
	guid := uuid.MustParse(order_id)
	g := &groceryStatus{}

	if err := c.Bind(g); err != nil {
		return err
	}
	DBDeliveredStatusGroceryRequest(g, guid)

	helperFirstName := DBGetFirstname(g.HelperID)
	NotificationAccepted(guid, g, fmt.Sprintf("%v hat deinen Einkauf (%v EUR) abgeliefert.", helperFirstName, g.ReceiptAmount))
	// TODO: Just return when DBAcceptStatusGroceryRequest has been executed without errors
	return c.String(http.StatusOK, fmt.Sprintf("Delivered groceries %s", order_id))

}

func repaidGroceries(c echo.Context) error {
	order_id := c.Param("id")
	guid := uuid.MustParse(order_id)
	g := &groceryStatus{}

	if err := c.Bind(g); err != nil {
		return err
	}

	DBRepaidStatusGroceryRequest(g, guid)
	// TODO: Just return when DBAcceptStatusGroceryRequest has been executed without errors
	return c.String(http.StatusOK, fmt.Sprintf("Repaid groceries %s", order_id))

}

func DBOpenStatusGroceryRequest(g *groceryStatus, guid uuid.UUID) {
	// DONE: Check first if order_id exists in groceryRequests table
	// TODO: Return HTTP Error when accepted twice
	sqlInsertStatement := `
    INSERT INTO delivery_status (order_id, helper_user_id, status, created_at)
    VALUES ($1, $2, $3, $4) returning order_id`
	var order_id string

	err := db.QueryRow(sqlInsertStatement,
		guid, g.HelperID, "open", time.Now()).Scan(&order_id)
	//_, err := db.Exec(sqlInsertStatement,
	//uuid.New(), g.Budget, g.ForSomeoneElse, g.InQuarantine, g.Elderly, g.RequestedItems, g.Location, time.Now())
	//_, err := db.Exec(sqlInsertStatement, u.FirstName, u.LastName, u.Email, u.Mobile, time.Now())
	if err != nil {
		panic(err)
	}
}

func DBAcceptStatusGroceryRequest(g *groceryStatus, guid uuid.UUID) {
	// DONE: Check first if order_id exists in groceryRequests table
	// TODO: Return HTTP Error when accepted twice by different users
	sqlInsertStatement := `
UPDATE delivery_status
SET status=$1, helper_user_id=$2
WHERE order_id=$3;
`
	//var order_id string

	_, err := db.Query(sqlInsertStatement,
		"accepted", g.HelperID, guid)
	if err != nil {
		panic(err)
	}

}
func DBPaidStatusGroceryRequest(g *groceryStatus, guid uuid.UUID) {
	// DONE: Check first if order_id exists in groceryRequests table
	// TODO: Return HTTP Error when accepted twice by different users
	sqlInsertStatement := `
UPDATE delivery_status
SET status=$1, receipt_amount=$2
WHERE order_id=$3 AND helper_user_id=$4;
`
	//var order_id string

	_, err := db.Query(sqlInsertStatement,
		"paid", g.ReceiptAmount, guid, g.HelperID)
	if err != nil {
		panic(err)
	}

}
func DBDeliveredStatusGroceryRequest(g *groceryStatus, guid uuid.UUID) {
	// DONE: Check first if order_id exists in groceryRequests table
	// TODO: Return HTTP Error when accepted twice by different users
	sqlInsertStatement := `
UPDATE delivery_status
SET status=$1
WHERE order_id=$2 AND helper_user_id=$3;
`
	//var order_id string

	_, err := db.Query(sqlInsertStatement,
		"delivered", guid, g.HelperID)
	if err != nil {
		panic(err)
	}

}
func DBRepaidStatusGroceryRequest(g *groceryStatus, guid uuid.UUID) {
	// DONE: Check first if order_id exists in groceryRequests table
	// TODO: Return HTTP Error when accepted twice by different users
	sqlInsertStatement := `
UPDATE delivery_status
SET status=$1
WHERE order_id=$2 AND helper_user_id=$3;
`
	//var order_id string

	_, err := db.Query(sqlInsertStatement,
		"closed", guid, g.HelperID)
	if err != nil {
		panic(err)
	}

}

func getGroceries(c echo.Context) error {
	// queryid := c.QueryParam("9ffbc79a-77b2-46a8-b3c4-a0dbed9ffa96")
	// queryid := "9ffbc79a-77b2-46a8-b3c4-a0dbed9ffa96"
	// select order_id, location, first_name, last_name from delivery_order inner join user_profile on delivery_order.request_user_id = user_profile.user_id
	sqlStatement := `SELECT delivery_order.order_id, status, first_name, location, budget, elderly, for_someone_else, in_quarantine, requested_Items
    FROM delivery_order
    INNER JOIN user_profile ON delivery_order.request_user_id = user_profile.user_id
    INNER JOIN delivery_status ON delivery_order.order_id = delivery_status.order_id
    WHERE delivery_status.status != 'closed'
    ORDER BY delivery_order.created_at DESC limit 30;`
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
		switch err := rows.Scan(&g.OrderID, &g.Status, &g.CreatedBy, &g.Location, &g.Budget, &g.Elderly, &g.ForSomeoneElse, &g.InQuarantine, &g.RequestedItems); err {
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
	u := &user{}
	if err := c.Bind(u); err != nil {
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
func updatePushtoken(c echo.Context) error {
	// DONE: Check first if order_id exists in groceryRequests table
	// TODO: Return HTTP Error when accepted twice by different users
	uid := c.Param("user_id")
	g := &pushtoken{}
	if err := c.Bind(g); err != nil {
		return err
	}

	sqlInsertStatement := `
UPDATE user_profile
SET firebase_pushtoken=$1
WHERE user_id=$2;
`
	//var order_id string

	_, err := db.Query(sqlInsertStatement,
		g.FirebasePushtoken, uid)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusCreated, "Token updated.")
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
		fmt.Println(u.UserID, u.FirstName)
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
	if username == "seb" && password == "pass" {
		cookie := &http.Cookie{}
		cookie.Name = "sessionID"
		cookie.Value = "hashstring"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)
		token, err := createJwtToken()
		if err != nil {
			fmt.Sprint("Error Creating JWT token $1", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "you were logged in!",
			"token":   token,
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
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
