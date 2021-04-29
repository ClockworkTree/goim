package time

import (
	"log"
	"testing"
	"time"
)

func TestTimer_Init(t1 *testing.T) {
	t := Timer{}
	t.init(4)
	time.Sleep(time.Second)
	log.Printf("%+v", *t.free)
	log.Printf("%+v", t.free)
}
