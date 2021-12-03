package service

type Services struct {
	Register             RegisterService
	GetAuthenticatedUser GetAuthenticatedUserService
	GetPeers             GetPeersService
	PostDirectMessage    PostDirectMessageService
	GetDirectMessages    GetDirectMessagesService
	SetMessageSeen       SetMessageSeenService
}