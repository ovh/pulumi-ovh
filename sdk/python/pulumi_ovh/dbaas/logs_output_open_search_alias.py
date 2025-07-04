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

__all__ = ['LogsOutputOpenSearchAliasArgs', 'LogsOutputOpenSearchAlias']

@pulumi.input_type
class LogsOutputOpenSearchAliasArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[builtins.str],
                 service_name: pulumi.Input[builtins.str],
                 suffix: pulumi.Input[builtins.str],
                 indexes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 nb_index: Optional[pulumi.Input[builtins.int]] = None,
                 nb_stream: Optional[pulumi.Input[builtins.int]] = None,
                 streams: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a LogsOutputOpenSearchAlias resource.
        :param pulumi.Input[builtins.str] description: Index description
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[builtins.str] suffix: Index suffix
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] indexes: List of attached indexes id
        :param pulumi.Input[builtins.int] nb_index: Number of indices linked
        :param pulumi.Input[builtins.int] nb_stream: Number of streams linked
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] streams: List of attached streams id
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "suffix", suffix)
        if indexes is not None:
            pulumi.set(__self__, "indexes", indexes)
        if nb_index is not None:
            pulumi.set(__self__, "nb_index", nb_index)
        if nb_stream is not None:
            pulumi.set(__self__, "nb_stream", nb_stream)
        if streams is not None:
            pulumi.set(__self__, "streams", streams)

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

    @property
    @pulumi.getter
    def indexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of attached indexes id
        """
        return pulumi.get(self, "indexes")

    @indexes.setter
    def indexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "indexes", value)

    @property
    @pulumi.getter(name="nbIndex")
    def nb_index(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Number of indices linked
        """
        return pulumi.get(self, "nb_index")

    @nb_index.setter
    def nb_index(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "nb_index", value)

    @property
    @pulumi.getter(name="nbStream")
    def nb_stream(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Number of streams linked
        """
        return pulumi.get(self, "nb_stream")

    @nb_stream.setter
    def nb_stream(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "nb_stream", value)

    @property
    @pulumi.getter
    def streams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of attached streams id
        """
        return pulumi.get(self, "streams")

    @streams.setter
    def streams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "streams", value)


@pulumi.input_type
class _LogsOutputOpenSearchAliasState:
    def __init__(__self__, *,
                 alias_id: Optional[pulumi.Input[builtins.str]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 current_size: Optional[pulumi.Input[builtins.int]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 indexes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 is_editable: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 nb_index: Optional[pulumi.Input[builtins.int]] = None,
                 nb_stream: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 streams: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 suffix: Optional[pulumi.Input[builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering LogsOutputOpenSearchAlias resources.
        :param pulumi.Input[builtins.str] alias_id: Alias Id
        :param pulumi.Input[builtins.str] created_at: Alias creation
        :param pulumi.Input[builtins.int] current_size: Current alias size (in bytes)
        :param pulumi.Input[builtins.str] description: Index description
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] indexes: List of attached indexes id
        :param pulumi.Input[builtins.bool] is_editable: Indicates if you are allowed to edit entry
        :param pulumi.Input[builtins.str] name: Alias name
        :param pulumi.Input[builtins.int] nb_index: Number of indices linked
        :param pulumi.Input[builtins.int] nb_stream: Number of streams linked
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] streams: List of attached streams id
        :param pulumi.Input[builtins.str] suffix: Index suffix
        :param pulumi.Input[builtins.str] updated_at: Input last update
        """
        if alias_id is not None:
            pulumi.set(__self__, "alias_id", alias_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if current_size is not None:
            pulumi.set(__self__, "current_size", current_size)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if indexes is not None:
            pulumi.set(__self__, "indexes", indexes)
        if is_editable is not None:
            pulumi.set(__self__, "is_editable", is_editable)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nb_index is not None:
            pulumi.set(__self__, "nb_index", nb_index)
        if nb_stream is not None:
            pulumi.set(__self__, "nb_stream", nb_stream)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if streams is not None:
            pulumi.set(__self__, "streams", streams)
        if suffix is not None:
            pulumi.set(__self__, "suffix", suffix)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="aliasId")
    def alias_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Alias Id
        """
        return pulumi.get(self, "alias_id")

    @alias_id.setter
    def alias_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "alias_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Alias creation
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="currentSize")
    def current_size(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Current alias size (in bytes)
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
    @pulumi.getter
    def indexes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of attached indexes id
        """
        return pulumi.get(self, "indexes")

    @indexes.setter
    def indexes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "indexes", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Alias name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nbIndex")
    def nb_index(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Number of indices linked
        """
        return pulumi.get(self, "nb_index")

    @nb_index.setter
    def nb_index(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "nb_index", value)

    @property
    @pulumi.getter(name="nbStream")
    def nb_stream(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Number of streams linked
        """
        return pulumi.get(self, "nb_stream")

    @nb_stream.setter
    def nb_stream(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "nb_stream", value)

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
    def streams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        List of attached streams id
        """
        return pulumi.get(self, "streams")

    @streams.setter
    def streams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "streams", value)

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
        Input last update
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)


@pulumi.type_token("ovh:Dbaas/logsOutputOpenSearchAlias:LogsOutputOpenSearchAlias")
class LogsOutputOpenSearchAlias(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 indexes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 nb_index: Optional[pulumi.Input[builtins.int]] = None,
                 nb_stream: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 streams: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 suffix: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates a DBaaS Logs Opensearch output alias.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        alias = ovh.dbaas.LogsOutputOpenSearchAlias("alias",
            service_name="....",
            description="my opensearch alias",
            suffix="alias")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Index description
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] indexes: List of attached indexes id
        :param pulumi.Input[builtins.int] nb_index: Number of indices linked
        :param pulumi.Input[builtins.int] nb_stream: Number of streams linked
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] streams: List of attached streams id
        :param pulumi.Input[builtins.str] suffix: Index suffix
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogsOutputOpenSearchAliasArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a DBaaS Logs Opensearch output alias.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        alias = ovh.dbaas.LogsOutputOpenSearchAlias("alias",
            service_name="....",
            description="my opensearch alias",
            suffix="alias")
        ```

        :param str resource_name: The name of the resource.
        :param LogsOutputOpenSearchAliasArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogsOutputOpenSearchAliasArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 indexes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 nb_index: Optional[pulumi.Input[builtins.int]] = None,
                 nb_stream: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 streams: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 suffix: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogsOutputOpenSearchAliasArgs.__new__(LogsOutputOpenSearchAliasArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["indexes"] = indexes
            __props__.__dict__["nb_index"] = nb_index
            __props__.__dict__["nb_stream"] = nb_stream
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["streams"] = streams
            if suffix is None and not opts.urn:
                raise TypeError("Missing required property 'suffix'")
            __props__.__dict__["suffix"] = suffix
            __props__.__dict__["alias_id"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["current_size"] = None
            __props__.__dict__["is_editable"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["updated_at"] = None
        super(LogsOutputOpenSearchAlias, __self__).__init__(
            'ovh:Dbaas/logsOutputOpenSearchAlias:LogsOutputOpenSearchAlias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias_id: Optional[pulumi.Input[builtins.str]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            current_size: Optional[pulumi.Input[builtins.int]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            indexes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            is_editable: Optional[pulumi.Input[builtins.bool]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            nb_index: Optional[pulumi.Input[builtins.int]] = None,
            nb_stream: Optional[pulumi.Input[builtins.int]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            streams: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            suffix: Optional[pulumi.Input[builtins.str]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None) -> 'LogsOutputOpenSearchAlias':
        """
        Get an existing LogsOutputOpenSearchAlias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] alias_id: Alias Id
        :param pulumi.Input[builtins.str] created_at: Alias creation
        :param pulumi.Input[builtins.int] current_size: Current alias size (in bytes)
        :param pulumi.Input[builtins.str] description: Index description
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] indexes: List of attached indexes id
        :param pulumi.Input[builtins.bool] is_editable: Indicates if you are allowed to edit entry
        :param pulumi.Input[builtins.str] name: Alias name
        :param pulumi.Input[builtins.int] nb_index: Number of indices linked
        :param pulumi.Input[builtins.int] nb_stream: Number of streams linked
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] streams: List of attached streams id
        :param pulumi.Input[builtins.str] suffix: Index suffix
        :param pulumi.Input[builtins.str] updated_at: Input last update
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogsOutputOpenSearchAliasState.__new__(_LogsOutputOpenSearchAliasState)

        __props__.__dict__["alias_id"] = alias_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["current_size"] = current_size
        __props__.__dict__["description"] = description
        __props__.__dict__["indexes"] = indexes
        __props__.__dict__["is_editable"] = is_editable
        __props__.__dict__["name"] = name
        __props__.__dict__["nb_index"] = nb_index
        __props__.__dict__["nb_stream"] = nb_stream
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["streams"] = streams
        __props__.__dict__["suffix"] = suffix
        __props__.__dict__["updated_at"] = updated_at
        return LogsOutputOpenSearchAlias(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aliasId")
    def alias_id(self) -> pulumi.Output[builtins.str]:
        """
        Alias Id
        """
        return pulumi.get(self, "alias_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Alias creation
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentSize")
    def current_size(self) -> pulumi.Output[builtins.int]:
        """
        Current alias size (in bytes)
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
    @pulumi.getter
    def indexes(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        List of attached indexes id
        """
        return pulumi.get(self, "indexes")

    @property
    @pulumi.getter(name="isEditable")
    def is_editable(self) -> pulumi.Output[builtins.bool]:
        """
        Indicates if you are allowed to edit entry
        """
        return pulumi.get(self, "is_editable")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Alias name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nbIndex")
    def nb_index(self) -> pulumi.Output[builtins.int]:
        """
        Number of indices linked
        """
        return pulumi.get(self, "nb_index")

    @property
    @pulumi.getter(name="nbStream")
    def nb_stream(self) -> pulumi.Output[builtins.int]:
        """
        Number of streams linked
        """
        return pulumi.get(self, "nb_stream")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def streams(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        List of attached streams id
        """
        return pulumi.get(self, "streams")

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
        Input last update
        """
        return pulumi.get(self, "updated_at")

