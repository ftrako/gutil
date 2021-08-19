package goutils

import (
	"github.com/fsnotify/fsnotify"
)

type FileWatchCB func(path string, err error)

// 监听文件
func DoFileWatch(path string, cb FileWatchCB) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		if cb != nil {
			cb(path, err)
		}
		return
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					continue
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					if cb != nil {
						cb(path, nil)
					}
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					CreateFile(path)
					err := watcher.Add(path)
					if err != nil {
						if cb != nil {
							cb(path, err)
						}
						return
					}
					if cb != nil {
						cb(path, nil)
					}
					continue
				}

			case err, ok := <-watcher.Errors:
				if err != nil {
					if cb != nil {
						cb(path, err)
					}
					return
				}
				if !ok {
					continue
				}
			}
		}
	}()

	err = CreateFile(path)
	if err != nil {
		if cb != nil {
			cb(path, err)
		}
		return
	}

	err = watcher.Add(path)
	if err != nil {
		if cb != nil {
			cb(path, err)
		}
		return
	}
	<-done
}
