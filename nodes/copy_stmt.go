// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type CopyStmt struct {
	Relation *RangeVar `json:"relation"` /* the relation to copy */
	Query    Node      `json:"query"`    /* the SELECT query to copy */
	Attlist  []Node    `json:"attlist"`  /* List of column names (as Strings), or NIL
	 * for all columns */

	IsFrom    bool    `json:"is_from"`    /* TO or FROM */
	IsProgram bool    `json:"is_program"` /* is 'filename' a program to popen? */
	Filename  *string `json:"filename"`   /* filename, or NULL for STDIN/STDOUT */
	Options   []Node  `json:"options"`    /* List of DefElem nodes */
}

func (node CopyStmt) MarshalJSON() ([]byte, error) {
	type CopyStmtMarshalAlias CopyStmt
	return json.Marshal(map[string]interface{}{
		"COPY": (*CopyStmtMarshalAlias)(&node),
	})
}

func (node *CopyStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
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

	if fields["query"] != nil {
		node.Query, err = UnmarshalNodeJSON(fields["query"])
		if err != nil {
			return
		}
	}

	if fields["attlist"] != nil {
		node.Attlist, err = UnmarshalNodeArrayJSON(fields["attlist"])
		if err != nil {
			return
		}
	}

	if fields["is_from"] != nil {
		err = json.Unmarshal(fields["is_from"], &node.IsFrom)
		if err != nil {
			return
		}
	}

	if fields["is_program"] != nil {
		err = json.Unmarshal(fields["is_program"], &node.IsProgram)
		if err != nil {
			return
		}
	}

	if fields["filename"] != nil {
		err = json.Unmarshal(fields["filename"], &node.Filename)
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
