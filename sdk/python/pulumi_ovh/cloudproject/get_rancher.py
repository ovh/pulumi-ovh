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
    'GetRancherResult',
    'AwaitableGetRancherResult',
    'get_rancher',
    'get_rancher_output',
]

@pulumi.output_type
class GetRancherResult:
    """
    A collection of values returned by getRancher.
    """
    def __init__(__self__, created_at=None, current_state=None, current_tasks=None, id=None, project_id=None, resource_status=None, target_spec=None, updated_at=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if current_state and not isinstance(current_state, dict):
            raise TypeError("Expected argument 'current_state' to be a dict")
        pulumi.set(__self__, "current_state", current_state)
        if current_tasks and not isinstance(current_tasks, list):
            raise TypeError("Expected argument 'current_tasks' to be a list")
        pulumi.set(__self__, "current_tasks", current_tasks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
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
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Date of the managed Rancher service creation
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentState")
    def current_state(self) -> 'outputs.GetRancherCurrentStateResult':
        """
        Current configuration applied to the managed Rancher service
        """
        return pulumi.get(self, "current_state")

    @property
    @pulumi.getter(name="currentTasks")
    def current_tasks(self) -> Sequence['outputs.GetRancherCurrentTaskResult']:
        """
        Asynchronous operations ongoing on the managed Rancher service
        """
        return pulumi.get(self, "current_tasks")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Unique identifier
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> str:
        """
        Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
        """
        return pulumi.get(self, "resource_status")

    @property
    @pulumi.getter(name="targetSpec")
    def target_spec(self) -> 'outputs.GetRancherTargetSpecResult':
        """
        Last target specification of the managed Rancher service
        """
        return pulumi.get(self, "target_spec")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        Date of the last managed Rancher service update
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetRancherResult(GetRancherResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRancherResult(
            created_at=self.created_at,
            current_state=self.current_state,
            current_tasks=self.current_tasks,
            id=self.id,
            project_id=self.project_id,
            resource_status=self.resource_status,
            target_spec=self.target_spec,
            updated_at=self.updated_at)


def get_rancher(id: Optional[str] = None,
                project_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRancherResult:
    """
    Retrieve information about a Managed Rancher Service in the given public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    rancher = ovh.CloudProject.get_rancher(id="<Rancher service ID>",
        project_id="<public cloud project ID>")
    ```


    :param str id: Unique identifier
    :param str project_id: Project ID
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getRancher:getRancher', __args__, opts=opts, typ=GetRancherResult).value

    return AwaitableGetRancherResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        current_state=pulumi.get(__ret__, 'current_state'),
        current_tasks=pulumi.get(__ret__, 'current_tasks'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        resource_status=pulumi.get(__ret__, 'resource_status'),
        target_spec=pulumi.get(__ret__, 'target_spec'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_rancher_output(id: Optional[pulumi.Input[str]] = None,
                       project_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRancherResult]:
    """
    Retrieve information about a Managed Rancher Service in the given public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    rancher = ovh.CloudProject.get_rancher(id="<Rancher service ID>",
        project_id="<public cloud project ID>")
    ```


    :param str id: Unique identifier
    :param str project_id: Project ID
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getRancher:getRancher', __args__, opts=opts, typ=GetRancherResult)
    return __ret__.apply(lambda __response__: GetRancherResult(
        created_at=pulumi.get(__response__, 'created_at'),
        current_state=pulumi.get(__response__, 'current_state'),
        current_tasks=pulumi.get(__response__, 'current_tasks'),
        id=pulumi.get(__response__, 'id'),
        project_id=pulumi.get(__response__, 'project_id'),
        resource_status=pulumi.get(__response__, 'resource_status'),
        target_spec=pulumi.get(__response__, 'target_spec'),
        updated_at=pulumi.get(__response__, 'updated_at')))
