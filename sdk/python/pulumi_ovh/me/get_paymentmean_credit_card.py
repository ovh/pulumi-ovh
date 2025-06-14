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

__all__ = [
    'GetPaymentmeanCreditCardResult',
    'AwaitableGetPaymentmeanCreditCardResult',
    'get_paymentmean_credit_card',
    'get_paymentmean_credit_card_output',
]

@pulumi.output_type
class GetPaymentmeanCreditCardResult:
    """
    A collection of values returned by getPaymentmeanCreditCard.
    """
    def __init__(__self__, default=None, description=None, description_regexp=None, id=None, state=None, states=None, use_default=None, use_last_to_expire=None):
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        pulumi.set(__self__, "default", default)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if description_regexp and not isinstance(description_regexp, str):
            raise TypeError("Expected argument 'description_regexp' to be a str")
        pulumi.set(__self__, "description_regexp", description_regexp)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if states and not isinstance(states, list):
            raise TypeError("Expected argument 'states' to be a list")
        pulumi.set(__self__, "states", states)
        if use_default and not isinstance(use_default, bool):
            raise TypeError("Expected argument 'use_default' to be a bool")
        pulumi.set(__self__, "use_default", use_default)
        if use_last_to_expire and not isinstance(use_last_to_expire, bool):
            raise TypeError("Expected argument 'use_last_to_expire' to be a bool")
        pulumi.set(__self__, "use_last_to_expire", use_last_to_expire)

    @property
    @pulumi.getter
    def default(self) -> builtins.bool:
        """
        a boolean which tells if the retrieved credit card is marked as the default payment mean
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        the description attribute of the credit card
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="descriptionRegexp")
    def description_regexp(self) -> Optional[builtins.str]:
        return pulumi.get(self, "description_regexp")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def state(self) -> builtins.str:
        """
        the state attribute of the credit card
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def states(self) -> Optional[Sequence[builtins.str]]:
        return pulumi.get(self, "states")

    @property
    @pulumi.getter(name="useDefault")
    def use_default(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "use_default")

    @property
    @pulumi.getter(name="useLastToExpire")
    def use_last_to_expire(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "use_last_to_expire")


class AwaitableGetPaymentmeanCreditCardResult(GetPaymentmeanCreditCardResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPaymentmeanCreditCardResult(
            default=self.default,
            description=self.description,
            description_regexp=self.description_regexp,
            id=self.id,
            state=self.state,
            states=self.states,
            use_default=self.use_default,
            use_last_to_expire=self.use_last_to_expire)


def get_paymentmean_credit_card(description_regexp: Optional[builtins.str] = None,
                                states: Optional[Sequence[builtins.str]] = None,
                                use_default: Optional[builtins.bool] = None,
                                use_last_to_expire: Optional[builtins.bool] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPaymentmeanCreditCardResult:
    """
    Use this data source to retrieve information about a credit card payment mean associated with an OVHcloud account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    cc = ovh.Me.get_paymentmean_credit_card(use_default=True)
    ```


    :param builtins.str description_regexp: a regexp used to filter credit cards on their `description` attributes.
    :param Sequence[builtins.str] states: Filter credit cards on their `state` attribute. Can be "expired", "valid", "tooManyFailures"
    :param builtins.bool use_default: Retrieve credit card marked as default payment mean.
    :param builtins.bool use_last_to_expire: Retrieve the credit card that will be the last to expire according to its expiration date.
    """
    __args__ = dict()
    __args__['descriptionRegexp'] = description_regexp
    __args__['states'] = states
    __args__['useDefault'] = use_default
    __args__['useLastToExpire'] = use_last_to_expire
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Me/getPaymentmeanCreditCard:getPaymentmeanCreditCard', __args__, opts=opts, typ=GetPaymentmeanCreditCardResult).value

    return AwaitableGetPaymentmeanCreditCardResult(
        default=pulumi.get(__ret__, 'default'),
        description=pulumi.get(__ret__, 'description'),
        description_regexp=pulumi.get(__ret__, 'description_regexp'),
        id=pulumi.get(__ret__, 'id'),
        state=pulumi.get(__ret__, 'state'),
        states=pulumi.get(__ret__, 'states'),
        use_default=pulumi.get(__ret__, 'use_default'),
        use_last_to_expire=pulumi.get(__ret__, 'use_last_to_expire'))
def get_paymentmean_credit_card_output(description_regexp: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                       states: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                                       use_default: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                       use_last_to_expire: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPaymentmeanCreditCardResult]:
    """
    Use this data source to retrieve information about a credit card payment mean associated with an OVHcloud account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    cc = ovh.Me.get_paymentmean_credit_card(use_default=True)
    ```


    :param builtins.str description_regexp: a regexp used to filter credit cards on their `description` attributes.
    :param Sequence[builtins.str] states: Filter credit cards on their `state` attribute. Can be "expired", "valid", "tooManyFailures"
    :param builtins.bool use_default: Retrieve credit card marked as default payment mean.
    :param builtins.bool use_last_to_expire: Retrieve the credit card that will be the last to expire according to its expiration date.
    """
    __args__ = dict()
    __args__['descriptionRegexp'] = description_regexp
    __args__['states'] = states
    __args__['useDefault'] = use_default
    __args__['useLastToExpire'] = use_last_to_expire
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Me/getPaymentmeanCreditCard:getPaymentmeanCreditCard', __args__, opts=opts, typ=GetPaymentmeanCreditCardResult)
    return __ret__.apply(lambda __response__: GetPaymentmeanCreditCardResult(
        default=pulumi.get(__response__, 'default'),
        description=pulumi.get(__response__, 'description'),
        description_regexp=pulumi.get(__response__, 'description_regexp'),
        id=pulumi.get(__response__, 'id'),
        state=pulumi.get(__response__, 'state'),
        states=pulumi.get(__response__, 'states'),
        use_default=pulumi.get(__response__, 'use_default'),
        use_last_to_expire=pulumi.get(__response__, 'use_last_to_expire')))
