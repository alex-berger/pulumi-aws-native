// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetVpcLink
    {
        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::VpcLink
        /// </summary>
        public static Task<GetVpcLinkResult> InvokeAsync(GetVpcLinkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcLinkResult>("aws-native:apigateway:getVpcLink", args ?? new GetVpcLinkArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ApiGateway::VpcLink
        /// </summary>
        public static Output<GetVpcLinkResult> Invoke(GetVpcLinkInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVpcLinkResult>("aws-native:apigateway:getVpcLink", args ?? new GetVpcLinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcLinkArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetVpcLinkArgs()
        {
        }
    }

    public sealed class GetVpcLinkInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetVpcLinkInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcLinkResult
    {
        public readonly string? Description;
        public readonly string? Id;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.VpcLinkTag> Tags;

        [OutputConstructor]
        private GetVpcLinkResult(
            string? description,

            string? id,

            string? name,

            ImmutableArray<Outputs.VpcLinkTag> tags)
        {
            Description = description;
            Id = id;
            Name = name;
            Tags = tags;
        }
    }
}
