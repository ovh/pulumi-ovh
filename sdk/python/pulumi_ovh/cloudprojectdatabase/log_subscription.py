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

__all__ = ['LogSubscriptionArgs', 'LogSubscription']

@pulumi.input_type
class LogSubscriptionArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[builtins.str],
                 engine: pulumi.Input[builtins.str],
                 kind: pulumi.Input[builtins.str],
                 stream_id: pulumi.Input[builtins.str],
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a LogSubscription resource.
        :param pulumi.Input[builtins.str] cluster_id: Cluster ID.
        :param pulumi.Input[builtins.str] engine: The database engine for which you want to manage a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[builtins.str] kind: Log kind name of this subscription.
        :param pulumi.Input[builtins.str] stream_id: Id of the target Log data platform stream.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "stream_id", stream_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[builtins.str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[builtins.str]:
        """
        The database engine for which you want to manage a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[builtins.str]:
        """
        Log kind name of this subscription.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="streamId")
    def stream_id(self) -> pulumi.Input[builtins.str]:
        """
        Id of the target Log data platform stream.
        """
        return pulumi.get(self, "stream_id")

    @stream_id.setter
    def stream_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "stream_id", value)

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
class _LogSubscriptionState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 engine: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 ldp_service_name: Optional[pulumi.Input[builtins.str]] = None,
                 operation_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 stream_id: Optional[pulumi.Input[builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering LogSubscription resources.
        :param pulumi.Input[builtins.str] cluster_id: Cluster ID.
        :param pulumi.Input[builtins.str] created_at: Creation date of the subscription.
        :param pulumi.Input[builtins.str] engine: The database engine for which you want to manage a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[builtins.str] kind: Log kind name of this subscription.
        :param pulumi.Input[builtins.str] ldp_service_name: Name of the destination log service.
        :param pulumi.Input[builtins.str] operation_id: Identifier of the operation.
        :param pulumi.Input[builtins.str] resource_name: Name of subscribed resource, where the logs come from.
        :param pulumi.Input[builtins.str] resource_type: Type of subscribed resource, where the logs come from.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] stream_id: Id of the target Log data platform stream.
        :param pulumi.Input[builtins.str] updated_at: Last update date of the subscription.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if ldp_service_name is not None:
            pulumi.set(__self__, "ldp_service_name", ldp_service_name)
        if operation_id is not None:
            pulumi.set(__self__, "operation_id", operation_id)
        if resource_name is not None:
            pulumi.set(__self__, "resource_name", resource_name)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if stream_id is not None:
            pulumi.set(__self__, "stream_id", stream_id)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Creation date of the subscription.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The database engine for which you want to manage a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Log kind name of this subscription.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="ldpServiceName")
    def ldp_service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the destination log service.
        """
        return pulumi.get(self, "ldp_service_name")

    @ldp_service_name.setter
    def ldp_service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ldp_service_name", value)

    @property
    @pulumi.getter(name="operationId")
    def operation_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Identifier of the operation.
        """
        return pulumi.get(self, "operation_id")

    @operation_id.setter
    def operation_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "operation_id", value)

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of subscribed resource, where the logs come from.
        """
        return pulumi.get(self, "resource_name")

    @resource_name.setter
    def resource_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_name", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Type of subscribed resource, where the logs come from.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_type", value)

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

    @property
    @pulumi.getter(name="streamId")
    def stream_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Id of the target Log data platform stream.
        """
        return pulumi.get(self, "stream_id")

    @stream_id.setter
    def stream_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "stream_id", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Last update date of the subscription.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)


@pulumi.type_token("ovh:CloudProjectDatabase/logSubscription:LogSubscription")
class LogSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 engine: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 stream_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates a log subscription for a cluster associated with a public cloud project.

        ## Example Usage

        Create a log subscription for a database.

        ```python
        import pulumi
        import pulumi_ovh as ovh

        stream = ovh.Dbaas.get_logs_output_graylog_stream(service_name="ldp-xx-xxxxx",
            title="my stream")
        db = ovh.CloudProjectDatabase.get_database(service_name="XXX",
            engine="YYY",
            id="ZZZ")
        subscription = ovh.cloud_project_database.LogSubscription("subscription",
            service_name=db.service_name,
            engine=db.engine,
            cluster_id=db.id,
            stream_id=stream.id,
            kind="customer_logs")
        ```

        ## Import

        OVHcloud Managed clusters logs subscription can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the subscription, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/logSubscription:LogSubscription sub service_name/engine/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_id: Cluster ID.
        :param pulumi.Input[builtins.str] engine: The database engine for which you want to manage a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[builtins.str] kind: Log kind name of this subscription.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] stream_id: Id of the target Log data platform stream.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a log subscription for a cluster associated with a public cloud project.

        ## Example Usage

        Create a log subscription for a database.

        ```python
        import pulumi
        import pulumi_ovh as ovh

        stream = ovh.Dbaas.get_logs_output_graylog_stream(service_name="ldp-xx-xxxxx",
            title="my stream")
        db = ovh.CloudProjectDatabase.get_database(service_name="XXX",
            engine="YYY",
            id="ZZZ")
        subscription = ovh.cloud_project_database.LogSubscription("subscription",
            service_name=db.service_name,
            engine=db.engine,
            cluster_id=db.id,
            stream_id=stream.id,
            kind="customer_logs")
        ```

        ## Import

        OVHcloud Managed clusters logs subscription can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the subscription, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/logSubscription:LogSubscription sub service_name/engine/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param LogSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 engine: Optional[pulumi.Input[builtins.str]] = None,
                 kind: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 stream_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogSubscriptionArgs.__new__(LogSubscriptionArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = kind
            __props__.__dict__["service_name"] = service_name
            if stream_id is None and not opts.urn:
                raise TypeError("Missing required property 'stream_id'")
            __props__.__dict__["stream_id"] = stream_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["ldp_service_name"] = None
            __props__.__dict__["operation_id"] = None
            __props__.__dict__["resource_name"] = None
            __props__.__dict__["resource_type"] = None
            __props__.__dict__["updated_at"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["ldpServiceName"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(LogSubscription, __self__).__init__(
            'ovh:CloudProjectDatabase/logSubscription:LogSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[builtins.str]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            engine: Optional[pulumi.Input[builtins.str]] = None,
            kind: Optional[pulumi.Input[builtins.str]] = None,
            ldp_service_name: Optional[pulumi.Input[builtins.str]] = None,
            operation_id: Optional[pulumi.Input[builtins.str]] = None,
            resource_name_: Optional[pulumi.Input[builtins.str]] = None,
            resource_type: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            stream_id: Optional[pulumi.Input[builtins.str]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None) -> 'LogSubscription':
        """
        Get an existing LogSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_id: Cluster ID.
        :param pulumi.Input[builtins.str] created_at: Creation date of the subscription.
        :param pulumi.Input[builtins.str] engine: The database engine for which you want to manage a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[builtins.str] kind: Log kind name of this subscription.
        :param pulumi.Input[builtins.str] ldp_service_name: Name of the destination log service.
        :param pulumi.Input[builtins.str] operation_id: Identifier of the operation.
        :param pulumi.Input[builtins.str] resource_name_: Name of subscribed resource, where the logs come from.
        :param pulumi.Input[builtins.str] resource_type: Type of subscribed resource, where the logs come from.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] stream_id: Id of the target Log data platform stream.
        :param pulumi.Input[builtins.str] updated_at: Last update date of the subscription.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogSubscriptionState.__new__(_LogSubscriptionState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["engine"] = engine
        __props__.__dict__["kind"] = kind
        __props__.__dict__["ldp_service_name"] = ldp_service_name
        __props__.__dict__["operation_id"] = operation_id
        __props__.__dict__["resource_name"] = resource_name_
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["stream_id"] = stream_id
        __props__.__dict__["updated_at"] = updated_at
        return LogSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[builtins.str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        Creation date of the subscription.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[builtins.str]:
        """
        The database engine for which you want to manage a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[builtins.str]:
        """
        Log kind name of this subscription.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="ldpServiceName")
    def ldp_service_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the destination log service.
        """
        return pulumi.get(self, "ldp_service_name")

    @property
    @pulumi.getter(name="operationId")
    def operation_id(self) -> pulumi.Output[builtins.str]:
        """
        Identifier of the operation.
        """
        return pulumi.get(self, "operation_id")

    @property
    @pulumi.getter(name="resourceName")
    def resource_name(self) -> pulumi.Output[builtins.str]:
        """
        Name of subscribed resource, where the logs come from.
        """
        return pulumi.get(self, "resource_name")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[builtins.str]:
        """
        Type of subscribed resource, where the logs come from.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="streamId")
    def stream_id(self) -> pulumi.Output[builtins.str]:
        """
        Id of the target Log data platform stream.
        """
        return pulumi.get(self, "stream_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        Last update date of the subscription.
        """
        return pulumi.get(self, "updated_at")

