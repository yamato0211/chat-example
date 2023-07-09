package healthz

import (
	"context"
	"fmt"
	"testing"

	"github.com/bufbuild/connect-go"
	"github.com/golang/mock/gomock"
	"github.com/sivchari/chat-example/pkg/log/mock_log"
	"github.com/sivchari/chat-example/proto/proto"
	"github.com/sivchari/chat-example/proto/proto/protoconnect"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	logger *mock_log.MockHandler
}

func newWithMocks(t *testing.T) (context.Context, protoconnect.HealthzHandler, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logger := mock_log.NewMockHandler(ctrl)
	return ctx, New(
			logger,
		), &mocks{
			logger,
		}
}

func TestServer_Check(t *testing.T) {
	name := "name"

	ctx, s, m := newWithMocks(t)
	m.logger.EXPECT().InfoCtx(ctx, "healthz check", "name", name)
	res, err := s.Check(ctx, connect.NewRequest(&proto.CheckRequest{
		Name: name,
	}))
	assert.Equal(t, connect.NewResponse(&proto.CheckResponse{
		Msg: fmt.Sprintf("Hello %s", name),
	}), res)
	assert.NoError(t, err)
}
