// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An aggregated metric of certain devices in your fleet
 */
export function getFleetMetric(args: GetFleetMetricArgs, opts?: pulumi.InvokeOptions): Promise<GetFleetMetricResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iot:getFleetMetric", {
        "metricName": args.metricName,
    }, opts);
}

export interface GetFleetMetricArgs {
    /**
     * The name of the fleet metric
     */
    metricName: string;
}

export interface GetFleetMetricResult {
    /**
     * The aggregation field to perform aggregation and metric emission
     */
    readonly aggregationField?: string;
    readonly aggregationType?: outputs.iot.FleetMetricAggregationType;
    /**
     * The creation date of a fleet metric
     */
    readonly creationDate?: number;
    /**
     * The description of a fleet metric
     */
    readonly description?: string;
    /**
     * The index name of a fleet metric
     */
    readonly indexName?: string;
    /**
     * The last modified date of a fleet metric
     */
    readonly lastModifiedDate?: number;
    /**
     * The Amazon Resource Number (ARN) of a fleet metric metric
     */
    readonly metricArn?: string;
    /**
     * The period of metric emission in seconds
     */
    readonly period?: number;
    /**
     * The Fleet Indexing query used by a fleet metric
     */
    readonly queryString?: string;
    /**
     * The version of a Fleet Indexing query used by a fleet metric
     */
    readonly queryVersion?: string;
    /**
     * An array of key-value pairs to apply to this resource
     */
    readonly tags?: outputs.iot.FleetMetricTag[];
    /**
     * The unit of data points emitted by a fleet metric
     */
    readonly unit?: string;
    /**
     * The version of a fleet metric
     */
    readonly version?: number;
}

export function getFleetMetricOutput(args: GetFleetMetricOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFleetMetricResult> {
    return pulumi.output(args).apply(a => getFleetMetric(a, opts))
}

export interface GetFleetMetricOutputArgs {
    /**
     * The name of the fleet metric
     */
    metricName: pulumi.Input<string>;
}
