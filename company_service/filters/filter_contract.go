package filters

import "go.mongodb.org/mongo-driver/bson"

type baseFilterTemplate struct {
	filter              bson.M
	key                 string
	conditionalOperator ConditionalOperator
	logicalOperator     LogicalOperator
	val                 any
	existingFilter      IFilter
}

func (base *baseFilterTemplate) Build() bson.M {
	base.filter = bson.M{base.key: conditionalOperatorFactory(base.conditionalOperator, base.val).Build()}

	if base.existingFilter == nil {
		return base.filter
	}

	base.filter = LogicalOperatorFactory(base.logicalOperator, base.existingFilter.Build(), base.filter).Build()
	return base.filter
}

type IFilter interface {
	Build() bson.M
}

type LogicalOperator string
type ConditionalOperator string

const (
	AND LogicalOperator = "$and"
	OR  LogicalOperator = "$or"
)

const (
	NOT_EQUAL ConditionalOperator = "$ne"
	EQUAL     ConditionalOperator = "$eq"
)
