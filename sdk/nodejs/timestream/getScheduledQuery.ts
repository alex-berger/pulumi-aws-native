// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::Timestream::ScheduledQuery resource creates a Timestream Scheduled Query.
 */
export function getScheduledQuery(args: GetScheduledQueryArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduledQueryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:timestream:getScheduledQuery", {
        "arn": args.arn,
    }, opts);
}

export interface GetScheduledQueryArgs {
    arn: string;
}

export interface GetScheduledQueryResult {
    readonly arn?: string;
    /**
     * Configuration for error reporting. Error reports will be generated when a problem is encountered when writing the query results.
     */
    readonly sQErrorReportConfiguration?: string;
    /**
     * The Amazon KMS key used to encrypt the scheduled query resource, at-rest. If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with alias/. If ErrorReportConfiguration uses SSE_KMS as encryption type, the same KmsKeyId is used to encrypt the error report at rest.
     */
    readonly sQKmsKeyId?: string;
    /**
     * The name of the scheduled query. Scheduled query names must be unique within each Region.
     */
    readonly sQName?: string;
    /**
     * Notification configuration for the scheduled query. A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.
     */
    readonly sQNotificationConfiguration?: string;
    /**
     * The query string to run. Parameter names can be specified in the query string @ character followed by an identifier. The named Parameter @scheduled_runtime is reserved and can be used in the query to get the time at which the query is scheduled to run. The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of @scheduled_runtime paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the @scheduled_runtime parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query.
     */
    readonly sQQueryString?: string;
    /**
     * Configuration for when the scheduled query is executed.
     */
    readonly sQScheduleConfiguration?: string;
    /**
     * The ARN for the IAM role that Timestream will assume when running the scheduled query.
     */
    readonly sQScheduledQueryExecutionRoleArn?: string;
    /**
     * Configuration of target store where scheduled query results are written to.
     */
    readonly sQTargetConfiguration?: string;
    readonly tags?: outputs.timestream.ScheduledQueryTag[];
}

export function getScheduledQueryOutput(args: GetScheduledQueryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScheduledQueryResult> {
    return pulumi.output(args).apply(a => getScheduledQuery(a, opts))
}

export interface GetScheduledQueryOutputArgs {
    arn: pulumi.Input<string>;
}
