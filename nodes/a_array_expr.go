// Auto-generated - DO NOT EDIT

package pg_query

import "encoding/json"

type A_ArrayExpr struct {
	Elements []Node `json:"elements"` /* array element expressions */
	Location int    `json:"location"` /* token location, or -1 if unknown */
}

func (node A_ArrayExpr) MarshalJSON() ([]byte, error) {
	type A_ArrayExprMarshalAlias A_ArrayExpr
	return json.Marshal(map[string]interface{}{
		"A_ARRAYEXPR": (*A_ArrayExprMarshalAlias)(&node),
	})
}

func (node *A_ArrayExpr) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["elements"] != nil {
		node.Elements, err = UnmarshalNodeArrayJSON(fields["elements"])
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

	return
}
