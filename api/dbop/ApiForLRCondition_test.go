package dbop

import "testing"

func TestInsertLRConditionByCom(t *testing.T) {
	roomCondition, err := InsertLRConditionByCom("22e106ca-ff5d-4876-be231d27-80900799","",1,2,1,1,20.00)
	if err != nil {
		t.Errorf("Error of insert LRcondition:%v", err)
	}
	t.Log(roomCondition)
}

func TestUpdateLRConditionByLid(t *testing.T) {
	roomCondition, err := UpdateLRConditionByLid("22e106ca-ff5d-4876-be231d27-80900799", "", "1769264507@qq.com", 1, 3, 1, 1, 30.234)
	if err != nil {
		t.Errorf("Error of update LRcondition:%v", err)
	}
	t.Log(roomCondition)
}

func TestRetrieveLRConditionByLid(t *testing.T) {
	roomCondition, err := RetrieveLRConditionByLid("22e106ca-ff5d-4876-be231d27-80900799")
	if err != err {
		t.Errorf("Error of retrieve LRcondition:%v", err)
	}
	t.Log(roomCondition)
}
