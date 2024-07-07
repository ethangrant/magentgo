package magentgo

import (
	"fmt"
	"strings"
)

type SearchCriteria struct {
	filterGroups []FilterGroup
	pageSize int
	currentPage int
}

type FilterGroup struct{
	filters []Filter
}

type Filter struct {
	field string
	value string
	conditionType string
}

type SearchCriteriaBuilder struct{
	searchCriteria *SearchCriteria
}

// TODO: addFilter, AddFilter group etc

func NewFilter(field string, value string, conditionType string) Filter {
	return Filter{
		field: field,
		value: value,
		conditionType: conditionType,
	}
}

func NewFilterGroup(filters []Filter) FilterGroup {
	return FilterGroup{filters: filters}
}

func NewSearchCriteriaBuilder() *SearchCriteriaBuilder {
	return &SearchCriteriaBuilder{searchCriteria: &SearchCriteria{}}
}

func (s *SearchCriteriaBuilder) addFilterGroup(filterGroup FilterGroup) *SearchCriteriaBuilder {
	s.searchCriteria.filterGroups = append(s.searchCriteria.filterGroups, filterGroup)
	return s
}

func (s *SearchCriteriaBuilder) setPageSize(size int) *SearchCriteriaBuilder {
	s.searchCriteria.pageSize = size
	return s
}

func (s *SearchCriteriaBuilder) setCurrentPage(page int) *SearchCriteriaBuilder {
	s.searchCriteria.currentPage = page
	return s
}

// creates a query string of searchCriteria params see https://developer.adobe.com/commerce/webapi/rest/use-rest/performing-searches/
func (s *SearchCriteriaBuilder) build() string {
	var query string

	for index, group := range s.searchCriteria.filterGroups {
		groupString := fmt.Sprintf("searchCriteria[filter_groups][%d]", index)
		for index, filter := range group.filters {
			field := fmt.Sprintf("[filters][%d][field]=%s&", index, filter.field)
			value := fmt.Sprintf("[filters][%d][value]=%s&", index, filter.value)
			conditionType := fmt.Sprintf("[filters][%d][condition_type]=%s&", index, filter.conditionType)

			query += groupString + field;
			query += groupString + value;
			query += groupString + conditionType;
		}
	}

	return strings.TrimSuffix(query, "&")
}