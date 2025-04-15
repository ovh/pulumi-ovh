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

__all__ = [
    'GetContainerRegistryResult',
    'AwaitableGetContainerRegistryResult',
    'get_container_registry',
    'get_container_registry_output',
]

@pulumi.output_type
class GetContainerRegistryResult:
    """
    A collection of values returned by getContainerRegistry.
    """
    def __init__(__self__, created_at=None, id=None, name=None, project_id=None, region=None, registry_id=None, service_name=None, size=None, status=None, updated_at=None, url=None, version=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if registry_id and not isinstance(registry_id, str):
            raise TypeError("Expected argument 'registry_id' to be a str")
        pulumi.set(__self__, "registry_id", registry_id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        """
        Registry creation date
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Registry name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> builtins.str:
        """
        Project ID of your registry
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        """
        Region of the registry
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> builtins.str:
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def size(self) -> builtins.int:
        """
        Current size of the registry (bytes)
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Registry status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> builtins.str:
        """
        Registry last update date
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def url(self) -> builtins.str:
        """
        Access url of the registry
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> builtins.str:
        """
        Version of your registry
        """
        return pulumi.get(self, "version")


class AwaitableGetContainerRegistryResult(GetContainerRegistryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerRegistryResult(
            created_at=self.created_at,
            id=self.id,
            name=self.name,
            project_id=self.project_id,
            region=self.region,
            registry_id=self.registry_id,
            service_name=self.service_name,
            size=self.size,
            status=self.status,
            updated_at=self.updated_at,
            url=self.url,
            version=self.version)


def get_container_registry(registry_id: Optional[builtins.str] = None,
                           service_name: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerRegistryResult:
    """
    Use this data source to get information about a container registry associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_registry = ovh.CloudProject.get_container_registry(service_name="XXXXXX",
        registry_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
    ```


    :param builtins.str registry_id: Registry ID
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['registryId'] = registry_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getContainerRegistry:getContainerRegistry', __args__, opts=opts, typ=GetContainerRegistryResult).value

    return AwaitableGetContainerRegistryResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        registry_id=pulumi.get(__ret__, 'registry_id'),
        service_name=pulumi.get(__ret__, 'service_name'),
        size=pulumi.get(__ret__, 'size'),
        status=pulumi.get(__ret__, 'status'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        url=pulumi.get(__ret__, 'url'),
        version=pulumi.get(__ret__, 'version'))
def get_container_registry_output(registry_id: Optional[pulumi.Input[builtins.str]] = None,
                                  service_name: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetContainerRegistryResult]:
    """
    Use this data source to get information about a container registry associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_registry = ovh.CloudProject.get_container_registry(service_name="XXXXXX",
        registry_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
    ```


    :param builtins.str registry_id: Registry ID
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['registryId'] = registry_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getContainerRegistry:getContainerRegistry', __args__, opts=opts, typ=GetContainerRegistryResult)
    return __ret__.apply(lambda __response__: GetContainerRegistryResult(
        created_at=pulumi.get(__response__, 'created_at'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        registry_id=pulumi.get(__response__, 'registry_id'),
        service_name=pulumi.get(__response__, 'service_name'),
        size=pulumi.get(__response__, 'size'),
        status=pulumi.get(__response__, 'status'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        url=pulumi.get(__response__, 'url'),
        version=pulumi.get(__response__, 'version')))
