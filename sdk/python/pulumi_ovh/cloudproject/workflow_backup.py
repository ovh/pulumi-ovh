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

__all__ = ['WorkflowBackupArgs', 'WorkflowBackup']

@pulumi.input_type
class WorkflowBackupArgs:
    def __init__(__self__, *,
                 cron: pulumi.Input[builtins.str],
                 instance_id: pulumi.Input[builtins.str],
                 region_name: pulumi.Input[builtins.str],
                 rotation: pulumi.Input[builtins.int],
                 backup_name: Optional[pulumi.Input[builtins.str]] = None,
                 max_execution_count: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a WorkflowBackup resource.
        :param pulumi.Input[builtins.str] cron: The cron periodicity at which the backup workflow is scheduled
               
               * `instanceId` the id of the instance to back up
        :param pulumi.Input[builtins.str] region_name: The name of the openstack region.
        :param pulumi.Input[builtins.int] rotation: The number of backup that are retained.
        :param pulumi.Input[builtins.str] backup_name: The name of the backup files that are created. If empty, the `name` attribute is used.
        :param pulumi.Input[builtins.int] max_execution_count: The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        :param pulumi.Input[builtins.str] name: The worflow name that is used in the UI
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        pulumi.set(__self__, "cron", cron)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "region_name", region_name)
        pulumi.set(__self__, "rotation", rotation)
        if backup_name is not None:
            pulumi.set(__self__, "backup_name", backup_name)
        if max_execution_count is not None:
            pulumi.set(__self__, "max_execution_count", max_execution_count)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def cron(self) -> pulumi.Input[builtins.str]:
        """
        The cron periodicity at which the backup workflow is scheduled

        * `instanceId` the id of the instance to back up
        """
        return pulumi.get(self, "cron")

    @cron.setter
    def cron(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cron", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the openstack region.
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter
    def rotation(self) -> pulumi.Input[builtins.int]:
        """
        The number of backup that are retained.
        """
        return pulumi.get(self, "rotation")

    @rotation.setter
    def rotation(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "rotation", value)

    @property
    @pulumi.getter(name="backupName")
    def backup_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the backup files that are created. If empty, the `name` attribute is used.
        """
        return pulumi.get(self, "backup_name")

    @backup_name.setter
    def backup_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backup_name", value)

    @property
    @pulumi.getter(name="maxExecutionCount")
    def max_execution_count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        """
        return pulumi.get(self, "max_execution_count")

    @max_execution_count.setter
    def max_execution_count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_execution_count", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The worflow name that is used in the UI
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _WorkflowBackupState:
    def __init__(__self__, *,
                 backup_name: Optional[pulumi.Input[builtins.str]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 cron: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 max_execution_count: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region_name: Optional[pulumi.Input[builtins.str]] = None,
                 rotation: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering WorkflowBackup resources.
        :param pulumi.Input[builtins.str] backup_name: The name of the backup files that are created. If empty, the `name` attribute is used.
        :param pulumi.Input[builtins.str] cron: The cron periodicity at which the backup workflow is scheduled
               
               * `instanceId` the id of the instance to back up
        :param pulumi.Input[builtins.int] max_execution_count: The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        :param pulumi.Input[builtins.str] name: The worflow name that is used in the UI
        :param pulumi.Input[builtins.str] region_name: The name of the openstack region.
        :param pulumi.Input[builtins.int] rotation: The number of backup that are retained.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        if backup_name is not None:
            pulumi.set(__self__, "backup_name", backup_name)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if cron is not None:
            pulumi.set(__self__, "cron", cron)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if max_execution_count is not None:
            pulumi.set(__self__, "max_execution_count", max_execution_count)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region_name is not None:
            pulumi.set(__self__, "region_name", region_name)
        if rotation is not None:
            pulumi.set(__self__, "rotation", rotation)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="backupName")
    def backup_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the backup files that are created. If empty, the `name` attribute is used.
        """
        return pulumi.get(self, "backup_name")

    @backup_name.setter
    def backup_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backup_name", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def cron(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The cron periodicity at which the backup workflow is scheduled

        * `instanceId` the id of the instance to back up
        """
        return pulumi.get(self, "cron")

    @cron.setter
    def cron(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cron", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="maxExecutionCount")
    def max_execution_count(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        """
        return pulumi.get(self, "max_execution_count")

    @max_execution_count.setter
    def max_execution_count(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_execution_count", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The worflow name that is used in the UI
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the openstack region.
        """
        return pulumi.get(self, "region_name")

    @region_name.setter
    def region_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region_name", value)

    @property
    @pulumi.getter
    def rotation(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The number of backup that are retained.
        """
        return pulumi.get(self, "rotation")

    @rotation.setter
    def rotation(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "rotation", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)


@pulumi.type_token("ovh:CloudProject/workflowBackup:WorkflowBackup")
class WorkflowBackup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_name: Optional[pulumi.Input[builtins.str]] = None,
                 cron: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 max_execution_count: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region_name: Optional[pulumi.Input[builtins.str]] = None,
                 rotation: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manage a worflow that schedules backups of public cloud instance. Note that upon deletion, the workflow is deleted but any backups that have been created by this workflow are not.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_backup = ovh.cloud_project.WorkflowBackup("my_backup",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            region_name="GRA11",
            cron="50 4 * * *",
            instance_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx",
            max_execution_count=0,
            name="Backup workflow for instance",
            rotation=7)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backup_name: The name of the backup files that are created. If empty, the `name` attribute is used.
        :param pulumi.Input[builtins.str] cron: The cron periodicity at which the backup workflow is scheduled
               
               * `instanceId` the id of the instance to back up
        :param pulumi.Input[builtins.int] max_execution_count: The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        :param pulumi.Input[builtins.str] name: The worflow name that is used in the UI
        :param pulumi.Input[builtins.str] region_name: The name of the openstack region.
        :param pulumi.Input[builtins.int] rotation: The number of backup that are retained.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkflowBackupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a worflow that schedules backups of public cloud instance. Note that upon deletion, the workflow is deleted but any backups that have been created by this workflow are not.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_backup = ovh.cloud_project.WorkflowBackup("my_backup",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            region_name="GRA11",
            cron="50 4 * * *",
            instance_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx",
            max_execution_count=0,
            name="Backup workflow for instance",
            rotation=7)
        ```

        :param str resource_name: The name of the resource.
        :param WorkflowBackupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkflowBackupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_name: Optional[pulumi.Input[builtins.str]] = None,
                 cron: Optional[pulumi.Input[builtins.str]] = None,
                 instance_id: Optional[pulumi.Input[builtins.str]] = None,
                 max_execution_count: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 region_name: Optional[pulumi.Input[builtins.str]] = None,
                 rotation: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkflowBackupArgs.__new__(WorkflowBackupArgs)

            __props__.__dict__["backup_name"] = backup_name
            if cron is None and not opts.urn:
                raise TypeError("Missing required property 'cron'")
            __props__.__dict__["cron"] = cron
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["max_execution_count"] = max_execution_count
            __props__.__dict__["name"] = name
            if region_name is None and not opts.urn:
                raise TypeError("Missing required property 'region_name'")
            __props__.__dict__["region_name"] = region_name
            if rotation is None and not opts.urn:
                raise TypeError("Missing required property 'rotation'")
            __props__.__dict__["rotation"] = rotation
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["created_at"] = None
        super(WorkflowBackup, __self__).__init__(
            'ovh:CloudProject/workflowBackup:WorkflowBackup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_name: Optional[pulumi.Input[builtins.str]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            cron: Optional[pulumi.Input[builtins.str]] = None,
            instance_id: Optional[pulumi.Input[builtins.str]] = None,
            max_execution_count: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            region_name: Optional[pulumi.Input[builtins.str]] = None,
            rotation: Optional[pulumi.Input[builtins.int]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None) -> 'WorkflowBackup':
        """
        Get an existing WorkflowBackup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backup_name: The name of the backup files that are created. If empty, the `name` attribute is used.
        :param pulumi.Input[builtins.str] cron: The cron periodicity at which the backup workflow is scheduled
               
               * `instanceId` the id of the instance to back up
        :param pulumi.Input[builtins.int] max_execution_count: The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        :param pulumi.Input[builtins.str] name: The worflow name that is used in the UI
        :param pulumi.Input[builtins.str] region_name: The name of the openstack region.
        :param pulumi.Input[builtins.int] rotation: The number of backup that are retained.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkflowBackupState.__new__(_WorkflowBackupState)

        __props__.__dict__["backup_name"] = backup_name
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["cron"] = cron
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["max_execution_count"] = max_execution_count
        __props__.__dict__["name"] = name
        __props__.__dict__["region_name"] = region_name
        __props__.__dict__["rotation"] = rotation
        __props__.__dict__["service_name"] = service_name
        return WorkflowBackup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupName")
    def backup_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the backup files that are created. If empty, the `name` attribute is used.
        """
        return pulumi.get(self, "backup_name")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def cron(self) -> pulumi.Output[builtins.str]:
        """
        The cron periodicity at which the backup workflow is scheduled

        * `instanceId` the id of the instance to back up
        """
        return pulumi.get(self, "cron")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="maxExecutionCount")
    def max_execution_count(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
        """
        return pulumi.get(self, "max_execution_count")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The worflow name that is used in the UI
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the openstack region.
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter
    def rotation(self) -> pulumi.Output[builtins.int]:
        """
        The number of backup that are retained.
        """
        return pulumi.get(self, "rotation")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

