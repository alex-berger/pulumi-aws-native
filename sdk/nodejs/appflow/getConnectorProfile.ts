// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppFlow::ConnectorProfile
 */
export function getConnectorProfile(args: GetConnectorProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectorProfileResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appflow:getConnectorProfile", {
        "connectorProfileName": args.connectorProfileName,
    }, opts);
}

export interface GetConnectorProfileArgs {
    /**
     * The maximum number of items to retrieve in a single batch.
     */
    connectorProfileName: string;
}

export interface GetConnectorProfileResult {
    /**
     * Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
     */
    readonly connectionMode?: enums.appflow.ConnectorProfileConnectionMode;
    /**
     * Unique identifier for connector profile resources
     */
    readonly connectorProfileArn?: string;
    /**
     * A unique Arn for Connector-Profile resource
     */
    readonly credentialsArn?: string;
}

export function getConnectorProfileOutput(args: GetConnectorProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectorProfileResult> {
    return pulumi.output(args).apply(a => getConnectorProfile(a, opts))
}

export interface GetConnectorProfileOutputArgs {
    /**
     * The maximum number of items to retrieve in a single batch.
     */
    connectorProfileName: pulumi.Input<string>;
}
