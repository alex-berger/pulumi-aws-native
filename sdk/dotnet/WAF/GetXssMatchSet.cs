// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAF
{
    public static class GetXssMatchSet
    {
        /// <summary>
        /// Resource Type definition for AWS::WAF::XssMatchSet
        /// </summary>
        public static Task<GetXssMatchSetResult> InvokeAsync(GetXssMatchSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetXssMatchSetResult>("aws-native:waf:getXssMatchSet", args ?? new GetXssMatchSetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::WAF::XssMatchSet
        /// </summary>
        public static Output<GetXssMatchSetResult> Invoke(GetXssMatchSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetXssMatchSetResult>("aws-native:waf:getXssMatchSet", args ?? new GetXssMatchSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetXssMatchSetArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetXssMatchSetArgs()
        {
        }
    }

    public sealed class GetXssMatchSetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetXssMatchSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetXssMatchSetResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.XssMatchSetXssMatchTuple> XssMatchTuples;

        [OutputConstructor]
        private GetXssMatchSetResult(
            string? id,

            ImmutableArray<Outputs.XssMatchSetXssMatchTuple> xssMatchTuples)
        {
            Id = id;
            XssMatchTuples = xssMatchTuples;
        }
    }
}
