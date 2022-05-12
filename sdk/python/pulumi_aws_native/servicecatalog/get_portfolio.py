# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetPortfolioResult',
    'AwaitableGetPortfolioResult',
    'get_portfolio',
    'get_portfolio_output',
]

@pulumi.output_type
class GetPortfolioResult:
    def __init__(__self__, accept_language=None, description=None, display_name=None, id=None, portfolio_name=None, provider_name=None, tags=None):
        if accept_language and not isinstance(accept_language, str):
            raise TypeError("Expected argument 'accept_language' to be a str")
        pulumi.set(__self__, "accept_language", accept_language)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if portfolio_name and not isinstance(portfolio_name, str):
            raise TypeError("Expected argument 'portfolio_name' to be a str")
        pulumi.set(__self__, "portfolio_name", portfolio_name)
        if provider_name and not isinstance(provider_name, str):
            raise TypeError("Expected argument 'provider_name' to be a str")
        pulumi.set(__self__, "provider_name", provider_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="acceptLanguage")
    def accept_language(self) -> Optional[str]:
        return pulumi.get(self, "accept_language")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="portfolioName")
    def portfolio_name(self) -> Optional[str]:
        return pulumi.get(self, "portfolio_name")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> Optional[str]:
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.PortfolioTag']]:
        return pulumi.get(self, "tags")


class AwaitableGetPortfolioResult(GetPortfolioResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPortfolioResult(
            accept_language=self.accept_language,
            description=self.description,
            display_name=self.display_name,
            id=self.id,
            portfolio_name=self.portfolio_name,
            provider_name=self.provider_name,
            tags=self.tags)


def get_portfolio(id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPortfolioResult:
    """
    Resource Type definition for AWS::ServiceCatalog::Portfolio
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:servicecatalog:getPortfolio', __args__, opts=opts, typ=GetPortfolioResult).value

    return AwaitableGetPortfolioResult(
        accept_language=__ret__.accept_language,
        description=__ret__.description,
        display_name=__ret__.display_name,
        id=__ret__.id,
        portfolio_name=__ret__.portfolio_name,
        provider_name=__ret__.provider_name,
        tags=__ret__.tags)


@_utilities.lift_output_func(get_portfolio)
def get_portfolio_output(id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPortfolioResult]:
    """
    Resource Type definition for AWS::ServiceCatalog::Portfolio
    """
    ...
