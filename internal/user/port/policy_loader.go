package port

import (
	"github.com/casbin/casbin/v2"
)

// LoadPolicies loads role-based policies into the Casbin enforcer.
func LoadPolicies(enforcer *casbin.Enforcer) error {
	// Define policies: [role, resource, action]
	policies := [][]string{
		{"superadmin", "/admin/*", "GET"},
		{"superadmin", "/admin/*", "POST"},
		{"admin", "/users/*", "GET"},
		{"user", "/profile", "GET"},
		{"owner", "/owner/*", "GET"},
	}

	// Add policies to the enforcer
	for _, policy := range policies {
		policyInterface := make([]interface{}, len(policy))
		for i, v := range policy {
			policyInterface[i] = v
		}

		_, err := enforcer.AddPolicy(policyInterface...)
		if err != nil {
			return err
		}
	}

	return enforcer.SavePolicy()
}
