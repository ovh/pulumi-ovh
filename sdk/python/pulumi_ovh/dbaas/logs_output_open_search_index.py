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

__all__ = ['LogsOutputOpenSearchIndexArgs', 'LogsOutputOpenSearchIndex']

@pulumi.input_type
class LogsOutputOpenSearchIndexArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[builtins.str],
                 nb_shard: pulumi.Input[builtins.int],
                 service_name: pulumi.Input[builtins.str],
                 suffix: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a LogsOutputOpenSearchIndex resource.
        :param pulumi.Input[builtins.str] description: Index description
        :param pulumi.Input[builtins.int] nb_shard: Number of shards
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[builtins.str] suffix: Index suffix
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "nb_shard", nb_shard)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "suffix", suffix)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[builtins.str]:
        """
        Index description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="nbShard")
    def nb_shard(self) -> pulumi.Input[builtins.int]:
        """
        Number of shards
        """
        return pulumi.get(self, "nb_shard")

    @nb_shard.setter
    def nb_shard(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "nb_shard", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def suffix(self) -> pulumi.Input[builtins.str]:
        """
        Index suffix
        """
        return pulumi.get(self, "suffix")

    @suffix.setter
    def suffix(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "suffix", value)


@pulumi.input_type
class _LogsOutputOpenSearchIndexState:
    def __init__(__self__, *,
                 alert_notify_enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 current_size: Optional[pulumi.Input[builtins.int]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 index_id: Optional[pulumi.Input[builtins.str]] = None,
                 is_editable: Optional[pulumi.Input[builtins.bool]] = None,
                 max_size: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 nb_shard: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 suffix: Optional[pulumi.Input[builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering LogsOutputOpenSearchIndex resources.
        :param pulumi.Input[builtins.bool] alert_notify_enabled: If set, notify when size is near 80, 90 or 100 % of its maximum capacity
        :param pulumi.Input[builtins.str] created_at: Index creation
        :param pulumi.Input[builtins.int] current_size: Current index size (in bytes)
        :param pulumi.Input[builtins.str] description: Index description
        :param pulumi.Input[builtins.str] index_id: Index ID
        :param pulumi.Input[builtins.bool] is_editable: Indicates if you are allowed to edit entry
        :param pulumi.Input[builtins.int] max_size: Maximum index size (in bytes)
        :param pulumi.Input[builtins.str] name: Index name
        :param pulumi.Input[builtins.int] nb_shard: Number of shards
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[builtins.str] suffix: Index suffix
        :param pulumi.Input[builtins.str] updated_at: Index last update
        """
        if alert_notify_enabled is not None:
            pulumi.set(__self__, "alert_notify_enabled", alert_notify_enabled)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if current_size is not None:
            pulumi.set(__self__, "current_size", current_size)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if index_id is not None:
            pulumi.set(__self__, "index_id", index_id)
        if is_editable is not None:
            pulumi.set(__self__, "is_editable", is_editable)
        if max_size is not None:
            pulumi.set(__self__, "max_size", max_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nb_shard is not None:
            pulumi.set(__self__, "nb_shard", nb_shard)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if suffix is not None:
            pulumi.set(__self__, "suffix", suffix)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="alertNotifyEnabled")
    def alert_notify_enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If set, notify when size is near 80, 90 or 100 % of its maximum capacity
        """
        return pulumi.get(self, "alert_notify_enabled")

    @alert_notify_enabled.setter
    def alert_notify_enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "alert_notify_enabled", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Index creation
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="currentSize")
    def current_size(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Current index size (in bytes)
        """
        return pulumi.get(self, "current_size")

    @current_size.setter
    def current_size(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "current_size", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Index description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Index ID
        """
        return pulumi.get(self, "index_id")

    @index_id.setter
    def index_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "index_id", value)

    @property
    @pulumi.getter(name="isEditable")
    def is_editable(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Indicates if you are allowed to edit entry
        """
        return pulumi.get(self, "is_editable")

    @is_editable.setter
    def is_editable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_editable", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Maximum index size (in bytes)
        """
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Index name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nbShard")
    def nb_shard(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Number of shards
        """
        return pulumi.get(self, "nb_shard")

    @nb_shard.setter
    def nb_shard(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "nb_shard", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def suffix(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Index suffix
        """
        return pulumi.get(self, "suffix")

    @suffix.setter
    def suffix(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "suffix", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Index last update
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)


class LogsOutputOpenSearchIndex(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 nb_shard: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 suffix: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates a DBaaS Logs Opensearch output index.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        index = ovh.dbaas.LogsOutputOpenSearchIndex("index",
            description="my opensearch index",
            service_name="....",
            suffix="index")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Index description
        :param pulumi.Input[builtins.int] nb_shard: Number of shards
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[builtins.str] suffix: Index suffix
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogsOutputOpenSearchIndexArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a DBaaS Logs Opensearch output index.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        index = ovh.dbaas.LogsOutputOpenSearchIndex("index",
            description="my opensearch index",
            service_name="....",
            suffix="index")
        ```

        :param str resource_name: The name of the resource.
        :param LogsOutputOpenSearchIndexArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogsOutputOpenSearchIndexArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 nb_shard: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 suffix: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogsOutputOpenSearchIndexArgs.__new__(LogsOutputOpenSearchIndexArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if nb_shard is None and not opts.urn:
                raise TypeError("Missing required property 'nb_shard'")
            __props__.__dict__["nb_shard"] = nb_shard
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if suffix is None and not opts.urn:
                raise TypeError("Missing required property 'suffix'")
            __props__.__dict__["suffix"] = suffix
            __props__.__dict__["alert_notify_enabled"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["current_size"] = None
            __props__.__dict__["index_id"] = None
            __props__.__dict__["is_editable"] = None
            __props__.__dict__["max_size"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["updated_at"] = None
        super(LogsOutputOpenSearchIndex, __self__).__init__(
            'ovh:Dbaas/logsOutputOpenSearchIndex:LogsOutputOpenSearchIndex',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_notify_enabled: Optional[pulumi.Input[builtins.bool]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            current_size: Optional[pulumi.Input[builtins.int]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            index_id: Optional[pulumi.Input[builtins.str]] = None,
            is_editable: Optional[pulumi.Input[builtins.bool]] = None,
            max_size: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            nb_shard: Optional[pulumi.Input[builtins.int]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            suffix: Optional[pulumi.Input[builtins.str]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None) -> 'LogsOutputOpenSearchIndex':
        """
        Get an existing LogsOutputOpenSearchIndex resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] alert_notify_enabled: If set, notify when size is near 80, 90 or 100 % of its maximum capacity
        :param pulumi.Input[builtins.str] created_at: Index creation
        :param pulumi.Input[builtins.int] current_size: Current index size (in bytes)
        :param pulumi.Input[builtins.str] description: Index description
        :param pulumi.Input[builtins.str] index_id: Index ID
        :param pulumi.Input[builtins.bool] is_editable: Indicates if you are allowed to edit entry
        :param pulumi.Input[builtins.int] max_size: Maximum index size (in bytes)
        :param pulumi.Input[builtins.str] name: Index name
        :param pulumi.Input[builtins.int] nb_shard: Number of shards
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[builtins.str] suffix: Index suffix
        :param pulumi.Input[builtins.str] updated_at: Index last update
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogsOutputOpenSearchIndexState.__new__(_LogsOutputOpenSearchIndexState)

        __props__.__dict__["alert_notify_enabled"] = alert_notify_enabled
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["current_size"] = current_size
        __props__.__dict__["description"] = description
        __props__.__dict__["index_id"] = index_id
        __props__.__dict__["is_editable"] = is_editable
        __props__.__dict__["max_size"] = max_size
        __props__.__dict__["name"] = name
        __props__.__dict__["nb_shard"] = nb_shard
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["suffix"] = suffix
        __props__.__dict__["updated_at"] = updated_at
        return LogsOutputOpenSearchIndex(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertNotifyEnabled")
    def alert_notify_enabled(self) -> pulumi.Output[builtins.bool]:
        """
        If set, notify when size is near 80, 90 or 100 % of its maximum capacity
        """
        return pulumi.get(self, "alert_notify_enabled")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Index creation
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentSize")
    def current_size(self) -> pulumi.Output[builtins.int]:
        """
        Current index size (in bytes)
        """
        return pulumi.get(self, "current_size")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[builtins.str]:
        """
        Index description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> pulumi.Output[builtins.str]:
        """
        Index ID
        """
        return pulumi.get(self, "index_id")

    @property
    @pulumi.getter(name="isEditable")
    def is_editable(self) -> pulumi.Output[builtins.bool]:
        """
        Indicates if you are allowed to edit entry
        """
        return pulumi.get(self, "is_editable")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Output[builtins.int]:
        """
        Maximum index size (in bytes)
        """
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Index name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nbShard")
    def nb_shard(self) -> pulumi.Output[builtins.int]:
        """
        Number of shards
        """
        return pulumi.get(self, "nb_shard")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def suffix(self) -> pulumi.Output[builtins.str]:
        """
        Index suffix
        """
        return pulumi.get(self, "suffix")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        Index last update
        """
        return pulumi.get(self, "updated_at")

