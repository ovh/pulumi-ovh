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

__all__ = ['IdentityUserArgs', 'IdentityUser']

@pulumi.input_type
class IdentityUserArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[builtins.str],
                 login: pulumi.Input[builtins.str],
                 password: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a IdentityUser resource.
        :param pulumi.Input[builtins.str] email: User's email.
        :param pulumi.Input[builtins.str] login: User's login suffix.
        :param pulumi.Input[builtins.str] password: User's password.
        :param pulumi.Input[builtins.str] description: User description.
        :param pulumi.Input[builtins.str] group: User's group.
        """
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "login", login)
        pulumi.set(__self__, "password", password)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group is not None:
            pulumi.set(__self__, "group", group)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[builtins.str]:
        """
        User's email.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def login(self) -> pulumi.Input[builtins.str]:
        """
        User's login suffix.
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[builtins.str]:
        """
        User's password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User's group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group", value)


@pulumi.input_type
class _IdentityUserState:
    def __init__(__self__, *,
                 user_urn: Optional[pulumi.Input[builtins.str]] = None,
                 creation: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 last_update: Optional[pulumi.Input[builtins.str]] = None,
                 login: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 password_last_update: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering IdentityUser resources.
        :param pulumi.Input[builtins.str] user_urn: URN of the user, used when writing IAM policies
        :param pulumi.Input[builtins.str] creation: Creation date of this user.
        :param pulumi.Input[builtins.str] description: User description.
        :param pulumi.Input[builtins.str] email: User's email.
        :param pulumi.Input[builtins.str] group: User's group.
        :param pulumi.Input[builtins.str] last_update: Last update of this user.
        :param pulumi.Input[builtins.str] login: User's login suffix.
        :param pulumi.Input[builtins.str] password: User's password.
        :param pulumi.Input[builtins.str] password_last_update: When the user changed his password for the last time.
        :param pulumi.Input[builtins.str] status: Current user's status.
        """
        if user_urn is not None:
            pulumi.set(__self__, "user_urn", user_urn)
        if creation is not None:
            pulumi.set(__self__, "creation", creation)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if last_update is not None:
            pulumi.set(__self__, "last_update", last_update)
        if login is not None:
            pulumi.set(__self__, "login", login)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if password_last_update is not None:
            pulumi.set(__self__, "password_last_update", password_last_update)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="UserURN")
    def user_urn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        URN of the user, used when writing IAM policies
        """
        return pulumi.get(self, "user_urn")

    @user_urn.setter
    def user_urn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "user_urn", value)

    @property
    @pulumi.getter
    def creation(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Creation date of this user.
        """
        return pulumi.get(self, "creation")

    @creation.setter
    def creation(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "creation", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User's email.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User's group.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="lastUpdate")
    def last_update(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Last update of this user.
        """
        return pulumi.get(self, "last_update")

    @last_update.setter
    def last_update(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "last_update", value)

    @property
    @pulumi.getter
    def login(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User's login suffix.
        """
        return pulumi.get(self, "login")

    @login.setter
    def login(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "login", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User's password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="passwordLastUpdate")
    def password_last_update(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        When the user changed his password for the last time.
        """
        return pulumi.get(self, "password_last_update")

    @password_last_update.setter
    def password_last_update(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "password_last_update", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Current user's status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)


@pulumi.type_token("ovh:Me/identityUser:IdentityUser")
class IdentityUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 login: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates an identity user.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_user = ovh.me.IdentityUser("my_user",
            description="Some custom description",
            email="my_login@example.com",
            group="DEFAULT",
            login="my_login",
            password="super-s3cr3t!password")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: User description.
        :param pulumi.Input[builtins.str] email: User's email.
        :param pulumi.Input[builtins.str] group: User's group.
        :param pulumi.Input[builtins.str] login: User's login suffix.
        :param pulumi.Input[builtins.str] password: User's password.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdentityUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an identity user.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_user = ovh.me.IdentityUser("my_user",
            description="Some custom description",
            email="my_login@example.com",
            group="DEFAULT",
            login="my_login",
            password="super-s3cr3t!password")
        ```

        :param str resource_name: The name of the resource.
        :param IdentityUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 group: Optional[pulumi.Input[builtins.str]] = None,
                 login: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdentityUserArgs.__new__(IdentityUserArgs)

            __props__.__dict__["description"] = description
            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            __props__.__dict__["group"] = group
            if login is None and not opts.urn:
                raise TypeError("Missing required property 'login'")
            __props__.__dict__["login"] = login
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["user_urn"] = None
            __props__.__dict__["creation"] = None
            __props__.__dict__["last_update"] = None
            __props__.__dict__["password_last_update"] = None
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(IdentityUser, __self__).__init__(
            'ovh:Me/identityUser:IdentityUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            user_urn: Optional[pulumi.Input[builtins.str]] = None,
            creation: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            email: Optional[pulumi.Input[builtins.str]] = None,
            group: Optional[pulumi.Input[builtins.str]] = None,
            last_update: Optional[pulumi.Input[builtins.str]] = None,
            login: Optional[pulumi.Input[builtins.str]] = None,
            password: Optional[pulumi.Input[builtins.str]] = None,
            password_last_update: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None) -> 'IdentityUser':
        """
        Get an existing IdentityUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] user_urn: URN of the user, used when writing IAM policies
        :param pulumi.Input[builtins.str] creation: Creation date of this user.
        :param pulumi.Input[builtins.str] description: User description.
        :param pulumi.Input[builtins.str] email: User's email.
        :param pulumi.Input[builtins.str] group: User's group.
        :param pulumi.Input[builtins.str] last_update: Last update of this user.
        :param pulumi.Input[builtins.str] login: User's login suffix.
        :param pulumi.Input[builtins.str] password: User's password.
        :param pulumi.Input[builtins.str] password_last_update: When the user changed his password for the last time.
        :param pulumi.Input[builtins.str] status: Current user's status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdentityUserState.__new__(_IdentityUserState)

        __props__.__dict__["user_urn"] = user_urn
        __props__.__dict__["creation"] = creation
        __props__.__dict__["description"] = description
        __props__.__dict__["email"] = email
        __props__.__dict__["group"] = group
        __props__.__dict__["last_update"] = last_update
        __props__.__dict__["login"] = login
        __props__.__dict__["password"] = password
        __props__.__dict__["password_last_update"] = password_last_update
        __props__.__dict__["status"] = status
        return IdentityUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="UserURN")
    def user_urn(self) -> pulumi.Output[builtins.str]:
        """
        URN of the user, used when writing IAM policies
        """
        return pulumi.get(self, "user_urn")

    @property
    @pulumi.getter
    def creation(self) -> pulumi.Output[builtins.str]:
        """
        Creation date of this user.
        """
        return pulumi.get(self, "creation")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        User description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[builtins.str]:
        """
        User's email.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        User's group.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="lastUpdate")
    def last_update(self) -> pulumi.Output[builtins.str]:
        """
        Last update of this user.
        """
        return pulumi.get(self, "last_update")

    @property
    @pulumi.getter
    def login(self) -> pulumi.Output[builtins.str]:
        """
        User's login suffix.
        """
        return pulumi.get(self, "login")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[builtins.str]:
        """
        User's password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="passwordLastUpdate")
    def password_last_update(self) -> pulumi.Output[builtins.str]:
        """
        When the user changed his password for the last time.
        """
        return pulumi.get(self, "password_last_update")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Current user's status.
        """
        return pulumi.get(self, "status")

