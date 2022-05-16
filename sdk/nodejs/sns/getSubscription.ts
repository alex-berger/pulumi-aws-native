// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SNS::Subscription
 */
export function getSubscription(args: GetSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetSubscriptionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sns:getSubscription", {
        "id": args.id,
    }, opts);
}

export interface GetSubscriptionArgs {
    id: string;
}

export interface GetSubscriptionResult {
    readonly deliveryPolicy?: any;
    readonly filterPolicy?: any;
    readonly id?: string;
    readonly rawMessageDelivery?: boolean;
    readonly redrivePolicy?: any;
    readonly region?: string;
    readonly subscriptionRoleArn?: string;
}

export function getSubscriptionOutput(args: GetSubscriptionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubscriptionResult> {
    return pulumi.output(args).apply(a => getSubscription(a, opts))
}

export interface GetSubscriptionOutputArgs {
    id: pulumi.Input<string>;
}