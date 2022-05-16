// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::RDS::DBProxyEndpoint.
 */
export function getDBProxyEndpoint(args: GetDBProxyEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetDBProxyEndpointResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:rds:getDBProxyEndpoint", {
        "dBProxyEndpointName": args.dBProxyEndpointName,
    }, opts);
}

export interface GetDBProxyEndpointArgs {
    /**
     * The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
     */
    dBProxyEndpointName: string;
}

export interface GetDBProxyEndpointResult {
    /**
     * The Amazon Resource Name (ARN) for the DB proxy endpoint.
     */
    readonly dBProxyEndpointArn?: string;
    /**
     * The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.
     */
    readonly endpoint?: string;
    /**
     * A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.
     */
    readonly isDefault?: boolean;
    /**
     * An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
     */
    readonly tags?: outputs.rds.DBProxyEndpointTagFormat[];
    /**
     * VPC ID to associate with the new DB proxy endpoint.
     */
    readonly vpcId?: string;
    /**
     * VPC security group IDs to associate with the new DB proxy endpoint.
     */
    readonly vpcSecurityGroupIds?: string[];
}

export function getDBProxyEndpointOutput(args: GetDBProxyEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDBProxyEndpointResult> {
    return pulumi.output(args).apply(a => getDBProxyEndpoint(a, opts))
}

export interface GetDBProxyEndpointOutputArgs {
    /**
     * The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
     */
    dBProxyEndpointName: pulumi.Input<string>;
}
