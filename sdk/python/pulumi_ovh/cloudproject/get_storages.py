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
from . import outputs

__all__ = [
    'GetStoragesResult',
    'AwaitableGetStoragesResult',
    'get_storages',
    'get_storages_output',
]

@pulumi.output_type
class GetStoragesResult:
    """
    A collection of values returned by getStorages.
    """
    def __init__(__self__, containers=None, id=None, region_name=None, service_name=None):
        if containers and not isinstance(containers, list):
            raise TypeError("Expected argument 'containers' to be a list")
        pulumi.set(__self__, "containers", containers)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region_name and not isinstance(region_name, str):
            raise TypeError("Expected argument 'region_name' to be a str")
        pulumi.set(__self__, "region_name", region_name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def containers(self) -> Sequence['outputs.GetStoragesContainerResult']:
        return pulumi.get(self, "containers")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> builtins.str:
        """
        Region name
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        """
        Service name
        """
        return pulumi.get(self, "service_name")


class AwaitableGetStoragesResult(GetStoragesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStoragesResult(
            containers=self.containers,
            id=self.id,
            region_name=self.region_name,
            service_name=self.service_name)


def get_storages(region_name: Optional[builtins.str] = None,
                 service_name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStoragesResult:
    """
    List your S3™* compatible storage container. \\* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    storage = ovh.CloudProject.get_storage(service_name="<public cloud project ID>",
        region_name="GRA")
    ```


    :param builtins.str region_name: Region name
    :param builtins.str service_name: Service name
    """
    __args__ = dict()
    __args__['regionName'] = region_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getStorages:getStorages', __args__, opts=opts, typ=GetStoragesResult).value

    return AwaitableGetStoragesResult(
        containers=pulumi.get(__ret__, 'containers'),
        id=pulumi.get(__ret__, 'id'),
        region_name=pulumi.get(__ret__, 'region_name'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_storages_output(region_name: Optional[pulumi.Input[builtins.str]] = None,
                        service_name: Optional[pulumi.Input[builtins.str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStoragesResult]:
    """
    List your S3™* compatible storage container. \\* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    storage = ovh.CloudProject.get_storage(service_name="<public cloud project ID>",
        region_name="GRA")
    ```


    :param builtins.str region_name: Region name
    :param builtins.str service_name: Service name
    """
    __args__ = dict()
    __args__['regionName'] = region_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getStorages:getStorages', __args__, opts=opts, typ=GetStoragesResult)
    return __ret__.apply(lambda __response__: GetStoragesResult(
        containers=pulumi.get(__response__, 'containers'),
        id=pulumi.get(__response__, 'id'),
        region_name=pulumi.get(__response__, 'region_name'),
        service_name=pulumi.get(__response__, 'service_name')))
