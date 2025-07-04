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

__all__ = ['LogsRoleArgs', 'LogsRole']

@pulumi.input_type
class LogsRoleArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[builtins.str],
                 service_name: pulumi.Input[builtins.str],
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a LogsRole resource.
        :param pulumi.Input[builtins.str] description: The role description
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[builtins.str] name: The role name
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "service_name", service_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[builtins.str]:
        """
        The role description
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
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The role name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _LogsRoleState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 nb_member: Optional[pulumi.Input[builtins.int]] = None,
                 nb_permission: Optional[pulumi.Input[builtins.int]] = None,
                 role_id: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering LogsRole resources.
        :param pulumi.Input[builtins.str] created_at: Role creation date
        :param pulumi.Input[builtins.str] description: The role description
        :param pulumi.Input[builtins.str] name: The role name
        :param pulumi.Input[builtins.int] nb_member: number of member for the role
        :param pulumi.Input[builtins.int] nb_permission: number of configured permission for the role
        :param pulumi.Input[builtins.str] role_id: Role identifier
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[builtins.str] updated_at: Role last update date
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nb_member is not None:
            pulumi.set(__self__, "nb_member", nb_member)
        if nb_permission is not None:
            pulumi.set(__self__, "nb_permission", nb_permission)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Role creation date
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The role description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The role name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nbMember")
    def nb_member(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        number of member for the role
        """
        return pulumi.get(self, "nb_member")

    @nb_member.setter
    def nb_member(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "nb_member", value)

    @property
    @pulumi.getter(name="nbPermission")
    def nb_permission(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        number of configured permission for the role
        """
        return pulumi.get(self, "nb_permission")

    @nb_permission.setter
    def nb_permission(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "nb_permission", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Role identifier
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role_id", value)

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
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Role last update date
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)


@pulumi.type_token("ovh:Dbaas/logsRole:LogsRole")
class LogsRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Reference a DBaaS logs role.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        ro = ovh.dbaas.LogsRole("ro",
            service_name="ldp-xx-xxxxx",
            name="Devops - RO",
            description="Devops - RO")
        ```

        ## Import

        OVHcloud DBaaS Log Role can be imported using the `service_name` and `role_id` of the role, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Dbaas/logsRole:LogsRole  ovh_dbaas_logs_role.ro ldp-ra-XX/dc145bc2-eb01-4efe-a802-XXXXXX
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: The role description
        :param pulumi.Input[builtins.str] name: The role name
        :param pulumi.Input[builtins.str] service_name: The service name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogsRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Reference a DBaaS logs role.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        ro = ovh.dbaas.LogsRole("ro",
            service_name="ldp-xx-xxxxx",
            name="Devops - RO",
            description="Devops - RO")
        ```

        ## Import

        OVHcloud DBaaS Log Role can be imported using the `service_name` and `role_id` of the role, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Dbaas/logsRole:LogsRole  ovh_dbaas_logs_role.ro ldp-ra-XX/dc145bc2-eb01-4efe-a802-XXXXXX
        ```

        :param str resource_name: The name of the resource.
        :param LogsRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogsRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogsRoleArgs.__new__(LogsRoleArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["created_at"] = None
            __props__.__dict__["nb_member"] = None
            __props__.__dict__["nb_permission"] = None
            __props__.__dict__["role_id"] = None
            __props__.__dict__["updated_at"] = None
        super(LogsRole, __self__).__init__(
            'ovh:Dbaas/logsRole:LogsRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            nb_member: Optional[pulumi.Input[builtins.int]] = None,
            nb_permission: Optional[pulumi.Input[builtins.int]] = None,
            role_id: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None) -> 'LogsRole':
        """
        Get an existing LogsRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] created_at: Role creation date
        :param pulumi.Input[builtins.str] description: The role description
        :param pulumi.Input[builtins.str] name: The role name
        :param pulumi.Input[builtins.int] nb_member: number of member for the role
        :param pulumi.Input[builtins.int] nb_permission: number of configured permission for the role
        :param pulumi.Input[builtins.str] role_id: Role identifier
        :param pulumi.Input[builtins.str] service_name: The service name
        :param pulumi.Input[builtins.str] updated_at: Role last update date
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogsRoleState.__new__(_LogsRoleState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["nb_member"] = nb_member
        __props__.__dict__["nb_permission"] = nb_permission
        __props__.__dict__["role_id"] = role_id
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["updated_at"] = updated_at
        return LogsRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Role creation date
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[builtins.str]:
        """
        The role description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The role name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nbMember")
    def nb_member(self) -> pulumi.Output[builtins.int]:
        """
        number of member for the role
        """
        return pulumi.get(self, "nb_member")

    @property
    @pulumi.getter(name="nbPermission")
    def nb_permission(self) -> pulumi.Output[builtins.int]:
        """
        number of configured permission for the role
        """
        return pulumi.get(self, "nb_permission")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[builtins.str]:
        """
        Role identifier
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The service name
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        Role last update date
        """
        return pulumi.get(self, "updated_at")

