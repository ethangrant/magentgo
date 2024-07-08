package magentgo

import (
	"testing"
)

// TODO: refactor testing process
func TestBuild(t *testing.T) {

	// test case 1
	expected1 := "searchCriteria[filter_groups][0][filters][0][field]=sku&searchCriteria[filter_groups][0][filters][0][value]=WSH%29%&searchCriteria[filter_groups][0][filters][0][condition_type]=like&searchCriteria[filter_groups][0][filters][1][field]=sku&searchCriteria[filter_groups][0][filters][1][value]=WP%29%&searchCriteria[filter_groups][0][filters][1][condition_type]=like&searchCriteria[filter_groups][1][filters][0][field]=price&searchCriteria[filter_groups][1][filters][0][value]=40&searchCriteria[filter_groups][1][filters][0][condition_type]=from&searchCriteria[filter_groups][2][filters][0][field]=price&searchCriteria[filter_groups][2][filters][0][value]=49.99&searchCriteria[filter_groups][2][filters][0][condition_type]=to"
	searchCriteriaBuilder := NewSearchCriteriaBuilder()

	var filters1 []Filter
	filters1 = append(filters1, NewFilter("sku", "WSH%29%", "like"))
	filters1 = append(filters1, NewFilter("sku", "WP%29%", "like"))
	filterGroup1 := NewFilterGroup(filters1)

	var filters2 []Filter
	filters2 = append(filters2, NewFilter("price", "40", "from"))
	filterGroup2 := NewFilterGroup(filters2)

	var filters3 []Filter
	filters3 = append(filters3, NewFilter("price", "49.99", "to"))
	filterGroup3 := NewFilterGroup(filters3)

	searchCriteriaBuilder.addFilterGroup(filterGroup1)
	searchCriteriaBuilder.addFilterGroup(filterGroup2)
	searchCriteriaBuilder.addFilterGroup(filterGroup3)

	actual1 := searchCriteriaBuilder.build()

	// test case 2
	expected2 := "searchCriteria[filter_groups][0][filters][0][field]=sku&searchCriteria[filter_groups][0][filters][0][value]=MGTI&searchCriteria[filter_groups][0][filters][0][condition_type]=eq&searchCriteria[sortOrders][0][field]=price&searchCriteria[sortOrders][0][direction]=asc&searchCriteria[pageSize]=100&searchCriteria[currentPage]=10"
	searchCriteriaBuilder = NewSearchCriteriaBuilder()
	var filters4 []Filter
	filters4 = append(filters4, NewFilter("sku", "MGTI", "eq"))
	filterGroup4 := NewFilterGroup(filters4)
	searchCriteriaBuilder.addFilterGroup(filterGroup4).addSortOrder(NewSortOrder("price", "asc")).setCurrentPage(10).setPageSize(100)

	actual2 := searchCriteriaBuilder.build()

	if actual1 != expected1 {
		t.Errorf("expected searchcriteria %s, got %s", expected1, actual1)
	}
	
	if actual2 != expected2 {
		t.Errorf("expected searchcriteria %s, got %s", expected2, actual2)
	}
}