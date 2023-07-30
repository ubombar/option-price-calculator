# Option Price Calculator
<img src="https://github.com/ubombar/option-price-calculator/blob/main/assets/images/robo.png?raw=true" width="150">

The Option Price Calculator is a simple yet powerful tool designed to calculate option prices for various financial instruments, helping traders and investors make informed decisions. This open-source project provides a user-friendly interface for inputting essential parameters and obtaining accurate option prices using popular pricing models.

## Table of Contents

- [Option Price Calculator](#option-price-calculator)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Getting Started](#getting-started)
    - [Docker Image](#docker-image)
    - [Running the Image](#running-the-image)
  - [Supported Pricing Models](#supported-pricing-models)
  - [Contributing](#contributing)

## Introduction

The financial markets are dynamic and ever-changing, and options play a significant role in managing risk and generating potential profits. However, accurately pricing options can be complex, involving various parameters and mathematical models. The Option Price Calculator simplifies this process, allowing users to estimate option prices quickly and efficiently.

This repository contains the source code for the Option Price Calculator. Feel free to explore the codebase, contribute, and use the calculator for your trading and investment needs.

## Getting Started

### Docker Image

The application uses gRPC for receiving requests. It is also containerized so the image can be accessed from DockerHub `ubombar/ubombar/option-price-calculator`.

### Running the Image

1. Run the following `docker` command.

```bash
docker run -itd -p 8080:8080 ubombar/option-price-calculator:latest
```

## Supported Pricing Models

The Option Price Calculator currently supports the following pricing models:

1. Black-Scholes Model

You can easily add more pricing models by following the guidelines provided in the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## Contributing

We welcome contributions to the Option Price Calculator! Whether you want to report a bug, suggest a new feature, or add support for a new pricing model, please refer to our [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to contribute.

