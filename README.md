# [![repo logo][]][Kicksware url]

<p align="center">
	<a href="https://www.kicksware.com">
		<img src="https://img.shields.io/website?label=Visit%20website&down_message=unavailable&up_color=teal&up_message=kicksware.com%20%7C%20online&url=https%3A%2F%2Fhealth.kicksware.com/ready">
	</a>
</p>

[![golang badge]](https://golang.org)&nbsp;
[![kubernetes badge]](https://kubernetes.io)&nbsp;
[![architecture badge]][microservice article]&nbsp;
[![maintainability badge]][maintainability source]&nbsp;
[![license badge]](https://www.gnu.org/licenses/agpl-3.0)

## Overview

_**Kicksware API**_ provides both RESTful and gRPC interfaces to grand access, control, and management of the Kicksware sneaker resale platform while it's designed as a microservice-based, cloud-native, composite solution that collectively represents a set of a lightweight, stateless web-services, each written in Golang.

## Table of contents

* [Overview](#overview)
* [Table of contents](#table-of-contents)
* [Endpoints](#endpoints)
	* [REST](#rest)
	* [gRPC](#grpc)
* [Authentication](#authentication)
* [Architecture](#architecture)
* [Internal design](#internal-design)
* [Requirements](#requirements)
* [Deployment](#deployment)
* [Wrap Up](#wrap-up)
* [License](#license)

## Endpoints

### REST

All API's endpoints are divided into 10 base resources which mostly correspond to their source microservices as follows:

| Service            | Server URL        | Endpoint base resources         |
|--------------------|-------------------|---------------------------------|
| users-service      | api.kicksware.com | /users, /auth, /mail, /interact |
| references-service | api.kicksware.com | /references/sneakers            |
| products-service   | api.kicksware.com | /products/sneakers              |
| search-service     | api.kicksware.com | /search                         |
| orders-service     | api.kicksware.com | /orders                         |
| cdn-service        | cdn.kicksware.com | /                               |
| \*                 | \*.kicksware.com  | /health                         |

RESTful API uses `api.kicksware.com` subdomain as it's base server URL.

An exception is the Content Deliver Network (CDN) service,
which as it's own subdomain `cdn.kicksware.com`.

API is accessible from both `:443` port via HTTPS and `:80` via plain HTTP.

The full API specification is available on [**Swagger**][swagger] and [**Readme.io**][readme.io].

### gRPC

As it is cloud-native and microservice-based application, Kicksware also provides RPC API using [gRPC framework][grpc] and [Protocol Buffers][protobuf] language.

Like many RPC systems, gRPC is based on the concept of defining a service in terms of functions (methods) that can be called remotely. This approach is especially useful for distributed, loose-coupled systems, as it provides a mechanism to write API specification ones as a set of `.proto` files from witch API implementation can be generated on any language.

All `.proto` specification files are available [here][proto files].

gRPC API logical division is also based on source microservice entries as follows:

| Service            | Server URL        | Proto endpoints                                                                    |
|--------------------|-------------------|------------------------------------------------------------------------------------|
| users-service      | rpc.kicksware.com | /proto.UserService, /proto.AuthService, /proto.MailService, /proto.InteractService |
| references-service | rpc.kicksware.com | /proto.ReferenceService                                                            |
| products-service   | rpc.kicksware.com | /proto.ProductService                                                              |
| search-service     | rpc.kicksware.com | /proto.SearchReferenceService, /proto.SearchProductService                         |
| orders-service     | rpc.kicksware.com | /proto.OrdersService                                                               |
| cdn-service        | cdn.kicksware.com | /proto.CDNService                                                                  |

## Authentication

[JSON Web Token (JWT)][jwt auth] is used to authenticate and authorize all REST requests.

To access gRPC based API, both secure TLS connection and [token interceptors][grpc interceptor] are required.

## Architecture

As it was mentioned earlier Kicksware API design is based completely on _[**microservice architecture**][microservice article] pattern_.

Like any other, this approach has its own pros and cons. The common right way to decide whether a specific architecture is suits a specific system or not, is by using what's calls [architecture trade-off analysis method (ATAM)][atam wiki]. Basically, this method compares software architectures relative to quality attribute goals and helps expose architectural risks that potentially inhibit the achievement of an organization’s business goals.

When evaluating a microservice architecture style, it is important to understand that this approach is generally harder to implement, maintain, and test and requires more staff, money, and resources, but as a trade-off microservices provides one the most effective method of Horizontal scaling. And as a bonus you'll get higher flexibility as each new web-service can be written in any language and using any technology as long as it has some kind of communication mechanism (API).

As for this particular project the main goal was to reverse engineer evaluating process and build the system that would fit MSA the best way possible.

## Internal design

While microservice architecture divides the entire system into small, independent services, it's important for a code to stay organized and clean even in scale of one microservice.

For this purpose Kicksware design adopts _[uncle Bob's][uncle Bob] [**Clean Architecture**][clean architecture]_ - another grate architecture pattern that separates the design elements into ring levels and its basic rule is that code dependencies can only come from the outer levels inward, so the further in you go, the higher level the software becomes. Simply put, the code on the inner layers can have no knowledge of the code on the outer layers.

You may have seen diagrams like the following, but this one, in particular, is Kicksware custom API microservice Clean Architecture representational chart:

![Clean architecture chart][clean architecture chart]

## Requirements

API microservice registry should be deployed after [Gateway][gateway repo] and [Tool Stack][tool-stack repo] projects. The reason for this is DB and Elasticsearch dependencies from [Tool Stack][tool-stack repo] project and Traefik Proxy from [Gateway][gateway repo] project to route outer traffic to services.

## Deployment

Kicksware project can be deployed using the following methods:

1. **Docker Compose file**

   This method requires single dedicated server with both installed [`docker`][docker-compose] and [`docker-compose`][docker-compose] utilities.

   Compose [configuration file][compose config] can be found in the root of the project. This file already contains setting for reverse proxy routing and load balancing.

   Gitlab CI deployment pipeline [configuration file][ci compose config] for compose method can be found in `.gitlab` directory.

2. **Kubernetes Helm charts**

   Deployment to Kubernetes cluster is the default and desired way.

   For more flexible and easier deployment [Helm package manager][helm] is used. It provides a simple, yet elegant way to write pre-configured, reusable Kubernetes resource configuration using YAML and Go Templates (or Lua scripts). Helm packages are called `charts`.

   Each microservice has it's own chart in the root of it's directory:

   | Service            | Helm chart                                                       |
   |--------------------|------------------------------------------------------------------|
   | users-service      | [~/user-service/users-chart][users-service chart]                |
   | references-service | [~/reference-service/references-chart][references-service chart] |
   | products-service   | [~/product-service/products-chart][products-service chart]       |
   | search-service     | [~/search-service/search-chart][search-service chart]            |
   | orders-service     | [~/order-service/orders-chart][orders-service chart]             |
   | cdn-service        | [~/cdn-service/cdn-chart][cdn-service chart]                     |

   Helm chart configuration already contains configuration of [Traefik IngressRoute][ingress route] [Custom Resource Definition (CRD)][k8s crd] for reverse proxy routing and load balancing.

   Gitlab CI deployment pipeline [configuration file][ci k8s config] for K8s method can be found in the root of the project.

## Wrap Up

**Kicksware API** is the accumulation of Kicksware's business logic in the form of distributed, atomically granulated, stateless web-services, where each is responsible only for its entities, use cases, API endpoints.

It exposes access to the data and its functionality as a set of both RESTfull endpoints and gRPC remote procedures.

There are two ways to perform Kicksware API deployment. To ensure top performance with longer uptime and lesser latency while having effective and automated control over scalability, Kicksware utilizes Kubernetes cluster with minimum of 3 basic spec nodes and 2 more for potential cluster autoscaling.

An alternative and in fact, a cheaper and easier way to achieve sufficient results is for deployment to be performed with just one VPS server and docker-compose utility. However, such temporary savings now may be offset by the costs of vertical scaling when demand rises in the future.

## More
See [other Kicksware project repositories][kicksware-main-repo].

## License

Licensed under the [GNU AGPL-3.0][license file].

[repo logo]: https://raw.githubusercontent.com/timoth-y/kicksware-api/master/assets/repo-logo.png
[kicksware url]: https://www.kicksware.com

[Website badge]: https://img.shields.io/website?label=Visit%20website&down_message=unavailable&up_color=teal&up_message=kicksware.com%20%7C%20online&url=https%3A%2F%2Fkicksware.com
[golang badge]: https://img.shields.io/badge/Code-Golang-informational?style=flat&logo=go&logoColor=white&color=6AD7E5
[commit activity badge]: https://img.shields.io/github/commit-activity/m/timoth-y/kicksware-api?label=Commit%20activity&color=teal
[repo commit activity]: https://github.com/timoth-y/kicksware-api/graphs/commit-activity
[lines counter]: https://img.shields.io/tokei/lines/github/timoth-y/kicksware-api?color=teal&label=Lines
[license badge]: https://img.shields.io/badge/License-AGPL%20v3-blue.svg?color=teal
[architecture badge]: https://img.shields.io/badge/Architecture-Microservices-informational?style=flat&logo=opslevel&logoColor=white&color=teal
[kubernetes badge]: https://img.shields.io/badge/DevOps-Kubernetes-informational?style=flat&logo=kubernetes&logoColor=white&color=316DE6
[gitlab badge]: https://img.shields.io/badge/CI-Gitlab_CE-informational?style=flat&logo=gitlab&logoColor=white&color=FCA326
[api pipeline]: https://ci.kicksware.com/kicksware/api/badges/master/pipeline.svg?key_text=API%20|%20pipeline&key_width=85
[maintainability badge]: https://api.codeclimate.com/v1/badges/367c3a861b61cc78d24c/maintainability
[maintainability source]: https://codeclimate.com/github/timoth-y/kicksware-api/maintainability

[microservice article]: https://martinfowler.com/articles/microservices.html

[jwt auth]: https://jwt.io/introduction
[grpc interceptor]: https://github.com/grpc/grpc-go/tree/master/examples/features/interceptor

[swagger]: https://app.swaggerhub.com/apis/timoth-y/kicksware-api/1.0.0
[readme.io]: https://kicksware-api.readme.io/reference
[grpc]: https://grpc.io
[protobuf]: https://developers.google.com/protocol-buffers
[proto files]: https://github.com/timoth-y/kicksware-api/tree/master/service-protos

[atam wiki]: https://en.wikipedia.org/wiki/Architecture_tradeoff_analysis_method
[uncle Bob]: http://cleancoder.com/products
[clean architecture]: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
[clean architecture chart]: https://raw.githubusercontent.com/timoth-y/kicksware-api/master/assets/clean-archtecture.png

[gateway repo]: https://github.com/timoth-y/kicksware-gateway
[tool-stack repo]: https://github.com/timoth-y/kicksware-tool-stack

[docker-desktop]: https://docs.docker.com/desktop/
[docker-compose]: https://docs.docker.com/compose/
[compose config]: https://github.com/timoth-y/kicksware-api/blob/master/docker-compose.yml
[ci compose config]: https://github.com/timoth-y/kicksware-api/blob/master/.gitlab/.gitlab-ci.compose.yml
[ci k8s config]: https://github.com/timoth-y/kicksware-api/blob/master/.gitlab-ci.yml
[k8s crd]: https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/
[ingress route]: https://docs.traefik.io/routing/providers/kubernetes-crd/

[helm]: https://helm.sh/
[users-service chart]: https://github.com/timoth-y/kicksware-api/tree/master/user-service/users-chart
[references-service chart]: https://github.com/timoth-y/kicksware-api/tree/master/reference-service/references-chart
[products-service chart]: https://github.com/timoth-y/kicksware-api/tree/master/product-service/products-chart
[search-service chart]: https://github.com/timoth-y/kicksware-api/tree/master/search-service/search-chart
[orders-service chart]: https://github.com/timoth-y/kicksware-api/tree/master/order-service/orders-chart
[cdn-service chart]: https://github.com/timoth-y/kicksware-api/tree/master/cdn-service/cdn-chart


[kicksware-main-repo]: https://github.com/timoth-y/kicksware-platform#components

[license file]: https://github.com/timoth-y/kicksware-api/blob/master/LICENSE
