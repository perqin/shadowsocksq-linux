// TODO: Support thread safety

package process

import (
	"os"
	"syscall"
)

type ProcessesService interface {
	StartOrRestartProcess(id int, binary string, args []string)
    StopProcess(id int)
	StopAllProcesses()
}

type processesServiceImpl struct {
	Processes	map[int]Process
}

var processesService ProcessesService

func init() {
	processesService = &processesServiceImpl{
		Processes: make(map[int]Process),
	}
}

func GetProcessesService() ProcessesService {
	return processesService
}

func (s *processesServiceImpl) StartOrRestartProcess(id int, binary string, args []string) {
	s.stopProcessLocked(id)
	// Create and start new process
	process := Process{}
	osProcess, err := os.StartProcess(binary, append([]string{binary}, args...), &os.ProcAttr{})
	if err == nil {
		process.OsProcess = osProcess
		s.Processes[id] = process
	}
}

func (s *processesServiceImpl) StopProcess(id int) {
	s.stopProcessLocked(id)
}

func (s *processesServiceImpl) StopAllProcesses() {
	for key := range s.Processes {
		s.stopProcessLocked(key)
	}
}

func (s *processesServiceImpl) stopProcessLocked(id int) {
	process, exist := s.Processes[id]
	if exist {
		// Restart process with new args and binary, so stop it first
		process.OsProcess.Signal(syscall.SIGINT)
		process.OsProcess.Wait()
		delete(s.Processes, id)
	}
}
