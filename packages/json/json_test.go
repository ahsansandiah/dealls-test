package json

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestWriteJSON(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	code := 200
// 	res := response{
// 		Data: true,
// 	}

// 	type args struct {
// 		w    http.ResponseWriter
// 		code int
// 		v    interface{}
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				w:    w,
// 				code: code,
// 				v:    res,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := WriteJSON(tt.args.w, tt.args.code, tt.args.v); (err != nil) != tt.wantErr {
// 				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestSuccessResponse(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/", nil)

// 	type args struct {
// 		w          http.ResponseWriter
// 		r          *http.Request
// 		statusCode int
// 		message    interface{}
// 		data       interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				w:          w,
// 				r:          r,
// 				statusCode: 200,
// 				message:    "success",
// 				data:       "anu",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			SuccessResponse(tt.args.w, tt.args.r, tt.args.statusCode, tt.args.message, tt.args.data)
// 		})
// 	}
// }

// func TestErrorResponse(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/", nil)

// 	type args struct {
// 		w          http.ResponseWriter
// 		r          *http.Request
// 		statusCode int
// 		message    interface{}
// 		code       string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				w:          w,
// 				r:          r,
// 				statusCode: 200,
// 				message:    "success",
// 				code:       "anu",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ErrorResponse(tt.args.w, tt.args.r, tt.args.statusCode, tt.args.message, tt.args.code)
// 		})
// 	}
// }
