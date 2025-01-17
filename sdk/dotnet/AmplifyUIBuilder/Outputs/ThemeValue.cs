// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUIBuilder.Outputs
{

    [OutputType]
    public sealed class ThemeValue
    {
        public readonly ImmutableArray<Outputs.ThemeValues> Children;
        public readonly string? Value;

        [OutputConstructor]
        private ThemeValue(
            ImmutableArray<Outputs.ThemeValues> children,

            string? value)
        {
            Children = children;
            Value = value;
        }
    }
}
