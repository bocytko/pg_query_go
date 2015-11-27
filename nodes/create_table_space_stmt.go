// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateTableSpaceStmt struct {
	Tablespacename *string `json:"tablespacename"`
	Owner          *string `json:"owner"`
	Location       *string `json:"location"`
	Options        []Node  `json:"options"`
}

func (node CreateTableSpaceStmt) MarshalJSON() ([]byte, error) {
	type CreateTableSpaceStmtMarshalAlias CreateTableSpaceStmt
	return json.Marshal(map[string]interface{}{
		"CREATETABLESPACESTMT": (*CreateTableSpaceStmtMarshalAlias)(&node),
	})
}

func (node *CreateTableSpaceStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["tablespacename"] != nil {
		err = json.Unmarshal(fields["tablespacename"], &node.Tablespacename)
		if err != nil {
			return
		}
	}

	if fields["owner"] != nil {
		err = json.Unmarshal(fields["owner"], &node.Owner)
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
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
