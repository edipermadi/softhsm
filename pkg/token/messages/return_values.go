package messages

type ReturnValue int

const (
	ReturnValue_OK              ReturnValue = 0x00000000
	ReturnValue_CANCEL          ReturnValue = 0x00000001
	ReturnValue_HOST_MEMORY     ReturnValue = 0x00000002
	ReturnValue_SLOT_ID_INVALID ReturnValue = 0x00000003

	ReturnValue_GENERAL_ERROR   ReturnValue = 0x00000005
	ReturnValue_FUNCTION_FAILED ReturnValue = 0x00000006

	ReturnValue_ARGUMENTS_BAD          ReturnValue = 0x00000007
	ReturnValue_NO_EVENT               ReturnValue = 0x00000008
	ReturnValue_NEED_TO_CREATE_THREADS ReturnValue = 0x00000009
	ReturnValue_CANT_LOCK              ReturnValue = 0x0000000A

	ReturnValue_ATTRIBUTE_READ_ONLY     ReturnValue = 0x00000010
	ReturnValue_ATTRIBUTE_SENSITIVE     ReturnValue = 0x00000011
	ReturnValue_ATTRIBUTE_TYPE_INVALID  ReturnValue = 0x00000012
	ReturnValue_ATTRIBUTE_VALUE_INVALID ReturnValue = 0x00000013

	ReturnValue_ACTION_PROHIBITED ReturnValue = 0x0000001B

	ReturnValue_DATA_INVALID             ReturnValue = 0x00000020
	ReturnValue_DATA_LEN_RANGE           ReturnValue = 0x00000021
	ReturnValue_DEVICE_ERROR             ReturnValue = 0x00000030
	ReturnValue_DEVICE_MEMORY            ReturnValue = 0x00000031
	ReturnValue_DEVICE_REMOVED           ReturnValue = 0x00000032
	ReturnValue_ENCRYPTED_DATA_INVALID   ReturnValue = 0x00000040
	ReturnValue_ENCRYPTED_DATA_LEN_RANGE ReturnValue = 0x00000041
	ReturnValue_AEAD_DECRYPT_FAILED      ReturnValue = 0x00000042
	ReturnValue_FUNCTION_CANCELED        ReturnValue = 0x00000050
	ReturnValue_FUNCTION_NOT_PARALLEL    ReturnValue = 0x00000051

	ReturnValue_FUNCTION_NOT_SUPPORTED ReturnValue = 0x00000054

	ReturnValue_KEY_HANDLE_INVALID ReturnValue = 0x00000060

	ReturnValue_KEY_SIZE_RANGE        ReturnValue = 0x00000062
	ReturnValue_KEY_TYPE_INCONSISTENT ReturnValue = 0x00000063

	ReturnValue_KEY_NOT_NEEDED             ReturnValue = 0x00000064
	ReturnValue_KEY_CHANGED                ReturnValue = 0x00000065
	ReturnValue_KEY_NEEDED                 ReturnValue = 0x00000066
	ReturnValue_KEY_INDIGESTIBLE           ReturnValue = 0x00000067
	ReturnValue_KEY_FUNCTION_NOT_PERMITTED ReturnValue = 0x00000068
	ReturnValue_KEY_NOT_WRAPPABLE          ReturnValue = 0x00000069
	ReturnValue_KEY_UNEXTRACTABLE          ReturnValue = 0x0000006A

	ReturnValue_MECHANISM_INVALID       ReturnValue = 0x00000070
	ReturnValue_MECHANISM_PARAM_INVALID ReturnValue = 0x00000071

	ReturnValue_OBJECT_HANDLE_INVALID     ReturnValue = 0x00000082
	ReturnValue_OPERATION_ACTIVE          ReturnValue = 0x00000090
	ReturnValue_OPERATION_NOT_INITIALIZED ReturnValue = 0x00000091
	ReturnValue_PIN_INCORRECT             ReturnValue = 0x000000A0
	ReturnValue_PIN_INVALID               ReturnValue = 0x000000A1
	ReturnValue_PIN_LEN_RANGE             ReturnValue = 0x000000A2

	ReturnValue_PIN_EXPIRED ReturnValue = 0x000000A3
	ReturnValue_PIN_LOCKED  ReturnValue = 0x000000A4

	ReturnValue_SESSION_CLOSED                 ReturnValue = 0x000000B0
	ReturnValue_SESSION_COUNT                  ReturnValue = 0x000000B1
	ReturnValue_SESSION_HANDLE_INVALID         ReturnValue = 0x000000B3
	ReturnValue_SESSION_PARALLEL_NOT_SUPPORTED ReturnValue = 0x000000B4
	ReturnValue_SESSION_READ_ONLY              ReturnValue = 0x000000B5
	ReturnValue_SESSION_EXISTS                 ReturnValue = 0x000000B6

	ReturnValue_SESSION_READ_ONLY_EXISTS     ReturnValue = 0x000000B7
	ReturnValue_SESSION_READ_WRITE_SO_EXISTS ReturnValue = 0x000000B8

	ReturnValue_SIGNATURE_INVALID                ReturnValue = 0x000000C0
	ReturnValue_SIGNATURE_LEN_RANGE              ReturnValue = 0x000000C1
	ReturnValue_TEMPLATE_INCOMPLETE              ReturnValue = 0x000000D0
	ReturnValue_TEMPLATE_INCONSISTENT            ReturnValue = 0x000000D1
	ReturnValue_TOKEN_NOT_PRESENT                ReturnValue = 0x000000E0
	ReturnValue_TOKEN_NOT_RECOGNIZED             ReturnValue = 0x000000E1
	ReturnValue_TOKEN_WRITE_PROTECTED            ReturnValue = 0x000000E2
	ReturnValue_UNWRAPPING_KEY_HANDLE_INVALID    ReturnValue = 0x000000F0
	ReturnValue_UNWRAPPING_KEY_SIZE_RANGE        ReturnValue = 0x000000F1
	ReturnValue_UNWRAPPING_KEY_TYPE_INCONSISTENT ReturnValue = 0x000000F2
	ReturnValue_USER_ALREADY_LOGGED_IN           ReturnValue = 0x00000100
	ReturnValue_USER_NOT_LOGGED_IN               ReturnValue = 0x00000101
	ReturnValue_USER_PIN_NOT_INITIALIZED         ReturnValue = 0x00000102
	ReturnValue_USER_TYPE_INVALID                ReturnValue = 0x00000103

	ReturnValue_USER_ANOTHER_ALREADY_LOGGED_IN ReturnValue = 0x00000104
	ReturnValue_USER_TOO_MANY_TYPES            ReturnValue = 0x00000105

	ReturnValue_WRAPPED_KEY_INVALID            ReturnValue = 0x00000110
	ReturnValue_WRAPPED_KEY_LEN_RANGE          ReturnValue = 0x00000112
	ReturnValue_WRAPPING_KEY_HANDLE_INVALID    ReturnValue = 0x00000113
	ReturnValue_WRAPPING_KEY_SIZE_RANGE        ReturnValue = 0x00000114
	ReturnValue_WRAPPING_KEY_TYPE_INCONSISTENT ReturnValue = 0x00000115
	ReturnValue_RANDOM_SEED_NOT_SUPPORTED      ReturnValue = 0x00000120

	ReturnValue_RANDOM_NO_RNG ReturnValue = 0x00000121

	ReturnValue_DOMAIN_PARAMS_INVALID ReturnValue = 0x00000130

	ReturnValue_CURVE_NOT_SUPPORTED ReturnValue = 0x00000140

	ReturnValue_BUFFER_TOO_SMALL      ReturnValue = 0x00000150
	ReturnValue_SAVED_STATE_INVALID   ReturnValue = 0x00000160
	ReturnValue_INFORMATION_SENSITIVE ReturnValue = 0x00000170
	ReturnValue_STATE_UNSAVEABLE      ReturnValue = 0x00000180

	ReturnValue_CRYPTOKI_NOT_INITIALIZED     ReturnValue = 0x00000190
	ReturnValue_CRYPTOKI_ALREADY_INITIALIZED ReturnValue = 0x00000191
	ReturnValue_MUTEX_BAD                    ReturnValue = 0x000001A0
	ReturnValue_MUTEX_NOT_LOCKED             ReturnValue = 0x000001A1

	ReturnValue_NEW_PIN_MODE ReturnValue = 0x000001B0
	ReturnValue_NEXT_OTP     ReturnValue = 0x000001B1

	ReturnValue_EXCEEDED_MAX_ITERATIONS ReturnValue = 0x000001B5
	ReturnValue_FIPS_SELF_TEST_FAILED   ReturnValue = 0x000001B6
	ReturnValue_LIBRARY_LOAD_FAILED     ReturnValue = 0x000001B7
	ReturnValue_PIN_TOO_WEAK            ReturnValue = 0x000001B8
	ReturnValue_PUBLIC_KEY_INVALID      ReturnValue = 0x000001B9

	ReturnValue_FUNCTION_REJECTED       ReturnValue = 0x00000200
	ReturnValue_TOKEN_RESOURCE_EXCEEDED ReturnValue = 0x00000201
	ReturnValue_OPERATION_CANCEL_FAILED ReturnValue = 0x00000202

	ReturnValue_VENDOR_DEFINED ReturnValue = 0x80000000
)
