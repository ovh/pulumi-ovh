# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'GetDatabasePostgreSQLConnectionPoolsResult',
    'AwaitableGetDatabasePostgreSQLConnectionPoolsResult',
    'get_database_postgre_sql_connection_pools',
    'get_database_postgre_sql_connection_pools_output',
]

@pulumi.output_type
class GetDatabasePostgreSQLConnectionPoolsResult:
    """
    A collection of values returned by getDatabasePostgreSQLConnectionPools.
    """
    def __init__(__self__, cluster_id=None, connection_pool_ids=None, id=None, service_name=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if connection_pool_ids and not isinstance(connection_pool_ids, list):
            raise TypeError("Expected argument 'connection_pool_ids' to be a list")
        pulumi.set(__self__, "connection_pool_ids", connection_pool_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="connectionPoolIds")
    def connection_pool_ids(self) -> Sequence[str]:
        """
        The list of patterns ids of the opensearch cluster associated with the project.
        """
        return pulumi.get(self, "connection_pool_ids")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetDatabasePostgreSQLConnectionPoolsResult(GetDatabasePostgreSQLConnectionPoolsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabasePostgreSQLConnectionPoolsResult(
            cluster_id=self.cluster_id,
            connection_pool_ids=self.connection_pool_ids,
            id=self.id,
            service_name=self.service_name)


def get_database_postgre_sql_connection_pools(cluster_id: Optional[str] = None,
                                              service_name: Optional[str] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabasePostgreSQLConnectionPoolsResult:
    """
    Use this data source to get the list of connection pools of a postgresql cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    test_pools = ovh.CloudProjectDatabase.get_database_postgre_sql_connection_pools(service_name="XXX",
        cluster_id="YYY")
    pulumi.export("connectionPoolIds", test_pools.connection_pool_ids)
    ```


    :param str cluster_id: Cluster ID.
    :param str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools', __args__, opts=opts, typ=GetDatabasePostgreSQLConnectionPoolsResult).value

    return AwaitableGetDatabasePostgreSQLConnectionPoolsResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        connection_pool_ids=pulumi.get(__ret__, 'connection_pool_ids'),
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_database_postgre_sql_connection_pools_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                                     service_name: Optional[pulumi.Input[str]] = None,
                                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDatabasePostgreSQLConnectionPoolsResult]:
    """
    Use this data source to get the list of connection pools of a postgresql cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    test_pools = ovh.CloudProjectDatabase.get_database_postgre_sql_connection_pools(service_name="XXX",
        cluster_id="YYY")
    pulumi.export("connectionPoolIds", test_pools.connection_pool_ids)
    ```


    :param str cluster_id: Cluster ID.
    :param str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools', __args__, opts=opts, typ=GetDatabasePostgreSQLConnectionPoolsResult)
    return __ret__.apply(lambda __response__: GetDatabasePostgreSQLConnectionPoolsResult(
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        connection_pool_ids=pulumi.get(__response__, 'connection_pool_ids'),
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name')))
