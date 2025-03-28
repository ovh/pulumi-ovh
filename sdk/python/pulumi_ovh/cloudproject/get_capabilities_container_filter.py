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
from . import outputs

__all__ = [
    'GetCapabilitiesContainerFilterResult',
    'AwaitableGetCapabilitiesContainerFilterResult',
    'get_capabilities_container_filter',
    'get_capabilities_container_filter_output',
]

@pulumi.output_type
class GetCapabilitiesContainerFilterResult:
    """
    A collection of values returned by getCapabilitiesContainerFilter.
    """
    def __init__(__self__, code=None, created_at=None, features=None, id=None, name=None, plan_name=None, region=None, registry_limits=None, service_name=None, updated_at=None):
        if code and not isinstance(code, str):
            raise TypeError("Expected argument 'code' to be a str")
        pulumi.set(__self__, "code", code)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        pulumi.set(__self__, "features", features)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if plan_name and not isinstance(plan_name, str):
            raise TypeError("Expected argument 'plan_name' to be a str")
        pulumi.set(__self__, "plan_name", plan_name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if registry_limits and not isinstance(registry_limits, list):
            raise TypeError("Expected argument 'registry_limits' to be a list")
        pulumi.set(__self__, "registry_limits", registry_limits)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Plan code from the catalog
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Plan creation date
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def features(self) -> Sequence['outputs.GetCapabilitiesContainerFilterFeatureResult']:
        """
        Features of the plan
        """
        return pulumi.get(self, "features")

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
        """
        Plan name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="planName")
    def plan_name(self) -> str:
        return pulumi.get(self, "plan_name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="registryLimits")
    def registry_limits(self) -> Sequence['outputs.GetCapabilitiesContainerFilterRegistryLimitResult']:
        """
        Container registry limits
        """
        return pulumi.get(self, "registry_limits")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Plan last update date
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetCapabilitiesContainerFilterResult(GetCapabilitiesContainerFilterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCapabilitiesContainerFilterResult(
            code=self.code,
            created_at=self.created_at,
            features=self.features,
            id=self.id,
            name=self.name,
            plan_name=self.plan_name,
            region=self.region,
            registry_limits=self.registry_limits,
            service_name=self.service_name,
            updated_at=self.updated_at)


def get_capabilities_container_filter(plan_name: Optional[str] = None,
                                      region: Optional[str] = None,
                                      service_name: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCapabilitiesContainerFilterResult:
    """
    Use this data source to filter the list of container registry capabilities associated with a public cloud project to match one and only one capability.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    capability = ovh.CloudProject.get_capabilities_container_filter(plan_name="SMALL",
        region="GRA",
        service_name="XXXXXX")
    ```


    :param str plan_name: The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'.
    :param str region: The region name
    :param str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['planName'] = plan_name
    __args__['region'] = region
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter', __args__, opts=opts, typ=GetCapabilitiesContainerFilterResult).value

    return AwaitableGetCapabilitiesContainerFilterResult(
        code=pulumi.get(__ret__, 'code'),
        created_at=pulumi.get(__ret__, 'created_at'),
        features=pulumi.get(__ret__, 'features'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        plan_name=pulumi.get(__ret__, 'plan_name'),
        region=pulumi.get(__ret__, 'region'),
        registry_limits=pulumi.get(__ret__, 'registry_limits'),
        service_name=pulumi.get(__ret__, 'service_name'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_capabilities_container_filter_output(plan_name: Optional[pulumi.Input[str]] = None,
                                             region: Optional[pulumi.Input[str]] = None,
                                             service_name: Optional[pulumi.Input[str]] = None,
                                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCapabilitiesContainerFilterResult]:
    """
    Use this data source to filter the list of container registry capabilities associated with a public cloud project to match one and only one capability.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    capability = ovh.CloudProject.get_capabilities_container_filter(plan_name="SMALL",
        region="GRA",
        service_name="XXXXXX")
    ```


    :param str plan_name: The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'.
    :param str region: The region name
    :param str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['planName'] = plan_name
    __args__['region'] = region
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter', __args__, opts=opts, typ=GetCapabilitiesContainerFilterResult)
    return __ret__.apply(lambda __response__: GetCapabilitiesContainerFilterResult(
        code=pulumi.get(__response__, 'code'),
        created_at=pulumi.get(__response__, 'created_at'),
        features=pulumi.get(__response__, 'features'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        plan_name=pulumi.get(__response__, 'plan_name'),
        region=pulumi.get(__response__, 'region'),
        registry_limits=pulumi.get(__response__, 'registry_limits'),
        service_name=pulumi.get(__response__, 'service_name'),
        updated_at=pulumi.get(__response__, 'updated_at')))
