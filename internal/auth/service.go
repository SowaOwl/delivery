package auth

import (
	"delivery/internal/cache"
	"delivery/internal/model"
	"delivery/internal/repository"
	"delivery/internal/util/debug"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type AuthService interface {
	Login(email string, password string) (string, error)
	Authenticate(token string) error
}

type AuthServiceImpl struct {
	jwtService     JWTService
	cacheService   cache.Cache
	userRepository repository.UserRepository
}

func NewAuthService(jwtService JWTService, cacheService cache.Cache, userRepo repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		jwtService:     jwtService,
		cacheService:   cacheService,
		userRepository: userRepo,
	}
}

func (a *AuthServiceImpl) Login(email string, password string) (string, error) {
	user, err := a.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if checkPassword(password, user.Password) {
		jsonUser, err := json.Marshal(user)
		if err != nil {
			return "", err
		}

		err = a.cacheService.Set("user_"+strconv.Itoa(int(user.ID)), jsonUser)
		if err != nil {
			return "", err
		}

		token, err := a.jwtService.CreateToken(user.ID)
		if err != nil {
			return "", err
		}

		return token, nil
	} else {
		return "", errors.New("invalid login or password")
	}
}

func (a *AuthServiceImpl) Authenticate(token string) error {
	claims, err := a.jwtService.ParseToken(token)
	if err != nil {
		return err
	}

	userID, ok := (*claims)["user_id"].(string)
	if !ok {
		return errors.New("user ID not found in token")
	}

	key := "user_" + userID

	cachedUser, err := a.cacheService.Get(key)
	if err != nil {
		return err
	}

	var user model.User

	err = json.Unmarshal([]byte(cachedUser), &user)
	if err != nil {
		return err
	}

	debug.JsonPrint(user)

	return nil
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
