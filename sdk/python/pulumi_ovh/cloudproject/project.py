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

__all__ = ['ProjectArgs', 'Project']

@pulumi.input_type
class ProjectArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectOrderArgs']]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['ProjectPlanArgs']] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectPlanOptionArgs']]]] = None):
        """
        The set of arguments for constructing a Project resource.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectOrderArgs']]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: Ovh Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input['ProjectPlanArgs'] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input['ProjectPlanOptionArgs']]] plan_options: Product Plan to order
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if orders is not None:
            pulumi.set(__self__, "orders", orders)
        if ovh_subsidiary is not None:
            pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        if payment_mean is not None:
            warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
            pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
        if payment_mean is not None:
            pulumi.set(__self__, "payment_mean", payment_mean)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if plan_options is not None:
            pulumi.set(__self__, "plan_options", plan_options)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def orders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectOrderArgs']]]]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @orders.setter
    def orders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectOrderArgs']]]]):
        pulumi.set(self, "orders", value)

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @ovh_subsidiary.setter
    def ovh_subsidiary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ovh_subsidiary", value)

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @payment_mean.setter
    def payment_mean(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_mean", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['ProjectPlanArgs']]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['ProjectPlanArgs']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectPlanOptionArgs']]]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @plan_options.setter
    def plan_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectPlanOptionArgs']]]]):
        pulumi.set(self, "plan_options", value)


@pulumi.input_type
class _ProjectState:
    def __init__(__self__, *,
                 project_urn: Optional[pulumi.Input[str]] = None,
                 access: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectOrderArgs']]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input['ProjectPlanArgs']] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectPlanOptionArgs']]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Project resources.
        :param pulumi.Input[Sequence[pulumi.Input['ProjectOrderArgs']]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: Ovh Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input['ProjectPlanArgs'] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input['ProjectPlanOptionArgs']]] plan_options: Product Plan to order
        """
        if project_urn is not None:
            pulumi.set(__self__, "project_urn", project_urn)
        if access is not None:
            pulumi.set(__self__, "access", access)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if orders is not None:
            pulumi.set(__self__, "orders", orders)
        if ovh_subsidiary is not None:
            pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        if payment_mean is not None:
            warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
            pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
        if payment_mean is not None:
            pulumi.set(__self__, "payment_mean", payment_mean)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if plan_options is not None:
            pulumi.set(__self__, "plan_options", plan_options)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="ProjectURN")
    def project_urn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_urn")

    @project_urn.setter
    def project_urn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_urn", value)

    @property
    @pulumi.getter
    def access(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access")

    @access.setter
    def access(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def orders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectOrderArgs']]]]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @orders.setter
    def orders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectOrderArgs']]]]):
        pulumi.set(self, "orders", value)

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @ovh_subsidiary.setter
    def ovh_subsidiary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ovh_subsidiary", value)

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> Optional[pulumi.Input[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @payment_mean.setter
    def payment_mean(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payment_mean", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['ProjectPlanArgs']]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['ProjectPlanArgs']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProjectPlanOptionArgs']]]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @plan_options.setter
    def plan_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProjectPlanOptionArgs']]]]):
        pulumi.set(self, "plan_options", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Project(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectOrderArgs', 'ProjectOrderArgsDict']]]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[Union['ProjectPlanArgs', 'ProjectPlanArgsDict']]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectPlanOptionArgs', 'ProjectPlanOptionArgsDict']]]]] = None,
                 __props__=None):
        """
        Create a Project resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProjectOrderArgs', 'ProjectOrderArgsDict']]]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: Ovh Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input[Union['ProjectPlanArgs', 'ProjectPlanArgsDict']] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProjectPlanOptionArgs', 'ProjectPlanOptionArgsDict']]]] plan_options: Product Plan to order
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProjectArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Project resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectOrderArgs', 'ProjectOrderArgsDict']]]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[str]] = None,
                 payment_mean: Optional[pulumi.Input[str]] = None,
                 plan: Optional[pulumi.Input[Union['ProjectPlanArgs', 'ProjectPlanArgsDict']]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectPlanOptionArgs', 'ProjectPlanOptionArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectArgs.__new__(ProjectArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["orders"] = orders
            __props__.__dict__["ovh_subsidiary"] = ovh_subsidiary
            __props__.__dict__["payment_mean"] = payment_mean
            __props__.__dict__["plan"] = plan
            __props__.__dict__["plan_options"] = plan_options
            __props__.__dict__["project_urn"] = None
            __props__.__dict__["access"] = None
            __props__.__dict__["project_id"] = None
            __props__.__dict__["project_name"] = None
            __props__.__dict__["status"] = None
        super(Project, __self__).__init__(
            'ovh:CloudProject/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project_urn: Optional[pulumi.Input[str]] = None,
            access: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectOrderArgs', 'ProjectOrderArgsDict']]]]] = None,
            ovh_subsidiary: Optional[pulumi.Input[str]] = None,
            payment_mean: Optional[pulumi.Input[str]] = None,
            plan: Optional[pulumi.Input[Union['ProjectPlanArgs', 'ProjectPlanArgsDict']]] = None,
            plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ProjectPlanOptionArgs', 'ProjectPlanOptionArgsDict']]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProjectOrderArgs', 'ProjectOrderArgsDict']]]] orders: Details about an Order
        :param pulumi.Input[str] ovh_subsidiary: Ovh Subsidiary
        :param pulumi.Input[str] payment_mean: Ovh payment mode
        :param pulumi.Input[Union['ProjectPlanArgs', 'ProjectPlanArgsDict']] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input[Union['ProjectPlanOptionArgs', 'ProjectPlanOptionArgsDict']]]] plan_options: Product Plan to order
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectState.__new__(_ProjectState)

        __props__.__dict__["project_urn"] = project_urn
        __props__.__dict__["access"] = access
        __props__.__dict__["description"] = description
        __props__.__dict__["orders"] = orders
        __props__.__dict__["ovh_subsidiary"] = ovh_subsidiary
        __props__.__dict__["payment_mean"] = payment_mean
        __props__.__dict__["plan"] = plan
        __props__.__dict__["plan_options"] = plan_options
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["status"] = status
        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ProjectURN")
    def project_urn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_urn")

    @property
    @pulumi.getter
    def access(self) -> pulumi.Output[str]:
        return pulumi.get(self, "access")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def orders(self) -> pulumi.Output[Sequence['outputs.ProjectOrder']]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> pulumi.Output[str]:
        """
        Ovh Subsidiary
        """
        return pulumi.get(self, "ovh_subsidiary")

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> pulumi.Output[Optional[str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output['outputs.ProjectPlan']:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> pulumi.Output[Optional[Sequence['outputs.ProjectPlanOption']]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

