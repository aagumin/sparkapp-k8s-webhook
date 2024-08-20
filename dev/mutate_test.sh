curl -X POST --location "https://localhost:443/mutate" \
    -H "Content-Type: application/json" \
    --cacert localhost.crt \
    -d '{
          "kind": "AdmissionReview",
          "apiVersion": "admission.k8s.io/v1",
          "request": {
            "uid": "d8aaf0a1-db2f-4b7d-b5ds-0e627ecx1229",
            "kind": {
              "group": "sparkoperator.k8s.io",
              "version": "v1beta2",
              "kind": "SparkApplication"
            },
            "resource": {
              "group": "sparkoperator.k8s.io",
              "version": "v1beta2",
              "resource": "sparkapplications"
            },
            "requestKind": {
              "group": "sparkoperator.k8s.io",
              "version": "v1beta2",
              "kind": "SparkApplication"
            },
            "requestResource": {
              "group": "sparkoperator.k8s.io",
              "version": "v1beta2",
              "resource": "sparkapplications"
            },
            "name": "spark-pi-test-hook-21",
            "namespace": "seldon",
            "operation": "CREATE",
            "userInfo": {
              "username": "12b120026de1facc00264f6893d",
              "groups": [
                "0f2762d3dsf00207c49c39",
                "system:authenticated"
              ]
            },
            "object": {
              "apiVersion": "sparkoperator.k8s.io/v1beta2",
              "kind": "SparkApplication",
              "metadata": {
                "annotations": {
                  "example.spark.app": "ml",
                  "example.webhook.jobType": "batch",
                  "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"sparkoperator.k8s.io/v1beta2\",\"kind\":\"SparkApplication\",\"metadata\":{\"annotations\":{\"example.spark.app\":\"ml\",\"example.webhook.jobType\":\"batch\"},\"name\":\"spark-pi-test-hook-21\",\"namespace\":\"seldon\"},\"spec\":{\"deps\":{\"packages\":[\"org.apache.hadoop:hadoop-aws:3.2.4\"]},\"driver\":{\"coreLimit\":\"1200m\",\"cores\":1,\"env\":[{\"name\":\"AWS_ACCESS_KEY_ID\",\"value\":\"xxxx\"},{\"name\":\"AWS_SECRET_ACCESS_KEY\",\"value\":\"xxx\"},{\"name\":\"KAFKA_USER\",\"value\":\"example-sasl\"},{\"name\":\"KAFKA_PASSWORD\",\"value\":\"ccccccccc\"}],\"javaOptions\":\"-Divy.cache.dir=/tmp -Divy.home=/tmp\",\"labels\":{\"version\":\"3.1.1\"},\"memory\":\"512m\",\"serviceAccount\":\"spark\",\"volumeMounts\":[{\"mountPath\":\"/tmp\",\"name\":\"test-volume\"}]},\"executor\":{\"cores\":1,\"env\":[{\"name\":\"AWS_ACCESS_KEY_ID\",\"value\":\"xxxxx\"},{\"name\":\"AWS_SECRET_ACCESS_KEY\",\"value\":\"xxxxx\"},{\"name\":\"KAFKA_USER\",\"value\":\"example-sasl\"},{\"name\":\"KAFKA_PASSWORD\",\"value\":\"xxxxxx\"}],\"instances\":1,\"labels\":{\"version\":\"3.1.1\"},\"memory\":\"512m\",\"volumeMounts\":[{\"mountPath\":\"/tmp\",\"name\":\"test-volume\"}]},\"image\":\"spark-3.2.4-0.0.3a\",\"imagePullPolicy\":\"Always\",\"mainApplicationFile\":\"local:///opt/spark/examples/jars/spark-examples_2.12-3.2.4.jar\",\"mainClass\":\"org.apache.spark.examples.SparkPi\",\"mode\":\"cluster\",\"restartPolicy\":{\"type\":\"Never\"},\"sparkConf\":{\"javaOptions\":\"-Divy.cache.dir=/tmp -Divy.home=/tmp\",\"spark.serializer\":\"org.apache.spark.serializer.KryoSerializer\"},\"sparkVersion\":\"3.2.4\",\"type\":\"Scala\",\"volumes\":[{\"hostPath\":{\"path\":\"/tmp\",\"type\":\"Directory\"},\"name\":\"test-volume\"}]}}\n"
                },
                "creationTimestamp": null,
                "managedFields": [
                  {
                    "apiVersion": "sparkoperator.k8s.io/v1beta2",
                    "fieldsType": "FieldsV1",
                    "fieldsV1": {
                      "f:metadata": {
                        "f:annotations": {
                          ".": {},
                          "f:example.spark.app": {},
                          "f:example.webhook.jobType": {},
                          "f:kubectl.kubernetes.io/last-applied-configuration": {}
                        }
                      },
                      "f:spec": {
                        ".": {},
                        "f:deps": {
                          ".": {},
                          "f:packages": {}
                        },
                        "f:driver": {
                          ".": {},
                          "f:coreLimit": {},
                          "f:cores": {},
                          "f:env": {},
                          "f:javaOptions": {},
                          "f:labels": {
                            ".": {},
                            "f:version": {}
                          },
                          "f:memory": {},
                          "f:serviceAccount": {},
                          "f:volumeMounts": {}
                        },
                        "f:executor": {
                          ".": {},
                          "f:cores": {},
                          "f:env": {},
                          "f:instances": {},
                          "f:labels": {
                            ".": {},
                            "f:version": {}
                          },
                          "f:memory": {},
                          "f:volumeMounts": {}
                        },
                        "f:image": {},
                        "f:imagePullPolicy": {},
                        "f:mainApplicationFile": {},
                        "f:mainClass": {},
                        "f:mode": {},
                        "f:restartPolicy": {
                          ".": {},
                          "f:type": {}
                        },
                        "f:sparkConf": {
                          ".": {},
                          "f:javaOptions": {},
                          "f:spark.serializer": {}
                        },
                        "f:sparkVersion": {},
                        "f:type": {},
                        "f:volumes": {}
                      }
                    },
                    "manager": "kubectl-client-side-apply",
                    "operation": "Update",
                    "time": "2024-02-16T12:41:21Z"
                  }
                ],
                "name": "spark-pi-test-hook-21",
                "namespace": "seldon"
              },
              "spec": {
                "deps": {
                  "packages": [
                    "org.apache.hadoop:hadoop-aws:3.2.4"
                  ]
                },
                "driver": {
                  "coreLimit": "1200m",
                  "cores": 1,
                  "env": [
                    {
                      "name": "AWS_ACCESS_KEY_ID",
                      "value": "xxxx"
                    },
                    {
                      "name": "AWS_SECRET_ACCESS_KEY",
                      "value": "xxxx"
                    },
                    {
                      "name": "KAFKA_USER",
                      "value": "example-sasl"
                    },
                    {
                      "name": "KAFKA_PASSWORD",
                      "value": "xxxxxx"
                    }
                  ],
                  "javaOptions": "-Divy.cache.dir=/tmp -Divy.home=/tmp",
                  "labels": {
                    "version": "3.1.1"
                  },
                  "memory": "512m",
                  "serviceAccount": "spark",
                  "volumeMounts": [
                    {
                      "mountPath": "/tmp",
                      "name": "test-volume"
                    }
                  ]
                },
                "executor": {
                  "cores": 1,
                  "env": [
                    {
                      "name": "AWS_ACCESS_KEY_ID",
                      "value": "xxxxxx"
                    },
                    {
                      "name": "AWS_SECRET_ACCESS_KEY",
                      "value": "xxxxxx"
                    },
                    {
                      "name": "KAFKA_USER",
                      "value": "example-sasl"
                    },
                    {
                      "name": "KAFKA_PASSWORD",
                      "value": "xxxxxx"
                    }
                  ],
                  "instances": 1,
                  "labels": {
                    "version": "3.1.1"
                  },
                  "memory": "512m",
                  "volumeMounts": [
                    {
                      "mountPath": "/tmp",
                      "name": "test-volume"
                    }
                  ]
                },
                "image": "spark-3.2.4-0.0.3a",
                "imagePullPolicy": "Always",
                "mainApplicationFile": "local:///opt/spark/examples/jars/spark-examples_2.12-3.2.4.jar",
                "mainClass": "org.apache.spark.examples.SparkPi",
                "mode": "cluster",
                "restartPolicy": {
                  "type": "Never"
                },
                "sparkConf": {
                  "javaOptions": "-Divy.cache.dir=/tmp -Divy.home=/tmp",
                  "spark.serializer": "org.apache.spark.serializer.KryoSerializer"
                },
                "sparkVersion": "3.2.4",
                "type": "Scala",
                "volumes": [
                  {
                    "hostPath": {
                      "path": "/tmp",
                      "type": "Directory"
                    },
                    "name": "test-volume"
                  }
                ]
              }
            },
            "oldObject": null,
            "dryRun": false,
            "options": {
              "kind": "CreateOptions",
              "apiVersion": "meta.k8s.io/v1",
              "fieldManager": "kubectl-client-side-apply"
            }
          }
        }'