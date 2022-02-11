package repo

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/luqman-v1/absent/repo/mocks"

	"github.com/luqman-v1/absent/helper"

	"github.com/stretchr/testify/assert"
)

func TestGC_ListEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	services := mocks.NewMockIServices(ctrl)
	//init the loc
	loc, _ := time.LoadLocation(helper.TimeZone)
	//set timezone,
	now := time.Now().In(loc)
	startDate := now.Format(time.RFC3339)
	//get end date
	tomorrow := now.AddDate(0, 0, 1)
	endDate := tomorrow.Format(time.RFC3339)

	g := NewGC()
	g.StartDate = startDate
	g.EndDate = endDate
	g.Services = services

	//set mock
	services.EXPECT().GetListEvents(startDate, endDate).Return([]byte("test"), nil)

	list, err := g.ListEvent()
	assert.Equal(t, string(list), "test")
	assert.Nil(t, err)
}
