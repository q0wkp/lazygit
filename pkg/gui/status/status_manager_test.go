package status

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithWaitingStatusShowsStatusRightAwayWithoutDelay(t *testing.T) {
	mgr := NewStatusManager()

	err := mgr.WithWaitingStatus("Working", 0, func() {}, func(*WaitingStatusHandle) error {
		assert.True(t, mgr.HasStatus())
		return nil
	})

	assert.NoError(t, err)
	assert.False(t, mgr.HasStatus())
}

func TestWithWaitingStatusNeverShowsStatusForQuickOperation(t *testing.T) {
	mgr := NewStatusManager()
	rendered := false

	err := mgr.WithWaitingStatus("Working", time.Hour, func() { rendered = true }, func(*WaitingStatusHandle) error {
		assert.False(t, mgr.HasStatus())
		return nil
	})

	assert.NoError(t, err)
	assert.False(t, rendered)
	assert.False(t, mgr.HasStatus())
}

func TestWithWaitingStatusShowsStatusOnceDelayElapses(t *testing.T) {
	mgr := NewStatusManager()

	err := mgr.WithWaitingStatus("Working", time.Millisecond, func() {}, func(*WaitingStatusHandle) error {
		assert.Eventually(t, mgr.HasStatus, time.Second, time.Millisecond)
		return nil
	})

	assert.NoError(t, err)
	assert.False(t, mgr.HasStatus())
}

func TestWithWaitingStatusShowCancelsPendingDelay(t *testing.T) {
	mgr := NewStatusManager()

	err := mgr.WithWaitingStatus("Working", time.Hour, func() {}, func(handle *WaitingStatusHandle) error {
		handle.Hide()
		handle.Show()
		assert.True(t, mgr.HasStatus())
		return nil
	})

	assert.NoError(t, err)
	assert.False(t, mgr.HasStatus())
}
