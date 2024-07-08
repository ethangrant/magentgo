package magentgo

import (
	"fmt"
	"strings"
)

type SearchCriteria struct {
	filterGroups []FilterGroup
	sortOrders   []SortOrder
	pageSize     int
	currentPage  int
}

type SortOrder struct {
	field     string
	direction string
}

type FilterGroup struct {
	filters []Filter
}

type Filter struct {
	field         string
	value         string
	conditionType string
}

type SearchCriteriaBuilder struct {
	searchCriteria *SearchCriteria
}

func NewSortOrder(field string, direction string) SortOrder {
	return SortOrder{field: field, direction: direction}
}

func NewFilter(field string, value string, conditionType string) Filter {
	return Filter{
		field:         field,
		value:         value,
		conditionType: conditionType,
	}
}

func NewFilterGroup(filters []Filter) FilterGroup {
	return FilterGroup{filters: filters}
}

func NewSearchCriteriaBuilder() *SearchCriteriaBuilder {
	return &SearchCriteriaBuilder{searchCriteria: &SearchCriteria{}}
}

func (s *SearchCriteriaBuilder) AddSortOrder(sortOrder SortOrder) *SearchCriteriaBuilder {
	s.searchCriteria.sortOrders = append(s.searchCriteria.sortOrders, sortOrder)
	return s
}

func (s *SearchCriteriaBuilder) AddFilterGroup(filterGroup FilterGroup) *SearchCriteriaBuilder {
	s.searchCriteria.filterGroups = append(s.searchCriteria.filterGroups, filterGroup)
	return s
}

func (s *SearchCriteriaBuilder) SetPageSize(size int) *SearchCriteriaBuilder {
	s.searchCriteria.pageSize = size
	return s
}

func (s *SearchCriteriaBuilder) SetCurrentPage(page int) *SearchCriteriaBuilder {
	s.searchCriteria.currentPage = page
	return s
}

// creates a query string of searchCriteria params see https://developer.adobe.com/commerce/webapi/rest/use-rest/performing-searches/
func (s *SearchCriteriaBuilder) Build() string {
	var query string

	for index, group := range s.searchCriteria.filterGroups {
		groupString := fmt.Sprintf("searchCriteria[filter_groups][%d]", index)
		for index, filter := range group.filters {
			field := fmt.Sprintf("[filters][%d][field]=%s&", index, filter.field)
			value := fmt.Sprintf("[filters][%d][value]=%s&", index, filter.value)
			conditionType := fmt.Sprintf("[filters][%d][condition_type]=%s&", index, filter.conditionType)

			query += groupString + field
			query += groupString + value
			query += groupString + conditionType
		}
	}

	for index, order := range s.searchCriteria.sortOrders {
		query += fmt.Sprintf("searchCriteria[sortOrders][%d][field]=%s&", index, order.field)
		query += fmt.Sprintf("searchCriteria[sortOrders][%d][direction]=%s&", index, order.direction)
	}

	if s.searchCriteria.pageSize > 0 {
		query += fmt.Sprintf("searchCriteria[pageSize]=%d&", s.searchCriteria.pageSize)
	}

	if s.searchCriteria.currentPage > 0 {
		query += fmt.Sprintf("searchCriteria[currentPage]=%d&", s.searchCriteria.currentPage)
	}

	return strings.TrimSuffix(query, "&")
}
