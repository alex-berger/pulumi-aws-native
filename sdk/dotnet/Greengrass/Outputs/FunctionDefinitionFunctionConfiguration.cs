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
    public sealed class FunctionDefinitionFunctionConfiguration
    {
        public readonly string? EncodingType;
        public readonly Outputs.FunctionDefinitionEnvironment? Environment;
        public readonly string? ExecArgs;
        public readonly string? Executable;
        public readonly int? MemorySize;
        public readonly bool? Pinned;
        public readonly int? Timeout;

        [OutputConstructor]
        private FunctionDefinitionFunctionConfiguration(
            string? encodingType,

            Outputs.FunctionDefinitionEnvironment? environment,

            string? execArgs,

            string? executable,

            int? memorySize,

            bool? pinned,

            int? timeout)
        {
            EncodingType = encodingType;
            Environment = environment;
            ExecArgs = execArgs;
            Executable = executable;
            MemorySize = memorySize;
            Pinned = pinned;
            Timeout = timeout;
        }
    }
}