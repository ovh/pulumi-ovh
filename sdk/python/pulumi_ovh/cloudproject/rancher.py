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
from . import outputs
from ._inputs import *

__all__ = ['RancherArgs', 'Rancher']

@pulumi.input_type
class RancherArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 target_spec: pulumi.Input['RancherTargetSpecArgs']):
        """
        The set of arguments for constructing a Rancher resource.
        :param pulumi.Input[str] project_id: Project ID
        :param pulumi.Input['RancherTargetSpecArgs'] target_spec: Target specification for the managed Rancher service
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "target_spec", target_spec)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="targetSpec")
    def target_spec(self) -> pulumi.Input['RancherTargetSpecArgs']:
        """
        Target specification for the managed Rancher service
        """
        return pulumi.get(self, "target_spec")

    @target_spec.setter
    def target_spec(self, value: pulumi.Input['RancherTargetSpecArgs']):
        pulumi.set(self, "target_spec", value)


@pulumi.input_type
class _RancherState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 current_state: Optional[pulumi.Input['RancherCurrentStateArgs']] = None,
                 current_tasks: Optional[pulumi.Input[Sequence[pulumi.Input['RancherCurrentTaskArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 resource_status: Optional[pulumi.Input[str]] = None,
                 target_spec: Optional[pulumi.Input['RancherTargetSpecArgs']] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Rancher resources.
        :param pulumi.Input[str] created_at: Date of the managed Rancher service creation
        :param pulumi.Input['RancherCurrentStateArgs'] current_state: Current configuration applied to the managed Rancher service
        :param pulumi.Input[Sequence[pulumi.Input['RancherCurrentTaskArgs']]] current_tasks: Asynchronous operations ongoing on the managed Rancher service
        :param pulumi.Input[str] project_id: Project ID
        :param pulumi.Input[str] resource_status: Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
        :param pulumi.Input['RancherTargetSpecArgs'] target_spec: Target specification for the managed Rancher service
        :param pulumi.Input[str] updated_at: Date of the last managed Rancher service update
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if current_state is not None:
            pulumi.set(__self__, "current_state", current_state)
        if current_tasks is not None:
            pulumi.set(__self__, "current_tasks", current_tasks)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if resource_status is not None:
            pulumi.set(__self__, "resource_status", resource_status)
        if target_spec is not None:
            pulumi.set(__self__, "target_spec", target_spec)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date of the managed Rancher service creation
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="currentState")
    def current_state(self) -> Optional[pulumi.Input['RancherCurrentStateArgs']]:
        """
        Current configuration applied to the managed Rancher service
        """
        return pulumi.get(self, "current_state")

    @current_state.setter
    def current_state(self, value: Optional[pulumi.Input['RancherCurrentStateArgs']]):
        pulumi.set(self, "current_state", value)

    @property
    @pulumi.getter(name="currentTasks")
    def current_tasks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RancherCurrentTaskArgs']]]]:
        """
        Asynchronous operations ongoing on the managed Rancher service
        """
        return pulumi.get(self, "current_tasks")

    @current_tasks.setter
    def current_tasks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RancherCurrentTaskArgs']]]]):
        pulumi.set(self, "current_tasks", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> Optional[pulumi.Input[str]]:
        """
        Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
        """
        return pulumi.get(self, "resource_status")

    @resource_status.setter
    def resource_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_status", value)

    @property
    @pulumi.getter(name="targetSpec")
    def target_spec(self) -> Optional[pulumi.Input['RancherTargetSpecArgs']]:
        """
        Target specification for the managed Rancher service
        """
        return pulumi.get(self, "target_spec")

    @target_spec.setter
    def target_spec(self, value: Optional[pulumi.Input['RancherTargetSpecArgs']]):
        pulumi.set(self, "target_spec", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date of the last managed Rancher service update
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class Rancher(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 target_spec: Optional[pulumi.Input[Union['RancherTargetSpecArgs', 'RancherTargetSpecArgsDict']]] = None,
                 __props__=None):
        """
        Manage a Rancher service in a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        rancher = ovh.cloud_project.Rancher("rancher",
            project_id="<public cloud project ID>",
            target_spec={
                "name": "MyRancher",
                "plan": "STANDARD",
            })
        ```

        ## Import

        A share in a public cloud project can be imported using the `project_id` and `id` attributes.

        Using the following configuration:

        hcl

        import {

          id = "<project_id>/<id>"

          to = ovh_cloud_project_rancher.rancher

        }

        You can then run:

        bash

        $ pulumi preview -generate-config-out=rancher.tf

        $ pulumi up

        The file `rancher.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.

        See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: Project ID
        :param pulumi.Input[Union['RancherTargetSpecArgs', 'RancherTargetSpecArgsDict']] target_spec: Target specification for the managed Rancher service
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RancherArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a Rancher service in a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        rancher = ovh.cloud_project.Rancher("rancher",
            project_id="<public cloud project ID>",
            target_spec={
                "name": "MyRancher",
                "plan": "STANDARD",
            })
        ```

        ## Import

        A share in a public cloud project can be imported using the `project_id` and `id` attributes.

        Using the following configuration:

        hcl

        import {

          id = "<project_id>/<id>"

          to = ovh_cloud_project_rancher.rancher

        }

        You can then run:

        bash

        $ pulumi preview -generate-config-out=rancher.tf

        $ pulumi up

        The file `rancher.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.

        See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.

        :param str resource_name: The name of the resource.
        :param RancherArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RancherArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 target_spec: Optional[pulumi.Input[Union['RancherTargetSpecArgs', 'RancherTargetSpecArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RancherArgs.__new__(RancherArgs)

            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if target_spec is None and not opts.urn:
                raise TypeError("Missing required property 'target_spec'")
            __props__.__dict__["target_spec"] = target_spec
            __props__.__dict__["created_at"] = None
            __props__.__dict__["current_state"] = None
            __props__.__dict__["current_tasks"] = None
            __props__.__dict__["resource_status"] = None
            __props__.__dict__["updated_at"] = None
        super(Rancher, __self__).__init__(
            'ovh:CloudProject/rancher:Rancher',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            current_state: Optional[pulumi.Input[Union['RancherCurrentStateArgs', 'RancherCurrentStateArgsDict']]] = None,
            current_tasks: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RancherCurrentTaskArgs', 'RancherCurrentTaskArgsDict']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            resource_status: Optional[pulumi.Input[str]] = None,
            target_spec: Optional[pulumi.Input[Union['RancherTargetSpecArgs', 'RancherTargetSpecArgsDict']]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'Rancher':
        """
        Get an existing Rancher resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: Date of the managed Rancher service creation
        :param pulumi.Input[Union['RancherCurrentStateArgs', 'RancherCurrentStateArgsDict']] current_state: Current configuration applied to the managed Rancher service
        :param pulumi.Input[Sequence[pulumi.Input[Union['RancherCurrentTaskArgs', 'RancherCurrentTaskArgsDict']]]] current_tasks: Asynchronous operations ongoing on the managed Rancher service
        :param pulumi.Input[str] project_id: Project ID
        :param pulumi.Input[str] resource_status: Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
        :param pulumi.Input[Union['RancherTargetSpecArgs', 'RancherTargetSpecArgsDict']] target_spec: Target specification for the managed Rancher service
        :param pulumi.Input[str] updated_at: Date of the last managed Rancher service update
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RancherState.__new__(_RancherState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["current_state"] = current_state
        __props__.__dict__["current_tasks"] = current_tasks
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["resource_status"] = resource_status
        __props__.__dict__["target_spec"] = target_spec
        __props__.__dict__["updated_at"] = updated_at
        return Rancher(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date of the managed Rancher service creation
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currentState")
    def current_state(self) -> pulumi.Output['outputs.RancherCurrentState']:
        """
        Current configuration applied to the managed Rancher service
        """
        return pulumi.get(self, "current_state")

    @property
    @pulumi.getter(name="currentTasks")
    def current_tasks(self) -> pulumi.Output[Sequence['outputs.RancherCurrentTask']]:
        """
        Asynchronous operations ongoing on the managed Rancher service
        """
        return pulumi.get(self, "current_tasks")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> pulumi.Output[str]:
        """
        Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
        """
        return pulumi.get(self, "resource_status")

    @property
    @pulumi.getter(name="targetSpec")
    def target_spec(self) -> pulumi.Output['outputs.RancherTargetSpec']:
        """
        Target specification for the managed Rancher service
        """
        return pulumi.get(self, "target_spec")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date of the last managed Rancher service update
        """
        return pulumi.get(self, "updated_at")

