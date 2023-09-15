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
    'GetContainerRegistryUsersResult',
    'AwaitableGetContainerRegistryUsersResult',
    'get_container_registry_users',
    'get_container_registry_users_output',
]

@pulumi.output_type
class GetContainerRegistryUsersResult:
    """
    A collection of values returned by getContainerRegistryUsers.
    """
    def __init__(__self__, id=None, registry_id=None, results=None, service_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if registry_id and not isinstance(registry_id, str):
            raise TypeError("Expected argument 'registry_id' to be a str")
        pulumi.set(__self__, "registry_id", registry_id)
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
    @pulumi.getter(name="registryId")
    def registry_id(self) -> str:
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetContainerRegistryUsersResultResult']:
        """
        The list of users of the container registry associated with the project.
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")


class AwaitableGetContainerRegistryUsersResult(GetContainerRegistryUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerRegistryUsersResult(
            id=self.id,
            registry_id=self.registry_id,
            results=self.results,
            service_name=self.service_name)


def get_container_registry_users(registry_id: Optional[str] = None,
                                 service_name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerRegistryUsersResult:
    """
    Use this data source to get the list of users of a container registry associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_registry = ovh.CloudProject.get_container_registry(service_name="XXXXXX",
        registry_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
    users = ovh.CloudProject.get_container_registry_users(service_name=ovh_cloud_project_containerregistry["registry"]["service_name"],
        registry_id=ovh_cloud_project_containerregistry["registry"]["id"])
    ```


    :param str registry_id: Registry ID
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['registryId'] = registry_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getContainerRegistryUsers:getContainerRegistryUsers', __args__, opts=opts, typ=GetContainerRegistryUsersResult).value

    return AwaitableGetContainerRegistryUsersResult(
        id=pulumi.get(__ret__, 'id'),
        registry_id=pulumi.get(__ret__, 'registry_id'),
        results=pulumi.get(__ret__, 'results'),
        service_name=pulumi.get(__ret__, 'service_name'))


@_utilities.lift_output_func(get_container_registry_users)
def get_container_registry_users_output(registry_id: Optional[pulumi.Input[str]] = None,
                                        service_name: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContainerRegistryUsersResult]:
    """
    Use this data source to get the list of users of a container registry associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_registry = ovh.CloudProject.get_container_registry(service_name="XXXXXX",
        registry_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
    users = ovh.CloudProject.get_container_registry_users(service_name=ovh_cloud_project_containerregistry["registry"]["service_name"],
        registry_id=ovh_cloud_project_containerregistry["registry"]["id"])
    ```


    :param str registry_id: Registry ID
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
