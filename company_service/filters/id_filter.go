package filters

type idFilter struct {
	baseFilterTemplate
}

func InitIdFilter(existingFilter IFilter, logicalOperator LogicalOperator, conditionalOperator ConditionalOperator, val any) IFilter {
	return &idFilter{
		baseFilterTemplate: baseFilterTemplate{
			existingFilter:      existingFilter,
			val:                 val,
			conditionalOperator: conditionalOperator,
			logicalOperator:     logicalOperator,
			key:                 "_id",
		},
	}
}
