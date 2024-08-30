package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

type mockHandler struct {
	uc *MockusecaseInterface
}

func (mh *mockHandler) toHandler() *Handler {
	return &Handler{
		uc: mh.uc,
	}
}

func newMockHandler(ctrl *gomock.Controller) *mockHandler {
	return &mockHandler{
		uc: NewMockusecaseInterface(ctrl),
	}
}

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := newMockHandler(ctrl)

	got := New(mu.uc)
	assert.NotNil(t, got)
}

func TestHandler_GetBooksList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name   string
		args   args
		mockFn func(*mockHandler, *args)
		want   http.ResponseWriter
	}{
		{
			name: "case bad request",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{`)),
			},
			mockFn: func(mu *mockHandler, a *args) {},
			want:   mockMissingParameterSubjectResponse(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mh := newMockHandler(ctrl)
			tt.mockFn(mh, &tt.args)
			h := mh.toHandler()

			h.GetBooksList(tt.args.w, tt.args.r)
			assert.Equal(t, tt.want, tt.args.w)
		})
	}
}
