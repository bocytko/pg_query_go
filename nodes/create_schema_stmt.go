// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateSchemaStmt struct {
	Schemaname  *string `json:"schemaname"`    /* the name of the schema to create */
	Authid      *string `json:"authid"`        /* the owner of the created schema */
	SchemaElts  []Node  `json:"schemaElts"`    /* schema components (list of parsenodes) */
	IfNotExists bool    `json:"if_not_exists"` /* just do nothing if schema already exists? */
}

func (node CreateSchemaStmt) MarshalJSON() ([]byte, error) {
	type CreateSchemaStmtMarshalAlias CreateSchemaStmt
	return json.Marshal(map[string]interface{}{
		"CREATE SCHEMA": (*CreateSchemaStmtMarshalAlias)(&node),
	})
}

func (node *CreateSchemaStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["schemaname"] != nil {
		err = json.Unmarshal(fields["schemaname"], &node.Schemaname)
		if err != nil {
			return
		}
	}

	if fields["authid"] != nil {
		err = json.Unmarshal(fields["authid"], &node.Authid)
		if err != nil {
			return
		}
	}

	if fields["schemaElts"] != nil {
		node.SchemaElts, err = UnmarshalNodeArrayJSON(fields["schemaElts"])
		if err != nil {
			return
		}
	}

	if fields["if_not_exists"] != nil {
		err = json.Unmarshal(fields["if_not_exists"], &node.IfNotExists)
		if err != nil {
			return
		}
	}

	return
}