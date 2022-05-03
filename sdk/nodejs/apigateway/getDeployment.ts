// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::Deployment
 */
export function getDeployment(args: GetDeploymentArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getDeployment", {
        "deploymentId": args.deploymentId,
        "restApiId": args.restApiId,
    }, opts);
}

export interface GetDeploymentArgs {
    /**
     * Primary Id for this resource
     */
    deploymentId: string;
    /**
     * The ID of the RestApi resource to deploy. 
     */
    restApiId: string;
}

export interface GetDeploymentResult {
    /**
     * Primary Id for this resource
     */
    readonly deploymentId?: string;
    /**
     * A description of the purpose of the API Gateway deployment.
     */
    readonly description?: string;
}

export function getDeploymentOutput(args: GetDeploymentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeploymentResult> {
    return pulumi.output(args).apply(a => getDeployment(a, opts))
}

export interface GetDeploymentOutputArgs {
    /**
     * Primary Id for this resource
     */
    deploymentId: pulumi.Input<string>;
    /**
     * The ID of the RestApi resource to deploy. 
     */
    restApiId: pulumi.Input<string>;
}
