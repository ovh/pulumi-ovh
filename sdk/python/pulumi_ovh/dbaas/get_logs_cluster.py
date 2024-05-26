# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetLogsClusterResult',
    'AwaitableGetLogsClusterResult',
    'get_logs_cluster',
    'get_logs_cluster_output',
]

@pulumi.output_type
class GetLogsClusterResult:
    """
    A collection of values returned by getLogsCluster.
    """
    def __init__(__self__, archive_allowed_networks=None, cluster_id=None, cluster_type=None, dedicated_input_pem=None, direct_input_allowed_networks=None, direct_input_pem=None, hostname=None, id=None, is_default=None, is_unlocked=None, query_allowed_networks=None, region=None, service_name=None, urn=None):
        if archive_allowed_networks and not isinstance(archive_allowed_networks, list):
            raise TypeError("Expected argument 'archive_allowed_networks' to be a list")
        pulumi.set(__self__, "archive_allowed_networks", archive_allowed_networks)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_type and not isinstance(cluster_type, str):
            raise TypeError("Expected argument 'cluster_type' to be a str")
        pulumi.set(__self__, "cluster_type", cluster_type)
        if dedicated_input_pem and not isinstance(dedicated_input_pem, str):
            raise TypeError("Expected argument 'dedicated_input_pem' to be a str")
        pulumi.set(__self__, "dedicated_input_pem", dedicated_input_pem)
        if direct_input_allowed_networks and not isinstance(direct_input_allowed_networks, list):
            raise TypeError("Expected argument 'direct_input_allowed_networks' to be a list")
        pulumi.set(__self__, "direct_input_allowed_networks", direct_input_allowed_networks)
        if direct_input_pem and not isinstance(direct_input_pem, str):
            raise TypeError("Expected argument 'direct_input_pem' to be a str")
        pulumi.set(__self__, "direct_input_pem", direct_input_pem)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_default and not isinstance(is_default, bool):
            raise TypeError("Expected argument 'is_default' to be a bool")
        pulumi.set(__self__, "is_default", is_default)
        if is_unlocked and not isinstance(is_unlocked, bool):
            raise TypeError("Expected argument 'is_unlocked' to be a bool")
        pulumi.set(__self__, "is_unlocked", is_unlocked)
        if query_allowed_networks and not isinstance(query_allowed_networks, list):
            raise TypeError("Expected argument 'query_allowed_networks' to be a list")
        pulumi.set(__self__, "query_allowed_networks", query_allowed_networks)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if urn and not isinstance(urn, str):
            raise TypeError("Expected argument 'urn' to be a str")
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="archiveAllowedNetworks")
    def archive_allowed_networks(self) -> Sequence[str]:
        """
        is allowed networks for ARCHIVE flow type
        """
        return pulumi.get(self, "archive_allowed_networks")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> str:
        """
        is type of cluster (DEDICATED, PRO or TRIAL)
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="dedicatedInputPem")
    def dedicated_input_pem(self) -> str:
        """
        is PEM for dedicated inputs
        """
        return pulumi.get(self, "dedicated_input_pem")

    @property
    @pulumi.getter(name="directInputAllowedNetworks")
    def direct_input_allowed_networks(self) -> Sequence[str]:
        """
        is allowed networks for DIRECT_INPUT flow type
        """
        return pulumi.get(self, "direct_input_allowed_networks")

    @property
    @pulumi.getter(name="directInputPem")
    def direct_input_pem(self) -> str:
        """
        is PEM for direct inputs
        """
        return pulumi.get(self, "direct_input_pem")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        is cluster hostname hosting the tenant
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> bool:
        """
        is true if all content generated by given service will be placed on this cluster
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter(name="isUnlocked")
    def is_unlocked(self) -> bool:
        """
        is true if given service can perform advanced operations on cluster
        """
        return pulumi.get(self, "is_unlocked")

    @property
    @pulumi.getter(name="queryAllowedNetworks")
    def query_allowed_networks(self) -> Sequence[str]:
        """
        is allowed networks for QUERY flow type
        """
        return pulumi.get(self, "query_allowed_networks")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        is datacenter localization
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        is the URN of the DBaas logs instance
        """
        return pulumi.get(self, "urn")


class AwaitableGetLogsClusterResult(GetLogsClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLogsClusterResult(
            archive_allowed_networks=self.archive_allowed_networks,
            cluster_id=self.cluster_id,
            cluster_type=self.cluster_type,
            dedicated_input_pem=self.dedicated_input_pem,
            direct_input_allowed_networks=self.direct_input_allowed_networks,
            direct_input_pem=self.direct_input_pem,
            hostname=self.hostname,
            id=self.id,
            is_default=self.is_default,
            is_unlocked=self.is_unlocked,
            query_allowed_networks=self.query_allowed_networks,
            region=self.region,
            service_name=self.service_name,
            urn=self.urn)


def get_logs_cluster(cluster_id: Optional[str] = None,
                     service_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLogsClusterResult:
    """
    Use this data source to retrieve informations about a DBaas logs cluster tenant.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    logstash = ovh.Dbaas.get_logs_cluster(cluster_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        service_name="ldp-xx-xxxxx")
    ```


    :param str cluster_id: Cluster ID. If not provided, the default cluster_id is returned
    :param str service_name: The service name. It's the ID of your Logs Data Platform instance.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dbaas/getLogsCluster:getLogsCluster', __args__, opts=opts, typ=GetLogsClusterResult).value

    return AwaitableGetLogsClusterResult(
        archive_allowed_networks=pulumi.get(__ret__, 'archive_allowed_networks'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        cluster_type=pulumi.get(__ret__, 'cluster_type'),
        dedicated_input_pem=pulumi.get(__ret__, 'dedicated_input_pem'),
        direct_input_allowed_networks=pulumi.get(__ret__, 'direct_input_allowed_networks'),
        direct_input_pem=pulumi.get(__ret__, 'direct_input_pem'),
        hostname=pulumi.get(__ret__, 'hostname'),
        id=pulumi.get(__ret__, 'id'),
        is_default=pulumi.get(__ret__, 'is_default'),
        is_unlocked=pulumi.get(__ret__, 'is_unlocked'),
        query_allowed_networks=pulumi.get(__ret__, 'query_allowed_networks'),
        region=pulumi.get(__ret__, 'region'),
        service_name=pulumi.get(__ret__, 'service_name'),
        urn=pulumi.get(__ret__, 'urn'))


@_utilities.lift_output_func(get_logs_cluster)
def get_logs_cluster_output(cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                            service_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLogsClusterResult]:
    """
    Use this data source to retrieve informations about a DBaas logs cluster tenant.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    logstash = ovh.Dbaas.get_logs_cluster(cluster_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        service_name="ldp-xx-xxxxx")
    ```


    :param str cluster_id: Cluster ID. If not provided, the default cluster_id is returned
    :param str service_name: The service name. It's the ID of your Logs Data Platform instance.
    """
    ...
