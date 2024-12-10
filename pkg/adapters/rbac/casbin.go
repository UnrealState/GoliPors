package rbac

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"log"
)

type CasbinAdapter struct {
	Enforcer *casbin.Enforcer
}

func NewCasbinAdapter(db *gorm.DB) CasbinAdapter {
	// Initialize a Gorm adapter with Casbin
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("Failed to initialize Casbin adapter: %v", err)
	}

	// Load Casbin model from configuration file
	enforcer, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		log.Fatalf("Failed to initialize Casbin enforcer: %v", err)
	}

	// Load policies from database
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Fatalf("Failed to load Casbin policies: %v", err)
	}

	return CasbinAdapter{
		Enforcer: enforcer,
	}
}
