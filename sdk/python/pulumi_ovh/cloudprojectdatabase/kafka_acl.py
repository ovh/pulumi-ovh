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

__all__ = ['KafkaAclArgs', 'KafkaAcl']

@pulumi.input_type
class KafkaAclArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 permission: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 topic: pulumi.Input[str],
                 username: pulumi.Input[str]):
        """
        The set of arguments for constructing a KafkaAcl resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] permission: Permission to give to this username on this topic. Available permissions:
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] topic: Topic affected by this ACL.
        :param pulumi.Input[str] username: Username affected by this ACL.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "permission", permission)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "topic", topic)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Input[str]:
        """
        Permission to give to this username on this topic. Available permissions:
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Input[str]:
        """
        Topic affected by this ACL.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        Username affected by this ACL.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _KafkaAclState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KafkaAcl resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] permission: Permission to give to this username on this topic. Available permissions:
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] topic: Topic affected by this ACL.
        :param pulumi.Input[str] username: Username affected by this ACL.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if permission is not None:
            pulumi.set(__self__, "permission", permission)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if topic is not None:
            pulumi.set(__self__, "topic", topic)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def permission(self) -> Optional[pulumi.Input[str]]:
        """
        Permission to give to this username on this topic. Available permissions:
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def topic(self) -> Optional[pulumi.Input[str]]:
        """
        Topic affected by this ACL.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Username affected by this ACL.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class KafkaAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an ACL for a kafka cluster associated with a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        kafka = ovh.CloudProjectDatabase.get_database(service_name="XXX",
            engine="kafka",
            id="ZZZ")
        acl = ovh.cloud_project_database.KafkaAcl("acl",
            service_name=kafka.service_name,
            cluster_id=kafka.id,
            permission="read",
            topic="mytopic",
            username="johndoe")
        ```

        ## Import

        OVHcloud Managed kafka clusters ACLs can be imported using the `service_name`, `cluster_id` and `id` of the acl, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl my_acl service_name/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] permission: Permission to give to this username on this topic. Available permissions:
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] topic: Topic affected by this ACL.
        :param pulumi.Input[str] username: Username affected by this ACL.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KafkaAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an ACL for a kafka cluster associated with a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        kafka = ovh.CloudProjectDatabase.get_database(service_name="XXX",
            engine="kafka",
            id="ZZZ")
        acl = ovh.cloud_project_database.KafkaAcl("acl",
            service_name=kafka.service_name,
            cluster_id=kafka.id,
            permission="read",
            topic="mytopic",
            username="johndoe")
        ```

        ## Import

        OVHcloud Managed kafka clusters ACLs can be imported using the `service_name`, `cluster_id` and `id` of the acl, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl my_acl service_name/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param KafkaAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KafkaAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KafkaAclArgs.__new__(KafkaAclArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if permission is None and not opts.urn:
                raise TypeError("Missing required property 'permission'")
            __props__.__dict__["permission"] = permission
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if topic is None and not opts.urn:
                raise TypeError("Missing required property 'topic'")
            __props__.__dict__["topic"] = topic
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        super(KafkaAcl, __self__).__init__(
            'ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            permission: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            topic: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'KafkaAcl':
        """
        Get an existing KafkaAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] permission: Permission to give to this username on this topic. Available permissions:
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] topic: Topic affected by this ACL.
        :param pulumi.Input[str] username: Username affected by this ACL.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KafkaAclState.__new__(_KafkaAclState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["permission"] = permission
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["topic"] = topic
        __props__.__dict__["username"] = username
        return KafkaAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Output[str]:
        """
        Permission to give to this username on this topic. Available permissions:
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Output[str]:
        """
        Topic affected by this ACL.
        """
        return pulumi.get(self, "topic")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        Username affected by this ACL.
        """
        return pulumi.get(self, "username")

