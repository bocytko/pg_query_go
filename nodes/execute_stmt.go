// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type ExecuteStmt struct {
	Name   *string `json:"name"`   /* The name of the plan to execute */
	Params []Node  `json:"params"` /* Values to assign to parameters */
}

func (node ExecuteStmt) MarshalJSON() ([]byte, error) {
	type ExecuteStmtMarshalAlias ExecuteStmt
	return json.Marshal(map[string]interface{}{
		"EXECUTESTMT": (*ExecuteStmtMarshalAlias)(&node),
	})
}

func (node *ExecuteStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["name"] != nil {
		err = json.Unmarshal(fields["name"], &node.Name)
		if err != nil {
			return
		}
	}

	if fields["params"] != nil {
		node.Params, err = UnmarshalNodeArrayJSON(fields["params"])
		if err != nil {
			return
		}
	}

	return
}
