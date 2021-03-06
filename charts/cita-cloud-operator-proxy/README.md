# cita-cloud-operator-proxy

![Version: 0.0.3](https://img.shields.io/badge/Version-0.0.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 6.4.0](https://img.shields.io/badge/AppVersion-6.4.0-informational?style=flat-square)

A Helm chart for Cita-Cloud Operator Proxy

**Homepage:** <https://github.com/cita-cloud/operator-proxy>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Rivtower Technologies | contact@rivtower.com |  |

## Requirements

Kubernetes: `>=1.18.0-0`

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| global.registry | string | `"docker.io"` | You can specify the image registry. |
| global.repository | string | `"citacloud"` | You can specify the image repository. |
| image.pullPolicy | string | `"IfNotPresent"` | You can specify the image pull policy. |
| image.tag | string | `"v0.0.3"` | You can specify the image tag. |
| replicas | int | `1` | You can specify the replica count. |
| serviceAccountName | string | `"cita-cloud-operator-proxy-sa"` | You can specify the service account. |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.6.0](https://github.com/norwoodj/helm-docs/releases/v1.6.0)
