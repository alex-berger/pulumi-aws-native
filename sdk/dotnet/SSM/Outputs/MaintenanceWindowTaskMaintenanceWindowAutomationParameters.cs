// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowautomationparameters.html
    /// </summary>
    [OutputType]
    public sealed class MaintenanceWindowTaskMaintenanceWindowAutomationParameters
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowautomationparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowautomationparameters-documentversion
        /// </summary>
        public readonly string? DocumentVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-maintenancewindowautomationparameters.html#cfn-ssm-maintenancewindowtask-maintenancewindowautomationparameters-parameters
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string>? Parameters;

        [OutputConstructor]
        private MaintenanceWindowTaskMaintenanceWindowAutomationParameters(
            string? documentVersion,

            Union<System.Text.Json.JsonElement, string>? parameters)
        {
            DocumentVersion = documentVersion;
            Parameters = parameters;
        }
    }
}
