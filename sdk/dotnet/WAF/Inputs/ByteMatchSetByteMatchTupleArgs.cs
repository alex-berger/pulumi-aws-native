// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAF.Inputs
{

    public sealed class ByteMatchSetByteMatchTupleArgs : Pulumi.ResourceArgs
    {
        [Input("fieldToMatch", required: true)]
        public Input<Inputs.ByteMatchSetFieldToMatchArgs> FieldToMatch { get; set; } = null!;

        [Input("positionalConstraint", required: true)]
        public Input<string> PositionalConstraint { get; set; } = null!;

        [Input("targetString")]
        public Input<string>? TargetString { get; set; }

        [Input("targetStringBase64")]
        public Input<string>? TargetStringBase64 { get; set; }

        [Input("textTransformation", required: true)]
        public Input<string> TextTransformation { get; set; } = null!;

        public ByteMatchSetByteMatchTupleArgs()
        {
        }
    }
}
