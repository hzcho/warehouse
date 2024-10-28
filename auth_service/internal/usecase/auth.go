package usecase

import (
	"auth/internal/config"
	"auth/internal/domain/model"
	"auth/internal/domain/net/request"
	"auth/internal/domain/net/response"
	"auth/internal/domain/repository"
	"auth/pkg/token"
	"context"
	"fmt"
	"time"

	"errors"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	userRepo     repository.User
	log          *logrus.Logger
	cfg          config.Auth
	tokenManager token.TokenManager
}

func NewAuth(userRepo repository.User, log *logrus.Logger, cfg config.Auth, tokenManager token.TokenManager) *Auth {
	return &Auth{
		userRepo:     userRepo,
		log:          log,
		cfg:          cfg,
		tokenManager: tokenManager,
	}
}

func (u *Auth) SignUp(ctx context.Context, req request.SignUp) error {
	log := u.log.WithFields(logrus.Fields{
		"op":           "internal/usecase/auth/SignUp",
		"login":        req.Login,
		"password":     req.Password,
		"phone_number": req.PhoneNumber,
		"email":        req.Email,
	})

	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return err
	}

	user := model.User{
		ID:           uuid.New(),
		Login:        req.Login,
		Password:     string(passHash),
		Role:         "employee",
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		RefreshToken: nil,
		TokenExpiry:  nil,
	}

	log.Info(user)

	_, err = u.userRepo.Create(ctx, user)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (u *Auth) SignIn(ctx context.Context, req request.SignIn) (response.Token, error) {
	log := u.log.WithFields(logrus.Fields{
		"op":       "internal/usecase/auth/SignIn",
		"login":    req.Login,
		"password": req.Password,
	})

	user, err := u.userRepo.Get(ctx, req.Login)
	if err != nil {
		log.Error(err.Error())
		return response.Token{}, err
	}
	if user == (model.User{}) {
		err := errors.New("the user does not exist")

		log.Error(err.Error())

		return response.Token{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Error(err.Error())
		return response.Token{}, err
	}

	accessToken, err := u.tokenManager.NewJWT(token.AuthInfo{
		UserID: user.ID.String(),
		Login:  user.Login,
		Role:   user.Role,
	})
	if err != nil {
		log.Error(err.Error())
		return response.Token{}, err
	}

	refreshToken, err := u.tokenManager.RefreshToken()
	if err != nil {
		log.Error(err.Error())
		return response.Token{}, err
	}

	user.RefreshToken = &refreshToken

	expTime := time.Now().Add(u.cfg.RFDuration)
	user.TokenExpiry = &expTime

	if err := u.userRepo.Update(ctx, user); err != nil {
		log.Error(err.Error())
		return response.Token{}, err
	}

	return response.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *Auth) RefreshToken(ctx context.Context, tkn request.RefreshToken) (response.Token, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/auth/RefreshToken",
	})

	claims, err := u.tokenManager.Parse(tkn.AccessToken)
	if err != nil {
		log.Errorf("access token: %v", err)
		return response.Token{}, err
	}

	user, err := u.userRepo.Get(ctx, claims.Login)
	if err != nil {
		log.Error(err)
		return response.Token{}, err
	}

	log.Infof("sisisisisis %s %v %s", user.Login, user.RefreshToken, user.Role)

	if user.TokenExpiry == nil {
		err := fmt.Errorf("refresh token is empty")
		log.Error(err)
		return response.Token{}, err
	}
	if time.Now().After(*user.TokenExpiry) {
		err := fmt.Errorf("refresh token expire")
		log.Error(err)
		return response.Token{}, err
	}

	ok, err := u.tokenManager.ValidateRefreshToken(tkn.RefreshToken)
	if err != nil {
		log.Errorf("refresh token: %v", err)
		return response.Token{}, err
	}
	if !ok {
		err := fmt.Errorf("invalid token")
		log.Error(err)
		return response.Token{}, err
	}

	accessToken, err := u.tokenManager.NewJWT(token.AuthInfo{
		UserID: user.ID.String(),
		Login:  user.Login,
		Role:   user.Role,
	})
	if err != nil {
		log.Error(err.Error())
		return response.Token{}, err
	}

	refreshToken, err := u.tokenManager.RefreshToken()
	if err != nil {
		log.Error(err.Error())
		return response.Token{}, err
	}

	user.RefreshToken = &refreshToken

	expTime := time.Now().Add(u.cfg.RFDuration)
	user.TokenExpiry = &expTime

	if err := u.userRepo.Update(ctx, user); err != nil {
		log.Error(err.Error())
		return response.Token{}, err
	}

	return response.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
