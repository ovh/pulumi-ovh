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
    'GetMeResult',
    'AwaitableGetMeResult',
    'get_me',
    'get_me_output',
]

@pulumi.output_type
class GetMeResult:
    """
    A collection of values returned by getMe.
    """
    def __init__(__self__, account_urn=None, address=None, area=None, birth_city=None, birth_day=None, city=None, company_national_identification_number=None, corporation_type=None, country=None, currencies=None, customer_code=None, email=None, fax=None, firstname=None, id=None, italian_sdi=None, language=None, legalform=None, name=None, national_identification_number=None, nichandle=None, organisation=None, ovh_company=None, ovh_subsidiary=None, phone=None, phone_country=None, sex=None, spare_email=None, state=None, vat=None, zip=None):
        if account_urn and not isinstance(account_urn, str):
            raise TypeError("Expected argument 'account_urn' to be a str")
        pulumi.set(__self__, "account_urn", account_urn)
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if area and not isinstance(area, str):
            raise TypeError("Expected argument 'area' to be a str")
        pulumi.set(__self__, "area", area)
        if birth_city and not isinstance(birth_city, str):
            raise TypeError("Expected argument 'birth_city' to be a str")
        pulumi.set(__self__, "birth_city", birth_city)
        if birth_day and not isinstance(birth_day, str):
            raise TypeError("Expected argument 'birth_day' to be a str")
        pulumi.set(__self__, "birth_day", birth_day)
        if city and not isinstance(city, str):
            raise TypeError("Expected argument 'city' to be a str")
        pulumi.set(__self__, "city", city)
        if company_national_identification_number and not isinstance(company_national_identification_number, str):
            raise TypeError("Expected argument 'company_national_identification_number' to be a str")
        pulumi.set(__self__, "company_national_identification_number", company_national_identification_number)
        if corporation_type and not isinstance(corporation_type, str):
            raise TypeError("Expected argument 'corporation_type' to be a str")
        pulumi.set(__self__, "corporation_type", corporation_type)
        if country and not isinstance(country, str):
            raise TypeError("Expected argument 'country' to be a str")
        pulumi.set(__self__, "country", country)
        if currencies and not isinstance(currencies, list):
            raise TypeError("Expected argument 'currencies' to be a list")
        pulumi.set(__self__, "currencies", currencies)
        if customer_code and not isinstance(customer_code, str):
            raise TypeError("Expected argument 'customer_code' to be a str")
        pulumi.set(__self__, "customer_code", customer_code)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if fax and not isinstance(fax, str):
            raise TypeError("Expected argument 'fax' to be a str")
        pulumi.set(__self__, "fax", fax)
        if firstname and not isinstance(firstname, str):
            raise TypeError("Expected argument 'firstname' to be a str")
        pulumi.set(__self__, "firstname", firstname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if italian_sdi and not isinstance(italian_sdi, str):
            raise TypeError("Expected argument 'italian_sdi' to be a str")
        pulumi.set(__self__, "italian_sdi", italian_sdi)
        if language and not isinstance(language, str):
            raise TypeError("Expected argument 'language' to be a str")
        pulumi.set(__self__, "language", language)
        if legalform and not isinstance(legalform, str):
            raise TypeError("Expected argument 'legalform' to be a str")
        pulumi.set(__self__, "legalform", legalform)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if national_identification_number and not isinstance(national_identification_number, str):
            raise TypeError("Expected argument 'national_identification_number' to be a str")
        pulumi.set(__self__, "national_identification_number", national_identification_number)
        if nichandle and not isinstance(nichandle, str):
            raise TypeError("Expected argument 'nichandle' to be a str")
        pulumi.set(__self__, "nichandle", nichandle)
        if organisation and not isinstance(organisation, str):
            raise TypeError("Expected argument 'organisation' to be a str")
        pulumi.set(__self__, "organisation", organisation)
        if ovh_company and not isinstance(ovh_company, str):
            raise TypeError("Expected argument 'ovh_company' to be a str")
        pulumi.set(__self__, "ovh_company", ovh_company)
        if ovh_subsidiary and not isinstance(ovh_subsidiary, str):
            raise TypeError("Expected argument 'ovh_subsidiary' to be a str")
        pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        if phone and not isinstance(phone, str):
            raise TypeError("Expected argument 'phone' to be a str")
        pulumi.set(__self__, "phone", phone)
        if phone_country and not isinstance(phone_country, str):
            raise TypeError("Expected argument 'phone_country' to be a str")
        pulumi.set(__self__, "phone_country", phone_country)
        if sex and not isinstance(sex, str):
            raise TypeError("Expected argument 'sex' to be a str")
        pulumi.set(__self__, "sex", sex)
        if spare_email and not isinstance(spare_email, str):
            raise TypeError("Expected argument 'spare_email' to be a str")
        pulumi.set(__self__, "spare_email", spare_email)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if vat and not isinstance(vat, str):
            raise TypeError("Expected argument 'vat' to be a str")
        pulumi.set(__self__, "vat", vat)
        if zip and not isinstance(zip, str):
            raise TypeError("Expected argument 'zip' to be a str")
        pulumi.set(__self__, "zip", zip)

    @property
    @pulumi.getter(name="AccountURN")
    def account_urn(self) -> str:
        return pulumi.get(self, "account_urn")

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def area(self) -> str:
        return pulumi.get(self, "area")

    @property
    @pulumi.getter(name="birthCity")
    def birth_city(self) -> str:
        return pulumi.get(self, "birth_city")

    @property
    @pulumi.getter(name="birthDay")
    def birth_day(self) -> str:
        return pulumi.get(self, "birth_day")

    @property
    @pulumi.getter
    def city(self) -> str:
        return pulumi.get(self, "city")

    @property
    @pulumi.getter(name="companyNationalIdentificationNumber")
    def company_national_identification_number(self) -> str:
        return pulumi.get(self, "company_national_identification_number")

    @property
    @pulumi.getter(name="corporationType")
    def corporation_type(self) -> str:
        return pulumi.get(self, "corporation_type")

    @property
    @pulumi.getter
    def country(self) -> str:
        return pulumi.get(self, "country")

    @property
    @pulumi.getter
    def currencies(self) -> Sequence['outputs.GetMeCurrencyResult']:
        return pulumi.get(self, "currencies")

    @property
    @pulumi.getter(name="customerCode")
    def customer_code(self) -> str:
        return pulumi.get(self, "customer_code")

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def fax(self) -> str:
        return pulumi.get(self, "fax")

    @property
    @pulumi.getter
    def firstname(self) -> str:
        return pulumi.get(self, "firstname")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="italianSdi")
    def italian_sdi(self) -> str:
        return pulumi.get(self, "italian_sdi")

    @property
    @pulumi.getter
    def language(self) -> str:
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def legalform(self) -> str:
        return pulumi.get(self, "legalform")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nationalIdentificationNumber")
    def national_identification_number(self) -> str:
        return pulumi.get(self, "national_identification_number")

    @property
    @pulumi.getter
    def nichandle(self) -> str:
        return pulumi.get(self, "nichandle")

    @property
    @pulumi.getter
    def organisation(self) -> str:
        return pulumi.get(self, "organisation")

    @property
    @pulumi.getter(name="ovhCompany")
    def ovh_company(self) -> str:
        return pulumi.get(self, "ovh_company")

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> str:
        return pulumi.get(self, "ovh_subsidiary")

    @property
    @pulumi.getter
    def phone(self) -> str:
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter(name="phoneCountry")
    def phone_country(self) -> str:
        return pulumi.get(self, "phone_country")

    @property
    @pulumi.getter
    def sex(self) -> str:
        return pulumi.get(self, "sex")

    @property
    @pulumi.getter(name="spareEmail")
    def spare_email(self) -> str:
        return pulumi.get(self, "spare_email")

    @property
    @pulumi.getter
    def state(self) -> str:
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def vat(self) -> str:
        return pulumi.get(self, "vat")

    @property
    @pulumi.getter
    def zip(self) -> str:
        return pulumi.get(self, "zip")


class AwaitableGetMeResult(GetMeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMeResult(
            account_urn=self.account_urn,
            address=self.address,
            area=self.area,
            birth_city=self.birth_city,
            birth_day=self.birth_day,
            city=self.city,
            company_national_identification_number=self.company_national_identification_number,
            corporation_type=self.corporation_type,
            country=self.country,
            currencies=self.currencies,
            customer_code=self.customer_code,
            email=self.email,
            fax=self.fax,
            firstname=self.firstname,
            id=self.id,
            italian_sdi=self.italian_sdi,
            language=self.language,
            legalform=self.legalform,
            name=self.name,
            national_identification_number=self.national_identification_number,
            nichandle=self.nichandle,
            organisation=self.organisation,
            ovh_company=self.ovh_company,
            ovh_subsidiary=self.ovh_subsidiary,
            phone=self.phone,
            phone_country=self.phone_country,
            sex=self.sex,
            spare_email=self.spare_email,
            state=self.state,
            vat=self.vat,
            zip=self.zip)


def get_me(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMeResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Me/getMe:getMe', __args__, opts=opts, typ=GetMeResult).value

    return AwaitableGetMeResult(
        account_urn=pulumi.get(__ret__, 'account_urn'),
        address=pulumi.get(__ret__, 'address'),
        area=pulumi.get(__ret__, 'area'),
        birth_city=pulumi.get(__ret__, 'birth_city'),
        birth_day=pulumi.get(__ret__, 'birth_day'),
        city=pulumi.get(__ret__, 'city'),
        company_national_identification_number=pulumi.get(__ret__, 'company_national_identification_number'),
        corporation_type=pulumi.get(__ret__, 'corporation_type'),
        country=pulumi.get(__ret__, 'country'),
        currencies=pulumi.get(__ret__, 'currencies'),
        customer_code=pulumi.get(__ret__, 'customer_code'),
        email=pulumi.get(__ret__, 'email'),
        fax=pulumi.get(__ret__, 'fax'),
        firstname=pulumi.get(__ret__, 'firstname'),
        id=pulumi.get(__ret__, 'id'),
        italian_sdi=pulumi.get(__ret__, 'italian_sdi'),
        language=pulumi.get(__ret__, 'language'),
        legalform=pulumi.get(__ret__, 'legalform'),
        name=pulumi.get(__ret__, 'name'),
        national_identification_number=pulumi.get(__ret__, 'national_identification_number'),
        nichandle=pulumi.get(__ret__, 'nichandle'),
        organisation=pulumi.get(__ret__, 'organisation'),
        ovh_company=pulumi.get(__ret__, 'ovh_company'),
        ovh_subsidiary=pulumi.get(__ret__, 'ovh_subsidiary'),
        phone=pulumi.get(__ret__, 'phone'),
        phone_country=pulumi.get(__ret__, 'phone_country'),
        sex=pulumi.get(__ret__, 'sex'),
        spare_email=pulumi.get(__ret__, 'spare_email'),
        state=pulumi.get(__ret__, 'state'),
        vat=pulumi.get(__ret__, 'vat'),
        zip=pulumi.get(__ret__, 'zip'))
def get_me_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMeResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Me/getMe:getMe', __args__, opts=opts, typ=GetMeResult)
    return __ret__.apply(lambda __response__: GetMeResult(
        account_urn=pulumi.get(__response__, 'account_urn'),
        address=pulumi.get(__response__, 'address'),
        area=pulumi.get(__response__, 'area'),
        birth_city=pulumi.get(__response__, 'birth_city'),
        birth_day=pulumi.get(__response__, 'birth_day'),
        city=pulumi.get(__response__, 'city'),
        company_national_identification_number=pulumi.get(__response__, 'company_national_identification_number'),
        corporation_type=pulumi.get(__response__, 'corporation_type'),
        country=pulumi.get(__response__, 'country'),
        currencies=pulumi.get(__response__, 'currencies'),
        customer_code=pulumi.get(__response__, 'customer_code'),
        email=pulumi.get(__response__, 'email'),
        fax=pulumi.get(__response__, 'fax'),
        firstname=pulumi.get(__response__, 'firstname'),
        id=pulumi.get(__response__, 'id'),
        italian_sdi=pulumi.get(__response__, 'italian_sdi'),
        language=pulumi.get(__response__, 'language'),
        legalform=pulumi.get(__response__, 'legalform'),
        name=pulumi.get(__response__, 'name'),
        national_identification_number=pulumi.get(__response__, 'national_identification_number'),
        nichandle=pulumi.get(__response__, 'nichandle'),
        organisation=pulumi.get(__response__, 'organisation'),
        ovh_company=pulumi.get(__response__, 'ovh_company'),
        ovh_subsidiary=pulumi.get(__response__, 'ovh_subsidiary'),
        phone=pulumi.get(__response__, 'phone'),
        phone_country=pulumi.get(__response__, 'phone_country'),
        sex=pulumi.get(__response__, 'sex'),
        spare_email=pulumi.get(__response__, 'spare_email'),
        state=pulumi.get(__response__, 'state'),
        vat=pulumi.get(__response__, 'vat'),
        zip=pulumi.get(__response__, 'zip')))
