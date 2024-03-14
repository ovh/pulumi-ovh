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
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/ovh/pulumi-ovh/provider/pkg/version"
	"github.com/ovh/terraform-provider-ovh/ovh"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	ovhPkg = "ovh"
	// modules:
	ovhMod             = "index"
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
		GitHubOrg: "ovh",
		Version:   version.Version,
		Config: map[string]*tfbridge.SchemaInfo{
			"endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OVH_ENDPOINT"},
				},
			},
			"application_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OVH_APPLICATION_KEY"},
				},
			},
			"application_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OVH_APPLICATION_SECRET"},
				},
				Secret: tfbridge.True(),
			},
			"consumer_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OVH_CONSUMER_KEY"},
				},
			},
		},
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
			"ovh_cloud_project_database_m3db_namespace": {
				Tok: ovhResource(cloudProjectDbMod, "M3DbNamespace"),
			},
			"ovh_cloud_project_database_m3db_user": {
				Tok: ovhResource(cloudProjectDbMod, "M3DbUser"),
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
			"ovh_cloud_project_network_private": {
				Tok: ovhResource(cloudProjectMod, "NetworkPrivate"),
			},
			"ovh_cloud_project_network_private_subnet": {
				Tok: ovhResource(cloudProjectMod, "NetworkPrivateSubnet"),
			},
			"ovh_cloud_project_region_storage_presign": {
				Tok: ovhResource(cloudProjectMod, "RegionStoragePresign"),
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
			"ovh_dbaas_logs_input": {
				Tok: ovhResource(dbaasMod, "LogsInput"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: ovhResource(dbaasMod, "LogsOutputGraylogStream"),
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
			"ovh_dedicated_server_install_task": {
				Tok: ovhResource(dedicatedMod, "ServerInstallTask"),
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
			"ovh_domain_zone": {
				Tok: ovhResource(domainMod, "Zone"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "ZoneURN",
					},
				},
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
			"ovh_me_installation_template": {
				Tok: ovhResource(meMod, "InstallationTemplate"),
			},
			"ovh_me_installation_template_partition_scheme": {
				Tok: ovhResource(meMod, "InstallationTemplatePartitionScheme"),
			},
			"ovh_me_installation_template_partition_scheme_hardware_raid": {
				Tok: ovhResource(meMod, "InstallationTemplatePartitionSchemeHardwareRaid"),
			},
			"ovh_me_installation_template_partition_scheme_partition": {
				Tok: ovhResource(meMod, "InstallationTemplatePartitionSchemePartition"),
			},
			"ovh_me_ssh_key": {
				Tok: ovhResource(meMod, "SshKey"),
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
			"ovh_cloud_project_workflow_backup": {
				Tok: ovhResource(cloudProjectMod, "WorkflowBackup"),
			},
			"ovh_dbaas_logs_cluster": {
				Tok: ovhResource(dbaasMod, "LogsCluster"),
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
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
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
			"ovh_cloud_project_region": {
				Tok: ovhDataSource(cloudProjectMod, "getRegion"),
			},
			"ovh_cloud_project_regions": {
				Tok: ovhDataSource(cloudProjectMod, "getRegions"),
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
			"ovh_dbaas_logs_clusters": {
				Tok: ovhDataSource(dbaasMod, "getLogsClusters"),
			},
			"ovh_dbaas_logs_input_engine": {
				Tok: ovhDataSource(dbaasMod, "getLogsInputEngine"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: ovhDataSource(dbaasMod, "getLogsOutputGraylogStream"),
			},
			"ovh_dedicated_ceph": {
				Tok: ovhDataSource(dedicatedMod, "getCeph"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "CephURN",
					},
				},
			},
			"ovh_dedicated_nasha": {
				Tok: ovhDataSource(dedicatedMod, "getNasHA"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"urn": {
						Name: "NasHAURN",
					},
				},
			},
			"ovh_dedicated_server_boots": {
				Tok: ovhDataSource(dedicatedMod, "getServerBoots"),
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
			"ovh_me_installation_template": {
				Tok: ovhDataSource(meMod, "getInstallationTemplate"),
			},
			"ovh_me_installation_templates": {
				Tok: ovhDataSource(meMod, "getInstallationTemplates"),
			},
			"ovh_me_paymentmean_bankaccount": {
				Tok: ovhDataSource(meMod, "getPaymentmeanBankAccount"),
			},
			"ovh_me_paymentmean_creditcard": {
				Tok: ovhDataSource(meMod, "getPaymentmeanCreditCard"),
			},
			"ovh_me_ssh_key": {
				Tok: ovhDataSource(meMod, "getSshKey"),
			},
			"ovh_me_ssh_keys": {
				Tok: ovhDataSource(meMod, "getSshKeys"),
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
			BasePackage: "com.ovh",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
