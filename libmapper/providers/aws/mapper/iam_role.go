/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package mapper

import (
	"encoding/json"
	"sort"

	"github.com/ernestio/definition-mapper/libmapper/providers/aws/components"
	"github.com/ernestio/definition-mapper/libmapper/providers/aws/definition"
	graph "gopkg.in/r3labs/graph.v2"
)

// MapIamRoles ...
func MapIamRoles(d *definition.Definition) []*components.IamRole {
	var rs []*components.IamRole

	for _, role := range d.IamRoles {
		cr := &components.IamRole{
			Name:                 role.Name,
			Path:                 role.Path,
			Description:          role.Description,
			Policies:             role.Policies,
			AssumePolicyDocument: role.AssumePolicyDocumentRaw,
		}

		if len(role.AssumePolicyDocument) > 0 {
			data, _ := json.Marshal(role.AssumePolicyDocument)
			cr.AssumePolicyDocument = string(data)
		}

		cr.SetDefaultVariables()

		rs = append(rs, cr)
	}

	return rs
}

// MapDefinitionIamRoles : Maps output iam roles into a definition defined iam roles
func MapDefinitionIamRoles(g *graph.Graph) []definition.IamRole {
	var roles []definition.IamRole
	var referenced []string

	for _, c := range g.GetComponents().ByType("iam_instance_profile") {
		profile := c.(*components.IamInstanceProfile)
		referenced = append(referenced, profile.Roles...)
	}

	for _, c := range g.GetComponents().ByType("iam_role") {
		var policyDoc map[string]interface{}

		r := c.(*components.IamRole)

		if sort.SearchStrings(referenced, r.Name) == -1 {
			g.DeleteComponent(c)
			continue
		}
		_ = json.Unmarshal([]byte(r.AssumePolicyDocument), &policyDoc)

		roles = append(roles, definition.IamRole{
			Name:                 r.Name,
			Path:                 r.Path,
			Description:          r.Description,
			Policies:             r.Policies,
			AssumePolicyDocument: policyDoc,
		})
	}

	return roles
}
