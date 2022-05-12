// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAF
{
    public static class GetSizeConstraintSet
    {
        /// <summary>
        /// Resource Type definition for AWS::WAF::SizeConstraintSet
        /// </summary>
        public static Task<GetSizeConstraintSetResult> InvokeAsync(GetSizeConstraintSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSizeConstraintSetResult>("aws-native:waf:getSizeConstraintSet", args ?? new GetSizeConstraintSetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::WAF::SizeConstraintSet
        /// </summary>
        public static Output<GetSizeConstraintSetResult> Invoke(GetSizeConstraintSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSizeConstraintSetResult>("aws-native:waf:getSizeConstraintSet", args ?? new GetSizeConstraintSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSizeConstraintSetArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSizeConstraintSetArgs()
        {
        }
    }

    public sealed class GetSizeConstraintSetInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSizeConstraintSetInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSizeConstraintSetResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.SizeConstraintSetSizeConstraint> SizeConstraints;

        [OutputConstructor]
        private GetSizeConstraintSetResult(
            string? id,

            ImmutableArray<Outputs.SizeConstraintSetSizeConstraint> sizeConstraints)
        {
            Id = id;
            SizeConstraints = sizeConstraints;
        }
    }
}
