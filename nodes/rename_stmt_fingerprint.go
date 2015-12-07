// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RenameStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RENAMESTMT")
	ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Newname != nil {
		ctx.WriteString(*node.Newname)
	}

	for _, subNode := range node.Objarg {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Object {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.RelationType)))
	ctx.WriteString(strconv.Itoa(int(node.RenameType)))

	if node.Subname != nil {
		ctx.WriteString(*node.Subname)
	}
}
