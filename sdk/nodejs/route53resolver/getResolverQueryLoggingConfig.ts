// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.
 */
export function getResolverQueryLoggingConfig(args: GetResolverQueryLoggingConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetResolverQueryLoggingConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53resolver:getResolverQueryLoggingConfig", {
        "id": args.id,
    }, opts);
}

export interface GetResolverQueryLoggingConfigArgs {
    /**
     * ResourceId
     */
    id: string;
}

export interface GetResolverQueryLoggingConfigResult {
    /**
     * Arn
     */
    readonly arn?: string;
    /**
     * Count
     */
    readonly associationCount?: number;
    /**
     * Rfc3339TimeString
     */
    readonly creationTime?: string;
    /**
     * The id of the creator request.
     */
    readonly creatorRequestId?: string;
    /**
     * ResourceId
     */
    readonly id?: string;
    /**
     * AccountId
     */
    readonly ownerId?: string;
    /**
     * ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.
     */
    readonly shareStatus?: enums.route53resolver.ResolverQueryLoggingConfigShareStatus;
    /**
     * ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.
     */
    readonly status?: enums.route53resolver.ResolverQueryLoggingConfigStatus;
}

export function getResolverQueryLoggingConfigOutput(args: GetResolverQueryLoggingConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResolverQueryLoggingConfigResult> {
    return pulumi.output(args).apply(a => getResolverQueryLoggingConfig(a, opts))
}

export interface GetResolverQueryLoggingConfigOutputArgs {
    /**
     * ResourceId
     */
    id: pulumi.Input<string>;
}
