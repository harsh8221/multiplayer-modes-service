# Multiplayer Modes Service

A Go-based microservice for tracking and reporting popular multiplayer game modes.

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [API Documentation](#api-documentation)
6. [Development](#development)
7. [Testing](#testing)
8. [Deployment](#deployment)
9. [Contributing](#contributing)

## Overview

This project implements a multiplayer modes service that allows tracking and reporting popular game modes. It uses gRPC for communication and MongoDB for data storage. The service provides two main functionalities:

1. Reporting mode playing: Increment player counts for specific game modes.
2. Getting popular modes: Retrieve a list of popular game modes sorted by player count.

## Features

- Real-time updating of mode popularity
- Efficient caching mechanism for frequently accessed data
- Scalable architecture using gRPC and MongoDB
- Docker support for easy deployment

## Installation

To install the necessary dependencies:
bash make build


This command will build the binary for the service.

## Usage

### Running the Service

To start the service locally:

bash make run


### Using Docker

To build and run the service using Docker:

bash docker-compose up --build


This will start both the service and a MongoDB instance.

### API Endpoints

The service exposes two main endpoints:

1. `ReportModePlaying`: Increments the player count for a specific mode.
2. `GetPopularModes`: Retrieves a list of popular modes sorted by player count.

See the [API Documentation](#api-documentation) section for more details.

## API Documentation

The service uses gRPC protocol. You can find the protobuf definitions in the `proto/service.proto` file.

To generate client code for your preferred language, use the following command:

bash protoc --go_out=. --go_opt=paths=source_relative
--go-grpc_out=. --go-grpc_opt=paths=source_relative
proto/service.proto


## Development

To start developing:

1. Clone the repository:
bash git clone https://github.com/your-repo/multiplayer-modes-service.git cd multiplayer-modes-service


2. Set up your development environment:
   - Install Go (version 1.23 or later)
   - Install Docker and Docker Compose

3. Run tests:
bash make test


4. Start the service locally:
bash make run


## Testing

The project includes unit tests for both business logic and storage layers. To run all tests:

bash make test


## Deployment

The service can be deployed using Docker. Here's a basic deployment process:

1. Build the Docker image:
bash docker build -t multiplayer-modes-app .


2. Run the container:
bash docker run -p 8080:8080 --env-file .env multiplayer-modes-app
