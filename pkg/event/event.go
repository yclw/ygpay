package event

// 事件

import (
	"context"
	"sync"
)

var event *SEvent

type EventFunc func(ctx context.Context, args ...interface{})

type SEvent struct {
	sync.Mutex
	list map[string][]EventFunc // 所有事件的列表
}

// Event 事件实例
func Event() *SEvent {
	if event == nil {
		event = &SEvent{
			list: make(map[string][]EventFunc),
		}
	}
	return event
}

// Register 往一个分组中注册事件
func (e *SEvent) Register(group string, callback EventFunc) {
	e.Lock()
	defer e.Unlock()
	e.list[group] = append(e.list[group], callback)
}

// Call 回调一个分组的事件
func (e *SEvent) Call(group string, ctx context.Context, args ...interface{}) {
	if events, ok := e.list[group]; ok {
		for _, f := range events {
			f(ctx, args...)
		}
	}
}

// Remove 移动一个分组的事件
func (e *SEvent) Remove(group string) {
	delete(e.list, group)
}

// Clear 清空事件列表
func (e *SEvent) Clear() {
	e.list = make(map[string][]EventFunc)
}
