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
    'GetLogsClustersResult',
    'AwaitableGetLogsClustersResult',
    'get_logs_clusters',
    'get_logs_clusters_output',
]

@pulumi.output_type
class GetLogsClustersResult:
    """
    A collection of values returned by getLogsClusters.
    """
    def __init__(__self__, id=None, service_name=None, urn=None, uuids=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if urn and not isinstance(urn, str):
            raise TypeError("Expected argument 'urn' to be a str")
        pulumi.set(__self__, "urn", urn)
        if uuids and not isinstance(uuids, list):
            raise TypeError("Expected argument 'uuids' to be a list")
        pulumi.set(__self__, "uuids", uuids)

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
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def urn(self) -> str:
        return pulumi.get(self, "urn")

    @property
    @pulumi.getter
    def uuids(self) -> Sequence[str]:
        return pulumi.get(self, "uuids")


class AwaitableGetLogsClustersResult(GetLogsClustersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLogsClustersResult(
            id=self.id,
            service_name=self.service_name,
            urn=self.urn,
            uuids=self.uuids)


def get_logs_clusters(service_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLogsClustersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dbaas/getLogsClusters:getLogsClusters', __args__, opts=opts, typ=GetLogsClustersResult).value

    return AwaitableGetLogsClustersResult(
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'),
        urn=pulumi.get(__ret__, 'urn'),
        uuids=pulumi.get(__ret__, 'uuids'))
def get_logs_clusters_output(service_name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLogsClustersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Dbaas/getLogsClusters:getLogsClusters', __args__, opts=opts, typ=GetLogsClustersResult)
    return __ret__.apply(lambda __response__: GetLogsClustersResult(
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name'),
        urn=pulumi.get(__response__, 'urn'),
        uuids=pulumi.get(__response__, 'uuids')))
