package lokr_test

import (
	"context"
	"testing"
	"time"

	"github.com/clavinjune/lokr"
	lokrv1 "github.com/clavinjune/lokr/api/lokr/v1"
	"github.com/clavinjune/lokr/mocks"
	"github.com/clavinjune/lokr/pkg"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLokr_RegisterLock(t *testing.T) {
	var tt = []struct {
		_           struct{}
		name        string
		key         string
		expectedErr error
	}{
		{
			name:        "key is filled",
			key:         "demo-key",
			expectedErr: nil,
		},
		{
			name:        "key is empty",
			key:         "",
			expectedErr: pkg.ErrEmptyLockKey,
		},
	}

	r := require.New(t)
	repo := mocks.NewRepository(t)

	l, err := lokr.ProvideLokr(repo, false)
	r.NoError(err)

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			lock := lokrv1.Lock{
				Key: tc.key,
			}

			if tc.key != "" {
				repo.EXPECT().
					StoreLock(mock.Anything, &lock).
					Return(tc.expectedErr)
			}

			err = l.RegisterLock(context.Background(), &lock)
			if tc.expectedErr == nil {
				r.NoError(err)
			} else {
				r.ErrorIs(err, tc.expectedErr)
			}
		})
	}
}

func TestLokr_TryObtain(t *testing.T) {
	var tt = []struct {
		_           struct{}
		name        string
		key         string
		isLocked    bool
		expectedOk  bool
		expectedErr error
	}{
		{
			name:        "key is empty",
			key:         "",
			isLocked:    false,
			expectedOk:  false,
			expectedErr: pkg.ErrEmptyLockKey,
		},
		{
			name:        "locked",
			key:         "demo-key",
			isLocked:    true,
			expectedOk:  false,
			expectedErr: nil,
		},
	}

	r := require.New(t)
	repo := mocks.NewRepository(t)

	l, err := lokr.ProvideLokr(repo, false)
	r.NoError(err)

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.key != "" {
				repo.EXPECT().
					Tx(mock.Anything, mock.Anything).
					Return(nil)
			}

			ok, err := l.TryObtain(context.Background(), tc.key)
			r.Equal(tc.expectedOk, ok)
			if tc.expectedErr == nil {
				r.NoError(err)
			} else {
				r.ErrorIs(err, tc.expectedErr)
			}
		})
	}
}

func TestLokr_Release(t *testing.T) {
	var tt = []struct {
		_           struct{}
		name        string
		key         string
		expectedErr error
	}{
		{
			name:        "key is empty",
			key:         "",
			expectedErr: pkg.ErrEmptyLockKey,
		},
		{
			name:        "locked",
			key:         "demo-key",
			expectedErr: nil,
		},
	}

	r := require.New(t)
	repo := mocks.NewRepository(t)

	l, err := lokr.ProvideLokr(repo, false)
	r.NoError(err)

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.key != "" {
				repo.EXPECT().
					PatchLockByKey(mock.Anything, mock.Anything).
					Return(nil)
			}

			err := l.Release(context.Background(), tc.key)
			if tc.expectedErr == nil {
				r.NoError(err)
			} else {
				r.ErrorIs(err, tc.expectedErr)
			}
		})
	}
}

func TestLokr_PollObtain(t *testing.T) {
	t.Parallel()

	r := require.New(t)
	repo := mocks.NewRepository(t)

	l, err := lokr.ProvideLokr(repo, true)
	r.NoError(err)

	repo.EXPECT().
		Tx(mock.Anything, mock.Anything).
		Return(nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	defer cancel()
	l.PollObtain(ctx, "demo-lock-key", time.Second/4, mocks.NewJobHandler(t))
}
