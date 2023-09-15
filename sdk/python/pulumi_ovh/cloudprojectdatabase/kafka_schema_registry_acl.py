# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KafkaSchemaRegistryAclArgs', 'KafkaSchemaRegistryAcl']

@pulumi.input_type
class KafkaSchemaRegistryAclArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 permission: pulumi.Input[str],
                 resource: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 username: pulumi.Input[str]):
        """
        The set of arguments for constructing a KafkaSchemaRegistryAcl resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] permission: Permission to give to this username on this resource.
               Available permissions:
        :param pulumi.Input[str] resource: Resource affected by this schema registry ACL.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] username: Username affected by this schema registry ACL.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "permission", permission)
        pulumi.set(__self__, "resource", resource)
        pulumi.set(__self__, "service_name", service_name)
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
        Permission to give to this username on this resource.
        Available permissions:
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: pulumi.Input[str]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Input[str]:
        """
        Resource affected by this schema registry ACL.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        Username affected by this schema registry ACL.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _KafkaSchemaRegistryAclState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KafkaSchemaRegistryAcl resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] permission: Permission to give to this username on this resource.
               Available permissions:
        :param pulumi.Input[str] resource: Resource affected by this schema registry ACL.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] username: Username affected by this schema registry ACL.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if permission is not None:
            pulumi.set(__self__, "permission", permission)
        if resource is not None:
            pulumi.set(__self__, "resource", resource)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
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
        Permission to give to this username on this resource.
        Available permissions:
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter
    def resource(self) -> Optional[pulumi.Input[str]]:
        """
        Resource affected by this schema registry ACL.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Username affected by this schema registry ACL.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class KafkaSchemaRegistryAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a schema registry ACL for a Kafka cluster associated with a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        kafka = ovh.CloudProjectDatabase.get_database(service_name="XXX",
            engine="kafka",
            id="ZZZ")
        schema_registry_acl = ovh.cloud_project_database.KafkaSchemaRegistryAcl("schemaRegistryAcl",
            service_name=kafka.service_name,
            cluster_id=kafka.id,
            permission="schema_registry_read",
            resource="Subject:myResource",
            username="johndoe")
        ```

        ## Import

        OVHcloud Managed Kafka clusters schema registry ACLs can be imported using the `service_name`, `cluster_id` and `id` of the schema registry ACL, separated by "/" E.g., bash

        ```sh
         $ pulumi import ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl my_schemaRegistryAcl service_name/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] permission: Permission to give to this username on this resource.
               Available permissions:
        :param pulumi.Input[str] resource: Resource affected by this schema registry ACL.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] username: Username affected by this schema registry ACL.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KafkaSchemaRegistryAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a schema registry ACL for a Kafka cluster associated with a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        kafka = ovh.CloudProjectDatabase.get_database(service_name="XXX",
            engine="kafka",
            id="ZZZ")
        schema_registry_acl = ovh.cloud_project_database.KafkaSchemaRegistryAcl("schemaRegistryAcl",
            service_name=kafka.service_name,
            cluster_id=kafka.id,
            permission="schema_registry_read",
            resource="Subject:myResource",
            username="johndoe")
        ```

        ## Import

        OVHcloud Managed Kafka clusters schema registry ACLs can be imported using the `service_name`, `cluster_id` and `id` of the schema registry ACL, separated by "/" E.g., bash

        ```sh
         $ pulumi import ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl my_schemaRegistryAcl service_name/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param KafkaSchemaRegistryAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KafkaSchemaRegistryAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KafkaSchemaRegistryAclArgs.__new__(KafkaSchemaRegistryAclArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if permission is None and not opts.urn:
                raise TypeError("Missing required property 'permission'")
            __props__.__dict__["permission"] = permission
            if resource is None and not opts.urn:
                raise TypeError("Missing required property 'resource'")
            __props__.__dict__["resource"] = resource
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        super(KafkaSchemaRegistryAcl, __self__).__init__(
            'ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            permission: Optional[pulumi.Input[str]] = None,
            resource: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'KafkaSchemaRegistryAcl':
        """
        Get an existing KafkaSchemaRegistryAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] permission: Permission to give to this username on this resource.
               Available permissions:
        :param pulumi.Input[str] resource: Resource affected by this schema registry ACL.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] username: Username affected by this schema registry ACL.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KafkaSchemaRegistryAclState.__new__(_KafkaSchemaRegistryAclState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["permission"] = permission
        __props__.__dict__["resource"] = resource
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["username"] = username
        return KafkaSchemaRegistryAcl(resource_name, opts=opts, __props__=__props__)

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
        Permission to give to this username on this resource.
        Available permissions:
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Output[str]:
        """
        Resource affected by this schema registry ACL.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        Username affected by this schema registry ACL.
        """
        return pulumi.get(self, "username")

