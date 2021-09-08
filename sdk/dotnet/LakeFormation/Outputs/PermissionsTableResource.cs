// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html
    /// </summary>
    [OutputType]
    public sealed class PermissionsTableResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html#cfn-lakeformation-permissions-tableresource-catalogid
        /// </summary>
        public readonly string? CatalogId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html#cfn-lakeformation-permissions-tableresource-databasename
        /// </summary>
        public readonly string? DatabaseName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html#cfn-lakeformation-permissions-tableresource-name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html#cfn-lakeformation-permissions-tableresource-tablewildcard
        /// </summary>
        public readonly Outputs.PermissionsTableWildcard? TableWildcard;

        [OutputConstructor]
        private PermissionsTableResource(
            string? catalogId,

            string? databaseName,

            string? name,

            Outputs.PermissionsTableWildcard? tableWildcard)
        {
            CatalogId = catalogId;
            DatabaseName = databaseName;
            Name = name;
            TableWildcard = tableWildcard;
        }
    }
}
