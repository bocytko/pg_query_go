// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterRoleStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ALTERROLESTMT")

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Role != nil {
		ctx.WriteString(*node.Role)
	}
}
