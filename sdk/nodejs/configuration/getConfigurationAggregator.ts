// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Config::ConfigurationAggregator
 */
export function getConfigurationAggregator(args: GetConfigurationAggregatorArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationAggregatorResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:configuration:getConfigurationAggregator", {
        "configurationAggregatorName": args.configurationAggregatorName,
    }, opts);
}

export interface GetConfigurationAggregatorArgs {
    /**
     * The name of the aggregator.
     */
    configurationAggregatorName: string;
}

export interface GetConfigurationAggregatorResult {
    readonly accountAggregationSources?: outputs.configuration.ConfigurationAggregatorAccountAggregationSource[];
    /**
     * The Amazon Resource Name (ARN) of the aggregator.
     */
    readonly configurationAggregatorArn?: string;
    readonly organizationAggregationSource?: outputs.configuration.ConfigurationAggregatorOrganizationAggregationSource;
    /**
     * The tags for the configuration aggregator.
     */
    readonly tags?: outputs.configuration.ConfigurationAggregatorTag[];
}

export function getConfigurationAggregatorOutput(args: GetConfigurationAggregatorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConfigurationAggregatorResult> {
    return pulumi.output(args).apply(a => getConfigurationAggregator(a, opts))
}

export interface GetConfigurationAggregatorOutputArgs {
    /**
     * The name of the aggregator.
     */
    configurationAggregatorName: pulumi.Input<string>;
}