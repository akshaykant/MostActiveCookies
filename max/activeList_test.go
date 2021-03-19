package max

import "testing"

func TestEvalList_ActiveCookieTest(t *testing.T){
	list := [][]string{
	{"AtY0laUfhglK3lC7", "2018-12-09"},
	{"SAZuXPGUrfbcn5UA", "2018-12-09"},
	{"5UAVanZf6UtGyKVS", "2018-12-09"},
	{"5UAVanZf6UtGyKVS", "2018-12-09"},
	{"AtY0laUfhglK3lC7", "2018-12-09"},
	}

	evalList := &EvalList{list}

	occurrence, err := evalList.ActiveCookie()

	if err == nil{
		t.Logf("Cookies with maximum occurrence : %s", occurrence)
	} else {
		t.Fatalf("Error while evaluating list : %s", err)
	}
}


func TestEmptyEvalList_ActiveCookieTest(t *testing.T){
	var list [][]string

	evalList := &EvalList{list}

	occurrence, err := evalList.ActiveCookie()

	if err != nil{
		t.Logf("Empty list expected : %s", err)
	} else {
		t.Fatalf("Fail : %s", occurrence)
	}
}