package messages

import (
	"encoding/asn1"
	"fmt"
)

func Unmarshal(bytes []byte) (Message, error) {
	type wrapper struct {
		Type Type
	}

	var wrapped wrapper
	_, err := asn1.Unmarshal(bytes, &wrapped)
	if err != nil {
		return nil, err
	}

	switch wrapped.Type {
	case TypeInitializeRequest:
		var decoded WrappedInitializeRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeInitializeResponse:
		var decoded WrappedInitializeResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeFinalizeRequest:
		var decoded WrappedFinalizeRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeFinalizeResponse:
		var decoded WrappedFinalizeResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetInfoRequest:
		var decoded WrappedGetInfoRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetInfoResponse:
		var decoded WrappedGetInfoResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetFunctionListRequest:
		var decoded WrappedGetFunctionListRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetFunctionListResponse:
		var decoded WrappedGetFunctionListResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetInterfaceListRequest:
		var decoded WrappedGetInterfaceListRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetInterfaceListResponse:
		var decoded WrappedGetInterfaceListResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetInterfaceRequest:
		var decoded WrappedGetInterfaceRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetInterfaceResponse:
		var decoded WrappedGetInterfaceResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetSlotListRequest:
		var decoded WrappedGetSlotListRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetSlotListResponse:
		var decoded WrappedGetSlotListResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetSlotInfoRequest:
		var decoded WrappedGetSlotInfoRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetSlotInfoResponse:
		var decoded WrappedGetSlotInfoResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetTokenInfoRequest:
		var decoded WrappedGetTokenInfoRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetTokenInfoResponse:
		var decoded WrappedGetTokenInfoResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeWaitForSlotEventRequest:
		var decoded WrappedWaitForSlotEventRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeWaitForSlotEventResponse:
		var decoded WrappedWaitForSlotEventResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetMechanismListRequest:
		var decoded WrappedGetMechanismListRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetMechanismListResponse:
		var decoded WrappedGetMechanismListResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetMechanismInfoRequest:
		var decoded WrappedGetMechanismInfoRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetMechanismInfoResponse:
		var decoded WrappedGetMechanismInfoResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeInitTokenRequest:
		var decoded WrappedInitTokenRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeInitTokenResponse:
		var decoded WrappedInitTokenResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeInitPINRequest:
		var decoded WrappedInitPINRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeInitPINResponse:
		var decoded WrappedInitPINResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSetPINRequest:
		var decoded WrappedSetPINRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSetPINResponse:
		var decoded WrappedSetPINResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeOpenSessionRequest:
		var decoded WrappedOpenSessionRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeOpenSessionResponse:
		var decoded WrappedOpenSessionResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCloseSessionRequest:
		var decoded WrappedCloseSessionRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCloseSessionResponse:
		var decoded WrappedCloseSessionResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCloseAllSessionsRequest:
		var decoded WrappedCloseAllSessionsRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCloseAllSessionsResponse:
		var decoded WrappedCloseAllSessionsResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetSessionInfoRequest:
		var decoded WrappedGetSessionInfoRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetSessionInfoResponse:
		var decoded WrappedGetSessionInfoResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSessionCancelRequest:
		var decoded WrappedSessionCancelRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSessionCancelResponse:
		var decoded WrappedSessionCancelResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetOperationStateRequest:
		var decoded WrappedGetOperationStateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetOperationStateResponse:
		var decoded WrappedGetOperationStateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSetOperationStateRequest:
		var decoded WrappedSetOperationStateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSetOperationStateResponse:
		var decoded WrappedSetOperationStateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeLoginRequest:
		var decoded WrappedLoginRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeLoginResponse:
		var decoded WrappedLoginResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeLoginUserRequest:
		var decoded WrappedLoginUserRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeLoginUserResponse:
		var decoded WrappedLoginUserResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeLogoutRequest:
		var decoded WrappedLogoutRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeLogoutResponse:
		var decoded WrappedLogoutResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCreateObjectRequest:
		var decoded WrappedCreateObjectRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCreateObjectResponse:
		var decoded WrappedCreateObjectResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCopyObjectRequest:
		var decoded WrappedCopyObjectRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCopyObjectResponse:
		var decoded WrappedCopyObjectResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDestroyObjectRequest:
		var decoded WrappedDestroyObjectRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDestroyObjectResponse:
		var decoded WrappedDestroyObjectResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetObjectSizeRequest:
		var decoded WrappedGetObjectSizeRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetObjectSizeResponse:
		var decoded WrappedGetObjectSizeResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetAttributeValueRequest:
		var decoded WrappedGetAttributeValueRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetAttributeValueResponse:
		var decoded WrappedGetAttributeValueResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSetAttributeValueRequest:
		var decoded WrappedSetAttributeValueRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSetAttributeValueResponse:
		var decoded WrappedSetAttributeValueResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeFindObjectsInitRequest:
		var decoded WrappedFindObjectsInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeFindObjectsInitResponse:
		var decoded WrappedFindObjectsInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeFindObjectsRequest:
		var decoded WrappedFindObjectsRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeFindObjectsResponse:
		var decoded WrappedFindObjectsResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeFindObjectsFinalRequest:
		var decoded WrappedFindObjectsFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeFindObjectsFinalResponse:
		var decoded WrappedFindObjectsFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptInitRequest:
		var decoded WrappedEncryptInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptInitResponse:
		var decoded WrappedEncryptInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptRequest:
		var decoded WrappedEncryptRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptResponse:
		var decoded WrappedEncryptResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptUpdateRequest:
		var decoded WrappedEncryptUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptUpdateResponse:
		var decoded WrappedEncryptUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptFinalRequest:
		var decoded WrappedEncryptFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptFinalResponse:
		var decoded WrappedEncryptFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageInitRequest:
		var decoded WrappedEncryptMessageInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageInitResponse:
		var decoded WrappedEncryptMessageInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageRequest:
		var decoded WrappedEncryptMessageRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageResponse:
		var decoded WrappedEncryptMessageResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageBeginRequest:
		var decoded WrappedEncryptMessageBeginRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageBeginResponse:
		var decoded WrappedEncryptMessageBeginResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageNextRequest:
		var decoded WrappedEncryptMessageNextRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageNextResponse:
		var decoded WrappedEncryptMessageNextResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageFinalRequest:
		var decoded WrappedEncryptMessageFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeEncryptMessageFinalResponse:
		var decoded WrappedEncryptMessageFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptInitRequest:
		var decoded WrappedDecryptInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptInitResponse:
		var decoded WrappedDecryptInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptRequest:
		var decoded WrappedDecryptRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptResponse:
		var decoded WrappedDecryptResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptUpdateRequest:
		var decoded WrappedDecryptUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptUpdateResponse:
		var decoded WrappedDecryptUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptFinalRequest:
		var decoded WrappedDecryptFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptFinalResponse:
		var decoded WrappedDecryptFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageInitRequest:
		var decoded WrappedDecryptMessageInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageInitResponse:
		var decoded WrappedDecryptMessageInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageRequest:
		var decoded WrappedDecryptMessageRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageResponse:
		var decoded WrappedDecryptMessageResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageBeginRequest:
		var decoded WrappedDecryptMessageBeginRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageBeginResponse:
		var decoded WrappedDecryptMessageBeginResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageNextRequest:
		var decoded WrappedDecryptMessageNextRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageNextResponse:
		var decoded WrappedDecryptMessageNextResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageFinalRequest:
		var decoded WrappedDecryptMessageFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptMessageFinalResponse:
		var decoded WrappedDecryptMessageFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestInitRequest:
		var decoded WrappedDigestInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestInitResponse:
		var decoded WrappedDigestInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestRequest:
		var decoded WrappedDigestRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestResponse:
		var decoded WrappedDigestResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestUpdateRequest:
		var decoded WrappedDigestUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestUpdateResponse:
		var decoded WrappedDigestUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestKeyRequest:
		var decoded WrappedDigestKeyRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestKeyResponse:
		var decoded WrappedDigestKeyResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestFinalRequest:
		var decoded WrappedDigestFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeDigestFinalResponse:
		var decoded WrappedDigestFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil

	case TypeSignInitRequest:
		var decoded WrappedSignInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignInitResponse:
		var decoded WrappedSignInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignRequest:
		var decoded WrappedSignRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignResponse:
		var decoded WrappedSignResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignUpdateRequest:
		var decoded WrappedSignUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignUpdateResponse:
		var decoded WrappedSignUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignFinalRequest:
		var decoded WrappedSignFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignFinalResponse:
		var decoded WrappedSignFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignRecoverInitRequest:
		var decoded WrappedSignRecoverInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignRecoverInitResponse:
		var decoded WrappedSignRecoverInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignRecoverRequest:
		var decoded WrappedSignRecoverRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignRecoverResponse:
		var decoded WrappedSignRecoverResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageInitRequest:
		var decoded WrappedSignMessageInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageInitResponse:
		var decoded WrappedSignMessageInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageRequest:
		var decoded WrappedSignMessageRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageResponse:
		var decoded WrappedSignMessageResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageBeginRequest:
		var decoded WrappedSignMessageBeginRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageBeginResponse:
		var decoded WrappedSignMessageBeginResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageNextRequest:
		var decoded WrappedSignMessageNextRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageNextResponse:
		var decoded WrappedSignMessageNextResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageFinalRequest:
		var decoded WrappedSignMessageFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignMessageFinalResponse:
		var decoded WrappedSignMessageFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyInitRequest:
		var decoded WrappedVerifyInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyInitResponse:
		var decoded WrappedVerifyInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyRequest:
		var decoded WrappedVerifyRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyResponse:
		var decoded WrappedVerifyResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyUpdateRequest:
		var decoded WrappedVerifyUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyUpdateResponse:
		var decoded WrappedVerifyUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyFinalRequest:
		var decoded WrappedVerifyFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyFinalResponse:
		var decoded WrappedVerifyFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyRecoverInitRequest:
		var decoded WrappedVerifyRecoverInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyRecoverInitResponse:
		var decoded WrappedVerifyRecoverInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyRecoverRequest:
		var decoded WrappedVerifyRecoverRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyRecoverResponse:
		var decoded WrappedVerifyRecoverResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageInitRequest:
		var decoded WrappedVerifyMessageInitRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageInitResponse:
		var decoded WrappedVerifyMessageInitResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageRequest:
		var decoded WrappedVerifyMessageRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageResponse:
		var decoded WrappedVerifyMessageResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageBeginRequest:
		var decoded WrappedVerifyMessageBeginRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageBeginResponse:
		var decoded WrappedVerifyMessageBeginResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageNextRequest:
		var decoded WrappedVerifyMessageNextRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageNextResponse:
		var decoded WrappedVerifyMessageNextResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageFinalRequest:
		var decoded WrappedVerifyMessageFinalRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeVerifyMessageFinalResponse:
		var decoded WrappedVerifyMessageFinalResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDigestEncryptUpdateRequest:
		var decoded WrappedDigestEncryptUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDigestEncryptUpdateResponse:
		var decoded WrappedDigestEncryptUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptDigestUpdateRequest:
		var decoded WrappedDecryptDigestUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptDigestUpdateResponse:
		var decoded WrappedDecryptDigestUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignEncryptUpdateRequest:
		var decoded WrappedSignEncryptUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSignEncryptUpdateResponse:
		var decoded WrappedSignEncryptUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptVerifyUpdateRequest:
		var decoded WrappedDecryptVerifyUpdateRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDecryptVerifyUpdateResponse:
		var decoded WrappedDecryptVerifyUpdateResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGenerateKeyRequest:
		var decoded WrappedGenerateKeyRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGenerateKeyResponse:
		var decoded WrappedGenerateKeyResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGenerateKeyPairRequest:
		var decoded WrappedGenerateKeyPairRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGenerateKeyPairResponse:
		var decoded WrappedGenerateKeyPairResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeWrapKeyRequest:
		var decoded WrappedWrapKeyRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeWrapKeyResponse:
		var decoded WrappedWrapKeyResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeUnwrapKeyRequest:
		var decoded WrappedUnwrapKeyRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeUnwrapKeyResponse:
		var decoded WrappedUnwrapKeyResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDeriveKeyRequest:
		var decoded WrappedDeriveKeyRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeDeriveKeyResponse:
		var decoded WrappedDeriveKeyResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSeedRandomRequest:
		var decoded WrappedSeedRandomRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeSeedRandomResponse:
		var decoded WrappedSeedRandomResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGenerateRandomRequest:
		var decoded WrappedGenerateRandomRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGenerateRandomResponse:
		var decoded WrappedGenerateRandomResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetFunctionStatusRequest:
		var decoded WrappedGetFunctionStatusRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeGetFunctionStatusResponse:
		var decoded WrappedGetFunctionStatusResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCancelFunctionRequest:
		var decoded WrappedCancelFunctionRequest
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	case TypeCancelFunctionResponse:
		var decoded WrappedCancelFunctionResponse
		if _, err = asn1.Unmarshal(bytes, &decoded); err != nil {
			return nil, err
		}
		return &decoded.Message, nil
	default:
		return nil, fmt.Errorf("unknown message type: %v", wrapped.Type)
	}
}
