# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AnalysisArgs', 'Analysis']

@pulumi.input_type
class AnalysisArgs:
    def __init__(__self__, *,
                 analysis_id: pulumi.Input[str],
                 aws_account_id: pulumi.Input[str],
                 source_entity: pulumi.Input['AnalysisSourceEntityArgs'],
                 errors: Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisErrorArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input['AnalysisParametersArgs']] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisResourcePermissionArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisTagArgs']]]] = None,
                 theme_arn: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Analysis resource.
        :param pulumi.Input[Sequence[pulumi.Input['AnalysisErrorArgs']]] errors: <p>Errors associated with the analysis.</p>
        :param pulumi.Input[str] name: <p>The descriptive name of the analysis.</p>
        :param pulumi.Input[Sequence[pulumi.Input['AnalysisResourcePermissionArgs']]] permissions: <p>A structure that describes the principals and the resource-level permissions on an
                           analysis. You can use the <code>Permissions</code> structure to grant permissions by
                           providing a list of AWS Identity and Access Management (IAM) action information for each
                           principal listed by Amazon Resource Name (ARN). </p>
               
                       <p>To specify no permissions, omit <code>Permissions</code>.</p>
        :param pulumi.Input[Sequence[pulumi.Input['AnalysisTagArgs']]] tags: <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
                           analysis.</p>
        :param pulumi.Input[str] theme_arn: <p>The ARN of the theme of the analysis.</p>
        """
        pulumi.set(__self__, "analysis_id", analysis_id)
        pulumi.set(__self__, "aws_account_id", aws_account_id)
        pulumi.set(__self__, "source_entity", source_entity)
        if errors is not None:
            pulumi.set(__self__, "errors", errors)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if theme_arn is not None:
            pulumi.set(__self__, "theme_arn", theme_arn)

    @property
    @pulumi.getter(name="analysisId")
    def analysis_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "analysis_id")

    @analysis_id.setter
    def analysis_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "analysis_id", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> pulumi.Input['AnalysisSourceEntityArgs']:
        return pulumi.get(self, "source_entity")

    @source_entity.setter
    def source_entity(self, value: pulumi.Input['AnalysisSourceEntityArgs']):
        pulumi.set(self, "source_entity", value)

    @property
    @pulumi.getter
    def errors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisErrorArgs']]]]:
        """
        <p>Errors associated with the analysis.</p>
        """
        return pulumi.get(self, "errors")

    @errors.setter
    def errors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisErrorArgs']]]]):
        pulumi.set(self, "errors", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        <p>The descriptive name of the analysis.</p>
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input['AnalysisParametersArgs']]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input['AnalysisParametersArgs']]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisResourcePermissionArgs']]]]:
        """
        <p>A structure that describes the principals and the resource-level permissions on an
                    analysis. You can use the <code>Permissions</code> structure to grant permissions by
                    providing a list of AWS Identity and Access Management (IAM) action information for each
                    principal listed by Amazon Resource Name (ARN). </p>

                <p>To specify no permissions, omit <code>Permissions</code>.</p>
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisResourcePermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisTagArgs']]]]:
        """
        <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
                    analysis.</p>
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AnalysisTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="themeArn")
    def theme_arn(self) -> Optional[pulumi.Input[str]]:
        """
        <p>The ARN of the theme of the analysis.</p>
        """
        return pulumi.get(self, "theme_arn")

    @theme_arn.setter
    def theme_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "theme_arn", value)


class Analysis(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 analysis_id: Optional[pulumi.Input[str]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 errors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisErrorArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[pulumi.InputType['AnalysisParametersArgs']]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisResourcePermissionArgs']]]]] = None,
                 source_entity: Optional[pulumi.Input[pulumi.InputType['AnalysisSourceEntityArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisTagArgs']]]]] = None,
                 theme_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Definition of the AWS::QuickSight::Analysis Resource Type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisErrorArgs']]]] errors: <p>Errors associated with the analysis.</p>
        :param pulumi.Input[str] name: <p>The descriptive name of the analysis.</p>
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisResourcePermissionArgs']]]] permissions: <p>A structure that describes the principals and the resource-level permissions on an
                           analysis. You can use the <code>Permissions</code> structure to grant permissions by
                           providing a list of AWS Identity and Access Management (IAM) action information for each
                           principal listed by Amazon Resource Name (ARN). </p>
               
                       <p>To specify no permissions, omit <code>Permissions</code>.</p>
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisTagArgs']]]] tags: <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
                           analysis.</p>
        :param pulumi.Input[str] theme_arn: <p>The ARN of the theme of the analysis.</p>
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AnalysisArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of the AWS::QuickSight::Analysis Resource Type.

        :param str resource_name: The name of the resource.
        :param AnalysisArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AnalysisArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 analysis_id: Optional[pulumi.Input[str]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 errors: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisErrorArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[pulumi.InputType['AnalysisParametersArgs']]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisResourcePermissionArgs']]]]] = None,
                 source_entity: Optional[pulumi.Input[pulumi.InputType['AnalysisSourceEntityArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AnalysisTagArgs']]]]] = None,
                 theme_arn: Optional[pulumi.Input[str]] = None,
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
            __props__ = AnalysisArgs.__new__(AnalysisArgs)

            if analysis_id is None and not opts.urn:
                raise TypeError("Missing required property 'analysis_id'")
            __props__.__dict__["analysis_id"] = analysis_id
            if aws_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'aws_account_id'")
            __props__.__dict__["aws_account_id"] = aws_account_id
            __props__.__dict__["errors"] = errors
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["permissions"] = permissions
            if source_entity is None and not opts.urn:
                raise TypeError("Missing required property 'source_entity'")
            __props__.__dict__["source_entity"] = source_entity
            __props__.__dict__["tags"] = tags
            __props__.__dict__["theme_arn"] = theme_arn
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["data_set_arns"] = None
            __props__.__dict__["last_updated_time"] = None
            __props__.__dict__["sheets"] = None
            __props__.__dict__["status"] = None
        super(Analysis, __self__).__init__(
            'aws-native:quicksight:Analysis',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Analysis':
        """
        Get an existing Analysis resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AnalysisArgs.__new__(AnalysisArgs)

        __props__.__dict__["analysis_id"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["created_time"] = None
        __props__.__dict__["data_set_arns"] = None
        __props__.__dict__["errors"] = None
        __props__.__dict__["last_updated_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["permissions"] = None
        __props__.__dict__["sheets"] = None
        __props__.__dict__["source_entity"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["theme_arn"] = None
        return Analysis(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="analysisId")
    def analysis_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "analysis_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        <p>The Amazon Resource Name (ARN) of the analysis.</p>
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        <p>The time that the analysis was created.</p>
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="dataSetArns")
    def data_set_arns(self) -> pulumi.Output[Sequence[str]]:
        """
        <p>The ARNs of the datasets of the analysis.</p>
        """
        return pulumi.get(self, "data_set_arns")

    @property
    @pulumi.getter
    def errors(self) -> pulumi.Output[Optional[Sequence['outputs.AnalysisError']]]:
        """
        <p>Errors associated with the analysis.</p>
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        """
        <p>The time that the analysis was last updated.</p>
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        <p>The descriptive name of the analysis.</p>
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional['outputs.AnalysisParameters']]:
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.AnalysisResourcePermission']]]:
        """
        <p>A structure that describes the principals and the resource-level permissions on an
                    analysis. You can use the <code>Permissions</code> structure to grant permissions by
                    providing a list of AWS Identity and Access Management (IAM) action information for each
                    principal listed by Amazon Resource Name (ARN). </p>

                <p>To specify no permissions, omit <code>Permissions</code>.</p>
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def sheets(self) -> pulumi.Output[Sequence['outputs.AnalysisSheet']]:
        """
        <p>A list of the associated sheets with the unique identifier and name of each sheet.</p>
        """
        return pulumi.get(self, "sheets")

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> pulumi.Output['outputs.AnalysisSourceEntity']:
        return pulumi.get(self, "source_entity")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['AnalysisResourceStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.AnalysisTag']]]:
        """
        <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
                    analysis.</p>
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="themeArn")
    def theme_arn(self) -> pulumi.Output[Optional[str]]:
        """
        <p>The ARN of the theme of the analysis.</p>
        """
        return pulumi.get(self, "theme_arn")

