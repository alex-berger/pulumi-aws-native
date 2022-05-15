// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::SES::ContactList.
 */
export function getContactList(args: GetContactListArgs, opts?: pulumi.InvokeOptions): Promise<GetContactListResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ses:getContactList", {
        "contactListName": args.contactListName,
    }, opts);
}

export interface GetContactListArgs {
    /**
     * The name of the contact list.
     */
    contactListName: string;
}

export interface GetContactListResult {
    /**
     * The description of the contact list.
     */
    readonly description?: string;
    /**
     * The tags (keys and values) associated with the contact list.
     */
    readonly tags?: outputs.ses.ContactListTag[];
    /**
     * The topics associated with the contact list.
     */
    readonly topics?: outputs.ses.ContactListTopic[];
}

export function getContactListOutput(args: GetContactListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetContactListResult> {
    return pulumi.output(args).apply(a => getContactList(a, opts))
}

export interface GetContactListOutputArgs {
    /**
     * The name of the contact list.
     */
    contactListName: pulumi.Input<string>;
}
