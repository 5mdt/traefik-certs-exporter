# Traefik certs exporter

[![Run revive linters](https://github.com/nett00n/traefik-certs-exporter/actions/workflows/linter_revive.yml/badge.svg?branch=main)](https://github.com/nett00n/traefik-certs-exporter/actions/workflows/linter_revive.yml)
[![Build linux binaries](https://github.com/nett00n/traefik-certs-exporter/actions/workflows/linux_build.yml/badge.svg?branch=main)](https://github.com/nett00n/traefik-certs-exporter/actions/workflows/linux_build.yml)
[![Push to Dockerhub](https://github.com/nett00n/traefik-certs-exporter/actions/workflows/docker_push.yml/badge.svg)](https://github.com/nett00n/traefik-certs-exporter/actions/workflows/docker_push.yml)

Script opens `acme.json` file used by traefik and export it's content to `output` folder as `*.cer` and `*.key` files.

There is template `acme.json` file in `input` directory for demonstration purposes. It is added in `.gitignore` - feel free to modify or delete it.

## Usage

Run app

```bash
go run traefik-certs-exporter.go -acmejson=/foo/acme.json -output=/bar/certs/
```

### defaults
- `acmejson`=`./input/acme.json`
- `output`=`./output`

or just use compiled one from Dockerhub

```bash
docker run -v ${PWD}/input:/input -v ${PWD}/output:/output: docker.io/nett00n/traefik-certs-exporter:1.0.0
```

## Links

- [Gihtub Project](https://github.com/nett00n/traefik-certs-exporter)
- [DockerHub Page](https://hub.docker.com/r/nett00n/traefik-certs-exporter)

## License

All code in this repository is licensed under the terms of the MIT License. For further information please refer to the LICENSE file.

## Authors

- Vladimir Budylnikov aka [@nett00n](https://github.com/nett00n)

---
2023, Yerevan
