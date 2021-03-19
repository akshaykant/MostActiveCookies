package search

import "testing"

func TestEvalList_GetListTest(t *testing.T) {
	list := [][]string{
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-09"},
		{"5UAVanZf6UtGyKVS", "2018-12-09"},
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-08"},
		{"4sMM2LxV07bPJzwf", "2018-12-08"},
		{"fbcn5UAVanZf6UtG", "2018-12-08"},
		{"4sMM2LxV07bPJzwf", "2018-12-07"},
	}

	evalList := &EvalList{list, "2018-12-09"}

	enrichedList, err := evalList.GetList()

	if err == nil{
		t.Logf("Enriched List : %s", enrichedList)
	} else {
		t.Fatalf("Error while enriching list : %s", err)
	}
}

func TestNoOccurrenceEvalList_GetListTest(t *testing.T) {
	list := [][]string{
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-09"},
		{"5UAVanZf6UtGyKVS", "2018-12-09"},
		{"AtY0laUfhglK3lC7", "2018-12-09"},
		{"SAZuXPGUrfbcn5UA", "2018-12-08"},
		{"4sMM2LxV07bPJzwf", "2018-12-08"},
		{"fbcn5UAVanZf6UtG", "2018-12-08"},
		{"4sMM2LxV07bPJzwf", "2018-12-07"},
	}

	evalList := &EvalList{list, "2018-12-01"}

	enrichedList, err := evalList.GetList()

	if err != nil{
		t.Logf("No occorrence expected : %s", err)
	} else {
		t.Fatalf("Error : %s", enrichedList)
	}
}



func TestEmptyEvalList_GetListTest(t *testing.T) {
	var list [][]string

	evalList := &EvalList{list, "2018-12-01"}

	enrichedList, err := evalList.GetList()

	if err != nil{
		t.Logf("Empty list expected : %s", err)
	} else {
		t.Fatalf("Error : %s", enrichedList)
	}
}