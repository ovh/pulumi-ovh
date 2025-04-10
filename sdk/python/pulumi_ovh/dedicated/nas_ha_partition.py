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

__all__ = ['NasHAPartitionArgs', 'NasHAPartition']

@pulumi.input_type
class NasHAPartitionArgs:
    def __init__(__self__, *,
                 protocol: pulumi.Input[builtins.str],
                 service_name: pulumi.Input[builtins.str],
                 size: pulumi.Input[builtins.int],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a NasHAPartition resource.
        :param pulumi.Input[builtins.str] protocol: one of "NFS", "CIFS" or "NFS_CIFS"
        :param pulumi.Input[builtins.str] service_name: The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        :param pulumi.Input[builtins.int] size: size of the partition in GB
        :param pulumi.Input[builtins.str] description: A brief description of the partition
        :param pulumi.Input[builtins.str] name: name of the partition
        """
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "size", size)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[builtins.str]:
        """
        one of "NFS", "CIFS" or "NFS_CIFS"
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[builtins.int]:
        """
        size of the partition in GB
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A brief description of the partition
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        name of the partition
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _NasHAPartitionState:
    def __init__(__self__, *,
                 capacity: Optional[pulumi.Input[builtins.int]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 size: Optional[pulumi.Input[builtins.int]] = None,
                 used_by_snapshots: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering NasHAPartition resources.
        :param pulumi.Input[builtins.int] capacity: Percentage of partition space used in %
        :param pulumi.Input[builtins.str] description: A brief description of the partition
        :param pulumi.Input[builtins.str] name: name of the partition
        :param pulumi.Input[builtins.str] protocol: one of "NFS", "CIFS" or "NFS_CIFS"
        :param pulumi.Input[builtins.str] service_name: The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        :param pulumi.Input[builtins.int] size: size of the partition in GB
        :param pulumi.Input[builtins.int] used_by_snapshots: Percentage of partition space used by snapshots in %
        """
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if used_by_snapshots is not None:
            pulumi.set(__self__, "used_by_snapshots", used_by_snapshots)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Percentage of partition space used in %
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A brief description of the partition
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        name of the partition
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        one of "NFS", "CIFS" or "NFS_CIFS"
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        size of the partition in GB
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="usedBySnapshots")
    def used_by_snapshots(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Percentage of partition space used by snapshots in %
        """
        return pulumi.get(self, "used_by_snapshots")

    @used_by_snapshots.setter
    def used_by_snapshots(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "used_by_snapshots", value)


class NasHAPartition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 size: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Provides a resource for managing partitions on HA-NAS services

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_partition = ovh.dedicated.NasHAPartition("myPartition",
            protocol="NFS",
            service_name="zpool-12345",
            size=20)
        ```

        ## Import

        HA-NAS can be imported using the `{service_name}/{name}`, e.g.

        ```sh
        $ pulumi import ovh:Dedicated/nasHAPartition:NasHAPartition my-partition zpool-12345/my-partition`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: A brief description of the partition
        :param pulumi.Input[builtins.str] name: name of the partition
        :param pulumi.Input[builtins.str] protocol: one of "NFS", "CIFS" or "NFS_CIFS"
        :param pulumi.Input[builtins.str] service_name: The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        :param pulumi.Input[builtins.int] size: size of the partition in GB
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NasHAPartitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource for managing partitions on HA-NAS services

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_partition = ovh.dedicated.NasHAPartition("myPartition",
            protocol="NFS",
            service_name="zpool-12345",
            size=20)
        ```

        ## Import

        HA-NAS can be imported using the `{service_name}/{name}`, e.g.

        ```sh
        $ pulumi import ovh:Dedicated/nasHAPartition:NasHAPartition my-partition zpool-12345/my-partition`
        ```

        :param str resource_name: The name of the resource.
        :param NasHAPartitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NasHAPartitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 protocol: Optional[pulumi.Input[builtins.str]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 size: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NasHAPartitionArgs.__new__(NasHAPartitionArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__.__dict__["size"] = size
            __props__.__dict__["capacity"] = None
            __props__.__dict__["used_by_snapshots"] = None
        super(NasHAPartition, __self__).__init__(
            'ovh:Dedicated/nasHAPartition:NasHAPartition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            capacity: Optional[pulumi.Input[builtins.int]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            protocol: Optional[pulumi.Input[builtins.str]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            size: Optional[pulumi.Input[builtins.int]] = None,
            used_by_snapshots: Optional[pulumi.Input[builtins.int]] = None) -> 'NasHAPartition':
        """
        Get an existing NasHAPartition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.int] capacity: Percentage of partition space used in %
        :param pulumi.Input[builtins.str] description: A brief description of the partition
        :param pulumi.Input[builtins.str] name: name of the partition
        :param pulumi.Input[builtins.str] protocol: one of "NFS", "CIFS" or "NFS_CIFS"
        :param pulumi.Input[builtins.str] service_name: The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        :param pulumi.Input[builtins.int] size: size of the partition in GB
        :param pulumi.Input[builtins.int] used_by_snapshots: Percentage of partition space used by snapshots in %
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NasHAPartitionState.__new__(_NasHAPartitionState)

        __props__.__dict__["capacity"] = capacity
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["size"] = size
        __props__.__dict__["used_by_snapshots"] = used_by_snapshots
        return NasHAPartition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[builtins.int]:
        """
        Percentage of partition space used in %
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A brief description of the partition
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        name of the partition
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[builtins.str]:
        """
        one of "NFS", "CIFS" or "NFS_CIFS"
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[builtins.int]:
        """
        size of the partition in GB
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="usedBySnapshots")
    def used_by_snapshots(self) -> pulumi.Output[builtins.int]:
        """
        Percentage of partition space used by snapshots in %
        """
        return pulumi.get(self, "used_by_snapshots")

