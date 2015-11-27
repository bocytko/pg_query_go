// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DeallocateStmt struct {
	Name *string `json:"name"` /* The name of the plan to remove */

	/* NULL means DEALLOCATE ALL */
}

func (node DeallocateStmt) MarshalJSON() ([]byte, error) {
	type DeallocateStmtMarshalAlias DeallocateStmt
	return json.Marshal(map[string]interface{}{
		"DEALLOCATESTMT": (*DeallocateStmtMarshalAlias)(&node),
	})
}

func (node *DeallocateStmt) UnmarshalJSON(input []byte) (err error) {
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

	return
}
