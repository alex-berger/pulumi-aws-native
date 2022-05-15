// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class AppImageConfigKernelSpec
    {
        /// <summary>
        /// The display name of the kernel.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The name of the kernel.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private AppImageConfigKernelSpec(
            string? displayName,

            string name)
        {
            DisplayName = displayName;
            Name = name;
        }
    }
}
