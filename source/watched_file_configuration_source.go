package source

import (
    "fmt"

    "github.com/nienie/manyfaced/poll"
    "github.com/nienie/manyfaced/parser"
    "github.com/go-fsnotify/fsnotify"
)

//WatchedFileConfigurationSource ...
type WatchedFileConfigurationSource struct {
    listeners      []poll.WatchedUpdateListener
    configContents map[string]map[string]interface{}
    configFiles    []string
    stop           chan bool
}

//NewWatchedFileConfigurationSource ...
func NewWatchedFileConfigurationSource(configFiles []string) (*WatchedFileConfigurationSource, error) {
    if configFiles == nil {
        return nil, fmt.Errorf("invalid parameters, configFiles is nil")
    }
    watchedFileSource := &WatchedFileConfigurationSource{
        listeners:          make([]poll.WatchedUpdateListener, 0),
        configContents:     make(map[string]map[string]interface{}, 0),
        configFiles:        configFiles,
        stop:               make(chan bool),
    }
    err := watchedFileSource.initialize()
    if err != nil {
        return nil, err
    }
    return watchedFileSource, nil
}

func (s *WatchedFileConfigurationSource)initialize() error {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }
    for _, configFile := range s.configFiles {
        err = watcher.Add(configFile)
        if err != nil {
            return err
        }
        s.configContents[configFile], err = parser.ParseConfigFile(configFile)
        if err != nil {
            return err
        }
    }
    go s.watch(watcher)
    return nil
}

func (s *WatchedFileConfigurationSource)watch(watcher *fsnotify.Watcher) {
    defer func(){
        if r := recover(); r != nil {
            //TODO: Add logger
        }
        watcher.Close()
    }()
    for {
        select {
        case event := <-watcher.Events:
            //TODO: sync or async ???
            s.handleWatchEvent(event)
        case <-watcher.Errors:
        //TODO: Add Log

        case <-s.stop:
            return
        }
    }
}

func (s WatchedFileConfigurationSource)handleWatchEvent(event fsnotify.Event) {
    switch event.Op {
    case fsnotify.Create:
        newConfigContent, err := parser.ParseConfigFile(event.Name)
        if err != nil {
            //TODO: Add Log
            return
        }
        updateResult := poll.NewIncrementalWatchUpdatedResult(newConfigContent, nil, nil)
        s.notifyListeners(updateResult)
        s.configContents[event.Name] = newConfigContent
    case fsnotify.Write:
        newConfigContent, err := parser.ParseConfigFile(event.Name)
        if err != nil {
            //TODO: Add Log
            return
        }
        added, changed, deleted := compareUpdateResult(newConfigContent, s.configContents[event.Name])
        updateResult := poll.NewIncrementalWatchUpdatedResult(added, changed, deleted)
        s.notifyListeners(updateResult)
        s.configContents[event.Name] = newConfigContent
    case fsnotify.Remove:
        updateResult := poll.NewIncrementalWatchUpdatedResult(nil, nil, s.configContents[event.Name])
        s.notifyListeners(updateResult)
        s.configContents[event.Name] = make(map[string]interface{})
    case fsnotify.Rename:
        //Nothing to do
    case fsnotify.Chmod:
        //Nothing to do
    }
}

func (s *WatchedFileConfigurationSource)notifyListeners(updateResult *poll.WatchedUpdateResult) {
    for _, listener := range s.listeners {
        listener.UpdateConfiguration(updateResult)
    }
}

func compareUpdateResult(newContent map[string]interface{}, oldContent map[string]interface{}) (added, changed, deleted map[string]interface{}){
    for key, val := range newContent {
        oldVal, ok := oldContent[key]
        if !ok {
            added[key] = val
            continue
        }
        if oldVal != val {
            changed[key] = val
        }
    }

    for key, val := range oldContent {
        _, ok := newContent[key]
        if !ok {
            deleted[key] = val
        }
    }
    return
}

//GetCurrentData ...
func (s *WatchedFileConfigurationSource)GetCurrentData() (map[string]interface{}, error) {
    currentData := make(map[string]interface{})
    for _, configFile := range s.configFiles {
        m, err := parser.ParseConfigFile(configFile)
        if err != nil {
            return nil, err
        }
        for key, val := range m {
            currentData[key] = val
        }
    }
    return currentData, nil
}

//AddUpdateListener ...
func (s *WatchedFileConfigurationSource)AddUpdateListener(l poll.WatchedUpdateListener) {
    if l != nil {
        s.listeners = append(s.listeners, l)
    }
    return
}

//RemoveUpdateListener ...
func (s *WatchedFileConfigurationSource)RemoveUpdateListener(l poll.WatchedUpdateListener) {
    if l == nil {
        return
    }
    var index = -1
    for i, listener := range s.listeners {
        if listener == l {
            index = i
            break
        }
    }
    if index == - 1 {
        return
    }
    s.listeners = append(s.listeners[:index], s.listeners[index+1:]...)
    return
}

func (s *WatchedFileConfigurationSource)StopWatching() {
    s.stop <- true
}