// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CreateForeignServerStmt struct {
	Servername *string `json:"servername"` /* server name */
	Servertype *string `json:"servertype"` /* optional server type */
	Version    *string `json:"version"`    /* optional server version */
	Fdwname    *string `json:"fdwname"`    /* FDW name */
	Options    []Node  `json:"options"`    /* generic options to server */
}

func (node CreateForeignServerStmt) MarshalJSON() ([]byte, error) {
	type CreateForeignServerStmtMarshalAlias CreateForeignServerStmt
	return json.Marshal(map[string]interface{}{
		"CREATEFOREIGNSERVERSTMT": (*CreateForeignServerStmtMarshalAlias)(&node),
	})
}

func (node *CreateForeignServerStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["servername"] != nil {
		err = json.Unmarshal(fields["servername"], &node.Servername)
		if err != nil {
			return
		}
	}

	if fields["servertype"] != nil {
		err = json.Unmarshal(fields["servertype"], &node.Servertype)
		if err != nil {
			return
		}
	}

	if fields["version"] != nil {
		err = json.Unmarshal(fields["version"], &node.Version)
		if err != nil {
			return
		}
	}

	if fields["fdwname"] != nil {
		err = json.Unmarshal(fields["fdwname"], &node.Fdwname)
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
