// Copyright 2020 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestActionsService_ListWorkflows(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"per_page": "2", "page": "2"})
		fmt.Fprint(w, `{"total_count":4,"workflows":[{"id":72844,"created_at":"2019-01-02T15:04:05Z","updated_at":"2020-01-02T15:04:05Z"},{"id":72845,"created_at":"2019-01-02T15:04:05Z","updated_at":"2020-01-02T15:04:05Z"}]}`)
	})

	opts := &ListOptions{Page: 2, PerPage: 2}
	workflows, _, err := client.Actions.ListWorkflows(context.Background(), "o", "r", opts)
	if err != nil {
		t.Errorf("Actions.ListWorkflows returned error: %v", err)
	}

	want := &Workflows{
		TotalCount: Int(4),
		Workflows: []*Workflow{
			{ID: Int64(72844), CreatedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, UpdatedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
			{ID: Int64(72845), CreatedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, UpdatedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
		},
	}
	if !reflect.DeepEqual(workflows, want) {
		t.Errorf("Actions.ListWorkflows returned %+v, want %+v", workflows, want)
	}
}

func TestActionsService_GetWorkflowByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows/72844", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":72844,"created_at":"2019-01-02T15:04:05Z","updated_at":"2020-01-02T15:04:05Z"}`)
	})

	workflow, _, err := client.Actions.GetWorkflowByID(context.Background(), "o", "r", 72844)
	if err != nil {
		t.Errorf("Actions.GetWorkflowByID returned error: %v", err)
	}

	want := &Workflow{
		ID:        Int64(72844),
		CreatedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)},
		UpdatedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)},
	}
	if !reflect.DeepEqual(workflow, want) {
		t.Errorf("Actions.GetWorkflowByID returned %+v, want %+v", workflow, want)
	}
}

func TestActionsService_GetWorkflowByFileName(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows/main.yml", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":72844,"created_at":"2019-01-02T15:04:05Z","updated_at":"2020-01-02T15:04:05Z"}`)
	})

	workflow, _, err := client.Actions.GetWorkflowByFileName(context.Background(), "o", "r", "main.yml")
	if err != nil {
		t.Errorf("Actions.GetWorkflowByFileName returned error: %v", err)
	}

	want := &Workflow{
		ID:        Int64(72844),
		CreatedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)},
		UpdatedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)},
	}
	if !reflect.DeepEqual(workflow, want) {
		t.Errorf("Actions.GetWorkflowByFileName returned %+v, want %+v", workflow, want)
	}
}

func TestActionsService_GetWorkflowUsageByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows/72844/timing", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"billable":{"UBUNTU":{"total_ms":180000},"MACOS":{"total_ms":240000},"WINDOWS":{"total_ms":300000}}}`)
	})

	workflowUsage, _, err := client.Actions.GetWorkflowUsageByID(context.Background(), "o", "r", 72844)
	if err != nil {
		t.Errorf("Actions.GetWorkflowUsageByID returned error: %v", err)
	}

	want := &WorkflowUsage{
		Billable: &WorkflowEnvironment{
			Ubuntu: &WorkflowBill{
				TotalMS: Int64(180000),
			},
			MacOS: &WorkflowBill{
				TotalMS: Int64(240000),
			},
			Windows: &WorkflowBill{
				TotalMS: Int64(300000),
			},
		},
	}
	if !reflect.DeepEqual(workflowUsage, want) {
		t.Errorf("Actions.GetWorkflowUsageByID returned %+v, want %+v", workflowUsage, want)
	}
}

func TestActionsService_GetWorkflowUsageByFileName(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows/main.yml/timing", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"billable":{"UBUNTU":{"total_ms":180000},"MACOS":{"total_ms":240000},"WINDOWS":{"total_ms":300000}}}`)
	})

	workflowUsage, _, err := client.Actions.GetWorkflowUsageByFileName(context.Background(), "o", "r", "main.yml")
	if err != nil {
		t.Errorf("Actions.GetWorkflowUsageByFileName returned error: %v", err)
	}

	want := &WorkflowUsage{
		Billable: &WorkflowEnvironment{
			Ubuntu: &WorkflowBill{
				TotalMS: Int64(180000),
			},
			MacOS: &WorkflowBill{
				TotalMS: Int64(240000),
			},
			Windows: &WorkflowBill{
				TotalMS: Int64(300000),
			},
		},
	}
	if !reflect.DeepEqual(workflowUsage, want) {
		t.Errorf("Actions.GetWorkflowUsageByFileName returned %+v, want %+v", workflowUsage, want)
	}
}

func TestActionsService_CreateWorkflowDispatchEventByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	event := CreateWorkflowDispatchEventRequest{
		Ref: "d4cfb6e7",
		Inputs: map[string]interface{}{
			"key": "value",
		},
	}
	mux.HandleFunc("/repos/o/r/actions/workflows/72844/dispatches", func(w http.ResponseWriter, r *http.Request) {
		var v CreateWorkflowDispatchEventRequest
		json.NewDecoder(r.Body).Decode(&v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, event) {
			t.Errorf("Request body = %+v, want %+v", v, event)
		}
	})

	_, err := client.Actions.CreateWorkflowDispatchEventByID(context.Background(), "o", "r", 72844, event)
	if err != nil {
		t.Errorf("Actions.CreateWorkflowDispatchEventByID returned error: %v", err)
	}

	// Test s.client.NewRequest failure
	client.BaseURL.Path = ""
	_, err = client.Actions.CreateWorkflowDispatchEventByID(context.Background(), "o", "r", 72844, event)
	if err == nil {
		t.Error("client.BaseURL.Path='' CreateWorkflowDispatchEventByID err = nil, want error")
	}
}

func TestActionsService_CreateWorkflowDispatchEventByFileName(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	event := CreateWorkflowDispatchEventRequest{
		Ref: "d4cfb6e7",
		Inputs: map[string]interface{}{
			"key": "value",
		},
	}
	mux.HandleFunc("/repos/o/r/actions/workflows/main.yml/dispatches", func(w http.ResponseWriter, r *http.Request) {
		var v CreateWorkflowDispatchEventRequest
		json.NewDecoder(r.Body).Decode(&v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, event) {
			t.Errorf("Request body = %+v, want %+v", v, event)
		}
	})

	_, err := client.Actions.CreateWorkflowDispatchEventByFileName(context.Background(), "o", "r", "main.yml", event)
	if err != nil {
		t.Errorf("Actions.CreateWorkflowDispatchEventByFileName returned error: %v", err)
	}

	// Test s.client.NewRequest failure
	client.BaseURL.Path = ""
	_, err = client.Actions.CreateWorkflowDispatchEventByFileName(context.Background(), "o", "r", "main.yml", event)
	if err == nil {
		t.Error("client.BaseURL.Path='' CreateWorkflowDispatchEventByFileName err = nil, want error")
	}
}

func TestActionsService_EnableWorkflowByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows/72844/enable", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if r.Body != http.NoBody {
			t.Errorf("Request body = %+v, want %+v", r.Body, http.NoBody)
		}
	})

	_, err := client.Actions.EnableWorkflowByID(context.Background(), "o", "r", 72844)
	if err != nil {
		t.Errorf("Actions.EnableWorkflowByID returned error: %v", err)
	}

	// Test s.client.NewRequest failure
	client.BaseURL.Path = ""
	_, err = client.Actions.EnableWorkflowByID(context.Background(), "o", "r", 72844)
	if err == nil {
		t.Error("client.BaseURL.Path='' EnableWorkflowByID err = nil, want error")
	}
}

func TestActionsService_EnableWorkflowByFilename(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows/main.yml/enable", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if r.Body != http.NoBody {
			t.Errorf("Request body = %+v, want %+v", r.Body, http.NoBody)
		}
	})

	_, err := client.Actions.EnableWorkflowByFileName(context.Background(), "o", "r", "main.yml")
	if err != nil {
		t.Errorf("Actions.EnableWorkflowByFilename returned error: %v", err)
	}

	// Test s.client.NewRequest failure
	client.BaseURL.Path = ""
	_, err = client.Actions.EnableWorkflowByFileName(context.Background(), "o", "r", "main.yml")
	if err == nil {
		t.Error("client.BaseURL.Path='' EnableWorkflowByFilename err = nil, want error")
	}
}

func TestActionsService_DisableWorkflowByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows/72844/disable", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if r.Body != http.NoBody {
			t.Errorf("Request body = %+v, want %+v", r.Body, http.NoBody)
		}
	})

	_, err := client.Actions.DisableWorkflowByID(context.Background(), "o", "r", 72844)
	if err != nil {
		t.Errorf("Actions.DisableWorkflowByID returned error: %v", err)
	}

	// Test s.client.NewRequest failure
	client.BaseURL.Path = ""
	_, err = client.Actions.DisableWorkflowByID(context.Background(), "o", "r", 72844)
	if err == nil {
		t.Error("client.BaseURL.Path='' DisableWorkflowByID err = nil, want error")
	}
}

func TestActionsService_DisableWorkflowByFileName(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/workflows/main.yml/disable", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		if r.Body != http.NoBody {
			t.Errorf("Request body = %+v, want %+v", r.Body, http.NoBody)
		}
	})

	_, err := client.Actions.DisableWorkflowByFileName(context.Background(), "o", "r", "main.yml")
	if err != nil {
		t.Errorf("Actions.DisableWorkflowByFileName returned error: %v", err)
	}

	// Test s.client.NewRequest failure
	client.BaseURL.Path = ""
	_, err = client.Actions.DisableWorkflowByFileName(context.Background(), "o", "r", "main.yml")
	if err == nil {
		t.Error("client.BaseURL.Path='' DisableWorkflowByFileName err = nil, want error")
	}
}
