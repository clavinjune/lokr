package lokr

import (
	"context"
	"log"
	"strings"
	"time"

	lokrv1 "github.com/clavinjune/lokr/api/lokr/v1"
	"github.com/clavinjune/lokr/pkg"
	"github.com/jmoiron/sqlx"
)

// Lokr provides simple distributed locking mechanism on SQL Database
type Lokr struct {
	_           struct{}
	repo        pkg.Repository
	isDebugMode bool
}

// RegisterLock registers lock key to the database
func (l *Lokr) RegisterLock(ctx context.Context, lock *lokrv1.Lock) error {
	if strings.TrimSpace(lock.Key) == "" {
		return pkg.ErrEmptyLockKey
	}

	return l.repo.StoreLock(ctx, lock)
}

// TryObtain tries to obtain the lock with the given key
// See PollObtain to repeatedly obtain the keys
// returns false if failed to obtain the lock
// returns true if the lock successfully obtained, you need to release the lock afterwards
// if error is not nil, consider the obtain process is failed and no need to release the lock
func (l *Lokr) TryObtain(ctx context.Context, key string) (bool, error) {
	if strings.TrimSpace(key) == "" {
		return false, pkg.ErrEmptyLockKey
	}

	repo := l.repo
	isObtained := false
	err := repo.Tx(ctx, l.tryObtainTxHandler(&isObtained, key))
	if err != nil {
		return false, err
	}
	return isObtained, nil
}

// Release releases the obtained lock with the given key
// this action is a must after user successfully obtain the
func (l *Lokr) Release(ctx context.Context, key string) error {
	if strings.TrimSpace(key) == "" {
		return pkg.ErrEmptyLockKey
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return l.repo.PatchLockByKey(ctx, &lokrv1.Lock{
		Key:      key,
		IsLocked: false,
	})
}

// PollObtain polls given key with given interval,
// will run given fn if successfully obtain the lock and release it afterwards
// PollObtain will stop polling if the given context is cancelled
func (l *Lokr) PollObtain(ctx context.Context, key string, interval time.Duration, fn pkg.JobHandler) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	isObtained := false
	for {
		select {
		case c := <-ctx.Done():
			l.log("context done", c)
			if !isObtained {
				l.log("has not obtained anything, returning")
				return
			}

			l.logf("has obtained %v, releasing lock", key)
			// using new context because the given context is reaching deadline
			if err := l.Release(context.Background(), key); err != nil {
				l.logf("there's an error when releasing %v: %v", key, err)
			}
		case c := <-ticker.C:
			l.log(c, "polling lock", key)
			var err error
			isObtained, err = l.TryObtain(ctx, key)
			if err != nil || !isObtained {
				continue
			}
			l.log(key, "lock obtained")
			fn.Handle(ctx, key)
			// using new context because the given context might reaching deadline
			if err := l.Release(context.Background(), key); err != nil {
				l.logf("there's an error when releasing lock %v: %v", key, err)
			} else {
				l.log("successfully releasing lock", key)
				isObtained = false
			}
		}
	}
}

func (l *Lokr) tryObtainTxHandler(isObtained *bool, key string) pkg.TxHandlerFunc {
	repo := l.repo
	return func(ctx context.Context, tx *sqlx.Tx) error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		lock, err := repo.FetchLockTx(ctx, tx, key)
		if err != nil {
			return err
		}

		if lock.IsLocked {
			return nil
		}

		lock.IsLocked = true
		if err := repo.PatchLockByKeyTx(ctx, tx, lock); err != nil {
			return err
		}

		*isObtained = true
		return nil
	}
}

func (l *Lokr) log(v ...any) {
	if !l.isDebugMode {
		return
	}

	log.Println(v...)
}

func (l *Lokr) logf(fmt string, v ...any) {
	if !l.isDebugMode {
		return
	}

	log.Printf(fmt, v...)
}
