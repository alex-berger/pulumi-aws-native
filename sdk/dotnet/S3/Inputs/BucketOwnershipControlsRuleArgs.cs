// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketOwnershipControlsRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies an object ownership rule.
        /// </summary>
        [Input("objectOwnership")]
        public Input<Pulumi.AwsNative.S3.BucketOwnershipControlsRuleObjectOwnership>? ObjectOwnership { get; set; }

        public BucketOwnershipControlsRuleArgs()
        {
        }
    }
}