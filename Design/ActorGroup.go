package Design

import (
	"fmt"
	"sync"
	"time"
)

type ActorGroup struct {
	timeout time.Duration
	actors  []*Actor
}

func (ag *ActorGroup) SetRunTimeout(timeout time.Duration) {
	ag.timeout = timeout
}

func (ag *ActorGroup) Add(actors ...*Actor) {
	for _, actor := range actors {
		ag.actors = append(ag.actors, actor)
	}

}

func (ag *ActorGroup) Run() error {
	if len(ag.actors) == 0 {
		return nil
	}

	var wg sync.WaitGroup

	for _, actor := range ag.actors {
		wg.Add(1)
		go actor.RunWithTimeout(ag.timeout, &wg)
	}

	wg.Wait()

	return nil

}

type Actor struct {
	Name       string
	Execute    func(inputParam interface{}) error
	IsTimeout  bool
	InputParam interface{}
	RetErr     error
}

func (a *Actor) RunWithTimeout(timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	if a.Execute == nil {
		return
	}

	// 设置默认的运行时间为 300s
	if timeout == 0 {
		timeout = 300 * time.Second
	}

	errChan := make(chan error)

	go func() {
		errChan <- a.Execute(a.InputParam)
	}()

	select {
	case a.RetErr = <-errChan:
		a.IsTimeout = false
	case <-time.After(timeout):
		a.RetErr = fmt.Errorf("execute timeout")
		a.IsTimeout = true

	}

	return
}
