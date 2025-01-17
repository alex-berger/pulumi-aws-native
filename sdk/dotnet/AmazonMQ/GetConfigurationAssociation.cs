// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmazonMQ
{
    public static class GetConfigurationAssociation
    {
        /// <summary>
        /// Resource Type definition for AWS::AmazonMQ::ConfigurationAssociation
        /// </summary>
        public static Task<GetConfigurationAssociationResult> InvokeAsync(GetConfigurationAssociationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationAssociationResult>("aws-native:amazonmq:getConfigurationAssociation", args ?? new GetConfigurationAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AmazonMQ::ConfigurationAssociation
        /// </summary>
        public static Output<GetConfigurationAssociationResult> Invoke(GetConfigurationAssociationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConfigurationAssociationResult>("aws-native:amazonmq:getConfigurationAssociation", args ?? new GetConfigurationAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationAssociationArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetConfigurationAssociationArgs()
        {
        }
    }

    public sealed class GetConfigurationAssociationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetConfigurationAssociationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigurationAssociationResult
    {
        public readonly Outputs.ConfigurationAssociationConfigurationId? Configuration;
        public readonly string? Id;

        [OutputConstructor]
        private GetConfigurationAssociationResult(
            Outputs.ConfigurationAssociationConfigurationId? configuration,

            string? id)
        {
            Configuration = configuration;
            Id = id;
        }
    }
}
