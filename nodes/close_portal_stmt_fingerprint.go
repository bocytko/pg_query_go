// Auto-generated - DO NOT EDIT

package pg_query

func (node ClosePortalStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CLOSEPORTALSTMT")

	if node.Portalname != nil {
		ctx.WriteString(*node.Portalname)
	}
}
