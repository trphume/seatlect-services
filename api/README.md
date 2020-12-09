# API

- [API](#api)
  - [**Overview**](#overview)
  - [**HTTP**](#http)
  - [**gRPC**](#grpc)
    - [`user`](#user)
    - [`token`](#token)
    - [`business`](#business)

## **Overview**

This repository contains files related to the definition, configuration and generated code of the Seatlect's backend services. The design principles used are based on Google's general guide in API design which can be found [here](https://cloud.google.com/apis/design).

In Seatlect there are two main API platform, HTTP and gRPC. The HTTP platform contains API used by our *business* users, it follows the RESTful design principles and is mostly CRUD based. Our gRPC API platform is used by Seatlect's mobile client for our standard users, the design is task oriented.

## **HTTP**

This section describes our HTTP API platform.

## **gRPC**

This section describes our gRPC API platform. The protocol buffer definition are located in the `/protobuf` directory. For ease of use, the makefile provides commands to easily generate the files. The generated files will be found (starting from the root directory) at `internal/genproto`.

### `user`

- **SignIn** - This endpoint attempts to authenticate the user with the given credentials. If successful, it returns token (refresh and jwt) and information on the user
- **SignUp** - This endpoints create a new user. If successful, it returns token (refresh and jwt) and information on the user

### `token`

- **FetchJWT** - This endpoints uses a refresh token to fetch a new jwt token

### `business`

- **ListBusiness** - This endpoint list (and sort) the business given some sort of search parameters, view the proto file for more information
