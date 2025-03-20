/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"
	"testing"

	"github.com/vmware-tanzu/velero/pkg/nodeagent"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLoadConcurrency_DeepCopyInto(t *testing.T) {
	// Create a source LoadConcurrency instance
	source := &LoadConcurrency{
		nodeagent.LoadConcurrency{
			GlobalConfig: 3,
			PerNodeConfig: []nodeagent.RuledConfigs{
				nodeagent.RuledConfigs{
					Number: 2,
					NodeSelector: v1.LabelSelector{
						MatchLabels: map[string]string{
							"test": "ting",
						},
					},
				},
			},
		},
	}
	
	// Create a destination LoadConcurrency instance
	destination := &LoadConcurrency{}
	
	// Call DeepCopyInto
	source.DeepCopyInto(destination)
	
	// Verify the copy is correct using DeepEqual
	if !reflect.DeepEqual(source, destination) {
		t.Errorf("DeepCopyInto() failed to create an equal copy")
	}
	
	// Verify that source and destination are different objects
	if source == destination {
		t.Errorf("DeepCopyInto() did not create a separate instance")
	}
	
	// Create a DeepCopy using the DeepCopy method (which should call DeepCopyInto)
	// This tests the full deep copy functionality
	copy := source.DeepCopy()
	
	// Verify the copy is correct
	if !reflect.DeepEqual(source, copy) {
		t.Errorf("DeepCopy() failed to create an equal copy")
	}
	
	// Verify that source and copy are different objects
	if source == copy {
		t.Errorf("DeepCopy() did not create a separate instance")
	}
	
	// Test that modifying destination doesn't affect source
	// Store original value for comparison
	originalGlobalConfig := source.GlobalConfig
	originalNodeSelectorValue := source.PerNodeConfig[0].NodeSelector.MatchLabels["test"]
	
	// Modify destination
	destination.GlobalConfig = 999
	destination.PerNodeConfig[0].NodeSelector.MatchLabels["test"] = "modified"
	
	// Verify source is unchanged
	if source.GlobalConfig != originalGlobalConfig {
		t.Errorf("Source was modified when destination was changed. Expected GlobalConfig: %d, got: %d", 
			originalGlobalConfig, source.GlobalConfig)
	}
	
	if source.PerNodeConfig[0].NodeSelector.MatchLabels["test"] != originalNodeSelectorValue {
		t.Errorf("Source was modified when destination was changed. Expected NodeSelector label: %s, got: %s", 
			originalNodeSelectorValue, source.PerNodeConfig[0].NodeSelector.MatchLabels["test"])
	}
	
	// Also test with DeepCopy
	copy.GlobalConfig = 888
	copy.PerNodeConfig[0].NodeSelector.MatchLabels["test"] = "also-modified"
	
	// Verify source is still unchanged
	if source.GlobalConfig != originalGlobalConfig {
		t.Errorf("Source was modified when copy was changed. Expected GlobalConfig: %d, got: %d", 
			originalGlobalConfig, source.GlobalConfig)
	}
	
	if source.PerNodeConfig[0].NodeSelector.MatchLabels["test"] != originalNodeSelectorValue {
		t.Errorf("Source was modified when copy was changed. Expected NodeSelector label: %s, got: %s", 
			originalNodeSelectorValue, source.PerNodeConfig[0].NodeSelector.MatchLabels["test"])
	}
}
