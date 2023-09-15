# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetCapabilitiesContainerRegistryResult',
    'AwaitableGetCapabilitiesContainerRegistryResult',
    'get_capabilities_container_registry',
    'get_capabilities_container_registry_output',
]

@pulumi.output_type
class GetCapabilitiesContainerRegistryResult:
    """
    A collection of values returned by getCapabilitiesContainerRegistry.
    """
    def __init__(__self__, id=None, results=None, service_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
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
    def results(self) -> Sequence['outputs.GetCapabilitiesContainerRegistryResultResult']:
        """
        List of container registry capability for a single region
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")


class AwaitableGetCapabilitiesContainerRegistryResult(GetCapabilitiesContainerRegistryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCapabilitiesContainerRegistryResult(
            id=self.id,
            results=self.results,
            service_name=self.service_name)


def get_capabilities_container_registry(service_name: Optional[str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCapabilitiesContainerRegistryResult:
    """
    Use this data source to get the container registry capabilities of a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    capabilities = ovh.CloudProject.get_capabilities_container_registry(service_name="XXXXXX")
    ```


    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getCapabilitiesContainerRegistry:getCapabilitiesContainerRegistry', __args__, opts=opts, typ=GetCapabilitiesContainerRegistryResult).value

    return AwaitableGetCapabilitiesContainerRegistryResult(
        id=pulumi.get(__ret__, 'id'),
        results=pulumi.get(__ret__, 'results'),
        service_name=pulumi.get(__ret__, 'service_name'))


@_utilities.lift_output_func(get_capabilities_container_registry)
def get_capabilities_container_registry_output(service_name: Optional[pulumi.Input[str]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCapabilitiesContainerRegistryResult]:
    """
    Use this data source to get the container registry capabilities of a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    capabilities = ovh.CloudProject.get_capabilities_container_registry(service_name="XXXXXX")
    ```


    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
