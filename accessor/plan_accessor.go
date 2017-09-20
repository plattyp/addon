package accessor

import (
	"github.com/plattyp/addon/db"
	"github.com/plattyp/addon/resources"
	upper "upper.io/db.v3"
)

// PlanAccessor is used to fetch plans
type PlanAccessor interface {
	GetPlanByCode(code string) (*resources.Plan, error)
}

const planTableName = "plans"

// PlanDataAccessor is used to interacts with plans
type PlanDataAccessor struct {
	Databaser *db.Databaser
}

// GetPlanByCode returns the Plan for the given Code if it exists
func (p PlanDataAccessor) GetPlanByCode(code string) (*resources.Plan, error) {
	var review resources.Plan

	err := p.plansTable().Find(upper.Cond{"code": code}).One(&review)
	if err != nil {
		return nil, err
	}

	return &review, nil
}

// plansTable returns back a collection
func (p PlanDataAccessor) plansTable() upper.Collection {
	return p.Databaser.Conn.Collection(planTableName)
}
