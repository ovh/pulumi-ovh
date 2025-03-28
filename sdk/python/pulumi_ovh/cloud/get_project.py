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

__all__ = [
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, access=None, creation_date=None, description=None, expiration=None, iam=None, id=None, manual_quota=None, order_id=None, plan_code=None, project_id=None, project_name=None, service_name=None, status=None, unleash=None):
        if access and not isinstance(access, str):
            raise TypeError("Expected argument 'access' to be a str")
        pulumi.set(__self__, "access", access)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if expiration and not isinstance(expiration, str):
            raise TypeError("Expected argument 'expiration' to be a str")
        pulumi.set(__self__, "expiration", expiration)
        if iam and not isinstance(iam, dict):
            raise TypeError("Expected argument 'iam' to be a dict")
        pulumi.set(__self__, "iam", iam)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if manual_quota and not isinstance(manual_quota, bool):
            raise TypeError("Expected argument 'manual_quota' to be a bool")
        pulumi.set(__self__, "manual_quota", manual_quota)
        if order_id and not isinstance(order_id, float):
            raise TypeError("Expected argument 'order_id' to be a float")
        pulumi.set(__self__, "order_id", order_id)
        if plan_code and not isinstance(plan_code, str):
            raise TypeError("Expected argument 'plan_code' to be a str")
        pulumi.set(__self__, "plan_code", plan_code)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if unleash and not isinstance(unleash, bool):
            raise TypeError("Expected argument 'unleash' to be a bool")
        pulumi.set(__self__, "unleash", unleash)

    @property
    @pulumi.getter
    def access(self) -> str:
        """
        Project access
        """
        return pulumi.get(self, "access")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        Project creation date
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of your project
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def expiration(self) -> str:
        """
        Expiration date of your project. After this date, your project will be deleted
        """
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter
    def iam(self) -> 'outputs.GetProjectIamResult':
        """
        IAM resource information
        """
        return pulumi.get(self, "iam")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="manualQuota")
    def manual_quota(self) -> bool:
        """
        Manual quota prevent automatic quota upgrade
        """
        return pulumi.get(self, "manual_quota")

    @property
    @pulumi.getter(name="orderId")
    def order_id(self) -> float:
        """
        Project order ID
        """
        return pulumi.get(self, "order_id")

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> str:
        """
        Order plan code
        """
        return pulumi.get(self, "plan_code")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        Project ID
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> str:
        """
        Project name
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        ID of the public cloud project
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Current status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def unleash(self) -> bool:
        """
        Project unleashed
        """
        return pulumi.get(self, "unleash")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            access=self.access,
            creation_date=self.creation_date,
            description=self.description,
            expiration=self.expiration,
            iam=self.iam,
            id=self.id,
            manual_quota=self.manual_quota,
            order_id=self.order_id,
            plan_code=self.plan_code,
            project_id=self.project_id,
            project_name=self.project_name,
            service_name=self.service_name,
            status=self.status,
            unleash=self.unleash)


def get_project(service_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Get the details of a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    project = ovh.Cloud.get_project(service_name="XXX")
    ```


    :param str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Cloud/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        access=pulumi.get(__ret__, 'access'),
        creation_date=pulumi.get(__ret__, 'creation_date'),
        description=pulumi.get(__ret__, 'description'),
        expiration=pulumi.get(__ret__, 'expiration'),
        iam=pulumi.get(__ret__, 'iam'),
        id=pulumi.get(__ret__, 'id'),
        manual_quota=pulumi.get(__ret__, 'manual_quota'),
        order_id=pulumi.get(__ret__, 'order_id'),
        plan_code=pulumi.get(__ret__, 'plan_code'),
        project_id=pulumi.get(__ret__, 'project_id'),
        project_name=pulumi.get(__ret__, 'project_name'),
        service_name=pulumi.get(__ret__, 'service_name'),
        status=pulumi.get(__ret__, 'status'),
        unleash=pulumi.get(__ret__, 'unleash'))
def get_project_output(service_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectResult]:
    """
    Get the details of a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    project = ovh.Cloud.get_project(service_name="XXX")
    ```


    :param str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Cloud/getProject:getProject', __args__, opts=opts, typ=GetProjectResult)
    return __ret__.apply(lambda __response__: GetProjectResult(
        access=pulumi.get(__response__, 'access'),
        creation_date=pulumi.get(__response__, 'creation_date'),
        description=pulumi.get(__response__, 'description'),
        expiration=pulumi.get(__response__, 'expiration'),
        iam=pulumi.get(__response__, 'iam'),
        id=pulumi.get(__response__, 'id'),
        manual_quota=pulumi.get(__response__, 'manual_quota'),
        order_id=pulumi.get(__response__, 'order_id'),
        plan_code=pulumi.get(__response__, 'plan_code'),
        project_id=pulumi.get(__response__, 'project_id'),
        project_name=pulumi.get(__response__, 'project_name'),
        service_name=pulumi.get(__response__, 'service_name'),
        status=pulumi.get(__response__, 'status'),
        unleash=pulumi.get(__response__, 'unleash')))
