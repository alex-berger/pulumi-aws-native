// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CertificateManager::Certificate
 *
 * @deprecated Certificate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Certificate {
        pulumi.log.warn("Certificate is deprecated: Certificate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Certificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:certificatemanager:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    public readonly certificateAuthorityArn!: pulumi.Output<string | undefined>;
    public readonly certificateTransparencyLoggingPreference!: pulumi.Output<string | undefined>;
    public readonly domainName!: pulumi.Output<string>;
    public readonly domainValidationOptions!: pulumi.Output<outputs.certificatemanager.CertificateDomainValidationOption[] | undefined>;
    public readonly subjectAlternativeNames!: pulumi.Output<string[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.certificatemanager.CertificateTag[] | undefined>;
    public readonly validationMethod!: pulumi.Output<string | undefined>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Certificate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Certificate is deprecated: Certificate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            resourceInputs["certificateAuthorityArn"] = args ? args.certificateAuthorityArn : undefined;
            resourceInputs["certificateTransparencyLoggingPreference"] = args ? args.certificateTransparencyLoggingPreference : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["domainValidationOptions"] = args ? args.domainValidationOptions : undefined;
            resourceInputs["subjectAlternativeNames"] = args ? args.subjectAlternativeNames : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["validationMethod"] = args ? args.validationMethod : undefined;
        } else {
            resourceInputs["certificateAuthorityArn"] = undefined /*out*/;
            resourceInputs["certificateTransparencyLoggingPreference"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["domainValidationOptions"] = undefined /*out*/;
            resourceInputs["subjectAlternativeNames"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["validationMethod"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Certificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    certificateAuthorityArn?: pulumi.Input<string>;
    certificateTransparencyLoggingPreference?: pulumi.Input<string>;
    domainName: pulumi.Input<string>;
    domainValidationOptions?: pulumi.Input<pulumi.Input<inputs.certificatemanager.CertificateDomainValidationOptionArgs>[]>;
    subjectAlternativeNames?: pulumi.Input<pulumi.Input<string>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.certificatemanager.CertificateTagArgs>[]>;
    validationMethod?: pulumi.Input<string>;
}