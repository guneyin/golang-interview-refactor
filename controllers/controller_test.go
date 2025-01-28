package controllers_test

import (
	"interview/controllers"
	"interview/testutils"
	"testing"
)

func TestNew(t *testing.T) {
	testutils.ChangeWorkDir()
	testutils.InitTestDB()

	cnt := controllers.New(testutils.NewTestRouter())
	if cnt == nil {
		t.Fatal("controller is nil")
	}
}
