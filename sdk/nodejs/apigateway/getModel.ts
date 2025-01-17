// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::Model
 */
export function getModel(args: GetModelArgs, opts?: pulumi.InvokeOptions): Promise<GetModelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getModel", {
        "name": args.name,
        "restApiId": args.restApiId,
    }, opts);
}

export interface GetModelArgs {
    /**
     * A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
     */
    name: string;
    /**
     * The ID of a REST API with which to associate this model.
     */
    restApiId: string;
}

export interface GetModelResult {
    /**
     * A description that identifies this model.
     */
    readonly description?: string;
    /**
     * The schema to use to transform data to one or more output formats. Specify null ({}) if you don't want to specify a schema.
     */
    readonly schema?: any;
}

export function getModelOutput(args: GetModelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelResult> {
    return pulumi.output(args).apply(a => getModel(a, opts))
}

export interface GetModelOutputArgs {
    /**
     * A name for the model. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the model name.
     */
    name: pulumi.Input<string>;
    /**
     * The ID of a REST API with which to associate this model.
     */
    restApiId: pulumi.Input<string>;
}
