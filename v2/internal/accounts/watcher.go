package accounts

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"tarkov-account-switcher/internal/config"
)

var (
	watcherMutex     sync.Mutex
	watcherRunning   bool
	watcherAccountID string
	stopChan         chan struct{}

	// SessionCapturedCallback is called when a session is captured
	SessionCapturedCallback func(accountID string)
)

// StartWatcher starts watching for session tokens for the given account.
// Any previously running watcher is stopped first.
func StartWatcher(accountID, expectedEmail string) {
	watcherMutex.Lock()
	if watcherRunning && stopChan != nil {
		close(stopChan)
	}
	watcherRunning = true
	watcherAccountID = accountID
	stopChan = make(chan struct{})
	localStopChan := stopChan
	watcherMutex.Unlock()

	paths := config.GetPaths()
	ticker := time.NewTicker(2 * time.Second)
	timeout := time.After(5 * time.Minute)
	defer ticker.Stop()

	clear := func() {
		watcherMutex.Lock()
		if watcherAccountID == accountID {
			watcherRunning = false
			watcherAccountID = ""
		}
		watcherMutex.Unlock()
	}

	for {
		select {
		case <-localStopChan:
			return

		case <-timeout:
			clear()
			return

		case <-ticker.C:
			data, err := os.ReadFile(paths.LauncherSettingsPath)
			if err != nil {
				continue
			}
			var launcherSettings map[string]any
			if err := json.Unmarshal(data, &launcherSettings); err != nil {
				continue
			}

			login, _ := launcherSettings["login"].(string)
			at, _ := launcherSettings["at"].(string)
			rt, _ := launcherSettings["rt"].(string)
			if login != expectedEmail || at == "" || rt == "" {
				continue
			}

			sessionData, err := json.Marshal(BuildAuthSession(launcherSettings))
			if err != nil {
				continue
			}
			if err := UpdateAccountSession(accountID, sessionData); err != nil {
				continue
			}

			clear()
			if SessionCapturedCallback != nil {
				SessionCapturedCallback(accountID)
			}
			return
		}
	}
}
