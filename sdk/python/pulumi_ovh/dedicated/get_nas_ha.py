# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
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
    'GetNasHAResult',
    'AwaitableGetNasHAResult',
    'get_nas_ha',
    'get_nas_ha_output',
]

@pulumi.output_type
class GetNasHAResult:
    """
    A collection of values returned by getNasHA.
    """
    def __init__(__self__, nas_haurn=None, can_create_partition=None, custom_name=None, datacenter=None, disk_type=None, id=None, ip=None, monitored=None, partitions_lists=None, service_name=None, zpool_capacity=None, zpool_size=None):
        if nas_haurn and not isinstance(nas_haurn, str):
            raise TypeError("Expected argument 'nas_haurn' to be a str")
        pulumi.set(__self__, "nas_haurn", nas_haurn)
        if can_create_partition and not isinstance(can_create_partition, bool):
            raise TypeError("Expected argument 'can_create_partition' to be a bool")
        pulumi.set(__self__, "can_create_partition", can_create_partition)
        if custom_name and not isinstance(custom_name, str):
            raise TypeError("Expected argument 'custom_name' to be a str")
        pulumi.set(__self__, "custom_name", custom_name)
        if datacenter and not isinstance(datacenter, str):
            raise TypeError("Expected argument 'datacenter' to be a str")
        pulumi.set(__self__, "datacenter", datacenter)
        if disk_type and not isinstance(disk_type, str):
            raise TypeError("Expected argument 'disk_type' to be a str")
        pulumi.set(__self__, "disk_type", disk_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if monitored and not isinstance(monitored, bool):
            raise TypeError("Expected argument 'monitored' to be a bool")
        pulumi.set(__self__, "monitored", monitored)
        if partitions_lists and not isinstance(partitions_lists, list):
            raise TypeError("Expected argument 'partitions_lists' to be a list")
        pulumi.set(__self__, "partitions_lists", partitions_lists)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if zpool_capacity and not isinstance(zpool_capacity, float):
            raise TypeError("Expected argument 'zpool_capacity' to be a float")
        pulumi.set(__self__, "zpool_capacity", zpool_capacity)
        if zpool_size and not isinstance(zpool_size, float):
            raise TypeError("Expected argument 'zpool_size' to be a float")
        pulumi.set(__self__, "zpool_size", zpool_size)

    @property
    @pulumi.getter(name="NasHAURN")
    def nas_haurn(self) -> builtins.str:
        """
        the URN of the HA-NAS instance
        """
        return pulumi.get(self, "nas_haurn")

    @property
    @pulumi.getter(name="canCreatePartition")
    def can_create_partition(self) -> builtins.bool:
        """
        True, if partition creation is allowed on this HA-NAS
        """
        return pulumi.get(self, "can_create_partition")

    @property
    @pulumi.getter(name="customName")
    def custom_name(self) -> builtins.str:
        """
        The name you give to the HA-NAS
        """
        return pulumi.get(self, "custom_name")

    @property
    @pulumi.getter
    def datacenter(self) -> builtins.str:
        """
        area of HA-NAS
        """
        return pulumi.get(self, "datacenter")

    @property
    @pulumi.getter(name="diskType")
    def disk_type(self) -> builtins.str:
        """
        the disk type of the HA-NAS. Possible values are: `hdd`, `ssd`, `nvme`
        """
        return pulumi.get(self, "disk_type")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> builtins.str:
        """
        Access IP of HA-NAS
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def monitored(self) -> builtins.bool:
        """
        Send an email to customer if any issue is detected
        """
        return pulumi.get(self, "monitored")

    @property
    @pulumi.getter(name="partitionsLists")
    def partitions_lists(self) -> Sequence[builtins.str]:
        """
        the list of the HA-NAS partitions name
        """
        return pulumi.get(self, "partitions_lists")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        """
        The storage service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="zpoolCapacity")
    def zpool_capacity(self) -> builtins.float:
        """
        percentage of HA-NAS space used in %
        """
        return pulumi.get(self, "zpool_capacity")

    @property
    @pulumi.getter(name="zpoolSize")
    def zpool_size(self) -> builtins.float:
        """
        the size of the HA-NAS in GB
        """
        return pulumi.get(self, "zpool_size")


class AwaitableGetNasHAResult(GetNasHAResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNasHAResult(
            nas_haurn=self.nas_haurn,
            can_create_partition=self.can_create_partition,
            custom_name=self.custom_name,
            datacenter=self.datacenter,
            disk_type=self.disk_type,
            id=self.id,
            ip=self.ip,
            monitored=self.monitored,
            partitions_lists=self.partitions_lists,
            service_name=self.service_name,
            zpool_capacity=self.zpool_capacity,
            zpool_size=self.zpool_size)


def get_nas_ha(service_name: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNasHAResult:
    """
    Use this data source to retrieve information about a dedicated HA-NAS.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_nas_ha = ovh.Dedicated.get_nas_ha(service_name="zpool-12345")
    ```


    :param builtins.str service_name: The service_name of your dedicated HA-NAS.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dedicated/getNasHA:getNasHA', __args__, opts=opts, typ=GetNasHAResult).value

    return AwaitableGetNasHAResult(
        nas_haurn=pulumi.get(__ret__, 'nas_haurn'),
        can_create_partition=pulumi.get(__ret__, 'can_create_partition'),
        custom_name=pulumi.get(__ret__, 'custom_name'),
        datacenter=pulumi.get(__ret__, 'datacenter'),
        disk_type=pulumi.get(__ret__, 'disk_type'),
        id=pulumi.get(__ret__, 'id'),
        ip=pulumi.get(__ret__, 'ip'),
        monitored=pulumi.get(__ret__, 'monitored'),
        partitions_lists=pulumi.get(__ret__, 'partitions_lists'),
        service_name=pulumi.get(__ret__, 'service_name'),
        zpool_capacity=pulumi.get(__ret__, 'zpool_capacity'),
        zpool_size=pulumi.get(__ret__, 'zpool_size'))
def get_nas_ha_output(service_name: Optional[pulumi.Input[builtins.str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNasHAResult]:
    """
    Use this data source to retrieve information about a dedicated HA-NAS.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_nas_ha = ovh.Dedicated.get_nas_ha(service_name="zpool-12345")
    ```


    :param builtins.str service_name: The service_name of your dedicated HA-NAS.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Dedicated/getNasHA:getNasHA', __args__, opts=opts, typ=GetNasHAResult)
    return __ret__.apply(lambda __response__: GetNasHAResult(
        nas_haurn=pulumi.get(__response__, 'nas_haurn'),
        can_create_partition=pulumi.get(__response__, 'can_create_partition'),
        custom_name=pulumi.get(__response__, 'custom_name'),
        datacenter=pulumi.get(__response__, 'datacenter'),
        disk_type=pulumi.get(__response__, 'disk_type'),
        id=pulumi.get(__response__, 'id'),
        ip=pulumi.get(__response__, 'ip'),
        monitored=pulumi.get(__response__, 'monitored'),
        partitions_lists=pulumi.get(__response__, 'partitions_lists'),
        service_name=pulumi.get(__response__, 'service_name'),
        zpool_capacity=pulumi.get(__response__, 'zpool_capacity'),
        zpool_size=pulumi.get(__response__, 'zpool_size')))
