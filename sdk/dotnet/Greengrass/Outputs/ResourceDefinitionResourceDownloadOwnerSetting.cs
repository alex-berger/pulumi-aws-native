// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Outputs
{

    [OutputType]
    public sealed class ResourceDefinitionResourceDownloadOwnerSetting
    {
        public readonly string GroupOwner;
        public readonly string GroupPermission;

        [OutputConstructor]
        private ResourceDefinitionResourceDownloadOwnerSetting(
            string groupOwner,

            string groupPermission)
        {
            GroupOwner = groupOwner;
            GroupPermission = groupPermission;
        }
    }
}