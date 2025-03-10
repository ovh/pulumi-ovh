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
    'GetNasHAPartitionResult',
    'AwaitableGetNasHAPartitionResult',
    'get_nas_ha_partition',
    'get_nas_ha_partition_output',
]

@pulumi.output_type
class GetNasHAPartitionResult:
    """
    A collection of values returned by getNasHAPartition.
    """
    def __init__(__self__, capacity=None, description=None, id=None, name=None, protocol=None, service_name=None, size=None, used_by_snapshots=None):
        if capacity and not isinstance(capacity, int):
            raise TypeError("Expected argument 'capacity' to be a int")
        pulumi.set(__self__, "capacity", capacity)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if used_by_snapshots and not isinstance(used_by_snapshots, int):
            raise TypeError("Expected argument 'used_by_snapshots' to be a int")
        pulumi.set(__self__, "used_by_snapshots", used_by_snapshots)

    @property
    @pulumi.getter
    def capacity(self) -> int:
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

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
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def size(self) -> int:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="usedBySnapshots")
    def used_by_snapshots(self) -> int:
        return pulumi.get(self, "used_by_snapshots")


class AwaitableGetNasHAPartitionResult(GetNasHAPartitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNasHAPartitionResult(
            capacity=self.capacity,
            description=self.description,
            id=self.id,
            name=self.name,
            protocol=self.protocol,
            service_name=self.service_name,
            size=self.size,
            used_by_snapshots=self.used_by_snapshots)


def get_nas_ha_partition(name: Optional[str] = None,
                         service_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNasHAPartitionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dedicated/getNasHAPartition:getNasHAPartition', __args__, opts=opts, typ=GetNasHAPartitionResult).value

    return AwaitableGetNasHAPartitionResult(
        capacity=pulumi.get(__ret__, 'capacity'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        protocol=pulumi.get(__ret__, 'protocol'),
        service_name=pulumi.get(__ret__, 'service_name'),
        size=pulumi.get(__ret__, 'size'),
        used_by_snapshots=pulumi.get(__ret__, 'used_by_snapshots'))
def get_nas_ha_partition_output(name: Optional[pulumi.Input[str]] = None,
                                service_name: Optional[pulumi.Input[str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNasHAPartitionResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Dedicated/getNasHAPartition:getNasHAPartition', __args__, opts=opts, typ=GetNasHAPartitionResult)
    return __ret__.apply(lambda __response__: GetNasHAPartitionResult(
        capacity=pulumi.get(__response__, 'capacity'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        protocol=pulumi.get(__response__, 'protocol'),
        service_name=pulumi.get(__response__, 'service_name'),
        size=pulumi.get(__response__, 'size'),
        used_by_snapshots=pulumi.get(__response__, 'used_by_snapshots')))
