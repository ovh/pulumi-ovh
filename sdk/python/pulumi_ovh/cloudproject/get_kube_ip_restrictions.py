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
    'GetKubeIpRestrictionsResult',
    'AwaitableGetKubeIpRestrictionsResult',
    'get_kube_ip_restrictions',
    'get_kube_ip_restrictions_output',
]

@pulumi.output_type
class GetKubeIpRestrictionsResult:
    """
    A collection of values returned by getKubeIpRestrictions.
    """
    def __init__(__self__, id=None, ips=None, kube_id=None, service_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if kube_id and not isinstance(kube_id, str):
            raise TypeError("Expected argument 'kube_id' to be a str")
        pulumi.set(__self__, "kube_id", kube_id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Sequence[str]:
        """
        The list of CIDRs that restricts the access to the API server.
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "kube_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetKubeIpRestrictionsResult(GetKubeIpRestrictionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubeIpRestrictionsResult(
            id=self.id,
            ips=self.ips,
            kube_id=self.kube_id,
            service_name=self.service_name)


def get_kube_ip_restrictions(kube_id: Optional[str] = None,
                             service_name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKubeIpRestrictionsResult:
    """
    Use this data source to get a OVHcloud Managed Kubernetes Service cluster IP restrictions.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    ip_restrictions = ovh.CloudProject.get_kube_ip_restrictions(service_name="XXXXXX",
        kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
    pulumi.export("ips", ip_restrictions.ips)
    ```


    :param str kube_id: The id of the managed kubernetes cluster.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['kubeId'] = kube_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getKubeIpRestrictions:getKubeIpRestrictions', __args__, opts=opts, typ=GetKubeIpRestrictionsResult).value

    return AwaitableGetKubeIpRestrictionsResult(
        id=pulumi.get(__ret__, 'id'),
        ips=pulumi.get(__ret__, 'ips'),
        kube_id=pulumi.get(__ret__, 'kube_id'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_kube_ip_restrictions_output(kube_id: Optional[pulumi.Input[str]] = None,
                                    service_name: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetKubeIpRestrictionsResult]:
    """
    Use this data source to get a OVHcloud Managed Kubernetes Service cluster IP restrictions.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    ip_restrictions = ovh.CloudProject.get_kube_ip_restrictions(service_name="XXXXXX",
        kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
    pulumi.export("ips", ip_restrictions.ips)
    ```


    :param str kube_id: The id of the managed kubernetes cluster.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['kubeId'] = kube_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getKubeIpRestrictions:getKubeIpRestrictions', __args__, opts=opts, typ=GetKubeIpRestrictionsResult)
    return __ret__.apply(lambda __response__: GetKubeIpRestrictionsResult(
        id=pulumi.get(__response__, 'id'),
        ips=pulumi.get(__response__, 'ips'),
        kube_id=pulumi.get(__response__, 'kube_id'),
        service_name=pulumi.get(__response__, 'service_name')))
