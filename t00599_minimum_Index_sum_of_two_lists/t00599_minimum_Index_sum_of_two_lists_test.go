package t00599_minimum_Index_sum_of_two_lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_t00599_minimum_Index_sum_of_two_lists(t *testing.T) {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

	res := findRestaurant(list1, list2)

	fmt.Printf("res=%+v\n", res)

	assert.ElementsMatch(
		t,
		[]string{"Shogun"},
		res,
	)
}
