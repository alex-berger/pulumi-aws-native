// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Outputs
{

    /// <summary>
    /// The destination repository for the container image.
    /// </summary>
    [OutputType]
    public sealed class DistributionConfigurationTargetContainerRepository
    {
        /// <summary>
        /// The repository name of target container repository.
        /// </summary>
        public readonly string? RepositoryName;
        /// <summary>
        /// The service of target container repository.
        /// </summary>
        public readonly Pulumi.AwsNative.ImageBuilder.DistributionConfigurationTargetContainerRepositoryService? Service;

        [OutputConstructor]
        private DistributionConfigurationTargetContainerRepository(
            string? repositoryName,

            Pulumi.AwsNative.ImageBuilder.DistributionConfigurationTargetContainerRepositoryService? service)
        {
            RepositoryName = repositoryName;
            Service = service;
        }
    }
}