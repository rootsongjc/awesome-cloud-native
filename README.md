# Awesome Cloud Native

A curated list of awesome cloud native architectures, tools and  softwares. Inspired by **[awesome-go](https://github.com/avelino/awesome-go)**.

### Contributing

Please take a quick gander at the **[contribution guidelines](https://github.com/rootsongjc/awesome-cloud-native/blob/master/CONTRIBUTING.md)** first. Thanks to all **[contributors](https://github.com/rootsongjc/awesome-cloud-native/graphs/contributors)**; you rock!

### Contents

- Awesome Cloud Native
  - [AI](#ai)
  - [API gateway](#api-gateway)
  - [Big Data](#big-data)
  - [Container engine](#container-engine)
  - [CI-CD](#ci-cd)
  - [Database](#database)
  - [Data Science](#data-science)
  - [Fault tolerant](#fault-tolerant)
  - [Logging](#logging)
  - [Message broker](#message-broker)
  - [Monitoring](#monitoring)
  - [Networking](#networking)
  - [Orchestration and scheduler](#orchestration-and-scheduler)
  - [Portability](#portability)
  - [Proxy and load balancer](#proxy-and-load-balancer)
  - [RPC](#rpc)
  - [Security and audit](#security-and-audit)
  - [Service broker](#service-broker)
  - [Service mesh](#service-mesh)
  - [Service registry and discovery](#service-registry-and-discovery)
  - [Serverless](#serverless)
  - [Storage](#storage)
  - [Tracing](#tracing)
  - [Tools](#tools)
  - [Tutorial](#tutorial)

### AI

- [Caffe2](https://github.com/caffe2/caffe2) - Caffe2 is a lightweight, modular, and scalable deep learning framework. [https://caffe2.ai](https://caffe2.ai/)
- [h2o](https://github.com/h2oai/h2o-3) - Open Source Fast Scalable Machine Learning API For Smarter Applications (Deep Learning, Gradient Boosting, Random Forest, Generalized Linear Modeling (Logistic Regression, Elastic Net), K-Means, PCA, Stacked Ensembles...) [http://h2o.ai](http://h2o.ai/)
- [kubeflow](https://github.com/google/kubeflow) - Machine Learning Toolkit for Kubernetes
- [leaf](https://github.com/autumnai/lea) - Open Machine Intelligence Framework for Hackers. (GPU/CPU) <http://autumnai.com/leaf/book>
- [Predictionio](https://github.com/apache/incubator-predictionio) - PredictionIO, a machine learning server for developers and ML engineers. Built on Apache Spark, HBase and Spray.[https://predictionio.incubator.apache…](https://predictionio.incubator.apache.org/)
- [TensorFlow](https://github.com/tensorflow/tensorflow) - Computation using data flow graphs for scalable machine learning [http://tensorflow.org](http://tensorflow.org/)

### API gateway

- [ambassador](https://github.com/datawire/ambassador) - Ambassador: a self-service API gateway for microservices built on Lyft Envoy [http://www.getambassador.io](http://www.getambassador.io/)
- [kong](https://github.com/Mashape/kong) - 🐒 The Microservice API Gateway <https://getkong.org/install>

### Big Data

- [spark](https://github.com/apache-spark-on-k8s/spark) - Apache Spark enhanced with native Kubernetes scheduler back-end

### Container engine

- [Containerd](https://github.com/containerd/containerd) - An open and reliable container runtime [https://containerd.io](https://containerd.io/)
- [cri-o](http://cri-o.io/) - Lightweight Container Runtime for Kubernetes
- [hyperd](https://github.com/hyperhq/hyperd) - HyperContainer Daemon [http://www.hypercontainer.io](http://www.hypercontainer.io/)
- [Kata Containers](https://katacontainers.io/) - Kata Containers is a new open source project building extremely lightweight virtual machines that seamlessly plug into the containers ecosystem.
- [Moby](https://github.com/moby/moby) - Moby Project - a collaborative project for the container ecosystem to assemble container-based systems[https://mobyproject.org/](https://mobyproject.org/)
- [pouch](https://github.com/alibaba/pouch) - Pouch is an open-source project created to promote the container technology movement.
- [railcar](https://github.com/oracle/railcar) - RailCar: Rust implementation of the Open Containers Initiative oci-runtime
- [rkt](https://github.com/rkt/rkt) - rkt is a pod-native container engine for Linux. It is composable, secure, and built on standards.

### CI-CD

- [argo](https://github.com/argoproj/argo) - Get stuff done with container-native workflows for Kubernetes.
- [buddy.works](https://buddy.works/) - Build, Test & Deploy Code in Seconds Continuous Delivery, simplified.
- [CircleCI](https://github.com/circleci) - Continuous Integration and Deployment
- [cross-cloud](https://github.com/cncf/cross-cloud/) - Cross Cloud Continuous Integration [https://cncf.io](https://cncf.io/)
- [Cyclone](https://github.com/caicloud/cyclone) - A cloud native CI/CD platform built for container workflow
- [Drone](https://github.com/drone/drone) - Drone is a Continuous Delivery platform built on Docker, written in Go [https://drone.io](https://drone.io/)
- [GitLab-CI](https://about.gitlab.com/features/gitlab-ci-cd) - GitLab has integrated CI/CD pipelines to build, test, deploy, and monitor your code 
- [Hygieia](https://github.com/capitalone/Hygieia) - CapitalOne DevOps Dashboard <http://www.capitalone.io/Hygieia>
- [Jenkins](http://jenkins.io) - The leading open source automation server, Jenkins provides hundreds of plugins to support building, deploying and automating any project.
- [Travis](https://github.com/travis-ci/travis-ci) - Free continuous integration platform for GitHub projects. [https://travis-ci.org](https://travis-ci.org/)
- [Wercker](https://github.com/wercker/wercker) - The Wercker CLI can be used to execute pipelines locally for both local development and easy introspection. [http://wercker.com](http://wercker.com/)

### Database

- [ArangoDB](https://github.com/arangodb/arangodb) - ArangoDB is a native multi-model database with flexible data models for documents, graphs, and key-values. Build high performance applications using a convenient SQL-like query language or JavaScript extensions.
- [beringei](https://github.com/facebookincubator/beringei) - Beringei is a high performance, in-memory storage engine for time series data.
- [CockroachDB](https://github.com/cockroachdb/cockroach/) - CockroachDB - the open source, cloud-native SQL database. [https://www.cockroachlabs.com](https://www.cockroachlabs.com/)
- [CouchDB](https://github.com/apache/couchdb) - Apache CouchDB is one of a new breed of database management systems. 
- [etcd](https://github.com/coreos/etcd) - Distributed reliable key-value store for the most critical data of a distributed system[https://coreos.com/etcd/docs/latest/](https://coreos.com/etcd/docs/latest/)
- [InfluxDB](https://github.com/influxdata/influxdb) - Scalable datastore for metrics, events, and real-time analytics [https://influxdata.com](https://influxdata.com/)
- [LevelDB](https://github.com/google/leveldb) - LevelDB is a fast key-value storage library written at Google that provides an ordered mapping from string keys to string values.
- [MongoDB](https://github.com/mongodb/mongo) - MongoDB is an open source database that uses a document-oriented data model.
- [OpenTSDB](https://github.com/OpenTSDB/opentsdb) - A scalable, distributed Time Series Database. [http://opentsdb.net](http://opentsdb.net/)
- [Redis](https://github.com/antirez/redis) - Redis is an in-memory database that persists on disk. The data model is key-value, but many different kind of values are supported: Strings, Lists, Sets, Sorted Sets, Hashes, HyperLogLogs, Bitmaps. [http://redis.io](http://redis.io/)
- [RethinkDB](https://github.com/rethinkdb/rethinkdb) - The open-source database for the realtime web. [https://rethinkdb.com](https://rethinkdb.com/)
- [TiDB](https://github.com/pingcap/tidb) - TiDB is a distributed NewSQL database compatible with MySQL protocol [https://pingcap.com](https://pingcap.com/)

### Data Science

- [Pachyderm](https://github.com/pachyderm/pachyderm) - Reproducible Data Science at Scale! [http://pachyderm.io](http://pachyderm.io/)
- [Wallaroo](https://github.com/WallarooLabs/wallaroo) - Ultrafast and elastic data processing [https://www.wallaroolabs.com](https://www.wallaroolabs.com/)

### DevOps

- [ContainerOps](https://github.com/Huawei/containerops) - DevOps Orchestration Platform [https://cncf.build](https://cncf.build/)
- [Crane](https://github.com/Dataman-Cloud/crane) - Yet another control plane based on docker built-in swarmkit [https://www.shurenyun.com/product-crane.html](https://www.shurenyun.com/product-crane.html)
- [Fabric8](https://github.com/fabric8io/fabric8) - fabric8 is an open source microservices platform based on Docker, Kubernetes and Jenkins [http://fabric8.io/](http://fabric8.io/)
- [Harbor](https://github.com/vmware/harbor) - An enterprise-class container registry server based on Docker Distribution [http://vmware.github.io/harbor/](http://vmware.github.io/harbor/)
- [Lastbackend](https://github.com/lastbackend/lastbackend) - Container orchestration with CI&CD, cli and amazing UI [https://lastbackend.com](https://lastbackend.com/)
- [Spinnaker](https://github.com/spinnaker/spinnaker) - Spinnaker is an open source, multi-cloud continuous delivery platform for releasing software changes with high velocity and confidence. [http://www.spinnaker.io/](http://www.spinnaker.io/)
- [Terraform](https://github.com/hashicorp/terraform) - Terraform is a tool for building, changing, and combining infrastructure safely and efficiently. [https://www.terraform.io/](https://www.terraform.io/)

### Fault tolerant

- [Chaosmonkey](https://github.com/Netflix/chaosmonkey) - Chaos Monkey is a resiliency tool that helps applications tolerate random instance failures.
- [Hystrix](https://github.com/Netflix/Hystrix) - Hystrix is a latency and fault tolerance library designed to isolate points of access to remote systems, services and 3rd party libraries, stop cascading failure and enable resilience in complex distributed systems where failure is inevitable.

### Logging

- [collectbeat](https://github.com/eBay/collectbeat) - Beats with discovery capabilities for environments like Kubernetes
- [Elasticsearch](https://github.com/elastic/elasticsearch) - Open Source, Distributed, RESTful Search Engine [https://www.elastic.co/products/elast…](https://www.elastic.co/products/elasticsearch)
- [Fluent-bit](https://github.com/fluent/fluent-bit) - Fast and Lightweight Log/Data Forwarder for Linux, BSD and OSX <http://fluentbit.io/>
- [Fluentd](https://github.com/fluent/fluentd) - Fluentd: Unified Logging Layer (project under CNCF) [http://www.fluentd.org/](http://www.fluentd.org/)
- [Fluentd-pilot](https://github.com/AliyunContainerService/fluentd-pilot) - Collect logs in docker containers
- [Flume](http://flume.apache.org/) - Flume is a distributed, reliable, and available service for efficiently collecting, aggregating, and moving large amounts of log data. 
- [Beats](https://github.com/elastic/beats) - Beats - Lightweight shippers for Elasticsearch & Logstash [https://www.elastic.co/products/beats](https://www.elastic.co/products/beats)
- [Heapster](https://github.com/kubernetes/heapster) - Compute Resource Usage Analysis and Monitoring of Container Clusters
- [Telegraf](https://github.com/influxdata/telegraf) - The plugin-driven server agent for collecting & reporting metrics.

### Orchestration and scheduler

- [Blox](https://github.com/blox/blox) - Open source tools for building custom schedulers on Amazon ECS
- [Compose](https://github.com/docker/compose) - Define and run multi-container applications with Docker
- [Conductor](https://github.com/Netflix/conductor) - Conductor is a microservices orchestration engine - [https://netflix.github.io/conductor/](https://netflix.github.io/conductor/)
- [DomeOS](https://github.com/domeos) - An enterprise application orchestration system based on docker by Sohu.com - http://domeos.org
- [Fleet](https://github.com/coreos/fleet) - fleet ties together systemd and etcd into a distributed init system
- [Kubernetes](https://github.com/kubernetes/kubernetes) - Production-Grade Container Scheduling and Management [http://kubernetes.io](http://kubernetes.io/)
- [Mesos](https://github.com/apache/mesos) - Apache Mesos abstracts CPU, memory, storage, and other compute resources away from machines (physical or virtual), enabling fault-tolerant and elastic distributed systems to easily be built and run effectively.
- [Marathon](https://github.com/mesosphere/marathon) - Deploy and manage containers (including Docker) on top of Apache Mesos at scale.[https://mesosphere.github.io/marathon/](https://mesosphere.github.io/marathon/)
- [Serf](https://github.com/hashicorp/serf) - Service orchestration and management tool by hashicorp. [https://www.serf.io/](https://www.serf.io/)
- [Swan](https://github.com/Dataman-Cloud/swan) - A Distributed, Highly Available Mesos Scheduler, Inspired by the design of Google Borg [https://www.shurenyun.com/product-swa…](https://www.shurenyun.com/product-swan.html)
- [Swarm](https://github.com/docker/swarm) - Swarm: a Docker-native clustering system

### Portability

- [Cloud Foundry](https://www.cloudfoundry.org) - Cloud Foundry is an open source, multi cloud application platform as a service (PaaS) governed by the Cloud Foundry Foundation.
- [conjure-up](https://github.com/conjure-up/conjure-up) - Deploying complex solutions, magically. [https://conjure-up.io](https://conjure-up.io/)
- [DC/OS](https://github.com/dcos) - Datacenter Operating System
- [Deis](https://github.com/deis/deis) - Deis v1, the CoreOS and Docker PaaS: Your PaaS. Your Rules. [https://deis.com/docs/](https://deis.com/docs/)
- [Kubernator](https://github.com/smpio/kubernator) - Alternative Kubernetes UI
- [OpenDCP](https://github.com/weibocom/opendcp) - Docker platform developed by weibo.com
- [Openshift](https://github.com/openshift/origin) - Enterprise Kubernetes for Developers [http://www.openshift.org](http://www.openshift.org/)
- [Rancher](https://github.com/rancher/rancher) - Platform for operating Docker in production [http://rancher.com](http://rancher.com/)
- [Supergiant](https://github.com/supergiant/supergiant) - Automatically scale hardware and easily run stateful applications using Kubernetes. <https://supergiant.io/>
- [vamp](https://github.com/magneticio/vamp) - Vamp - canary releasing and autoscaling for microservice systems [http://vamp.io](http://vamp.io/)

### Proxy and load balancer

- [Contour](https://github.com/heptio/contour) - Contour is a Kubernetes ingress controller for Lyft's Envoy proxy.
- [Envoy](https://github.com/lyft/envoy) - C++ front/service proxy [https://lyft.github.io/envoy](https://lyft.github.io/envoy)
- [Haproxy](https://github.com/haproxy/haproxy) - HAProxy is a free, very fast and reliable solution offering high availability, load balancing, and proxying for TCP and HTTP-based applications. 
- [Kong-ingress](https://github.com/koli/kong-ingress) - A Kubernetes Ingress for Kong
- [Nginx-kubernetes-ingress](https://github.com/nginxinc/kubernetes-ingress) - NGINX and NGINX Plus Ingress Controllers for Kubernetes
- [Nginx](https://github.com/nginx/nginx) - nginx [engine x] is an HTTP and reverse proxy server, a mail proxy server, and a generic TCP/UDP proxy server, originally written by [Igor Sysoev](http://sysoev.ru/en/). 
- [Ribbon](https://github.com/Netflix/ribbon) - Ribbon is a Inter Process Communication (remote procedure calls) library with built in software load balancers. The primary usage model involves REST calls with various serialization scheme support.
- [Traefik](https://github.com/containous/traefik) - Træfik, a modern reverse proxy [https://traefik.io](https://traefik.io/)
- [voyager](https://github.com/appscode/voyager) - ✈️️ Secure Ingress Controller for Kubernetes by <https://appscode.com>

### RPC

- [brpc](https://github.com/brpc/brpc) - Most common RPC framework used throughout Baidu, with 600,000+ instances and 500+ kinds of services, called "baidu-rpc" inside Baidu.
- [finagle](https://github.com/twitter/finagle) - A fault tolerant, protocol-agnostic RPC system [http://twitter.github.io/finagle](http://twitter.github.io/finagle)
- [gRPC](https://github.com/grpc) - A high performance, open source, general-purpose RPC framework
- [proxygen](https://github.com/facebook/proxygen) - A collection of C++ HTTP libraries including an easy to use HTTP server.
- [thrift](https://github.com/apache/thrift) - Apache thrift

### Message broker

- [Flume](https://github.com/apache/flume) - Apache Flume is a distributed, reliable, and available service for efficiently collecting, aggregating, and moving large amounts of log data. 
- [Kafka](https://github.com/apache/kafka) - A distributed streaming platform.
- [nsq](https://github.com/nsqio/nsq) - A realtime distributed messaging platform [http://nsq.io/](http://nsq.io/)
- [RabbitMQ](https://github.com/rabbitmq) - RabbitMQ is the most widely deployed open source message broker

### Monitoring

- [cAdvisor](https://github.com/google/cadvisor) - Analyzes resource usage and performance characteristics of running containers. 
- [Cortex](https://github.com/weaveworks/cortex) - A multitenant, horizontally scalable Prometheus as a Service
- [elasticsearch-HQ](https://github.com/royrusso/elasticsearch-HQ) - Monitoring and Management Web Application for ElasticSearch instances and clusters.
- [Grafana](https://github.com/grafana/grafana) - The tool for beautiful monitoring and metric analytics & dashboards for Graphite, InfluxDB & Prometheus & More [https://grafana.com](https://grafana.com/)
- [kubernetes-zabbix](https://github.com/monitoringartist/kubernetes-zabbix) - Kubernetes Zabbix/Grafana cluster (bare metal, Google Computer Engine - GCE, Google Container Engine - GKE)
- [open-falcon](https://github.com/XiaoMi/open-falcon) - Enterprise Internet monitoring system from Xiaomi [http://open-falcon.com/](http://open-falcon.com/)
- [owl](https://github.com/TalkingData/owl) - Distributed monitoring system from TalkingData
- [Prometheus](https://github.com/prometheus/prometheus) - The Prometheus monitoring system and time series database. [https://prometheus.io/](https://prometheus.io/)
- [Scope](https://github.com/weaveworks/scope) - Monitoring, visualisation & management for Docker & Kubernetes [https://www.weave.works/product/weave-scope](https://www.weave.works/product/weave-scope/)
- [Kibana](https://github.com/elastic/kibana) - Kibana analytics and search dashboard for Elasticsearch [https://www.elastic.co/products/kibana](https://www.elastic.co/products/kibana)

### Networking

- [CNI](https://github.com/containernetworking/cni) - Container Network Interface - networking for Linux containers
- [Calico](https://github.com/projectcalico) - A Pure Layer 3 Approach to Virtual Networking for Highly Scalable Data Centers
- [Contiv](https://github.com/contiv) - Container networking for various use cases
- [Flannel](https://github.com/coreos/flannel) - flannel is a network fabric for containers, designed for Kubernetes
- [MatchBox](https://github.com/coreos/matchbox/) - Network boot and provision Container Linux clusters (e.g. etcd3, Kubernetes, more) <https://coreos.com/matchbox/docs/latest/>
- [Weave](https://github.com/weaveworks/weave) - Simple, resilient multi-host Docker networking and more. [https://www.weave.works](https://www.weave.works/)

### Security and audit

- [AppArmor](http://wiki.apparmor.net/) - AppArmor is an effective and easy-to-use Linux application security system.
- [authenticator](https://github.com/heptio/authenticator) - A tool for using AWS IAM credentials to authenticate to a Kubernetes cluster
- [Cilium](https://github.com/cilium/cilium) - Linux Native, HTTP Aware Networking and Security for Containers
- [Clair](https://github.com/coreos/clair) - Vulnerability Static Analysis for Containers
- [OpenSCAP](https://www.open-scap.org/) - Discover a wide array of tools for managing system security and standards compliance.
- [dex](https://github.com/coreos/dex) - OpenID Connect Identity (OIDC) and OAuth 2.0 Provider with Pluggable Connectors [https://coreos.com/blog/announcing-de…](https://coreos.com/blog/announcing-dex.html)
- [docker-bench-security](https://github.com/docker/docker-bench-security) - The Docker Bench for Security is a script that checks for dozens of common best-practices around deploying Docker containers in production.
- [dockscan](https://github.com/kost/dockscan) - dockscan is security vulnerability and audit scanner for Docker installations
- [drydock](https://github.com/zuBux/drydock) - drydock provides a flexible way of assessing the security of your Docker daemon configuration and containers using editable audit templates
- [falco](https://github.com/draios/falco) - Behavioral Activity Monitoring With Container Support
- [goldfish](https://github.com/Caiyeon/goldfish) - A HashiCorp Vault UI panel written with VueJS and Vault native Go API [https://vault-ui.io](https://vault-ui.io/)
- [grafeas](https://github.com/Grafeas/Grafeas) - Cloud artifact metadata CRUD API and resource specifications
- [guard](https://github.com/appscode/guard) - Kubernetes Authentication WebHook Server
- [k8guard](https://github.com/k8guard) - An auditing system for Kubernetes
- [kube-lego](https://github.com/jetstack/kube-lego) - Automatically request certificates for Kubernetes Ingress resources from Let's Encrypt
- [kube2iam](https://github.com/jtblin/kube2iam) - kube2iam provides different AWS IAM roles for pods running on Kubernetes
- [kubed](https://github.com/appscode/kubed) - 🛡️ A Kubernetes Cluster Operator Daemon by <https://appscode.com>
- [metallb](https://github.com/google/metallb) - A network load-balancer implementation for Kubernetes using BGP
- [notary](https://github.com/docker/notary) - Notary is a Docker project that allows anyone to have trust over arbitrary collections of data [https://docker.com](https://docker.com/)
- [vault](https://github.com/hashicorp/vault) - A tool for managing secrets. <https://www.vaultproject.io>

### Service broker

- [open-service-broker-sdk](https://github.com/openshift/open-service-broker-sdk) - A starting point for creating service brokers implementing the Open Service Broker API
- [service-catalog](https://github.com/kubernetes-incubator/service-catalog) - Consume services in Kubernetes using the Open Service Broker API
- [service-broker](https://github.com/openservicebrokerapi/servicebroker) - Open Service Broker API Specification <https://openservicebrokerapi.org/>

### Service mesh

- [Amalgam8](https://github.com/amalgam8/amalgam8) - Content and Version-based Routing Fabric for Polyglot Microservices
- [Conduit](https://conduit.io) - The Ultralight Service Mesh for Kubernetes [https://conduit.io](https://conduit.io/)
- [Istio](https://github.com/istio) - An open platform to connect, manage, and secure microservices.
- [Linkerd](https://github.com/linkerd/linkerd) - Resilient service mesh for cloud native apps [https://linkerd.io](https://linkerd.io/)
- [nginmesh](https://github.com/nginmesh/nginmesh) - Service Mesh using Nginx
- [nginx-unit](https://github.com/nginx/unit) - NGINX Unit is a new, lightweight, open source application server built to meet the demands of today’s dynamic and distributed applications. 
- [ServiceComb](https://github.com/ServiceComb) - ServiceComb is a microservice framework that provides an easy way to develop and deploy applications in the cloud. 

### Service registry and discovery

- [Confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul
- [Consul](https://github.com/hashicorp/consul) - Service Discovery and Configuration Made Easy [https://www.consul.io/](https://www.consul.io/)
- [CoreDNS](https://github.com/coredns/coredns) - CoreDNS is a DNS server that chains middleware [https://coredns.io](https://coredns.io/)
- [Eureka](https://github.com/Netflix/eureka) - AWS Service registry for resilient mid-tier load balancing and failover.
- [Registrator](https://github.com/gliderlabs/registrator) - Service registry bridge for Docker with pluggable adapters [http://gliderlabs.com/registrator](http://gliderlabs.com/registrator)
- [SkyDNS](https://github.com/skynetservices/skydns1) - DNS for skynet or any other service discovery
- [Steward](https://github.com/deis/steward) - The Kubernetes-native Service Broker. [https://deis.com/](https://deis.com/)
- [Synapse](https://github.com/airbnb/synapse) - A transparent service discovery framework for connecting an SOA
- [Vulcand](https://github.com/vulcand/vulcand) - Programmatic load balancer backed by Etcd [http://vulcand.readthedocs.io/](http://vulcand.readthedocs.io/)
- [Zookeeper](https://github.com/apache/zookeeper) - Apache ZooKeeper is an effort to develop and maintain an open-source server which enables highly reliable distributed coordination.

### Serverless

- [faas-netes](https://github.com/alexellis/faas-netes) - Enable Kubernetes as a backend for Functions as a Service (OpenFaaS) <https://github.com/alexellis/faas>
- [fn](https://github.com/fnproject/fn) - The container native, cloud agnostic serverless platform. [http://fnproject.io](http://fnproject.io/)
- [funktion](https://github.com/funktionio/funktion/) - a CLI tool for working with funktion <https://funktion.fabric8.io/>
- [fx](https://github.com/metrue/fx) - Poor man's serverless framework based on Docker, Function as a Service with painless.
- [IronFunctions](https://github.com/iron-io/functions) - IronFunctions - the serverless microservices platform. [http://iron.io](http://iron.io/)
- [kubeless](https://github.com/kubeless/kubeless) - Kubernetes Native Serverless Framework [http://kubeless.io](http://kubeless.io/)
- [OpenFaaS](https://github.com/openfaas/faas) - OpenFaaS - Serverless Functions Made Simple for Docker & Kubernetes [https://blog.alexellis.io/introducing-functions-as-a-service/](https://blog.alexellis.io/introducing-functions-as-a-service/) 
- [OpenWhisk](http://openwhisk.incubator.apache.org/) - Apache OpenWhisk (Incubating) is a [serverless](http://openwhisk.incubator.apache.org/serverless), open source cloud platform that executes functions in response to events at any scale.

### Storage

- [Ceph](https://github.com/ceph/ceph) - Ceph is a distributed object, block, and file storage platform http://ceph.com
- [Convoy](https://github.com/rancher/convoy) - A Docker volume plugin, managing persistent container volumes.
- [Flocker](https://github.com/ClusterHQ/flocker) - Container data volume manager for your Dockerized application https://clusterhq.com
- [GlusterFS](https://github.com/gluster/glusterfs) - Gluster is a software defined distributed storage that can scale to several petabytes. It provides interfaces for object, block and file storage.
- [Heketi](https://github.com/heketi/heketi) - RESTful based volume management framework for GlusterFS
- [Infinit](https://github.com/infinit/infinit) - The Infinit policy-based software-defined storage platform. http://infinit.sh
- [LeoFS](https://leo-project.net/leofs/) - The LeoFS Storage System https://leo-project.net/leofs/
- [Longhorn](https://github.com/rancher/longhorn) - We put storage on cows and move them around from rancher.
- [Minio](https://github.com/minio/minio) - Minio is an open source object storage server compatible with Amazon S3 APIs https://minio.io
- [OpenEBS](https://github.com/openebs/openebs) - OpenEBS is containerized block storage written in Go for cloud native and other environments w/ per container (or pod) QoS SLAs, tiering and replica policies across AZs and environments, and predictable and scalable performance. [https://www.openebs.io](https://www.openebs.io/)
- [Rook](https://github.com/rook/rook) - File, Block, and Object Storage Services for your Cloud-Native Environment https://rook.io
- [StorageOS](https://storageos.com/) - Enterprise persistent storage for containers and the cloud
- [Torus](https://github.com/coreos/torus) - Torus Distributed Storage https://coreos.com/blog/torus-distributed-storage-by-cores.html

### Tools
- [Aglio](https://github.com/danielgtaylor/aglio) - An API Blueprint renderer with theme support that outputs static HTML
- [Ark](https://github.com/heptio/ark) - Heptio Ark is a utility for managing disaster recovery, specifically for your Kubernetes cluster resources and persistent volumes. Brought to you by Heptio. [http://www.heptio.com](http://www.heptio.com/)
- [Draft](https://github.com/Azure/draft) - A tool for developers to create cloud-native applications on Kubernetes.
- [Habitus](https://github.com/cloud66/habitus) - A Build Flow Tool for Docker [http://www.habitus.io](http://www.habitus.io/)
- [Helm](https://github.com/kubernetes/helm) - The Kubernetes Package Manager
- [Smith](https://github.com/oracle/Smith) - Smith: A microcontainer builder
- [Swagger](https://github.com/swagger-api/swagger-ui) - Swagger UI is a collection of HTML, Javascript, and CSS assets that dynamically generate beautiful documentation from a Swagger-compliant API. [http://swagger.io](http://swagger.io/)
- [Vagrant](https://github.com/mitchellh/vagrant) - Vagrant is a tool for building and distributing development environments. [https://www.vagrantup.com](https://www.vagrantup.com/)
- [charitify](https://github.com/appscode/chartify) - 📈 Generate Helm Charts from Kubernetes objects by <https://appscode.com>
- [client-go](https://github.com/kubernetes/client-go) - Go client for Kubernetes.
- [container-transform](https://github.com/micahhausler/container-transform) - Transforms docker-compose, ECS, and Marathon configurations
- [crashcart](https://github.com/oracle/crashcart) - CrashCart: sideload binaries into a running container
- [cri-tools](https://github.com/kubernetes-incubator/cri-tools) - CLI and validation tools for Kubelet Container Runtime Interface (CRI) .
- [docker-elk](https://github.com/deviantony/docker-elk) - The ELK stack powered by Docker and Compose.
- [dockersh](https://github.com/Yelp/dockersh) - A shell which places users into individual docker containers
- [drakov](https://github.com/Aconex/drakov) - Mock Server that implements the API Blueprint specification
- [flux](https://github.com/weaveworks/flux) - A tool for turning container images into running Kubernetes services
- [gockerize](https://github.com/aerofs/gockerize) - Package golang service into minimal docker containers.
- [jsonnet](https://github.com/google/jsonnet) - Jsonnet - The data templating language [http://jsonnet.org](http://jsonnet.org/)
- [kail](https://github.com/boz/kail) - kubernetes log viewer
- [kd](https://github.com/UKHomeOffice/kd) - Minimalistic kubernetes resources deployment tool with templating
- [kedge](https://github.com/kedgeproject/kedge) - Kedge - Concise Application Definition for Kubernetes [http://kedgeproject.org](http://kedgeproject.org/)
- [kismatic](https://github.com/apprenda/kismatic) - Kismatic Enterprise Toolkit: Fully-Automated, Production-Grade Kubernetes Operations
- [kompose](https://github.com/kubernetes/kompose) - Go from Docker Compose to Kubernetes [http://kompose.io](http://kompose.io/)
- [kops](https://github.com/kubernetes/kops) - Kubernetes Operations (kops) - Production Grade K8s Installation, Upgrades, and Management
- [ksonnet-lib](https://github.com/ksonnet/ksonnet-lib) - (technical preview) Simplify working with Kubernetes [http://ksonnet.heptio.com](http://ksonnet.heptio.com/)
- [ktmpl](https://github.com/InQuicker/ktmpl) - Parameterized templates for Kubernetes manifests.
- [kube-shell](https://github.com/cloudnativelabs/kube-shell) - Kubernetes shell: An integrated shell for working with the Kubernetes CLI
- [kubeadm](https://github.com/kubernetes/kubeadm) - Aggregator for issues filed against kubeadm
- [kubeadm-offline-installer](https://github.com/fleeto/kubeadm-offline-installer) - Setup a cluster with kubeadm, without internet connections.
- [kubeapps](https://kubeapps.com/) - Discover & launch great Kubernetes-ready apps
- [kubecfg](https://github.com/ksonnet/kubecfg) - A tool for managing complex enterprise Kubernetes environments as code.￼
- [kubedb](https://github.com/k8sdb/cli) - KubeDB CLI [https://kubedb.com](https://kubedb.com/) to manage kubernetes ready production-grade Databases
- [kubegen](https://github.com/errordeveloper/kubegen) - kubegen – simple way to describe Kubernetes resources
- [kubernetes-client](https://github.com/fabric8io/kubernetes-client) - Java client for Kubernetes & OpenShift 3 [http://fabric8.io](http://fabric8.io/)
- [kubernetes-deploy](https://github.com/Shopify/kubernetes-deploy) - A command-line tool that helps you ship changes to a Kubernetes namespace and understand the result
- [kubespray](https://github.com/kubernetes-incubator/kubespray) - Setup a kubernetes cluster also mentioned as kargo
- [kubeup](https://github.com/kubeup/archon) - Cluster operation the Kubernetes way
- [minikube](https://github.com/kubernetes/minikube) - Run Kubernetes locally
- [opencompose](https://github.com/redhat-developer/opencompose) - OpenCompose - A higher level abstraction for Kubernetes Resource
- [prometheus-operator](https://github.com/coreos/prometheus-operator) - Prometheus Operator creates/configures/manages Prometheus clusters atop Kubernetes<https://coreos.com/operators/prometheus>
- [searchlight](https://github.com/appscode/searchlight) - 🔦 Alerts for Kubernetes
- [sonobuoy](https://github.com/heptio/sonobuoy) - Heptio Sonobuoy is a diagnostic tool that makes it easier to understand the state of a Kubernetes cluster by running a set of Kubernetes conformance tests in an accessible and non-destructive manner. Brought to you by Heptio.[http://www.heptio.com](http://www.heptio.com/)
- [source-to-image](https://github.com/openshift/source-to-image) - A tool for building/building artifacts from source and injecting into docker images
- [stash](https://github.com/appscode/stash) - 🛅 Backup your Kubernetes Volumes by <htts://appscode.com>
- [stern](https://github.com/wercker/stern) - Multi pod and container log tailing for Kubernetes
- [tectonic-installer](https://github.com/coreos/tectonic-installer) - Install a Kubernetes cluster the CoreOS Tectonic Way: HA, self-hosted, RBAC, etcd Operator, and more
- [telepresence](https://github.com/datawire/telepresence) - Local development against a remote Kubernetes or OpenShift cluster [http://www.telepresence.io](http://www.telepresence.io/)
- [test-infra](https://github.com/kubernetes/test-infra) - Test infrastructure for the Kubernetes project.

### Tracing

- [appdash](https://github.com/sourcegraph/appdash) - Application tracing system for Go, based on Google's Dapper. [https://sourcegraph.com](https://sourcegraph.com/)
- [jaeger](https://github.com/uber/jaeger) - Jaeger, a Distributed Tracing System <http://uber.github.io/jaeger/>
- [OpenTracing](https://github.com/opentracing) - Consistent, expressive, vendor-neutral APIs for distributed tracing and context propagation
- [Sentry](https://github.com/getsentry/sentry) - Sentry is a cross-platform crash reporting and aggregation platform. [https://sentry.io](https://sentry.io/)
- [Zipkin](https://github.com/openzipkin/zipkin) - Zipkin is a distributed tracing system [http://zipkin.io](http://zipkin.io/)

### Tutorial

- [aws-workshop-for-kubernetes](https://github.com/aws-samples/aws-workshop-for-kubernetes) - AWS Workshop for Kubernetes
- [Cloud Native Go](http://rootsongjc.github.io/cloud-native-go) - Building Web Applications and Microservices for the Cloud with Go and React 
  written by Kevin Hoffman and Dan Nemeth translated to Chinese by four guys from [TalkingData](http://www.talkingdata.com/) with ❤️
- [Istio-ingress-tutorial](https://github.com/kelseyhightower/istio-ingress-tutorial) - How to run the Istio Ingress Controller on Kubernetes
- [ks](https://github.com/red-gate/ks) - A series of Kubernetes walk-throughs
- [kubernetes-on-aws](https://github.com/zalando-incubator/kubernetes-on-aws) - Deploying Kubernetes on AWS with CloudFormation and Container Linux [https://kubernetes-on-aws.readthedocs.io](https://kubernetes-on-aws.readthedocs.io/)
- [kubeadm-workshop](https://github.com/luxas/kubeadm-workshop) - Showcasing a bare-metal multi-platform kubeadm setup with persistent storage and monitoring
- [kubernetes-handbook](https://github.com/rootsongjc/kubernetes-handbook) - kubernetes中文指南/实践手册 https://jimmysong.io/kubernetes-handbook
- [kubernetes-java-simple](https://github.com/arun-gupta/kubernetes-java-sample) - Kubernetes Hands-on Workshop for Java Developers
- [kubernetes-the-hard-way](https://github.com/kelseyhightower/kubernetes-the-hard-way) - Bootstrap Kubernetes the hard way on Google Cloud Platform. No scripts.
- [kubicorn](https://github.com/kris-nova/kubicorn) - Create, manage, snapshot, and scale Kubernetes infrastructure in the public cloud.
- [Migrating to Cloud Native Application Architectures](https://github.com/rootsongjc/migrating-to-cloud-native-application-architectures) - 《迁移到云原生应用架构》中文版 https://jimmysong.io/migrating-to-cloud-native-application-architectures/
