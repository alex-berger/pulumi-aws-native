// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx
{
    public static class GetFileSystem
    {
        /// <summary>
        /// Resource Type definition for AWS::FSx::FileSystem
        /// </summary>
        public static Task<GetFileSystemResult> InvokeAsync(GetFileSystemArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFileSystemResult>("aws-native:fsx:getFileSystem", args ?? new GetFileSystemArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::FSx::FileSystem
        /// </summary>
        public static Output<GetFileSystemResult> Invoke(GetFileSystemInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFileSystemResult>("aws-native:fsx:getFileSystem", args ?? new GetFileSystemInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileSystemArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetFileSystemArgs()
        {
        }
    }

    public sealed class GetFileSystemInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetFileSystemInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFileSystemResult
    {
        public readonly string? DNSName;
        public readonly string? Id;
        public readonly Outputs.FileSystemLustreConfiguration? LustreConfiguration;
        public readonly string? LustreMountName;
        public readonly Outputs.FileSystemOntapConfiguration? OntapConfiguration;
        public readonly Outputs.FileSystemOpenZFSConfiguration? OpenZFSConfiguration;
        public readonly string? RootVolumeId;
        public readonly int? StorageCapacity;
        public readonly ImmutableArray<Outputs.FileSystemTag> Tags;
        public readonly Outputs.FileSystemWindowsConfiguration? WindowsConfiguration;

        [OutputConstructor]
        private GetFileSystemResult(
            string? dNSName,

            string? id,

            Outputs.FileSystemLustreConfiguration? lustreConfiguration,

            string? lustreMountName,

            Outputs.FileSystemOntapConfiguration? ontapConfiguration,

            Outputs.FileSystemOpenZFSConfiguration? openZFSConfiguration,

            string? rootVolumeId,

            int? storageCapacity,

            ImmutableArray<Outputs.FileSystemTag> tags,

            Outputs.FileSystemWindowsConfiguration? windowsConfiguration)
        {
            DNSName = dNSName;
            Id = id;
            LustreConfiguration = lustreConfiguration;
            LustreMountName = lustreMountName;
            OntapConfiguration = ontapConfiguration;
            OpenZFSConfiguration = openZFSConfiguration;
            RootVolumeId = rootVolumeId;
            StorageCapacity = storageCapacity;
            Tags = tags;
            WindowsConfiguration = windowsConfiguration;
        }
    }
}
