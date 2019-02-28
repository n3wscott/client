// Copyright © 2019 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var eventsListPrintFlags *genericclioptions.PrintFlags

// NewEventsListCommand represents the list command for events.
func NewEventsListCommand(p *KnParams) *cobra.Command {

	eventsListPrintFlags := genericclioptions.NewPrintFlags("").WithDefaultOutput(
		"jsonpath={range .items[*]}{.metadata.name}{\"\\n\"}{end}")
	eventsListCommand := &cobra.Command{
		Use:   "list",
		Short: "List available events.",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace := cmd.Flag("namespace").Value.String()
			broker := cmd.Flag("namespace").Value.String()

			printer, err := eventsListPrintFlags.ToPrinter()
			if err != nil {
				return err
			}

			objs, err := doEventsList(namespace, broker)
			if err != nil {
				return err
			}

			if err = printer.PrintObj(objs, cmd.OutOrStdout()); err != nil {
				return err
			}
			return nil
		},
	}
	eventsListPrintFlags.AddFlags(eventsListCommand)
	return eventsListCommand
}

func doEventsList(namespace, broker string) (runtime.Object, error) {

	eventFoo := unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "dev.knative.eventing/v1alpha1",
			"kind":       "eventType",
			"metadata": map[string]interface{}{
				"name": "foo",
			},
			"status": map[string]interface{}{
				"extra": "fields",
				"fooable": map[string]interface{}{
					"field1": "foo",
					"field2": "bar",
				},
			},
		}}

	eventBar := unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "dev.knative.eventing/v1alpha1",
			"kind":       "eventType",
			"metadata": map[string]interface{}{
				"name": "bar",
			},
			"status": map[string]interface{}{
				"extra": "fields",
				"fooable": map[string]interface{}{
					"field1": "foo",
					"field2": "bar",
				},
			},
		}}

	events := &unstructured.UnstructuredList{Items: []unstructured.Unstructured{eventFoo, eventBar}}

	return events, nil
}
