package decode

type AdsRulesEngine struct {
	RuleID       string `json:"rule_id"`
	AccountID    string `json:"account_id"`
	ObjectID     string `json:"object_id"`
	ObjectType   string `json:"object_type"`
	TriggerType  string `json:"trigger_type"`
	TriggerField string `json:"trigger_field"`
	CurrentValue string `json:"current_field"`
}
