package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/luqman-v1/absent/repo/mocks"
)

func TestTalenta_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	services := mocks.NewMockIServices(ctrl)
	g := NewTalenta()
	g.Services = services
	response := `
{
	"data": {
		"token": "test"
	}
}
`
	services.EXPECT().LoginTalenta().Return([]byte(response), nil)
	b, err := g.Login()
	assert.Equal(t, "test", b)
	assert.Nil(t, err)
}

func TestTalenta_Present(t *testing.T) {
	ctrl := gomock.NewController(t)
	services := mocks.NewMockIServices(ctrl)
	g := NewTalenta()
	g.Services = services
	response := `
{
	"data": {
		"token": "test"
	}
}
`
	services.EXPECT().LoginTalenta().Return([]byte(response), nil)
	services.EXPECT().PresentTalenta(gomock.Any(), gomock.Any()).Return([]byte("test"), nil)
	_, err := g.Present("checkin")
	assert.Nil(t, err)
}
