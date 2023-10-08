package filters

import "go.mongodb.org/mongo-driver/bson"

type IConditionalOperatorFilterStatergy interface {
	Build() bson.M
}

type equalConditionalOperatorFilterStatergy struct {
	filter bson.M
}

func (statergy *equalConditionalOperatorFilterStatergy) Build() bson.M {
	return statergy.filter
}

func InitEqualConditionalOperatorFilterStatergy(obj any) IConditionalOperatorFilterStatergy {
	return &equalConditionalOperatorFilterStatergy{
		filter: bson.M{string(EQUAL): obj},
	}
}

type notEqualConditionalOperatorFilterStatergy struct {
	filter bson.M
}

func (statergy *notEqualConditionalOperatorFilterStatergy) Build() bson.M {
	return statergy.filter
}

func InitNotEqualConditionalOperatorFilterStatergy(obj any) IConditionalOperatorFilterStatergy {
	return &equalConditionalOperatorFilterStatergy{
		filter: bson.M{string(NOT_EQUAL): obj},
	}
}

func conditionalOperatorFactory(operator ConditionalOperator, obj any) ILogicalOperatorFilterStatergy {
	switch operator {
	case EQUAL:
		return InitEqualConditionalOperatorFilterStatergy(obj)
	case NOT_EQUAL:
		return InitNotEqualConditionalOperatorFilterStatergy(obj)
	}
	return InitEqualConditionalOperatorFilterStatergy(obj)
}
