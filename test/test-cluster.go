package test

import (
	"os"
	"os/exec"
	"path/filepath"
)

// testCluster represents abstraction that enables Docker containers management
type testCluster struct {
	path string
}

// newTestCluster is a constructor function for testCluster instance
func newTestCluster(path string) (*testCluster, error) {
	if path == "" {
		workDir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path = filepath.Dir(workDir)
	}

	if err := os.Chdir(path); err != nil {
		return nil, err
	}

	return &testCluster{path: path}, nil
}

// start starts required Docker containers
func (t *testCluster) start() ([]byte, error) {
	return exec.Command("make", "run-docker").CombinedOutput()
}

// stop stop and destroys running Docker containers
func (t *testCluster) stop() ([]byte, error) {
	return exec.Command("make", "destroy-docker").CombinedOutput()
}
