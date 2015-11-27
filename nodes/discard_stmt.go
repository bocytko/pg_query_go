// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type DiscardStmt struct {
	Target DiscardMode `json:"target"`
}

func (node DiscardStmt) MarshalJSON() ([]byte, error) {
	type DiscardStmtMarshalAlias DiscardStmt
	return json.Marshal(map[string]interface{}{
		"DISCARDSTMT": (*DiscardStmtMarshalAlias)(&node),
	})
}

func (node *DiscardStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["target"] != nil {
		err = json.Unmarshal(fields["target"], &node.Target)
		if err != nil {
			return
		}
	}

	return
}
