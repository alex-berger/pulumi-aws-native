// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html
    /// </summary>
    [AwsNativeResourceType("aws-native:lakeformation:DataLakeSettings")]
    public partial class DataLakeSettings : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-admins
        /// </summary>
        [Output("admins")]
        public Output<Outputs.DataLakeSettingsAdmins?> Admins { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-trustedresourceowners
        /// </summary>
        [Output("trustedResourceOwners")]
        public Output<ImmutableArray<string>> TrustedResourceOwners { get; private set; } = null!;


        /// <summary>
        /// Create a DataLakeSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataLakeSettings(string name, DataLakeSettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:lakeformation:DataLakeSettings", name, args ?? new DataLakeSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataLakeSettings(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:lakeformation:DataLakeSettings", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataLakeSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataLakeSettings Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataLakeSettings(name, id, options);
        }
    }

    public sealed class DataLakeSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-admins
        /// </summary>
        [Input("admins")]
        public Input<Inputs.DataLakeSettingsAdminsArgs>? Admins { get; set; }

        [Input("trustedResourceOwners")]
        private InputList<string>? _trustedResourceOwners;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html#cfn-lakeformation-datalakesettings-trustedresourceowners
        /// </summary>
        public InputList<string> TrustedResourceOwners
        {
            get => _trustedResourceOwners ?? (_trustedResourceOwners = new InputList<string>());
            set => _trustedResourceOwners = value;
        }

        public DataLakeSettingsArgs()
        {
        }
    }
}
