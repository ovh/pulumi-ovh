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
from . import _utilities

__all__ = [
    'GetServersResult',
    'AwaitableGetServersResult',
    'get_servers',
    'get_servers_output',
]

@pulumi.output_type
class GetServersResult:
    """
    A collection of values returned by getServers.
    """
    def __init__(__self__, id=None, results=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def results(self) -> Sequence[str]:
        return pulumi.get(self, "results")


class AwaitableGetServersResult(GetServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServersResult(
            id=self.id,
            results=self.results)


def get_servers(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:index/getServers:getServers', __args__, opts=opts, typ=GetServersResult).value

    return AwaitableGetServersResult(
        id=pulumi.get(__ret__, 'id'),
        results=pulumi.get(__ret__, 'results'))
def get_servers_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:index/getServers:getServers', __args__, opts=opts, typ=GetServersResult)
    return __ret__.apply(lambda __response__: GetServersResult(
        id=pulumi.get(__response__, 'id'),
        results=pulumi.get(__response__, 'results')))
