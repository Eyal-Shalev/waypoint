package singleprocess

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/cap/oidc"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/hashicorp/waypoint/internal/server/gen"
	serverptypes "github.com/hashicorp/waypoint/internal/server/ptypes"
)

const (
	// oidcAuthExpiry is the duration that an OIDC-based login is valid for.
	// We default this to 30 days for now but that is arbitrary. We can change
	// this default anytime or choose to make it configurable one day on the
	// server.
	oidcAuthExpiry = 30 * 24 * time.Hour

	// oidcReqExpiry is the time that an OIDC auth request is valid for.
	// 5 minutes should be plenty of time to complete auth.
	oidcReqExpiry = 5 * 60 * time.Minute
)

func (s *service) ListOIDCAuthMethods(
	ctx context.Context,
	req *empty.Empty,
) (*pb.ListOIDCAuthMethodsResponse, error) {
	// We implement this by just requesting all the auth methods. We could
	// index OIDC methods specifically and do this more efficiently but
	// realistically we don't expect there to ever be that many auth methods.
	// Even if there were thousands (why????) this would be okay.
	values, err := s.state.AuthMethodList()
	if err != nil {
		return nil, err
	}

	// Go through and extract the auth methods
	var result []*pb.OIDCAuthMethod
	for _, method := range values {
		_, ok := method.Method.(*pb.AuthMethod_Oidc)
		if !ok {
			continue
		}

		result = append(result, &pb.OIDCAuthMethod{
			Name:        method.Name,
			DisplayName: method.DisplayName,
			Kind:        pb.OIDCAuthMethod_UNKNOWN,
		})
	}

	return &pb.ListOIDCAuthMethodsResponse{AuthMethods: result}, nil
}

func (s *service) GetOIDCAuthURL(
	ctx context.Context,
	req *pb.GetOIDCAuthURLRequest,
) (*pb.GetOIDCAuthURLResponse, error) {
	if err := serverptypes.ValidateGetOIDCAuthURLRequest(req); err != nil {
		return nil, err
	}

	// Get the auth method
	am, err := s.state.AuthMethodGet(req.AuthMethod)
	if err != nil {
		return nil, err
	}

	// The auth method must be OIDC
	amMethod, ok := am.Method.(*pb.AuthMethod_Oidc)
	if !ok {
		return nil, status.Errorf(codes.FailedPrecondition,
			"auth method is not OIDC")
	}

	// We need our server config.
	sc, err := s.state.ServerConfigGet()
	if err != nil {
		return nil, err
	}

	// Get our OIDC provider
	provider, err := s.oidcCache.Get(ctx, am, sc)
	if err != nil {
		return nil, err
	}

	// Create a minimal request to get the auth URL
	oidcReqOpts := []oidc.Option{
		oidc.WithNonce(req.Nonce),
	}
	if v := amMethod.Oidc.Scopes; len(v) > 0 {
		oidcReqOpts = append(oidcReqOpts, oidc.WithScopes(v...))
	}
	oidcReq, err := oidc.NewRequest(
		oidcReqExpiry,
		req.RedirectUri,
		oidcReqOpts...,
	)
	if err != nil {
		return nil, err
	}

	// Get the auth URL
	url, err := provider.AuthURL(ctx, oidcReq)
	if err != nil {
		return nil, err
	}

	return &pb.GetOIDCAuthURLResponse{
		Url: url,
	}, nil
}

func (s *service) CompleteOIDCAuth(
	ctx context.Context,
	req *pb.CompleteOIDCAuthRequest,
) (*pb.CompleteOIDCAuthResponse, error) {
	log := hclog.FromContext(ctx)

	if err := serverptypes.ValidateCompleteOIDCAuthRequest(req); err != nil {
		return nil, err
	}

	// Get the auth method
	am, err := s.state.AuthMethodGet(req.AuthMethod)
	if err != nil {
		return nil, err
	}

	// The auth method must be OIDC
	amMethod, ok := am.Method.(*pb.AuthMethod_Oidc)
	if !ok {
		return nil, status.Errorf(codes.FailedPrecondition,
			"auth method is not OIDC")
	}

	// We need our server config.
	sc, err := s.state.ServerConfigGet()
	if err != nil {
		return nil, err
	}

	// Get our OIDC provider
	provider, err := s.oidcCache.Get(ctx, am, sc)
	if err != nil {
		return nil, err
	}

	// Create a minimal request to get the auth URL
	oidcReqOpts := []oidc.Option{
		oidc.WithNonce(req.Nonce),
		oidc.WithState(req.State),
	}
	if v := amMethod.Oidc.Scopes; len(v) > 0 {
		oidcReqOpts = append(oidcReqOpts, oidc.WithScopes(v...))
	}
	if v := amMethod.Oidc.Auds; len(v) > 0 {
		oidcReqOpts = append(oidcReqOpts, oidc.WithAudiences(v...))
	}
	oidcReq, err := oidc.NewRequest(
		oidcReqExpiry,
		req.RedirectUri,
		oidcReqOpts...,
	)
	if err != nil {
		return nil, err
	}

	// Exchange our code for our token
	oidcToken, err := provider.Exchange(ctx, oidcReq, req.State, req.Code)
	if err != nil {
		return nil, err
	}

	// Extract the claims as a raw JSON message.
	var jsonClaims json.RawMessage
	if err := oidcToken.IDToken().Claims(&jsonClaims); err != nil {
		return nil, err
	}

	// Structurally extract only the claim fields we care about.
	var idClaimVals idClaims
	if err := json.Unmarshal([]byte(jsonClaims), &idClaimVals); err != nil {
		return nil, err
	}

	// Valid OIDC providers should never behave this way.
	if idClaimVals.Sub == "" {
		return nil, status.Errorf(codes.Internal, "OIDC provider returned empty subscriber ID")
	}

	// Look up a user by sub.
	user, err := s.oidcInitUser(log, &idClaimVals)
	if err != nil {
		return nil, err
	}

	// Generate a token for this user
	token, err := s.newToken(oidcAuthExpiry, DefaultKeyId, nil, &pb.Token{
		Kind: &pb.Token_Login_{
			Login: &pb.Token_Login{UserId: user.Id},
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CompleteOIDCAuthResponse{
		Token:      token,
		User:       user,
		ClaimsJson: string(jsonClaims),
	}, nil
}

// oidcInitUser finds or creates the user for the given OIDC information.
func (s *service) oidcInitUser(log hclog.Logger, claims *idClaims) (*pb.User, error) {
	// This method attempts to find, link, or create a new user to the
	// given OIDC result in the following order:
	//
	//   (1) find user with exact account link (iss, sub match)
	//   (2) find user with matching email and then link it
	//   (3) create new user and link it
	//

	// The email for the user. We only set this if the email is known and
	// verified. This prevents impersonation.
	var email string
	if claims.Email != "" && claims.EmailVerified {
		email = claims.Email
	}

	// First look up by exact account link.
	user, err := s.state.UserGetOIDC(claims.Iss, claims.Sub)
	if err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, err
		}

		// Just ensure user is nil cause that's the check we'll keep using.
		user = nil
	}
	if user != nil {
		return user, nil
	}

	// Look up the user by email if we don't have a user by sub.
	if email != "" {
		user, err = s.state.UserGetEmail(email)
		if err != nil {
			if status.Code(err) != codes.NotFound {
				return nil, err
			}

			// Just ensure user is nil cause that's the check we'll keep using.
			user = nil
		}
	}

	// If the user still doesn't exist, we create a new user.
	if user == nil {
		// Random username to start.
		// NOTE(mitchellh): we can improve this in a ton of ways in
		// the future by using their preferred username claim, first name,
		// etc.
		username := fmt.Sprintf("user_%d", time.Now().Unix())

		user = &pb.User{
			Username: username,
			Email:    email,
		}
	}

	// Setup their link
	user.Links = append(user.Links, &pb.User_Link{
		Method: &pb.User_Link_Oidc{
			Oidc: &pb.User_Link_OIDC{
				Iss: claims.Iss,
				Sub: claims.Sub,
			},
		},
	})

	if err := s.state.UserPut(user); err != nil {
		return nil, err
	}

	log.Info("new OIDC user linked",
		"user_id", user.Id,
		"username", user.Username,
		"iss", claims.Iss,
		"sub", claims.Sub,
	)

	return user, nil
}

// idClaims are the claims for the ID token that we care about. There
// are many more claims[1] but we only add what we need.
//
// [1]: https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims
type idClaims struct {
	Iss           string `json:"iss"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}