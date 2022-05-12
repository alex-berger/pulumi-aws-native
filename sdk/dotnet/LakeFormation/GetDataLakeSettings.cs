// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation
{
    public static class GetDataLakeSettings
    {
        /// <summary>
        /// Resource Type definition for AWS::LakeFormation::DataLakeSettings
        /// </summary>
        public static Task<GetDataLakeSettingsResult> InvokeAsync(GetDataLakeSettingsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataLakeSettingsResult>("aws-native:lakeformation:getDataLakeSettings", args ?? new GetDataLakeSettingsArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::LakeFormation::DataLakeSettings
        /// </summary>
        public static Output<GetDataLakeSettingsResult> Invoke(GetDataLakeSettingsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDataLakeSettingsResult>("aws-native:lakeformation:getDataLakeSettings", args ?? new GetDataLakeSettingsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataLakeSettingsArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDataLakeSettingsArgs()
        {
        }
    }

    public sealed class GetDataLakeSettingsInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDataLakeSettingsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataLakeSettingsResult
    {
        public readonly Outputs.DataLakeSettingsAdmins? Admins;
        public readonly string? Id;
        public readonly ImmutableArray<string> TrustedResourceOwners;

        [OutputConstructor]
        private GetDataLakeSettingsResult(
            Outputs.DataLakeSettingsAdmins? admins,

            string? id,

            ImmutableArray<string> trustedResourceOwners)
        {
            Admins = admins;
            Id = id;
            TrustedResourceOwners = trustedResourceOwners;
        }
    }
}
