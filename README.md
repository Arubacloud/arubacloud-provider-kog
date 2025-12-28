# Aruba Cloud Provider KOG Blueprint

[![GitHub release](https://img.shields.io/github/tag/arubacloud/arubacloud-resource-operator.svg?label=release)](https://github.com/arubacloud/arubacloud-resource-operator/releases/latest) [![Tests](https://github.com/arubacloud/arubacloud-resource-operator/actions/workflows/test.yml/badge.svg)](https://github.com/arubacloud/arubacloud-resource-operator/actions/workflows/test.yml) [![Release](https://github.com/arubacloud/arubacloud-resource-operator/actions/workflows/release.yml/badge.svg)](https://github.com/arubacloud/arubacloud-resource-operator/actions/workflows/release.yml)

> **⚠️ Development Status**: This operator is currently under active development and is **not production-ready yet**. APIs and resource schemas may change. Use at your own risk in production environments.

***KOG***: (*Krateo Operator Generator*)

This is a Krateo Blueprint that deploys the Aruba Cloud Provider KOG leveraging the [OASGen Provider](https://github.com/krateoplatformops/oasgen-provider) and the [Aruba Cloud API](https://api.arubacloud.com/docs/intro).
This provider allows you to manage Aruba Cloud resources in a cloud-native way using the Krateo platform.

**Supported resource categories:**
- **Compute** - Virtual servers and SSH key pairs (CloudServer, KeyPair)
- **Container** - Kubernetes and container registries (KaaS, ContainerRegistry)
- **Database** - Managed database services (DBaaS, Database, User, Grant, Backup)
- **Network** - Networking resources (VPC, Subnet, SecurityGroup, SecurityRule, ElasticIP, LoadBalancer, VPNTunnel, VPCPeering, VPCPeeringRoute)
- **Storage** - Block storage and backups (BlockStorage, Snapshot, Backup, Restore)
- **Project** - Project management (Project)
- **Schedule** - Job scheduling (Job)
- **Security** - Key management (KMS)

## Summary

- [Requirements](#requirements)
- [Project structure](#project-structure)
- [How to install](#how-to-install)
  - [Full provider installation](#full-provider-installation)
  - [Single resource installation](#single-resource-installation)
- [OpenAPI Specification](#openapi-specification)
- [Supported resources](#supported-resources)
  - [Available Blueprints](#available-blueprints)
  - [Resource details](#resource-details)
    - [Subnet](#subnet)
    - [CloudServer](#cloudserver)
    - [DBaaS](#dbaas)
  - [Resource examples](#resource-examples)
- [Authentication](#authentication)
- [Configuration](#configuration)
  - [Configuration resources](#configuration-resources)
  - [values.yaml](#valuesyaml)
  - [Verbose logging](#verbose-logging)
- [Charts structure](#charts-structure)
- [Troubleshooting](#troubleshooting)
- [Release process](#release-process)

## Requirements

[OASGen Provider](https://github.com/krateoplatformops/oasgen-provider) should be installed in your cluster with version >= 0.7.1.

Follow the related Helm Chart [README](https://github.com/krateoplatformops/oasgen-provider-chart) for installation instructions.
Note that a standard installation of Krateo contains the OASGen Provider.

## Project structure

This project is composed by the following folders:
- **arubacloud-provider-kog-*-blueprint**: Helm charts that deploys single resources supported by this provider. These charts are useful if you want to deploy only one of the supported resources.
- **arubacloud-provider-kog-blueprint**: a Helm chart that can deploy all resources supported by this provider. It is useful if you want to manage multiple of the supported resources.
- **plugins**: a folder that is a monorepo containing multiple Go plugins. If needed, they are deployed as part of the Helm chart of the specific resource.

## How to install

### Full provider installation

To install the **arubacloud-provider-kog-blueprint** Helm chart (full provider), use the following command:

```sh
helm install arubacloud-provider-kog arubacloud-provider-kog \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --version 1.0.0 \
  --wait
```

> [!NOTE]
> Due to the nature of the providers leveraging the [OASGen Provider](https://github.com/krateoplatformops/oasgen-provider), this chart will install a set of RestDefinitions that will in turn trigger the deployment of a set controllers in the cluster. These controllers need to be up and running before you can create or manage resources using the Custom Resources (CRs) defined by this provider. This may take a few minutes after the chart is installed. The RestDefinitions will reach the condition `Ready` when the related CRDs are installed and the controllers are up and running.

You can check the status of the RestDefinitions with the following commands:

```sh
kubectl get restdefinitions.ogen.krateo.io --all-namespaces | awk 'NR==1 || /arubacloud/'
```
You should see output similar to this:
```sh
NAMESPACE       NAME                                 READY   AGE
krateo-system   arubacloud-provider-kog-subnet       False   59s
```

You can also wait for a specific RestDefinition (`arubacloud-provider-kog-subnet` in this case) to be ready with a command like this:
```sh
kubectl wait restdefinitions.ogen.krateo.io arubacloud-provider-kog-subnet --for condition=Ready=True --namespace krateo-system --timeout=300s
```

Note that the names of the RestDefinitions and the namespace where the RestDefinitions are installed may vary based on your configuration.

### Single resource installation

To manage resources from a specific category, you can install the dedicated Helm chart. Here are examples for each category:

**Subnet resources:**
```sh
helm install arubacloud-provider-kog-subnet arubacloud-provider-kog-subnet \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --version 1.0.0 \
  --wait
```

**Compute resources (CloudServer, KeyPair):**
```sh
helm install arubacloud-provider-kog-compute arubacloud-provider-kog-compute-blueprint \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --wait
```

**Container resources (KaaS, ContainerRegistry):**
```sh
helm install arubacloud-provider-kog-container arubacloud-provider-kog-container-blueprint \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --wait
```

**Database resources (DBaaS, Database, User, Grant, Backup):**
```sh
helm install arubacloud-provider-kog-database arubacloud-provider-kog-database-blueprint \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --wait
```

**Network resources (VPC, Subnet, SecurityGroup, SecurityRule, ElasticIP, LoadBalancer, VPNTunnel, VPCPeering, VPCPeeringRoute):**
```sh
helm install arubacloud-provider-kog-network arubacloud-provider-kog-network-blueprint \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --wait
```

**Storage resources (BlockStorage, Snapshot, Backup, Restore):**
```sh
helm install arubacloud-provider-kog-storage arubacloud-provider-kog-storage-blueprint \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --wait
```

**Project resources:**
```sh
helm install arubacloud-provider-kog-project arubacloud-provider-kog-project-blueprint \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --wait
```

**Schedule resources (Job):**
```sh
helm install arubacloud-provider-kog-schedule arubacloud-provider-kog-schedule-blueprint \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --wait
```

**Security resources (KMS):**
```sh
helm install arubacloud-provider-kog-security arubacloud-provider-kog-security-blueprint \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --wait
```

## OpenAPI Specification

The OpenAPI Specification used for this provider is derived from the one provided by Aruba Cloud which can be found at the following URL: https://api.arubacloud.com/openapi/network-provider.json.

## Supported resources

This provider supports the following resources across multiple categories:

### Compute Resources
| Resource     | Get  | Create | Update | Delete |
|--------------|------|--------|--------|--------|
| CloudServer  | ✅   | ✅     | ✅     | ✅     |
| KeyPair      | ✅   | ✅     | ✅     | ✅     |

### Container Resources
| Resource           | Get  | Create | Update | Delete |
|-------------------|------|--------|--------|--------|
| KaaS              | ✅   | ✅     | ✅     | ✅     |
| ContainerRegistry | ✅   | ✅     | ✅     | ✅     |

### Database Resources
| Resource   | Get  | Create | Update | Delete |
|-----------|------|--------|--------|--------|
| DBaaS     | ✅   | ✅     | ✅     | ✅     |
| Database  | ✅   | ✅     | ✅     | ✅     |
| User      | ✅   | ✅     | ✅     | ✅     |
| Grant     | ✅   | ✅     | ✅     | ✅     |
| Backup    | ✅   | ✅     | ✅     | ✅     |

### Network Resources
| Resource         | Get  | Create | Update | Delete |
|-----------------|------|--------|--------|--------|
| VPC             | ✅   | ✅     | ✅     | ✅     |
| Subnet          | ✅   | ✅     | ✅     | ✅     |
| SecurityGroup   | ✅   | ✅     | ✅     | ✅     |
| SecurityRule    | ✅   | ✅     | ✅     | ✅     |
| ElasticIP       | ✅   | ✅     | ✅     | ✅     |
| LoadBalancer    | ✅   | ❌     | ❌     | ❌     |
| VPNTunnel       | ✅   | ✅     | ✅     | ✅     |
| VPCPeering      | ✅   | ✅     | ✅     | ✅     |
| VPCPeeringRoute | ✅   | ✅     | ✅     | ✅     |

> **Note**: LoadBalancer is read-only. Only GET (list and retrieve) operations are supported. Create, update, and delete operations are not available via the REST API.

### Storage Resources
| Resource      | Get  | Create | Update | Delete |
|--------------|------|--------|--------|--------|
| BlockStorage | ✅   | ✅     | ✅     | ✅     |
| Snapshot     | ✅   | ✅     | ✅     | ✅     |
| Backup       | ✅   | ✅     | ✅     | ✅     |
| Restore      | ✅   | ✅     | ✅     | ✅     |

### Project Resources
| Resource | Get  | Create | Update | Delete |
|---------|------|--------|--------|--------|
| Project | ✅   | ✅     | ✅     | ✅     |

### Schedule Resources
| Resource | Get  | Create | Update | Delete |
|---------|------|--------|--------|--------|
| Job     | ✅   | ✅     | ✅     | ✅     |

### Security Resources
| Resource | Get  | Create | Update | Delete |
|---------|------|--------|--------|--------|
| KMS     | ✅   | ✅     | ✅     | ✅     |

The resources listed above are Custom Resources (CRs) defined in resource group-specific API groups (e.g., `compute.ogen-krateo.arubacloud.com`, `network.ogen-krateo.arubacloud.com`, etc.). They are used to manage Aruba Cloud resources in a Kubernetes-native way, allowing you to create, update, and delete Aruba Cloud resources using Kubernetes manifests.

## Available Blueprints

This project provides separate Helm charts for each resource category:

- **arubacloud-provider-kog-compute-blueprint** - Manages compute resources (CloudServer, KeyPair)
- **arubacloud-provider-kog-container-blueprint** - Manages container resources (KaaS, ContainerRegistry)
- **arubacloud-provider-kog-database-blueprint** - Manages database resources (DBaaS, Database, User, Grant, Backup)
- **arubacloud-provider-kog-network-blueprint** - Manages network resources (VPC, Subnet, SecurityGroup, SecurityRule, ElasticIP, LoadBalancer, VPNTunnel, VPCPeering, VPCPeeringRoute)
- **arubacloud-provider-kog-storage-blueprint** - Manages storage resources (BlockStorage, Snapshot, Backup, Restore)
- **arubacloud-provider-kog-project-blueprint** - Manages project resources
- **arubacloud-provider-kog-schedule-blueprint** - Manages scheduling resources (Job)
- **arubacloud-provider-kog-security-blueprint** - Manages security resources (KMS)
- **arubacloud-provider-kog-blueprint** - Umbrella chart that includes all of the above

### Resource details

#### Subnet

The `Subnet` resource allows you to create, update, and delete Aruba Cloud subnets.
You can specify the subnet name, location, tags, type, and other settings such as DHCP configuration and routes.

An example of a Subnet resource is:
```yaml
apiVersion: network.ogen-krateo.arubacloud.com/v1alpha1
kind: Subnet
metadata:
  name: test-subnet-kog-123-complete
  namespace: default
  annotations:
    krateo.io/connector-verbose: "true"
spec:
  configurationRef:
    name: example-configuration
    namespace: config-namespace
  projectId: "proj-12345"
  vpcId: "vpc-67890"
  name: "test-subnet-kog-123-complete"
  location:
    value: "ITBG-Bergamo"
  #newDefaultSubnet: "" # URI for existing subnet to set as default, if needed during deletion of this subnet
  tags:
  - "tag1"
  - "tag2"
  properties:
    default: false
    type: "Advanced" # allowed values: {Basic, Advanced}
    network:
      address: "10.1.0.0/24"
    dhcp:
      enabled: true
      dns:
        - "8.8.8.8"
        - "8.8.4.4"
      range:
        start: "10.1.0.10"
        count: 200
      #routes:
      #  - address: "192.168.0.0/16"
      #    gateway: "10.1.0.11"
      #  - address: "172.16.0.0/12"
      #    gateway: "10.1.0.12"
```

#### CloudServer

The `CloudServer` resource allows you to create, update, and delete Aruba Cloud virtual servers.
You can specify the server name, location, instance type, image, network configuration, and other settings.

An example of a CloudServer resource is:
```yaml
apiVersion: compute.ogen-krateo.arubacloud.com/v1alpha1
kind: CloudServer
metadata:
  name: my-cloud-server
  namespace: default
  annotations:
    krateo.io/connector-verbose: "true"
spec:
  configurationRef:
    name: cloudserver-config
    namespace: default
  projectId: "proj-12345"
  name: "my-cloud-server"
  location:
    value: "ITBG-Bergamo"
  properties:
    instanceType: "LS"  # Small instance
    image: "ubuntu-22.04"
    networkConfiguration:
      vpcId: "vpc-67890"
      subnetId: "subnet-12345"
    tags:
      - "production"
      - "web-server"
```

#### DBaaS

The `DBaaS` resource allows you to create, update, and delete Aruba Cloud managed database instances.
You can specify the database engine, version, instance type, storage, and other configuration options.

An example of a DBaaS resource is:
```yaml
apiVersion: database.ogen-krateo.arubacloud.com/v1alpha1
kind: DBaaS
metadata:
  name: my-postgres-db
  namespace: default
  annotations:
    krateo.io/connector-verbose: "true"
spec:
  configurationRef:
    name: dbaas-config
    namespace: default
  projectId: "proj-12345"
  name: "my-postgres-db"
  location:
    value: "ITBG-Bergamo"
  properties:
    engine: "postgresql"
    version: "15"
    instanceType: "db.t3.medium"
    storageSize: 100  # GB
    backupRetentionPeriod: 7  # days
```

### Resource examples

You can find example resources for each supported resource type in the `/samples` folder of each blueprint chart.
For instance:
- Subnet examples: `arubacloud-provider-kog-subnet-blueprint/samples/`
- Compute examples: `arubacloud-provider-kog-compute-blueprint/samples/`
- Database examples: `arubacloud-provider-kog-database-blueprint/samples/`
- And so on for each category...

The umbrella chart (`arubacloud-provider-kog-blueprint`) also includes sample files for quick reference.

**Sample Configuration Files**: Ready-to-use configuration examples for all resource types are available in `arubacloud-provider-kog-blueprint/samples/configs/`. These files are pre-configured to reference the `aruba-access-token` secret in the `default` namespace and use the correct API versions for each resource group.

## Authentication

The authentication to the Aruba Cloud API is managed using 2 kinds of resources (both are required):

- **Kubernetes Secret**: This resource is used to store the Aruba Cloud Token that is used to authenticate with the Aruba Cloud API. 

In order to generate a Aruba Cloud token, follow these instructions: https://api.arubacloud.com/docs/authentication/.

Note that the token has a limited validity (default 1 hour) and needs to be regenerated periodically.
Specific solution for token rotation are not covered in this chart and should be implemented by the user if needed.

Example of a Kubernetes Secret that you can apply to your cluster:
```sh
kubectl apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: aruba-access-token
  namespace: default
type: Opaque
stringData:
  token: <YOUR_TOKEN>
EOF
```

Replace `<YOUR_TOKEN>` with your actual Aruba Cloud Token (without quotes and without `Bearer ` prefix).

- **\<Resource\>Configuration**: These resources can reference the Kubernetes Secret and are used to authenticate with the Aruba Cloud API. They must be referenced with the `configurationRef` field of the resources defined in this chart. The configuration resource can be in a different namespace than the resource itself.

Note that the specific configuration resource type depends on the resource you are managing:
- For `Subnet` resources: use `SubnetConfiguration`
- For `CloudServer` resources: use `CloudServerConfiguration`
- For `KeyPair` resources: use `KeyPairConfiguration`
- For `KaaS` resources: use `KaaSConfiguration`
- For `DBaaS` resources: use `DbaasConfiguration`
- For `VPC` resources: use `VpcConfiguration`
- For `BlockStorage` resources: use `BlockStorageConfiguration`
- And so on for each resource type...

An example of a `SubnetConfiguration` resource that references the Kubernetes Secret:
```sh
kubectl apply -f - <<EOF
apiVersion: network.ogen-krateo.arubacloud.com/v1alpha1
kind: SubnetConfiguration
metadata:
  name: my-subnet-config
  namespace: default
spec:
  authentication:
    bearer:
      tokenRef:
        name: aruba-access-token
        namespace: default
        key: token
  configuration:
    query:
      create:
        api-version: "1.0"
      delete:
        api-version: "1.0"
      get:
        api-version: "1.0"
      update:
        api-version: "1.0"
      findby:
        api-version: "1.0"
EOF
```

Then, in the `Subnet` resource, you can reference the `SubnetConfiguration` resource as follows:
```yaml
apiVersion: network.ogen-krateo.arubacloud.com/v1alpha1
kind: Subnet
metadata:
  name: test-subnet-kog-123
  namespace: default
  annotations:
    krateo.io/connector-verbose: "true"
spec:
  configurationRef:
    name: my-subnet-config
    namespace: default 
  projectId: ABCDEFGHIJKLMN
  vpcId: ABC1234567890
  name: test-subnet-kog-123
```

Similarly, for a `CloudServer` resource, you would create a `CloudServerConfiguration`:
```sh
kubectl apply -f - <<EOF
apiVersion: compute.ogen-krateo.arubacloud.com/v1alpha1
kind: CloudServerConfiguration
metadata:
  name: my-cloudserver-config
  namespace: default
spec:
  authentication:
    bearer:
      tokenRef:
        name: aruba-access-token
        namespace: default
        key: token
  configuration:
    query:
      create:
        api-version: "1.0"
      delete:
        api-version: "1.0"
      get:
        api-version: "1.0"
      update:
        api-version: "1.0"
      findby:
        api-version: "1.0"
EOF
```

And reference it in your `CloudServer` resource:
```yaml
apiVersion: compute.ogen-krateo.arubacloud.com/v1alpha1
kind: CloudServer
metadata:
  name: my-server
  namespace: default
spec:
  configurationRef:
    name: my-cloudserver-config
    namespace: default
  projectId: ABCDEFGHIJKLMN
  name: my-server
  # ... other properties
```

More details about the configuration resources in the [Configuration resources](#configuration-resources) section below.

## Configuration

### Configuration resources

Each resource type requires a specific configuration resource to be created in the cluster.

**Currently supported configuration resources:**
- `CloudServerConfiguration` - For CloudServer resources (API: `compute.ogen-krateo.arubacloud.com/v1alpha1`)
- `KeyPairConfiguration` - For KeyPair resources (API: `compute.ogen-krateo.arubacloud.com/v1alpha1`)
- `KaaSConfiguration` - For KaaS resources (API: `container.ogen-krateo.arubacloud.com/v1alpha1`)
- `ContainerRegistryConfiguration` - For ContainerRegistry resources (API: `container.ogen-krateo.arubacloud.com/v1alpha1`)
- `DBaaSConfiguration` - For DBaaS resources (API: `database.ogen-krateo.arubacloud.com/v1alpha1`)
- `DBaaSDatabaseConfiguration` - For Database resources (API: `database.ogen-krateo.arubacloud.com/v1alpha1`)
- `DBaaSUserConfiguration` - For User resources (API: `database.ogen-krateo.arubacloud.com/v1alpha1`)
- `DBaaSGrantConfiguration` - For Grant resources (API: `database.ogen-krateo.arubacloud.com/v1alpha1`)
- `BackupConfiguration` - For Backup resources (database) (API: `database.ogen-krateo.arubacloud.com/v1alpha1`)
- `VPCConfiguration` - For VPC resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`)
- `SubnetConfiguration` - For Subnet resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`)
- `SecurityGroupConfiguration` - For SecurityGroup resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`)
- `SecurityRuleConfiguration` - For SecurityRule resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`)
- `ElasticIPConfiguration` - For ElasticIP resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`)
- `LoadBalancerConfiguration` - For LoadBalancer resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`) - **Read-only resource**
- `VPNTunnelConfiguration` - For VPNTunnel resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`)
- `VPCPeeringConfiguration` - For VPCPeering resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`)
- `VPCPeeringRouteConfiguration` - For VPCPeeringRoute resources (API: `network.ogen-krateo.arubacloud.com/v1alpha1`)
- `BlockStorageConfiguration` - For BlockStorage resources (API: `storage.ogen-krateo.arubacloud.com/v1alpha1`)
- `SnapshotConfiguration` - For Snapshot resources (API: `storage.ogen-krateo.arubacloud.com/v1alpha1`)
- `BackupConfiguration` - For Backup resources (storage) (API: `storage.ogen-krateo.arubacloud.com/v1alpha1`)
- `RestoreConfiguration` - For Restore resources (API: `storage.ogen-krateo.arubacloud.com/v1alpha1`)
- `ProjectConfiguration` - For Project resources (API: `project.ogen-krateo.arubacloud.com/v1alpha1`)
- `JobConfiguration` - For Job resources (API: `schedule.ogen-krateo.arubacloud.com/v1alpha1`)
- `KMSConfiguration` - For KMS resources (API: `security.ogen-krateo.arubacloud.com/v1alpha1`)

These configuration resources are used to store the authentication information (i.e., reference to the Kubernetes Secret containing the Aruba Cloud Token) and other configuration options for the resource type.

**Important**: Each configuration resource uses a specific API version based on its resource group:
- Compute resources: `compute.ogen-krateo.arubacloud.com/v1alpha1`
- Container resources: `container.ogen-krateo.arubacloud.com/v1alpha1`
- Database resources: `database.ogen-krateo.arubacloud.com/v1alpha1`
- Network resources: `network.ogen-krateo.arubacloud.com/v1alpha1`
- Storage resources: `storage.ogen-krateo.arubacloud.com/v1alpha1`
- Project resources: `project.ogen-krateo.arubacloud.com/v1alpha1`
- Schedule resources: `schedule.ogen-krateo.arubacloud.com/v1alpha1`
- Security resources: `security.ogen-krateo.arubacloud.com/v1alpha1`

You can find example configuration files for all resource types in the `/samples/configs` folder of the `arubacloud-provider-kog-blueprint` chart. These sample files are ready to use and reference the `aruba-access-token` secret in the `default` namespace.

Note that a single configuration resource can be used by multiple resources of the same type.
For example, you can create a single `SubnetConfiguration` resource and reference it in multiple `Subnet` resources.

### values.yaml

You can customize the **arubacloud-provider-kog-blueprint** chart (the umbrella chart) by modifying the `values.yaml` file.
For instance, you can select which category of resources the provider should support by enabling or disabling individual blueprints.
This may be useful if you want to limit the resources managed by the provider to only those you need, reducing the overhead of managing unnecessary controllers.

The umbrella chart supports the following blueprints that can be enabled/disabled:
- `arubacloud-provider-kog-subnet-blueprint` - Subnet resources
- `arubacloud-provider-kog-compute-blueprint` - Compute resources (CloudServer, KeyPair)
- `arubacloud-provider-kog-container-blueprint` - Container resources (KaaS, ContainerRegistry)
- `arubacloud-provider-kog-database-blueprint` - Database resources (DBaaS and related)
- `arubacloud-provider-kog-network-blueprint` - Network resources (VPC, SecurityGroup, etc.)
- `arubacloud-provider-kog-storage-blueprint` - Storage resources (BlockStorage, Snapshot, Backup, Restore)
- `arubacloud-provider-kog-project-blueprint` - Project resources
- `arubacloud-provider-kog-schedule-blueprint` - Schedule resources (Job)
- `arubacloud-provider-kog-security-blueprint` - Security resources (KMS)

By default, all blueprints are enabled.

### Verbose logging

In order to enable verbose logging for the controllers, you can add the `krateo.io/connector-verbose: "true"` annotation to the metadata of the resources you want to manage, as shown in the examples above. 
This will enable verbose logging for those specific resources, which can be useful for debugging and troubleshooting as it will provide more detailed information about the operations performed by the controllers.

## Charts structure

Main components of the charts:

- **RestDefinitions**: These are the core resources needed to manage resources leveraging the OASGen Provider. In this case, they refers to the OpenAPI Specification to be used for the creation of the Custom Resources (CRs) that represent Aruba Cloud resources.
They also define the operations that can be performed on those resources. Once the chart is installed, RestDefinitions will be created and as a result, specific controllers will be deployed in the cluster to manage the resources defined with those RestDefinitions.

- **ConfigMaps**: Refer directly to the OpenAPI Specification content in the `/assets` folder.

- **/assets** folder: Contains the selected OpenAPI Specification files for the Aruba Cloud API.

- **Deployment** (optional): Deploys a plugin that is used as a proxy to resolve some integration issue with Aruba Cloud. The specific endpoins managed by the plugin are described in the [plugins README](./plugins/README.md)

- **Service** (optional): Exposes the plugin described above, allowing the resource controllers to communicate with the Aruba Cloud API through the plugin, only if needed.

## Troubleshooting

For troubleshooting, you can refer to the [Troubleshooting guide](./arubacloud-provider-kog-blueprint/docs/troubleshooting.md) in the `/docs` folder of the main blueprint (chart). 
It contains common issues and solutions related to this chart.

## Release process

Please refer to the [Release guide](./docs/release.md) in the `/docs` folder for detailed instructions on how to release new versions of the chart and its components.


