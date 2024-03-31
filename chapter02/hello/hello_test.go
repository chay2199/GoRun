package hello

import "testing"

func TestSay(t *testing.T) {
	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "Hello World!",
		},
		{
			result: "Hello Mat, Damon!",
			items: []string{"Mat", "Damon"},
		},
	}

	for _, subtest := range subtests {
		if s := Say(subtest.items); s != subtest.result {
			t.Errorf("Wanted %s (%v), got %s", subtest.result, subtest.items, s)
		} 
	}
}