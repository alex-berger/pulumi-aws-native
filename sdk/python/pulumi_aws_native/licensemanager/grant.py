# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['GrantArgs', 'Grant']

@pulumi.input_type
class GrantArgs:
    def __init__(__self__, *,
                 allowed_operations: Optional[pulumi.Input['GrantAllowedOperationListArgs']] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 filters: Optional[pulumi.Input['GrantFilterListArgs']] = None,
                 grant_arns: Optional[pulumi.Input['GrantArnListArgs']] = None,
                 grant_name: Optional[pulumi.Input[str]] = None,
                 grant_status: Optional[pulumi.Input[str]] = None,
                 granted_operations: Optional[pulumi.Input['GrantAllowedOperationListArgs']] = None,
                 grantee_principal_arn: Optional[pulumi.Input[str]] = None,
                 home_region: Optional[pulumi.Input[str]] = None,
                 license_arn: Optional[pulumi.Input[str]] = None,
                 max_results: Optional[pulumi.Input[int]] = None,
                 next_token: Optional[pulumi.Input[str]] = None,
                 parent_arn: Optional[pulumi.Input[str]] = None,
                 principals: Optional[pulumi.Input['GrantArnListArgs']] = None,
                 source_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_reason: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input['GrantTagListArgs']] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Grant resource.
        :param pulumi.Input['GrantAllowedOperationListArgs'] allowed_operations: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-allowedoperations
        :param pulumi.Input[str] client_token: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-clienttoken
        :param pulumi.Input['GrantFilterListArgs'] filters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-filters
        :param pulumi.Input['GrantArnListArgs'] grant_arns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantarns
        :param pulumi.Input[str] grant_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantname
        :param pulumi.Input[str] grant_status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantstatus
        :param pulumi.Input['GrantAllowedOperationListArgs'] granted_operations: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantedoperations
        :param pulumi.Input[str] grantee_principal_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-granteeprincipalarn
        :param pulumi.Input[str] home_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-homeregion
        :param pulumi.Input[str] license_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-licensearn
        :param pulumi.Input[int] max_results: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-maxresults
        :param pulumi.Input[str] next_token: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-nexttoken
        :param pulumi.Input[str] parent_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-parentarn
        :param pulumi.Input['GrantArnListArgs'] principals: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-principals
        :param pulumi.Input[str] source_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-sourceversion
        :param pulumi.Input[str] status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-status
        :param pulumi.Input[str] status_reason: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-statusreason
        :param pulumi.Input['GrantTagListArgs'] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-tags
        :param pulumi.Input[str] version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-version
        """
        if allowed_operations is not None:
            pulumi.set(__self__, "allowed_operations", allowed_operations)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if grant_arns is not None:
            pulumi.set(__self__, "grant_arns", grant_arns)
        if grant_name is not None:
            pulumi.set(__self__, "grant_name", grant_name)
        if grant_status is not None:
            pulumi.set(__self__, "grant_status", grant_status)
        if granted_operations is not None:
            pulumi.set(__self__, "granted_operations", granted_operations)
        if grantee_principal_arn is not None:
            pulumi.set(__self__, "grantee_principal_arn", grantee_principal_arn)
        if home_region is not None:
            pulumi.set(__self__, "home_region", home_region)
        if license_arn is not None:
            pulumi.set(__self__, "license_arn", license_arn)
        if max_results is not None:
            pulumi.set(__self__, "max_results", max_results)
        if next_token is not None:
            pulumi.set(__self__, "next_token", next_token)
        if parent_arn is not None:
            pulumi.set(__self__, "parent_arn", parent_arn)
        if principals is not None:
            pulumi.set(__self__, "principals", principals)
        if source_version is not None:
            pulumi.set(__self__, "source_version", source_version)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if status_reason is not None:
            pulumi.set(__self__, "status_reason", status_reason)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="allowedOperations")
    def allowed_operations(self) -> Optional[pulumi.Input['GrantAllowedOperationListArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-allowedoperations
        """
        return pulumi.get(self, "allowed_operations")

    @allowed_operations.setter
    def allowed_operations(self, value: Optional[pulumi.Input['GrantAllowedOperationListArgs']]):
        pulumi.set(self, "allowed_operations", value)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-clienttoken
        """
        return pulumi.get(self, "client_token")

    @client_token.setter
    def client_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_token", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input['GrantFilterListArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-filters
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input['GrantFilterListArgs']]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter(name="grantArns")
    def grant_arns(self) -> Optional[pulumi.Input['GrantArnListArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantarns
        """
        return pulumi.get(self, "grant_arns")

    @grant_arns.setter
    def grant_arns(self, value: Optional[pulumi.Input['GrantArnListArgs']]):
        pulumi.set(self, "grant_arns", value)

    @property
    @pulumi.getter(name="grantName")
    def grant_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantname
        """
        return pulumi.get(self, "grant_name")

    @grant_name.setter
    def grant_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_name", value)

    @property
    @pulumi.getter(name="grantStatus")
    def grant_status(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantstatus
        """
        return pulumi.get(self, "grant_status")

    @grant_status.setter
    def grant_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_status", value)

    @property
    @pulumi.getter(name="grantedOperations")
    def granted_operations(self) -> Optional[pulumi.Input['GrantAllowedOperationListArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantedoperations
        """
        return pulumi.get(self, "granted_operations")

    @granted_operations.setter
    def granted_operations(self, value: Optional[pulumi.Input['GrantAllowedOperationListArgs']]):
        pulumi.set(self, "granted_operations", value)

    @property
    @pulumi.getter(name="granteePrincipalArn")
    def grantee_principal_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-granteeprincipalarn
        """
        return pulumi.get(self, "grantee_principal_arn")

    @grantee_principal_arn.setter
    def grantee_principal_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grantee_principal_arn", value)

    @property
    @pulumi.getter(name="homeRegion")
    def home_region(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-homeregion
        """
        return pulumi.get(self, "home_region")

    @home_region.setter
    def home_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "home_region", value)

    @property
    @pulumi.getter(name="licenseArn")
    def license_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-licensearn
        """
        return pulumi.get(self, "license_arn")

    @license_arn.setter
    def license_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license_arn", value)

    @property
    @pulumi.getter(name="maxResults")
    def max_results(self) -> Optional[pulumi.Input[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-maxresults
        """
        return pulumi.get(self, "max_results")

    @max_results.setter
    def max_results(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_results", value)

    @property
    @pulumi.getter(name="nextToken")
    def next_token(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-nexttoken
        """
        return pulumi.get(self, "next_token")

    @next_token.setter
    def next_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_token", value)

    @property
    @pulumi.getter(name="parentArn")
    def parent_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-parentarn
        """
        return pulumi.get(self, "parent_arn")

    @parent_arn.setter
    def parent_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_arn", value)

    @property
    @pulumi.getter
    def principals(self) -> Optional[pulumi.Input['GrantArnListArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-principals
        """
        return pulumi.get(self, "principals")

    @principals.setter
    def principals(self, value: Optional[pulumi.Input['GrantArnListArgs']]):
        pulumi.set(self, "principals", value)

    @property
    @pulumi.getter(name="sourceVersion")
    def source_version(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-sourceversion
        """
        return pulumi.get(self, "source_version")

    @source_version.setter
    def source_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_version", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="statusReason")
    def status_reason(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-statusreason
        """
        return pulumi.get(self, "status_reason")

    @status_reason.setter
    def status_reason(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status_reason", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input['GrantTagListArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input['GrantTagListArgs']]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-version
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Grant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_operations: Optional[pulumi.Input[pulumi.InputType['GrantAllowedOperationListArgs']]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 filters: Optional[pulumi.Input[pulumi.InputType['GrantFilterListArgs']]] = None,
                 grant_arns: Optional[pulumi.Input[pulumi.InputType['GrantArnListArgs']]] = None,
                 grant_name: Optional[pulumi.Input[str]] = None,
                 grant_status: Optional[pulumi.Input[str]] = None,
                 granted_operations: Optional[pulumi.Input[pulumi.InputType['GrantAllowedOperationListArgs']]] = None,
                 grantee_principal_arn: Optional[pulumi.Input[str]] = None,
                 home_region: Optional[pulumi.Input[str]] = None,
                 license_arn: Optional[pulumi.Input[str]] = None,
                 max_results: Optional[pulumi.Input[int]] = None,
                 next_token: Optional[pulumi.Input[str]] = None,
                 parent_arn: Optional[pulumi.Input[str]] = None,
                 principals: Optional[pulumi.Input[pulumi.InputType['GrantArnListArgs']]] = None,
                 source_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_reason: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['GrantTagListArgs']]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GrantAllowedOperationListArgs']] allowed_operations: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-allowedoperations
        :param pulumi.Input[str] client_token: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-clienttoken
        :param pulumi.Input[pulumi.InputType['GrantFilterListArgs']] filters: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-filters
        :param pulumi.Input[pulumi.InputType['GrantArnListArgs']] grant_arns: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantarns
        :param pulumi.Input[str] grant_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantname
        :param pulumi.Input[str] grant_status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantstatus
        :param pulumi.Input[pulumi.InputType['GrantAllowedOperationListArgs']] granted_operations: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantedoperations
        :param pulumi.Input[str] grantee_principal_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-granteeprincipalarn
        :param pulumi.Input[str] home_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-homeregion
        :param pulumi.Input[str] license_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-licensearn
        :param pulumi.Input[int] max_results: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-maxresults
        :param pulumi.Input[str] next_token: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-nexttoken
        :param pulumi.Input[str] parent_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-parentarn
        :param pulumi.Input[pulumi.InputType['GrantArnListArgs']] principals: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-principals
        :param pulumi.Input[str] source_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-sourceversion
        :param pulumi.Input[str] status: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-status
        :param pulumi.Input[str] status_reason: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-statusreason
        :param pulumi.Input[pulumi.InputType['GrantTagListArgs']] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-tags
        :param pulumi.Input[str] version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-version
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GrantArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html

        :param str resource_name: The name of the resource.
        :param GrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_operations: Optional[pulumi.Input[pulumi.InputType['GrantAllowedOperationListArgs']]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 filters: Optional[pulumi.Input[pulumi.InputType['GrantFilterListArgs']]] = None,
                 grant_arns: Optional[pulumi.Input[pulumi.InputType['GrantArnListArgs']]] = None,
                 grant_name: Optional[pulumi.Input[str]] = None,
                 grant_status: Optional[pulumi.Input[str]] = None,
                 granted_operations: Optional[pulumi.Input[pulumi.InputType['GrantAllowedOperationListArgs']]] = None,
                 grantee_principal_arn: Optional[pulumi.Input[str]] = None,
                 home_region: Optional[pulumi.Input[str]] = None,
                 license_arn: Optional[pulumi.Input[str]] = None,
                 max_results: Optional[pulumi.Input[int]] = None,
                 next_token: Optional[pulumi.Input[str]] = None,
                 parent_arn: Optional[pulumi.Input[str]] = None,
                 principals: Optional[pulumi.Input[pulumi.InputType['GrantArnListArgs']]] = None,
                 source_version: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 status_reason: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[pulumi.InputType['GrantTagListArgs']]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GrantArgs.__new__(GrantArgs)

            __props__.__dict__["allowed_operations"] = allowed_operations
            __props__.__dict__["client_token"] = client_token
            __props__.__dict__["filters"] = filters
            __props__.__dict__["grant_arns"] = grant_arns
            __props__.__dict__["grant_name"] = grant_name
            __props__.__dict__["grant_status"] = grant_status
            __props__.__dict__["granted_operations"] = granted_operations
            __props__.__dict__["grantee_principal_arn"] = grantee_principal_arn
            __props__.__dict__["home_region"] = home_region
            __props__.__dict__["license_arn"] = license_arn
            __props__.__dict__["max_results"] = max_results
            __props__.__dict__["next_token"] = next_token
            __props__.__dict__["parent_arn"] = parent_arn
            __props__.__dict__["principals"] = principals
            __props__.__dict__["source_version"] = source_version
            __props__.__dict__["status"] = status
            __props__.__dict__["status_reason"] = status_reason
            __props__.__dict__["tags"] = tags
            __props__.__dict__["version"] = version
            __props__.__dict__["grant_arn"] = None
        super(Grant, __self__).__init__(
            'aws-native:LicenseManager:Grant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Grant':
        """
        Get an existing Grant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GrantArgs.__new__(GrantArgs)

        __props__.__dict__["allowed_operations"] = None
        __props__.__dict__["client_token"] = None
        __props__.__dict__["filters"] = None
        __props__.__dict__["grant_arn"] = None
        __props__.__dict__["grant_arns"] = None
        __props__.__dict__["grant_name"] = None
        __props__.__dict__["grant_status"] = None
        __props__.__dict__["granted_operations"] = None
        __props__.__dict__["grantee_principal_arn"] = None
        __props__.__dict__["home_region"] = None
        __props__.__dict__["license_arn"] = None
        __props__.__dict__["max_results"] = None
        __props__.__dict__["next_token"] = None
        __props__.__dict__["parent_arn"] = None
        __props__.__dict__["principals"] = None
        __props__.__dict__["source_version"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_reason"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["version"] = None
        return Grant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedOperations")
    def allowed_operations(self) -> pulumi.Output[Optional['outputs.GrantAllowedOperationList']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-allowedoperations
        """
        return pulumi.get(self, "allowed_operations")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-clienttoken
        """
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional['outputs.GrantFilterList']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-filters
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter(name="grantArn")
    def grant_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "grant_arn")

    @property
    @pulumi.getter(name="grantArns")
    def grant_arns(self) -> pulumi.Output[Optional['outputs.GrantArnList']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantarns
        """
        return pulumi.get(self, "grant_arns")

    @property
    @pulumi.getter(name="grantName")
    def grant_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantname
        """
        return pulumi.get(self, "grant_name")

    @property
    @pulumi.getter(name="grantStatus")
    def grant_status(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantstatus
        """
        return pulumi.get(self, "grant_status")

    @property
    @pulumi.getter(name="grantedOperations")
    def granted_operations(self) -> pulumi.Output[Optional['outputs.GrantAllowedOperationList']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-grantedoperations
        """
        return pulumi.get(self, "granted_operations")

    @property
    @pulumi.getter(name="granteePrincipalArn")
    def grantee_principal_arn(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-granteeprincipalarn
        """
        return pulumi.get(self, "grantee_principal_arn")

    @property
    @pulumi.getter(name="homeRegion")
    def home_region(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-homeregion
        """
        return pulumi.get(self, "home_region")

    @property
    @pulumi.getter(name="licenseArn")
    def license_arn(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-licensearn
        """
        return pulumi.get(self, "license_arn")

    @property
    @pulumi.getter(name="maxResults")
    def max_results(self) -> pulumi.Output[Optional[int]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-maxresults
        """
        return pulumi.get(self, "max_results")

    @property
    @pulumi.getter(name="nextToken")
    def next_token(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-nexttoken
        """
        return pulumi.get(self, "next_token")

    @property
    @pulumi.getter(name="parentArn")
    def parent_arn(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-parentarn
        """
        return pulumi.get(self, "parent_arn")

    @property
    @pulumi.getter
    def principals(self) -> pulumi.Output[Optional['outputs.GrantArnList']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-principals
        """
        return pulumi.get(self, "principals")

    @property
    @pulumi.getter(name="sourceVersion")
    def source_version(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-sourceversion
        """
        return pulumi.get(self, "source_version")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusReason")
    def status_reason(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-statusreason
        """
        return pulumi.get(self, "status_reason")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional['outputs.GrantTagList']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html#cfn-licensemanager-grant-version
        """
        return pulumi.get(self, "version")

