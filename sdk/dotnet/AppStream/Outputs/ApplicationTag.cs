// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Outputs
{

    [OutputType]
    public sealed class ApplicationTag
    {
        public readonly string TagKey;
        public readonly string TagValue;

        [OutputConstructor]
        private ApplicationTag(
            string tagKey,

            string tagValue)
        {
            TagKey = tagKey;
            TagValue = tagValue;
        }
    }
}
