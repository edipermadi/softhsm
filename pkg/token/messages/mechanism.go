package messages

type MechanismType int

const (
	/* the following mechanism types are defined: */
	MechanismType_RSA_PKCS_KEY_PAIR_GEN MechanismType = 0x00000000
	MechanismType_RSA_PKCS              MechanismType = 0x00000001
	MechanismType_RSA_9796              MechanismType = 0x00000002
	MechanismType_RSA_X_509             MechanismType = 0x00000003

	MechanismType_MD2_RSA_PKCS  MechanismType = 0x00000004
	MechanismType_MD5_RSA_PKCS  MechanismType = 0x00000005
	MechanismType_SHA1_RSA_PKCS MechanismType = 0x00000006

	MechanismType_RIPEMD128_RSA_PKCS MechanismType = 0x00000007
	MechanismType_RIPEMD160_RSA_PKCS MechanismType = 0x00000008
	MechanismType_RSA_PKCS_OAEP      MechanismType = 0x00000009

	MechanismType_RSA_X9_31_KEY_PAIR_GEN MechanismType = 0x0000000A
	MechanismType_RSA_X9_31              MechanismType = 0x0000000B
	MechanismType_SHA1_RSA_X9_31         MechanismType = 0x0000000C
	MechanismType_RSA_PKCS_PSS           MechanismType = 0x0000000D
	MechanismType_SHA1_RSA_PKCS_PSS      MechanismType = 0x0000000E

	MechanismType_DSA_KEY_PAIR_GEN MechanismType = 0x00000010
	MechanismType_DSA              MechanismType = 0x00000011
	MechanismType_DSA_SHA1         MechanismType = 0x00000012
	MechanismType_DSA_SHA224       MechanismType = 0x00000013
	MechanismType_DSA_SHA256       MechanismType = 0x00000014
	MechanismType_DSA_SHA384       MechanismType = 0x00000015
	MechanismType_DSA_SHA512       MechanismType = 0x00000016
	MechanismType_DSA_SHA3_224     MechanismType = 0x00000018
	MechanismType_DSA_SHA3_256     MechanismType = 0x00000019
	MechanismType_DSA_SHA3_384     MechanismType = 0x0000001A
	MechanismType_DSA_SHA3_512     MechanismType = 0x0000001B

	MechanismType_DH_PKCS_KEY_PAIR_GEN MechanismType = 0x00000020
	MechanismType_DH_PKCS_DERIVE       MechanismType = 0x00000021

	MechanismType_X9_42_DH_KEY_PAIR_GEN  MechanismType = 0x00000030
	MechanismType_X9_42_DH_DERIVE        MechanismType = 0x00000031
	MechanismType_X9_42_DH_HYBRID_DERIVE MechanismType = 0x00000032
	MechanismType_X9_42_MQV_DERIVE       MechanismType = 0x00000033

	MechanismType_SHA256_RSA_PKCS     MechanismType = 0x00000040
	MechanismType_SHA384_RSA_PKCS     MechanismType = 0x00000041
	MechanismType_SHA512_RSA_PKCS     MechanismType = 0x00000042
	MechanismType_SHA256_RSA_PKCS_PSS MechanismType = 0x00000043
	MechanismType_SHA384_RSA_PKCS_PSS MechanismType = 0x00000044
	MechanismType_SHA512_RSA_PKCS_PSS MechanismType = 0x00000045

	MechanismType_SHA224_RSA_PKCS     MechanismType = 0x00000046
	MechanismType_SHA224_RSA_PKCS_PSS MechanismType = 0x00000047

	MechanismType_SHA512_224                MechanismType = 0x00000048
	MechanismType_SHA512_224_HMAC           MechanismType = 0x00000049
	MechanismType_SHA512_224_HMAC_GENERAL   MechanismType = 0x0000004A
	MechanismType_SHA512_224_KEY_DERIVATION MechanismType = 0x0000004B
	MechanismType_SHA512_256                MechanismType = 0x0000004C
	MechanismType_SHA512_256_HMAC           MechanismType = 0x0000004D
	MechanismType_SHA512_256_HMAC_GENERAL   MechanismType = 0x0000004E
	MechanismType_SHA512_256_KEY_DERIVATION MechanismType = 0x0000004F

	MechanismType_SHA512_T                MechanismType = 0x00000050
	MechanismType_SHA512_T_HMAC           MechanismType = 0x00000051
	MechanismType_SHA512_T_HMAC_GENERAL   MechanismType = 0x00000052
	MechanismType_SHA512_T_KEY_DERIVATION MechanismType = 0x00000053

	MechanismType_SHA3_256_RSA_PKCS     MechanismType = 0x00000060
	MechanismType_SHA3_384_RSA_PKCS     MechanismType = 0x00000061
	MechanismType_SHA3_512_RSA_PKCS     MechanismType = 0x00000062
	MechanismType_SHA3_256_RSA_PKCS_PSS MechanismType = 0x00000063
	MechanismType_SHA3_384_RSA_PKCS_PSS MechanismType = 0x00000064
	MechanismType_SHA3_512_RSA_PKCS_PSS MechanismType = 0x00000065
	MechanismType_SHA3_224_RSA_PKCS     MechanismType = 0x00000066
	MechanismType_SHA3_224_RSA_PKCS_PSS MechanismType = 0x00000067

	MechanismType_RC2_KEY_GEN MechanismType = 0x00000100
	MechanismType_RC2_ECB     MechanismType = 0x00000101
	MechanismType_RC2_CBC     MechanismType = 0x00000102
	MechanismType_RC2_MAC     MechanismType = 0x00000103

	MechanismType_RC2_MAC_GENERAL MechanismType = 0x00000104
	MechanismType_RC2_CBC_PAD     MechanismType = 0x00000105

	MechanismType_RC4_KEY_GEN MechanismType = 0x00000110
	MechanismType_RC4         MechanismType = 0x00000111
	MechanismType_DES_KEY_GEN MechanismType = 0x00000120
	MechanismType_DES_ECB     MechanismType = 0x00000121
	MechanismType_DES_CBC     MechanismType = 0x00000122
	MechanismType_DES_MAC     MechanismType = 0x00000123

	MechanismType_DES_MAC_GENERAL MechanismType = 0x00000124
	MechanismType_DES_CBC_PAD     MechanismType = 0x00000125

	MechanismType_DES2_KEY_GEN MechanismType = 0x00000130
	MechanismType_DES3_KEY_GEN MechanismType = 0x00000131
	MechanismType_DES3_ECB     MechanismType = 0x00000132
	MechanismType_DES3_CBC     MechanismType = 0x00000133
	MechanismType_DES3_MAC     MechanismType = 0x00000134

	MechanismType_DES3_MAC_GENERAL  MechanismType = 0x00000135
	MechanismType_DES3_CBC_PAD      MechanismType = 0x00000136
	MechanismType_DES3_CMAC_GENERAL MechanismType = 0x00000137
	MechanismType_DES3_CMAC         MechanismType = 0x00000138
	MechanismType_CDMF_KEY_GEN      MechanismType = 0x00000140
	MechanismType_CDMF_ECB          MechanismType = 0x00000141
	MechanismType_CDMF_CBC          MechanismType = 0x00000142
	MechanismType_CDMF_MAC          MechanismType = 0x00000143
	MechanismType_CDMF_MAC_GENERAL  MechanismType = 0x00000144
	MechanismType_CDMF_CBC_PAD      MechanismType = 0x00000145

	MechanismType_DES_OFB64 MechanismType = 0x00000150
	MechanismType_DES_OFB8  MechanismType = 0x00000151
	MechanismType_DES_CFB64 MechanismType = 0x00000152
	MechanismType_DES_CFB8  MechanismType = 0x00000153

	MechanismType_MD2 MechanismType = 0x00000200

	MechanismType_MD2_HMAC         MechanismType = 0x00000201
	MechanismType_MD2_HMAC_GENERAL MechanismType = 0x00000202

	MechanismType_MD5 MechanismType = 0x00000210

	MechanismType_MD5_HMAC         MechanismType = 0x00000211
	MechanismType_MD5_HMAC_GENERAL MechanismType = 0x00000212

	MechanismType_SHA_1 MechanismType = 0x00000220

	MechanismType_SHA_1_HMAC         MechanismType = 0x00000221
	MechanismType_SHA_1_HMAC_GENERAL MechanismType = 0x00000222

	MechanismType_RIPEMD128              MechanismType = 0x00000230
	MechanismType_RIPEMD128_HMAC         MechanismType = 0x00000231
	MechanismType_RIPEMD128_HMAC_GENERAL MechanismType = 0x00000232
	MechanismType_RIPEMD160              MechanismType = 0x00000240
	MechanismType_RIPEMD160_HMAC         MechanismType = 0x00000241
	MechanismType_RIPEMD160_HMAC_GENERAL MechanismType = 0x00000242

	MechanismType_SHA256              MechanismType = 0x00000250
	MechanismType_SHA256_HMAC         MechanismType = 0x00000251
	MechanismType_SHA256_HMAC_GENERAL MechanismType = 0x00000252
	MechanismType_SHA224              MechanismType = 0x00000255
	MechanismType_SHA224_HMAC         MechanismType = 0x00000256
	MechanismType_SHA224_HMAC_GENERAL MechanismType = 0x00000257
	MechanismType_SHA384              MechanismType = 0x00000260
	MechanismType_SHA384_HMAC         MechanismType = 0x00000261
	MechanismType_SHA384_HMAC_GENERAL MechanismType = 0x00000262
	MechanismType_SHA512              MechanismType = 0x00000270
	MechanismType_SHA512_HMAC         MechanismType = 0x00000271
	MechanismType_SHA512_HMAC_GENERAL MechanismType = 0x00000272
	MechanismType_SECURID_KEY_GEN     MechanismType = 0x00000280
	MechanismType_SECURID             MechanismType = 0x00000282
	MechanismType_HOTP_KEY_GEN        MechanismType = 0x00000290
	MechanismType_HOTP                MechanismType = 0x00000291
	MechanismType_ACTI                MechanismType = 0x000002A0
	MechanismType_ACTI_KEY_GEN        MechanismType = 0x000002A1

	MechanismType_SHA3_256              MechanismType = 0x000002B0
	MechanismType_SHA3_256_HMAC         MechanismType = 0x000002B1
	MechanismType_SHA3_256_HMAC_GENERAL MechanismType = 0x000002B2
	MechanismType_SHA3_256_KEY_GEN      MechanismType = 0x000002B3
	MechanismType_SHA3_224              MechanismType = 0x000002B5
	MechanismType_SHA3_224_HMAC         MechanismType = 0x000002B6
	MechanismType_SHA3_224_HMAC_GENERAL MechanismType = 0x000002B7
	MechanismType_SHA3_224_KEY_GEN      MechanismType = 0x000002B8
	MechanismType_SHA3_384              MechanismType = 0x000002C0
	MechanismType_SHA3_384_HMAC         MechanismType = 0x000002C1
	MechanismType_SHA3_384_HMAC_GENERAL MechanismType = 0x000002C2
	MechanismType_SHA3_384_KEY_GEN      MechanismType = 0x000002C3
	MechanismType_SHA3_512              MechanismType = 0x000002D0
	MechanismType_SHA3_512_HMAC         MechanismType = 0x000002D1
	MechanismType_SHA3_512_HMAC_GENERAL MechanismType = 0x000002D2
	MechanismType_SHA3_512_KEY_GEN      MechanismType = 0x000002D3

	MechanismType_CAST_KEY_GEN      MechanismType = 0x00000300
	MechanismType_CAST_ECB          MechanismType = 0x00000301
	MechanismType_CAST_CBC          MechanismType = 0x00000302
	MechanismType_CAST_MAC          MechanismType = 0x00000303
	MechanismType_CAST_MAC_GENERAL  MechanismType = 0x00000304
	MechanismType_CAST_CBC_PAD      MechanismType = 0x00000305
	MechanismType_CAST3_KEY_GEN     MechanismType = 0x00000310
	MechanismType_CAST3_ECB         MechanismType = 0x00000311
	MechanismType_CAST3_CBC         MechanismType = 0x00000312
	MechanismType_CAST3_MAC         MechanismType = 0x00000313
	MechanismType_CAST3_MAC_GENERAL MechanismType = 0x00000314
	MechanismType_CAST3_CBC_PAD     MechanismType = 0x00000315
	/* Note that CAST128 and CAST5 are the same algorithm */
	MechanismType_CAST5_KEY_GEN             MechanismType = 0x00000320
	MechanismType_CAST128_KEY_GEN           MechanismType = 0x00000320
	MechanismType_CAST5_ECB                 MechanismType = 0x00000321
	MechanismType_CAST128_ECB               MechanismType = 0x00000321
	MechanismType_CAST5_CBC                 MechanismType = 0x00000322 /* Deprecated */
	MechanismType_CAST128_CBC               MechanismType = 0x00000322
	MechanismType_CAST5_MAC                 MechanismType = 0x00000323 /* Deprecated */
	MechanismType_CAST128_MAC               MechanismType = 0x00000323
	MechanismType_CAST5_MAC_GENERAL         MechanismType = 0x00000324 /* Deprecated */
	MechanismType_CAST128_MAC_GENERAL       MechanismType = 0x00000324
	MechanismType_CAST5_CBC_PAD             MechanismType = 0x00000325 /* Deprecated */
	MechanismType_CAST128_CBC_PAD           MechanismType = 0x00000325
	MechanismType_RC5_KEY_GEN               MechanismType = 0x00000330
	MechanismType_RC5_ECB                   MechanismType = 0x00000331
	MechanismType_RC5_CBC                   MechanismType = 0x00000332
	MechanismType_RC5_MAC                   MechanismType = 0x00000333
	MechanismType_RC5_MAC_GENERAL           MechanismType = 0x00000334
	MechanismType_RC5_CBC_PAD               MechanismType = 0x00000335
	MechanismType_IDEA_KEY_GEN              MechanismType = 0x00000340
	MechanismType_IDEA_ECB                  MechanismType = 0x00000341
	MechanismType_IDEA_CBC                  MechanismType = 0x00000342
	MechanismType_IDEA_MAC                  MechanismType = 0x00000343
	MechanismType_IDEA_MAC_GENERAL          MechanismType = 0x00000344
	MechanismType_IDEA_CBC_PAD              MechanismType = 0x00000345
	MechanismType_GENERIC_SECRET_KEY_GEN    MechanismType = 0x00000350
	MechanismType_CONCATENATE_BASE_AND_KEY  MechanismType = 0x00000360
	MechanismType_CONCATENATE_BASE_AND_DATA MechanismType = 0x00000362
	MechanismType_CONCATENATE_DATA_AND_BASE MechanismType = 0x00000363
	MechanismType_XOR_BASE_AND_DATA         MechanismType = 0x00000364
	MechanismType_EXTRACT_KEY_FROM_KEY      MechanismType = 0x00000365
	MechanismType_SSL3_PRE_MASTER_KEY_GEN   MechanismType = 0x00000370
	MechanismType_SSL3_MASTER_KEY_DERIVE    MechanismType = 0x00000371
	MechanismType_SSL3_KEY_AND_MAC_DERIVE   MechanismType = 0x00000372

	MechanismType_SSL3_MASTER_KEY_DERIVE_DH MechanismType = 0x00000373
	MechanismType_TLS_PRE_MASTER_KEY_GEN    MechanismType = 0x00000374
	MechanismType_TLS_MASTER_KEY_DERIVE     MechanismType = 0x00000375
	MechanismType_TLS_KEY_AND_MAC_DERIVE    MechanismType = 0x00000376
	MechanismType_TLS_MASTER_KEY_DERIVE_DH  MechanismType = 0x00000377

	MechanismType_TLS_PRF MechanismType = 0x00000378

	MechanismType_SSL3_MD5_MAC        MechanismType = 0x00000380
	MechanismType_SSL3_SHA1_MAC       MechanismType = 0x00000381
	MechanismType_MD5_KEY_DERIVATION  MechanismType = 0x00000390
	MechanismType_MD2_KEY_DERIVATION  MechanismType = 0x00000391
	MechanismType_SHA1_KEY_DERIVATION MechanismType = 0x00000392

	MechanismType_SHA256_KEY_DERIVATION    MechanismType = 0x00000393
	MechanismType_SHA384_KEY_DERIVATION    MechanismType = 0x00000394
	MechanismType_SHA512_KEY_DERIVATION    MechanismType = 0x00000395
	MechanismType_SHA224_KEY_DERIVATION    MechanismType = 0x00000396
	MechanismType_SHA3_256_KEY_DERIVATION  MechanismType = 0x00000397
	MechanismType_SHA3_224_KEY_DERIVATION  MechanismType = 0x00000398
	MechanismType_SHA3_384_KEY_DERIVATION  MechanismType = 0x00000399
	MechanismType_SHA3_512_KEY_DERIVATION  MechanismType = 0x0000039A
	MechanismType_SHAKE_128_KEY_DERIVATION MechanismType = 0x0000039B
	MechanismType_SHAKE_256_KEY_DERIVATION MechanismType = 0x0000039C
	MechanismType_SHA3_256_KEY_DERIVE      MechanismType = MechanismType_SHA3_256_KEY_DERIVATION
	MechanismType_SHA3_224_KEY_DERIVE      MechanismType = MechanismType_SHA3_224_KEY_DERIVATION
	MechanismType_SHA3_384_KEY_DERIVE      MechanismType = MechanismType_SHA3_384_KEY_DERIVATION
	MechanismType_SHA3_512_KEY_DERIVE      MechanismType = MechanismType_SHA3_512_KEY_DERIVATION
	MechanismType_SHAKE_128_KEY_DERIVE     MechanismType = MechanismType_SHAKE_128_KEY_DERIVATION
	MechanismType_SHAKE_256_KEY_DERIVE     MechanismType = MechanismType_SHAKE_256_KEY_DERIVATION

	MechanismType_PBE_MD2_DES_CBC       MechanismType = 0x000003A0
	MechanismType_PBE_MD5_DES_CBC       MechanismType = 0x000003A1
	MechanismType_PBE_MD5_CAST_CBC      MechanismType = 0x000003A2
	MechanismType_PBE_MD5_CAST3_CBC     MechanismType = 0x000003A3
	MechanismType_PBE_MD5_CAST5_CBC     MechanismType = 0x000003A4 /* Deprecated */
	MechanismType_PBE_MD5_CAST128_CBC   MechanismType = 0x000003A4
	MechanismType_PBE_SHA1_CAST5_CBC    MechanismType = 0x000003A5 /* Deprecated */
	MechanismType_PBE_SHA1_CAST128_CBC  MechanismType = 0x000003A5
	MechanismType_PBE_SHA1_RC4_128      MechanismType = 0x000003A6
	MechanismType_PBE_SHA1_RC4_40       MechanismType = 0x000003A7
	MechanismType_PBE_SHA1_DES3_EDE_CBC MechanismType = 0x000003A8
	MechanismType_PBE_SHA1_DES2_EDE_CBC MechanismType = 0x000003A9
	MechanismType_PBE_SHA1_RC2_128_CBC  MechanismType = 0x000003AA
	MechanismType_PBE_SHA1_RC2_40_CBC   MechanismType = 0x000003AB

	MechanismType_PKCS5_PBKD2 MechanismType = 0x000003B0

	MechanismType_PBA_SHA1_WITH_SHA1_HMAC MechanismType = 0x000003C0

	MechanismType_WTLS_PRE_MASTER_KEY_GEN        MechanismType = 0x000003D0
	MechanismType_WTLS_MASTER_KEY_DERIVE         MechanismType = 0x000003D1
	MechanismType_WTLS_MASTER_KEY_DERIVE_DH_ECC  MechanismType = 0x000003D2
	MechanismType_WTLS_PRF                       MechanismType = 0x000003D3
	MechanismType_WTLS_SERVER_KEY_AND_MAC_DERIVE MechanismType = 0x000003D4
	MechanismType_WTLS_CLIENT_KEY_AND_MAC_DERIVE MechanismType = 0x000003D5

	MechanismType_TLS12_MAC                  MechanismType = 0x000003D8
	MechanismType_TLS12_KDF                  MechanismType = 0x000003D9
	MechanismType_TLS12_MASTER_KEY_DERIVE    MechanismType = 0x000003E0
	MechanismType_TLS12_KEY_AND_MAC_DERIVE   MechanismType = 0x000003E1
	MechanismType_TLS12_MASTER_KEY_DERIVE_DH MechanismType = 0x000003E2
	MechanismType_TLS12_KEY_SAFE_DERIVE      MechanismType = 0x000003E3
	MechanismType_TLS_MAC                    MechanismType = 0x000003E4
	MechanismType_TLS_KDF                    MechanismType = 0x000003E5

	MechanismType_KEY_WRAP_LYNKS    MechanismType = 0x00000400
	MechanismType_KEY_WRAP_SET_OAEP MechanismType = 0x00000401

	MechanismType_CMS_SIG    MechanismType = 0x00000500
	MechanismType_KIP_DERIVE MechanismType = 0x00000510
	MechanismType_KIP_WRAP   MechanismType = 0x00000511
	MechanismType_KIP_MAC    MechanismType = 0x00000512

	MechanismType_CAMELLIA_KEY_GEN          MechanismType = 0x00000550
	MechanismType_CAMELLIA_ECB              MechanismType = 0x00000551
	MechanismType_CAMELLIA_CBC              MechanismType = 0x00000552
	MechanismType_CAMELLIA_MAC              MechanismType = 0x00000553
	MechanismType_CAMELLIA_MAC_GENERAL      MechanismType = 0x00000554
	MechanismType_CAMELLIA_CBC_PAD          MechanismType = 0x00000555
	MechanismType_CAMELLIA_ECB_ENCRYPT_DATA MechanismType = 0x00000556
	MechanismType_CAMELLIA_CBC_ENCRYPT_DATA MechanismType = 0x00000557
	MechanismType_CAMELLIA_CTR              MechanismType = 0x00000558

	MechanismType_ARIA_KEY_GEN          MechanismType = 0x00000560
	MechanismType_ARIA_ECB              MechanismType = 0x00000561
	MechanismType_ARIA_CBC              MechanismType = 0x00000562
	MechanismType_ARIA_MAC              MechanismType = 0x00000563
	MechanismType_ARIA_MAC_GENERAL      MechanismType = 0x00000564
	MechanismType_ARIA_CBC_PAD          MechanismType = 0x00000565
	MechanismType_ARIA_ECB_ENCRYPT_DATA MechanismType = 0x00000566
	MechanismType_ARIA_CBC_ENCRYPT_DATA MechanismType = 0x00000567

	MechanismType_SEED_KEY_GEN          MechanismType = 0x00000650
	MechanismType_SEED_ECB              MechanismType = 0x00000651
	MechanismType_SEED_CBC              MechanismType = 0x00000652
	MechanismType_SEED_MAC              MechanismType = 0x00000653
	MechanismType_SEED_MAC_GENERAL      MechanismType = 0x00000654
	MechanismType_SEED_CBC_PAD          MechanismType = 0x00000655
	MechanismType_SEED_ECB_ENCRYPT_DATA MechanismType = 0x00000656
	MechanismType_SEED_CBC_ENCRYPT_DATA MechanismType = 0x00000657

	MechanismType_SKIPJACK_KEY_GEN      MechanismType = 0x00001000
	MechanismType_SKIPJACK_ECB64        MechanismType = 0x00001001
	MechanismType_SKIPJACK_CBC64        MechanismType = 0x00001002
	MechanismType_SKIPJACK_OFB64        MechanismType = 0x00001003
	MechanismType_SKIPJACK_CFB64        MechanismType = 0x00001004
	MechanismType_SKIPJACK_CFB32        MechanismType = 0x00001005
	MechanismType_SKIPJACK_CFB16        MechanismType = 0x00001006
	MechanismType_SKIPJACK_CFB8         MechanismType = 0x00001007
	MechanismType_SKIPJACK_WRAP         MechanismType = 0x00001008
	MechanismType_SKIPJACK_PRIVATE_WRAP MechanismType = 0x00001009
	MechanismType_SKIPJACK_RELAYX       MechanismType = 0x0000100a
	MechanismType_KEA_KEY_PAIR_GEN      MechanismType = 0x00001010
	MechanismType_KEA_KEY_DERIVE        MechanismType = 0x00001011
	MechanismType_KEA_DERIVE            MechanismType = 0x00001012
	MechanismType_FORTEZZA_TIMESTAMP    MechanismType = 0x00001020
	MechanismType_BATON_KEY_GEN         MechanismType = 0x00001030
	MechanismType_BATON_ECB128          MechanismType = 0x00001031
	MechanismType_BATON_ECB96           MechanismType = 0x00001032
	MechanismType_BATON_CBC128          MechanismType = 0x00001033
	MechanismType_BATON_COUNTER         MechanismType = 0x00001034
	MechanismType_BATON_SHUFFLE         MechanismType = 0x00001035
	MechanismType_BATON_WRAP            MechanismType = 0x00001036

	MechanismType_ECDSA_KEY_PAIR_GEN MechanismType = 0x00001040 /* Deprecated */
	MechanismType_EC_KEY_PAIR_GEN    MechanismType = 0x00001040

	MechanismType_ECDSA                        MechanismType = 0x00001041
	MechanismType_ECDSA_SHA1                   MechanismType = 0x00001042
	MechanismType_ECDSA_SHA224                 MechanismType = 0x00001043
	MechanismType_ECDSA_SHA256                 MechanismType = 0x00001044
	MechanismType_ECDSA_SHA384                 MechanismType = 0x00001045
	MechanismType_ECDSA_SHA512                 MechanismType = 0x00001046
	MechanismType_EC_KEY_PAIR_GEN_W_EXTRA_BITS MechanismType = 0x0000140B

	MechanismType_ECDH1_DERIVE          MechanismType = 0x00001050
	MechanismType_ECDH1_COFACTOR_DERIVE MechanismType = 0x00001051
	MechanismType_ECMQV_DERIVE          MechanismType = 0x00001052

	MechanismType_ECDH_AES_KEY_WRAP MechanismType = 0x00001053
	MechanismType_RSA_AES_KEY_WRAP  MechanismType = 0x00001054

	MechanismType_JUNIPER_KEY_GEN MechanismType = 0x00001060
	MechanismType_JUNIPER_ECB128  MechanismType = 0x00001061
	MechanismType_JUNIPER_CBC128  MechanismType = 0x00001062
	MechanismType_JUNIPER_COUNTER MechanismType = 0x00001063
	MechanismType_JUNIPER_SHUFFLE MechanismType = 0x00001064
	MechanismType_JUNIPER_WRAP    MechanismType = 0x00001065
	MechanismType_FASTHASH        MechanismType = 0x00001070

	MechanismType_AES_XTS          MechanismType = 0x00001071
	MechanismType_AES_XTS_KEY_GEN  MechanismType = 0x00001072
	MechanismType_AES_KEY_GEN      MechanismType = 0x00001080
	MechanismType_AES_ECB          MechanismType = 0x00001081
	MechanismType_AES_CBC          MechanismType = 0x00001082
	MechanismType_AES_MAC          MechanismType = 0x00001083
	MechanismType_AES_MAC_GENERAL  MechanismType = 0x00001084
	MechanismType_AES_CBC_PAD      MechanismType = 0x00001085
	MechanismType_AES_CTR          MechanismType = 0x00001086
	MechanismType_AES_GCM          MechanismType = 0x00001087
	MechanismType_AES_CCM          MechanismType = 0x00001088
	MechanismType_AES_CTS          MechanismType = 0x00001089
	MechanismType_AES_CMAC         MechanismType = 0x0000108A
	MechanismType_AES_CMAC_GENERAL MechanismType = 0x0000108B

	MechanismType_AES_XCBC_MAC    MechanismType = 0x0000108C
	MechanismType_AES_XCBC_MAC_96 MechanismType = 0x0000108D
	MechanismType_AES_GMAC        MechanismType = 0x0000108E

	MechanismType_BLOWFISH_KEY_GEN MechanismType = 0x00001090
	MechanismType_BLOWFISH_CBC     MechanismType = 0x00001091
	MechanismType_TWOFISH_KEY_GEN  MechanismType = 0x00001092
	MechanismType_TWOFISH_CBC      MechanismType = 0x00001093
	MechanismType_BLOWFISH_CBC_PAD MechanismType = 0x00001094
	MechanismType_TWOFISH_CBC_PAD  MechanismType = 0x00001095

	MechanismType_DES_ECB_ENCRYPT_DATA  MechanismType = 0x00001100
	MechanismType_DES_CBC_ENCRYPT_DATA  MechanismType = 0x00001101
	MechanismType_DES3_ECB_ENCRYPT_DATA MechanismType = 0x00001102
	MechanismType_DES3_CBC_ENCRYPT_DATA MechanismType = 0x00001103
	MechanismType_AES_ECB_ENCRYPT_DATA  MechanismType = 0x00001104
	MechanismType_AES_CBC_ENCRYPT_DATA  MechanismType = 0x00001105

	MechanismType_GOSTR3410_KEY_PAIR_GEN          MechanismType = 0x00001200
	MechanismType_GOSTR3410                       MechanismType = 0x00001201
	MechanismType_GOSTR3410_WITH_GOSTR3411        MechanismType = 0x00001202
	MechanismType_GOSTR3410_KEY_WRAP              MechanismType = 0x00001203
	MechanismType_GOSTR3410_DERIVE                MechanismType = 0x00001204
	MechanismType_GOSTR3411                       MechanismType = 0x00001210
	MechanismType_GOSTR3411_HMAC                  MechanismType = 0x00001211
	MechanismType_GOST28147_KEY_GEN               MechanismType = 0x00001220
	MechanismType_GOST28147_ECB                   MechanismType = 0x00001221
	MechanismType_GOST28147                       MechanismType = 0x00001222
	MechanismType_GOST28147_MAC                   MechanismType = 0x00001223
	MechanismType_GOST28147_KEY_WRAP              MechanismType = 0x00001224
	MechanismType_CHACHA20_KEY_GEN                MechanismType = 0x00001225
	MechanismType_CHACHA20                        MechanismType = 0x00001226
	MechanismType_POLY1305_KEY_GEN                MechanismType = 0x00001227
	MechanismType_POLY1305                        MechanismType = 0x00001228
	MechanismType_DSA_PARAMETER_GEN               MechanismType = 0x00002000
	MechanismType_DH_PKCS_PARAMETER_GEN           MechanismType = 0x00002001
	MechanismType_X9_42_DH_PARAMETER_GEN          MechanismType = 0x00002002
	MechanismType_DSA_PROBABILISTIC_PARAMETER_GEN MechanismType = 0x00002003
	MechanismType_DSA_PROBABLISTIC_PARAMETER_GEN  MechanismType = MechanismType_DSA_PROBABILISTIC_PARAMETER_GEN
	MechanismType_DSA_SHAWE_TAYLOR_PARAMETER_GEN  MechanismType = 0x00002004
	MechanismType_DSA_FIPS_G_GEN                  MechanismType = 0x00002005

	MechanismType_AES_OFB    MechanismType = 0x00002104
	MechanismType_AES_CFB64  MechanismType = 0x00002105
	MechanismType_AES_CFB8   MechanismType = 0x00002106
	MechanismType_AES_CFB128 MechanismType = 0x00002107

	MechanismType_AES_CFB1         MechanismType = 0x00002108
	MechanismType_AES_KEY_WRAP     MechanismType = 0x00002109 /* WAS: MechanismType = 0x00001090 */
	MechanismType_AES_KEY_WRAP_PAD MechanismType = 0x0000210A /* WAS: MechanismType = 0x00001091 */
	MechanismType_AES_KEY_WRAP_KWP MechanismType = 0x0000210B

	MechanismType_RSA_PKCS_TPM_1_1      MechanismType = 0x00004001
	MechanismType_RSA_PKCS_OAEP_TPM_1_1 MechanismType = 0x00004002

	MechanismType_SHA_1_KEY_GEN            MechanismType = 0x00004003
	MechanismType_SHA224_KEY_GEN           MechanismType = 0x00004004
	MechanismType_SHA256_KEY_GEN           MechanismType = 0x00004005
	MechanismType_SHA384_KEY_GEN           MechanismType = 0x00004006
	MechanismType_SHA512_KEY_GEN           MechanismType = 0x00004007
	MechanismType_SHA512_224_KEY_GEN       MechanismType = 0x00004008
	MechanismType_SHA512_256_KEY_GEN       MechanismType = 0x00004009
	MechanismType_SHA512_T_KEY_GEN         MechanismType = 0x0000400a
	MechanismType_NULL                     MechanismType = 0x0000400b
	MechanismType_BLAKE2B_160              MechanismType = 0x0000400c
	MechanismType_BLAKE2B_160_HMAC         MechanismType = 0x0000400d
	MechanismType_BLAKE2B_160_HMAC_GENERAL MechanismType = 0x0000400e
	MechanismType_BLAKE2B_160_KEY_DERIVE   MechanismType = 0x0000400f
	MechanismType_BLAKE2B_160_KEY_GEN      MechanismType = 0x00004010
	MechanismType_BLAKE2B_256              MechanismType = 0x00004011
	MechanismType_BLAKE2B_256_HMAC         MechanismType = 0x00004012
	MechanismType_BLAKE2B_256_HMAC_GENERAL MechanismType = 0x00004013
	MechanismType_BLAKE2B_256_KEY_DERIVE   MechanismType = 0x00004014
	MechanismType_BLAKE2B_256_KEY_GEN      MechanismType = 0x00004015
	MechanismType_BLAKE2B_384              MechanismType = 0x00004016
	MechanismType_BLAKE2B_384_HMAC         MechanismType = 0x00004017
	MechanismType_BLAKE2B_384_HMAC_GENERAL MechanismType = 0x00004018
	MechanismType_BLAKE2B_384_KEY_DERIVE   MechanismType = 0x00004019
	MechanismType_BLAKE2B_384_KEY_GEN      MechanismType = 0x0000401a
	MechanismType_BLAKE2B_512              MechanismType = 0x0000401b
	MechanismType_BLAKE2B_512_HMAC         MechanismType = 0x0000401c
	MechanismType_BLAKE2B_512_HMAC_GENERAL MechanismType = 0x0000401d
	MechanismType_BLAKE2B_512_KEY_DERIVE   MechanismType = 0x0000401e
	MechanismType_BLAKE2B_512_KEY_GEN      MechanismType = 0x0000401f
	MechanismType_SALSA20                  MechanismType = 0x00004020
	MechanismType_CHACHA20_POLY1305        MechanismType = 0x00004021
	MechanismType_SALSA20_POLY1305         MechanismType = 0x00004022
	MechanismType_X3DH_INITIALIZE          MechanismType = 0x00004023
	MechanismType_X3DH_RESPOND             MechanismType = 0x00004024
	MechanismType_X2RATCHET_INITIALIZE     MechanismType = 0x00004025
	MechanismType_X2RATCHET_RESPOND        MechanismType = 0x00004026
	MechanismType_X2RATCHET_ENCRYPT        MechanismType = 0x00004027
	MechanismType_X2RATCHET_DECRYPT        MechanismType = 0x00004028
	MechanismType_XEDDSA                   MechanismType = 0x00004029
	MechanismType_HKDF_DERIVE              MechanismType = 0x0000402a
	MechanismType_HKDF_DATA                MechanismType = 0x0000402b
	MechanismType_HKDF_KEY_GEN             MechanismType = 0x0000402c
	MechanismType_SALSA20_KEY_GEN          MechanismType = 0x0000402d

	MechanismType_ECDSA_SHA3_224                MechanismType = 0x00001047
	MechanismType_ECDSA_SHA3_256                MechanismType = 0x00001048
	MechanismType_ECDSA_SHA3_384                MechanismType = 0x00001049
	MechanismType_ECDSA_SHA3_512                MechanismType = 0x0000104a
	MechanismType_EC_EDWARDS_KEY_PAIR_GEN       MechanismType = 0x00001055
	MechanismType_EC_MONTGOMERY_KEY_PAIR_GEN    MechanismType = 0x00001056
	MechanismType_EDDSA                         MechanismType = 0x00001057
	MechanismType_SP800_108_COUNTER_KDF         MechanismType = 0x000003ac
	MechanismType_SP800_108_FEEDBACK_KDF        MechanismType = 0x000003ad
	MechanismType_SP800_108_DOUBLE_PIPELINE_KDF MechanismType = 0x000003ae

	MechanismType_VENDOR_DEFINED MechanismType = 0x80000000
)

type Mechanism struct {
	Mechanism MechanismType
	Parameter MechanismParameter
}

type MechanismParameter struct {
	SymmetricKey []byte
	DigestLength int
}
