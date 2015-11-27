// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterObjectSchemaStmt struct {
	ObjectType ObjectType `json:"objectType"` /* OBJECT_TABLE, OBJECT_TYPE, etc */
	Relation   *RangeVar  `json:"relation"`   /* in case it's a table */
	Object     []Node     `json:"object"`     /* in case it's some other object */
	Objarg     []Node     `json:"objarg"`     /* argument types, if applicable */
	Newschema  *string    `json:"newschema"`  /* the new schema */
	MissingOk  bool       `json:"missing_ok"` /* skip error if missing? */
}

func (node AlterObjectSchemaStmt) MarshalJSON() ([]byte, error) {
	type AlterObjectSchemaStmtMarshalAlias AlterObjectSchemaStmt
	return json.Marshal(map[string]interface{}{
		"ALTEROBJECTSCHEMASTMT": (*AlterObjectSchemaStmtMarshalAlias)(&node),
	})
}

func (node *AlterObjectSchemaStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["objectType"] != nil {
		err = json.Unmarshal(fields["objectType"], &node.ObjectType)
		if err != nil {
			return
		}
	}

	if fields["relation"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["relation"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(RangeVar)
			node.Relation = &val
		}
	}

	if fields["object"] != nil {
		node.Object, err = UnmarshalNodeArrayJSON(fields["object"])
		if err != nil {
			return
		}
	}

	if fields["objarg"] != nil {
		node.Objarg, err = UnmarshalNodeArrayJSON(fields["objarg"])
		if err != nil {
			return
		}
	}

	if fields["newschema"] != nil {
		err = json.Unmarshal(fields["newschema"], &node.Newschema)
		if err != nil {
			return
		}
	}

	if fields["missing_ok"] != nil {
		err = json.Unmarshal(fields["missing_ok"], &node.MissingOk)
		if err != nil {
			return
		}
	}

	return
}
