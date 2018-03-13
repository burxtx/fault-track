package service

import (
	"fmt"
	"os"
	"testing"
	"context"
	
	"github.com/burxtx/fault/config"
	"github.com/burxtx/fault/fault/pkg/model"
	"github.com/burxtx/fault/db"
)

func BenchmarkNewFault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := new(model.Fault)
		fmt.Println("%p", &f)
	}
}
func TestAddService(t *testing.T) {
	b := new(basicFaultService)
	f := model.Fault{
		Description: "asasas",
		UserName: "txtx",
	}
	var ctx context.Context
	affected, err := b.Add(f)
	if err != nil {
		t.Errorf("insert error: %s", err)
	}
	if affected < 1 {
		t.Errorf("not inserted")
	}
}

func TestMain(m *testing.M) {
	env := "test"
	config.Init(env)
	db.Init()
	os.Exit(m.Run())
}
