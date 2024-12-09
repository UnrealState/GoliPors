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
		{"superadmin", "admin:manageUsers", "admin:manageSettings", "dashboard:view", "settings:update", "content:create", "content:update", "content:delete"},

		// Admin role with permissions for user management and content creation
		{"admin", "user:view", "user:edit", "content:create", "content:update"},

		// Editor role with content management permissions
		{"editor", "content:create", "content:update", "content:delete"},

		// User role with profile management permissions
		{"user", "profile:view", "profile:update"},

		// Owner role with permission to manage and view their profile
		{"owner", "owner:manageProfile", "owner:viewProfile"},
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
