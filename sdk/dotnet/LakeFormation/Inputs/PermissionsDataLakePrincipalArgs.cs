// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-datalakeprincipal.html
    /// </summary>
    public sealed class PermissionsDataLakePrincipalArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-datalakeprincipal.html#cfn-lakeformation-permissions-datalakeprincipal-datalakeprincipalidentifier
        /// </summary>
        [Input("dataLakePrincipalIdentifier")]
        public Input<string>? DataLakePrincipalIdentifier { get; set; }

        public PermissionsDataLakePrincipalArgs()
        {
        }
    }
}
