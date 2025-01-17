// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::AppRunner::ObservabilityConfiguration resource  is an AWS App Runner resource type that specifies an App Runner observability configuration
 */
export function getObservabilityConfiguration(args: GetObservabilityConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetObservabilityConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apprunner:getObservabilityConfiguration", {
        "observabilityConfigurationArn": args.observabilityConfigurationArn,
    }, opts);
}

export interface GetObservabilityConfigurationArgs {
    /**
     * The Amazon Resource Name (ARN) of this ObservabilityConfiguration
     */
    observabilityConfigurationArn: string;
}

export interface GetObservabilityConfigurationResult {
    /**
     * It's set to true for the configuration with the highest Revision among all configurations that share the same Name. It's set to false otherwise.
     */
    readonly latest?: boolean;
    /**
     * The Amazon Resource Name (ARN) of this ObservabilityConfiguration
     */
    readonly observabilityConfigurationArn?: string;
    /**
     * The revision of this observability configuration. It's unique among all the active configurations ('Status': 'ACTIVE') that share the same ObservabilityConfigurationName.
     */
    readonly observabilityConfigurationRevision?: number;
}

export function getObservabilityConfigurationOutput(args: GetObservabilityConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetObservabilityConfigurationResult> {
    return pulumi.output(args).apply(a => getObservabilityConfiguration(a, opts))
}

export interface GetObservabilityConfigurationOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of this ObservabilityConfiguration
     */
    observabilityConfigurationArn: pulumi.Input<string>;
}
