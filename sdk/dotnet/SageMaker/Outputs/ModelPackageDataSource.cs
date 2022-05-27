// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Describes the input source of a transform job and the way the transform job consumes it.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageDataSource
    {
        public readonly Outputs.ModelPackageS3DataSource S3DataSource;

        [OutputConstructor]
        private ModelPackageDataSource(Outputs.ModelPackageS3DataSource s3DataSource)
        {
            S3DataSource = s3DataSource;
        }
    }
}