// Copyright 2023 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mountutils

import (
	mount "k8s.io/mount-utils"
	utilexec "k8s.io/utils/exec"
)

//go:generate $MOCKGEN -package mocks -destination=../../../pkg/driver/mocks/mountutils.go -source ../../../pkg/driver/mountutils/mountutils.go

// MountWrapper is the interface implemented by NodeMounter. A mix & match of
// functions defined in upstream libraries. (FormatAndMount from struct
// SafeFormatAndMount). Defined it explicitly so that it can be mocked.
type MountWrapper interface {
	mount.Interface
	FormatAndMount(source string, target string, fstype string, options []string) error
}

// NodeMounter implements MountWrapper.
// A superstruct of SafeFormatAndMount.
type NodeMounter struct {
	*mount.SafeFormatAndMount
}

func NewNodeMounter() (MountWrapper, error) {
	safeMounter, err := newSafeMounter()
	if err != nil {
		return nil, err
	}
	return &NodeMounter{safeMounter}, nil
}

// newSafeMounter returns a SafeFormatAndMount
func newSafeMounter() (*mount.SafeFormatAndMount, error) {
	return &mount.SafeFormatAndMount{
		Interface: mount.New(""),
		Exec:      utilexec.New(),
	}, nil
}