package dbop

import (
	"log"
	"testing"
)

func TestDeleteResourse(t *testing.T) {
	err := DeleteResourse("382d65b2-a8b6-4a83-a71b5241-9c527aa9")
	log.Println(err)
}
