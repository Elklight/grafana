package state

import (
	"context"
	"time"

	"github.com/grafana/grafana/pkg/services/ngalert/models"
)

// InstanceStore represents the ability to fetch and write alert instances.
type InstanceStore interface {
	FetchOrgIds(ctx context.Context) ([]int64, error)
	ListAlertInstances(ctx context.Context, cmd *models.ListAlertInstancesQuery) error
	SaveAlertInstances(ctx context.Context, cmd ...models.AlertInstance) error
	DeleteAlertInstances(ctx context.Context, keys ...models.AlertInstanceKey) error
	DeleteAlertInstancesByRule(ctx context.Context, key models.AlertRuleKey) error
}

// RuleReader represents the ability to fetch alert rules.
type RuleReader interface {
	ListAlertRules(ctx context.Context, query *models.ListAlertRulesQuery) error
}

// Historian maintains an audit log of alert state history.
type Historian interface {
	RecordState(ctx context.Context, rule *models.AlertRule, currentState *State, evaluatedAt time.Time, currentData, previousData InstanceStateAndReason)
}
