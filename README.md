# bitfusion-license-manager-for-kubernetes

## Overview
Microservice to manage [Bitfusion FlexDirect](https://bitfusion.io/product/flexdirect/) licenses
on [Kubernetes](https://kubernetes.io).

The microservice consists of two components:

1. A [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset) that manages and reservices bitfusion pre-created licenses; and
1. A [Service](https://kubernetes.io/docs/concepts/services-networking/service) and supporting [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment) that provides for a remote mutex lock request.

The gist of operations is as follows:

The DaemonSet provides that a License Holder process (pod in this case) executes on all worker
nodes of a cluster into which it is deployed. The process therein has a
responsibility to request a mutex lock on a [Persistent
Volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes) that
contains a set of Bitfusion licenses, one per subdirectory on that volume.

Once a lock is obtained, the License Holder process scans the subdirectories to
find a license that is not currently locked by another License Holder (pod). If
one is found, the process locks (reserves) the license for its use and copies
it to the worker node's /etc/bitfusionio directory.

The reason for copying the license to the /etc/bitfusionio directory is so that
pod containers needing to utilize FlexDirect for access to GPUs need do nothing
more than mount the /etc/bitfusionio as a
[hostPath](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath) to
the container's /etc/bitfusionio directory to be appropriately licensed to take
advantage of the GPU.

## Try it out
To try out bitfusion-license-manager-for-kubernetes, it must be built and
executed in a suitable environment. Details are [further
below](#build-and-run).

### Prerequisites
There are some basic requirements in order build and use the
bitfusion-license-manager-for-kubernetes microservices:

1. A kubernetes cluster sufficient to run bitfusion-license-manager-for-kubernetes and its helper [license locking deployment](deployment/kubernetes/token).
2. A Docker registry and relevant credentials for pushing and pulling containers;
3. A build environment that includes golang 12.x (or greater), gcc compiler
and [Docker](https://docker.com) installed and running.

#### Kubernetes Cluster
Running the bitfusion-license-manager-for-kubernetes generally targets a kubenernetes cluster.
It runs in the default namespace unless you specify another, e.g.:

    kubectl -n bitfusion ...

when deploying.

#### Docker
Building the microservice requires that the local build machine has a valid,
running Docker setup. To see if you have such a system, you can run the
following command:

    docker images

If that works, your build should succeed absent other issues not related to
Docker.

### Build and Run
Building the code involves very few steps and there are some CI/CD options to
help out.

1. Get the code
2. Build the code
3. Deploy the code

#### Get the Code
To get the code, get it similarly to the following:

    git clone https://github.com/vmware/bitfusion-license-manager-for-kubernetes
    cd bitfusion-license-manager-for-kubernetes
    git submodule init
    git submodule update --recursive

#### Build the Code
The project includes a Makefile for use in building the
bitfusion-license-manager-for-kubernetes microservice. To build a docker
container, you must provide a container name for pushing to a registry.

    make clean && make

That will build two containers on the local Docker service: flexdirect and
tokenmgr. The former is the container used by the DaemonSet to establish pods
on all worker nodes. The latter is the container used for the locking (mutex)
service that the flexdirect pods ultimate request in order to have sole access
to scan the licenses.

#### Run
To run the microservice, you should register the two new containers with your
Docker registry, such as [Harbor](https://github.com/goharbor/harbor). For example:

    docker push flexdirect:1.0.0
    docker push token:1.0.0

Note also that the sample files provided in the [deployment
directory](deployment/kubernetes) assume an open registry (i.e., no required
secrets). You should also create (or otherwise reqeuest) a namespace in which
to place the various Kubernetes objects. The process to run the set (assuming
kubernetes has access to a Docker registry that contains the containers built
above) is as follows (which provides a namespace "bitfusion"):

    kubectl apply -f examples/kubernetes/namespace.yaml
    kubectl -n bitfusion apply -f examples/kubernetes/tokenmgr/service.yaml
    kubectl -n bitfusion apply -f examples/kubernetes/tokenmgr/deployment.yaml
    kubectl -n bitfusion apply -f examples/kubernetes/daemonset/configmap.yaml
    kubectl -n bitfusion apply -f examples/kubernetes/daemonset/persistentVolume.yaml
    kubectl -n bitfusion apply -f examples/kubernetes/daemonset/persistentVolumeClaim.yaml
    kubectl -n bitfusion apply -f examples/kubernetes/daemonset/daemonset.yaml

It is a good idea to check on the pods that are running at that point to
assure they are properly executing:

    kubtctl -n bitfusion get po

All pods should be be ready and running normally.

## API and HTML Documentation
The microservice provides two mechanisms for interacting with the service:

1. an API for requesting and releaseing mutex locks; and
1. a web interface for visually doing the same.

### REST API
Upon running the service, a REST API exists within the cluster at
http://tokenmgr.bitfusion:8080/api/reminders and paths further thereafter
pursuant to the pattern:

- GET /api/tokenmgr/lock/:timeout

Where "timeout" is a time in seconds for which a unique lock is requested, but
will timeout at :timeout seconds if the mutex lock was not released by another
holder.  Returns all reminders currently in the database.

The return (body) of the request contains a unique id for the lock, if
obtained:

    {
      "id": "bea9eccb-f571-498d-ba46-42a5cbe25bd7"
    }


- POST /api/tokenmgr/release/:id

Releases the mutex lock with the :id value provided by the lock API.

-GET /api/stats

Returns a JSON body that provides a list of URLs serviced by the microservice
and how many times that URL was hit.

#### The HTML Interface
To reach the HTML interface (given the same sample as above), browse to:
http://tokenmgr.bitfusion/html/tmpl/index and the bulk  of the HTML paths are
available from that page or others as appropriate given traversal of the 'site.'

## Releases & Major Branches

## Contributing

The bitfusion-license-manager-for-kubernetes project team welcomes
contributions from the community. Before you start working with
bitfusion-license-manager-for-kubernetes, please read our [Developer
Certificate of Origin](https://cla.vmware.com/dco). All contributions to this
repository must be signed as described on that page.  Your signature certifies
that you wrote the patch or have the right to pass it on as an open-source
patch. For more detailed information, refer to
[CONTRIBUTING.md](CONTRIBUTING.md).



## License
Copyright (c) 2019 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: [https://spdx.org/licenses/MIT.html](https://spdx.org/licenses/MIT.html)
