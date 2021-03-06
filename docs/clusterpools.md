# Cluster Pools

## Overview

Hive exposes a `ClusterPool` API which allows users to maintain a pool of "hot"
precreated `ClusterDeployments`, ready to be claimed when needed. The pool size
can be configured and Hive will attempt to maintain that set number of
clusters.

When a user needs a cluster they create a `ClusterClaim` resource which will be
filled immediately with details on where to find their cluster.  (or as soon as
a cluster is available in the pool if none were presently available). Once
claimed, a cluster is removed from the pool and a new one will be created to
replace it. Claimed clusters never return to the pool, they are intended to be
destroyed when no longer needed. The `ClusterClaim.Spec.Namespace` will be
populated once the claim has been filled, and the `ClusterDeployment` will be
present in that namespace with an identical name.

`ClusterPools` are a namespaced resource, and can be used to centralize billing
by using a namespace limited to a team via Kubernetes RBAC. All clusters in the
pool will use the same set of cloud credentials specified in the platform for
the pool. `ClusterClaims` must be created in the same namespace as their
`ClusterPool`, but each actual `ClusterDeployment` is given its own namespace.
The user who claims a cluster can be given RBAC to their clusters namespace to
prevent anyone else from being able to access it.

Presently once a `ClusterDeployment` is ready, it will be
[hibernated](./hibernating-clusters.md) automatically. Once claimed it will be
automatically resumed, meaning that the typical time to claim a cluster and be
ready to go is in the 2-5 minute range while the cluster starts up. An option
to keep the clusters in the pool always running so they can instantly be used
may be added in the future.

When done with a cluster, users can just delete their `ClusterClaim` and the
`ClusterDeployment` will be automatically deprovisioned. An optional
`ClusterClaim.Spec.Lifetime` can be specified after which a cluster claim will
automatically be deleted. The namespace created
for each cluster will eventually be cleaned up once deprovision has finished.

Note that at present, the shared credentials used for a pool will be visible
in-cluster. This may improve in the future for some clouds.

## Supported Cloud Platforms

`ClusterPool` currently supports the following cloud platforms:

  * AWS
  * Azure
  * GCP

## Sample Cluster Pool

```yaml
apiVersion: hive.openshift.io/v1
kind: ClusterPool
metadata:
  name: openshift-46-aws-us-east-1
  namespace: hive
spec:
  baseDomain: new-installer.openshift.com
  imageSetRef:
    name: openshift-4.6
  platform:
    aws:
      credentialsSecretRef:
        name: hive-team-aws-creds
      region: us-east-1
  pullSecretRef:
    name: hive-team-pull-secret
  size: 1
```

## Sample Cluster Claim

```yaml
apiVersion: hive.openshift.io/v1
kind: ClusterClaim
metadata:
  name: dgood46
  namespace: hive
spec:
  clusterPoolName: openshift-46-aws-us-east-1
  lifetime: 8h
  namespace: openshift-46-aws-us-east-1-j495p # populated by Hive once claim is filled and should not be set by the user on creation
status:
  conditions:
  - lastProbeTime: "2020-11-05T14:49:26Z"
    lastTransitionTime: "2020-11-05T14:49:26Z"
    message: Cluster claimed
    reason: ClusterClaimed
    status: "False"
    type: Pending
```

## Managing admins for Cluster Pools

Role bindings in the **namespace** of a `ClusterPool` that bind to the Cluster Role `hive-cluster-pool-admin`
are used to provide the **subjects** same permission in the namespaces created for various clusterprovisions for the cluster pool.
This allows operators to define adminstrators for a `ClusterPool` allowing them visibility to all the resources created for it. This is
most useful to debug `ClusterProvisions` associated with the pool that have failed and therefore cannot be claimed.

NOTE: You can only define such administrators for the entire namespace and not a specific `ClusterPool`.

To make any `User` or `Group` `hive-cluster-pool-admin` for a namespace you can,

```sh
oc -n <namespace> adm policy add-role-to-group hive-cluster-pool-admin <user>
```

or,

```sh
oc -n <namespace> adm policy add-role-to-group hive-cluster-pool-admin <group>
```

## Install Config Template

To control parts of the cluster deployments that are not directly supported by Hive, such as controlPlane Nodes and types, you can load a valid `install-config.yaml` which will be passed directly to the openshift-installer, only updating `metadata.name` and `baseDomain`

Load the install-config.yaml template as a secret (assuming that the `install-config.yaml` you want to use as a template is in the active directory)

```bash
kubectl  -n hive create secret generic my-install-config-template --from-file=install-config.yaml=./install-config.yaml
```

With this secret created, you can create a pool that references the install config secret template

```yaml
apiVersion: hive.openshift.io/v1
kind: ClusterPool
metadata:
  name: testpool
  namespace: hive
spec:
  baseDomain: hive.mytests.io
  imageSetRef:
    name: openshift-v4.5.13
  installConfigSecretTemplateRef: 
    name: my-install-config-template
  skipMachinePools: true
  platform:
    aws:
      credentialsSecretRef:
        name: global-aws-creds
      region: eu-west-1
  size: 1
```

**Note** When you use installConfigSecretTemplate you will most likely want to disable MachinePools, so that Hive does not reconcile away from the machine config specified in install-config.yaml
