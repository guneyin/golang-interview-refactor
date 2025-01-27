package controllers_test

import (
	"interview/controllers"
	"interview/test"
	"testing"
)

func TestNew(t *testing.T) {
	test.ChangeWorkDir()
	test.InitTestDB()

	cnt := controllers.New(test.NewTestRouter())
	if cnt == nil {
		t.Fatal("controller is nil")
	}
}
