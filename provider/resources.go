// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ovh

import (
	"context"
	_ "embed"
	"fmt"
	"path"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/ovh/pulumi-ovh/provider/v2/pkg/version"
	"github.com/ovh/terraform-provider-ovh/v2/ovh"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	ovhPkg = "ovh"
	// modules:
	ovhMod             = "index"
	cloudMod           = "Cloud"
	cloudProjectMod    = "CloudProject"
	cloudProjectDbMod  = "CloudProjectDatabase"
	dbaasMod           = "Dbaas"
	dedicatedMod       = "Dedicated"
	domainMod          = "Domain"
	hostingMod         = "Hosting"
	ipMod              = "Ip"
	ipLoadBalancingMod = "IpLoadBalancing"
	meMod              = "Me"
	vrackMod           = "Vrack"
	orderMod           = "Order"
	vpsMod             = "Vps"
	iamMod             = "Iam"
	okmsMod            = "Okms"
	savingPlansMod     = "SavingsPlan"
	ovhCloudMod        = "OVHcloud"
	vmwareMod          = "VMware"
)

// ovhDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Ovh package and names the file by simply lower casing the data
// source's first character.
func ovhDataSource(mod string, res string) tokens.ModuleMember {
	return tfbridge.MakeDataSource(ovhPkg, mod, res)
}

// ovhResource manufactures a standard resource token given a module and resource name.
// It automatically uses the ovh package and names the file by simply lower casing the resource's
// first character.
func ovhResource(mod string, res string) tokens.Type {
	return tfbridge.MakeResource(ovhPkg, mod, res)
}

//go:embed cmd/pulumi-resource-ovh/bridge-metadata.json
var bridgeMetadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pfbridge.MuxShimWithPF(
		context.Background(),
		shimv2.NewProvider(ovh.Provider()),
		&ovh.OvhProvider{},
	)

	// ComputeID
	delegateID := func(pulumiField string) tfbridge.ComputeID {
		return tfbridge.DelegateIDField(resource.PropertyKey(pulumiField),
			"ovh", "https://github.com/ovh/pulumi-ovh")
	}

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "ovh",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "OVHCloud",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "OVHcloud",

		// This is now required.
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),

		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/ovh/pulumi-ovh",
		Description:       "A Pulumi package for creating and managing OVH resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "ovh", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/ovh/pulumi-ovh",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:               "ovh",
		TFProviderModuleVersion: "v2",
		Version:                 version.Version,
		Config:                  map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"ovh_cloud_project": {
				Tok: ovhResource(cloudProjectMod, "Project"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "ProjectURN",
					},
				},
			},
			"ovh_cloud_project_alerting": {
				Tok: ovhResource(cloudProjectMod, "Alerting"),
			},
			"ovh_cloud_project_containerregistry": {
				Tok: ovhResource(cloudProjectMod, "ContainerRegistry"),
			},
			"ovh_cloud_project_containerregistry_user": {
				Tok: ovhResource(cloudProjectMod, "ContainerRegistryUser"),
			},
			"ovh_cloud_project_containerregistry_oidc": {
				Tok: ovhResource(cloudProjectMod, "ContainerRegistryOIDC"),
			},
			"ovh_cloud_project_containerregistry_ip_restrictions_management": {
				Tok: ovhResource(cloudProjectMod, "ContainerRegistryIPRestrictionsManagement"),
			},
			"ovh_cloud_project_containerregistry_ip_restrictions_registry": {
				Tok: ovhResource(cloudProjectMod, "ContainerRegistryIPRestrictionsRegistry"),
			},
			"ovh_cloud_project_database": {
				Tok: ovhResource(cloudProjectMod, "Database"),
			},
			"ovh_cloud_project_database_database": {
				Tok: ovhResource(cloudProjectDbMod, "DatabaseInstance"),
			},
			"ovh_cloud_project_database_ip_restriction": {
				Tok: ovhResource(cloudProjectDbMod, "IpRestriction"),
			},
			"ovh_cloud_project_database_integration": {
				Tok: ovhResource(cloudProjectDbMod, "Integration"),
			},
			"ovh_cloud_project_database_log_subscription": {
				Tok: ovhResource(cloudProjectDbMod, "LogSubscription"),
			},
			"ovh_cloud_project_database_kafka_acl": {
				Tok: ovhResource(cloudProjectDbMod, "KafkaAcl"),
			},
			"ovh_cloud_project_database_kafka_schemaregistryacl": {
				Tok: ovhResource(cloudProjectDbMod, "KafkaSchemaRegistryAcl"),
			},
			"ovh_cloud_project_database_kafka_topic": {
				Tok: ovhResource(cloudProjectDbMod, "KafkaTopic"),
			},
			"ovh_cloud_project_database_postgresql_user": {
				Tok: ovhResource(cloudProjectDbMod, "PostgresSqlUser"),
			},
			"ovh_cloud_project_database_postgresql_connection_pool": {
				Tok: ovhResource(cloudProjectDbMod, "PostgresSqlConnectionPool"),
			},
			"ovh_cloud_project_database_prometheus": {
				Tok: ovhResource(cloudProjectDbMod, "Prometheus"),
			},
			"ovh_cloud_project_database_m3db_namespace": {
				Tok: ovhResource(cloudProjectDbMod, "M3DbNamespace"),
			},
			"ovh_cloud_project_database_m3db_user": {
				Tok: ovhResource(cloudProjectDbMod, "M3DbUser"),
			},
			"ovh_cloud_project_database_mongodb_prometheus": {
				Tok: ovhResource(cloudProjectDbMod, "MongoDbPrometheus"),
			},
			"ovh_cloud_project_database_mongodb_user": {
				Tok: ovhResource(cloudProjectDbMod, "MongoDbUser"),
			},
			"ovh_cloud_project_database_opensearch_pattern": {
				Tok: ovhResource(cloudProjectDbMod, "OpensearchPattern"),
			},
			"ovh_cloud_project_database_opensearch_user": {
				Tok: ovhResource(cloudProjectDbMod, "OpensearchUser"),
			},
			"ovh_cloud_project_database_redis_user": {
				Tok: ovhResource(cloudProjectDbMod, "RedisUser"),
			},
			"ovh_cloud_project_database_user": {
				Tok: ovhResource(cloudProjectDbMod, "User"),
			},
			"ovh_cloud_project_failover_ip_attach": {
				Tok: ovhResource(cloudProjectMod, "FailoverIpAttach"),
			},
			"ovh_cloud_project_gateway": {
				Tok: ovhResource(cloudProjectMod, "Gateway"),
			},
			"ovh_cloud_project_gateway_interface": {
				Tok:       ovhResource(cloudProjectMod, "GatewayInterface"),
				ComputeID: delegateID("id"),
			},
			"ovh_cloud_project_kube": {
				Tok: ovhResource(cloudProjectMod, "Kube"),
			},
			"ovh_cloud_project_kube_iprestrictions": {
				Tok: ovhResource(cloudProjectMod, "KubeIpRestrictions"),
			},
			"ovh_cloud_project_kube_nodepool": {
				Tok: ovhResource(cloudProjectMod, "KubeNodePool"),
			},
			"ovh_cloud_project_kube_oidc": {
				Tok: ovhResource(cloudProjectMod, "KubeOidc"),
			},
			"ovh_cloud_project_instance": {
				Tok: ovhResource(cloudProjectMod, "Instance"),
			},
			"ovh_cloud_project_instance_snapshot": {
				Tok: ovhResource(cloudProjectMod, "InstanceSnapshot"),
			},
			"ovh_cloud_project_loadbalancer": {
				Tok: ovhResource(cloudProjectMod, "LoadBalancer"),
			},
			"ovh_cloud_project_network_private": {
				Tok: ovhResource(cloudProjectMod, "NetworkPrivate"),
			},
			"ovh_cloud_project_network_private_subnet": {
				Tok: ovhResource(cloudProjectMod, "NetworkPrivateSubnet"),
			},
			"ovh_cloud_project_network_private_subnet_v2": {
				Tok: ovhResource(cloudProjectMod, "NetworkPrivateSubnetV2"),
			},
			"ovh_cloud_project_region_network": {
				Tok: ovhResource(cloudProjectMod, "RegionNetwork"),
			},
			"ovh_cloud_project_rancher": {
				Tok: ovhResource(cloudProjectMod, "Rancher"),
			},
			"ovh_cloud_project_region_storage_presign": {
				Tok: ovhResource(cloudProjectMod, "RegionStoragePresign"),
			},
			"ovh_cloud_project_storage": {
				Tok:       ovhResource(cloudProjectMod, "Storage"),
				ComputeID: delegateID("name"),
			},
			"ovh_cloud_project_region_loadbalancer_log_subscription": {
				Tok: ovhResource(cloudProjectMod, "RegionLoadBalancerLogSubscription"),
			},
			"ovh_cloud_project_user": {
				Tok: ovhResource(cloudProjectMod, "User"),
			},
			"ovh_cloud_project_user_s3_credential": {
				Tok: ovhResource(cloudProjectMod, "S3Credential"),
			},
			"ovh_cloud_project_user_s3_policy": {
				Tok: ovhResource(cloudProjectMod, "S3Policy"),
			},
			"ovh_cloud_project_volume": {
				Tok: ovhResource(cloudProjectMod, "Volume"),
			},
			"ovh_cloud_project_volume_backup": {
				Tok: ovhResource(cloudProjectMod, "VolumeBackup"),
			},
			"ovh_dbaas_logs_input": {
				Tok: ovhResource(dbaasMod, "LogsInput"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: ovhResource(dbaasMod, "LogsOutputGraylogStream"),
			},
			"ovh_dbaas_logs_output_opensearch_alias": {
				Tok: ovhResource(dbaasMod, "LogsOutputOpenSearchAlias"),
			},
			"ovh_dbaas_logs_output_opensearch_index": {
				Tok: ovhResource(dbaasMod, "LogsOutputOpenSearchIndex"),
			},
			"ovh_dbaas_logs_role": {
				Tok: ovhResource(dbaasMod, "LogsRole"),
			},
			"ovh_dbaas_logs_role_permission_stream": {
				Tok: ovhResource(dbaasMod, "LogsRolePermissionStream"),
			},
			"ovh_dedicated_ceph_acl": {
				Tok: ovhResource(dedicatedMod, "CephAcl"),
			},
			"ovh_dedicated_nasha_partition": {
				Tok: ovhResource(dedicatedMod, "NasHAPartition"),
			},
			"ovh_dedicated_nasha_partition_access": {
				Tok: ovhResource(dedicatedMod, "NasHAPartitionAccess"),
			},
			"ovh_dedicated_nasha_partition_snapshot": {
				Tok: ovhResource(dedicatedMod, "NasHAPartitionSnapshot"),
			},
			"ovh_dedicated_server": {
				Tok:       ovhResource(dedicatedMod, "Server"),
				ComputeID: delegateID("display_name"),
			},
			"ovh_dedicated_server_reinstall_task": {
				Tok: ovhResource(dedicatedMod, "ServerReinstallTask"),
			},
			"ovh_dedicated_server_networking": {
				Tok: ovhResource(dedicatedMod, "ServerNetworking"),
			},
			"ovh_dedicated_server_reboot_task": {
				Tok: ovhResource(dedicatedMod, "ServerRebootTask"),
			},
			"ovh_dedicated_server_update": {
				Tok: ovhResource(dedicatedMod, "ServerUpdate"),
			},
			"ovh_domain_ds_records": {
				Tok: ovhResource(domainMod, "DSRecords"),
			},
			"ovh_domain_name": {
				Tok: ovhResource(domainMod, "Name"),
			},
			"ovh_domain_name_servers": {
				Tok: ovhResource(domainMod, "NameServers"),
			},
			"ovh_domain_zone": {
				Tok: ovhResource(domainMod, "Zone"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "ZoneURN",
					},
				},
			},
			"ovh_domain_zone_dnssec": {
				Tok:       ovhResource(domainMod, "ZoneDNSSec"),
				ComputeID: delegateID("zone_name"),
			},
			"ovh_domain_zone_import": {
				Tok:       ovhResource(domainMod, "ZoneImport"),
				ComputeID: delegateID("zone_name"),
			},
			"ovh_domain_zone_record": {
				Tok: ovhResource(domainMod, "ZoneRecord"),
			},
			"ovh_domain_zone_redirection": {
				Tok: ovhResource(domainMod, "ZoneRedirection"),
			},
			"ovh_hosting_privatedatabase": {
				Tok: ovhResource(hostingMod, "PrivateDatabase"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "DatabaseURN",
					},
				},
			},
			"ovh_hosting_privatedatabase_database": {
				Tok: ovhResource(hostingMod, "PrivateDatabaseDb"),
			},
			"ovh_hosting_privatedatabase_user": {
				Tok: ovhResource(hostingMod, "PrivateDatabaseUser"),
			},
			"ovh_hosting_privatedatabase_user_grant": {
				Tok: ovhResource(hostingMod, "PrivateDatabaseUserGrant"),
			},
			"ovh_hosting_privatedatabase_whitelist": {
				Tok: ovhResource(hostingMod, "PrivateDatabaseAllowlist"),
			},
			"ovh_ip_firewall": {
				Tok:       ovhResource(ipMod, "Firewall"),
				ComputeID: delegateID("ip"),
			},
			"ovh_ip_firewall_rule": {
				Tok:       ovhResource(ipMod, "FirewallRule"),
				ComputeID: delegateID("ip"),
			},
			"ovh_ip_mitigation": {
				Tok:       ovhResource(ipMod, "Mitigation"),
				ComputeID: delegateID("ip"),
			},
			"ovh_ip_move": {
				Tok: ovhResource(ipMod, "Move"),
			},
			"ovh_ip_reverse": {
				Tok: ovhResource(ipMod, "Reverse"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ip_reverse": {
						Name: "ReverseIp",
					},
					"reverse": {
						Name: "ReverseValue",
					},
				},
			},
			"ovh_ip_service": {
				Tok: ovhResource(ipMod, "IpService"),
			},
			"ovh_iploadbalancing": {
				Tok: ovhResource(ipLoadBalancingMod, "LoadBalancer"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "LoadBalancerURN",
					},
				},
			},
			"ovh_iploadbalancing_http_farm": {
				Tok: ovhResource(ipLoadBalancingMod, "HttpFarm"),
			},
			"ovh_iploadbalancing_http_farm_server": {
				Tok: ovhResource(ipLoadBalancingMod, "HttpFarmServer"),
			},
			"ovh_iploadbalancing_http_frontend": {
				Tok: ovhResource(ipLoadBalancingMod, "HttpFrontend"),
			},
			"ovh_iploadbalancing_http_route": {
				Tok: ovhResource(ipLoadBalancingMod, "HttpRoute"),
			},
			"ovh_iploadbalancing_http_route_rule": {
				Tok: ovhResource(ipLoadBalancingMod, "HttpRouteRule"),
			},
			"ovh_iploadbalancing_refresh": {
				Tok: ovhResource(ipLoadBalancingMod, "Refresh"),
			},
			"ovh_iploadbalancing_ssl": {
				Tok:       ovhResource(ipLoadBalancingMod, "Ssl"),
				ComputeID: delegateID("display_name"),
			},
			"ovh_iploadbalancing_tcp_farm": {
				Tok: ovhResource(ipLoadBalancingMod, "TcpFarm"),
			},
			"ovh_iploadbalancing_tcp_farm_server": {
				Tok: ovhResource(ipLoadBalancingMod, "TcpFarmServer"),
			},
			"ovh_iploadbalancing_tcp_frontend": {
				Tok: ovhResource(ipLoadBalancingMod, "TcpFrontend"),
			},
			"ovh_iploadbalancing_tcp_route": {
				Tok: ovhResource(ipLoadBalancingMod, "TcpRoute"),
			},
			"ovh_iploadbalancing_tcp_route_rule": {
				Tok: ovhResource(ipLoadBalancingMod, "TcpRouteRule"),
			},
			"ovh_iploadbalancing_udp_farm": {
				Tok:       ovhResource(ipLoadBalancingMod, "UdpFarm"),
				ComputeID: delegateID("display_name"),
			},
			"ovh_iploadbalancing_udp_farm_server": {
				Tok:       ovhResource(ipLoadBalancingMod, "UdpFarmServer"),
				ComputeID: delegateID("display_name"),
			},
			"ovh_iploadbalancing_udp_frontend": {
				Tok:       ovhResource(ipLoadBalancingMod, "UdpFrontend"),
				ComputeID: delegateID("display_name"),
			},
			"ovh_iploadbalancing_vrack_network": {
				Tok: ovhResource(ipLoadBalancingMod, "VrackNetwork"),
			},
			"ovh_me_identity_user": {
				Tok: ovhResource(meMod, "IdentityUser"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "UserURN",
					},
				},
			},
			"ovh_okms": {
				Tok: ovhResource(okmsMod, "Okms"),
			},
			"ovh_okms_credential": {
				Tok: ovhResource(okmsMod, "Credential"),
			},
			"ovh_okms_service_key": {
				Tok: ovhResource(okmsMod, "ServiceKey"),
			},
			"ovh_okms_service_key_jwk": {
				Tok: ovhResource(okmsMod, "ServiceKeyJWK"),
			},
			"ovh_savings_plan": {
				Tok: ovhResource(savingPlansMod, "SavingsPlan"),
			},
			"ovh_vrack": {
				Tok: ovhResource(vrackMod, "Vrack"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "VrackURN",
					},
				},
			},
			"ovh_vrack_cloudproject": {
				Tok: ovhResource(vrackMod, "CloudProject"),
			},
			"ovh_vrack_dedicated_cloud": {
				Tok:       ovhResource(vrackMod, "VrackDedicatedCloud"), // VrackDedicatedCloud instead of DedicatedCloud fix until we don't have to create a ComputeID based on a dedicated_cloud field (error in dotnet sdk generation)
				ComputeID: delegateID("dedicated_cloud"),
			},
			"ovh_vrack_dedicated_server": {
				Tok: ovhResource(vrackMod, "DedicatedServer"),
			},
			"ovh_vrack_dedicated_server_interface": {
				Tok: ovhResource(vrackMod, "DedicatedServerInterface"),
			},
			"ovh_vrack_ip": {
				Tok: ovhResource(vrackMod, "IpAddress"),
			},
			"ovh_vrack_iploadbalancing": {
				Tok: ovhResource(vrackMod, "IpLoadbalancing"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"ip_loadbalancing": {
						Name: "LoadbalancingId",
					},
				},
			},
			"ovh_vrack_ipv6": {
				Tok: ovhResource(vrackMod, "IpV6"),
			},
			"ovh_vrack_ovhcloudconnect": {
				Tok:       ovhResource(vrackMod, "OVHcloudConnect"),
				ComputeID: delegateID("ovh_cloud_connect"),
			},
			"ovh_vrack_vrackservices": {
				Tok: ovhResource(vrackMod, "Vrackservices"), // "Vrackservices" Mandatory name to avoid CSharp issue: /workspaces/pulumi-ovh/sdk/dotnet/Vrack/VrackServices.cs(57,31): error CS0542: 'VrackServices': member names cannot be the same as their enclosing type [/workspaces/pulumi-ovh/sdk/dotnet/Pulumi.Ovh.csproj]
			},
			"ovh_cloud_project_workflow_backup": {
				Tok: ovhResource(cloudProjectMod, "WorkflowBackup"),
			},
			"ovh_dbaas_logs_cluster": {
				Tok: ovhResource(dbaasMod, "LogsCluster"),
			},
			"ovh_dbaas_logs_token": {
				Tok:       ovhResource(dbaasMod, "LogsToken"),
				ComputeID: delegateID("name"),
			},
			"ovh_iam_permissions_group": {
				Tok: ovhResource(iamMod, "PermissionsGroup"),
			},
			"ovh_iam_policy": {
				Tok: ovhResource(iamMod, "Policy"),
			},
			"ovh_iam_resource_group": {
				Tok: ovhResource(iamMod, "ResourceGroup"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "GroupURN",
					},
				},
			},
			"ovh_me_api_oauth2_client": {
				Tok: ovhResource(meMod, "APIOAuth2Client"),
			},
			"ovh_me_identity_group": {
				Tok: ovhResource(meMod, "IdentityGroup"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "GroupURN",
					},
				},
			},
			"ovh_vps": {
				Tok:       ovhResource(vpsMod, "Vps"),
				ComputeID: delegateID("display_name"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"ovh_cloud_project": {
				Tok: ovhDataSource(cloudMod, "getProject"),
			},
			"ovh_cloud_project_capabilities_containerregistry": {
				Tok: ovhDataSource(cloudProjectMod, "getCapabilitiesContainerRegistry"),
			},
			"ovh_cloud_project_capabilities_containerregistry_filter": {
				Tok: ovhDataSource(cloudProjectMod, "getCapabilitiesContainerFilter"),
			},
			"ovh_cloud_project_containerregistries": {
				Tok: ovhDataSource(cloudProjectMod, "getContainerRegistries"),
			},
			"ovh_cloud_project_containerregistry": {
				Tok: ovhDataSource(cloudProjectMod, "getContainerRegistry"),
			},
			"ovh_cloud_project_containerregistry_users": {
				Tok: ovhDataSource(cloudProjectMod, "getContainerRegistryUsers"),
			},
			"ovh_cloud_project_containerregistry_oidc": {
				Tok: ovhDataSource(cloudProjectMod, "getContainerRegistryOIDC"),
			},
			"ovh_cloud_project_containerregistry_ip_restrictions_management": {
				Tok: ovhDataSource(cloudProjectMod, "getContainerRegistryIPRestrictionsManagement"),
			},
			"ovh_cloud_project_containerregistry_ip_restrictions_registry": {
				Tok: ovhDataSource(cloudProjectMod, "getContainerRegistryIPRestrictionsRegistry"),
			},
			"ovh_cloud_project_database": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabase"),
			},
			"ovh_cloud_project_database_capabilities": {
				Tok: ovhDataSource(cloudProjectDbMod, "getCapabilities"),
			},
			"ovh_cloud_project_database_database": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabaseInstance"),
			},
			"ovh_cloud_project_database_databases": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabaseInstances"),
			},
			"ovh_cloud_project_database_integration": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabaseIntegration"),
			},
			"ovh_cloud_project_database_integrations": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabaseIntegrations"),
			},
			"ovh_cloud_project_database_log_subscription": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabaseLogSubscription"),
			},
			"ovh_cloud_project_database_log_subscriptions": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabaseLogSubscriptions"),
			},
			"ovh_cloud_project_database_postgresql_connection_pools": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabasePostgreSQLConnectionPools"),
			},
			"ovh_cloud_project_database_certificates": {
				Tok: ovhDataSource(cloudProjectDbMod, "getCertificates"),
			},
			"ovh_cloud_project_database_ip_restrictions": {
				Tok: ovhDataSource(cloudProjectDbMod, "getIpRestrictions"),
			},
			"ovh_cloud_project_database_kafka_acl": {
				Tok: ovhDataSource(cloudProjectDbMod, "getKafkaAcl"),
			},
			"ovh_cloud_project_database_kafka_acls": {
				Tok: ovhDataSource(cloudProjectDbMod, "getKafkaAcls"),
			},
			"ovh_cloud_project_database_kafka_schemaregistryacl": {
				Tok: ovhDataSource(cloudProjectDbMod, "getKafkaSchemaRegistryAcl"),
			},
			"ovh_cloud_project_database_kafka_schemaregistryacls": {
				Tok: ovhDataSource(cloudProjectDbMod, "getKafkaSchemaRegistryAcls"),
			},
			"ovh_cloud_project_database_kafka_topic": {
				Tok: ovhDataSource(cloudProjectDbMod, "getKafkaTopic"),
			},
			"ovh_cloud_project_database_kafka_topics": {
				Tok: ovhDataSource(cloudProjectDbMod, "getKafkaTopics"),
			},
			"ovh_cloud_project_database_kafka_user_access": {
				Tok: ovhDataSource(cloudProjectDbMod, "getKafkaUserAccess"),
			},
			"ovh_cloud_project_database_postgresql_user": {
				Tok: ovhDataSource(cloudProjectDbMod, "getPostgresSqlUser"),
			},
			"ovh_cloud_project_database_postgresql_connection_pool": {
				Tok: ovhDataSource(cloudProjectDbMod, "getPostgresSqlConnectionPool"),
			},
			"ovh_cloud_project_database_user": {
				Tok: ovhDataSource(cloudProjectDbMod, "getUser"),
			},
			"ovh_cloud_project_database_users": {
				Tok: ovhDataSource(cloudProjectDbMod, "getUsers"),
			},
			"ovh_cloud_project_databases": {
				Tok: ovhDataSource(cloudProjectDbMod, "getDatabases"),
			},
			"ovh_cloud_project_failover_ip_attach": {
				Tok: ovhDataSource(cloudProjectMod, "getFailoverIpAttach"),
			},
			"ovh_cloud_project_flavors": {
				Tok: ovhDataSource(cloudProjectMod, "getFlavors"),
			},
			"ovh_cloud_project_floatingips": {
				Tok: ovhDataSource(cloudProjectMod, "getFloatingIPs"),
			},
			"ovh_cloud_project_gateway_interface": {
				Tok: ovhDataSource(cloudProjectMod, "getGatewayInterface"),
			},
			"ovh_cloud_project_image": {
				Tok: ovhDataSource(cloudProjectMod, "getImage"),
			},
			"ovh_cloud_project_images": {
				Tok: ovhDataSource(cloudProjectMod, "getImages"),
			},
			"ovh_cloud_project_kube": {
				Tok: ovhDataSource(cloudProjectMod, "getKube"),
			},
			"ovh_cloud_project_kube_iprestrictions": {
				Tok: ovhDataSource(cloudProjectMod, "getKubeIpRestrictions"),
			},
			"ovh_cloud_project_kube_nodepool": {
				Tok: ovhDataSource(cloudProjectMod, "getKubeNodePool"),
			},
			"ovh_cloud_project_kube_nodepool_nodes": {
				Tok: ovhDataSource(cloudProjectMod, "getKubeNodePoolNodes"),
			},
			"ovh_cloud_project_kube_nodes": {
				Tok: ovhDataSource(cloudProjectMod, "getKubeNodes"),
			},
			"ovh_cloud_project_kube_oidc": {
				Tok: ovhDataSource(cloudProjectMod, "getKubeOidc"),
			},
			"ovh_cloud_project_database_m3db_namespace": {
				Tok: ovhDataSource(cloudProjectMod, "getM3dbNamespace"),
			},
			"ovh_cloud_project_database_m3db_namespaces": {
				Tok: ovhDataSource(cloudProjectMod, "getM3dbNamespaces"),
			},
			"ovh_cloud_project_database_m3db_user": {
				Tok: ovhDataSource(cloudProjectMod, "getM3dbUser"),
			},
			"ovh_cloud_project_database_mongodb_user": {
				Tok: ovhDataSource(cloudProjectMod, "getMongoDbUser"),
			},
			"ovh_cloud_project_database_mongodb_prometheus": {
				Tok: ovhDataSource(cloudProjectMod, "getMongoDbPrometheus"),
			},
			"ovh_cloud_project_database_prometheus": {
				Tok: ovhDataSource(cloudProjectMod, "getPrometheus"),
			},
			"ovh_cloud_project_database_opensearch_pattern": {
				Tok: ovhDataSource(cloudProjectMod, "getOpenSearchPattern"),
			},
			"ovh_cloud_project_database_opensearch_patterns": {
				Tok: ovhDataSource(cloudProjectMod, "getOpenSearchPatterns"),
			},
			"ovh_cloud_project_database_opensearch_user": {
				Tok: ovhDataSource(cloudProjectMod, "getOpenSearchUser"),
			},
			"ovh_cloud_project_database_redis_user": {
				Tok: ovhDataSource(cloudProjectMod, "getRedisUser"),
			},
			"ovh_cloud_project_loadbalancer": {
				Tok: ovhDataSource(cloudProjectMod, "getLoadBalancer"),
			},
			"ovh_cloud_project_instance": {
				Tok: ovhDataSource(cloudProjectMod, "getInstance"),
			},
			"ovh_cloud_project_instances": {
				Tok: ovhDataSource(cloudProjectMod, "getInstances"),
			},
			"ovh_cloud_project_loadbalancers": {
				Tok: ovhDataSource(cloudProjectMod, "getLoadBalancers"),
			},
			"ovh_cloud_project_loadbalancer_flavors": {
				Tok: ovhDataSource(cloudProjectMod, "getLoadBalancerFlavors"),
			},
			"ovh_cloud_project_rancher": {
				Tok: ovhDataSource(cloudProjectMod, "getRancher"),
			},
			"ovh_cloud_project_rancher_plan": {
				Tok: ovhDataSource(cloudProjectMod, "getRancherPlan"),
			},
			"ovh_cloud_project_rancher_version": {
				Tok: ovhDataSource(cloudProjectMod, "getRancherVersion"),
			},
			"ovh_cloud_project_region": {
				Tok: ovhDataSource(cloudProjectMod, "getRegion"),
			},
			"ovh_cloud_project_network_private": {
				Tok: ovhDataSource(cloudProjectMod, "getNetworkPrivate"),
			},
			"ovh_cloud_project_network_privates": {
				Tok: ovhDataSource(cloudProjectMod, "getNetworkPrivates"),
			},
			"ovh_cloud_project_network_private_subnets": {
				Tok: ovhDataSource(cloudProjectMod, "getNetworkPrivateSubnets"),
			},
			"ovh_cloud_project_regions": {
				Tok: ovhDataSource(cloudProjectMod, "getRegions"),
			},
			"ovh_cloud_project_region_loadbalancer_log_subscription": {
				Tok: ovhDataSource(cloudProjectMod, "getRegionLoadBalancerLogSubscription"),
			},
			"ovh_cloud_project_region_loadbalancer_log_subscriptions": {
				Tok: ovhDataSource(cloudProjectMod, "getRegionLoadBalancerLogSubscriptions"),
			},
			"ovh_cloud_project_storage": {
				Tok: ovhDataSource(cloudProjectMod, "getStorage"),
			},
			"ovh_cloud_project_storages": {
				Tok: ovhDataSource(cloudProjectMod, "getStorages"),
			},
			"ovh_cloud_project_storage_object": {
				Tok: ovhDataSource(cloudProjectMod, "getStorageObject"),
			},
			"ovh_cloud_project_storage_objects": {
				Tok: ovhDataSource(cloudProjectMod, "getStorageObjects"),
			},
			"ovh_cloud_project_user": {
				Tok: ovhDataSource(cloudProjectMod, "getUser"),
			},
			"ovh_cloud_project_user_s3_credential": {
				Tok: ovhDataSource(cloudProjectMod, "getUserS3Credential"),
			},
			"ovh_cloud_project_user_s3_credentials": {
				Tok: ovhDataSource(cloudProjectMod, "getUserS3Credentials"),
			},
			"ovh_cloud_project_user_s3_policy": {
				Tok: ovhDataSource(cloudProjectMod, "getUserS3Policy"),
			},
			"ovh_cloud_project_users": {
				Tok: ovhDataSource(cloudProjectMod, "getUsers"),
			},
			"ovh_cloud_project_vrack": {
				Tok: ovhDataSource(cloudProjectMod, "getVRack"),
			},
			"ovh_cloud_project_volume": {
				Tok: ovhDataSource(cloudProjectMod, "getVolume"),
			},
			"ovh_cloud_project_volumes": {
				Tok: ovhDataSource(cloudProjectMod, "getVolumes"),
			},
			"ovh_cloud_projects": {
				Tok: ovhDataSource(cloudMod, "getProjects"),
			},
			"ovh_dbaas_logs_clusters": {
				Tok: ovhDataSource(dbaasMod, "getLogsClusters"),
			},
			"ovh_dbaas_logs_cluster_retention": {
				Tok: ovhDataSource(dbaasMod, "getLogsClustersRetention"),
			},
			"ovh_dbaas_logs_input_engine": {
				Tok: ovhDataSource(dbaasMod, "getLogsInputEngine"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: ovhDataSource(dbaasMod, "getLogsOutputGraylogStream"),
			},
			"ovh_dbaas_logs_output_opensearch_index": {
				Tok: ovhDataSource(dbaasMod, "getLogsOutputOpenSearchIndex"),
			},
			"ovh_dedicated_ceph": {
				Tok: ovhDataSource(dedicatedMod, "getCeph"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "CephURN",
					},
				},
			},
			"ovh_dedicated_cloud": {
				Tok: ovhDataSource(dedicatedMod, "getCloud"),
			},
			"ovh_dedicated_nasha": {
				Tok: ovhDataSource(dedicatedMod, "getNasHA"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "NasHAURN",
					},
				},
			},
			"ovh_dedicated_nasha_partition": {
				Tok: ovhDataSource(dedicatedMod, "getNasHAPartition"),
			},
			"ovh_dedicated_server_boots": {
				Tok: ovhDataSource(dedicatedMod, "getServerBoots"),
			},
			"ovh_dedicated_server_specifications_hardware": {
				Tok: ovhDataSource(dedicatedMod, "getServerSpecificationsHardware"),
			},
			"ovh_dedicated_server_specifications_network": {
				Tok: ovhDataSource(dedicatedMod, "getServerSpecificationsNetwork"),
			},
			"ovh_hosting_privatedatabase": {
				Tok: ovhDataSource(hostingMod, "getPrivateDatabase"),
			},
			"ovh_hosting_privatedatabase_database": {
				Tok: ovhDataSource(hostingMod, "getPrivateDatabaseDb"),
			},
			"ovh_hosting_privatedatabase_user": {
				Tok: ovhDataSource(hostingMod, "getPrivateDatabaseUser"),
			},
			"ovh_hosting_privatedatabase_user_grant": {
				Tok: ovhDataSource(hostingMod, "getPrivateDatabaseUserGrant"),
			},
			"ovh_hosting_privatedatabase_whitelist": {
				Tok: ovhDataSource(hostingMod, "getPrivateDatabaseAllowlist"),
			},
			"ovh_dedicated_installation_template": {
				Tok: ovhDataSource(ovhMod, "getInstallationTemplate"),
			},
			"ovh_dedicated_installation_templates": {
				Tok: ovhDataSource(ovhMod, "getInstallationTemplates"),
			},
			"ovh_dedicated_server": {
				Tok: ovhDataSource(ovhMod, "getServer"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "ServerURN",
					},
				},
			},
			"ovh_dedicated_servers": {
				Tok: ovhDataSource(ovhMod, "getServers"),
			},
			"ovh_domain_zone": {
				Tok: ovhDataSource(domainMod, "getZone"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "ZoneURN",
					},
				},
			},
			"ovh_domain_zone_dnssec": {
				Tok: ovhDataSource(domainMod, "getZoneDNSSec"),
			},
			"ovh_ip_firewall": {
				Tok: ovhDataSource(ipMod, "getFirewall"),
			},
			"ovh_ip_firewall_rule": {
				Tok: ovhDataSource(ipMod, "getFirewallRule"),
			},
			"ovh_ip_mitigation": {
				Tok: ovhDataSource(ipMod, "getMitigation"),
			},
			"ovh_ip_service": {
				Tok: ovhDataSource(ipMod, "getService"),
			},
			"ovh_iploadbalancing_vrack_network": {
				Tok: ovhDataSource(ipLoadBalancingMod, "getVrackNetwork"),
			},
			"ovh_iploadbalancing": {
				Tok: ovhDataSource(ipLoadBalancingMod, "getIpLoadBalancing"),
			},
			"ovh_iploadbalancing_vrack_networks": {
				Tok: ovhDataSource(ovhMod, "getVrackNetworks"),
			},
			"ovh_okms_credential": {
				Tok: ovhDataSource(okmsMod, "getOkmsCredential"),
			},
			"ovh_okms_resource": {
				Tok: ovhDataSource(okmsMod, "getOkmsResource"),
			},
			"ovh_okms_service_key": {
				Tok: ovhDataSource(okmsMod, "getOkmsServiceKey"),
			},
			"ovh_okms_service_key_jwk": {
				Tok: ovhDataSource(okmsMod, "getOkmsServiceKeyJwk"),
			},
			"ovh_okms_service_key_pem": {
				Tok: ovhDataSource(okmsMod, "getOkmsServiceKeyPem"),
			},
			"ovh_ovhcloud_connect": {
				Tok: ovhDataSource(ovhCloudMod, "Connect"),
			},
			"ovh_ovhcloud_connects": {
				Tok: ovhDataSource(ovhCloudMod, "Connects"),
			},
			"ovh_me": {
				Tok: ovhDataSource(meMod, "getMe"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "AccountURN",
					},
				},
			},
			"ovh_me_identity_user": {
				Tok: ovhDataSource(meMod, "getIdentityUser"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "UserURN",
					},
				},
			},
			"ovh_me_identity_users": {
				Tok: ovhDataSource(meMod, "getIdentityUsers"),
			},
			"ovh_me_paymentmean_bankaccount": {
				Tok: ovhDataSource(meMod, "getPaymentmeanBankAccount"),
			},
			"ovh_me_paymentmean_creditcard": {
				Tok: ovhDataSource(meMod, "getPaymentmeanCreditCard"),
			},
			"ovh_order_cart": {
				Tok: ovhDataSource(orderMod, "getCart"),
			},
			"ovh_order_cart_product": {
				Tok: ovhDataSource(orderMod, "getCartProduct"),
			},
			"ovh_order_cart_product_options": {
				Tok: ovhDataSource(orderMod, "getCartProductOptions"),
			},
			"ovh_order_cart_product_options_plan": {
				Tok: ovhDataSource(orderMod, "getCartProductOptionsPlan"),
			},
			"ovh_order_cart_product_plan": {
				Tok: ovhDataSource(orderMod, "getCartProductPlan"),
			},
			"ovh_vps": {
				Tok: ovhDataSource(vpsMod, "getVps"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "VpsURN",
					},
				},
			},
			"ovh_vpss": {
				Tok: ovhDataSource(vpsMod, "getVpss"),
			},
			"ovh_vracks": {
				Tok: ovhDataSource(vrackMod, "getVracks"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "VrackURN",
					},
				},
			},
			"ovh_dbaas_logs_cluster": {
				Tok: ovhDataSource(dbaasMod, "getLogsCluster"),
			},
			"ovh_iam_permissions_group": {
				Tok: ovhDataSource(iamMod, "getPermissionsGroup"),
			},
			"ovh_iam_permissions_groups": {
				Tok: ovhDataSource(iamMod, "getPermissionsGroups"),
			},
			"ovh_iam_policies": {
				Tok: ovhDataSource(iamMod, "getPolicies"),
			},
			"ovh_iam_policy": {
				Tok: ovhDataSource(iamMod, "getPolicy"),
			},
			"ovh_iam_reference_actions": {
				Tok: ovhDataSource(iamMod, "getReferenceActions"),
			},
			"ovh_iam_resource_group": {
				Tok: ovhDataSource(iamMod, "getResourceGroup"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "GroupURN",
					},
				},
			},
			"ovh_iam_resource_groups": {
				Tok: ovhDataSource(iamMod, "getResourceGroups"),
			},
			"ovh_me_api_oauth2_client": {
				Tok: ovhDataSource(meMod, "getAPIOAuth2Client"),
			},
			"ovh_me_api_oauth2_clients": {
				Tok: ovhDataSource(meMod, "getAPIOAuth2Clients"),
			},
			"ovh_me_identity_group": {
				Tok: ovhDataSource(meMod, "getIdentityGroup"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "GroupURN",
					},
				},
			},
			"ovh_me_identity_groups": {
				Tok: ovhDataSource(meMod, "getIdentityGroups"),
			},
			"ovh_iam_reference_resource_type": {
				Tok: ovhDataSource(iamMod, "getReferenceResourceType"),
			},
			"ovh_vmware_cloud_director_backup": {
				Tok: ovhDataSource(vmwareMod, "getCloudDirectorBackup"),
			},
			"ovh_vmware_cloud_director_organization": {
				Tok: ovhDataSource(vmwareMod, "getCloudDirectorOrganization"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@ovhcloud/pulumi-ovh",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumi_ovh",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject: struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/ovh/pulumi-%[1]s/sdk/", ovhPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				ovhPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumi",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.ovhcloud.pulumi",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
