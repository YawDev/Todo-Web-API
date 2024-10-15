package Storage

import (
	"os"
	"strconv"
	models "todo-web-api/Models"
	sqlite "todo-web-api/StorageLite"

	"github.com/joho/godotenv"
)

var UserManager IUserManager
var TaskManager ITaskManager
var ListManager IListManager
var StoreManager IDatabase

func ConfigureDb() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	str := os.Getenv("UseSQLite")
	useSQLite, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}

	if useSQLite {
		Sqlite()
	} else {
		SqlServer()
	}
}

func Sqlite() {
	UserManager = &sqlite.UserStoreLite{}
	TaskManager = &sqlite.TaskStoreLite{}
	ListManager = &sqlite.ListStoreLite{}
	StoreManager = &sqlite.StoreManagerLite{}
}

func SqlServer() {
	UserManager = &UserStore{}
	TaskManager = &TaskStore{}
	ListManager = &ListStore{}
	StoreManager = &StoreDbManager{}
}

type IListManager interface {
	CreateList(list *models.List) (ID int, err error)
	DeleteList(id int) (success bool, err error)
	GetListForUser(id int) (*models.List, error)
	GetList(id int) (*models.List, error)
}

type ITaskManager interface {
	CreateTask(task *models.Task, listId int) (ID int, err error)
	DeleteTask(id int) (success bool, err error)
	GetTask(id int) (*models.Task, error)
	UpdateTask(task *models.Task) (ID int, err error)
}

type IUserManager interface {
	CreateUser(user *models.User) (ID int, err error)
	DeleteUser(id int) (success bool, err error)
	GetUser(id int) (*models.User, error)
	FindExistingAccount(username string, password string) (*models.User, error)
}

type IDatabase interface {
	Connect()
}
