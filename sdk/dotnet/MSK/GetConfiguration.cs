// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MSK
{
    public static class GetConfiguration
    {
        /// <summary>
        /// Resource Type definition for AWS::MSK::Configuration
        /// </summary>
        public static Task<GetConfigurationResult> InvokeAsync(GetConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigurationResult>("aws-native:msk:getConfiguration", args ?? new GetConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MSK::Configuration
        /// </summary>
        public static Output<GetConfigurationResult> Invoke(GetConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConfigurationResult>("aws-native:msk:getConfiguration", args ?? new GetConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigurationArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetConfigurationArgs()
        {
        }
    }

    public sealed class GetConfigurationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetConfigurationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConfigurationResult
    {
        public readonly string? Arn;
        public readonly string? Description;

        [OutputConstructor]
        private GetConfigurationResult(
            string? arn,

            string? description)
        {
            Arn = arn;
            Description = description;
        }
    }
}
