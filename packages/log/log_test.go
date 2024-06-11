package log

// import (
// 	"context"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"gitlab.com/otoklix/backend/payment-service/package/config"
// )

// func TestAddlog(t *testing.T) {
// 	r := httptest.NewRequest("GET", "/?id=1", nil)
// 	ctx := context.WithValue(r.Context(), config.ContextKey("startTime"), time.Now())
// 	ctx = context.WithValue(ctx, config.ContextKey("body"), []byte(`{"id":1}`))
// 	r.Header.Set("Authorization", "Bearer abc")
// 	r = r.WithContext(ctx)

// 	r2 := httptest.NewRequest("GET", "/", nil)
// 	ctx = context.WithValue(ctx, config.ContextKey("body"), []byte{})
// 	r2 = r2.WithContext(ctx)

// 	type args struct {
// 		r     *http.Request
// 		level string
// 		data  interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				r:     r,
// 				level: "SUCCESS",
// 			},
// 		},
// 		{
// 			name: "success",
// 			args: args{
// 				r:     r2,
// 				level: "ERROR",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Addlog(tt.args.r, tt.args.level, tt.args.data)
// 		})
// 	}
// }
