package accessor

import (
	"errors"

	"github.com/plattyp/addon/db"
	"github.com/plattyp/addon/resources"
	upper "upper.io/db.v3"
)

// AppAccessor is used to fetch plans
type AppAccessor interface {
	CreateApp(userID int64) (*resources.App, error)
	FetchApp(ID int64) (*resources.App, error)
}

const appTableName = "apps"

// AppDataAccessor is used to interacts with apps
type AppDataAccessor struct {
	Databaser *db.Databaser
}

// CreateApp will create an app based on the params
func (a AppDataAccessor) CreateApp(userID int64) (*resources.App, error) {
	app := resources.App{
		UserID: userID,
	}

	result, err := a.appsTable().Insert(app)
	if err != nil {
		return nil, err
	}

	if id, ok := result.(int64); ok {
		app, err := a.FetchApp(id)
		if err != nil {
			return nil, err
		}
		return app, nil
	}

	return nil, errors.New("Failed to create app")
}

// FetchApp retrieves a app by their ID
func (a AppDataAccessor) FetchApp(id int64) (*resources.App, error) {
	var app resources.App

	err := a.appsTable().Find(upper.Cond{"id": id}).One(&app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

// appsTable returns back a collection
func (a AppDataAccessor) appsTable() upper.Collection {
	return a.Databaser.Conn.Collection(appTableName)
}
