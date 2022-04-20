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
        "functionArn": args.functionArn,
    }, opts);
}

export interface GetUrlArgs {
    /**
     * The full Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    functionArn: string;
}

export interface GetUrlResult {
    /**
     * Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.
     */
    readonly authType?: enums.lambda.UrlAuthType;
    readonly cors?: outputs.lambda.UrlCors;
    /**
     * The full Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    readonly functionArn?: string;
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
     * The full Amazon Resource Name (ARN) of the function associated with the Function URL.
     */
    functionArn: pulumi.Input<string>;
}
