// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Indicates whether a slot can return multiple values.
    /// </summary>
    [OutputType]
    public sealed class BotMultipleValuesSetting
    {
        public readonly bool? AllowMultipleValues;

        [OutputConstructor]
        private BotMultipleValuesSetting(bool? allowMultipleValues)
        {
            AllowMultipleValues = allowMultipleValues;
        }
    }
}
