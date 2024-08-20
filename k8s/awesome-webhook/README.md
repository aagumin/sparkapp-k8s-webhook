# awesome-webhook

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.1.0](https://img.shields.io/badge/AppVersion-0.1.0-informational?style=flat-square)

A Helm chart for Kubernetes

**Homepage:** <https://github.com/aagumin/sparkapp-k8s-webhook>

## Source Code

* <https://github.com/aagumin/blob/master/k8s/awesome-webhook>

## Values

| Key                                                 | Type   | Default                                             | Description |
|-----------------------------------------------------|--------|-----------------------------------------------------|-------------|
| affinity                                            | object | `{}`                                                |             |
| appConfig.spark.configMapName                       | string | `"asmwebhook-configmap"`                            |             |
| appConfig.spark.configMapValue                      | string | ``                                                  |             |
| autoscaling.enabled                                 | bool   | `false`                                             |             |
| autoscaling.maxReplicas                             | int    | `100`                                               |             |
| autoscaling.minReplicas                             | int    | `1`                                                 |             |
| autoscaling.targetCPUUtilizationPercentage          | int    | `80`                                                |             |
| autoscaling.targetMemoryUtilizationPercentage       | int    | `80`                                                |             |
| containerEnv[0].name                                | string | `"LOG_LEVEL"`                                       |             |
| containerEnv[0].value                               | string | `"INFO"`                                            |             |
| containerEnv[1].name                                | string | `"SPARK_CFG_PATH"`                                  |             |
| containerEnv[1].value                               | string | `"/etc/configs/spark.yaml"`                         |             |
| fullnameOverride                                    | string | `""`                                                |             |
| image.pullPolicy                                    | string | `"IfNotPresent"`                                    |             |
| image.repository                                    | string | `"awesome-webhook"`                                 |             |
| image.tag                                           | string | `"0.1.0"`                                           |             |
| imagePullSecrets                                    | list   | `[]`                                                |             |
| livenessProbePath                                   | string | `"/health"`                                         |             |
| mutate.apiGroups[0]                                 | string | `"*"`                                               |             |
| mutate.apiVersions[0]                               | string | `"*"`                                               |             |
| mutate.mutateUrl                                    | string | `"/mutate"`                                         |             |
| mutate.objectSelector.matchExpressions[0].key       | string | `"deploymentApiType"`                               |             |
| mutate.objectSelector.matchExpressions[0].operator  | string | `"NotIn"`                                           |             |
| mutate.objectSelector.matchExpressions[0].values[0] | string | `"nrt_job"`                                         |             |
| mutate.objectSelector.matchExpressions[1].key       | string | `"example.webhook.status"`                          |             |
| mutate.objectSelector.matchExpressions[1].operator  | string | `"NotIn"`                                           |             |
| mutate.objectSelector.matchExpressions[1].values[0] | string | `"disable"`                                         |             |
| mutate.operations[0]                                | string | `"CREATE"`                                          |             |
| mutate.resources[0]                                 | string | `"SparkApplication"`                                |             |
| mutate.resources[1]                                 | string | `"SparkApplications"`                               |             |
| mutate.resources[2]                                 | string | `"sparkapplications"`                               |             |
| mutate.resources[3]                                 | string | `"sparkapplication"`                                |             |
| mutate.resources[4]                                 | string | `"sparkapp"`                                        |             |
| mutate.resources[5]                                 | string | `"sparkapplications.sparkoperator.k8s.io"`          |             |
| mutate.resources[6]                                 | string | `"scheduledsparkapplications.sparkoperator.k8s.io"` |             |
| mutate.scope                                        | string | `"Namespaced"`                                      |             |
| mutate.sideEffects                                  | string | `"None"`                                            |             |
| mutate.timeoutSeconds                               | int    | `20`                                                |             |
| nameOverride                                        | string | `""`                                                |             |
| nodeSelector                                        | object | `{}`                                                |             |
| podAnnotations                                      | object | `{}`                                                |             |
| podSecurityContext                                  | object | `{}`                                                |             |
| readinessProbePath                                  | string | `"/health"`                                         |             |
| replicaCount                                        | int    | `1`                                                 |             |
| resources.limits.cpu                                | string | `"2"`                                               |             |
| resources.limits.memory                             | string | `"4Gi"`                                             |             |
| resources.requests.cpu                              | string | `"2"`                                               |             |
| resources.requests.memory                           | string | `"2Gi"`                                             |             |
| securityContext.allowPrivilegeEscalation            | bool   | `false`                                             |             |
| securityContext.privileged                          | bool   | `false`                                             |             |
| securityContext.readOnlyRootFilesystem              | bool   | `true`                                              |             |
| securityContext.runAsNonRoot                        | bool   | `true`                                              |             |
| securityContext.runAsUser                           | int    | `10001`                                             |             |
| service.port                                        | string | `"9443"`                                            |             |
| service.protocol                                    | string | `"https"`                                           |             |
| service.type                                        | string | `"ClusterIP"`                                       |             |
| tolerations                                         | list   | `[]`                                                |             |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
