// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Outputs
{

    [OutputType]
    public sealed class MaintenanceWindowTaskMaintenanceWindowStepFunctionsParameters
    {
        public readonly string? Input;
        public readonly string? Name;

        [OutputConstructor]
        private MaintenanceWindowTaskMaintenanceWindowStepFunctionsParameters(
            string? input,

            string? name)
        {
            Input = input;
            Name = name;
        }
    }
}
