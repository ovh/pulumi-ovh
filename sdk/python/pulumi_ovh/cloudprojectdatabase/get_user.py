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
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, cluster_id=None, created_at=None, engine=None, id=None, name=None, service_name=None, status=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date of the creation of the user.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Current status of the user.
        """
        return pulumi.get(self, "status")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            cluster_id=self.cluster_id,
            created_at=self.created_at,
            engine=self.engine,
            id=self.id,
            name=self.name,
            service_name=self.service_name,
            status=self.status)


def get_user(cluster_id: Optional[str] = None,
             engine: Optional[str] = None,
             name: Optional[str] = None,
             service_name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Use this data source to get information about a user of a database cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    user = ovh.CloudProjectDatabase.get_user(service_name="XXX",
        engine="YYY",
        cluster_id="ZZZ",
        name="UUU")
    pulumi.export("userName", user.name)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want user information. To get a full list of available engine visit :
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
           Available engines:
    :param str name: Name of the user.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        engine=pulumi.get(__ret__, 'engine'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        service_name=pulumi.get(__ret__, 'service_name'),
        status=pulumi.get(__ret__, 'status'))
def get_user_output(cluster_id: Optional[pulumi.Input[str]] = None,
                    engine: Optional[pulumi.Input[str]] = None,
                    name: Optional[pulumi.Input[str]] = None,
                    service_name: Optional[pulumi.Input[str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserResult]:
    """
    Use this data source to get information about a user of a database cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    user = ovh.CloudProjectDatabase.get_user(service_name="XXX",
        engine="YYY",
        cluster_id="ZZZ",
        name="UUU")
    pulumi.export("userName", user.name)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want user information. To get a full list of available engine visit :
           [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
           Available engines:
    :param str name: Name of the user.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProjectDatabase/getUser:getUser', __args__, opts=opts, typ=GetUserResult)
    return __ret__.apply(lambda __response__: GetUserResult(
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        created_at=pulumi.get(__response__, 'created_at'),
        engine=pulumi.get(__response__, 'engine'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        service_name=pulumi.get(__response__, 'service_name'),
        status=pulumi.get(__response__, 'status')))
