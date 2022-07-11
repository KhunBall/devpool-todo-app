package users

import (
	db "goapi/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {

	userDB, err := db.All()
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	// user := []User{
	// 	ID:    userDB.ID,
	// 	Name:  userDB.Name,
	// 	Email: userDB.Email,
	// 	Todo:  userDB.Todo,
	// }
	// fmt.Println("hello")
	// fmt.Println(userDB)

	// fmt.Println(reflect.TypeOf(userDB))

	return c.JSON(http.StatusOK, userDB)
}

func GetbyMail(c echo.Context) error {
	email := c.Param("email")

	userDB, err := db.Bymail(email)
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}
	// fmt.Println(email)

	return c.JSON(http.StatusOK, userDB)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	userDB, err := db.Read(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	// user := User{
	// 	ID:    userDB.ID,
	// 	Name:  userDB.Name,
	// 	Email: userDB.Email,
	// 	Todo:  userDB.Todo,
	// }
	return c.JSON(http.StatusOK, userDB)
}

func Save(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// arr := strings.FieldsFunc(user.Email, func(r rune) bool {
	// 	return r == '@'
	// })

	// name := arr[0]

	userDB := db.UserDB{
		Name:  user.Email,
		Email: user.Email,
		Todo:  user.Todo,
	}

	if err := db.Create(userDB); err != nil {
		return c.String(http.StatusExpectationFailed, "create fail")
	}

	return c.JSON(http.StatusCreated, user)
}

// /users/:id
func Update(c echo.Context) error {
	id := c.Param("id")

	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	userDB := db.UserDB{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Todo:  user.Todo,
	}

	if err := db.Update(userDB); err != nil {
		return c.String(http.StatusExpectationFailed, "update fail")
	}

	return c.String(http.StatusOK, "id = "+id)
}

func Delete(c echo.Context) error {
	id := c.Param("id")

	err := db.Delete(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Not Found")
	}
	return c.JSON(http.StatusOK, "Delete "+id+" success")
}
