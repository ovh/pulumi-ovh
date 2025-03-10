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
    'GetDatabaseLogSubscriptionResult',
    'AwaitableGetDatabaseLogSubscriptionResult',
    'get_database_log_subscription',
    'get_database_log_subscription_output',
]

@pulumi.output_type
class GetDatabaseLogSubscriptionResult:
    """
    A collection of values returned by getDatabaseLogSubscription.
    """
    def __init__(__self__, cluster_id=None, created_at=None, engine=None, id=None, kind=None, ldp_service_name=None, resource_name=None, resource_type=None, service_name=None, stream_id=None, updated_at=None):
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
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if ldp_service_name and not isinstance(ldp_service_name, str):
            raise TypeError("Expected argument 'ldp_service_name' to be a str")
        pulumi.set(__self__, "ldp_service_name", ldp_service_name)
        if resource_name and not isinstance(resource_name, str):
            raise TypeError("Expected argument 'resource_name' to be a str")
        pulumi.set(__self__, "resource_name", resource_name)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if stream_id and not isinstance(stream_id, str):
            raise TypeError("Expected argument 'stream_id' to be a str")
        pulumi.set(__self__, "stream_id", stream_id)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def engine(self) -> str:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="ldpServiceName")
    def ldp_service_name(self) -> str:
        return pulumi.get(self, "ldp_service_name")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> str:
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="streamId")
    def stream_id(self) -> str:
        return pulumi.get(self, "stream_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetDatabaseLogSubscriptionResult(GetDatabaseLogSubscriptionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseLogSubscriptionResult(
            cluster_id=self.cluster_id,
            created_at=self.created_at,
            engine=self.engine,
            id=self.id,
            kind=self.kind,
            ldp_service_name=self.ldp_service_name,
            resource_name=self.resource_name,
            resource_type=self.resource_type,
            service_name=self.service_name,
            stream_id=self.stream_id,
            updated_at=self.updated_at)


def get_database_log_subscription(cluster_id: Optional[str] = None,
                                  engine: Optional[str] = None,
                                  id: Optional[str] = None,
                                  service_name: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseLogSubscriptionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['id'] = id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getDatabaseLogSubscription:getDatabaseLogSubscription', __args__, opts=opts, typ=GetDatabaseLogSubscriptionResult).value

    return AwaitableGetDatabaseLogSubscriptionResult(
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        engine=pulumi.get(__ret__, 'engine'),
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        ldp_service_name=pulumi.get(__ret__, 'ldp_service_name'),
        resource_name=pulumi.get(__ret__, 'resource_name'),
        resource_type=pulumi.get(__ret__, 'resource_type'),
        service_name=pulumi.get(__ret__, 'service_name'),
        stream_id=pulumi.get(__ret__, 'stream_id'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_database_log_subscription_output(cluster_id: Optional[pulumi.Input[str]] = None,
                                         engine: Optional[pulumi.Input[str]] = None,
                                         id: Optional[pulumi.Input[str]] = None,
                                         service_name: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDatabaseLogSubscriptionResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['id'] = id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProjectDatabase/getDatabaseLogSubscription:getDatabaseLogSubscription', __args__, opts=opts, typ=GetDatabaseLogSubscriptionResult)
    return __ret__.apply(lambda __response__: GetDatabaseLogSubscriptionResult(
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        created_at=pulumi.get(__response__, 'created_at'),
        engine=pulumi.get(__response__, 'engine'),
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        ldp_service_name=pulumi.get(__response__, 'ldp_service_name'),
        resource_name=pulumi.get(__response__, 'resource_name'),
        resource_type=pulumi.get(__response__, 'resource_type'),
        service_name=pulumi.get(__response__, 'service_name'),
        stream_id=pulumi.get(__response__, 'stream_id'),
        updated_at=pulumi.get(__response__, 'updated_at')))
