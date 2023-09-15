# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ZoneRecordArgs', 'ZoneRecord']

@pulumi.input_type
class ZoneRecordArgs:
    def __init__(__self__, *,
                 fieldtype: pulumi.Input[str],
                 target: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 subdomain: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ZoneRecord resource.
        :param pulumi.Input[str] fieldtype: The type of the record
        :param pulumi.Input[str] target: The value of the record
        :param pulumi.Input[str] zone: The domain to add the record to
        :param pulumi.Input[str] subdomain: The name of the record. It can be an empty string.
        :param pulumi.Input[int] ttl: The TTL of the record, it shall be >= to 60.
        """
        pulumi.set(__self__, "fieldtype", fieldtype)
        pulumi.set(__self__, "target", target)
        pulumi.set(__self__, "zone", zone)
        if subdomain is not None:
            pulumi.set(__self__, "subdomain", subdomain)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def fieldtype(self) -> pulumi.Input[str]:
        """
        The type of the record
        """
        return pulumi.get(self, "fieldtype")

    @fieldtype.setter
    def fieldtype(self, value: pulumi.Input[str]):
        pulumi.set(self, "fieldtype", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        """
        The value of the record
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        The domain to add the record to
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter
    def subdomain(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the record. It can be an empty string.
        """
        return pulumi.get(self, "subdomain")

    @subdomain.setter
    def subdomain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdomain", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The TTL of the record, it shall be >= to 60.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class _ZoneRecordState:
    def __init__(__self__, *,
                 fieldtype: Optional[pulumi.Input[str]] = None,
                 subdomain: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ZoneRecord resources.
        :param pulumi.Input[str] fieldtype: The type of the record
        :param pulumi.Input[str] subdomain: The name of the record. It can be an empty string.
        :param pulumi.Input[str] target: The value of the record
        :param pulumi.Input[int] ttl: The TTL of the record, it shall be >= to 60.
        :param pulumi.Input[str] zone: The domain to add the record to
        """
        if fieldtype is not None:
            pulumi.set(__self__, "fieldtype", fieldtype)
        if subdomain is not None:
            pulumi.set(__self__, "subdomain", subdomain)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def fieldtype(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the record
        """
        return pulumi.get(self, "fieldtype")

    @fieldtype.setter
    def fieldtype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fieldtype", value)

    @property
    @pulumi.getter
    def subdomain(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the record. It can be an empty string.
        """
        return pulumi.get(self, "subdomain")

    @subdomain.setter
    def subdomain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdomain", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        The value of the record
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The TTL of the record, it shall be >= to 60.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The domain to add the record to
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class ZoneRecord(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fieldtype: Optional[pulumi.Input[str]] = None,
                 subdomain: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        # Add a record to a sub-domain
        test = ovh.domain.ZoneRecord("test",
            fieldtype="A",
            subdomain="test",
            target="0.0.0.0",
            ttl=3600,
            zone="testdemo.ovh")
        ```

        ## Import

        OVHcloud domain zone record can be imported using the `id`, which can be retrieved by using [OVH API portal](https://api.ovh.com/console/#/domain/zone/%7BzoneName%7D/record~GET), and the `zone`, separated by "." E.g., bash

        ```sh
         $ pulumi import ovh:Domain/zoneRecord:ZoneRecord test id.zone
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fieldtype: The type of the record
        :param pulumi.Input[str] subdomain: The name of the record. It can be an empty string.
        :param pulumi.Input[str] target: The value of the record
        :param pulumi.Input[int] ttl: The TTL of the record, it shall be >= to 60.
        :param pulumi.Input[str] zone: The domain to add the record to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ZoneRecordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        # Add a record to a sub-domain
        test = ovh.domain.ZoneRecord("test",
            fieldtype="A",
            subdomain="test",
            target="0.0.0.0",
            ttl=3600,
            zone="testdemo.ovh")
        ```

        ## Import

        OVHcloud domain zone record can be imported using the `id`, which can be retrieved by using [OVH API portal](https://api.ovh.com/console/#/domain/zone/%7BzoneName%7D/record~GET), and the `zone`, separated by "." E.g., bash

        ```sh
         $ pulumi import ovh:Domain/zoneRecord:ZoneRecord test id.zone
        ```

        :param str resource_name: The name of the resource.
        :param ZoneRecordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneRecordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 fieldtype: Optional[pulumi.Input[str]] = None,
                 subdomain: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZoneRecordArgs.__new__(ZoneRecordArgs)

            if fieldtype is None and not opts.urn:
                raise TypeError("Missing required property 'fieldtype'")
            __props__.__dict__["fieldtype"] = fieldtype
            __props__.__dict__["subdomain"] = subdomain
            if target is None and not opts.urn:
                raise TypeError("Missing required property 'target'")
            __props__.__dict__["target"] = target
            __props__.__dict__["ttl"] = ttl
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(ZoneRecord, __self__).__init__(
            'ovh:Domain/zoneRecord:ZoneRecord',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            fieldtype: Optional[pulumi.Input[str]] = None,
            subdomain: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[int]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'ZoneRecord':
        """
        Get an existing ZoneRecord resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] fieldtype: The type of the record
        :param pulumi.Input[str] subdomain: The name of the record. It can be an empty string.
        :param pulumi.Input[str] target: The value of the record
        :param pulumi.Input[int] ttl: The TTL of the record, it shall be >= to 60.
        :param pulumi.Input[str] zone: The domain to add the record to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneRecordState.__new__(_ZoneRecordState)

        __props__.__dict__["fieldtype"] = fieldtype
        __props__.__dict__["subdomain"] = subdomain
        __props__.__dict__["target"] = target
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["zone"] = zone
        return ZoneRecord(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fieldtype(self) -> pulumi.Output[str]:
        """
        The type of the record
        """
        return pulumi.get(self, "fieldtype")

    @property
    @pulumi.getter
    def subdomain(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the record. It can be an empty string.
        """
        return pulumi.get(self, "subdomain")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[str]:
        """
        The value of the record
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[int]]:
        """
        The TTL of the record, it shall be >= to 60.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The domain to add the record to
        """
        return pulumi.get(self, "zone")

