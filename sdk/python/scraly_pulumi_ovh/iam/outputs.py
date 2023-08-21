# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetReferenceActionsActionResult',
]

@pulumi.output_type
class GetReferenceActionsActionResult(dict):
    def __init__(__self__, *,
                 action: str,
                 categories: Sequence[str],
                 description: str,
                 resource_type: str):
        """
        :param str action: Name of the action
        :param Sequence[str] categories: List of the categories of the action
        :param str description: Description of the action
        :param str resource_type: Resource type the action is related to
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "categories", categories)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        Name of the action
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def categories(self) -> Sequence[str]:
        """
        List of the categories of the action
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the action
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        Resource type the action is related to
        """
        return pulumi.get(self, "resource_type")

