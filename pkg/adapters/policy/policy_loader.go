package policy

import (
	"errors"
	"github.com/casbin/casbin/v2"
)

// LoadPolicies loads role-based policies into the Casbin enforcer.
func LoadPolicies(enforcer *casbin.Enforcer) error {
	// Define policies: [role, permission1, permission2, ...]
	policies := [][]string{
		// Superadmin role with full permissions
		{"superadmin", "user:create", "user:view", "user:edit", "user:delete", "survey:create", "survey:manage", "survey:view", "answer:manage", "answer:view", "role:create", "admin:create"},

		// Admin role with permissions for user management and content creation
		{"admin", "user:view", "user:edit", "survey:create", "survey:manage", "survey:view", "answer:manage", "answer:view"},

		// Editor role with content management permissions
		{"editor", "survey:create", "survey:manage", "survey:view", "answer:manage", "answer:view"},

		// User role with profile management permissions
		{"user", "survey:create", "survey:view", "answer:create"},

		// Owner role with permission to manage and view their profile
		{"owner", "survey:start", "survey:finish", "survey:view", "answer:manage", "answer:view", "survey:assignUserPermission"},
	}

	// Add policies to the enforcer
	for _, policy := range policies {
		// Convert each policy (a slice of strings) into a slice of interface{}
		policyInterface := make([]interface{}, len(policy))
		for i, v := range policy {
			policyInterface[i] = v
		}

		// Add policy to enforcer
		_, err := enforcer.AddPolicy(policyInterface...)
		if err != nil {
			return err
		}
	}

	// Optionally, save the policy to persistent storage
	err := enforcer.SavePolicy()
	if err != nil {
		return errors.New("failed to save policies to enforcer")
	}

	return nil
}
