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
    'GetLogsClustersRetentionResult',
    'AwaitableGetLogsClustersRetentionResult',
    'get_logs_clusters_retention',
    'get_logs_clusters_retention_output',
]

@pulumi.output_type
class GetLogsClustersRetentionResult:
    """
    A collection of values returned by getLogsClustersRetention.
    """
    def __init__(__self__, cluster_id=None, duration=None, id=None, is_supported=None, retention_id=None, service_name=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if duration and not isinstance(duration, str):
            raise TypeError("Expected argument 'duration' to be a str")
        pulumi.set(__self__, "duration", duration)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_supported and not isinstance(is_supported, bool):
            raise TypeError("Expected argument 'is_supported' to be a bool")
        pulumi.set(__self__, "is_supported", is_supported)
        if retention_id and not isinstance(retention_id, str):
            raise TypeError("Expected argument 'retention_id' to be a str")
        pulumi.set(__self__, "retention_id", retention_id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def duration(self) -> str:
        """
        Indexed duration expressed in ISO-8601 format
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isSupported")
    def is_supported(self) -> bool:
        """
        Indicates if a new stream can use it
        """
        return pulumi.get(self, "is_supported")

    @property
    @pulumi.getter(name="retentionId")
    def retention_id(self) -> str:
        """
        ID of the retention that can be used when creating a stream
        """
        return pulumi.get(self, "retention_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")


class AwaitableGetLogsClustersRetentionResult(GetLogsClustersRetentionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLogsClustersRetentionResult(
            cluster_id=self.cluster_id,
            duration=self.duration,
            id=self.id,
            is_supported=self.is_supported,
            retention_id=self.retention_id,
            service_name=self.service_name)


def get_logs_clusters_retention(cluster_id: Optional[str] = None,
                                duration: Optional[str] = None,
                                retention_id: Optional[str] = None,
                                service_name: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLogsClustersRetentionResult:
    """
    Use this data source to retrieve information about a DBaas logs cluster retention.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    retention = ovh.Dbaas.get_logs_clusters_retention(cluster_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        retention_id="yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy",
        service_name="ldp-xx-xxxxx")
    ```

    It is also possible to retrieve a retention using its duration:

    ```python
    import pulumi
    import pulumi_ovh as ovh

    retention = ovh.Dbaas.get_logs_clusters_retention(cluster_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        duration="P14D",
        service_name="ldp-xx-xxxxx")
    ```


    :param str cluster_id: Cluster ID
    :param str duration: Indexed duration expressed in ISO-8601 format
    :param str retention_id: ID of the retention object
    :param str service_name: The service name. It's the ID of your Logs Data Platform instance.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['duration'] = duration
    __args__['retentionId'] = retention_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dbaas/getLogsClustersRetention:getLogsClustersRetention', __args__, opts=opts, typ=GetLogsClustersRetentionResult).value

    return AwaitableGetLogsClustersRetentionResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        duration=pulumi.get(__ret__, 'duration'),
        id=pulumi.get(__ret__, 'id'),
        is_supported=pulumi.get(__ret__, 'is_supported'),
        retention_id=pulumi.get(__ret__, 'retention_id'),
        service_name=pulumi.get(__ret__, 'service_name'))


@_utilities.lift_output_func(get_logs_clusters_retention)
def get_logs_clusters_retention_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                       duration: Optional[pulumi.Input[Optional[str]]] = None,
                                       retention_id: Optional[pulumi.Input[Optional[str]]] = None,
                                       service_name: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLogsClustersRetentionResult]:
    """
    Use this data source to retrieve information about a DBaas logs cluster retention.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    retention = ovh.Dbaas.get_logs_clusters_retention(cluster_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        retention_id="yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy",
        service_name="ldp-xx-xxxxx")
    ```

    It is also possible to retrieve a retention using its duration:

    ```python
    import pulumi
    import pulumi_ovh as ovh

    retention = ovh.Dbaas.get_logs_clusters_retention(cluster_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        duration="P14D",
        service_name="ldp-xx-xxxxx")
    ```


    :param str cluster_id: Cluster ID
    :param str duration: Indexed duration expressed in ISO-8601 format
    :param str retention_id: ID of the retention object
    :param str service_name: The service name. It's the ID of your Logs Data Platform instance.
    """
    ...