# Operator

The Elemental operator is responsible for managing the Elemental versions
and maintaining a machine inventory to assist with edge or baremetal installations.

## Installation

The Elemental operator can be added to a cluster running Rancher Multi Cluster
Management server.  It is a helm chart and can be installed as follows:

```bash
helm -n cattle-elemental-operator-system install --create-namespace elemental-operator https://github.com/rancher/elemental-operator/releases/download/v0.1.0/elemental-operator-0.1.0.tgz
```

## Managing Upgrades

The Elemental operator will manage the upgrade of the local cluster where the operator
is running and also any downstream cluster managed by Rancher Multi-Cluster
Manager.

### ManagedOSImage

The ManagedOSImage kind used to define what version of Elemental should be
running on each node. The simplest example of this type would be to change
the image of the local nodes.

```bash
kubectl edit -n fleet-default default-os-image
```
```yaml
apiVersion: elemental.cattle.io/v1beta1
kind: ManagedOSImage
metadata:
  name: default-os-image
  namespace: fleet-default
spec:
  osImage: rancher/elemental:v0.0.0
```

A `ManagedOSImage` can also use a `ManagedOSVersion` to drive upgrades. 
To use a `ManagedOSVersion` specify a `managedOSVersionName`, as `osImage` takes precedence, mind to set back as empty:

```bash
kubectl edit -n fleet-default default-os-image
```
```yaml
apiVersion: elemental.cattle.io/v1beta1
kind: ManagedOSImage
metadata:
  name: default-os-image
  namespace: fleet-local
spec:
  osImage: ""
  managedOSVersionName: "version-name"
```


#### Reference

Below is reference of the full type

```yaml
apiVersion: elemental.cattle.io/v1beta1
kind: ManagedOSImage
metadata:
  name: arbitrary
  
  # There are two special namespaces to consider.  If you wish to manage
  # nodes on the local cluster this namespace must be `fleet-local`. If
  # you wish to manage nodes in Rancher MCM managed clusters then the
  # namespace must match the namespace of the clusters.provisioning.cattle.io resource
  # which is typically fleet-default.
  namespace: fleet-default
spec:
  # The image name to pull for the OS. Overrides managedOSVersionName when specified
  osImage: rancher/os2:v0.0.0

  # The ManagedOSVersion to use for the upgrade
  managedOSVersionName: ""
  
  # The selector for which nodes will be select.  If null then all nodes
  # will be selected
  nodeSelector:
    matchLabels: {}
    
  # How many nodes in parallel to update.  If empty the default is 1 and
  # if set to 0 the rollout will be paused
  concurrency: 2
    
  # Arbitrary action to perform on the node prior to upgrade
  prepare:
    image: ubuntu
    command: ["/bin/sh"]
    args: ["-c", "true"]
    env:
    - name: TEST_ENV
      value: testValue
      
  # Parameters to control the drain behavior.  If null no draining will happen
  # on the node.
  drain:
    # Refer to kubectl drain --help for the definition of these values
    timeout: 5m
    gracePeriod: 5m
    deleteLocalData: false
    ignoreDaemonSets: true
    force: false
    disableEviction: false
    skipWaitForDeleteTimeout: 5
    
  # Which clusters to target
  # This is used if you are running Rancher MCM and managing
  # multiple clusters.  The syntax of this field matches the
  # Fleet targets and is described at https://fleet.rancher.io/gitrepo-targets/
  clusterTargets: []

  # Overrides the default container created for running the upgrade with a custom one
  # This is optional and used only if specific upgrading mechanisms needs to be applied
  # in place of the default behavior.
  # The image used here overrides ones specified in osImage, depending on the upgrade strategy.
  upgradeContainer:
    image: ubuntu
    command: ["/bin/sh"]
    args: ["-c", "true"]
    env:
    - name: TEST_ENV
      value: testValue
```

## Inventory Management

The Elemental operator can hold an inventory of machines and
the mapping of the machine to it's configuration and assigned cluster.

### MachineInventory

#### Reference

```yaml
apiVersion: elemental.cattle.io/v1beta1
kind: MachineInventory
metadata:
  name: machine-a
  # The namespace must match the namespace of the cluster
  # assigned to the clusters.provisioning.cattle.io resource
  namespace: fleet-default
spec:
  # The cluster that this machine is assigned to
  clusterName: some-cluster
  # The hash of the TPM EK public key. This is used if you are
  # using TPM2 to identifiy nodes.  You can obtain the TPM by
  # running `rancherd get-tpm-hash` on the node. Or nodes can
  # report their TPM hash by using the MachineRegister
  tpm: d68795c6192af9922692f050b...
  # Generic SMBIOS fields that are typically populated with
  # the MachineRegister approach
  smbios: {}
  # A reference to a secret that contains a shared secret value to
  # identify a node.  The secret must be of type "elemental.cattle.io/token"
  # and have on field "token" which is the value of the shared secret
  machineTokenSecretName: some-secret-name
  # Arbitrary cloud config that will be added to the machines cloud config
  # during the rancherd bootstrap phase.  The one important field that should
  # be set is the role.
  config:
    role: server
```

### MachineRegistration

#### Reference

```yaml
apiVersion: elemental.cattle.io/v1beta1
kind: MachineRegistration
metadata:
  name: machine-registration
  # The namespace must match the namespace of the cluster
  # assigned to the clusters.provisioning.cattle.io resource
  namespace: fleet-default
spec:
  # Labels to be added to the created MachineInventory object
  machineInventoryLabels: {}
  # Annotations to be added to the created MachineInventory object
  machineInventoryAnnotations: {}
  # The cloud config that will be used to provision the node
  cloudConfig: {}
```
