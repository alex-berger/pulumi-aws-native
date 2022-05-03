// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECR.Inputs
{

    /// <summary>
    /// The image scanning configuration for the repository. This setting determines whether images are scanned for known vulnerabilities after being pushed to the repository.
    /// </summary>
    public sealed class RepositoryImageScanningConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("scanOnPush")]
        public Input<bool>? ScanOnPush { get; set; }

        public RepositoryImageScanningConfigurationArgs()
        {
        }
    }
}
