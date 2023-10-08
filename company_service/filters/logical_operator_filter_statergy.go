package filters

import "go.mongodb.org/mongo-driver/bson"

type ILogicalOperatorFilterStatergy interface {
	Build() bson.M
}

type andLogicalOperatorFilterStatergy struct {
	filter bson.M
}

func (statergy *andLogicalOperatorFilterStatergy) Build() bson.M {
	return statergy.filter
}

func InitAndLogicalOperatorFilterStatergy(filters ...bson.M) ILogicalOperatorFilterStatergy {
	return &andLogicalOperatorFilterStatergy{
		filter: bson.M{
			string(AND): filters,
		},
	}
}

type orLogicalOperatorFilterStatergy struct {
	filter bson.M
}

func (statergy *orLogicalOperatorFilterStatergy) Build() bson.M {
	return statergy.filter
}

func InitOrLogicalOperatorFilterStatergy(filters ...bson.M) ILogicalOperatorFilterStatergy {
	return &andLogicalOperatorFilterStatergy{
		filter: bson.M{
			string(OR): filters,
		},
	}
}

func LogicalOperatorFactory(operator LogicalOperator, filter ...bson.M) ILogicalOperatorFilterStatergy {
	switch operator {
	case AND:
		return InitAndLogicalOperatorFilterStatergy(filter...)
	case OR:
		return InitOrLogicalOperatorFilterStatergy(filter...)
	}
	return InitAndLogicalOperatorFilterStatergy(filter...)
}
