package archive

import (
	"gainsmith/cmd/server"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tests"
)

const testDataDir = "../../test_pb_data"

func setupTestApp(t testing.TB) *tests.TestApp {
	testApp, err := tests.NewTestApp(testDataDir)
	if err != nil {
		t.Fatal(err)
	}

	// Bootstrap on serve
	testApp.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Create the Pocketbase Superuser
		server.CreateSuperuser(se.App)

		// register archive program route
		se.Router.POST("/api/archive/create", HandlerArchive)

		return se.Next()
	})

	return testApp
}

func TestHandler(t *testing.T) {
	// assess
	scenarios := []tests.ApiScenario{
		{
			Name:   "Default Request",
			Method: "POST",
			URL:    "/api/archive/create",
			Body: strings.NewReader(`{
				"name": "Test",
				"from": "2026-01-01T00:00:00.000Z",
				"to": "2027-01-01T00:00:00.000Z",
				"notes": "These are notes"
				}`),
			ExpectedStatus:  200,
			ExpectedContent: []string{`"id":`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Failing Request",
			Method: "POST",
			URL:    "/api/archive/create",
			Body: strings.NewReader(`{
				"name": "",
				"from": "2026-01-01T00:00:00.000Z",
				"to": "2027-01-01T00:00:00.000Z",
				"notes": "These are notes"
				}`),
			ExpectedStatus:  400,
			ExpectedContent: []string{"Invalid archive data"},
			TestAppFactory:  setupTestApp,
		},
	}

	// act & assert
	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
