// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// A function's environment variable settings.
    /// </summary>
    [OutputType]
    public sealed class FunctionEnvironment
    {
        /// <summary>
        /// Environment variable key-value pairs.
        /// </summary>
        public readonly object? Variables;

        [OutputConstructor]
        private FunctionEnvironment(object? variables)
        {
            Variables = variables;
        }
    }
}
