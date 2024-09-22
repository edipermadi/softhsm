package services

import (
	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/sessions"
)

type Service interface {
	Initialize(request *messages.InitializeRequest) (*messages.InitializeResponse, error)
	Finalize(request *messages.FinalizeRequest) (*messages.FinalizeResponse, error)
	GetInfo(request *messages.GetInfoRequest) (*messages.GetInfoResponse, error)
	GetFunctionList(request *messages.GetFunctionListRequest) (*messages.GetFunctionListResponse, error)
	GetInterfaceList(request *messages.GetInterfaceListRequest) (*messages.GetInterfaceListResponse, error)
	GetInterface(request *messages.GetInterfaceRequest) (*messages.GetInterfaceResponse, error)

	GetSlotList(request *messages.GetSlotListRequest) (*messages.GetSlotListResponse, error)
	GetSlotInfo(request *messages.GetSlotInfoRequest) (*messages.GetSlotInfoResponse, error)
	GetTokenInfo(request *messages.GetTokenInfoRequest) (*messages.GetTokenInfoResponse, error)
	WaitForSlotEvent(request *messages.WaitForSlotEventRequest) (*messages.WaitForSlotEventResponse, error)
	GetMechanismList(request *messages.GetMechanismListRequest) (*messages.GetMechanismListResponse, error)
	GetMechanismInfo(request *messages.GetMechanismInfoRequest) (*messages.GetMechanismInfoResponse, error)
	InitToken(request *messages.InitTokenRequest) (*messages.InitTokenResponse, error)
	InitPIN(request *messages.InitPINRequest) (*messages.InitPINResponse, error)
	SetPIN(request *messages.SetPINRequest) (*messages.SetPINResponse, error)

	OpenSession(request *messages.OpenSessionRequest) (*messages.OpenSessionResponse, error)
	CloseSession(request *messages.CloseSessionRequest) (*messages.CloseSessionResponse, error)
	CloseAllSessions(request *messages.CloseAllSessionsRequest) (*messages.CloseAllSessionsResponse, error)
	GetSessionInfo(request *messages.GetSessionInfoRequest) (*messages.GetSessionInfoResponse, error)
	SessionCancel(request *messages.SessionCancelRequest) (*messages.SessionCancelResponse, error)
	GetOperationState(request *messages.GetOperationStateRequest) (*messages.GetOperationStateResponse, error)
	SetOperationState(request *messages.SetOperationStateRequest) (*messages.SetOperationStateResponse, error)
	Login(request *messages.LoginRequest) (*messages.LoginResponse, error)
	LoginUser(request *messages.LoginUserRequest) (*messages.LoginUserResponse, error)
	Logout(request *messages.LogoutRequest) (*messages.LogoutResponse, error)

	CreateObject(request *messages.CreateObjectRequest) (*messages.CreateObjectResponse, error)
	CopyObject(request *messages.CopyObjectRequest) (*messages.CopyObjectResponse, error)
	DestroyObject(request *messages.DestroyObjectRequest) (*messages.DestroyObjectResponse, error)
	GetObjectSize(request *messages.GetObjectSizeRequest) (*messages.GetObjectSizeResponse, error)
	GetAttributeValue(request *messages.GetAttributeValueRequest) (*messages.GetAttributeValueResponse, error)
	SetAttributeValue(request *messages.SetAttributeValueRequest) (*messages.SetAttributeValueResponse, error)
	FindObjectsInit(request *messages.FindObjectsInitRequest) (*messages.FindObjectsInitResponse, error)
	FindObjects(request *messages.FindObjectsRequest) (*messages.FindObjectsResponse, error)
	FindObjectsFinal(request *messages.FindObjectsFinalRequest) (*messages.FindObjectsFinalResponse, error)

	EncryptInit(request *messages.EncryptInitRequest) (*messages.EncryptInitResponse, error)
	Encrypt(request *messages.EncryptRequest) (*messages.EncryptResponse, error)
	EncryptUpdate(request *messages.EncryptUpdateRequest) (*messages.EncryptUpdateResponse, error)
	EncryptFinal(request *messages.EncryptFinalRequest) (*messages.EncryptFinalResponse, error)

	EncryptMessageInit(request *messages.EncryptMessageInitRequest) (*messages.EncryptMessageInitResponse, error)
	EncryptMessage(request *messages.EncryptMessageRequest) (*messages.EncryptMessageResponse, error)
	EncryptMessageBegin(request *messages.EncryptMessageBeginRequest) (*messages.EncryptMessageBeginResponse, error)
	EncryptMessageNext(request *messages.EncryptMessageNextRequest) (*messages.EncryptMessageNextResponse, error)
	EncryptMessageFinal(request *messages.EncryptMessageFinalRequest) (*messages.EncryptMessageFinalResponse, error)

	DecryptInit(request *messages.DecryptInitRequest) (*messages.DecryptInitResponse, error)
	Decrypt(request *messages.DecryptRequest) (*messages.DecryptResponse, error)
	DecryptUpdate(request *messages.DecryptUpdateRequest) (*messages.DecryptUpdateResponse, error)
	DecryptFinal(request *messages.DecryptFinalRequest) (*messages.DecryptFinalResponse, error)

	DecryptMessageInit(request *messages.DecryptMessageInitRequest) (*messages.DecryptMessageInitResponse, error)
	DecryptMessage(request *messages.DecryptMessageRequest) (*messages.DecryptMessageResponse, error)
	DecryptMessageBegin(request *messages.DecryptMessageBeginRequest) (*messages.DecryptMessageBeginResponse, error)
	DecryptMessageNext(request *messages.DecryptMessageNextRequest) (*messages.DecryptMessageNextResponse, error)
	DecryptMessageFinal(request *messages.DecryptMessageFinalRequest) (*messages.DecryptMessageFinalResponse, error)

	DigestInit(request *messages.DigestInitRequest) (*messages.DigestInitResponse, error)
	Digest(request *messages.DigestRequest) (*messages.DigestResponse, error)
	DigestUpdate(request *messages.DigestUpdateRequest) (*messages.DigestUpdateResponse, error)
	DigestKey(request *messages.DigestKeyRequest) (*messages.DigestKeyResponse, error)
	DigestFinal(request *messages.DigestFinalRequest) (*messages.DigestFinalResponse, error)

	SignInit(request *messages.SignInitRequest) (*messages.SignInitResponse, error)
	Sign(request *messages.SignRequest) (*messages.SignResponse, error)
	SignUpdate(request *messages.SignUpdateRequest) (*messages.SignUpdateResponse, error)
	SignFinal(request *messages.SignFinalRequest) (*messages.SignFinalResponse, error)
	SignRecoverInit(request *messages.SignRecoverInitRequest) (*messages.SignRecoverInitResponse, error)
	SignRecover(request *messages.SignRecoverRequest) (*messages.SignRecoverResponse, error)

	SignMessageInit(request *messages.SignMessageInitRequest) (*messages.SignMessageInitResponse, error)
	SignMessage(request *messages.SignMessageRequest) (*messages.SignMessageResponse, error)
	SignMessageBegin(request *messages.SignMessageBeginRequest) (*messages.SignMessageBeginResponse, error)
	SignMessageNext(request *messages.SignMessageNextRequest) (*messages.SignMessageNextResponse, error)
	SignMessageFinal(request *messages.SignMessageFinalRequest) (*messages.SignMessageFinalResponse, error)

	VerifyInit(request *messages.VerifyInitRequest) (*messages.VerifyInitResponse, error)
	Verify(request *messages.VerifyRequest) (*messages.VerifyResponse, error)
	VerifyUpdate(request *messages.VerifyUpdateRequest) (*messages.VerifyUpdateResponse, error)
	VerifyFinal(request *messages.VerifyFinalRequest) (*messages.VerifyFinalResponse, error)
	VerifyRecoverInit(request *messages.VerifyRecoverInitRequest) (*messages.VerifyRecoverInitResponse, error)
	VerifyRecover(request *messages.VerifyRecoverRequest) (*messages.VerifyRecoverResponse, error)

	VerifyMessageInit(request *messages.VerifyMessageInitRequest) (*messages.VerifyMessageInitResponse, error)
	VerifyMessage(request *messages.VerifyMessageRequest) (*messages.VerifyMessageResponse, error)
	VerifyMessageBegin(request *messages.VerifyMessageBeginRequest) (*messages.VerifyMessageBeginResponse, error)
	VerifyMessageNext(request *messages.VerifyMessageNextRequest) (*messages.VerifyMessageNextResponse, error)
	VerifyMessageFinal(request *messages.VerifyMessageFinalRequest) (*messages.VerifyMessageFinalResponse, error)

	DigestEncryptUpdate(request *messages.DigestEncryptUpdateRequest) (*messages.DigestEncryptUpdateResponse, error)
	DecryptDigestUpdate(request *messages.DecryptDigestUpdateRequest) (*messages.DecryptDigestUpdateResponse, error)
	SignEncryptUpdate(request *messages.SignEncryptUpdateRequest) (*messages.SignEncryptUpdateResponse, error)
	DecryptVerifyUpdate(request *messages.DecryptVerifyUpdateRequest) (*messages.DecryptVerifyUpdateResponse, error)

	GenerateKey(request *messages.GenerateKeyRequest) (*messages.GenerateKeyResponse, error)
	GenerateKeyPair(request *messages.GenerateKeyPairRequest) (*messages.GenerateKeyPairResponse, error)
	WrapKey(request *messages.WrapKeyRequest) (*messages.WrapKeyResponse, error)
	UnwrapKey(request *messages.UnwrapKeyRequest) (*messages.UnwrapKeyResponse, error)
	DeriveKey(request *messages.DeriveKeyRequest) (*messages.DeriveKeyResponse, error)

	SeedRandom(request *messages.SeedRandomRequest) (*messages.SeedRandomResponse, error)
	GenerateRandom(request *messages.GenerateRandomRequest) (*messages.GenerateRandomResponse, error)

	GetFunctionStatus(request *messages.GetFunctionStatusRequest) (*messages.GetFunctionStatusResponse, error)
	CancelFunction(request *messages.CancelFunctionRequest) (*messages.CancelFunctionResponse, error)
}

func NewService() Service {
	return &service{sessions: make(map[int]sessions.Session)}
}

type service struct {
	sessions map[int]sessions.Session
}
