/*
Copyright © 2022 SUSE LLC

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

package clients

import (
	"context"

	elmscheme "github.com/rancher/elemental-operator/pkg/generated/clientset/versioned/scheme"
	capicontrollers "github.com/rancher/elemental-operator/pkg/generated/controllers/cluster.x-k8s.io"
	capi "github.com/rancher/elemental-operator/pkg/generated/controllers/cluster.x-k8s.io/v1beta1"
	"github.com/rancher/elemental-operator/pkg/generated/controllers/elemental.cattle.io"
	elmcontrollers "github.com/rancher/elemental-operator/pkg/generated/controllers/elemental.cattle.io/v1beta1"
	"github.com/rancher/elemental-operator/pkg/generated/controllers/fleet.cattle.io"
	fleetcontrollers "github.com/rancher/elemental-operator/pkg/generated/controllers/fleet.cattle.io/v1alpha1"
	"github.com/rancher/elemental-operator/pkg/generated/controllers/management.cattle.io"
	ranchercontrollers "github.com/rancher/elemental-operator/pkg/generated/controllers/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/clients"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	corev1Typed "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
)

const SystemNamespace = "cattle-system"

type Clients struct {
	*clients.Clients
	Fleet     fleetcontrollers.Interface
	Elemental elmcontrollers.Interface
	Events    corev1Typed.EventInterface
	Rancher   ranchercontrollers.Interface
	CAPI      capi.Interface
}

// EventRecorder creates an event recorder associated to a controller nome for the schema (arbitrary)
func (c *Clients) EventRecorder(name string) record.EventRecorder {
	// Create event broadcaster
	utilruntime.Must(elmscheme.AddToScheme(scheme.Scheme))
	logrus.Info("Creating event broadcaster for " + name)
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(logrus.Infof)
	eventBroadcaster.StartRecordingToSink(&corev1Typed.EventSinkImpl{Interface: c.Events})
	return eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: name})
}

func NewFromConfig(restConfig *rest.Config) (*Clients, error) {
	c, err := clients.NewFromConfig(restConfig, nil)
	if err != nil {
		return nil, err
	}

	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	opts := &generic.FactoryOptions{
		SharedControllerFactory: c.SharedControllerFactory,
	}
	return &Clients{
		Clients:   c,
		Events:    kubeClient.CoreV1().Events(""),
		Fleet:     fleet.NewFactoryFromConfigWithOptionsOrDie(restConfig, opts).Fleet().V1alpha1(),
		Elemental: elemental.NewFactoryFromConfigWithOptionsOrDie(restConfig, opts).Elemental().V1beta1(),
		CAPI:      capicontrollers.NewFactoryFromConfigWithOptionsOrDie(restConfig, opts).Cluster().V1beta1(),
		Rancher:   management.NewFactoryFromConfigWithOptionsOrDie(restConfig, opts).Management().V3(),
	}, nil
}

func (c *Clients) Start(ctx context.Context) error {
	return c.SharedControllerFactory.Start(ctx, 5)
}
