# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'KeyKeySpec',
    'KeyKeyUsage',
]


class KeyKeySpec(str, Enum):
    """
    Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.
    """
    SYMMETRIC_DEFAULT = "SYMMETRIC_DEFAULT"
    RSA2048 = "RSA_2048"
    RSA3072 = "RSA_3072"
    RSA4096 = "RSA_4096"
    ECC_NIST_P256 = "ECC_NIST_P256"
    ECC_NIST_P384 = "ECC_NIST_P384"
    ECC_NIST_P521 = "ECC_NIST_P521"
    ECC_SECG_P256K1 = "ECC_SECG_P256K1"


class KeyKeyUsage(str, Enum):
    """
    Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.
    """
    ENCRYPT_DECRYPT = "ENCRYPT_DECRYPT"
    SIGN_VERIFY = "SIGN_VERIFY"