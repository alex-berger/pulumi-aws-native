# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['SessionArgs', 'Session']

@pulumi.input_type
class SessionArgs:
    def __init__(__self__, *,
                 finding_publishing_frequency: Optional[pulumi.Input['SessionFindingPublishingFrequency']] = None,
                 status: Optional[pulumi.Input['SessionStatus']] = None):
        """
        The set of arguments for constructing a Session resource.
        :param pulumi.Input['SessionFindingPublishingFrequency'] finding_publishing_frequency: A enumeration value that specifies how frequently finding updates are published.
        :param pulumi.Input['SessionStatus'] status: A enumeration value that specifies the status of the Macie Session.
        """
        if finding_publishing_frequency is not None:
            pulumi.set(__self__, "finding_publishing_frequency", finding_publishing_frequency)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="findingPublishingFrequency")
    def finding_publishing_frequency(self) -> Optional[pulumi.Input['SessionFindingPublishingFrequency']]:
        """
        A enumeration value that specifies how frequently finding updates are published.
        """
        return pulumi.get(self, "finding_publishing_frequency")

    @finding_publishing_frequency.setter
    def finding_publishing_frequency(self, value: Optional[pulumi.Input['SessionFindingPublishingFrequency']]):
        pulumi.set(self, "finding_publishing_frequency", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['SessionStatus']]:
        """
        A enumeration value that specifies the status of the Macie Session.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['SessionStatus']]):
        pulumi.set(self, "status", value)


class Session(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 finding_publishing_frequency: Optional[pulumi.Input['SessionFindingPublishingFrequency']] = None,
                 status: Optional[pulumi.Input['SessionStatus']] = None,
                 __props__=None):
        """
        The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['SessionFindingPublishingFrequency'] finding_publishing_frequency: A enumeration value that specifies how frequently finding updates are published.
        :param pulumi.Input['SessionStatus'] status: A enumeration value that specifies the status of the Macie Session.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SessionArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.

        :param str resource_name: The name of the resource.
        :param SessionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SessionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 finding_publishing_frequency: Optional[pulumi.Input['SessionFindingPublishingFrequency']] = None,
                 status: Optional[pulumi.Input['SessionStatus']] = None,
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
            __props__ = SessionArgs.__new__(SessionArgs)

            __props__.__dict__["finding_publishing_frequency"] = finding_publishing_frequency
            __props__.__dict__["status"] = status
            __props__.__dict__["aws_account_id"] = None
            __props__.__dict__["service_role"] = None
        super(Session, __self__).__init__(
            'aws-native:macie:Session',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Session':
        """
        Get an existing Session resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SessionArgs.__new__(SessionArgs)

        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["finding_publishing_frequency"] = None
        __props__.__dict__["service_role"] = None
        __props__.__dict__["status"] = None
        return Session(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        """
        AWS account ID of customer
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="findingPublishingFrequency")
    def finding_publishing_frequency(self) -> pulumi.Output[Optional['SessionFindingPublishingFrequency']]:
        """
        A enumeration value that specifies how frequently finding updates are published.
        """
        return pulumi.get(self, "finding_publishing_frequency")

    @property
    @pulumi.getter(name="serviceRole")
    def service_role(self) -> pulumi.Output[str]:
        """
        Service role used by Macie
        """
        return pulumi.get(self, "service_role")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['SessionStatus']]:
        """
        A enumeration value that specifies the status of the Macie Session.
        """
        return pulumi.get(self, "status")

