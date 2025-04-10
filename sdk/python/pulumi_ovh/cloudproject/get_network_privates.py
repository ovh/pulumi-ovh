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
from . import outputs

__all__ = [
    'GetNetworkPrivatesResult',
    'AwaitableGetNetworkPrivatesResult',
    'get_network_privates',
    'get_network_privates_output',
]

@pulumi.output_type
class GetNetworkPrivatesResult:
    """
    A collection of values returned by getNetworkPrivates.
    """
    def __init__(__self__, id=None, networks=None, service_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.GetNetworkPrivatesNetworkResult']:
        """
        List of network
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        """
        ID of the public cloud project
        """
        return pulumi.get(self, "service_name")


class AwaitableGetNetworkPrivatesResult(GetNetworkPrivatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkPrivatesResult(
            id=self.id,
            networks=self.networks,
            service_name=self.service_name)


def get_network_privates(service_name: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkPrivatesResult:
    """
    List public cloud project private networks.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    private_network_privates = ovh.CloudProject.get_network_privates(service_name="XXXXXX")
    pulumi.export("private", private_network_privates)
    ```


    :param builtins.str service_name: The ID of the public cloud project.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getNetworkPrivates:getNetworkPrivates', __args__, opts=opts, typ=GetNetworkPrivatesResult).value

    return AwaitableGetNetworkPrivatesResult(
        id=pulumi.get(__ret__, 'id'),
        networks=pulumi.get(__ret__, 'networks'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_network_privates_output(service_name: Optional[pulumi.Input[builtins.str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNetworkPrivatesResult]:
    """
    List public cloud project private networks.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    private_network_privates = ovh.CloudProject.get_network_privates(service_name="XXXXXX")
    pulumi.export("private", private_network_privates)
    ```


    :param builtins.str service_name: The ID of the public cloud project.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getNetworkPrivates:getNetworkPrivates', __args__, opts=opts, typ=GetNetworkPrivatesResult)
    return __ret__.apply(lambda __response__: GetNetworkPrivatesResult(
        id=pulumi.get(__response__, 'id'),
        networks=pulumi.get(__response__, 'networks'),
        service_name=pulumi.get(__response__, 'service_name')))
