// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterUserMappingStmt struct {
	Username   *string `json:"username"`   /* username or PUBLIC/CURRENT_USER */
	Servername *string `json:"servername"` /* server name */
	Options    []Node  `json:"options"`    /* generic options to server */
}

func (node AlterUserMappingStmt) MarshalJSON() ([]byte, error) {
	type AlterUserMappingStmtMarshalAlias AlterUserMappingStmt
	return json.Marshal(map[string]interface{}{
		"ALTERUSERMAPPINGSTMT": (*AlterUserMappingStmtMarshalAlias)(&node),
	})
}

func (node *AlterUserMappingStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["username"] != nil {
		err = json.Unmarshal(fields["username"], &node.Username)
		if err != nil {
			return
		}
	}

	if fields["servername"] != nil {
		err = json.Unmarshal(fields["servername"], &node.Servername)
		if err != nil {
			return
		}
	}

	if fields["options"] != nil {
		node.Options, err = UnmarshalNodeArrayJSON(fields["options"])
		if err != nil {
			return
		}
	}

	return
}
