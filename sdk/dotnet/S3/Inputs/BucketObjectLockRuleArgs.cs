// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// The Object Lock rule in place for the specified object.
    /// </summary>
    public sealed class BucketObjectLockRuleArgs : Pulumi.ResourceArgs
    {
        [Input("defaultRetention")]
        public Input<Inputs.BucketDefaultRetentionArgs>? DefaultRetention { get; set; }

        public BucketObjectLockRuleArgs()
        {
        }
    }
}