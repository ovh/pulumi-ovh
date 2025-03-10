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
    'GetRegionLoadBalancerLogSubscriptionsResult',
    'AwaitableGetRegionLoadBalancerLogSubscriptionsResult',
    'get_region_load_balancer_log_subscriptions',
    'get_region_load_balancer_log_subscriptions_output',
]

@pulumi.output_type
class GetRegionLoadBalancerLogSubscriptionsResult:
    """
    A collection of values returned by getRegionLoadBalancerLogSubscriptions.
    """
    def __init__(__self__, id=None, kind=None, loadbalancer_id=None, region_name=None, service_name=None, subscription_ids=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if loadbalancer_id and not isinstance(loadbalancer_id, str):
            raise TypeError("Expected argument 'loadbalancer_id' to be a str")
        pulumi.set(__self__, "loadbalancer_id", loadbalancer_id)
        if region_name and not isinstance(region_name, str):
            raise TypeError("Expected argument 'region_name' to be a str")
        pulumi.set(__self__, "region_name", region_name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if subscription_ids and not isinstance(subscription_ids, list):
            raise TypeError("Expected argument 'subscription_ids' to be a list")
        pulumi.set(__self__, "subscription_ids", subscription_ids)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="loadbalancerId")
    def loadbalancer_id(self) -> str:
        return pulumi.get(self, "loadbalancer_id")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> str:
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="subscriptionIds")
    def subscription_ids(self) -> Sequence[str]:
        return pulumi.get(self, "subscription_ids")


class AwaitableGetRegionLoadBalancerLogSubscriptionsResult(GetRegionLoadBalancerLogSubscriptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionLoadBalancerLogSubscriptionsResult(
            id=self.id,
            kind=self.kind,
            loadbalancer_id=self.loadbalancer_id,
            region_name=self.region_name,
            service_name=self.service_name,
            subscription_ids=self.subscription_ids)


def get_region_load_balancer_log_subscriptions(kind: Optional[str] = None,
                                               loadbalancer_id: Optional[str] = None,
                                               region_name: Optional[str] = None,
                                               service_name: Optional[str] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionLoadBalancerLogSubscriptionsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['kind'] = kind
    __args__['loadbalancerId'] = loadbalancer_id
    __args__['regionName'] = region_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getRegionLoadBalancerLogSubscriptions:getRegionLoadBalancerLogSubscriptions', __args__, opts=opts, typ=GetRegionLoadBalancerLogSubscriptionsResult).value

    return AwaitableGetRegionLoadBalancerLogSubscriptionsResult(
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        loadbalancer_id=pulumi.get(__ret__, 'loadbalancer_id'),
        region_name=pulumi.get(__ret__, 'region_name'),
        service_name=pulumi.get(__ret__, 'service_name'),
        subscription_ids=pulumi.get(__ret__, 'subscription_ids'))
def get_region_load_balancer_log_subscriptions_output(kind: Optional[pulumi.Input[Optional[str]]] = None,
                                                      loadbalancer_id: Optional[pulumi.Input[str]] = None,
                                                      region_name: Optional[pulumi.Input[str]] = None,
                                                      service_name: Optional[pulumi.Input[str]] = None,
                                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRegionLoadBalancerLogSubscriptionsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['kind'] = kind
    __args__['loadbalancerId'] = loadbalancer_id
    __args__['regionName'] = region_name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getRegionLoadBalancerLogSubscriptions:getRegionLoadBalancerLogSubscriptions', __args__, opts=opts, typ=GetRegionLoadBalancerLogSubscriptionsResult)
    return __ret__.apply(lambda __response__: GetRegionLoadBalancerLogSubscriptionsResult(
        id=pulumi.get(__response__, 'id'),
        kind=pulumi.get(__response__, 'kind'),
        loadbalancer_id=pulumi.get(__response__, 'loadbalancer_id'),
        region_name=pulumi.get(__response__, 'region_name'),
        service_name=pulumi.get(__response__, 'service_name'),
        subscription_ids=pulumi.get(__response__, 'subscription_ids')))
