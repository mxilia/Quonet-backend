package rest

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/internal/user/dto"
	"github.com/mxilia/Conflux-backend/internal/user/usecase"
	appError "github.com/mxilia/Conflux-backend/pkg/apperror"
	"github.com/mxilia/Conflux-backend/pkg/config"
	"github.com/mxilia/Conflux-backend/pkg/responses"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type HttpUserHandler struct {
	userUseCase       usecase.UserUseCase
	oauthGoogleConfig *oauth2.Config
}

func NewHttpUserHandler(userUseCase usecase.UserUseCase, cfg *config.Config) *HttpUserHandler {
	return &HttpUserHandler{
		userUseCase: userUseCase,
		oauthGoogleConfig: &oauth2.Config{
			ClientID:     cfg.GOOGLE_CLIENT_ID,
			ClientSecret: cfg.GOOGLE_CLIENT_SECRET,
			RedirectURL:  cfg.GOOGLE_OAUTH_REDIRECT_URL,
			Scopes:       []string{"email", "profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (h *HttpUserHandler) GoogleLogin(c *fiber.Ctx) error {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HTTPOnly: true,
	})
	url := h.oauthGoogleConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce, oauth2.SetAuthURLParam("prompt", "consent select_account"))
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func (h *HttpUserHandler) GoogleCallBack(c *fiber.Ctx) error {
	state := c.Cookies("oauthstate")
	if c.Query("state") != state {
		return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "invalid oauth state")
	}

	code := c.Query("code")
	if code == "" {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "code not found")
	}

	token, err := h.oauthGoogleConfig.Exchange(c.Context(), code)
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInternalServer, "cannot exchange token")
	}

	client := h.oauthGoogleConfig.Client(c.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return responses.ErrorWithMessage(c, err, "failed to get user info from google api")
	}

	var googleReq dto.CreateUserByGoogleRequest
	if err := json.NewDecoder(resp.Body).Decode(&googleReq); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to decode user info")
	}

	userInfo := &entities.User{Email: googleReq.Email, ProfileUrl: googleReq.ProfileUrl}

	user, err := h.userUseCase.GoogleUserEntry(userInfo)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "failed to create user via google login")
	}
	return c.JSON(dto.ToUserResponse(user))
}

func (h *HttpUserHandler) FindAllUsers(c *fiber.Ctx) error {
	users, err := h.userUseCase.FindAllUsers()
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToUserResponseList(users))
}

func (h *HttpUserHandler) FindUserByID(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	user, err := h.userUseCase.FindUserByID(userID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "failed to find user by id")
	}
	return c.JSON(dto.ToUserResponse(user))

}

func (h *HttpUserHandler) FindUserByHandler(c *fiber.Ctx) error {
	user, err := h.userUseCase.FindUserByHandler(c.Params("handler"))
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToUserResponse(user))
}

func (h *HttpUserHandler) FindUserByEmail(c *fiber.Ctx) error {
	user, err := h.userUseCase.FindUserByEmail(c.Params("email"))
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToUserResponse(user))
}

func (h *HttpUserHandler) PatchUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	var req dto.UserPatchRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, err)
	}

	userInfo := dto.FromUserPatchRequest(&req)
	if err := h.userUseCase.PatchUser(userID, userInfo); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "patch successfully")
}
func (h *HttpUserHandler) DeleteUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	if err := h.userUseCase.DeleteUser(userID); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to delete user by id")
	}

	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
