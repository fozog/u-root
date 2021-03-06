// Copyright 2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"
)

var (
	RWLock   sync.RWMutex
	Registry = make(map[string]uint)
)

func read(serviceName string) (port uint, exists bool) {
	RWLock.RLock()
	defer RWLock.RUnlock()
	port, exists = Registry[serviceName]
	return
}

func register(serviceName string, portNum uint) error {
	RWLock.Lock()
	defer RWLock.Unlock()
	_, exists := Registry[serviceName]
	if exists {
		return fmt.Errorf("%v already exists", serviceName)
	}
	Registry[serviceName] = portNum
	return nil
}

func unregister(serviceName string) {
	RWLock.Lock()
	defer RWLock.Unlock()
	delete(Registry, serviceName)
}

func snapshotRegistry() map[string]uint {
	RWLock.RLock()
	defer RWLock.RUnlock()
	snapshot := make(map[string]uint)
	for name, port := range Registry {
		snapshot[name] = port
	}
	return snapshot
}

func main() {
	startServer()
}
