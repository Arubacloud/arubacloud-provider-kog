
# Aruba Cloud Provider KOG Blueprint

This is a Helm chart for deploying the Aruba Cloud Provider KOG Blueprint.
It acts as an umbrella chart and it includes all the other blueprints:
- arubacloud-provider-kog-compute-blueprint (CloudServer, KeyPair)
- arubacloud-provider-kog-container-blueprint (KaaS, ContainerRegistry)
- arubacloud-provider-kog-database-blueprint (DBaaS, Database, User, Grant, Backup)
- arubacloud-provider-kog-project-blueprint (Project)
- arubacloud-provider-kog-schedule-blueprint (Job)
- arubacloud-provider-kog-storage-blueprint (BlockStorage, Snapshot, Backup, Restore)
- arubacloud-provider-kog-security-blueprint (KMS)
- arubacloud-provider-kog-network-blueprint (VPC, Subnet, SecurityGroup, SecurityRule, ElasticIP, LoadBalancer, VPNTunnel, VPCPeering, VPCPeeringRoute)

## Requirements

**OASGen Provider** should be installed in your cluster with version >= 0.7.1.

Follow the related [Helm Chart README](https://github.com/krateoplatformops/oasgen-provider) for installation instructions. Note that a standard installation of Krateo contains the OASGen Provider.

To install OASGen Provider:

```bash
helm repo add krateo https://charts.krateo.io
helm repo update
helm install oasgen-provider krateo/oasgen-provider --namespace krateo-system --create-namespace
```

## Installation

### Production Installation (from Krateo Marketplace)

Install the chart from the Krateo marketplace:

```bash
helm install arubacloud-provider-kog arubacloud-provider-kog \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --create-namespace \
  --version 1.0.0 \
  --wait
```

### Local Development Installation

For local development, first update the chart dependencies:

```bash
cd arubacloud-provider-kog-blueprint
helm dependency update
```

Then install the chart from the local directory:

```bash
helm upgrade --install arubacloud-provider-kog . \
  --namespace arubacloud-system \
  --create-namespace \
  --wait
```

Or install with a custom values file:

```bash
helm install arubacloud-provider-kog . \
  --namespace <release-namespace> \
  --create-namespace \
  --values custom-values.yaml \
  --wait
```

### Upgrade

To upgrade an existing installation:

```bash
# For production
helm upgrade arubacloud-provider-kog arubacloud-provider-kog \
  --repo https://marketplace.krateo.io \
  --namespace <release-namespace> \
  --version 1.0.0 \
  --wait

# For local development
helm upgrade arubacloud-provider-kog . \
  --namespace <release-namespace> \
  --wait
```

### Uninstall

To uninstall the chart:

```bash
helm uninstall arubacloud-provider-kog --namespace <release-namespace>
```

## Configuration

You can selectively enable or disable specific blueprints by setting their `enabled` value to `false` in your values file:

```yaml
arubacloud-provider-kog-compute-blueprint:
  enabled: false  # Disable compute resources

arubacloud-provider-kog-database-blueprint:
  enabled: false  # Disable database resources
```

