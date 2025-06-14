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
    'GetCloudDirectorBackupResult',
    'AwaitableGetCloudDirectorBackupResult',
    'get_cloud_director_backup',
    'get_cloud_director_backup_output',
]

@pulumi.output_type
class GetCloudDirectorBackupResult:
    """
    A collection of values returned by getCloudDirectorBackup.
    """
    def __init__(__self__, backup_id=None, created_at=None, current_state=None, current_tasks=None, iam=None, id=None, resource_status=None, target_spec=None, updated_at=None):
        if backup_id and not isinstance(backup_id, str):
            raise TypeError("Expected argument 'backup_id' to be a str")
        pulumi.set(__self__, "backup_id", backup_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
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
    @pulumi.getter(name="backupId")
    def backup_id(self) -> builtins.str:
        """
        Backup ID
        """
        return pulumi.get(self, "backup_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> builtins.str:
        """
        Datetime when backup was enabled
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentState")
    def current_state(self) -> 'outputs.GetCloudDirectorBackupCurrentStateResult':
        """
        VMware Cloud Director Backup service current state
        """
        return pulumi.get(self, "current_state")

    @property
    @pulumi.getter(name="currentTasks")
    def current_tasks(self) -> Sequence['outputs.GetCloudDirectorBackupCurrentTaskResult']:
        """
        Asynchronous operations ongoing on the VMware Cloud Director organization backup service
        """
        return pulumi.get(self, "current_tasks")

    @property
    @pulumi.getter
    def iam(self) -> 'outputs.GetCloudDirectorBackupIamResult':
        """
        IAM resource metadata
        """
        return pulumi.get(self, "iam")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Unique identifier of the VMware Cloud Director backup
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> builtins.str:
        """
        Reflects the readiness of the VMware Cloud Director organization backup service
        """
        return pulumi.get(self, "resource_status")

    @property
    @pulumi.getter(name="targetSpec")
    def target_spec(self) -> 'outputs.GetCloudDirectorBackupTargetSpecResult':
        """
        VMware Cloud Director Backup target spec
        """
        return pulumi.get(self, "target_spec")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> builtins.str:
        """
        Datetime when backup is modified
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetCloudDirectorBackupResult(GetCloudDirectorBackupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudDirectorBackupResult(
            backup_id=self.backup_id,
            created_at=self.created_at,
            current_state=self.current_state,
            current_tasks=self.current_tasks,
            iam=self.iam,
            id=self.id,
            resource_status=self.resource_status,
            target_spec=self.target_spec,
            updated_at=self.updated_at)


def get_cloud_director_backup(backup_id: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudDirectorBackupResult:
    """
    Get information about a VMware Cloud Director Backup service

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    backup = ovh.VMware.get_cloud_director_backup(backup_id="<VCD backup ID>")
    ```


    :param builtins.str backup_id: Backup ID
    """
    __args__ = dict()
    __args__['backupId'] = backup_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:VMware/getCloudDirectorBackup:getCloudDirectorBackup', __args__, opts=opts, typ=GetCloudDirectorBackupResult).value

    return AwaitableGetCloudDirectorBackupResult(
        backup_id=pulumi.get(__ret__, 'backup_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        current_state=pulumi.get(__ret__, 'current_state'),
        current_tasks=pulumi.get(__ret__, 'current_tasks'),
        iam=pulumi.get(__ret__, 'iam'),
        id=pulumi.get(__ret__, 'id'),
        resource_status=pulumi.get(__ret__, 'resource_status'),
        target_spec=pulumi.get(__ret__, 'target_spec'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_cloud_director_backup_output(backup_id: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCloudDirectorBackupResult]:
    """
    Get information about a VMware Cloud Director Backup service

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    backup = ovh.VMware.get_cloud_director_backup(backup_id="<VCD backup ID>")
    ```


    :param builtins.str backup_id: Backup ID
    """
    __args__ = dict()
    __args__['backupId'] = backup_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:VMware/getCloudDirectorBackup:getCloudDirectorBackup', __args__, opts=opts, typ=GetCloudDirectorBackupResult)
    return __ret__.apply(lambda __response__: GetCloudDirectorBackupResult(
        backup_id=pulumi.get(__response__, 'backup_id'),
        created_at=pulumi.get(__response__, 'created_at'),
        current_state=pulumi.get(__response__, 'current_state'),
        current_tasks=pulumi.get(__response__, 'current_tasks'),
        iam=pulumi.get(__response__, 'iam'),
        id=pulumi.get(__response__, 'id'),
        resource_status=pulumi.get(__response__, 'resource_status'),
        target_spec=pulumi.get(__response__, 'target_spec'),
        updated_at=pulumi.get(__response__, 'updated_at')))
