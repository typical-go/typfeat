package feato

type Manager struct {
	ToggleRouter ToggleRouter
}

type RunFunc func() error

// Run run functions according feature and toggleRouter
func (m *Manager) Run(feature *Feature, runFuncs ...RunFunc) error {
	index := m.ToggleRouter.Route(feature)
	if len(runFuncs) <= int(index) {
		return ErrOutOfRunFunctionsIndex
	}

	return runFuncs[index]()

}
