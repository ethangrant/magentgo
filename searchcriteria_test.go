package magentgo

import (
	"testing"
)

func TestBuild(t *testing.T) {
	expected := "searchCriteria[filter_groups][0][filters][0][field]=sku&searchCriteria[filter_groups][0][filters][0][value]=WSH%29%&searchCriteria[filter_groups][0][filters][0][condition_type]=like&searchCriteria[filter_groups][0][filters][1][field]=sku&searchCriteria[filter_groups][0][filters][1][value]=WP%29%&searchCriteria[filter_groups][0][filters][1][condition_type]=like&searchCriteria[filter_groups][1][filters][0][field]=price&searchCriteria[filter_groups][1][filters][0][value]=40&searchCriteria[filter_groups][1][filters][0][condition_type]=from&searchCriteria[filter_groups][2][filters][0][field]=price&searchCriteria[filter_groups][2][filters][0][value]=49.99&searchCriteria[filter_groups][2][filters][0][condition_type]=to"
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

	actual := searchCriteriaBuilder.build()

	if actual != expected {
		t.Errorf("expected searchcriteria %s, got %s", expected, actual)
	}
}