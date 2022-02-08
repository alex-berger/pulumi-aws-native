// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SDB
{
    public static class GetDomain
    {
        /// <summary>
        /// Resource Type definition for AWS::SDB::Domain
        /// </summary>
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("aws-native:sdb:getDomain", args ?? new GetDomainArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SDB::Domain
        /// </summary>
        public static Output<GetDomainResult> Invoke(GetDomainInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDomainResult>("aws-native:sdb:getDomain", args ?? new GetDomainInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDomainArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDomainArgs()
        {
        }
    }

    public sealed class GetDomainInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDomainInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainResult
    {
        public readonly string? Description;
        public readonly string? Id;

        [OutputConstructor]
        private GetDomainResult(
            string? description,

            string? id)
        {
            Description = description;
            Id = id;
        }
    }
}