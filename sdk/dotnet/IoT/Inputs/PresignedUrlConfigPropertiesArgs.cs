// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// Configuration for pre-signed S3 URLs.
    /// </summary>
    public sealed class PresignedUrlConfigPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("expiresInSec")]
        public Input<int>? ExpiresInSec { get; set; }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public PresignedUrlConfigPropertiesArgs()
        {
        }
    }
}
