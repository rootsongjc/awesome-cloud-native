# Awesome Cloud Native![](https://ws1.sinaimg.cn/large/00704eQkgy1frzchbek3ej300w00wgle.jpg)

A curated list of awesome cloud native tools, softwares and tutorials. Inspired by **[awesome-go](https://github.com/avelino/awesome-go)**.

### Contributing

Please take a quick gander at the **[contribution guidelines](https://github.com/rootsongjc/awesome-cloud-native/blob/master/CONTRIBUTING.md)** first. View in [GitHub](https://github.com/rootsongjc/awesome-cloud-native). Thanks to all **[contributors](https://github.com/rootsongjc/awesome-cloud-native/graphs/contributors)**, you rock🤟!

### Contents

- Awesome Cloud Native
  - [AI](#ai)
  - [API Gateway](#api-gateway)
  - [Big Data](#big-data)
  - [Container Runtime](#container-runtime)
  - [CI-CD](#ci-cd)
  - [Database](#database)
  - [Data Science](#data-science)
  - [Fault Tolerant](#fault-tolerant)
  - [IoT](#iot)
  - [Kubernetes Operator](#kubernetes-operator)
  - [Logging](#logging)
  - [Message Broker](#message-broker)
  - [Miscellaneous](#miscellaneous)
  - [Monitoring](#monitoring)
  - [Network](#network)
  - [Observability](#observability)
  - [Orchestration and Scheduler](#orchestration-and-scheduler)
  - [PaaS](#paas)
  - [Proxy and Load Balancer](#proxy-and-load-balancer)
  - [RPC](#rpc)
  - [Security and Audit](#security-and-audit)
  - [Service Broker](#service-broker)
  - [Service Mesh](#service-mesh)
  - [Service Registry and Discovery](#service-registry-and-discovery)
  - [Serverless](#serverless)
  - [Storage](#storage)
  - [Tracing](#tracing)
  - [Tools](#tools)
  - [Tutorials](#tutorials)

### AI

- [allennlp](https://github.com/allenai/allennlp) - An open-source NLP research library, built on PyTorch. 
[http://www.allennlp.org](http://www.allennlp.org/)
- [caffe2](https://github.com/caffe2/caffe2) - Caffe2 is a lightweight, modular, and scalable deep learning framework. [https://caffe2.ai](https://caffe2.ai/)
- [h2o](https://github.com/h2oai/h2o-3) - Open Source Fast Scalable Machine Learning API For Smarter Applications (Deep Learning, Gradient Boosting, Random Forest, Generalized Linear Modeling (Logistic Regression, Elastic Net), K-Means, PCA, Stacked Ensembles...) [http://h2o.ai](http://h2o.ai/)
- [keras](https://github.com/keras-team/keras) - Keras is a high-level neural networks API, written in Python and capable of running on top of TensorFlow, CNTK, or Theano.
[http://keras.io/](http://keras.io/)
- [kubeflow](https://github.com/google/kubeflow) - Machine Learning Toolkit for Kubernetes.
[https://www.kubeflow.org](https://www.kubeflow.org/)
- [leaf](https://github.com/autumnai/leaf) - Open Machine Intelligence Framework for Hackers. (GPU/CPU) <http://autumnai.com/leaf/book>
- [PaddlePaddle](https://github.com/PaddlePaddle/Paddle) - PArallel Distributed Deep LEarning. （PaddlePaddle核心框架，高性能单机、分布式训练和跨平台部署）
[http://www.paddlepaddle.org](http://www.paddlepaddle.org/)
- [predictionio](https://github.com/apache/incubator-predictionio) - PredictionIO, a machine learning server for developers and ML engineers. Built on Apache Spark, HBase and Spray.[https://predictionio.incubator.apache…](https://predictionio.incubator.apache.org/)
- [pytorch](https://github.com/pytorch/pytorch) -An open source deep learning platform that provides a seamless path from research prototyping to production deployment.
[https://pytorch.org](https://pytorch.org)
- [tensorflow](https://github.com/tensorflow/tensorflow) - Computation using data flow graphs for scalable machine learning [http://tensorflow.org](http://tensorflow.org/)

### API Gateway

- [ambassador](https://github.com/datawire/ambassador) - Ambassador: a self-service API gateway for microservices built on Lyft Envoy [http://www.getambassador.io](http://www.getambassador.io/)
- [express-gateway](https://github.com/ExpressGateway/express-gateway) - A microservices API Gateway built on top of ExpressJS [https://www.express-gateway.io](https://www.express-gateway.io/)
- [kong](https://github.com/Mashape/kong) - 🐒 The Microservice API Gateway <https://getkong.org/install>
- [krakend](https://github.com/devopsfaith/krakend) - Ultra performant API Gateway with middlewares. <http://krakend.io/>
- [orange](https://github.com/sumory/orange) - OpenResty/Nginx Gateway for API Monitoring and Management. [http://orange.sumory.com](http://orange.sumory.com/)
- [ocelot](https://github.com/ThreeMammals/Ocelot) - .NET core API Gateway. [http://threemammals.com/ocelot](http://threemammals.com/ocelot)

### Big Data

- [fast-data-dev](https://github.com/Landoop/fast-data-dev) - Kafka Docker for development. Kafka, Zookeeper, Schema Registry, Kafka-Connect, Landoop Tools, 20+ connectors
- [spark](https://github.com/apache-spark-on-k8s/spark) - Apache Spark enhanced with native Kubernetes scheduler back-end
- [wallaroo](https://github.com/WallarooLabs/wallaroo) - Ultrafast and elastic data processing [https://www.wallaroolabs.com](https://www.wallaroolabs.com/)

### Container Runtime

- [clear-containers](https://github.com/clearcontainers) - OCI (Open Containers Initiative) compatible runtime using Virtual Machines
- [containerd](https://github.com/containerd/containerd) - An open and reliable container runtime [https://containerd.io](https://containerd.io/)
- [cri-containerd](https://github.com/containerd/cri-containerd) - Containerd-based implementation of Kubernetes Container Runtime Interface
- [cri-o](http://cri-o.io/) - Lightweight Container Runtime for Kubernetes
- [frakti](https://github.com/kubernetes/frakti) - The hypervisor-based container runtime for Kubernetes.
- [gvisor](https://github.com/google/gvisor) - Sandboxed Container Runtime
- [hyperd](https://github.com/hyperhq/hyperd) - HyperContainer Daemon [http://www.hypercontainer.io](http://www.hypercontainer.io/)
- [katacontainers](https://katacontainers.io/) - Kata Containers is a new open source project building extremely lightweight virtual machines that seamlessly plug into the containers ecosystem.
- [moby](https://github.com/moby/moby) - Moby Project - a collaborative project for the container ecosystem to assemble container-based systems[https://mobyproject.org/](https://mobyproject.org/)
- [pouch](https://github.com/alibaba/pouch) - Pouch is an open-source project created to promote the container technology movement.
- [railcar](https://github.com/oracle/railcar) - RailCar: Rust implementation of the Open Containers Initiative oci-runtime
- [rkt](https://github.com/rkt/rkt) - rkt is a pod-native container engine for Linux. It is composable, secure, and built on standards.

### CI-CD

- [argo](https://github.com/argoproj/argo) - Get stuff done with container-native workflows for Kubernetes.
- [binderhub](https://github.com/jupyterhub/binderhub) - Deterministically build docker images from a git repository + commit [https://binderhub.readthedocs.io](https://binderhub.readthedocs.io/)
- [buddy.works](https://buddy.works/) - Build, Test & Deploy Code in Seconds Continuous Delivery, simplified.
- [circleci](https://github.com/circleci) - Continuous Integration and Deployment
- [concourse](https://github.com/concourse/concourse) - BOSH release and development workspace for Concourse [https://concourse.ci](https://concourse.ci/)
- [cross-cloud](https://github.com/cncf/cross-cloud/) - Cross Cloud Continuous Integration [https://cncf.io](https://cncf.io/)
- [cyclone](https://github.com/caicloud/cyclone) - Powerful workflow engine and end-to-end pipeline solutions implemented with native Kubernetes resources.
- [drone](https://github.com/drone/drone) - Drone is a Continuous Delivery platform built on Docker, written in Go [https://drone.io](https://drone.io/)
- [gitLab-ci](https://about.gitlab.com/features/gitlab-ci-cd) - GitLab has integrated CI/CD pipelines to build, test, deploy, and monitor your code
- [gitkube](https://github.com/hasura/gitkube) - Gitkube: Build and deploy docker images to Kubernetes using git push. [https://gitkube.sh](https://gitkube.sh/)
- [hygieia](https://github.com/capitalone/Hygieia) - CapitalOne DevOps Dashboard <http://www.capitalone.io/Hygieia>
- [jenkins](http://jenkins.io) - The leading open source automation server, Jenkins provides hundreds of plugins to support building, deploying and automating any project.
- [jx](https://github.com/jenkins-x/jx) - A command line tool for installing and working with Jenkins X <http://jenkins-x.io/>
- [kenyata](https://github.com/spinnaker/kayenta) - Automated Canary Service
- [pipeline](https://github.com/banzaicloud/pipeline) - REST API to provision or reuse managed Kubernetes clusters in the cloud and deploy cloud native apps
- [skaffold](https://github.com/GoogleCloudPlatform/skaffold) - Easy and Repeatable Kubernetes Development
- [travis](https://github.com/travis-ci/travis-ci) - Free continuous integration platform for GitHub projects. [https://travis-ci.org](https://travis-ci.org/)
- [wercker](https://github.com/wercker/wercker) - The Wercker CLI can be used to execute pipelines locally for both local development and easy introspection. [http://wercker.com](http://wercker.com/)

### Database

- [arangodb](https://github.com/arangodb/arangodb) - ArangoDB is a native multi-model database with flexible data models for documents, graphs, and key-values. Build high performance applications using a convenient SQL-like query language or JavaScript extensions.
- [beringei](https://github.com/facebookincubator/beringei) - Beringei is a high performance, in-memory storage engine for time series data.
- [cockroachdb](https://github.com/cockroachdb/cockroach/) - CockroachDB - the open source, cloud-native SQL database. [https://www.cockroachlabs.com](https://www.cockroachlabs.com/)
- [couchdb](https://github.com/apache/couchdb) - Apache CouchDB is one of a new breed of database management systems.
- [etcd](https://github.com/coreos/etcd) - Distributed reliable key-value store for the most critical data of a distributed system[https://coreos.com/etcd/docs/latest/](https://coreos.com/etcd/docs/latest/)
- [influxdb](https://github.com/influxdata/influxdb) - Scalable datastore for metrics, events, and real-time analytics [https://influxdata.com](https://influxdata.com/)
- [leveldb](https://github.com/google/leveldb) - LevelDB is a fast key-value storage library written at Google that provides an ordered mapping from string keys to string values.
- [mehdb](https://github.com/mhausenblas/mehdb) - Educational Kubernetes-native NoSQL datastore using StatefulSet and persistent volumes [https://blog.openshift.com/kubernetes…](https://blog.openshift.com/kubernetes-statefulset-in-action/)
- [mongodb](https://github.com/mongodb/mongo) - MongoDB is an open source database that uses a document-oriented data model.
- [opentsdb](https://github.com/OpenTSDB/opentsdb) - A scalable, distributed Time Series Database. [http://opentsdb.net](http://opentsdb.net/)
- [redis](https://github.com/antirez/redis) - Redis is an in-memory database that persists on disk. The data model is key-value, but many different kind of values are supported: Strings, Lists, Sets, Sorted Sets, Hashes, HyperLogLogs, Bitmaps. [http://redis.io](http://redis.io/)
- [rethinkdb](https://github.com/rethinkdb/rethinkdb) - The open-source database for the realtime web. [https://rethinkdb.com](https://rethinkdb.com/)
- [sharding-sphere](https://github.com/sharding-sphere/sharding-sphere) - Distributed database middleware
- [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed NewSQL database compatible with MySQL protocol [https://pingcap.com](https://pingcap.com/)

### Data Science

- [pachyderm](https://github.com/pachyderm/pachyderm) - Reproducible Data Science at Scale! [http://pachyderm.io](http://pachyderm.io/)
- [wallaroo](https://github.com/WallarooLabs/wallaroo) - Ultrafast and elastic data processing [https://www.wallaroolabs.com](https://www.wallaroolabs.com/)

### DevOps

- [containerops](https://github.com/Huawei/containerops) - DevOps Orchestration Platform [https://cncf.build](https://cncf.build/)
- [crane](https://github.com/Dataman-Cloud/crane) - Yet another control plane based on docker built-in swarmkit [https://www.shurenyun.com/product-crane.html](https://www.shurenyun.com/product-crane.html)
- [fabric8](https://github.com/fabric8io/fabric8) - fabric8 is an open source microservices platform based on Docker, Kubernetes and Jenkins [http://fabric8.io/](http://fabric8.io/)
- [lastbackend](https://github.com/lastbackend/lastbackend) - Container orchestration with CI&CD, cli and amazing UI [https://lastbackend.com](https://lastbackend.com/)
- [spinnaker](https://github.com/spinnaker/spinnaker) - Spinnaker is an open source, multi-cloud continuous delivery platform for releasing software changes with high velocity and confidence. [http://www.spinnaker.io/](http://www.spinnaker.io/)
- [terraform](https://github.com/hashicorp/terraform) - Terraform is a tool for building, changing, and combining infrastructure safely and efficiently. [https://www.terraform.io/](https://www.terraform.io/)

### Fault Tolerant

- [chaosmonkey](https://github.com/Netflix/chaosmonkey) - Chaos Monkey is a resiliency tool that helps applications tolerate random instance failures.
- [concurrency-limits](https://github.com/Netflix/concurrency-limits) - Java Library that implements and integrates concepts from TCP congestion control to auto-detect concurrency limits to achieve optimal throughput with optimal latency.
- [hystrix](https://github.com/Netflix/Hystrix) - Hystrix is a latency and fault tolerance library designed to isolate points of access to remote systems, services and 3rd party libraries, stop cascading failure and enable resilience in complex distributed systems where failure is inevitable.
- [ratelimit](https://github.com/lyft/ratelimit) - Go/gRPC service designed to enable generic rate limit scenarios from different types of applications.

### IoT

- [eliot](https://github.com/ernoaapa/eliot) - Open source system for managing containerized applications in IoT device [http://eliot.run](http://eliot.run/)

### Kubernetes operator

- [keel](https://github.com/keel-hq/keel) - Kubernetes Operator to automate Helm, DaemonSet, StatefulSet & Deployment updates [https://keel.sh](https://keel.sh/)
- [kubevirt](https://github.com/kubevirt/kubevirt) - Kubernetes Virtualization Operator with API and runtime in order to define and manage virtual machines.
- [operator-sdk](https://github.com/operator-framework/operator-sdk) - SDK for building Kubernetes applications. Provides high level APIs, useful abstractions, and project scaffolding. <https://coreos.com/operators>
- [prometheus-operator](https://github.com/coreos/prometheus-operator) - Prometheus Operator creates/configures/manages Prometheus clusters atop Kubernetes<https://coreos.com/operators/prometheus>
- [spark-on-k8s-operator](https://github.com/GoogleCloudPlatform/spark-on-k8s-operator) - Kubernetes operator for managing the lifecycle of Apache Spark applications on Kubernetes.

### Logging

- [beats](https://github.com/elastic/beats) - Beats - Lightweight shippers for Elasticsearch & Logstash [https://www.elastic.co/products/beats](https://www.elastic.co/products/beats)
- [collectbeat](https://github.com/eBay/collectbeat) - Beats with discovery capabilities for environments like Kubernetes
- [elasticsearch](https://github.com/elastic/elasticsearch) - Open Source, Distributed, RESTful Search Engine [https://www.elastic.co/products/elast…](https://www.elastic.co/products/elasticsearch)
- [fluent-bit](https://github.com/fluent/fluent-bit) - Fast and Lightweight Log/Data Forwarder for Linux, BSD and OSX <http://fluentbit.io/>
- [fluentd-pilot](https://github.com/AliyunContainerService/fluentd-pilot) - Collect logs in docker containers
- [fluentd](https://github.com/fluent/fluentd) - Fluentd: Unified Logging Layer (project under CNCF) [http://www.fluentd.org/](http://www.fluentd.org/)
- [flume](http://flume.apache.org/) - Flume is a distributed, reliable, and available service for efficiently collecting, aggregating, and moving large amounts of log data.
- [heapster](https://github.com/kubernetes/heapster) - Compute Resource Usage Analysis and Monitoring of Container Clusters
- [log-pilot](https://github.com/AliyunContainerService/log-pilot) - Collect logs in docker containers
- [loki](https://github.com/grafana/loki) - Like Prometheus, but for logs.
- [telegraf](https://github.com/influxdata/telegraf) - The plugin-driven server agent for collecting & reporting metrics.

### Observability

- [envoy-ui](https://github.com/Nitro/envoy-ui) - Dead simple server-side UI for Envoy proxy (like HAproxy stats)
- [goldpinger](https://github.com/bloomberg/goldpinger) - Debugging tool for Kubernetes which tests and displays connectivity between nodes in the cluster.
- [istio-ui](https://github.com/jukylin/istio-ui) - Istio config management backend
- [kiali](https://github.com/kiali/kiali) - kiali project to help istio service mesh observability
- [kube-ops-view](https://github.com/hjacobs/kube-ops-view) - Kubernetes Operational View - read-only system dashboard for multiple K8s clusters [https://kubernetes-operational-view.readthedocs.io](https://kubernetes-operational-view.readthedocs.io/)
- [naftis](https://github.com/XiaoMi/naftis) - An excellent dashboard for Istio built with love.
- [vistio](https://github.com/nmnellis/vistio) - Visualize your Istio mesh using Netflix's Vizceral
- [vizceral](https://github.com/Netflix/vizceral) - WebGL visualization for displaying animated traffic graphs

### Orchestration and Scheduler

- [alameda](https://github.com/containers-ai/alameda) - Intelligent Resources Orchestrator for Kubernetes by using machine learning
- [blox](https://github.com/blox/blox) - Open source tools for building custom schedulers on Amazon ECS
- [compose](https://github.com/docker/compose) - Define and run multi-container applications with Docker
- [conductor](https://github.com/Netflix/conductor) - Conductor is a microservices orchestration engine - [https://netflix.github.io/conductor/](https://netflix.github.io/conductor/)
- [descheduler](https://github.com/kubernetes-incubator/descheduler) - Descheduler for Kubernetes [https://github.com/kubernetes-incubator/descheduler](https://github.com/kubernetes-incubator/descheduler)
- [domeos](https://github.com/domeos) - An enterprise application orchestration system based on docker by Sohu.com - http://domeos.org
- [fleet](https://github.com/coreos/fleet) - fleet ties together systemd and etcd into a distributed init system
- [kubernetes](https://github.com/kubernetes/kubernetes) - Production-Grade Container Scheduling and Management [http://kubernetes.io](http://kubernetes.io/)
- [marathon](https://github.com/mesosphere/marathon) - Deploy and manage containers (including Docker) on top of Apache Mesos at scale.[https://mesosphere.github.io/marathon/](https://mesosphere.github.io/marathon/)
- [mesos](https://github.com/apache/mesos) - Apache Mesos abstracts CPU, memory, storage, and other compute resources away from machines (physical or virtual), enabling fault-tolerant and elastic distributed systems to easily be built and run effectively.
- [serf](https://github.com/hashicorp/serf) - Service orchestration and management tool by hashicorp. [https://www.serf.io/](https://www.serf.io/)
- [service-fabric](https://github.com/Microsoft/service-fabric) - Service Fabric is a distributed systems platform for packaging, deploying, and managing stateless and stateful distributed applications and containers at large scale. [https://docs.microsoft.com/en-us/azure/service-fabric](https://docs.microsoft.com/en-us/azure/service-fabric/)
- [swan](https://github.com/Dataman-Cloud/swan) - A Distributed, Highly Available Mesos Scheduler, Inspired by the design of Google Borg [https://www.shurenyun.com/product-swa…](https://www.shurenyun.com/product-swan.html)
- [swarm](https://github.com/docker/swarm) - Swarm: a Docker-native clustering system

### PaaS

- [breeze](https://github.com/wise2c-devops/breeze) -  Wise2C ansible playbook for Kubernetes cluster installation
- [choerodon](https://github.com/choerodon/choerodon) - The open source PaaS for Kubernetes.[http://choerodon.io](http://choerodon.io/)
- [cloudfoundry](https://www.cloudfoundry.org) - Cloud Foundry is an open source, multi cloud application platform as a service (PaaS) governed by the Cloud Foundry Foundation.
- [conjure-up](https://github.com/conjure-up/conjure-up) - Deploying complex solutions, magically. [https://conjure-up.io](https://conjure-up.io/)
- [crossplane](https://github.com/crossplaneio/crossplane) - An Open Source Multicloud Control Plane [https://crossplane.io](https://crossplane.io/)
- [dc/os](https://github.com/dcos) - Datacenter Operating System
- [deis](https://github.com/deis/deis) - Deis v1, the CoreOS and Docker PaaS: Your PaaS. Your Rules. [https://deis.com/docs/](https://deis.com/docs/)
- [flynn](https://github.com/flynn/flynn) - A next generation open source platform as a service (PaaS) [https://flynn.io](https://flynn.io/)
- [kqeen](https://github.com/Mirantis/kqueen) - Kubernetes queen - cluster manager
- [kubernator](https://github.com/smpio/kubernator) - Alternative Kubernetes UI
- [kubesphere](https://github.com/kubesphere/kubesphere) - Enterprise Container Managent Platform https://kubesphere.io
- [opendcp](https://github.com/weibocom/opendcp) - Docker platform developed by weibo.com
- [openshift](https://github.com/openshift/origin) - Enterprise Kubernetes for Developers [http://www.openshift.org](http://www.openshift.org/)
- [rainbond](https://github.com/goodrain/rainbond) - Serverless PaaS , A new generation of easy-to-use cloud management platforms based on kubernetes.[http://www.rainbond.com](http://www.rainbond.com/)
- [rancher](https://github.com/rancher/rancher) - Complete container management platform [http://rancher.com](http://rancher.com/)
- [supergiant](https://github.com/supergiant/supergiant) - Automatically scale hardware and easily run stateful applications using Kubernetes. <https://supergiant.io/>
- [wayne](https://github.com/Qihoo360/wayne) - Web UI for Kubernetes multi-clusters
- [vamp](https://github.com/magneticio/vamp) - Vamp - canary releasing and autoscaling for microservice systems [http://vamp.io](http://vamp.io/)

### Proxy and Load Balancer

- [caddy](https://github.com/mholt/caddy) - Fast, cross-platform HTTP/2 web server with automatic HTTPS [https://caddyserver.com](https://caddyserver.com/)
- [contour](https://github.com/heptio/contour) - Contour is a Kubernetes ingress controller for Lyft's Envoy proxy.
- [envoy-docker-shim](https://github.com/Nitro/envoy-docker-shim) - Run Envoy in place of docker-proxy
- [envoy](https://github.com/lyft/envoy) - C++ front/service proxy [https://lyft.github.io/envoy](https://lyft.github.io/envoy)
- [gimbal](https://github.com/heptio/gimbal) - Heptio Gimbal is an ingress load balancing platform capable of routing traffic to multiple Kubernetes and OpenStack clusters. Built by Heptio in partnership with Actapio. [https://www.heptio.com](https://www.heptio.com/)
- [haproxy](https://github.com/haproxy/haproxy) - HAProxy is a free, very fast and reliable solution offering high availability, load balancing, and proxying for TCP and HTTP-based applications.
- [kEdge](https://github.com/improbable-eng/kedge) - kEdge - Kubernetes Edge Proxy for gRPC and HTTP Microservices
- [katran](https://github.com/facebookincubator/katran) - A high performance layer 4 load balancer
- [kong-ingress](https://github.com/koli/kong-ingress) - A Kubernetes Ingress for Kong
- [kong/kubernetes-ingress-controller](https://github.com/Kong/kubernetes-ingress-controller) - Deploy Kong in a native Kubernetes Ingress Controller <https://konghq.com/>
- [metallb](https://github.com/google/metallb) - A network load-balancer implementation for Kubernetes using standard routing protocols [https://metallb.universe.tf](https://metallb.universe.tf/)
- [nginx-kubernetes-ingress](https://github.com/nginxinc/kubernetes-ingress) - NGINX and NGINX Plus Ingress Controllers for Kubernetes
- [nginx](https://github.com/nginx/nginx) - nginx [engine x] is an HTTP and reverse proxy server, a mail proxy server, and a generic TCP/UDP proxy server, originally written by [Igor Sysoev](http://sysoev.ru/en/).
- [ribbon](https://github.com/Netflix/ribbon) - Ribbon is a Inter Process Communication (remote procedure calls) library with built in software load balancers. The primary usage model involves REST calls with various serialization scheme support.
- [skipper](https://github.com/zalando/skipper) - An HTTP router and reverse proxy for service composition, including use cases like Kubernetes Ingress
- [sofa-mosn](https://github.com/alipay/sofa-mosn) - SOFA MOSN is a modular observable smart proxy which can be used in service mesh deployed as a data plane sidecar.
- [traefik](https://github.com/containous/traefik) - Træfik, a modern reverse proxy [https://traefik.io](https://traefik.io/)
- [voyager](https://github.com/appscode/voyager) - ✈️️ Secure Ingress Controller for Kubernetes by <https://appscode.com>

### RPC

- [brpc](https://github.com/brpc/brpc) - Most common RPC framework used throughout Baidu, with 600,000+ instances and 500+ kinds of services, called "baidu-rpc" inside Baidu.
- [finagle](https://github.com/twitter/finagle) - A fault tolerant, protocol-agnostic RPC system [http://twitter.github.io/finagle](http://twitter.github.io/finagle)
- [grpc](https://github.com/grpc) - A high performance, open source, general-purpose RPC framework
- [proxygen](https://github.com/facebook/proxygen) - A collection of C++ HTTP libraries including an easy to use HTTP server.
- [rsocket](https://github.com/rsocket) - Streaming message protocol with Reactive Extension/Stream semantics
- [sofa-bolt](https://github.com/alipay/sofa-bolt) -  SOFABolt is a lightweight, easy to use and high performance remoting framework based on Netty.
- [sofa-rpc](https://github.com/alipay/sofa-rpc) - SOFARPC is a high-performance, high-extensibility, production-level Java RPC framework.
- [thrift](https://github.com/apache/thrift) - Apache thrift

### Message Broker

- [flume](https://github.com/apache/flume) - Apache Flume is a distributed, reliable, and available service for efficiently collecting, aggregating, and moving large amounts of log data.
- [gnatsd](https://github.com/nats-io/gnatsd) - High-Performance server for NATS, the cloud native messaging system. [https://nats.io](https://nats.io/)
- [jocko](https://github.com/travisjeffery/jocko) - Kafka implemented in Golang with built-in coordination (No ZK dep, single binary install, Cloud Native)<https://twitter.com/travisjeffery>
- [kafka](https://github.com/apache/kafka) - A distributed streaming platform.
- [nsq](https://github.com/nsqio/nsq) - A realtime distributed messaging platform [http://nsq.io/](http://nsq.io/)
- [rabbitmq](https://github.com/rabbitmq) - RabbitMQ is the most widely deployed open source message broker
- [rocketmq](https://github.com/apache/rocketmq) - Apache RocketMQ is a distributed messaging and streaming platform with low latency, high performance and reliability, trillion-level capacity and flexible scalability.

### Miscellaneous

- [ballerina-lang](https://github.com/ballerina-platform/ballerina-lang) - Ballerina is a new programming language for integration built on a sequence diagram metaphor.
- [firecracker](https://github.com/firecracker-microvm/firecracker) - Secure and fast microVMs for serverless computing. [http://firecracker-microvm.io](http://firecracker-microvm.io/)
- [harbor](https://github.com/vmware/harbor) - An enterprise-class container registry server based on Docker Distribution [http://vmware.github.io/harbor/](http://vmware.github.io/harbor/)
- [kind](https://github.com/kubernetes-sigs/kind) - Kubernetes IN Docker - local clusters for testing Kubernetes
- [portainer](https://github.com/portainer/portainer) - Simple management UI for Docker [http://portainer.io](http://portainer.io/)
- [pulumi](https://github.com/pulumi/pulumi) -  A multi-language, multi-cloud development platform -- your code, your cloud, your team
- [skopeo](https://github.com/projectatomic/skopeo) - Work with remote images registries - retrieving information, images, signing content

### Monitoring

- [cadvisor](https://github.com/google/cadvisor) - Analyzes resource usage and performance characteristics of running containers.
- [cortex](https://github.com/weaveworks/cortex) - A multitenant, horizontally scalable Prometheus as a Service
- [elasticsearch-hq](https://github.com/royrusso/elasticsearch-HQ) - Monitoring and Management Web Application for ElasticSearch instances and clusters.
- [grafana](https://github.com/grafana/grafana) - The tool for beautiful monitoring and metric analytics & dashboards for Graphite, InfluxDB & Prometheus & More [https://grafana.com](https://grafana.com/)
- [hawkular-metrics](https://github.com/hawkular/hawkular-metrics) - Time Series Metrics Engine based on Cassandra
- [kibana](https://github.com/elastic/kibana) - Kibana analytics and search dashboard for Elasticsearch [https://www.elastic.co/products/kibana](https://www.elastic.co/products/kibana)
- [kubernetes-zabbix](https://github.com/monitoringartist/kubernetes-zabbix) - Kubernetes Zabbix/Grafana cluster (bare metal, Google Computer Engine - GCE, Google Container Engine - GKE)
- [open-falcon](https://github.com/XiaoMi/open-falcon) - Enterprise Internet monitoring system from Xiaomi [http://open-falcon.com/](http://open-falcon.com/)
- [owl](https://github.com/TalkingData/owl) - Distributed monitoring system from TalkingData
- [prometheus](https://github.com/prometheus/prometheus) - The Prometheus monitoring system and time series database. [https://prometheus.io/](https://prometheus.io/)
- [scope](https://github.com/weaveworks/scope) - Monitoring, visualisation & management for Docker & Kubernetes [https://www.weave.works/product/weave-scope](https://www.weave.works/product/weave-scope/)
- [sofa-lookout](https://github.com/alipay/sofa-lookout) - Lookout can help you to measure and monitor the status of the target system with its multi-dimensional metrics
- [statsd](https://github.com/etsy/statsd) - Daemon for easy but powerful stats aggregation

### Network

- [calico](https://github.com/projectcalico) - A Pure Layer 3 Approach to Virtual Networking for Highly Scalable Data Centers
- [cni](https://github.com/containernetworking/cni) - Container Network Interface - networking for Linux containers
- [contiv](https://github.com/contiv) - Container networking for various use cases
- [flannel](https://github.com/coreos/flannel) - flannel is a network fabric for containers, designed for Kubernetes
- [istio-cni](https://github.com/tiswanso/istio-cni) - Istio CNI to setup kubernetes pod namespaces to redirect traffic to sidecar proxy.
- [knitter](https://github.com/ZTE/Knitter) - Kubernetes network solution
- [kube-router](https://github.com/cloudnativelabs/kube-router) - Kube-router, a turnkey solution for Kubernetes networking. [https://kube-router.io](https://kube-router.io/)
- [matchBox](https://github.com/coreos/matchbox/) - Network boot and provision Container Linux clusters (e.g. etcd3, Kubernetes, more) <https://coreos.com/matchbox/docs/latest/>
- [weave](https://github.com/weaveworks/weave) - Simple, resilient multi-host Docker networking and more. [https://www.weave.works](https://www.weave.works/)

### Security and Audit

- [apparmor](http://wiki.apparmor.net/) - AppArmor is an effective and easy-to-use Linux application security system.
- [authenticator](https://github.com/heptio/authenticator) - A tool for using AWS IAM credentials to authenticate to a Kubernetes cluster
- [cert-manager](https://github.com/jetstack/cert-manager) - Automatically provision and manage TLS certificates in Kubernetes
- [cilium](https://github.com/cilium/cilium) - API Aware Networking and Security using BPF and XDP
- [clair](https://github.com/coreos/clair) - Vulnerability Static Analysis for Containers
- [dex](https://github.com/coreos/dex) - OpenID Connect Identity (OIDC) and OAuth 2.0 Provider with Pluggable Connectors [https://coreos.com/blog/announcing-de…](https://coreos.com/blog/announcing-dex.html)
- [docker-bench-security](https://github.com/docker/docker-bench-security) - The Docker Bench for Security is a script that checks for dozens of common best-practices around deploying Docker containers in production.
- [dockscan](https://github.com/kost/dockscan) - dockscan is security vulnerability and audit scanner for Docker installations
- [drydock](https://github.com/zuBux/drydock) - drydock provides a flexible way of assessing the security of your Docker daemon configuration and containers using editable audit templates
- [falco](https://github.com/draios/falco) - Behavioral Activity Monitoring With Container Support
- [goldfish](https://github.com/Caiyeon/goldfish) - A HashiCorp Vault UI panel written with VueJS and Vault native Go API [https://vault-ui.io](https://vault-ui.io/)
- [grafeas](https://github.com/Grafeas/Grafeas) - Cloud artifact metadata CRUD API and resource specifications
- [guard](https://github.com/appscode/guard) - Kubernetes Authentication WebHook Server
- [k8guard](https://github.com/k8guard) - An auditing system for Kubernetes
- [keycloak](https://github.com/keycloak/keycloak) - Open Source Identity and Access Management For Modern Applications and Services [http://www.keycloak.org](http://www.keycloak.org/)
- [kritis](https://github.com/grafeas/kritis) - Software supply chain security for #Kubernetes apps [https://grafeas.io/docs/concepts/what-is-kritis](https://grafeas.io/docs/concepts/what-is-kritis)
- [kube-bench](https://github.com/aquasecurity/kube-bench) - The Kubernetes Bench for Security is a Go application that checks whether Kubernetes is deployed according to security best practices
- [kube-lego](https://github.com/jetstack/kube-lego) - Automatically request certificates for Kubernetes Ingress resources from Let's Encrypt
- [kube2iam](https://github.com/jtblin/kube2iam) - kube2iam provides different AWS IAM roles for pods running on Kubernetes
- [kubed](https://github.com/appscode/kubed) - 🛡️ A Kubernetes Cluster Operator Daemon by <https://appscode.com>
- [notary](https://github.com/docker/notary) - Notary is a Docker project that allows anyone to have trust over arbitrary collections of data [https://docker.com](https://docker.com/)
- [opa](https://github.com/open-policy-agent/opa) - An open source project to policy-enable your service. [http://openpolicyagent.org](http://openpolicyagent.org/)
- [openscap](https://www.open-scap.org/) - Discover a wide array of tools for managing system security and standards compliance.
- [spiffe](https://github.com/spiffe/spiffe) - The SPIFFE Project [http://spiffe.io](http://spiffe.io/)
- [vault](https://github.com/hashicorp/vault) - A tool for managing secrets. <https://www.vaultproject.io>

### Service Broker

- [open-service-broker-sdk](https://github.com/openshift/open-service-broker-sdk) - A starting point for creating service brokers implementing the Open Service Broker API
- [rotor](https://github.com/turbinelabs/rotor) - Rotor is a fast, lightweight bridge between your service discovery and Envoy’s configuration APIs. Rotor supports Kubernetes, Consul, AWS (EC2 and ECS), DC/OS, flat files, and even other EDS/CDS implementations. 
- [service-catalog](https://github.com/kubernetes-incubator/service-catalog) - Consume services in Kubernetes using the Open Service Broker API
- [service-broker](https://github.com/openservicebrokerapi/servicebroker) - Open Service Broker API Specification <https://openservicebrokerapi.org/>

### Service Mesh

- [amalgam8](https://github.com/amalgam8/amalgam8) - Content and Version-based Routing Fabric for Polyglot Microservices
- [conduit](https://conduit.io) - The Ultralight Service Mesh for Kubernetes [https://conduit.io](https://conduit.io/)
- [istio](https://github.com/istio) - An open platform to connect, manage, and secure microservices.
- [linkerd](https://github.com/linkerd/linkerd) - Resilient service mesh for cloud native apps [https://linkerd.io](https://linkerd.io/)
- [mesher](https://github.com/go-chassis/mesher) -  A light weight service mesh implementation based on [go chassis](https://github.com/ServiceComb/go-chassis).
- [nginmesh](https://github.com/nginmesh/nginmesh) - Service Mesh using Nginx
- [nginx-unit](https://github.com/nginx/unit) - NGINX Unit is a new, lightweight, open source application server built to meet the demands of today’s dynamic and distributed applications.
- [servicecomb](https://github.com/ServiceComb) - ServiceComb is a microservice framework that provides an easy way to develop and deploy applications in the cloud.
- [sofa-mesh](https://github.com/alipay/sofa-mesh) - A solution for large-scale Service Mesh based on Istio. <http://www.sofastack.tech/>
- [supergloo](https://github.com/solo-io/supergloo) - The Service Mesh Orchestration Platform [https://supergloo.solo.io](https://supergloo.solo.io/)

### Service Registry and Discovery

- [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul
- [consul](https://github.com/hashicorp/consul) - Service Discovery and Configuration Made Easy [https://www.consul.io/](https://www.consul.io/)
- [coredns](https://github.com/coredns/coredns) - CoreDNS is a DNS server that chains middleware [https://coredns.io](https://coredns.io/)
- [eureka](https://github.com/Netflix/eureka) - AWS Service registry for resilient mid-tier load balancing and failover.
- [registrator](https://github.com/gliderlabs/registrator) - Service registry bridge for Docker with pluggable adapters [http://gliderlabs.com/registrator](http://gliderlabs.com/registrator)
- [skydns](https://github.com/skynetservices/skydns1) - DNS for skynet or any other service discovery
- [steward](https://github.com/deis/steward) - The Kubernetes-native Service Broker. [https://deis.com/](https://deis.com/)
- [synapse](https://github.com/airbnb/synapse) - A transparent service discovery framework for connecting an SOA
- [vulcand](https://github.com/vulcand/vulcand) - Programmatic load balancer backed by Etcd [http://vulcand.readthedocs.io/](http://vulcand.readthedocs.io/)
- [zookeeper](https://github.com/apache/zookeeper) - Apache ZooKeeper is an effort to develop and maintain an open-source server which enables highly reliable distributed coordination.

### Serverless

- [dispatch](https://github.com/vmware/dispatch) - Dispatch is a framework for deploying and managing serverless style applications.
- [eventing](https://github.com/knative/eventing) - Open source specification and implementation of Knative event binding and delivery
- [faas-netes](https://github.com/alexellis/faas-netes) - Enable Kubernetes as a backend for Functions as a Service (OpenFaaS) <https://github.com/alexellis/faas>
- [firecamp](https://github.com/cloudstax/firecamp) - Serverless Platform for the stateful services [https://www.cloudstax.io](https://www.cloudstax.io/)
- [fission](https://github.com/fission/fission) - Fast Serverless Functions for Kubernetes http://fission.io
- [fn](https://github.com/fnproject/fn) - The container native, cloud agnostic serverless platform. [http://fnproject.io](http://fnproject.io/)
- [funktion](https://github.com/funktionio/funktion/) - a CLI tool for working with funktion <https://funktion.fabric8.io/>
- [fx](https://github.com/metrue/fx) - Poor man's serverless framework based on Docker, Function as a Service with painless.
- [gloo](https://github.com/solo-io/gloo) - The Function Gateway built on top of Envoy
- [ironfunctions](https://github.com/iron-io/functions) - IronFunctions - the serverless microservices platform. [http://iron.io](http://iron.io/)
- [knative-lambda-runtime](https://github.com/triggermesh/knative-lambda-runtime) - Running AWS Lambda Functions on Knative/Kubernetes Clusters [https://triggermesh.com](https://triggermesh.com/)
- [kubeless](https://github.com/kubeless/kubeless) - Kubernetes Native Serverless Framework [http://kubeless.io](http://kubeless.io/)
- [nuclio](https://github.com/nuclio/nuclio) - High-Performance Serverless event and data processing platform
- [openfaas](https://github.com/openfaas/faas) - OpenFaaS - Serverless Functions Made Simple for Docker & Kubernetes [https://blog.alexellis.io/introducing-functions-as-a-service/](https://blog.alexellis.io/introducing-functions-as-a-service/)
- [openwhisk](http://openwhisk.incubator.apache.org/) - Apache OpenWhisk (Incubating) is a [serverless](http://openwhisk.incubator.apache.org/serverless), open source cloud platform that executes functions in response to events at any scale.
- [serverless](https://github.com/serverless/serverless) - Serverless Framework – Build web, mobile and IoT applications with serverless architectures using AWS Lambda, Azure Functions, Google CloudFunctions & more! – [https://serverless.com](https://serverless.com/)
- [spec](https://github.com/cloudevents/spec) - CloudEvents Specification [https://cloudevents.io](https://cloudevents.io/)
- [thanos](https://github.com/improbable-eng/thanos) - Highly available Prometheus setup with long term storage capabilities.

### Storage

- [ceph](https://github.com/ceph/ceph) - Ceph is a distributed object, block, and file storage platform http://ceph.com
- [convoy](https://github.com/rancher/convoy) - A Docker volume plugin, managing persistent container volumes.
- [fastdfs](https://github.com/happyfish100/fastdfs) - FastDFS is an open source high performance distributed file system (DFS). It's major functions include: file storing, file syncing and file accessing, and design for high capacity and load balance.
- [flocker](https://github.com/ClusterHQ/flocker) - Container data volume manager for your Dockerized application https://clusterhq.com
- [glusterd2](https://github.com/gluster/glusterd2) -  GlusterD-2.0 is the distributed management framework to be used for GlusterFS-4.0
- [glusterfs](https://github.com/gluster/glusterfs) - Gluster is a software defined distributed storage that can scale to several petabytes. It provides interfaces for object, block and file storage.
- [heketi](https://github.com/heketi/heketi) - RESTful based volume management framework for GlusterFS
- [infinit](https://github.com/infinit/infinit) - The Infinit policy-based software-defined storage platform. http://infinit.sh
- [leofs](https://leo-project.net/leofs/) - The LeoFS Storage System https://leo-project.net/leofs/
- [longhorn](https://github.com/rancher/longhorn) - We put storage on cows and move them around from rancher.
- [minio](https://github.com/minio/minio) - Minio is an open source object storage server compatible with Amazon S3 APIs https://minio.io
- [openebs](https://github.com/openebs/openebs) - OpenEBS is containerized block storage written in Go for cloud native and other environments w/ per container (or pod) QoS SLAs, tiering and replica policies across AZs and environments, and predictable and scalable performance. [https://www.openebs.io](https://www.openebs.io/)
- [rook](https://github.com/rook/rook) - File, Block, and Object Storage Services for your Cloud-Native Environment https://rook.io
- [storageos](https://storageos.com/) - Enterprise persistent storage for containers and the cloud
- [torus](https://github.com/coreos/torus) - Torus Distributed Storage https://coreos.com/blog/torus-distributed-storage-by-cores.html
- [vitess](https://github.com/youtube/vitess) - Vitess is a database clustering system for horizontal scaling of MySQL. [http://vitess.io](http://vitess.io/)
- [zenko](https://github.com/scality/Zenko) - Because everyone should be in control of their data. [http://zenko.io](http://zenko.io/)

### Tools
- [aglio](https://github.com/danielgtaylor/aglio) - An API Blueprint renderer with theme support that outputs static HTML
- [ark](https://github.com/heptio/ark) - Heptio Ark is a utility for managing disaster recovery, specifically for your Kubernetes cluster resources and persistent volumes. Brought to you by Heptio. [http://www.heptio.com](http://www.heptio.com/)
- [armada](https://github.com/att-comdev/armada) - A python orchestrator for a installing, upgrading, and managing a collection of helm charts, dependencies, and values overrides.
- [autoapply](https://github.com/autoapply/autoapply) - Automatically apply changes from a git repository to Kubernetes.
- [build](https://github.com/knative/build) - A Kubernetes-native Build resource.
- [chaostoolkit](https://github.com/chaostoolkit/chaostoolkit/) - An Open API to Chaos Engineering [http://chaostoolkit.org](http://chaostoolkit.org/)
- [charitify](https://github.com/appscode/chartify) - 📈 Generate Helm Charts from Kubernetes objects by <https://appscode.com>
- [client-go](https://github.com/kubernetes/client-go) - Go client for Kubernetes.
- [cloud-native-sandbox](https://github.com/rootsongjc/cloud-native-sandbox) - Cloud Native Sandbox can help you setup a standalone Kubernetes and Istio environment with Docker on you own laptop.
- [container-structure-test](https://github.com/GoogleCloudPlatform/container-structure-test) - validate the structure of your container images
- [container-transform](https://github.com/micahhausler/container-transform) - Transforms docker-compose, ECS, and Marathon configurations
- [crashcart](https://github.com/oracle/crashcart) - CrashCart: sideload binaries into a running container
- [cri-tools](https://github.com/kubernetes-incubator/cri-tools) - CLI and validation tools for Kubelet Container Runtime Interface (CRI) .
- [devspace](https://github.com/covexo/devspace) - Cloud Native Software Development with Kubernetes and Docker - simply run "devspace up" in any of your projects and start coding directly on top of Kubernetes (works with minikube, self-hosted and cloud-based clusters) [https://devspace-cloud.com](https://devspace-cloud.com/)
- [docker-elk](https://github.com/deviantony/docker-elk) - The ELK stack powered by Docker and Compose.
- [dockersh](https://github.com/Yelp/dockersh) - A shell which places users into individual docker containers
- [dotmesh](https://github.com/dotmesh-io/dotmesh) - dotmesh (dm) is like git for your data volumes (databases, files etc) in Docker and Kubernetes [https://dotmesh.com](https://dotmesh.com/)
- [draft](https://github.com/Azure/draft) - A tool for developers to create cloud-native applications on Kubernetes.
- [dragonfly](https://github.com/alibaba/Dragonfly) - Dragonfly is an intelligent P2P based file distribution system.
- [drakov](https://github.com/Aconex/drakov) - Mock Server that implements the API Blueprint specification
- [eksctl](https://github.com/weaveworks/eksctl) - a CLI for Amazon EKS [https://eksctl.io](https://eksctl.io/)
- [escalator](https://github.com/atlassian/escalator) - Escalator is a batch or job optimized horizontal autoscaler for Kubernetes
- [flux](https://github.com/weaveworks/flux) - A tool for turning container images into running Kubernetes services
- [freshpod](https://github.com/googlecloudplatform/freshpod) - Restart Pods on Minikube automatically on image rebuilds
- [gardener](https://github.com/gardener/gardener) - Kubernetes API server extension and controller manager providing conformant Kubernetes clusters (a.k.a. (off)shoot clusters) as a service (with day-2 ops) on Alibaba, AWS, Azure, GCP, and OpenStack.
- [gockerize](https://github.com/aerofs/gockerize) - Package golang service into minimal docker containers.
- [habitus](https://github.com/cloud66/habitus) - A Build Flow Tool for Docker [http://www.habitus.io](http://www.habitus.io/)
- [helm](https://github.com/kubernetes/helm) - The Kubernetes Package Manager
- [hiboot](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web and cli application framework with dependency injection support [https://hiboot.hidevops.io](https://hiboot.hidevops.io/)
- [istio-pod-network-controller](https://github.com/sabre1041/istio-pod-network-controller) - Controller to manage Istio Pod Network
- [jib](https://github.com/GoogleContainerTools/jib) -  ⛵️ Build container images for your Java applications.
- [jsonnet](https://github.com/google/jsonnet) - Jsonnet - The data templating language [http://jsonnet.org](http://jsonnet.org/)
- [kail](https://github.com/boz/kail) - kubernetes log viewer
- [kaniko](https://github.com/GoogleCloudPlatform/kaniko) - Build Container Images In Kubernetes
- [kd](https://github.com/UKHomeOffice/kd) - Minimalistic kubernetes resources deployment tool with templating
- [kedge](https://github.com/kedgeproject/kedge) - Kedge - Concise Application Definition for Kubernetes [http://kedgeproject.org](http://kedgeproject.org/)
- [kismatic](https://github.com/apprenda/kismatic) - Kismatic Enterprise Toolkit: Fully-Automated, Production-Grade Kubernetes Operations
- [kompose](https://github.com/kubernetes/kompose) - Go from Docker Compose to Kubernetes [http://kompose.io](http://kompose.io/)
- [kops](https://github.com/kubernetes/kops) - Kubernetes Operations (kops) - Production Grade K8s Installation, Upgrades, and Management
- [ksonnet-lib](https://github.com/ksonnet/ksonnet-lib) - (technical preview) Simplify working with Kubernetes [http://ksonnet.heptio.com](http://ksonnet.heptio.com/)
- [ksonnet](https://github.com/ksonnet/ksonnet) - A CLI-supported framework that streamlines writing and deployment of Kubernetes configurations to multiple clusters.<https://ksonnet.io/>
- [ksync](https://github.com/vapor-ware/ksync) -  Sync files between your local system and a kubernetes cluster. <https://vapor-ware.github.io/ksync>
- [ktmpl](https://github.com/InQuicker/ktmpl) - Parameterized templates for Kubernetes manifests.
- [kube-fledged](https://github.com/senthilrch/kube-fledged) - A kubernetes add-on for creating and managing a cache of container images in a kubernetes cluster
- [kube-ps1](https://github.com/jonmosco/kube-ps1) - Kubernetes prompt info for bash and zsh
- [kube-shell](https://github.com/cloudnativelabs/kube-shell) - Kubernetes shell: An integrated shell for working with the Kubernetes CLI
- [kube-version-converter](https://github.com/fleeto/kube-version-converter) - Convert API Object file into specified version.
- [kubeasz](https://github.com/gjmzj/kubeasz) - 使用Ansible脚本安装K8S集群，介绍组件交互原理，方便直接，不受国内网络环境影响
- [kubeadm-offline-installer](https://github.com/fleeto/kubeadm-offline-installer) - Setup a cluster with kubeadm, without internet connections.
- [kubeadm](https://github.com/kubernetes/kubeadm) - Aggregator for issues filed against kubeadm
- [kubeapps](https://kubeapps.com/) - Discover & launch great Kubernetes-ready apps
- [kubebox](https://github.com/astefanutti/kubebox) - ⎈❏ Terminal console for Kubernetes clusters [https://kube.sh](https://kube.sh/)
- [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) - Kubebuilder - SDK for building Kubernetes APIs using CRDs [http://book.kubebuilder.io](http://book.kubebuilder.io/)
- [kubecfg](https://github.com/ksonnet/kubecfg) - A tool for managing complex enterprise Kubernetes environments as code.￼
- [kubectl-trace](https://github.com/fntlnz/kubectl-trace) - Schedule bpftrace programs on your kubernetes cluster using the kubectl
- [kubedb](https://github.com/k8sdb/cli) - KubeDB CLI [https://kubedb.com](https://kubedb.com/) to manage kubernetes ready production-grade Databases
- [kubedirector](https://github.com/bluek8s/kubedirector) - Kubernetes Director (aka KubeDirector) for deploying and managing stateful applications on Kubernetes
- [kubefwd](https://github.com/txn2/kubefwd) - Bulk port forwarding Kubernetes services for local development. [https://imti.co/kubernetes-port-forwarding](https://imti.co/kubernetes-port-forwarding/)
- [kubegen](https://github.com/errordeveloper/kubegen) - kubegen – simple way to describe Kubernetes resources
- [kubehandler](https://github.com/gojektech/kubehandler) - A framework for writing Kubernetes controllers
- [kubeiql](https://github.com/yipeeio/kubeiql) - A GraphQL interface for Kubernetes. <https://kubeiql.io/>
- [kubernetes-client](https://github.com/fabric8io/kubernetes-client) - Java client for Kubernetes & OpenShift 3 [http://fabric8.io](http://fabric8.io/)
- [kubernetes-deploy](https://github.com/Shopify/kubernetes-deploy) - A command-line tool that helps you ship changes to a Kubernetes namespace and understand the result
- [kubernetes-vagrant-centos-cluster](https://github.com/rootsongjc/kubernetes-vagrant-centos-cluster) - Setting up a distributed Kubernetes cluster along with Istio service mesh locally with Vagrant and VirtualBox
- [kubespray](https://github.com/kubernetes-incubator/kubespray) - Setup a kubernetes cluster also mentioned as kargo
- [kubeup](https://github.com/kubeup/archon) - Cluster operation the Kubernetes way
- [kustomize](https://github.com/kubernetes-sigs/kustomize) - Customization of kubernetes YAML configurations
- [microk8s](https://github.com/ubuntu/microk8s) - A kubernetes cluster in a snap [https://microk8s.io](https://microk8s.io/)
- [minikube](https://github.com/kubernetes/minikube) - Run Kubernetes locally
- [opencompose](https://github.com/redhat-developer/opencompose) - OpenCompose - A higher level abstraction for Kubernetes Resource
- [searchlight](https://github.com/appscode/searchlight) - 🔦 Alerts for Kubernetes
- [serving](https://github.com/knative/serving) - Kubernetes-based, scale-to-zero, request-driven compute
- [smith](https://github.com/oracle/Smith) - Smith: A microcontainer builder
- [sonobuoy](https://github.com/heptio/sonobuoy) - Heptio Sonobuoy is a diagnostic tool that makes it easier to understand the state of a Kubernetes cluster by running a set of Kubernetes conformance tests in an accessible and non-destructive manner. Brought to you by Heptio.[http://www.heptio.com](http://www.heptio.com/)
- [source-to-image](https://github.com/openshift/source-to-image) - A tool for building/building artifacts from source and injecting into docker images
- [squash](https://github.com/solo-io/squash) - The debugger for microservices
- [stash](https://github.com/appscode/stash) - 🛅 Backup your Kubernetes Volumes by <htts://appscode.com>
- [stern](https://github.com/wercker/stern) - Multi pod and container log tailing for Kubernetes
- [swagger](https://github.com/swagger-api/swagger-ui) - Swagger UI is a collection of HTML, Javascript, and CSS assets that dynamically generate beautiful documentation from a Swagger-compliant API. [http://swagger.io](http://swagger.io/)
- [tectonic-installer](https://github.com/coreos/tectonic-installer) - Install a Kubernetes cluster the CoreOS Tectonic Way: HA, self-hosted, RBAC, etcd Operator, and more
- [telepresence](https://github.com/datawire/telepresence) - Local development against a remote Kubernetes or OpenShift cluster [http://www.telepresence.io](http://www.telepresence.io/)
- [terminus](https://github.com/godaddy/terminus) - Graceful shutdown and Kubernetes readiness / liveness checks for any Node.js HTTP applications
- [test-infra](https://github.com/kubernetes/test-infra) - Test infrastructure for the Kubernetes project.
- [tor-controller](https://github.com/kragniz/tor-controller) - Run Tor onion services on Kubernetes
- [usernetes](https://github.com/rootless-containers/usernetes) - Kubernetes installable under $HOME, without the root privileges
- [vagrant](https://github.com/mitchellh/vagrant) - Vagrant is a tool for building and distributing development environments. [https://www.vagrantup.com](https://www.vagrantup.com/)
- [watchtower](https://github.com/v2tec/watchtower) - Automatically update running Docker containers

### Tracing

- [appdash](https://github.com/sourcegraph/appdash) - Application tracing system for Go, based on Google's Dapper. [https://sourcegraph.com](https://sourcegraph.com/)
- [jaeger](https://github.com/uber/jaeger) - Jaeger, a Distributed Tracing System <http://uber.github.io/jaeger/>
- [opencensus](https://github.com/census-instrumentation) - A single distribution of libraries that automatically collect traces and metrics from your app, display them locally, and send them to any backend. - https://opencensus.io
- [opentracing](https://github.com/opentracing) - Consistent, expressive, vendor-neutral APIs for distributed tracing and context propagation
- [pinpoint](https://github.com/naver/pinpoint) - Pinpoint is an open source APM (Application Performance Management) tool for large-scale distributed systems written in Java. <http://naver.github.io/pinpoint/>
- [sentry](https://github.com/getsentry/sentry) - Sentry is a cross-platform crash reporting and aggregation platform. [https://sentry.io](https://sentry.io/)
- [skywalking](https://github.com/apache/incubator-skywalking) - An APM system for tracing, monitoring, diagnosing distributed systems, especially based on microservices, cloud native and container
- [sofa-tracker](https://github.com/alipay/sofa-tracer) -  SOFATracer is a component for the distributed system call trace. And through a unified traceId logging the logs of various network calls in the invoking link . These logs can be used for quick discovery of faults, service governance, etc.
- [zipkin](https://github.com/openzipkin/zipkin) - Zipkin is a distributed tracing system [http://zipkin.io](http://zipkin.io/)

### Tutorials

- [aws-workshop-for-kubernetes](https://github.com/aws-samples/aws-workshop-for-kubernetes) - AWS Workshop for Kubernetes
- [envoy-steps](https://github.com/datawire/envoy-steps) - Envoy Step by Step
- [envoy-tutorial](https://github.com/rootsongjc/envoy-tutorial) - Envoy mesh in kubernetes tutorial
- [istio-handbook](https://github.com/rootsongjc/istio-handbook) - Istio服务网格实践指南 by Jimmy Song and [ServiceMesher Community](http://www.servicemesher.com) - https://jimmysong.io/istio-handbook
- [istio-in-action](https://www.manning.com/books/istio-in-action) - Istio in Action by Christian E. Posta - <https://www.manning.com/books/istio-in-action>
- [istio-index-conf2018](https://github.com/todkap/istio-index-conf2018) -  Istio is not just for Microservices: Secure your Kubernetes services using Istio Service Mesh [https://developer.ibm.com/indexconf/s…](https://developer.ibm.com/indexconf/sessions/#!?id=5399)
- [istio-ingress-tutorial](https://github.com/kelseyhightower/istio-ingress-tutorial) - How to run the Istio Ingress Controller on Kubernetes
- [istio-service-mesh-workshop](https://github.com/leecalcote/istio-service-mesh-workshop) - Using Istio Workshop [https://layer5.io](https://layer5.io/)
- [istio-tutorial](https://github.com/redhat-developer-demos/istio-tutorial) - Istio Tutorial for Java Microservices
- [istio101](https://github.com/IBM/istio101) - Istio 101 workshop from IBM
- [ks](https://github.com/red-gate/ks) - A series of Kubernetes walk-throughs
- [kubeadm-workshop](https://github.com/luxas/kubeadm-workshop) - Showcasing a bare-metal multi-platform kubeadm setup with persistent storage and monitoring
- [kubernetes-handbook](https://github.com/rootsongjc/kubernetes-handbook) - Kubernetes中文指南/云原生应用架构实践手册 - https://jimmysong.io/kubernetes-handbook
- [kubernetes-java-simple](https://github.com/arun-gupta/kubernetes-java-sample) - Kubernetes Hands-on Workshop for Java Developers
- [kubernetes-on-aws](https://github.com/zalando-incubator/kubernetes-on-aws) - Deploying Kubernetes on AWS with CloudFormation and Container Linux [https://kubernetes-on-aws.readthedocs.io](https://kubernetes-on-aws.readthedocs.io/)
- [kubernetes-security-best-practice](https://github.com/freach/kubernetes-security-best-practice) - Kubernetes Security - Best Practice Guide
- [kubernetes-the-hard-way](https://github.com/kelseyhightower/kubernetes-the-hard-way) - Bootstrap Kubernetes the hard way on Google Cloud Platform. No scripts.
- [kubicorn](https://github.com/kris-nova/kubicorn) - Create, manage, snapshot, and scale Kubernetes infrastructure in the public cloud.
