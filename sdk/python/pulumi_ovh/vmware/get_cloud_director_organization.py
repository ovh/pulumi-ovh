# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
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
    'GetCloudDirectorOrganizationResult',
    'AwaitableGetCloudDirectorOrganizationResult',
    'get_cloud_director_organization',
    'get_cloud_director_organization_output',
]

@pulumi.output_type
class GetCloudDirectorOrganizationResult:
    """
    A collection of values returned by getCloudDirectorOrganization.
    """
    def __init__(__self__, current_state=None, current_tasks=None, iam=None, id=None, organization_id=None, resource_status=None, target_spec=None, updated_at=None):
        if current_state and not isinstance(current_state, dict):
            raise TypeError("Expected argument 'current_state' to be a dict")
        pulumi.set(__self__, "current_state", current_state)
        if current_tasks and not isinstance(current_tasks, list):
            raise TypeError("Expected argument 'current_tasks' to be a list")
        pulumi.set(__self__, "current_tasks", current_tasks)
        if iam and not isinstance(iam, dict):
            raise TypeError("Expected argument 'iam' to be a dict")
        pulumi.set(__self__, "iam", iam)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if resource_status and not isinstance(resource_status, str):
            raise TypeError("Expected argument 'resource_status' to be a str")
        pulumi.set(__self__, "resource_status", resource_status)
        if target_spec and not isinstance(target_spec, dict):
            raise TypeError("Expected argument 'target_spec' to be a dict")
        pulumi.set(__self__, "target_spec", target_spec)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="currentState")
    def current_state(self) -> 'outputs.GetCloudDirectorOrganizationCurrentStateResult':
        """
        Current state of VMware Cloud Director organization service
        """
        return pulumi.get(self, "current_state")

    @property
    @pulumi.getter(name="currentTasks")
    def current_tasks(self) -> Sequence['outputs.GetCloudDirectorOrganizationCurrentTaskResult']:
        """
        Asynchronous operations ongoing on the VMware Cloud Director organization
        """
        return pulumi.get(self, "current_tasks")

    @property
    @pulumi.getter
    def iam(self) -> 'outputs.GetCloudDirectorOrganizationIamResult':
        """
        IAM resource metadata
        """
        return pulumi.get(self, "iam")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Unique identifier
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> builtins.str:
        """
        Organization ID
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> builtins.str:
        """
        Reflects the readiness of the VMware Cloud Director organization.
        """
        return pulumi.get(self, "resource_status")

    @property
    @pulumi.getter(name="targetSpec")
    def target_spec(self) -> 'outputs.GetCloudDirectorOrganizationTargetSpecResult':
        """
        Target specification of VMware Cloud Director organization service
        """
        return pulumi.get(self, "target_spec")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> builtins.str:
        return pulumi.get(self, "updated_at")


class AwaitableGetCloudDirectorOrganizationResult(GetCloudDirectorOrganizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudDirectorOrganizationResult(
            current_state=self.current_state,
            current_tasks=self.current_tasks,
            iam=self.iam,
            id=self.id,
            organization_id=self.organization_id,
            resource_status=self.resource_status,
            target_spec=self.target_spec,
            updated_at=self.updated_at)


def get_cloud_director_organization(organization_id: Optional[builtins.str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudDirectorOrganizationResult:
    """
    Get VMware Cloud Director organization details

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    organization = ovh.VMware.get_cloud_director_organization(organization_id="<VCD organization ID>")
    ```


    :param builtins.str organization_id: Organization ID
    """
    __args__ = dict()
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:VMware/getCloudDirectorOrganization:getCloudDirectorOrganization', __args__, opts=opts, typ=GetCloudDirectorOrganizationResult).value

    return AwaitableGetCloudDirectorOrganizationResult(
        current_state=pulumi.get(__ret__, 'current_state'),
        current_tasks=pulumi.get(__ret__, 'current_tasks'),
        iam=pulumi.get(__ret__, 'iam'),
        id=pulumi.get(__ret__, 'id'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        resource_status=pulumi.get(__ret__, 'resource_status'),
        target_spec=pulumi.get(__ret__, 'target_spec'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_cloud_director_organization_output(organization_id: Optional[pulumi.Input[builtins.str]] = None,
                                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCloudDirectorOrganizationResult]:
    """
    Get VMware Cloud Director organization details

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    organization = ovh.VMware.get_cloud_director_organization(organization_id="<VCD organization ID>")
    ```


    :param builtins.str organization_id: Organization ID
    """
    __args__ = dict()
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:VMware/getCloudDirectorOrganization:getCloudDirectorOrganization', __args__, opts=opts, typ=GetCloudDirectorOrganizationResult)
    return __ret__.apply(lambda __response__: GetCloudDirectorOrganizationResult(
        current_state=pulumi.get(__response__, 'current_state'),
        current_tasks=pulumi.get(__response__, 'current_tasks'),
        iam=pulumi.get(__response__, 'iam'),
        id=pulumi.get(__response__, 'id'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        resource_status=pulumi.get(__response__, 'resource_status'),
        target_spec=pulumi.get(__response__, 'target_spec'),
        updated_at=pulumi.get(__response__, 'updated_at')))
