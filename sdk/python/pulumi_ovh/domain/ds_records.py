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
from . import outputs
from ._inputs import *

__all__ = ['DSRecordsArgs', 'DSRecords']

@pulumi.input_type
class DSRecordsArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[builtins.str],
                 ds_records: pulumi.Input[Sequence[pulumi.Input['DSRecordsDsRecordArgs']]]):
        """
        The set of arguments for constructing a DSRecords resource.
        :param pulumi.Input[builtins.str] domain: Domain name for which to manage DS records
        :param pulumi.Input[Sequence[pulumi.Input['DSRecordsDsRecordArgs']]] ds_records: Details about a DS record
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "ds_records", ds_records)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[builtins.str]:
        """
        Domain name for which to manage DS records
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="dsRecords")
    def ds_records(self) -> pulumi.Input[Sequence[pulumi.Input['DSRecordsDsRecordArgs']]]:
        """
        Details about a DS record
        """
        return pulumi.get(self, "ds_records")

    @ds_records.setter
    def ds_records(self, value: pulumi.Input[Sequence[pulumi.Input['DSRecordsDsRecordArgs']]]):
        pulumi.set(self, "ds_records", value)


@pulumi.input_type
class _DSRecordsState:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[builtins.str]] = None,
                 ds_records: Optional[pulumi.Input[Sequence[pulumi.Input['DSRecordsDsRecordArgs']]]] = None):
        """
        Input properties used for looking up and filtering DSRecords resources.
        :param pulumi.Input[builtins.str] domain: Domain name for which to manage DS records
        :param pulumi.Input[Sequence[pulumi.Input['DSRecordsDsRecordArgs']]] ds_records: Details about a DS record
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if ds_records is not None:
            pulumi.set(__self__, "ds_records", ds_records)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Domain name for which to manage DS records
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="dsRecords")
    def ds_records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DSRecordsDsRecordArgs']]]]:
        """
        Details about a DS record
        """
        return pulumi.get(self, "ds_records")

    @ds_records.setter
    def ds_records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DSRecordsDsRecordArgs']]]]):
        pulumi.set(self, "ds_records", value)


class DSRecords(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[builtins.str]] = None,
                 ds_records: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DSRecordsDsRecordArgs', 'DSRecordsDsRecordArgsDict']]]]] = None,
                 __props__=None):
        """
        Use this resource to manage a domain's DS records.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        ds_records = ovh.domain.DSRecords("dsRecords",
            domain="mydomain.ovh",
            ds_records=[{
                "algorithm": "RSASHA1_NSEC3_SHA1",
                "flags": "KEY_SIGNING_KEY",
                "public_key": "my_base64_encoded_public_key",
                "tag": 12345,
            }])
        ```

        ## Import

        DS records can be imported using their `domain`.

        Using the following configuration:

        terraform

        import {

          to = ovh_domain_ds_records.ds_records

          id = "<domain name>"

        }

        You can then run:

        bash

        $ pulumi preview -generate-config-out=ds_records.tf

        $ pulumi up

        The file `ds_records.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] domain: Domain name for which to manage DS records
        :param pulumi.Input[Sequence[pulumi.Input[Union['DSRecordsDsRecordArgs', 'DSRecordsDsRecordArgsDict']]]] ds_records: Details about a DS record
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DSRecordsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to manage a domain's DS records.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        ds_records = ovh.domain.DSRecords("dsRecords",
            domain="mydomain.ovh",
            ds_records=[{
                "algorithm": "RSASHA1_NSEC3_SHA1",
                "flags": "KEY_SIGNING_KEY",
                "public_key": "my_base64_encoded_public_key",
                "tag": 12345,
            }])
        ```

        ## Import

        DS records can be imported using their `domain`.

        Using the following configuration:

        terraform

        import {

          to = ovh_domain_ds_records.ds_records

          id = "<domain name>"

        }

        You can then run:

        bash

        $ pulumi preview -generate-config-out=ds_records.tf

        $ pulumi up

        The file `ds_records.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.

        :param str resource_name: The name of the resource.
        :param DSRecordsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DSRecordsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[builtins.str]] = None,
                 ds_records: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DSRecordsDsRecordArgs', 'DSRecordsDsRecordArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DSRecordsArgs.__new__(DSRecordsArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            if ds_records is None and not opts.urn:
                raise TypeError("Missing required property 'ds_records'")
            __props__.__dict__["ds_records"] = ds_records
        super(DSRecords, __self__).__init__(
            'ovh:Domain/dSRecords:DSRecords',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain: Optional[pulumi.Input[builtins.str]] = None,
            ds_records: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DSRecordsDsRecordArgs', 'DSRecordsDsRecordArgsDict']]]]] = None) -> 'DSRecords':
        """
        Get an existing DSRecords resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] domain: Domain name for which to manage DS records
        :param pulumi.Input[Sequence[pulumi.Input[Union['DSRecordsDsRecordArgs', 'DSRecordsDsRecordArgsDict']]]] ds_records: Details about a DS record
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DSRecordsState.__new__(_DSRecordsState)

        __props__.__dict__["domain"] = domain
        __props__.__dict__["ds_records"] = ds_records
        return DSRecords(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[builtins.str]:
        """
        Domain name for which to manage DS records
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="dsRecords")
    def ds_records(self) -> pulumi.Output[Sequence['outputs.DSRecordsDsRecord']]:
        """
        Details about a DS record
        """
        return pulumi.get(self, "ds_records")

