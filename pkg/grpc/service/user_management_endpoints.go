package service

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/influenzanet/go-utils/pkg/constants"
	loggingAPI "github.com/influenzanet/logging-service/pkg/api"
	messageAPI "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/user-management-service/pkg/api"
	"github.com/influenzanet/user-management-service/pkg/models"
	"github.com/influenzanet/user-management-service/pkg/pwhash"
	"github.com/influenzanet/user-management-service/pkg/tokens"
	"github.com/influenzanet/user-management-service/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userManagementServer) CreateUser(ctx context.Context, req *api.CreateUserReq) (*api.User, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) || req.AccountId == "" || req.InitialPassword == "" {
		return nil, status.Error(codes.InvalidArgument, "missing arguments")
	}
	if !utils.CheckRoleInToken(req.Token, constants.USER_ROLE_ADMIN) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	req.AccountId = utils.SanitizeEmail(req.AccountId)
	if !utils.CheckEmailFormat(req.AccountId) {
		return nil, status.Error(codes.InvalidArgument, "account id not a valid email")
	}
	if !utils.CheckPasswordFormat(req.InitialPassword) {
		return nil, status.Error(codes.InvalidArgument, "password too weak")
	}

	password, err := pwhash.HashPassword(req.InitialPassword)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Create user DB object from request:
	newUser := models.User{
		Account: models.Account{
			Type:               "email",
			AccountID:          req.AccountId,
			AccountConfirmedAt: 0, // not confirmed yet
			Password:           password,
			PreferredLanguage:  req.PreferredLanguage,
		},
		Roles: req.Roles,
		Profiles: []models.Profile{
			{
				ID:                 primitive.NewObjectID(),
				Alias:              req.AccountId,
				AvatarID:           "default",
				ConsentConfirmedAt: time.Now().Unix(),
				MainProfile:        true,
			},
		},
		Timestamps: models.Timestamps{
			CreatedAt: time.Now().Unix() + userCreationTimestampOffset,
		},
	}
	newUser.AddNewEmail(req.AccountId, false)
	if req.Use_2Fa {
		newUser.Account.AuthType = "2FA"
	}
	newUser.ContactPreferences.ReceiveWeeklyMessageDayOfWeek = int32(rand.Intn(7))

	instanceID := req.Token.InstanceId
	id, err := s.userDBservice.AddUser(instanceID, newUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	newUser.ID, _ = primitive.ObjectIDFromHex(id)

	// TempToken for contact verification:
	tempTokenInfos := models.TempToken{
		UserID:     id,
		InstanceID: instanceID,
		Purpose:    constants.TOKEN_PURPOSE_INVITATION,
		Info: map[string]string{
			"type":  "email",
			"email": newUser.Account.AccountID,
		},
		Expiration: tokens.GetExpirationTime(time.Hour * 24 * 7),
	}
	tempToken, err := s.globalDBService.AddTempToken(tempTokenInfos)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// ---> Trigger message sending
	_, err = s.clients.MessagingService.SendInstantEmail(ctx, &messageAPI.SendEmailReq{
		InstanceId:  instanceID,
		To:          []string{newUser.Account.AccountID},
		MessageType: constants.EMAIL_TYPE_INVITATION,
		ContentInfos: map[string]string{
			"token": tempToken,
		},
		PreferredLanguage: newUser.Account.PreferredLanguage,
		UseLowPrio:        true,
	})
	if err != nil {
		log.Printf("CreateUser: %s", err.Error())
	}
	// <---

	s.SaveLogEvent(req.Token.InstanceId, req.Token.Id, loggingAPI.LogEventType_LOG, constants.LOG_EVENT_ACCOUNT_CREATED, "by admin - "+newUser.ID.Hex()+" - "+newUser.Account.AccountID)

	return newUser.ToAPI(), nil
}

func (s *userManagementServer) AddRoleForUser(ctx context.Context, req *api.RoleMsg) (*api.User, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) || req.AccountId == "" || req.Role == "" {
		return nil, status.Error(codes.InvalidArgument, "missing arguments")
	}
	if !utils.CheckRoleInToken(req.Token, constants.USER_ROLE_ADMIN) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	user, err := s.userDBservice.GetUserByAccountID(req.Token.InstanceId, req.AccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := user.AddRole(req.Role); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user, err = s.userDBservice.UpdateUser(req.Token.InstanceId, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.SaveLogEvent(req.Token.InstanceId, req.Token.Id, loggingAPI.LogEventType_LOG, constants.LOG_EVENT_ACCOUNT_ROLE_ADDED, user.Account.AccountID+"("+user.ID.Hex()+") + "+req.Role)

	return user.ToAPI(), nil
}

func (s *userManagementServer) RemoveRoleForUser(ctx context.Context, req *api.RoleMsg) (*api.User, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) || req.AccountId == "" || req.Role == "" {
		return nil, status.Error(codes.InvalidArgument, "missing arguments")
	}
	if !utils.CheckRoleInToken(req.Token, constants.USER_ROLE_ADMIN) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	user, err := s.userDBservice.GetUserByAccountID(req.Token.InstanceId, req.AccountId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := user.RemoveRole(req.Role); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user, err = s.userDBservice.UpdateUser(req.Token.InstanceId, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	s.SaveLogEvent(req.Token.InstanceId, req.Token.Id, loggingAPI.LogEventType_LOG, constants.LOG_EVENT_ACCOUNT_ROLE_REMOVED, user.Account.AccountID+"("+user.ID.Hex()+") - "+req.Role)
	return user.ToAPI(), nil
}

func (s *userManagementServer) FindNonParticipantUsers(ctx context.Context, req *api.FindNonParticipantUsersMsg) (*api.UserListMsg, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) {
		return nil, status.Error(codes.InvalidArgument, "missing arguments")
	}
	if !utils.CheckRoleInToken(req.Token, constants.USER_ROLE_ADMIN) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	users, err := s.userDBservice.FindNonParticipantUsers(req.Token.InstanceId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	resp := api.UserListMsg{
		Users: make([]*api.User, len(users)),
	}
	for i, u := range users {
		resp.Users[i] = u.ToAPI()
	}
	return &resp, nil
}

func (s *userManagementServer) StreamUsers(req *api.StreamUsersMsg, stream api.UserManagementApi_StreamUsersServer) error {
	if req == nil || stream == nil || req.InstanceId == "" {
		return status.Error(codes.InvalidArgument, "missing arguments")
	}

	sendUserOverGrpc := func(instanceID string, user models.User, args ...interface{}) error {
		if len(args) != 1 {
			return errors.New("StreamUsers callback: unexpected number of args")
		}
		stream, ok := args[0].(api.UserManagementApi_StreamUsersServer)
		if !ok {
			return errors.New(("StreamUsers callback: can't parse stream"))
		}

		if err := stream.Send(user.ToAPI()); err != nil {
			return err
		}
		return nil
	}

	err := s.userDBservice.PerfomActionForUsers(req.InstanceId, sendUserOverGrpc, stream)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
