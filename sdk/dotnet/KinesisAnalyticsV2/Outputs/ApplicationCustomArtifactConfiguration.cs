// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    [OutputType]
    public sealed class ApplicationCustomArtifactConfiguration
    {
        public readonly string ArtifactType;
        public readonly Outputs.ApplicationMavenReference? MavenReference;
        public readonly Outputs.ApplicationS3ContentLocation? S3ContentLocation;

        [OutputConstructor]
        private ApplicationCustomArtifactConfiguration(
            string artifactType,

            Outputs.ApplicationMavenReference? mavenReference,

            Outputs.ApplicationS3ContentLocation? s3ContentLocation)
        {
            ArtifactType = artifactType;
            MavenReference = mavenReference;
            S3ContentLocation = s3ContentLocation;
        }
    }
}