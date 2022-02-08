// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::Url
 */
export function getUrl(args: GetUrlArgs, opts?: pulumi.InvokeOptions): Promise<GetUrlResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:lambda:getUrl", {
        "arn": args.arn,
    }, opts);
}

export interface GetUrlArgs {
    /**
     * The Amazon Resource Name (ARN) of the Function URL.
     */
    arn: string;
}

export interface GetUrlResult {
    /**
     * The Amazon Resource Name (ARN) of the Function URL.
     */
    readonly arn?: string;
    /**
     * Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
     */
    readonly authorizationType?: enums.lambda.UrlAuthorizationType;
    readonly cors?: outputs.lambda.UrlCors;
    /**
     * The generated url for this resource.
     */
    readonly functionUrl?: string;
}

export function getUrlOutput(args: GetUrlOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUrlResult> {
    return pulumi.output(args).apply(a => getUrl(a, opts))
}

export interface GetUrlOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the Function URL.
     */
    arn: pulumi.Input<string>;
}