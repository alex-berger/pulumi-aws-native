// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    public static class GetLocationSMB
    {
        /// <summary>
        /// Resource schema for AWS::DataSync::LocationSMB.
        /// </summary>
        public static Task<GetLocationSMBResult> InvokeAsync(GetLocationSMBArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLocationSMBResult>("aws-native:datasync:getLocationSMB", args ?? new GetLocationSMBArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataSync::LocationSMB.
        /// </summary>
        public static Output<GetLocationSMBResult> Invoke(GetLocationSMBInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLocationSMBResult>("aws-native:datasync:getLocationSMB", args ?? new GetLocationSMBInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocationSMBArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the SMB location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public string LocationArn { get; set; } = null!;

        public GetLocationSMBArgs()
        {
        }
    }

    public sealed class GetLocationSMBInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the SMB location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public Input<string> LocationArn { get; set; } = null!;

        public GetLocationSMBInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLocationSMBResult
    {
        /// <summary>
        /// The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
        /// </summary>
        public readonly ImmutableArray<string> AgentArns;
        /// <summary>
        /// The name of the Windows domain that the SMB server belongs to.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the SMB location that is created.
        /// </summary>
        public readonly string? LocationArn;
        /// <summary>
        /// The URL of the SMB location that was described.
        /// </summary>
        public readonly string? LocationUri;
        public readonly Outputs.LocationSMBMountOptions? MountOptions;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationSMBTag> Tags;
        /// <summary>
        /// The user who can mount the share, has the permissions to access files and folders in the SMB share.
        /// </summary>
        public readonly string? User;

        [OutputConstructor]
        private GetLocationSMBResult(
            ImmutableArray<string> agentArns,

            string? domain,

            string? locationArn,

            string? locationUri,

            Outputs.LocationSMBMountOptions? mountOptions,

            ImmutableArray<Outputs.LocationSMBTag> tags,

            string? user)
        {
            AgentArns = agentArns;
            Domain = domain;
            LocationArn = locationArn;
            LocationUri = locationUri;
            MountOptions = mountOptions;
            Tags = tags;
            User = user;
        }
    }
}
