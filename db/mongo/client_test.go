package mongo

import (
	"os"
	"reflect"
	"testing"

	"github.com/kashifkhan0771/HostService/db"
	"github.com/kashifkhan0771/HostService/models"
)

func Test_client_AddHost(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	type args struct {
		host *models.Host
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add host in mongo db",
			args: args{host: &models.Host{IP: "189.191.0.1", Name: "Linux", Metadata: map[string]interface{}{
				"OS":      "Linux",
				"Version": 20.4,
			}, Status: true}},
			wantErr: false,
		},
		{
			name: "fail - add wrong credentials",
			args: args{host: &models.Host{ID: "Host 1", IP: "190.190.1.1", Name: "Windows", Metadata: map[string]interface{}{
				"OS":      "Windows XP",
				"Version": 06,
			}, Status: false}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			if c != nil {
				_, err := c.AddHost(tt.args.host)
				if (err != nil) != tt.wantErr {
					t.Errorf("AddHost() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func Test_client_DeleteHost(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")
	c, _ := NewClient(db.Option{})
	host := &models.Host{IP: "191.191.191.0", Name: "Ubuntu", Metadata: map[string]interface{}{
		"OS":      "Linux",
		"Version": 20.4,
	}, Status: true}
	if c != nil {
		_, _ = c.AddHost(host)
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success - delete host from db",
			args:    args{id: host.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteHost(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteHost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetHost(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})
	host := &models.Host{IP: "189.191.0.1", Name: "Linux", Metadata: map[string]interface{}{
		"OS":      "Linux",
		"Version": 20.4,
	}, Status: true}
	if c != nil {
		_, _ = c.AddHost(host)
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Host
		wantErr bool
	}{
		{
			name:    "Success - get student from db",
			args:    args{host.ID},
			want:    host,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetHost(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateHost(t *testing.T) {
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})
	type args struct {
		host *models.Host
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success - update host in db",
			args: args{host: &models.Host{IP: "189.191.0.1", Name: "Linux", Metadata: map[string]interface{}{
				"OS":      "Linux",
				"Version": 20.4,
			}, Status: true}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if c != nil {
			if err := c.UpdateHost(tt.args.host); (err != nil) != tt.wantErr {
				t.Errorf("UpdateHost() error = %v, wantErr %v", err, tt.wantErr)
			}
		}
	}
}
