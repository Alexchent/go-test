package redis

import (
	"github.com/go-redis/redis/v8"
	"testing"
)

// go test -v -run TestCacheService_Get
func TestCacheService_Get(t *testing.T) {
	clientSvc := NewCacheService(redis.Options{Addr: "127.0.0.1:6379"})
	type fields struct {
		Client *redis.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantVal string
	}{
		{name: "1", wantVal: "18", fields: fields{Client: clientSvc.Client}, args: args{key: "alex"}},
		{name: "2", wantVal: "20", fields: fields{Client: clientSvc.Client}, args: args{key: "tom"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CacheService{
				Client: tt.fields.Client,
			}
			if gotVal := s.Get(tt.args.key); gotVal != tt.wantVal {
				t.Errorf("Get() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}
