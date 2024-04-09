# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetMitigationResult',
    'AwaitableGetMitigationResult',
    'get_mitigation',
    'get_mitigation_output',
]

@pulumi.output_type
class GetMitigationResult:
    """
    A collection of values returned by getMitigation.
    """
    def __init__(__self__, auto=None, id=None, ip=None, ip_on_mitigation=None, permanent=None, state=None):
        if auto and not isinstance(auto, bool):
            raise TypeError("Expected argument 'auto' to be a bool")
        pulumi.set(__self__, "auto", auto)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if ip_on_mitigation and not isinstance(ip_on_mitigation, str):
            raise TypeError("Expected argument 'ip_on_mitigation' to be a str")
        pulumi.set(__self__, "ip_on_mitigation", ip_on_mitigation)
        if permanent and not isinstance(permanent, bool):
            raise TypeError("Expected argument 'permanent' to be a bool")
        pulumi.set(__self__, "permanent", permanent)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def auto(self) -> bool:
        """
        Set on true if the IP is on auto-mitigation
        """
        return pulumi.get(self, "auto")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        The IP or the CIDR
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="ipOnMitigation")
    def ip_on_mitigation(self) -> str:
        """
        IPv4 address
        * `permanent ` - Set on true if the IP is on permanent mitigation
        """
        return pulumi.get(self, "ip_on_mitigation")

    @property
    @pulumi.getter
    def permanent(self) -> bool:
        return pulumi.get(self, "permanent")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the IP on mitigation
        """
        return pulumi.get(self, "state")


class AwaitableGetMitigationResult(GetMitigationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMitigationResult(
            auto=self.auto,
            id=self.id,
            ip=self.ip,
            ip_on_mitigation=self.ip_on_mitigation,
            permanent=self.permanent,
            state=self.state)


def get_mitigation(ip: Optional[str] = None,
                   ip_on_mitigation: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMitigationResult:
    """
    Use this resource to retrieve information about an IP permanent mitigation.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ovh as ovh

    mitigation_data = ovh.Ip.get_mitigation(ip="XXXXXX",
        ip_on_mitigation="XXXXXX")
    ```
    <!--End PulumiCodeChooser -->


    :param str ip: The IP or the CIDR
    :param str ip_on_mitigation: IPv4 address
    """
    __args__ = dict()
    __args__['ip'] = ip
    __args__['ipOnMitigation'] = ip_on_mitigation
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Ip/getMitigation:getMitigation', __args__, opts=opts, typ=GetMitigationResult).value

    return AwaitableGetMitigationResult(
        auto=pulumi.get(__ret__, 'auto'),
        id=pulumi.get(__ret__, 'id'),
        ip=pulumi.get(__ret__, 'ip'),
        ip_on_mitigation=pulumi.get(__ret__, 'ip_on_mitigation'),
        permanent=pulumi.get(__ret__, 'permanent'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_mitigation)
def get_mitigation_output(ip: Optional[pulumi.Input[str]] = None,
                          ip_on_mitigation: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMitigationResult]:
    """
    Use this resource to retrieve information about an IP permanent mitigation.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_ovh as ovh

    mitigation_data = ovh.Ip.get_mitigation(ip="XXXXXX",
        ip_on_mitigation="XXXXXX")
    ```
    <!--End PulumiCodeChooser -->


    :param str ip: The IP or the CIDR
    :param str ip_on_mitigation: IPv4 address
    """
    ...
