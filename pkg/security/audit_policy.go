package security

type AuditRule struct {
	RuleID   string
	Severity string
}

func GetDefaultAuditRules() []AuditRule {
	return []AuditRule{
		{RuleID: "SEC-001", Severity: "CRITICAL"},
		{RuleID: "SEC-002", Severity: "HIGH"},
	}
}
