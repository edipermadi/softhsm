package servers

import (
	"errors"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/services"
)

type Server interface {
	Process(data []byte) ([]byte, error)
}

func NewServer(service services.Service) Server {
	return &server{service: service}
}

type server struct {
	service services.Service
}

func (s *server) Process(data []byte) ([]byte, error) {
	message, err := messages.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	var response messages.Message
	switch message.Type() {
	case messages.TypeInitializeRequest:
		response, err = s.service.Initialize(message.(*messages.InitializeRequest))
	case messages.TypeFinalizeRequest:
		response, err = s.service.Finalize(message.(*messages.FinalizeRequest))
	case messages.TypeGetInfoRequest:
		response, err = s.service.GetInfo(message.(*messages.GetInfoRequest))
	case messages.TypeGetFunctionListRequest:
		response, err = s.service.GetFunctionList(message.(*messages.GetFunctionListRequest))
	case messages.TypeGetInterfaceListRequest:
		response, err = s.service.GetInterfaceList(message.(*messages.GetInterfaceListRequest))
	case messages.TypeGetInterfaceRequest:
		response, err = s.service.GetInterface(message.(*messages.GetInterfaceRequest))

	case messages.TypeGetSlotListRequest:
		response, err = s.service.GetSlotList(message.(*messages.GetSlotListRequest))
	case messages.TypeGetSlotInfoRequest:
		response, err = s.service.GetSlotInfo(message.(*messages.GetSlotInfoRequest))
	case messages.TypeGetTokenInfoRequest:
		response, err = s.service.GetTokenInfo(message.(*messages.GetTokenInfoRequest))
	case messages.TypeWaitForSlotEventRequest:
		response, err = s.service.WaitForSlotEvent(message.(*messages.WaitForSlotEventRequest))
	case messages.TypeGetMechanismListRequest:
		response, err = s.service.GetMechanismList(message.(*messages.GetMechanismListRequest))
	case messages.TypeGetMechanismInfoRequest:
		response, err = s.service.GetMechanismInfo(message.(*messages.GetMechanismInfoRequest))
	case messages.TypeInitTokenRequest:
		response, err = s.service.InitToken(message.(*messages.InitTokenRequest))
	case messages.TypeInitPINRequest:
		response, err = s.service.InitPIN(message.(*messages.InitPINRequest))
	case messages.TypeSetPINRequest:
		response, err = s.service.SetPIN(message.(*messages.SetPINRequest))

	case messages.TypeOpenSessionRequest:
		response, err = s.service.OpenSession(message.(*messages.OpenSessionRequest))
	case messages.TypeCloseSessionRequest:
		response, err = s.service.CloseSession(message.(*messages.CloseSessionRequest))
	case messages.TypeCloseAllSessionsRequest:
		response, err = s.service.CloseAllSessions(message.(*messages.CloseAllSessionsRequest))
	case messages.TypeGetSessionInfoRequest:
		response, err = s.service.GetSessionInfo(message.(*messages.GetSessionInfoRequest))
	case messages.TypeSessionCancelRequest:
		response, err = s.service.SessionCancel(message.(*messages.SessionCancelRequest))
	case messages.TypeGetOperationStateRequest:
		response, err = s.service.GetOperationState(message.(*messages.GetOperationStateRequest))
	case messages.TypeSetOperationStateRequest:
		response, err = s.service.SetOperationState(message.(*messages.SetOperationStateRequest))
	case messages.TypeLoginRequest:
		response, err = s.service.Login(message.(*messages.LoginRequest))
	case messages.TypeLoginUserRequest:
		response, err = s.service.LoginUser(message.(*messages.LoginUserRequest))
	case messages.TypeLogoutRequest:
		response, err = s.service.Logout(message.(*messages.LogoutRequest))

	case messages.TypeCreateObjectRequest:
		response, err = s.service.CreateObject(message.(*messages.CreateObjectRequest))
	case messages.TypeCopyObjectRequest:
		response, err = s.service.CopyObject(message.(*messages.CopyObjectRequest))
	case messages.TypeDestroyObjectRequest:
		response, err = s.service.DestroyObject(message.(*messages.DestroyObjectRequest))
	case messages.TypeGetObjectSizeRequest:
		response, err = s.service.GetObjectSize(message.(*messages.GetObjectSizeRequest))
	case messages.TypeGetAttributeValueRequest:
		response, err = s.service.GetAttributeValue(message.(*messages.GetAttributeValueRequest))
	case messages.TypeSetAttributeValueRequest:
		response, err = s.service.SetAttributeValue(message.(*messages.SetAttributeValueRequest))
	case messages.TypeFindObjectsInitRequest:
		response, err = s.service.FindObjectsInit(message.(*messages.FindObjectsInitRequest))
	case messages.TypeFindObjectsRequest:
		response, err = s.service.FindObjects(message.(*messages.FindObjectsRequest))
	case messages.TypeFindObjectsFinalRequest:
		response, err = s.service.FindObjectsFinal(message.(*messages.FindObjectsFinalRequest))

	case messages.TypeEncryptInitRequest:
		response, err = s.service.EncryptInit(message.(*messages.EncryptInitRequest))
	case messages.TypeEncryptRequest:
		response, err = s.service.Encrypt(message.(*messages.EncryptRequest))
	case messages.TypeEncryptUpdateRequest:
		response, err = s.service.EncryptUpdate(message.(*messages.EncryptUpdateRequest))
	case messages.TypeEncryptFinalRequest:
		response, err = s.service.EncryptFinal(message.(*messages.EncryptFinalRequest))

	case messages.TypeEncryptMessageInitRequest:
		response, err = s.service.EncryptMessageInit(message.(*messages.EncryptMessageInitRequest))
	case messages.TypeEncryptMessageRequest:
		response, err = s.service.EncryptMessage(message.(*messages.EncryptMessageRequest))
	case messages.TypeEncryptMessageBeginRequest:
		response, err = s.service.EncryptMessageBegin(message.(*messages.EncryptMessageBeginRequest))
	case messages.TypeEncryptMessageNextRequest:
		response, err = s.service.EncryptMessageNext(message.(*messages.EncryptMessageNextRequest))
	case messages.TypeEncryptMessageFinalRequest:
		response, err = s.service.EncryptMessageFinal(message.(*messages.EncryptMessageFinalRequest))

	case messages.TypeDecryptInitRequest:
		response, err = s.service.DecryptInit(message.(*messages.DecryptInitRequest))
	case messages.TypeDecryptRequest:
		response, err = s.service.Decrypt(message.(*messages.DecryptRequest))
	case messages.TypeDecryptUpdateRequest:
		response, err = s.service.DecryptUpdate(message.(*messages.DecryptUpdateRequest))
	case messages.TypeDecryptFinalRequest:
		response, err = s.service.DecryptFinal(message.(*messages.DecryptFinalRequest))

	case messages.TypeDecryptMessageInitRequest:
		response, err = s.service.DecryptMessageInit(message.(*messages.DecryptMessageInitRequest))
	case messages.TypeDecryptMessageRequest:
		response, err = s.service.DecryptMessage(message.(*messages.DecryptMessageRequest))
	case messages.TypeDecryptMessageBeginRequest:
		response, err = s.service.DecryptMessageBegin(message.(*messages.DecryptMessageBeginRequest))
	case messages.TypeDecryptMessageNextRequest:
		response, err = s.service.DecryptMessageNext(message.(*messages.DecryptMessageNextRequest))
	case messages.TypeDecryptMessageFinalRequest:
		response, err = s.service.DecryptMessageFinal(message.(*messages.DecryptMessageFinalRequest))

	case messages.TypeDigestInitRequest:
		response, err = s.service.DigestInit(message.(*messages.DigestInitRequest))
	case messages.TypeDigestRequest:
		response, err = s.service.Digest(message.(*messages.DigestRequest))
	case messages.TypeDigestUpdateRequest:
		response, err = s.service.DigestUpdate(message.(*messages.DigestUpdateRequest))
	case messages.TypeDigestKeyRequest:
		response, err = s.service.DigestKey(message.(*messages.DigestKeyRequest))
	case messages.TypeDigestFinalRequest:
		response, err = s.service.DigestFinal(message.(*messages.DigestFinalRequest))

	case messages.TypeSignInitRequest:
		response, err = s.service.SignInit(message.(*messages.SignInitRequest))
	case messages.TypeSignRequest:
		response, err = s.service.Sign(message.(*messages.SignRequest))
	case messages.TypeSignUpdateRequest:
		response, err = s.service.SignUpdate(message.(*messages.SignUpdateRequest))
	case messages.TypeSignFinalRequest:
		response, err = s.service.SignFinal(message.(*messages.SignFinalRequest))
	case messages.TypeSignRecoverInitRequest:
		response, err = s.service.SignRecoverInit(message.(*messages.SignRecoverInitRequest))
	case messages.TypeSignRecoverRequest:
		response, err = s.service.SignRecover(message.(*messages.SignRecoverRequest))

	case messages.TypeSignMessageInitRequest:
		response, err = s.service.SignMessageInit(message.(*messages.SignMessageInitRequest))
	case messages.TypeSignMessageRequest:
		response, err = s.service.SignMessage(message.(*messages.SignMessageRequest))
	case messages.TypeSignMessageBeginRequest:
		response, err = s.service.SignMessageBegin(message.(*messages.SignMessageBeginRequest))
	case messages.TypeSignMessageNextRequest:
		response, err = s.service.SignMessageNext(message.(*messages.SignMessageNextRequest))
	case messages.TypeSignMessageFinalRequest:
		response, err = s.service.SignMessageFinal(message.(*messages.SignMessageFinalRequest))

	case messages.TypeVerifyInitRequest:
		response, err = s.service.VerifyInit(message.(*messages.VerifyInitRequest))
	case messages.TypeVerifyRequest:
		response, err = s.service.Verify(message.(*messages.VerifyRequest))
	case messages.TypeVerifyUpdateRequest:
		response, err = s.service.VerifyUpdate(message.(*messages.VerifyUpdateRequest))
	case messages.TypeVerifyFinalRequest:
		response, err = s.service.VerifyFinal(message.(*messages.VerifyFinalRequest))
	case messages.TypeVerifyRecoverInitRequest:
		response, err = s.service.VerifyRecoverInit(message.(*messages.VerifyRecoverInitRequest))
	case messages.TypeVerifyRecoverRequest:
		response, err = s.service.VerifyRecover(message.(*messages.VerifyRecoverRequest))

	case messages.TypeVerifyMessageInitRequest:
		response, err = s.service.VerifyMessageInit(message.(*messages.VerifyMessageInitRequest))
	case messages.TypeVerifyMessageRequest:
		response, err = s.service.VerifyMessage(message.(*messages.VerifyMessageRequest))
	case messages.TypeVerifyMessageBeginRequest:
		response, err = s.service.VerifyMessageBegin(message.(*messages.VerifyMessageBeginRequest))
	case messages.TypeVerifyMessageNextRequest:
		response, err = s.service.VerifyMessageNext(message.(*messages.VerifyMessageNextRequest))
	case messages.TypeVerifyMessageFinalRequest:
		response, err = s.service.VerifyMessageFinal(message.(*messages.VerifyMessageFinalRequest))

	case messages.TypeDigestEncryptUpdateRequest:
		response, err = s.service.DigestEncryptUpdate(message.(*messages.DigestEncryptUpdateRequest))
	case messages.TypeDecryptDigestUpdateRequest:
		response, err = s.service.DecryptDigestUpdate(message.(*messages.DecryptDigestUpdateRequest))
	case messages.TypeSignEncryptUpdateRequest:
		response, err = s.service.SignEncryptUpdate(message.(*messages.SignEncryptUpdateRequest))
	case messages.TypeDecryptVerifyUpdateRequest:
		response, err = s.service.DecryptVerifyUpdate(message.(*messages.DecryptVerifyUpdateRequest))

	case messages.TypeGenerateKeyRequest:
		response, err = s.service.GenerateKey(message.(*messages.GenerateKeyRequest))
	case messages.TypeGenerateKeyPairRequest:
		response, err = s.service.GenerateKeyPair(message.(*messages.GenerateKeyPairRequest))
	case messages.TypeWrapKeyRequest:
		response, err = s.service.WrapKey(message.(*messages.WrapKeyRequest))
	case messages.TypeUnwrapKeyRequest:
		response, err = s.service.UnwrapKey(message.(*messages.UnwrapKeyRequest))
	case messages.TypeDeriveKeyRequest:
		response, err = s.service.DeriveKey(message.(*messages.DeriveKeyRequest))

	case messages.TypeSeedRandomRequest:
		response, err = s.service.SeedRandom(message.(*messages.SeedRandomRequest))
	case messages.TypeGenerateRandomRequest:
		response, err = s.service.GenerateRandom(message.(*messages.GenerateRandomRequest))

	case messages.TypeGetFunctionStatusRequest:
		response, err = s.service.GetFunctionStatus(message.(*messages.GetFunctionStatusRequest))
	case messages.TypeCancelFunctionRequest:
		response, err = s.service.CancelFunction(message.(*messages.CancelFunctionRequest))

	default:
		return nil, errors.New("unsupported message")
	}

	if err != nil {
		return nil, err
	}

	return messages.Marshall(response)
}
