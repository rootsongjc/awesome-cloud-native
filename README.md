# Awesome Cloud Native [![Awesome](https://awesome.re/badge-flat2.svg)](https://github.com/sindresorhus/awesome#readme) 

> A curated list of cloud native tools, software, and tutorials.

Cloud Native is a behavior and design philosophy. At its essence, any behavior or approach that improves resource utilization and application delivery efficiency in the cloud is called Cloud Native.

- [Join the Cloud Native Community](https://cloudnative.to)

## Contents

- [AI](#ai)
- [API Gateway](#api-gateway)
- [Application Delivery](#application-delivery)
- [Big Data](#big-data)
- [Container Runtime](#container-runtime)
- [Database](#database)
- [Edge Computing](#edge-computing)
- [Fault Tolerant](#fault-tolerant)
- [Kubernetes Operators](#kubernetes-operators)
- [Logging](#logging)
- [Message Broker](#message-broker)
- [Miscellaneous](#miscellaneous)
- [Monitoring](#monitoring)
- [Network](#network)
- [Observability](#observability)
- [Orchestration and Scheduler](#orchestration-and-scheduler)
- [PaaS Manager](#paas-manager)
- [Proxy](#proxy)
- [RPC](#rpc)
- [Security and Audit](#security-and-audit)
- [Service Mesh](#service-mesh)
- [Service Registry and Discovery](#service-registry-and-discovery)
- [Serverless](#serverless)
- [Storage](#storage)
- [Tools](#tools)
- [Tracing](#tracing)
- [Tutorials](#tutorials)
- [Community](#community)

## AI

- [allennlp](https://github.com/allenai/allennlp) - An open-source NLP research library, built on PyTorch. [https://allennlp.org/](https://allennlp.org/)
- [caffe2](https://github.com/facebookarchive/caffe2) - Caffe2 is a lightweight, modular, and scalable deep learning framework. [https://caffe2.ai](https://caffe2.ai/)
- [elasticdl](https://github.com/sql-machine-learning/elasticdl) - Kubernetes-native Deep Learning Framework.
- [h2o-3](https://github.com/h2oai/h2o-3) - Open Source Fast Scalable Machine Learning API For Smarter Applications (Deep Learning, Gradient Boosting, Random Forest, Generalized Linear Modeling (Logistic Regression, Elastic Net), K-Means, PCA, Stacked Ensembles.) [http://www.h2o.ai/](http://www.h2o.ai/)
- [keras](https://github.com/keras-team/keras) - Keras is a high-level neural networks API, written in Python and capable of running on top of TensorFlow, CNTK, or Theano. [https://keras.io/](https://keras.io/)
- [kubeflow](https://github.com/kubeflow/kubeflow) - Machine Learning Toolkit for Kubernetes. [https://www.kubeflow.org](https://www.kubeflow.org/)
- [leaf](https://github.com/autumnai/leaf) - Open Machine Intelligence Framework for Hackers. (GPU/CPU) <http://ww1.autumnai.com>
- [PaddlePaddle](https://github.com/PaddlePaddle/Paddle) - PArallel Distributed Deep LEarning: Machine Learning Framework from Industrial Practice （『飞桨』核心框架，深度学习&机器学习高性能单机、分布式训练和跨平台部署） http://www.paddlepaddle.org.cn/http:/www.paddlepaddle.org/)
- [predictionio](https://github.com/apache/predictionio) - PredictionIO, a machine learning server for developers and ML engineers.
- [pytorch](https://github.com/pytorch/pytorch) - Tensors and Dynamic neural networks in Python with strong GPU acceleration [https://pytorch.org](https://pytorch.org/)
- [seldon-core](https://github.com/SeldonIO/seldon-core) - A framework to deploy, manage and scale your production machine learning to thousands of models.
- [sqlflow](https://github.com/sql-machine-learning/sqlflow) - Brings SQL and AI together.
- [tensorflow](https://github.com/tensorflow/tensorflow) - Computation using data flow graphs for scalable machine learning [https://www.tensorflow.org/](https://www.tensorflow.org/)

## API Gateway

- [ambassador](https://github.com/datawire/ambassador) - Ambassador: a self-service API gateway for microservices built on Lyft Envoy [https://www.getambassador.io/](https://www.getambassador.io/)
- [express-gateway](https://github.com/ExpressGateway/express-gateway) - A microservices API Gateway built on top of ExpressJS [https://www.express-gateway.io](https://www.express-gateway.io/)
- [kong](https://github.com/Kong/kong) - The Microservice API Gateway <https://konghq.com/install/>
- [krakend](https://github.com/devopsfaith/krakend) - Ultra performant API Gateway with middlewares. <https://www.krakend.io/>
- [orange](https://github.com/orlabs/orange) - OpenResty/Nginx Gateway for API Monitoring and Management. [http://orange.sumory.com](http://orange.sumory.com/)
- [ocelot](https://github.com/ThreeMammals/Ocelot) - .NET core API Gateway. [https://threemammals.com/ocelot](https://threemammals.com/ocelot)

## Application Delivery

- [argo-cd](https://github.com/argoproj/argo-cd/) - Declarative continuous deployment for Kubernetes.
- [argo](https://github.com/argoproj/argo) - Get stuff done with container-native workflows for Kubernetes.
- [arkade](https://github.com/alexellis/arkade) - Kubernetes apps for developers.
- [armada](https://github.com/att-comdev/armada) - A python orchestrator for a installing, upgrading, and managing a collection of helm charts, dependencies, and values overrides.
- [autoapply](https://github.com/autoapply/autoapply) - Automatically apply changes from a git repository to Kubernetes.
- [ballerina-lang](https://github.com/ballerina-platform/ballerina-lang) - Ballerina is a new programming language for integration built on a sequence diagram metaphor.
- [beetle](https://github.com/Clivern/Beetle) - Kubernetes multi-cluster deployment automation service.
- [binderhub](https://github.com/jupyterhub/binderhub) - Deterministically build docker images from a git repository + commit [https://binderhub.readthedocs.io/en/latest/en/latest/](https://binderhub.readthedocs.io/en/latest/en/latest/)
- [build](https://github.com/knative/build) - A Kubernetes-native Build resource.
- [cdk8s](https://github.com/awslabs/cdk8s) - Define Kubernetes native apps and abstractions using object-oriented programming [https://cdk8s.io](https://cdk8s.io/)
- [cds](https://github.com/ovh/cds) - Enterprise-Grade Continuous Delivery & DevOps Automation Open Source Platform.
- [charitify](https://github.com/kubepack/chartify) - Generate Helm Charts from Kubernetes objects by <https://appscode.com>
- [circleci](https://github.com/circleci) - Continuous Integration and Deployment.
- [cnab-spec](https://github.com/cnabio/cnab-spec) - Cloud Native Application Bundle Specification [https://cnab.io](https://cnab.io/)
- [commandeer](https://github.com/commandeer/open) - Cloud management desktop app for macOS, Windows, and Linux [https://getcommandeer.com](https://getcommandeer.com)
- [containerops](https://github.com/Huawei/containerops) - DevOps Orchestration Platform.
- [crane](https://github.com/Dataman-Cloud/crane) - Yet another control plane based on docker built-in swarmkit.
- [cross-cloud](https://github.com/crosscloudci/cross-cloud) - Cross Cloud Continuous Integration [https://www.cncf.io/](https://www.cncf.io/)
- [cue](https://github.com/cuelang/cue) - Validate and define text-based and dynamic configuration.
- [cyclone](https://github.com/caicloud/cyclone) - Powerful workflow engine and end-to-end pipeline solutions implemented with native Kubernetes resources.
- [draft](https://github.com/azure/draft) - A tool for developers to create cloud-native applications on Kubernetes. [https://draft.sh](https://draft.sh)
- [drone](https://github.com/drone/drone) - Drone is a Continuous Delivery platform built on Docker, written in Go [https://drone.io](https://drone.io/)
- [fabric8](https://github.com/fabric8io/fabric8) - fabric8 is an open source microservices platform based on Docker, Kubernetes and Jenkins [http://fabric8.io/](http://fabric8.io/)
- [flagger](https://github.com/weaveworks/flagger) - Progressive delivery Kubernetes operator (Canary, A/B Testing and Blue/Green deployments) [https://flagger.app](https://flagger.app/)
- [flux](https://github.com/fluxcd/flux) - A tool for turning container images into running Kubernetes services.
- [gitkube](https://github.com/hasura/gitkube) - Gitkube: Build and deploy docker images to Kubernetes using git push. [https://gitkube.sh](https://gitkube.sh/)
- [gockerize](https://github.com/redbooth/gockerize) - Package golang service into minimal docker containers.
- [habitus](https://github.com/cloud66-oss/habitus) - A build flow tool for Docker. https://help.cloud66.com/habitus/
- [helm](https://github.com/helm/helm) - The Kubernetes Package Manager.
- [helmsman](https://github.com/Praqma/helmsman) - Helm Charts as Code.
- [hiboot](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web and cli application framework with dependency injection support [https://hiboot.hidevops.io](https://hiboot.hidevops.io/)
- [hygieia](https://github.com/Hygieia/Hygieia) - CapitalOne DevOps Dashboard.
- [hyscale](https://github.com/hyscale/hyscale) - All things HyScale https://www.hyscale.io/
- [jenkins](https://github.com/jenkinsci/jenkins) - Jenkins automation server http://www.jenkins.io/
- [jib](https://github.com/GoogleContainerTools/jib) -  Build container images for your Java applications.
- [jsonnet](https://github.com/google/jsonnet) - Jsonnet - The data templating language [https://jsonnet.org/](https://jsonnet.org/)
- [jx](https://github.com/jenkins-x/jx) - A command line tool for installing and working with Jenkins X <https://jenkins-x.io/>
- [kaniko](https://github.com/GoogleContainerTools/kaniko) - Build Container Images In Kubernetes.
- [kd](https://github.com/UKHomeOffice/kd) - Minimalistic kubernetes resources deployment tool with templating.
- [kdo](https://github.com/stepro/kdo) - Deployless Development on Kubernetes.
- [kedge](https://github.com/kedgeproject/kedge) - Kedge - Concise Application Definition for Kubernetes [http://kedgeproject.org](http://kedgeproject.org/)
- [kenyata](https://github.com/spinnaker/kayenta) - Automated Canary Service.
- [keptn](https://github.com/keptn/keptn) - keptn is a control-plane for continuous delivery and operations enable cloud-native applications to run autonomously [https://keptn.sh](https://keptn.sh/)
- [kismatic](https://github.com/apprenda/kismatic) - Kismatic Enterprise Toolkit: Fully-Automated, Production-Grade Kubernetes Operations.
- [kompose](https://github.com/kubernetes/kompose) - Go from Docker Compose to Kubernetes [https://kompose.io/](https://kompose.io/)
- [kpt](https://github.com/GoogleContainerTools/kpt) - Kpt is a toolkit to help you manage, manipulate, customize, and apply Kubernetes Resource configuration data files.
- [kubeapps](https://github.com/kubeapps/kubeapps) - A web-based UI for deploying and managing applications in Kubernetes clusters [https://kubeapps.com](https://kubeapps.com/)
- [kubegen](https://github.com/errordeveloper/kubegen) - Kubegen – simple way to describe Kubernetes resources.
- [kubernetes-deploy](https://github.com/Shopify/krane) - A command-line tool that helps you ship changes to a Kubernetes namespace and understand the result.
- [kustomize](https://github.com/kubernetes-sigs/kustomize) - Customization of kubernetes YAML configurations.
- [lastbackend](https://github.com/lastbackend/lastbackend) - Container orchestration with CI&CD, cli and amazing UI [https://lastbackend.com](https://lastbackend.com/)
- [mkit](https://github.com/darkbitio/mkit) - MKIT is a Managed Kubernetes Inspection Tool that validates several common security-related configuration settings of managed Kubernetes cluster objects and the workloads/resources running inside the cluster.
- [opencompose](https://github.com/redhat-developer/opencompose) - OpenCompose - A higher level abstraction for Kubernetes Resource.
- [pipeline](https://github.com/banzaicloud/pipeline) - REST API to provision or reuse managed Kubernetes clusters in the cloud and deploy cloud native apps.
- [pipeline](https://github.com/tektoncd/pipeline) - A K8s-native Pipeline resource. [https://tekton.dev](https://tekton.dev/)
- [pulumi](https://github.com/pulumi/pulumi) -  A multi-language, multi-cloud development platform -- your code, your cloud, your team.
- [rudr](https://github.com/oam-dev/rudr) - A Kubernetes implementation of the Open Application Model specification.
- [skaffold](https://github.com/GoogleContainerTools/skaffold) - Easy and Repeatable Kubernetes Development.
- [smith](https://github.com/oracle/Smith) - Smith: A microcontainer builder.
- [source-to-image](https://github.com/openshift/source-to-image) - A tool for building/building artifacts from source and injecting into docker images.
- [spec](https://github.com/oam-dev/spec) - The Open Application Model specification [https://oam.dev](https://oam.dev/)
- [spinnaker](https://github.com/spinnaker/spinnaker) - Spinnaker is an open source, multi-cloud continuous delivery platform for releasing software changes with high velocity and confidence. [https://spinnaker.io/](https://spinnaker.io/)
- [terraform](https://github.com/hashicorp/terraform) - Terraform is a tool for building, changing, and combining infrastructure safely and efficiently. [https://www.terraform.io/](https://www.terraform.io/)
- [tilt](https://github.com/tilt-dev/tilt) - A multi-service dev environment for teams on Kubernetes [https://tilt.dev/](https://tilt.dev/)
- [wercker](https://github.com/wercker/wercker) - The Wercker CLI can be used to execute pipelines locally for both local development and easy introspection. [http://www.wercker.com/](http://www.wercker.com/)
- [woodpecker](https://github.com/laszlocph/woodpecker) - Fork of drone.io v0.8 since drone is not fully opensource anymore. [https://woodpecker.laszlo.cloud](https://woodpecker.laszlo.cloud)

## Big Data

- [fast-data-dev](https://github.com/lensesio/fast-data-dev) - Kafka Docker for development. Kafka, Zookeeper, Schema Registry, Kafka-Connect, Landoop Tools, 20+ connectors.
- [pachyderm](https://github.com/pachyderm/pachyderm) - Reproducible Data Science at Scale! <https://www.pachyderm.com/>
- [spark](https://github.com/apache-spark-on-k8s/spark) - Apache Spark enhanced with native Kubernetes scheduler back-end.
- [spark-on-kubernetes-helm](https://github.com/jahstreet/spark-on-kubernetes-helm) - Spark on Kubernetes infrastructure Helm charts repo.
- [wallaroo](https://github.com/WallarooLabs/wallaroo) - Ultrafast and elastic data processing [https://www.wallaroolabs.com](https://www.wallaroolabs.com/)

## Container Runtime

- [clear-containers](https://github.com/clearcontainers) - OCI (Open Containers Initiative) compatible runtime using Virtual Machines.
- [containerd](https://github.com/containerd/containerd) - An open and reliable container runtime [https://containerd.io](https://containerd.io/)
- [cri-containerd](https://github.com/containerd/cri) - Containerd-based implementation of Kubernetes Container Runtime Interface.
- [cri-o](https://github.com/cri-o/cri-o) - Open Container Initiative-based implementation of Kubernetes Container Runtime Interface [https://cri-o.io](https://cri-o.io/)
- [frakti](https://github.com/kubernetes/frakti) - The hypervisor-based container runtime for Kubernetes.
- [gvisor](https://github.com/google/gvisor) - Sandboxed Container Runtime.
- [hyperd](https://github.com/hyperhq/hyperd) - HyperContainer Daemon [http://www.hypercontainer.io](http://www.hypercontainer.io/)
- [katacontainers](https://katacontainers.io/) - Kata Containers is a new open source project building extremely lightweight virtual machines that seamlessly plug into the containers ecosystem.
- [moby](https://github.com/moby/moby) - Moby Project - a collaborative project for the container ecosystem to assemble container-based systems. [https://mobyproject.org/](https://mobyproject.org/)
- [pouch](https://github.com/alibaba/pouch) - Pouch is an open-source project created to promote the container technology movement.
- [railcar](https://github.com/oracle/railcar) - RailCar: Rust implementation of the Open Containers Initiative oci-runtime.
- [rkt](https://github.com/rkt/rkt) - `rkt` is a pod-native container engine for Linux. It is composable, secure, and built on standards.

## Database

- [arangodb](https://github.com/arangodb/arangodb) - ArangoDB is a native multi-model database with flexible data models for documents, graphs, and key-values. Build high performance applications using a convenient SQL-like query language or JavaScript extensions.
- [beringei](https://github.com/facebookarchive/beringei) - Beringei is a high performance, in-memory storage engine for time series data.
- [cockroachdb](https://github.com/cockroachdb/cockroach/) - CockroachDB - the open source, cloud-native SQL database. [https://www.cockroachlabs.com](https://www.cockroachlabs.com/)
- [couchdb](https://github.com/apache/couchdb) - Apache CouchDB is one of a new breed of database management systems.
- [etcd](https://github.com/etcd-io/etcd) - Distributed reliable key-value store for the most critical data of a distributed system[https://etcd.io/](https://etcd.io/)
- [influxdb](https://github.com/influxdata/influxdb) - Scalable datastore for metrics, events, and real-time analytics [https://www.influxdata.com/](https://www.influxdata.com/)
- [leveldb](https://github.com/google/leveldb) - LevelDB is a fast key-value storage library written at Google that provides an ordered mapping from string keys to string values.
- [mehdb](https://github.com/mhausenblas/mehdb) - Educational Kubernetes-native NoSQL datastore using StatefulSet and persistent volumes [https://www.openshift.com/blog/kubernetes-statatefulset-in-action](https://www.openshift.com/blog/kubernetes-statefulset-in-action/)
- [mongodb](https://github.com/mongodb/mongo) - MongoDB is an open source database that uses a document-oriented data model.
- [opentsdb](https://github.com/OpenTSDB/opentsdb) - A scalable, distributed Time Series Database. [http://opentsdb.net](http://opentsdb.net/)
- [redis](https://github.com/redis/redis) - Redis is an in-memory database that persists on disk. The data model is key-value, but many different kind of values are supported: Strings, Lists, Sets, Sorted Sets, Hashes, HyperLogLogs, Bitmaps. [https://redis.io/](https://redis.io/)
- [rethinkdb](https://github.com/rethinkdb/rethinkdb) - The open-source database for the realtime web. [https://rethinkdb.com](https://rethinkdb.com/)
- [sharding-sphere](https://github.com/apache/shardingsphere) - Distributed database middleware.
- [stolon](https://github.com/sorintlab/stolon) - PostgreSQL cloud native High Availability and more.
- [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed NewSQL database compatible with MySQL protocol [https://pingcap.com](https://pingcap.com/)

## Edge Computing

- [baetyl](https://github.com/baetyl/baetyl) - Extend cloud computing, data and service seamlessly to edge devices. https://baetyl.io/
- [eliot](https://github.com/ernoaapa/eliot) - Open source system for managing containerized applications in IoT device [https://docs.eliot.run/](https://docs.eliot.run/)
- [iotedge](https://github.com/Azure/iotedge) - The IoT Edge OSS project.
- [k3s](https://github.com/rancher/k3s) - Lightweight Kubernetes. 5 less than k8s. [https://k3s.io](https://k3s.io/)
- [kubeedge](https://github.com/kubeedge/kubeedge) - Kubernetes Native Edge Computing Framework (project under CNCF) [https://kubeedge.io](https://kubeedge.io/)
- [openyurt](https://github.com/alibaba/openyurt) - Extending your native Kubernetes to the edge.

## Fault Tolerant

- [chaosblade](https://github.com/chaosblade-io/chaosblade) - An easy to use and powerful chaos engineering experiment toolkit（阿里巴巴开源的一款简单易用、功能强大的混沌实验注入工具）.
- [chaosmonkey](https://github.com/Netflix/chaosmonkey) - Chaos Monkey is a resiliency tool that helps applications tolerate random instance failures.
- [chaos-mesh](https://github.com/pingcap/chaos-mesh) - A Chaos Engineering Platform for Kubernetes.
- [concurrency-limits](https://github.com/Netflix/concurrency-limits) - Java Library that implements and integrates concepts from TCP congestion control to auto-detect concurrency limits to achieve optimal throughput with optimal latency.
- [hystrix](https://github.com/Netflix/Hystrix) - Hystrix is a latency and fault tolerance library designed to isolate points of access to remote systems, services and 3rd party libraries, stop cascading failure and enable resilience in complex distributed systems where failure is inevitable.
- [ratelimit](https://github.com/envoyproxy/ratelimit) - Go/gRPC service designed to enable generic rate limit scenarios from different types of applications.
- [sentinel](https://github.com/alibaba/sentinel) - A powerful flow control component enabling reliability, resilience and monitoring for microservices. (面向云原生微服务的高可用流控防护组件)

## Kubernetes Operators

- [banzaicloud/bank-vaults](https://github.com/banzaicloud/bank-vaults) - A Vault swiss-army knife: a K8s operator, Go client with automatic token renewal, automatic configuration, multiple unseal options and more. A CLI tool to init, unseal and configure Vault (auth methods, secret engines). Direct secret injection into Pods.
- [eunomia](https://github.com/KohlsTechnology/eunomia) - A GitOps Operator for Kubernetes.
- [flagger](https://github.com/weaveworks/flagger) - Istio progressive delivery Kubernetes operator [https://docs.flagger.app](https://docs.flagger.app/)
- [keel](https://github.com/keel-hq/keel) - Kubernetes Operator to automate Helm, DaemonSet, StatefulSet & Deployment updates [https://keel.sh](https://keel.sh/)
- [kopf](https://github.com/zalando-incubator/kopf) - A Python framework to write Kubernetes operators in just few lines of code. [https://kopf.readthedocs.io/en/latest/en/latest/](https://kopf.readthedocs.io/en/latest/en/latest/)
- [kudo](https://github.com/kudobuilder/kudo) - Kubernetes Universal Declarative Operator (KUDO) [https://kudo.dev](https://kudo.dev/)
- [kubevirt](https://github.com/kubevirt/kubevirt) - Kubernetes Virtualization Operator with API and runtime in order to define and manage virtual machines.
- [operator-lifecycle-manager](https://github.com/operator-framework/operator-lifecycle-manager) - A management framework for extending Kubernetes with Operators.
- [operator-sdk](https://github.com/operator-framework/operator-sdk) - SDK for building Kubernetes applications. Provides high level APIs, useful abstractions, and project scaffolding. <https://coreos.com/operators/>
- [prometheus-operator](https://github.com/coreos/prometheus-operator) - Prometheus Operator creates/configures/manages Prometheus clusters atop Kubernetes<https://coreos.com/operators//prometheus>
- [spark-on-k8s-operator](https://github.com/GoogleCloudPlatform/spark-on-k8s-operator) - Kubernetes operator for managing the lifecycle of Apache Spark applications on Kubernetes.
- [strimzi-kafka-operator](https://github.com/strimzi/strimzi-kafka-operator) - Apache Kafka running on Kubernetes https://strimzi.io/
- [tidb-operator](https://github.com/pingcap/tidb-operator) - TiDB operator creates and manages TiDB clusters running in Kubernetes.

## Logging

- [beats](https://github.com/elastic/beats) - Beats - Lightweight shippers for Elasticsearch & Logstash [https://www.elastic.co/beats/](https://www.elastic.co/beats/)
- [collectbeat](https://github.com/eBay/collectbeat) - Beats with discovery capabilities for environments like Kubernetes.
- [elasticsearch](https://github.com/elastic/elasticsearch) - Open Source, Distributed, RESTful Search Engine [https://www.elastic.co/elasticsearch](https://www.elastic.co/elasticsearch)
- [fluent-bit](https://github.com/fluent/fluent-bit) - Fast and Lightweight Log/Data Forwarder for Linux, BSD and macOS <https://fluentbit.io/>
- [fluentd-pilot](https://github.com/AliyunContainerService/log-pilot) - Collect logs in docker containers.
- [fluentd](https://github.com/fluent/fluentd) - Fluentd: Unified Logging Layer (project under CNCF) [https://www.fluentd.org/](https://www.fluentd.org/)
- [flume](http://flume.apache.org/) - Flume is a distributed, reliable, and available service for efficiently collecting, aggregating, and moving large amounts of log data.
- [heapster](https://github.com/kubernetes-retired/heapster) - Compute Resource Usage Analysis and Monitoring of Container Clusters.
- [log-pilot](https://github.com/AliyunContainerService/log-pilot) - Collect logs in docker containers.
- [loki](https://github.com/grafana/loki) - Like Prometheus, but for logs.
- [telegraf](https://github.com/influxdata/telegraf) - The plugin-driven server agent for collecting & reporting metrics.

## Message Broker

- [emqx](https://github.com/emqx/emqx) - EMQ X Broker - Scalable Distributed MQTT Message Broker for IoT in 5G Era.
- [flume](https://github.com/apache/flume) - Apache Flume is a distributed, reliable, and available service for efficiently collecting, aggregating, and moving large amounts of log data.
- [gnatsd](https://github.com/nats-io/nats-server) - High-Performance server for NATS, the cloud native messaging system. [https://nats.io](https://nats.io/)
- [jocko](https://github.com/travisjeffery/jocko) - Kafka implemented in Golang with built-in coordination (No ZK dep, single binary install, Cloud Native) <https://twitter.com/travisjeffery>
- [kafka](https://github.com/apache/kafka) - A distributed streaming platform.
- [nsq](https://github.com/nsqio/nsq) - A realtime distributed messaging platform [https://nsq.io/](https://nsq.io/)
- [rabbitmq](https://github.com/rabbitmq) - RabbitMQ is the most widely deployed open source message broker.
- [rocketmq](https://github.com/apache/rocketmq) - Apache RocketMQ is a distributed messaging and streaming platform with low latency, high performance and reliability, trillion-level capacity and flexible scalability.

## Miscellaneous

- [agones](https://github.com/googleforgames/agones) - Dedicated Game Server Hosting and Scaling for Multiplayer Games on Kubernetes.

## Monitoring

- [cadvisor](https://github.com/google/cadvisor) - Analyzes resource usage and performance characteristics of running containers.
- [cortex](https://github.com/cortexproject/cortex) - A multitenant, horizontally scalable Prometheus as a Service.
- [elasticsearch-hq](https://github.com/ElasticHQ/elasticsearch-HQ) - Monitoring and Management Web Application for ElasticSearch instances and clusters.
- [grafana](https://github.com/grafana/grafana) - The tool for beautiful monitoring and metric analytics & dashboards for Graphite, InfluxDB & Prometheus & More [https://grafana.com](https://grafana.com/)
- [hawkular-metrics](https://github.com/hawkular/hawkular-metrics) - Time Series Metrics Engine based on Cassandra.
- [kibana](https://github.com/elastic/kibana) - Kibana analytics and search dashboard for Elasticsearch [https://www.elastic.co/kibana](https://www.elastic.co/kibana)
- [kubernetes-zabbix](https://github.com/monitoringartist/kubernetes-zabbix) - Kubernetes Zabbix/Grafana cluster (bare metal, Google Computer Engine - GCE, Google Container Engine - GKE).
- [nexclipper](https://github.com/NexClipper/NexClipper) - An open source software for monitoring Kubernetes and containers.
- [open-falcon](https://github.com/XiaoMi/open-falcon) - Enterprise Internet monitoring system from Xiaomi [http://open-falcon.com/](http://open-falcon.com/)
- [owl](https://github.com/TalkingData/owl) - Distributed monitoring system from TalkingData.
- [prometheus](https://github.com/prometheus/prometheus) - The Prometheus monitoring system and time series database. [https://prometheus.io/](https://prometheus.io/)
- [scope](https://github.com/weaveworks/scope) - Monitoring, visualisation & management for Docker & Kubernetes [https://www.weave.works/oss/scope/](https://www.weave.works/oss/scope/)
- [sofa-lookout](https://github.com/sofastack/sofa-lookout) - Lookout can help you to measure and monitor the status of the target system with its multi-dimensional metrics.
- [statsd](https://github.com/statsd/statsd) - Daemon for easy but powerful stats aggregation.

## Network

- [calico](https://github.com/projectcalico) - A Pure Layer 3 Approach to Virtual Networking for Highly Scalable Data Centers.
- [cilium](https://github.com/cilium/cilium) - API-aware Networking and Security using eBPF and XDP [https://cilium.io/](https://cilium.io/)
- [cni](https://github.com/containernetworking/cni) - Container Network Interface - networking for Linux containers.
- [cni-genie](https://github.com/cni-genie/CNI-Genie) - CNI-Genie for choosing pod network of your choice during deployment time. Supported pod networks - Calico, Flannel, Romana, Weave.
- [contiv](https://github.com/contiv) - Container networking for various use cases.
- [flannel](https://github.com/coreos/flannel) - Flannel is a network fabric for containers, designed for Kubernetes.
- [hubble](https://github.com/cilium/hubble) - Hubble - Network, Service & Security Observability for Kubernetes.
- [istio-cni](https://github.com/istio/cni) - Istio CNI to setup kubernetes pod namespaces to redirect traffic to sidecar proxy.
- [knitter](https://github.com/ZTE/Knitter) - Kubernetes network solution.
- [kube-router](https://github.com/cloudnativelabs/kube-router) - Kube-router, a turnkey solution for Kubernetes networking. [https://www.kube-router.io/](https://www.kube-router.io/)
- [kube-ovn](https://github.com/alauda/kube-ovn) - Kube-OVN, a Kubernetes network fabric for enterprises that is rich in functions and easy in operations.
- [matchbox](https://github.com/poseidon/matchbox) - Network boot and provision Container Linux clusters (e.g. etcd3, Kubernetes, more) <https://coreos.com/matchbox/docs/latest/>
- [submariner](https://github.com/submariner-io/submariner) - Connect all your Kubernetes clusters, no matter where they are in the world. [https://submariner.io](https://submariner.io/)
- [weave](https://github.com/weaveworks/weave) - Simple, resilient multi-host Docker networking and more. [https://www.weave.works](https://www.weave.works/)

## Observability

- [envoy-ui](https://github.com/Nitro/envoy-ui) - Dead simple server-side UI for Envoy proxy (like HAproxy stats).
- [goldpinger](https://github.com/bloomberg/goldpinger) - Debugging tool for Kubernetes which tests and displays connectivity between nodes in the cluster.
- [istio-ui](https://github.com/jukylin/istio-ui) - Istio config management backend.
- [kiali](https://github.com/kiali/kiali) - Kiali project to help istio service mesh observability.
- [konstellate](https://github.com/containership/konstellate) - Free and Open Source GUI to Visualize Kubernetes Applications.
- [kube-ops-view](https://github.com/hjacobs/kube-ops-view) - Kubernetes Operational View - read-only system dashboard for multiple K8s clusters [https://kubernetes-operational-view.readthedocs.io/en/latest/en/latest/](https://kubernetes-operational-view.readthedocs.io/en/latest/en/latest/)
- [naftis](https://github.com/XiaoMi/naftis) - An excellent dashboard for Istio built with love.
- [octant](https://github.com/vmware-tanzu/octant) - Highly extensible platform for developers to better understand the complexity of Kubernetes clusters. [https://octant.dev](https://octant.dev)
- [vistio](https://github.com/nmnellis/vistio) - Visualize your Istio mesh using Netflix's Vizceral.
- [vizceral](https://github.com/Netflix/vizceral) - WebGL visualization for displaying animated traffic graphs.

## Orchestration and Scheduler

- [alameda](https://github.com/containers-ai/alameda) - Intelligent Resources Orchestrator for Kubernetes by using machine learning.
- [blox](https://github.com/blox/blox) - Open source tools for building custom schedulers on Amazon ECS.
- [compose](https://github.com/docker/compose) - Define and run multi-container applications with Docker.
- [conductor](https://github.com/Netflix/conductor) - Conductor is a microservices orchestration engine.
- [descheduler](https://github.com/kubernetes-sigs/descheduler) - Descheduler for Kubernetes.
- [domeos](https://github.com/domeos) - An enterprise application orchestration system based on docker by Sohu.com - http://domeos.org
- [fleet](https://github.com/coreos/fleet) - Fleet ties together systemd and etcd into a distributed init system.
- [kruise](https://github.com/openkruise/kruise) - Automate application workloads management on Kubernetes [https://openkruise.io](https://openkruise.io/)
- [kubernetes](https://github.com/kubernetes/kubernetes) - Production-Grade Container Scheduling and Management [https://kubernetes.io/](https://kubernetes.io/)
- [marathon](https://github.com/mesosphere/marathon) - Deploy and manage containers (including Docker) on top of Apache Mesos at scale.
- [mesos](https://github.com/apache/mesos) - Apache Mesos abstracts CPU, memory, storage, and other compute resources away from machines (physical or virtual), enabling fault-tolerant and elastic distributed systems to easily be built and run effectively.
- [serf](https://github.com/hashicorp/serf) - Service orchestration and management tool by hashicorp. [https://www.serf.io/](https://www.serf.io/)
- [service-fabric](https://github.com/Microsoft/service-fabric) - Service Fabric is a distributed systems platform for packaging, deploying, and managing stateless and stateful distributed applications and containers at large scale. [https://docs.microsoft.com/en-us/azure/service-fabric/](https://docs.microsoft.com/en-us/azure/service-fabric/)
- [swan](https://github.com/Dataman-Cloud/swan) - A Distributed, Highly Available Mesos Scheduler, Inspired by the design of Google Borg.
- [swarm](https://github.com/docker/classicswarm) - Swarm: a Docker-native clustering system.

## PaaS Manager

- [breeze](https://github.com/wise2c-devops/breeze) -  Wise2C ansible playbook for Kubernetes cluster installation.
- [choerodon](https://github.com/choerodon/choerodon) - The open source PaaS for Kubernetes.[http://choerodon.io](http://choerodon.io/)
- [cloudfoundry](https://github.com/cloudfoundry) - Cloud Foundry is an open source, multi cloud application platform as a service (PaaS) governed by the Cloud Foundry Foundation.
- [conjure-up](https://github.com/conjure-up/conjure-up) - Deploying complex solutions, magically. [https://conjure-up.io](https://conjure-up.io/)
- [crossplane](https://github.com/crossplane/crossplane) - An Open Source Multicloud Control Plane [https://crossplane.io](https://crossplane.io/)
- [dc/os](https://github.com/dcos) - Datacenter Operating System.
- [deis](https://github.com/deis/deis) - Deis v1, the CoreOS and Docker PaaS: Your PaaS. Your Rules.
- [flynn](https://github.com/flynn/flynn) - A next generation open source platform as a service (PaaS) [https://flynn.io](https://flynn.io/)
- [kqeen](https://github.com/Mirantis/kqueen) - Kubernetes queen - cluster manager.
- [kubermatic](https://github.com/kubermatic/kubermatic) - The Central Kubernetes Management Platform For Any Infrastructure.
- [kubernator](https://github.com/smpio/kubernator) - Alternative Kubernetes UI.
- [kubesphere](https://github.com/kubesphere/kubesphere) - Enterprise Container Managent Platform https://kubesphere.io
- [oneinfra](https://github.com/oneinfra/oneinfra) - Kubernetes as a Service.
- [opendcp](https://github.com/weibocom/opendcp) - Docker platform developed by [weibo.com](http://weibo.com/us)
- [openshift](https://github.com/openshift/origin) - Enterprise Kubernetes for Developers [https://www.openshift.org/](https://www.openshift.org/)
- [rainbond](https://github.com/goodrain/rainbond) - Serverless PaaS , A new generation of easy-to-use cloud management platforms based on kubernetes.[https://www.rainbond.com/](https://www.rainbond.com/)
- [rancher](https://github.com/rancher/rancher) - Complete container management platform [https://rancher.com/](https://rancher.com/)
- [supergiant](https://github.com/supergiant/control) - Automatically scale hardware and easily run stateful applications using Kubernetes. <https://supergiant.io/thank-you/>
- [wayne](https://github.com/Qihoo360/wayne) - Web UI for Kubernetes multi-clusters.
- [vamp](https://github.com/magneticio/vamp) - Vamp - canary releasing and autoscaling for microservice systems [https://vamp.io/](https://vamp.io/)
- [volcano](https://github.com/volcano-sh/volcano) - A Kubernetes Native Batch System (Project under CNCF).

## Proxy

- [caddy](https://github.com/caddyserver/caddy) - Fast, cross-platform HTTP/2 web server with automatic HTTPS [https://caddyserver.com](https://caddyserver.com/)
- [contour](https://github.com/projectcontour/contour) - Contour is a Kubernetes ingress controller for Lyft's Envoy proxy.
- [envoy-docker-shim](https://github.com/Nitro/envoy-docker-shim) - Run Envoy in place of docker-proxy.
- [envoy](https://github.com/envoyproxy/envoy) - C++ front/service proxy [https://www.envoyproxy.io/](https://www.envoyproxy.io/)
- [gimbal](https://github.com/projectcontour/gimbal) - Heptio Gimbal is an ingress load balancing platform capable of routing traffic to multiple Kubernetes and OpenStack clusters. Built by Heptio in partnership with Actapio. [https://heptio.cloud.vmware.com/](https://heptio.cloud.vmware.com/)
- [gobetween](https://github.com/yyyar/gobetween) - Modern & minimalistic load balancer for the Сloud era [http://gobetween.io](http://gobetween.io/)
- [haproxy](https://github.com/haproxy/haproxy) - HAProxy is a free, very fast and reliable solution offering high availability, load balancing, and proxying for TCP and HTTP-based applications.
- [kEdge](https://github.com/improbable-eng/kedge) - kEdge - Kubernetes Edge Proxy for gRPC and HTTP Microservices.
- [katran](https://github.com/facebookincubator/katran) - A high performance layer 4 load balancer.
- [kong-ingress](https://github.com/koli/kong-ingress) - A Kubernetes Ingress for Kong.
- [kong/kubernetes-ingress-controller](https://github.com/Kong/kubernetes-ingress-controller) - Deploy Kong in a native Kubernetes Ingress Controller <https://konghq.com/>
- [metallb](https://github.com/metallb/metallb) - A network load-balancer implementation for Kubernetes using standard routing protocols [https://metallb.universe.tf](https://metallb.universe.tf/)
- [mosn](https://github.com/mosn/mosn) - MOSN is a cloud native proxy for edge or service mesh. https://mosn.io
- [nginx-kubernetes-ingress](https://github.com/nginxinc/kubernetes-ingress) - NGINX and NGINX Plus Ingress Controllers for Kubernetes.
- [nginx](https://github.com/nginx/nginx) - nginx is an HTTP and reverse proxy server, a mail proxy server, and a generic TCP/UDP proxy server, originally written by [Igor Sysoev](http://sysoev.ru/en/).
- [ribbon](https://github.com/Netflix/ribbon) - Ribbon is a Inter Process Communication (remote procedure calls) library with built in software load balancers. The primary usage model involves REST calls with various serialization scheme support.
- [skipper](https://github.com/zalando/skipper) - An HTTP router and reverse proxy for service composition, including use cases like Kubernetes Ingress.
- [traefik](https://github.com/containous/traefik) - Træfik, a modern reverse proxy [https://containo.us/traefik/](https://containo.us/traefik/)
- [voyager](https://github.com/voyagermesh/voyager) - Secure Ingress Controller for Kubernetes by <https://appscode.com>

## RPC

- [brpc](https://github.com/apache/incubator-brpc) - Most common RPC framework used throughout Baidu, with 600,000+ instances and 500+ kinds of services, called "baidu-rpc" inside Baidu.
- [finagle](https://github.com/twitter/finagle) - A fault tolerant, protocol-agnostic RPC system.
- [grpc](https://github.com/grpc) - A high performance, open source, general-purpose RPC framework.
- [proxygen](https://github.com/facebook/proxygen) - A collection of C++ HTTP libraries including an easy to use HTTP server.
- [rsocket](https://github.com/rsocket) - Streaming message protocol with Reactive Extension/Stream semantics.
- [sofa-bolt](https://github.com/sofastack/sofa-bolt) -  SOFABolt is a lightweight, easy to use and high performance remoting framework based on Netty.
- [sofa-rpc](https://github.com/sofastack/sofa-rpc) - SOFARPC is a high-performance, high-extensibility, production-level Java RPC framework.
- [tars](https://github.com/TarsCloud/Tars) - Tars is a high-performance RPC framework based on name service and Tars protocol, also integrated administration platform, and implemented hosting-service via flexible schedule.
- [thrift](https://github.com/apache/thrift) - Apache thrift.

## Security and Audit

- [apparmor](https://gitlab.com/apparmor/apparmor/wikis/home/) - AppArmor is an effective and easy-to-use Linux application security system.
- [authenticator](https://github.com/kubernetes-sigs/aws-iam-authenticator) - A tool for using AWS IAM credentials to authenticate to a Kubernetes cluster.
- [cert-manager](https://github.com/jetstack/cert-manager) - Automatically provision and manage TLS certificates in Kubernetes.
- [checkov](https://github.com/bridgecrewio/checkov/) - A static analysis tool for infrastructure as code - to prevent misconfigs at build time.
- [clair](https://github.com/quay/clair) - Vulnerability Static Analysis for Containers.
- [cost-model](https://github.com/kubecost/cost-model) - Cross-cloud cost allocation models for workloads running on Kubernetes [https://kubecost.com/](https://kubecost.com/)
- [dex](https://github.com/dexidp/dex) - OpenID Connect Identity (OIDC) and OAuth 2.0 Provider with Pluggable Connectors [https://coreos.com/blog/announcing-dex.html](https://coreos.com/blog/announcing-dex.html)
- [docker-bench-security](https://github.com/docker/docker-bench-security) - The Docker Bench for Security is a script that checks for dozens of common best-practices around deploying Docker containers in production.
- [dockscan](https://github.com/kost/dockscan) - Dockscan is security vulnerability and audit scanner for Docker installations.
- [drydock](https://github.com/zuBux/drydock) - Drydock provides a flexible way of assessing the security of your Docker daemon configuration and containers using editable audit templates.
- [falco](https://github.com/falcosecurity/falco) - Behavioral Activity Monitoring With Container Support.
- [goldfish](https://github.com/Caiyeon/goldfish) - A HashiCorp Vault UI panel written with VueJS and Vault native Go API.
- [grafeas](https://github.com/Grafeas/Grafeas) - Cloud artifact metadata CRUD API and resource specifications.
- [guard](https://github.com/appscode/guard) - Kubernetes Authentication WebHook Server.
- [k8guard](https://github.com/k8guard) - An auditing system for Kubernetes.
- [keycloak](https://github.com/keycloak/keycloak) - Open Source Identity and Access Management For Modern Applications and Services [https://www.keycloak.org/](https://www.keycloak.org/)
- [kritis](https://github.com/grafeas/kritis) - Software supply chain security for #Kubernetes apps [https://grafeas.io/docs/concepts/what-is-kritis](https://grafeas.io/docs/concepts/what-is-kritis)
- [kube-bench](https://github.com/aquasecurity/kube-bench) - The Kubernetes Bench for Security is a Go application that checks whether Kubernetes is deployed according to security best practices.
- [kube-lego](https://github.com/jetstack/kube-lego) - Automatically request certificates for Kubernetes Ingress resources from Let's Encrypt.
- [kube2iam](https://github.com/jtblin/kube2iam) - kube2iam provides different AWS IAM roles for pods running on Kubernetes.
- [kubed](https://github.com/appscode/kubed) - A Kubernetes Cluster Operator Daemon by <https://appscode.com>
- [notary](https://github.com/theupdateframework/notary) - Notary is a Docker project that allows anyone to have trust over arbitrary collections of data [https://www.docker.com/](https://www.docker.com/)
- [opa](https://github.com/open-policy-agent/opa) - An open source project to policy-enable your service. [https://openpolicyagent.org/](https://openpolicyagent.org/)
- [openscap](https://www.open-scap.org/) - Discover a wide array of tools for managing system security and standards compliance.
- [pomerium](https://github.com/pomerium/pomerium/) - Pomerium is a zero-trust context and identity aware access gateway inspired by BeyondCorp.
- [spiffe](https://github.com/spiffe/spiffe) - The SPIFFE Project [https://spiffe.io/](https://spiffe.io/)
- [trivy](https://github.com/aquasecurity/trivy) - A Simple and Comprehensive Vulnerability Scanner for Containers, Suitable for CI.
- [vault](https://github.com/hashicorp/vault) - A tool for managing secrets. <https://www.vaultproject.io>

## Service Mesh

- [amalgam8](https://github.com/amalgam8/amalgam8) - Content and Version-based Routing Fabric for Polyglot Microservices.
- [conduit](https://conduit.io) - The Ultralight Service Mesh for Kubernetes [https://conduit.io](https://conduit.io/)
- [consul](https://github.com/hashicorp/consul) - Consul is a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure. [https://www.consul.io](https://www.consul.io)
- [istio](https://github.com/istio/istio) - Connect, secure, control, and observe services. [https://istio.io](https://istio.io/)
- [kuma](https://github.com/Kong/kuma) - Universal Control Plane for your Service Mesh https://kuma.io/install
- [linkerd](https://github.com/linkerd/linkerd) - Resilient service mesh for cloud native apps [https://linkerd.io](https://linkerd.io/)
- [maesh](https://github.com/containous/maesh) - Simpler Service Mesh [https://containo.us/maesh/](https://containo.us/maesh/)
- [nginmesh](https://github.com/nginxinc/nginmesh) - Service Mesh using Nginx.
- [nginx-unit](https://github.com/nginx/unit) - NGINX Unit is a new, lightweight, open source application server built to meet the demands of dynamic and distributed applications.
- [servicecomb](https://github.com/ServiceComb) - ServiceComb is a microservice framework that provides an easy way to develop and deploy applications in the cloud.
- [supergloo](https://github.com/solo-io/service-mesh-hub) - The Service Mesh Orchestration Platform [https://docs.solo.io/service-mesh-hub/latest/](https://docs.solo.io/service-mesh-hub/latest/)

## Service Registry and Discovery

- [apollo](https://github.com/ctripcorp/apollo) - Apollo（阿波罗）是携程框架部门研发的分布式配置中心，能够集中化管理应用不同环境、不同集群的配置，配置修改后能够实时推送到应用端，并且具备规范的权限、流程治理等特性，适用于微服务配置管理场景.
- [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul.
- [coredns](https://github.com/coredns/coredns) - CoreDNS is a DNS server that chains middleware [https://coredns.io](https://coredns.io/)
- [eureka](https://github.com/Netflix/eureka) - AWS Service registry for resilient mid-tier load balancing and failover.
- [open-service-broker-sdk](https://github.com/openshift/open-service-broker-sdk) - A starting point for creating service brokers implementing the Open Service Broker API.
- [registrator](https://github.com/gliderlabs/registrator) - Service registry bridge for Docker with pluggable adapters.
- [rotor](https://github.com/turbinelabs/rotor) - Rotor is a fast, lightweight bridge between your service discovery and the configuration APIs of Envoy. Rotor supports Kubernetes, Consul, AWS (EC2 and ECS), DC/OS, flat files, and even other EDS/CDS implementations.
- [service-broker](https://github.com/openservicebrokerapi/servicebroker) - Open Service Broker API Specification <https://www.openservicebrokerapi.org/>
- [service-catalog](https://github.com/kubernetes-sigs/service-catalog) - Consume services in Kubernetes using the Open Service Broker API.
- [skydns](https://github.com/skynetservices/skydns1) - DNS for skynet or any other service discovery.
- [steward](https://github.com/deis/steward) - The Kubernetes-native Service Broker.
- [synapse](https://github.com/airbnb/synapse) - A transparent service discovery framework for connecting an SOA.
- [vulcand](https://github.com/vulcand/vulcand) - Programmatic load balancer backed by Etcd [https://vulcand.readthedocs.io/](https://vulcand.readthedocs.io/)
- [zookeeper](https://github.com/apache/zookeeper) - Apache ZooKeeper is an effort to develop and maintain an open-source server which enables highly reliable distributed coordination.

## Serverless

- [dapr](https://github.com/dapr/dapr) - Dapr is a portable, event-driven, runtime for building distributed applications across cloud and edge.
- [dispatch](https://github.com/vmware/dispatch) - Dispatch is a framework for deploying and managing serverless style applications.
- [eventing](https://github.com/knative/eventing) - Open source specification and implementation of Knative event binding and delivery.
- [faas-netes](https://github.com/openfaas/faas-netes) - Enable Kubernetes as a backend for Functions as a Service (OpenFaaS).
- [firecamp](https://github.com/cloudstax/firecamp) - Serverless Platform for the stateful services.
- [firecracker](https://github.com/firecracker-microvm/firecracker) - Secure and fast microVMs for serverless computing.
- [fission](https://github.com/fission/fission) - Fast Serverless Functions for Kubernetes https://fission.io/
- [fn](https://github.com/fnproject/fn) - The container native, cloud agnostic serverless platform. [https://fnproject.io/](https://fnproject.io/)
- [funktion](https://github.com/funktionio/funktion/) - a CLI tool for working with funktion <https://funktion.fabric8.io/>
- [fx](https://github.com/metrue/fx) - Poor man's serverless framework based on Docker, Function as a Service with painless.
- [gloo](https://github.com/solo-io/gloo) - The Function Gateway built on top of Envoy.
- [ironfunctions](https://github.com/iron-io/functions) - IronFunctions - the serverless microservices platform. [https://www.iron.io/](https://www.iron.io/)
- [keda](https://github.com/kedacore/keda) - KEDA is a Kubernetes-based Event Driven Autoscaling component. It provides event driven scale for any container running in Kubernetes [https://keda.sh](https://keda.sh/)
- [knative-lambda-runtime](https://github.com/triggermesh/knative-lambda-runtime) - Running AWS Lambda Functions on Knative/Kubernetes Clusters [https://triggermesh.com](https://triggermesh.com/)
- [knix](https://github.com/knix-microfunctions/knix) - KNIX MicroFunctions is a serverless computing platform that combines container-based resource isolation with a lightweight execution model using processes to significantly improve resource efficiency and decrease the function startup latency. KNIX MicroFunctions works in Knative as well as bare metal or virtual machine-based environments.
- [kubeless](https://github.com/kubeless/kubeless) - Kubernetes Native Serverless Framework [https://kubeless.io/](https://kubeless.io/)
- [nuclio](https://github.com/nuclio/nuclio) - High-Performance Serverless event and data processing platform.
- [openfaas](https://github.com/openfaas/faas) - OpenFaaS - Serverless Functions Made Simple for Docker & Kubernetes [https://blog.alexellis.io/introducing-functions-as-a-service/](https://blog.alexellis.io/introducing-functions-as-a-service/)
- [openwhisk](http://openwhisk.apache.org/) - Apache OpenWhisk (Incubating) is a [serverless](https://en.wikipedia.org/wiki/Serverless_computing), open source cloud platform that executes functions in response to events at any scale.
- [osiris](https://github.com/deislabs/osiris) - A general purpose, scale-to-zero component for Kubernetes.
- [riff](https://github.com/projectriff/riff) - riff is for functions [https://projectriff.io](https://projectriff.io/)
- [serverless](https://github.com/serverless/serverless) - Serverless Framework – Build web, mobile and IoT applications with serverless architectures using AWS Lambda, Azure Functions, Google CloudFunctions & more! – [https://www.serverless.com/](https://www.serverless.com/)
- [serving](https://github.com/knative/serving) - Kubernetes-based, scale-to-zero, request-driven compute.
- [spec](https://github.com/cloudevents/spec) - CloudEvents Specification [https://cloudevents.io](https://cloudevents.io/)
- [sqoop](https://github.com/solo-io/sqoop) - The GraphQL Engine powered by Gloo [https://sqoop.solo.io](https://sqoop.solo.io/)
- [thanos](https://github.com/thanos-io/thanos) - Highly available Prometheus setup with long term storage capabilities.

## Storage

- [ceph](https://github.com/ceph/ceph) - Ceph is a distributed object, block, and file storage platform https://ceph.com/
- [chubaofs](https://github.com/chubaofs/chubaofs) - A distributed storage system for cloud native applications http://chubao.io/
- [convoy](https://github.com/rancher/convoy) - A Docker volume plugin, managing persistent container volumes.
- [fastdfs](https://github.com/happyfish100/fastdfs) - FastDFS is an open source high performance distributed file system (DFS). It's major functions include: file storing, file syncing and file accessing, and design for high capacity and load balance.
- [flocker](https://github.com/ClusterHQ/flocker) - Container data volume manager for your Dockerized application.
- [glusterd2](https://github.com/gluster/glusterd2) -  GlusterD-2.0 is the distributed management framework to be used for GlusterFS-4.0.
- [glusterfs](https://github.com/gluster/glusterfs) - Gluster is a software defined distributed storage that can scale to several petabytes. It provides interfaces for object, block and file storage.
- [harbor](https://github.com/goharbor/harbor) - An open source trusted cloud native registry project that stores, signs, and scans content. https://goharbor.io/
- [heketi](https://github.com/heketi/heketi) - RESTful based volume management framework for GlusterFS.
- [infinit](https://github.com/infinit/infinit) - The Infinit policy-based software-defined storage platform. https://infinit.sh/
- [leofs](https://leo-project.net/leofs/) - The LeoFS Storage System https://leo-project.net/leofs/
- [longhorn](https://github.com/longhorn/longhorn) - We put storage on cows and move them around from rancher.
- [minio](https://github.com/minio/minio) - Minio is an open source object storage server compatible with Amazon S3 APIs https://min.io/
- [openebs](https://github.com/openebs/openebs) - OpenEBS is containerized block storage written in Go for cloud native and other environments w/ per container (or pod) QoS SLAs, tiering and replica policies across AZs and environments, and predictable and scalable performance. [https://openebs.io/](https://openebs.io/)
- [rook](https://github.com/rook/rook) - File, Block, and Object Storage Services for your Cloud-Native Environment https://rook.io
- [storageos](https://storageos.com/) - Enterprise persistent storage for containers and the cloud.
- [torus](https://github.com/coreos/torus) - Torus Distributed Storage https://coreos.com/blog/torus-distributed-storage-by-coreos.html
- [vitess](https://github.com/vitessio/vitess) - Vitess is a database clustering system for horizontal scaling of MySQL. [https://vitess.io/](https://vitess.io/)
- [zenko](https://github.com/scality/Zenko) - Because everyone should be in control of their data. [http://www.zenko.io/](http://www.zenko.io/)

## Tools
- [aglio](https://github.com/danielgtaylor/aglio) - An API Blueprint renderer with theme support that outputs static HTML.
- [ansible](https://github.com/ansible/ansible) - Ansible is a radically simple IT automation platform that makes your applications and systems easier to deploy. Avoid writing scripts or custom code to deploy and update your applications — automate in a language that approaches plain English, using SSH, with no agents to install on remote systems. http://docs.ansible.com/ansible/latest/index.html
- [ark](https://github.com/vmware-tanzu/velero) - Heptio Ark is a utility for managing disaster recovery, specifically for your Kubernetes cluster resources and persistent volumes. Brought to you by Heptio. [https://heptio.cloud.vmware.com//](https://heptio.cloud.vmware.com//)
- [chaostoolkit](https://github.com/chaostoolkit/chaostoolkit/) - An Open API to Chaos Engineering [https://chaostoolkit.org/](https://chaostoolkit.org/)
- [che](https://github.com/eclipse/che) - Eclipse Che: Next-generation Eclipse IDE. Open source workspace server and cloud IDE.
- [client-go](https://github.com/kubernetes/client-go) - Go client for Kubernetes.
- [cloud-native-sandbox](https://github.com/rootsongjc/cloud-native-sandbox) - Cloud Native Sandbox can help you setup a standalone Kubernetes and Istio environment with Docker on you own laptop.
- [cluster-lifecycle-manager](https://github.com/zalando-incubator/cluster-lifecycle-manager) - Cluster Lifecycle Manager (CLM) to provision and update multiple Kubernetes clusters.
- [cost-model](https://github.com/kubecost/cost-model) - Cross-cloud cost allocation models for workloads running on Kubernetes.
- [container-structure-test](https://github.com/GoogleContainerTools/container-structure-test) - Validate the structure of your container images.
- [container-transform](https://github.com/micahhausler/container-transform) - Transforms docker-compose, ECS, and Marathon configurations.
- [crashcart](https://github.com/oracle/crashcart) - CrashCart: sideload binaries into a running container.
- [cri-tools](https://github.com/kubernetes-sigs/cri-tools) - CLI and validation tools for Kubelet Container Runtime Interface (CRI) .
- [devspace](https://github.com/devspace-cloud/devspace) - Cloud Native Software Development with Kubernetes and Docker - simply run "devspace up" in any of your projects and start coding directly on top of Kubernetes (works with minikube, self-hosted and cloud-based clusters) [https://devspace.cloud](https://devspace.cloud/)
- [docker-elk](https://github.com/deviantony/docker-elk) - The ELK stack powered by Docker and Compose.
- [docker-pushrm](https://github.com/christian-korneck/docker-pushrm) - A Docker CLI plugin that that lets you push the README.md file from the current directory to Docker Hub. Also supports Quay and Harbor.
- [dockersh](https://github.com/Yelp/dockersh) - A shell which places users into individual docker containers.
- [dotmesh](https://github.com/dotmesh-io/dotmesh) - dotmesh (dm) is like git for your data volumes (databases, files etc) in Docker and Kubernetes [https://dotmesh.com](https://dotmesh.com/)
- [dragonfly](https://github.com/dragonflyoss/Dragonfly) - Dragonfly is an intelligent P2P based file distribution system.
- [drakov](https://github.com/Aconex/drakov) - Mock Server that implements the API Blueprint specification.
- [eksctl](https://github.com/weaveworks/eksctl) - a CLI for Amazon EKS [https://eksctl.io](https://eksctl.io/)
- [escalator](https://github.com/atlassian/escalator) - Escalator is a batch or job optimized horizontal autoscaler for Kubernetes.
- [fleet](https://github.com/rancher/fleet) - Manage large fleets of Kubernetes clusters.
- [freshpod](https://github.com/googlecloudplatform/freshpod) - Restart Pods on Minikube automatically on image rebuilds.
- [garden](https://github.com/garden-io/garden) - Development orchestrator for Kubernetes, containers and serverless functions. [https://garden.io](https://garden.io/)
- [gardener](https://github.com/gardener/gardener) - Kubernetes API server extension and controller manager providing conformant Kubernetes clusters (a.k.a. (off)shoot clusters) as a service (with day-2 ops) on Alibaba, AWS, Azure, GCP, and OpenStack.
- [go-kubectx](https://github.com/aca/go-kubectx) - 5x-10x faster alternative to kubectx. Uses client-go.
- [istio-pod-network-controller](https://github.com/sabre1041/istio-pod-network-controller) - Controller to manage Istio Pod Network.
- [kail](https://github.com/boz/kail) - Kubernetes log viewer.
- [kind](https://github.com/kubernetes-sigs/kind) - Kubernetes IN Docker - local clusters for testing Kubernetes.
- [kops](https://github.com/kubernetes/kops) - Kubernetes Operations (kops) - Production Grade K8s Installation, Upgrades, and Management.
- [krane](https://github.com/Shopify/krane) - A command-line tool that helps you ship changes to a Kubernetes namespace and understand the result.
- [krustlet](https://github.com/deislabs/krustlet) - Kubernetes Rust Kubelet.
- [ksonnet-lib](https://github.com/ksonnet/ksonnet-lib) - (technical preview) Simplify working with Kubernetes.
- [ksonnet](https://github.com/ksonnet/ksonnet) - A CLI-supported framework that streamlines writing and deployment of Kubernetes configurations to multiple clusters.
- [ksync](https://github.com/ksync/ksync) -  Sync files between your local system and a kubernetes cluster.
- [kt-connect](https://github.com/alibaba/kt-connect) - Manage and Integration with your Kubernetes dev environment more efficient.
- [ktmpl](https://github.com/InQuicker/ktmpl) - Parameterized templates for Kubernetes manifests.
- [kube-capacity](https://github.com/robscott/kube-capacity) - A simple CLI that provides an overview of the resource requests, limits, and utilization in a Kubernetes cluster.
- [kube-downscaler](https://github.com/hjacobs/kube-downscaler) - Scale down Kubernetes deployments after work hours https://hub.docker.com/r/hjacobs/kube-downscaler
- [kube-fledged](https://github.com/senthilrch/kube-fledged) - A kubernetes add-on for creating and managing a cache of container images in a kubernetes cluster.
- [kube-ps1](https://github.com/jonmosco/kube-ps1) - Kubernetes prompt info for bash and zsh.
- [kube-shell](https://github.com/cloudnativelabs/kube-shell) - Kubernetes shell: An integrated shell for working with the Kubernetes CLI.
- [kube-version-converter](https://github.com/fleeto/kube-version-converter) - Convert API Object file into specified version.
- [kubeasz](https://github.com/easzlab/kubeasz) - 使用Ansible脚本安装K8S集群，介绍组件交互原理，方便直接，不受国内网络环境影响.
- [kubeadm-offline-installer](https://github.com/fleeto/kubeadm-offline-installer) - Setup a cluster with kubeadm, without internet connections.
- [kubeadm](https://github.com/kubernetes/kubeadm) - Aggregator for issues filed against kubeadm.
- [kubebox](https://github.com/astefanutti/kubebox) - Terminal console for Kubernetes clusters [https://kube.sh](https://kube.sh/)
- [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) - Kubebuilder - SDK for building Kubernetes APIs using CRDs [https://book.kubebuilder.io/](https://book.kubebuilder.io/)
- [kubecarrier](https://github.com/kubermatic/kubecarrier) - KubeCarrier - Service Management at Scale.
- [kubecdn](https://github.com/ilhaan/kubeCDN) - Self-hosted CDN based on Kubernetes.
- [kubecfg](https://github.com/bitnami/kubecfg) - A tool for managing complex enterprise Kubernetes environments as code.
- [kubectl-doctor](https://github.com/emirozer/kubectl-doctor) - `kubectl` cluster triage plugin for Kubernetes (brew doctor equivalent).
- [kubectl-trace](https://github.com/iovisor/kubectl-trace) - Schedule bpftrace programs on your kubernetes cluster using the kubectl.
- [kubedb](https://github.com/k8sdb/cli) - KubeDB CLI [https://kubedb.com](https://kubedb.com/) to manage kubernetes ready production-grade Databases.
- [kubedirector](https://github.com/bluek8s/kubedirector) - Kubernetes Director (aka KubeDirector) for deploying and managing stateful applications on Kubernetes.
- [kubefwd](https://github.com/txn2/kubefwd) - Bulk port forwarding Kubernetes services for local development. [https://imti.co/kubernetes-port-forwarding/](https://imti.co/kubernetes-port-forwarding/)
- [kubehandler](https://github.com/gojektech/kubehandler) - A framework for writing Kubernetes controllers.
- [kubeiql](https://github.com/yipeeio/kubeiql) - A GraphQL interface for Kubernetes.
- [kuberlr](https://github.com/flavio/kuberlr) - A tool that simplifies the management of multiple versions of kubectl.
- [kubernetes-client](https://github.com/fabric8io/kubernetes-client) - Java client for Kubernetes & OpenShift 3 [http://fabric8.io](http://fabric8.io/)
- [kubernetes-vagrant-centos-cluster](https://github.com/rootsongjc/kubernetes-vagrant-centos-cluster) - Setting up a distributed Kubernetes cluster along with Istio service mesh locally with Vagrant and VirtualBox.
- [kubespray](https://github.com/kubernetes-sigs/kubespray) - Setup a kubernetes cluster also mentioned as kargo.
- [kubesql](https://github.com/xuxinkun/kubesql) - A tool using sql to query the resources of kubernetes, such as pod, node and so on.
- [kubetap](https://github.com/soluble-ai/kubetap) - Kubectl plugin to interactively proxy Kubernetes Services with ease.
- [kubeup](https://github.com/kubeup/archon) - Cluster operation the Kubernetes way.
- [kubie](https://github.com/sbstp/kubie) - A more powerful alternative to kubectx and kubens [https://blog.sbstp.ca/introducing-kubie/](https://blog.sbstp.ca/introducing-kubie/)
- [kuui](https://github.com/viveksinghggits/kuui) - UI that can be used to edit configmaps/secrets of your kubernetes cluster.
- [microconfig](https://github.com/microconfig/microconfig) - Modern and simple way of microservice configuration management.
- [microk8s](https://github.com/ubuntu/microk8s) - A kubernetes cluster in a snap [https://microk8s.io](https://microk8s.io/)
- [minikube](https://github.com/kubernetes/minikube) - Run Kubernetes locally.
- [monday](https://github.com/eko/monday) - A dev tool for microservice developers that run local applications and/or forward some others from Kubernetes or over SSH.
- [okteto](https://github.com/okteto/okteto) - Local development experience for Kubernetes apps.
- [packer](https://github.com/hashicorp/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
- [pangolin](https://github.com/dpeckett/pangolin) - An enhanced Horizontal Pod Autoscaler for Kubernetes.
- [podtnl](https://github.com/narendranathreddythota/podtnl) - A Powerful CLI that makes your pod available to online without exposing a Kubernetes service.
- [portainer](https://github.com/portainer/portainer) - Simple management UI for Docker [http://www.portainer.io/](http://www.portainer.io/)
- [rback](https://github.com/team-soteria/rback) - RBAC in Kubernetes visualizer.
- [reloader](https://github.com/stakater/Reloader) - A Kubernetes controller to watch changes in ConfigMap and Secrets and do rolling upgrades on Pods with their associated Deployment, StatefulSet, DaemonSet and DeploymentConfig.
- [searchlight](https://github.com/searchlight/searchlight) - Alerts for Kubernetes.
- [skopeo](https://github.com/containers/skopeo) - Work with remote images registries - retrieving information, images, signing content.
- [sloop](https://github.com/salesforce/sloop) - Kubernetes History Visualization.
- [sonobuoy](https://github.com/vmware-tanzu/sonobuoy) - Heptio Sonobuoy is a diagnostic tool that makes it easier to understand the state of a Kubernetes cluster by running a set of Kubernetes conformance tests in an accessible and non-destructive manner. Brought to you by Heptio.[https://heptio.cloud.vmware.com//](https://heptio.cloud.vmware.com//)
- [squash](https://github.com/solo-io/squash) - The debugger for microservices.
- [stash](https://github.com/stashed/stash) - Backup your Kubernetes Volumes by <htts://appscode.com>
- [statusbay](https://github.com/similarweb/statusbay) - Kubernetes deployment visibility like a pro.
- [stern](https://github.com/wercker/stern) - Multi pod and container log tailing for Kubernetes.
- [swagger](https://github.com/swagger-api/swagger-ui) - Swagger UI is a collection of HTML, JavaScript, and CSS assets that dynamically generate beautiful documentation from a Swagger-compliant API. [https://swagger.io/](https://swagger.io/)
- [talos](https://github.com/talos-systems/talos) - A modern OS for Kubernetes. [https://www.talos.dev](https://www.talos.dev/)
- [tectonic-installer](https://github.com/coreos/tectonic-installer) - Install a Kubernetes cluster the CoreOS Tectonic Way: HA, self-hosted, RBAC, etcd Operator, and more.
- [telepresence](https://github.com/telepresenceio/telepresence) - Local development against a remote Kubernetes or OpenShift cluster [https://www.telepresence.io/](https://www.telepresence.io/)
- [terminus](https://github.com/godaddy/terminus) - Graceful shutdown and Kubernetes readiness / liveness checks for any Node.js HTTP applications.
- [test-infra](https://github.com/kubernetes/test-infra) - Test infrastructure for the Kubernetes project.
- [tini](https://github.com/krallin/tini) - A tiny but valid `init` for containers.
- [tor-controller](https://github.com/kragniz/tor-controller) - Run Tor onion services on Kubernetes.
- [usernetes](https://github.com/rootless-containers/usernetes) - Kubernetes installable under $HOME, without the root privileges.
- [vagrant](https://github.com/hashicorp/vagrant) - Vagrant is a tool for building and distributing development environments. [https://www.vagrantup.com](https://www.vagrantup.com/)
- [watchtower](https://github.com/containrrr/watchtower) - Automatically update running Docker containers.
- [wksctl](https://github.com/weaveworks/wksctl) - Open Source Weaveworks Kubernetes System.
- [xlskubectl](https://github.com/learnk8s/xlskubectl) - a spreadsheet to control your Kubernetes cluster https://learnk8s.io/real-time-dashboard

## Tracing

- [appdash](https://github.com/sourcegraph/appdash) - Application tracing system for Go, based on Google's Dapper. [https://about.sourcegraph.com](https://about.sourcegraph.com/)
- [jaeger](https://github.com/jaegertracing/jaeger) - Jaeger, a Distributed Tracing System.
- [opencensus](https://github.com/census-instrumentation) - A single distribution of libraries that automatically collect traces and metrics from your app, display them locally, and send them to any backend. - https://opencensus.io
- [opentracing](https://github.com/opentracing) - Consistent, expressive, vendor-neutral APIs for distributed tracing and context propagation.
- [pinpoint](https://github.com/naver/pinpoint) - Pinpoint is an open source APM (Application Performance Management) tool for large-scale distributed systems written in Java.
- [sentry](https://github.com/getsentry/sentry) - Sentry is a cross-platform crash reporting and aggregation platform. [https://sentry.io/welcome/welcome/](https://sentry.io/welcome/welcome/)
- [skywalking](https://github.com/apache/skywalking) - An APM system for tracing, monitoring, diagnosing distributed systems, especially based on microservices, cloud native and container.
- [sofa-tracker](https://github.com/sofastack/sofa-tracer) -  SOFATracer is a component for the distributed system call trace. And through a unified traceId logging the logs of various network calls in the invoking link . These logs can be used for quick discovery of faults, service governance, etc.
- [zipkin](https://github.com/openzipkin/zipkin) - Zipkin is a distributed tracing system [https://zipkin.io/](https://zipkin.io/)

## Tutorials

- [aws-workshop-for-kubernetes](https://github.com/aws-samples/aws-workshop-for-kubernetes) - AWS Workshop for Kubernetes.
- [envoy-steps](https://github.com/datawire/envoy-steps) - Envoy Step by Step.
- [envoy-tutorial](https://github.com/rootsongjc/envoy-tutorial) - Envoy mesh in kubernetes tutorial.
- [istio-handbook](https://github.com/servicemesher/istio-handbook) - Istio Service Mesh Advanced Practical Istio服务网格进阶实战 https://www.servicemesher.com/istio-handbook/
- [istio-index-conf2018](https://github.com/todkap/istio-index-conf2018) -  Istio is not just for Microservices: Secure your Kubernetes services using Istio Service Mesh <https://developer.ibm.com/>
- [istio-ingress-tutorial](https://github.com/kelseyhightower/istio-ingress-tutorial) - How to run the Istio Ingress Controller on Kubernetes.
- [istio-service-mesh-workshop](https://github.com/layer5io/istio-service-mesh-workshop) - Using Istio Workshop [https://layer5.io](https://layer5.io/)
- [istio-tutorial](https://github.com/redhat-developer-demos/istio-tutorial) - Istio Tutorial for Java Microservices.
- [istio101](https://github.com/IBM/istio101) - Istio 101 workshop from IBM.
- [ks](https://github.com/red-gate/ks) - A series of Kubernetes walk-throughs.
- [kube-ladder](https://github.com/caicloud/kube-ladder) - Learning Kubernetes, The Chinese Taoist Way.
- [kubeadm-workshop](https://github.com/luxas/kubeadm-workshop) - Showcasing a bare-metal multi-platform kubeadm setup with persistent storage and monitoring.
- [kubernetes-handbook](https://github.com/rootsongjc/kubernetes-handbook) - Kubernetes中文指南/云原生应用架构实践手册 - https://jimmysong.io/kubernetes-handbook/
- [kubernetes-java-simple](https://github.com/arun-gupta/kubernetes-java-sample) - Kubernetes Hands-on Workshop for Java Developers.
- [kubernetes-on-aws](https://github.com/zalando-incubator/kubernetes-on-aws) - Deploying Kubernetes on AWS with CloudFormation and Container Linux [https://kubernetes-on-aws.readthedocs.io/en/latest/en/latest/](https://kubernetes-on-aws.readthedocs.io/en/latest/en/latest/)
- [kubernetes-security-best-practice](https://github.com/freach/kubernetes-security-best-practice) - Kubernetes Security - Best Practice Guide.
- [kubernetes-the-hard-way](https://github.com/kelseyhightower/kubernetes-the-hard-way) - Bootstrap Kubernetes the hard way on Google Cloud Platform. No scripts.
- [kubicorn](https://github.com/kris-nova/kubicorn-fork) - Create, manage, snapshot, and scale Kubernetes infrastructure in the public cloud.
- [mosn-tutorial](https://github.com/mosn/mosn-tutorial) - Tutorial for MOSN and Istio Service Mesh - https://www.katacoda.com/mosn
- [rpi-handbook](https://github.com/rootsongjc/rpi-handbook) - Raspberry Pi Handbook/树莓派实践手册 - https://jimmysong.io/rpi-handbook/

## Community

- [Cloud Native Community](https://cloudnative.to)
- [Cloud Native Computing Foundation](https://www.cncf.io/)
- [ServiceMesher](https://www.servicemesher.com)

## Contribute

This website is hosted on GitHub Pages with [rootsongjc/awesome-cloud-native](https://github.com/rootsongjc/awesome-cloud-native) repository.

Please take a quick gander at the **[contribution guidelines](https://github.com/rootsongjc/awesome-cloud-native/blob/master/CONTRIBUTING.md)** first. Thanks to all **[contributors](https://github.com/rootsongjc/awesome-cloud-native/graphs/contributors)**, you rock🤟!

## License

[![CC0](https://camo.githubusercontent.com/60561947585c982aee67ed3e3b25388184cc0aa3/687474703a2f2f6d6972726f72732e6372656174697665636f6d6d6f6e732e6f72672f70726573736b69742f627574746f6e732f38387833312f7376672f63632d7a65726f2e737667)](https://creativecommons.org/publicdomain/zero/1.0/)

To the extent possible under law, the authors have waived all copyright and related or neighbouring rights to this work.
