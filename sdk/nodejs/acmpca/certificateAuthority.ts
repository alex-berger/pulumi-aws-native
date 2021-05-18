// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html
 */
export class CertificateAuthority extends pulumi.CustomResource {
    /**
     * Get an existing CertificateAuthority resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CertificateAuthority {
        return new CertificateAuthority(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ACMPCA:CertificateAuthority';

    /**
     * Returns true if the given object is an instance of CertificateAuthority.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateAuthority {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateAuthority.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly certificateSigningRequest!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-keyalgorithm
     */
    public readonly keyAlgorithm!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-revocationconfiguration
     */
    public readonly revocationConfiguration!: pulumi.Output<outputs.ACMPCA.CertificateAuthorityRevocationConfiguration | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-signingalgorithm
     */
    public readonly signingAlgorithm!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-subject
     */
    public readonly subject!: pulumi.Output<outputs.ACMPCA.CertificateAuthoritySubject>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-type
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a CertificateAuthority resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateAuthorityArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.keyAlgorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyAlgorithm'");
            }
            if ((!args || args.signingAlgorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'signingAlgorithm'");
            }
            if ((!args || args.subject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subject'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["keyAlgorithm"] = args ? args.keyAlgorithm : undefined;
            inputs["revocationConfiguration"] = args ? args.revocationConfiguration : undefined;
            inputs["signingAlgorithm"] = args ? args.signingAlgorithm : undefined;
            inputs["subject"] = args ? args.subject : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificateSigningRequest"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["certificateSigningRequest"] = undefined /*out*/;
            inputs["keyAlgorithm"] = undefined /*out*/;
            inputs["revocationConfiguration"] = undefined /*out*/;
            inputs["signingAlgorithm"] = undefined /*out*/;
            inputs["subject"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CertificateAuthority.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CertificateAuthority resource.
 */
export interface CertificateAuthorityArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-keyalgorithm
     */
    readonly keyAlgorithm: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-revocationconfiguration
     */
    readonly revocationConfiguration?: pulumi.Input<inputs.ACMPCA.CertificateAuthorityRevocationConfigurationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-signingalgorithm
     */
    readonly signingAlgorithm: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-subject
     */
    readonly subject: pulumi.Input<inputs.ACMPCA.CertificateAuthoritySubjectArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-tags
     */
    readonly tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthority.html#cfn-acmpca-certificateauthority-type
     */
    readonly type: pulumi.Input<string>;
}
