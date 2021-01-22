package mysql

import (
	"os"
	"reflect"
	"testing"

	"github.com/kashifkhan0771/HostService/db"
	"github.com/kashifkhan0771/HostService/models"
)

func Test_client_AddHost(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "3306")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_USER", "root")
	type args struct {
		host *models.Host
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add host t db",
			args: args{host: &models.Host{IP: "127.0.0.1", Name: "Ubuntu", Metadata: map[string]interface{}{
				"OS":      "Linux",
				"Version": "20.4",
			}, Status: true}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
	t.Parallel()
	_ = os.Setenv("DB_PORT", "3306")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_USER", "root")
	c, _ := NewClient(db.Option{})
	host := &models.Host{IP: "127.1.1.1", Name: "RedHat", Metadata: map[string]interface{}{
		"OS":      "Linux",
		"Version": "12.5",
	}, Status: false}
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
			name:    "success - delete host form db",
			args:    args{id: host.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if c != nil {
				if err := c.DeleteHost(tt.args.id); (err != nil) != tt.wantErr {
					t.Errorf("DeleteHost() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func Test_client_GetHost(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "3306")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_USER", "root")
	c, _ := NewClient(db.Option{})
	host := &models.Host{IP: "127.0.0.1", Name: "Ubuntu", Metadata: map[string]interface{}{
		"OS":      "Linux",
		"Version": "20.4",
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
			name:    "success - get host data from db",
			args:    args{id: host.ID},
			want:    host,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if c != nil {
				got, err := c.GetHost(tt.args.id)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetHost() error = %v, wantErr %v", err, tt.wantErr)

					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetHost() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_client_UpdateHost(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "3306")
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_USER", "root")
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
			name: "success - update host in db",
			args: args{host: &models.Host{ID: "1", IP: "127.0.01.1", Name: "Ubuntu", Metadata: map[string]interface{}{
				"OS":      "Linux",
				"Version": "20.4",
			}, Status: true}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if c != nil {
				if err := c.UpdateHost(tt.args.host); (err != nil) != tt.wantErr {
					t.Errorf("UpdateHost() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
