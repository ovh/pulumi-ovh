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

__all__ = ['IpRestrictionArgs', 'IpRestriction']

@pulumi.input_type
class IpRestrictionArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[builtins.str],
                 engine: pulumi.Input[builtins.str],
                 ip: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a IpRestriction resource.
        :param pulumi.Input[builtins.str] cluster_id: Cluster ID.
        :param pulumi.Input[builtins.str] engine: The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[builtins.str] ip: Authorized IP.
        :param pulumi.Input[builtins.str] description: Description of the IP restriction.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "ip", ip)
        if description is not None:
            pulumi.set(__self__, "description", description)
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
        The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[builtins.str]:
        """
        Authorized IP.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the IP restriction.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

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
class _IpRestrictionState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 engine: Optional[pulumi.Input[builtins.str]] = None,
                 ip: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering IpRestriction resources.
        :param pulumi.Input[builtins.str] cluster_id: Cluster ID.
        :param pulumi.Input[builtins.str] description: Description of the IP restriction.
        :param pulumi.Input[builtins.str] engine: The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[builtins.str] ip: Authorized IP.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] status: Current status of the IP restriction.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the IP restriction.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Authorized IP.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ip", value)

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
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Current status of the IP restriction.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)


@pulumi.type_token("ovh:CloudProjectDatabase/ipRestriction:IpRestriction")
class IpRestriction(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 engine: Optional[pulumi.Input[builtins.str]] = None,
                 ip: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Deprecated: Use ip_restriction field in cloud_project_database resource instead. Continuing to use the CloudProjectDatabase.IpRestriction resource to add an IP restriction to a cloud_project_database resource will cause the cloud_project_database resource to be updated on every apply

        Apply IP restrictions to an OVHcloud Managed Database cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        db = ovh.CloudProjectDatabase.get_database(service_name="XXXX",
            engine="YYYY",
            id="ZZZZ")
        ip_restriction = ovh.cloud_project_database.IpRestriction("ip_restriction",
            service_name=db.service_name,
            engine=db.engine,
            cluster_id=db.id,
            ip="178.97.6.0/24")
        ```

        ## Import

        OVHcloud Managed database cluster IP restrictions can be imported using the `service_name`, `engine`, `cluster_id` and the `ip`, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/ipRestriction:IpRestriction my_ip_restriction service_name/engine/cluster_id/178.97.6.0/24
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_id: Cluster ID.
        :param pulumi.Input[builtins.str] description: Description of the IP restriction.
        :param pulumi.Input[builtins.str] engine: The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[builtins.str] ip: Authorized IP.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpRestrictionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Deprecated: Use ip_restriction field in cloud_project_database resource instead. Continuing to use the CloudProjectDatabase.IpRestriction resource to add an IP restriction to a cloud_project_database resource will cause the cloud_project_database resource to be updated on every apply

        Apply IP restrictions to an OVHcloud Managed Database cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        db = ovh.CloudProjectDatabase.get_database(service_name="XXXX",
            engine="YYYY",
            id="ZZZZ")
        ip_restriction = ovh.cloud_project_database.IpRestriction("ip_restriction",
            service_name=db.service_name,
            engine=db.engine,
            cluster_id=db.id,
            ip="178.97.6.0/24")
        ```

        ## Import

        OVHcloud Managed database cluster IP restrictions can be imported using the `service_name`, `engine`, `cluster_id` and the `ip`, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProjectDatabase/ipRestriction:IpRestriction my_ip_restriction service_name/engine/cluster_id/178.97.6.0/24
        ```

        :param str resource_name: The name of the resource.
        :param IpRestrictionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpRestrictionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 engine: Optional[pulumi.Input[builtins.str]] = None,
                 ip: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpRestrictionArgs.__new__(IpRestrictionArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["description"] = description
            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            if ip is None and not opts.urn:
                raise TypeError("Missing required property 'ip'")
            __props__.__dict__["ip"] = ip
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["status"] = None
        super(IpRestriction, __self__).__init__(
            'ovh:CloudProjectDatabase/ipRestriction:IpRestriction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            engine: Optional[pulumi.Input[builtins.str]] = None,
            ip: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None) -> 'IpRestriction':
        """
        Get an existing IpRestriction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_id: Cluster ID.
        :param pulumi.Input[builtins.str] description: Description of the IP restriction.
        :param pulumi.Input[builtins.str] engine: The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        :param pulumi.Input[builtins.str] ip: Authorized IP.
        :param pulumi.Input[builtins.str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[builtins.str] status: Current status of the IP restriction.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpRestrictionState.__new__(_IpRestrictionState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["description"] = description
        __props__.__dict__["engine"] = engine
        __props__.__dict__["ip"] = ip
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        return IpRestriction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[builtins.str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the IP restriction.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[builtins.str]:
        """
        The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[builtins.str]:
        """
        Authorized IP.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Current status of the IP restriction.
        """
        return pulumi.get(self, "status")

