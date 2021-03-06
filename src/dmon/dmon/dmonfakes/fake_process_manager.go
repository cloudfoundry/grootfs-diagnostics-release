// Code generated by counterfeiter. DO NOT EDIT.
package dmonfakes

import (
	"dmon/dmon"
	"os/exec"
	"sync"
)

type FakeProcessManager struct {
	SpawnProcessStub        func(cmd *exec.Cmd) (int, error)
	spawnProcessMutex       sync.RWMutex
	spawnProcessArgsForCall []struct {
		cmd *exec.Cmd
	}
	spawnProcessReturns struct {
		result1 int
		result2 error
	}
	spawnProcessReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	WaitStub        func(pid int) (int, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
		pid int
	}
	waitReturns struct {
		result1 int
		result2 error
	}
	waitReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcessManager) SpawnProcess(cmd *exec.Cmd) (int, error) {
	fake.spawnProcessMutex.Lock()
	ret, specificReturn := fake.spawnProcessReturnsOnCall[len(fake.spawnProcessArgsForCall)]
	fake.spawnProcessArgsForCall = append(fake.spawnProcessArgsForCall, struct {
		cmd *exec.Cmd
	}{cmd})
	fake.recordInvocation("SpawnProcess", []interface{}{cmd})
	fake.spawnProcessMutex.Unlock()
	if fake.SpawnProcessStub != nil {
		return fake.SpawnProcessStub(cmd)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.spawnProcessReturns.result1, fake.spawnProcessReturns.result2
}

func (fake *FakeProcessManager) SpawnProcessCallCount() int {
	fake.spawnProcessMutex.RLock()
	defer fake.spawnProcessMutex.RUnlock()
	return len(fake.spawnProcessArgsForCall)
}

func (fake *FakeProcessManager) SpawnProcessArgsForCall(i int) *exec.Cmd {
	fake.spawnProcessMutex.RLock()
	defer fake.spawnProcessMutex.RUnlock()
	return fake.spawnProcessArgsForCall[i].cmd
}

func (fake *FakeProcessManager) SpawnProcessReturns(result1 int, result2 error) {
	fake.SpawnProcessStub = nil
	fake.spawnProcessReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessManager) SpawnProcessReturnsOnCall(i int, result1 int, result2 error) {
	fake.SpawnProcessStub = nil
	if fake.spawnProcessReturnsOnCall == nil {
		fake.spawnProcessReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.spawnProcessReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessManager) Wait(pid int) (int, error) {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
		pid int
	}{pid})
	fake.recordInvocation("Wait", []interface{}{pid})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub(pid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.waitReturns.result1, fake.waitReturns.result2
}

func (fake *FakeProcessManager) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeProcessManager) WaitArgsForCall(i int) int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return fake.waitArgsForCall[i].pid
}

func (fake *FakeProcessManager) WaitReturns(result1 int, result2 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessManager) WaitReturnsOnCall(i int, result1 int, result2 error) {
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeProcessManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.spawnProcessMutex.RLock()
	defer fake.spawnProcessMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProcessManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ dmon.ProcessManager = new(FakeProcessManager)
