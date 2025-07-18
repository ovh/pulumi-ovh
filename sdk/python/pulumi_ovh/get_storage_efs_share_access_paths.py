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
from . import _utilities
from . import outputs

__all__ = [
    'GetStorageEfsShareAccessPathsResult',
    'AwaitableGetStorageEfsShareAccessPathsResult',
    'get_storage_efs_share_access_paths',
    'get_storage_efs_share_access_paths_output',
]

@pulumi.output_type
class GetStorageEfsShareAccessPathsResult:
    """
    A collection of values returned by getStorageEfsShareAccessPaths.
    """
    def __init__(__self__, access_paths=None, id=None, service_name=None, share_id=None):
        if access_paths and not isinstance(access_paths, list):
            raise TypeError("Expected argument 'access_paths' to be a list")
        pulumi.set(__self__, "access_paths", access_paths)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if share_id and not isinstance(share_id, str):
            raise TypeError("Expected argument 'share_id' to be a str")
        pulumi.set(__self__, "share_id", share_id)

    @property
    @pulumi.getter(name="accessPaths")
    def access_paths(self) -> Sequence['outputs.GetStorageEfsShareAccessPathsAccessPathResult']:
        return pulumi.get(self, "access_paths")

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
        Service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="shareId")
    def share_id(self) -> builtins.str:
        """
        Share ID
        """
        return pulumi.get(self, "share_id")


class AwaitableGetStorageEfsShareAccessPathsResult(GetStorageEfsShareAccessPathsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageEfsShareAccessPathsResult(
            access_paths=self.access_paths,
            id=self.id,
            service_name=self.service_name,
            share_id=self.share_id)


def get_storage_efs_share_access_paths(service_name: Optional[builtins.str] = None,
                                       share_id: Optional[builtins.str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageEfsShareAccessPathsResult:
    """
    List access paths for a share belonging to an EFS service.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    efs = ovh.get_storage_efs(service_name="XXX")
    access_paths = ovh.get_storage_efs_share_access_paths(service_name=efs.service_name,
        share_id="XXX")
    ```


    :param builtins.str service_name: Service name
    :param builtins.str share_id: Share ID
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['shareId'] = share_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:index/getStorageEfsShareAccessPaths:getStorageEfsShareAccessPaths', __args__, opts=opts, typ=GetStorageEfsShareAccessPathsResult).value

    return AwaitableGetStorageEfsShareAccessPathsResult(
        access_paths=pulumi.get(__ret__, 'access_paths'),
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'),
        share_id=pulumi.get(__ret__, 'share_id'))
def get_storage_efs_share_access_paths_output(service_name: Optional[pulumi.Input[builtins.str]] = None,
                                              share_id: Optional[pulumi.Input[builtins.str]] = None,
                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetStorageEfsShareAccessPathsResult]:
    """
    List access paths for a share belonging to an EFS service.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    efs = ovh.get_storage_efs(service_name="XXX")
    access_paths = ovh.get_storage_efs_share_access_paths(service_name=efs.service_name,
        share_id="XXX")
    ```


    :param builtins.str service_name: Service name
    :param builtins.str share_id: Share ID
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['shareId'] = share_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:index/getStorageEfsShareAccessPaths:getStorageEfsShareAccessPaths', __args__, opts=opts, typ=GetStorageEfsShareAccessPathsResult)
    return __ret__.apply(lambda __response__: GetStorageEfsShareAccessPathsResult(
        access_paths=pulumi.get(__response__, 'access_paths'),
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name'),
        share_id=pulumi.get(__response__, 'share_id')))
