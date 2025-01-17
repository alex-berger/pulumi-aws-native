// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class ComponentVersionLambdaVolumeMount
    {
        public readonly bool? AddGroupOwner;
        public readonly string? DestinationPath;
        public readonly Pulumi.AwsNative.GreengrassV2.ComponentVersionLambdaFilesystemPermission? Permission;
        public readonly string? SourcePath;

        [OutputConstructor]
        private ComponentVersionLambdaVolumeMount(
            bool? addGroupOwner,

            string? destinationPath,

            Pulumi.AwsNative.GreengrassV2.ComponentVersionLambdaFilesystemPermission? permission,

            string? sourcePath)
        {
            AddGroupOwner = addGroupOwner;
            DestinationPath = destinationPath;
            Permission = permission;
            SourcePath = sourcePath;
        }
    }
}
