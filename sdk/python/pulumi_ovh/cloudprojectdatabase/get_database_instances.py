# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'GetDatabaseInstancesResult',
    'AwaitableGetDatabaseInstancesResult',
    'get_database_instances',
    'get_database_instances_output',
]

@pulumi.output_type
class GetDatabaseInstancesResult:
    """
    A collection of values returned by getDatabaseInstances.
    """
    def __init__(__self__, cluster_id=None, database_ids=None, engine=None, id=None, service_name=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if database_ids and not isinstance(database_ids, list):
            raise TypeError("Expected argument 'database_ids' to be a list")
        pulumi.set(__self__, "database_ids", database_ids)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="databaseIds")
    def database_ids(self) -> Sequence[builtins.str]:
        """
        The list of databases ids of the database cluster associated with the project.
        """
        return pulumi.get(self, "database_ids")

    @property
    @pulumi.getter
    def engine(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetDatabaseInstancesResult(GetDatabaseInstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseInstancesResult(
            cluster_id=self.cluster_id,
            database_ids=self.database_ids,
            engine=self.engine,
            id=self.id,
            service_name=self.service_name)


def get_database_instances(cluster_id: Optional[builtins.str] = None,
                           engine: Optional[builtins.str] = None,
                           service_name: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseInstancesResult:
    """
    Use this data source to get the list of databases of a database cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    databases = ovh.CloudProjectDatabase.get_database_instances(service_name="XXXX",
        engine="YYYY",
        cluster_id="ZZZ")
    pulumi.export("databaseIds", databases.database_ids)
    ```


    :param builtins.str cluster_id: Cluster ID
    :param builtins.str engine: The engine of the database cluster you want to list databases. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getDatabaseInstances:getDatabaseInstances', __args__, opts=opts, typ=GetDatabaseInstancesResult).value

    return AwaitableGetDatabaseInstancesResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        database_ids=pulumi.get(__ret__, 'database_ids'),
        engine=pulumi.get(__ret__, 'engine'),
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_database_instances_output(cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                                  engine: Optional[pulumi.Input[builtins.str]] = None,
                                  service_name: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDatabaseInstancesResult]:
    """
    Use this data source to get the list of databases of a database cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    databases = ovh.CloudProjectDatabase.get_database_instances(service_name="XXXX",
        engine="YYYY",
        cluster_id="ZZZ")
    pulumi.export("databaseIds", databases.database_ids)
    ```


    :param builtins.str cluster_id: Cluster ID
    :param builtins.str engine: The engine of the database cluster you want to list databases. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProjectDatabase/getDatabaseInstances:getDatabaseInstances', __args__, opts=opts, typ=GetDatabaseInstancesResult)
    return __ret__.apply(lambda __response__: GetDatabaseInstancesResult(
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        database_ids=pulumi.get(__response__, 'database_ids'),
        engine=pulumi.get(__response__, 'engine'),
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name')))
