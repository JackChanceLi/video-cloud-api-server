package dbop

import (
	"fmt"
	"testing"
)

func TestRetrieveLiveRoomByCid(t *testing.T) {
	room, err := RetrieveLiveRoomByCid("dfc30895-f79f-460d-ba5f83ff-a7bd23fd")
	if err != nil {
		t.Errorf("Error of retrieve liveroom: %v", err)
	}

	fmt.Println(room)
}
