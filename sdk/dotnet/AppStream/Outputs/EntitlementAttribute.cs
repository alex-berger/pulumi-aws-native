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
    public sealed class EntitlementAttribute
    {
        public readonly string Name;
        public readonly string Value;

        [OutputConstructor]
        private EntitlementAttribute(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
