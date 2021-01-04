package models

import (
	"reflect"
	"testing"
)

func TestHost_Map(t *testing.T) {
	type fields struct {
		ID       string
		IP       string
		Name     string
		Metadata map[string]interface{}
		Status   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - host 1",
			fields: fields{
				ID:   "Host 1",
				IP:   "192.168.122.236",
				Name: "Linux",
				Metadata: map[string]interface{}{
					"Os":      "Linux",
					"Version": "17.2",
				},
				Status: true,
			},
			want: map[string]interface{}{
				"id":   "Host 1",
				"ip":   "192.168.122.236",
				"name": "Linux",
				"metadata": map[string]interface{}{
					"Os":      "Linux",
					"Version": "17.2",
				},
				"status": true,
			},
		},
		{
			name: "success - host2 with fewer fields",
			fields: fields{
				ID:       "Host 2",
				IP:       "192.168.0.242",
				Metadata: map[string]interface{}{},
			},
			want: map[string]interface{}{
				"id":       "Host 2",
				"ip":       "192.168.0.242",
				"name":     "",
				"metadata": map[string]interface{}{},
				"status":   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Host{
				ID:       tt.fields.ID,
				IP:       tt.fields.IP,
				Name:     tt.fields.Name,
				Metadata: tt.fields.Metadata,
				Status:   tt.fields.Status,
			}
			if got := h.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHost_Names(t *testing.T) {
	type fields struct {
		ID       string
		IP       string
		Name     string
		Metadata map[string]interface{}
		Status   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Host 1",
			fields: fields{
				ID:   "Host 1",
				IP:   "192.168.122.236",
				Name: "Linux",
				Metadata: map[string]interface{}{
					"Os":      "Linux",
					"Version": "17.2",
				},
				Status: true,
			},
			want: []string{"id", "ip", "name", "metadata", "status"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Host{
				ID:       tt.fields.ID,
				IP:       tt.fields.IP,
				Name:     tt.fields.Name,
				Metadata: tt.fields.Metadata,
				Status:   tt.fields.Status,
			}
			if got := h.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
