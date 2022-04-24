// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudWatch::AnomalyDetector
 *
 * @deprecated AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class AnomalyDetector extends pulumi.CustomResource {
    /**
     * Get an existing AnomalyDetector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AnomalyDetector {
        pulumi.log.warn("AnomalyDetector is deprecated: AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new AnomalyDetector(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudwatch:AnomalyDetector';

    /**
     * Returns true if the given object is an instance of AnomalyDetector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnomalyDetector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnomalyDetector.__pulumiType;
    }

    public readonly configuration!: pulumi.Output<outputs.cloudwatch.AnomalyDetectorConfiguration | undefined>;
    public readonly dimensions!: pulumi.Output<outputs.cloudwatch.AnomalyDetectorDimension[] | undefined>;
    public readonly metricMathAnomalyDetector!: pulumi.Output<outputs.cloudwatch.AnomalyDetectorMetricMathAnomalyDetector | undefined>;
    public readonly metricName!: pulumi.Output<string | undefined>;
    public readonly namespace!: pulumi.Output<string | undefined>;
    public readonly singleMetricAnomalyDetector!: pulumi.Output<outputs.cloudwatch.AnomalyDetectorSingleMetricAnomalyDetector | undefined>;
    public readonly stat!: pulumi.Output<string | undefined>;

    /**
     * Create a AnomalyDetector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: AnomalyDetectorArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("AnomalyDetector is deprecated: AnomalyDetector is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["dimensions"] = args ? args.dimensions : undefined;
            resourceInputs["metricMathAnomalyDetector"] = args ? args.metricMathAnomalyDetector : undefined;
            resourceInputs["metricName"] = args ? args.metricName : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["singleMetricAnomalyDetector"] = args ? args.singleMetricAnomalyDetector : undefined;
            resourceInputs["stat"] = args ? args.stat : undefined;
        } else {
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["dimensions"] = undefined /*out*/;
            resourceInputs["metricMathAnomalyDetector"] = undefined /*out*/;
            resourceInputs["metricName"] = undefined /*out*/;
            resourceInputs["namespace"] = undefined /*out*/;
            resourceInputs["singleMetricAnomalyDetector"] = undefined /*out*/;
            resourceInputs["stat"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AnomalyDetector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AnomalyDetector resource.
 */
export interface AnomalyDetectorArgs {
    configuration?: pulumi.Input<inputs.cloudwatch.AnomalyDetectorConfigurationArgs>;
    dimensions?: pulumi.Input<pulumi.Input<inputs.cloudwatch.AnomalyDetectorDimensionArgs>[]>;
    metricMathAnomalyDetector?: pulumi.Input<inputs.cloudwatch.AnomalyDetectorMetricMathAnomalyDetectorArgs>;
    metricName?: pulumi.Input<string>;
    namespace?: pulumi.Input<string>;
    singleMetricAnomalyDetector?: pulumi.Input<inputs.cloudwatch.AnomalyDetectorSingleMetricAnomalyDetectorArgs>;
    stat?: pulumi.Input<string>;
}
