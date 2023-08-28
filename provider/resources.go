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
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/ovh/terraform-provider-ovh/ovh"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/scraly/pulumi-ovh/provider/pkg/version"
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

// ovhMember manufactures a type token for the Ovh package and the given module and type.
func ovhMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(ovhPkg + ":" + mod + ":" + mem)
}

// ovhType manufactures a type token for the Ovh package and the given module and type.
func ovhType(mod string, typ string) tokens.Type {
	return tokens.Type(ovhMember(mod, typ))
}

// ovhDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Ovh package and names the file by simply lower casing the data
// source's first character.
func ovhDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return ovhMember(mod+"/"+fn, res)
}

// ovhResource manufactures a standard resource token given a module and resource name.
// It automatically uses the ovh package and names the file by simply lower casing the resource's
// first character.
func ovhResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return ovhType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(ovh.Provider())

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
		Publisher: "scraly",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/scraly/pulumi-ovh",
		Description:       "A Pulumi package for creating and managing OVH resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "ovh", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/scraly/pulumi-ovh",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "ovh",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"ovh_cloud_project": {Tok: ovhResource(cloudProjectMod, "Project")},
			"ovh_cloud_project_containerregistry": {
				Tok: ovhResource(cloudProjectMod, "ContainerRegistry"),
			},
			"ovh_cloud_project_containerregistry_user": {
				Tok: ovhResource(cloudProjectMod, "ContainerRegistryUser"),
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
			},
			"ovh_domain_zone_record": {
				Tok: ovhResource(domainMod, "ZoneRecord"),
			},
			"ovh_domain_zone_redirection": {
				Tok: ovhResource(domainMod, "ZoneRedirection"),
			},
			"ovh_hosting_privatedatabase": {
				Tok: ovhResource(hostingMod, "PrivateDatabase"),
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
			"ovh_me_ipxe_script": {
				Tok: ovhResource(meMod, "IpxeScript"),
			},
			"ovh_me_ssh_key": {
				Tok: ovhResource(meMod, "SshKey"),
			},
			"ovh_vrack": {
				Tok: ovhResource(vrackMod, "Vrack"),
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
			"ovh_iam_policy": {
				Tok: ovhResource(iamMod, "Policy"),
			},
			"ovh_me_identity_group": {
				Tok: ovhResource(meMod, "IdentityGroup"),
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
			"ovh_dbaas_logs_input_engine": {
				Tok: ovhDataSource(dbaasMod, "getLogsInputEngine"),
			},
			"ovh_dbaas_logs_output_graylog_stream": {
				Tok: ovhDataSource(dbaasMod, "getLogsOutputGraylogStream"),
			},
			"ovh_dedicated_ceph": {
				Tok: ovhDataSource(dedicatedMod, "getCeph"),
			},
			"ovh_dedicated_nasha": {
				Tok: ovhDataSource(dedicatedMod, "getNasHA"),
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
			},
			"ovh_dedicated_servers": {
				Tok: ovhDataSource(ovhMod, "getServers"),
			},
			"ovh_domain_zone": {
				Tok: ovhDataSource(domainMod, "getZone"),
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
			},
			"ovh_me_identity_user": {
				Tok: ovhDataSource(meMod, "getIdentityUser"),
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
			"ovh_me_ipxe_script": {
				Tok: ovhDataSource(meMod, "getIpxeScript"),
			},
			"ovh_me_ipxe_scripts": {
				Tok: ovhDataSource(meMod, "getIpxeScripts"),
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
			},
			"ovh_vpss": {
				Tok: ovhDataSource(vpsMod, "getVpss"),
			},
			"ovh_vracks": {
				Tok: ovhDataSource(vrackMod, "getVracks"),
			},
			"ovh_dbaas_logs_cluster": {
				Tok: ovhDataSource(dbaasMod, "getLogsCluster"),
			},
			/*
			 "" not mapped to the Pulumi provider
			        * TF data source "" not mapped to the Pulumi provider
			        * TF data source "" not mapped to the Pulumi provider
			        * TF data source "" not mapped to the Pulumi provider
			        * TF data source "" not mapped to the Pulumi provider
			        * TF data source "
			*/
			"ovh_iam_policies": {
				Tok: ovhDataSource(iamMod, "getPolicies"),
			},
			"ovh_iam_policy": {
				Tok: ovhDataSource(iamMod, "getPolicy"),
			},
			"ovh_iam_reference_actions": {
				Tok: ovhDataSource(iamMod, "getReferenceActions"),
			},
			"ovh_iam_reference_resource_types": {
				Tok: ovhDataSource(iamMod, "getReferenceResourceTypes"),
			},
			"ovh_me_identity_group": {
				Tok: ovhDataSource(meMod, "getIdentityGroup"),
			},
			"ovh_me_identity_groups": {
				Tok: ovhDataSource(meMod, "getIdentityGroups"),
			},
			"ovh_iam_reference_resource_type": {
				Tok: ovhDataSource(iamMod, "getReferenceResourceType"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@scraly/pulumi-ovh",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "scraly_pulumi_ovh",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/scraly/pulumi-%[1]s/sdk/", ovhPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				ovhPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Scraly.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.scraly",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
