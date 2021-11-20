// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class EndpointConfigProductionVariant
    {
        public readonly string? AcceleratorType;
        public readonly int? InitialInstanceCount;
        public readonly double InitialVariantWeight;
        public readonly string? InstanceType;
        public readonly string ModelName;
        public readonly string VariantName;

        [OutputConstructor]
        private EndpointConfigProductionVariant(
            string? acceleratorType,

            int? initialInstanceCount,

            double initialVariantWeight,

            string? instanceType,

            string modelName,

            string variantName)
        {
            AcceleratorType = acceleratorType;
            InitialInstanceCount = initialInstanceCount;
            InitialVariantWeight = initialVariantWeight;
            InstanceType = instanceType;
            ModelName = modelName;
            VariantName = variantName;
        }
    }
}
