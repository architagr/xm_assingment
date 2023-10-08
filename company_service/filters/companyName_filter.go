package filters

type companynameFilter struct {
	baseFilterTemplate
}

func InitCompanynameFilter(existingFilter IFilter, logicalOperator LogicalOperator, conditionalOperator ConditionalOperator, val any) IFilter {
	return &companynameFilter{
		baseFilterTemplate: baseFilterTemplate{
			existingFilter:      existingFilter,
			val:                 val,
			conditionalOperator: conditionalOperator,
			logicalOperator:     logicalOperator,
			key:                 "name",
		},
	}

}
