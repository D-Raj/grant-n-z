package service
//
//import (
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/labstack/echo"
//	"github.com/satori/go.uuid"
//	"github.com/tomoyane/grant-n-z/domain/entity"
//	"github.com/tomoyane/grant-n-z/domain/repository"
//	"github.com/tomoyane/grant-n-z/handler"
//	"github.com/tomoyane/grant-n-z/infra"
//	"golang.org/x/crypto/bcrypt"
//	"net/http"
//	"strings"
//	"time"
//)
//
//type TokenService struct {
//	TokenRepository repository.TokenRepository
//	UserRepository repository.UserRepository
//	PrincipalRepository repository.PrincipalRepository
//	MemberRepository repository.MemberRepository
//}
//
//func (t TokenService) ComparePw(passwordHash string, password string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
//	if err != nil {
//		return false
//	}
//
//	return true
//}
//
//func (t TokenService) GenerateJwt(username string, userUuid uuid.UUID, members []*entity.Member) string {
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	claims := token.Claims.(jwt.MapClaims)
//	claims["username"] = username
//	claims["user_uuid"] = userUuid.String()
//	claims["expires"] = time.Now().Add(time.Hour * 365).String()
//
//	if len(members) != 0 {
//		var memberUuid string
//		for _, member := range members {
//			memberUuid += member.Uuid.String() + ","
//		}
//
//		fmt.Println(memberUuid)
//		claims["member_uuid"] = memberUuid
//	}
//
//	signedToken, err := token.SignedString([]byte(infra.Yaml.App.PrivateKey))
//	if err != nil {
//		handler.ErrorResponse{}.Print(http.StatusInternalServerError, "failed generate jwt", "")
//		return ""
//	}
//
//	return signedToken
//}
//
//func (t TokenService) ParseJwt(token string) (map[string]string, bool) {
//	resultMap := map[string]string{}
//
//	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
//		return []byte(infra.Yaml.App.PrivateKey), nil
//	})
//
//	if err != nil || !parseToken.Valid {
//		return resultMap, false
//	}
//
//	claims := parseToken.Claims.(jwt.MapClaims)
//	if _, ok := claims["username"].(string); !ok {
//		return resultMap, false
//	}
//
//	if _, ok := claims["user_uuid"].(string); !ok {
//		return resultMap, false
//	}
//
//	if _, ok := claims["expires"].(string); !ok {
//		return resultMap, false
//	}
//
//	if _, ok := claims["member_uuid"].(string); ok {
//		resultMap["member_uuid"] = claims["member_uuid"].(string)
//	}
//
//	resultMap["username"] = claims["username"].(string)
//	resultMap["user_uuid"] = claims["user_uuid"].(string)
//	resultMap["expires"] = claims["expires"].(string)
//
//	return resultMap, true
//}
//
//func (t TokenService) GetTokenByUserUuid(userUuid string) *entity.Token {
//	return t.TokenRepository.FindByUserUuid(userUuid)
//}
//
//func (t TokenService) InsertToken(userUuid uuid.UUID, token string, refreshToken string) *entity.Token {
//	data := entity.Token{
//		TokenType: "Bearer",
//		Token: token,
//		RefreshToken: refreshToken,
//		UserUuid: userUuid,
//	}
//	return t.TokenRepository.Save(data)
//}
//
//func (t TokenService) VerifyToken(c echo.Context, token string, option string) (*handler.ErrorResponse) {
//	if token == "" {
//		return handler.Unauthorized("")
//	}
//
//	resultMap, result := t.ParseJwt(token)
//	if !result {
//		return handler.Unauthorized("")
//	}
//
//	user := t.UserRepository.FindByUuid(resultMap["user_uuid"])
//	if user == nil {
//		return handler.InternalServerError("")
//	}
//
//	if len(user.Email) == 0 {
//		return handler.Unauthorized("")
//	}
//
//	// Any permission
//	if option != "" {
//		var memberUuids []string
//		memberUuids = strings.Split(resultMap["member_uuid"], ",")
//
//		principals := t.PrincipalRepository.FindByMemberUuid(memberUuids)
//		if principals == nil {
//			return handler.InternalServerError("")
//		}
//
//		for _, principal := range principals {
//			if strings.Contains(principal.Role.Permission, option) {
//				return nil
//			}
//		}
//
//		return handler.Forbidden("")
//
//	} else {
//		return nil
//	}
//
//	//role := roleService.GetRoleByUserUuid(user.Uuid.String())
//	//if role == nil {
//	//	return echo.NewHTTPError(http.StatusInternalServerError, handler.InternalServerError(""))
//	//}
//
//	//if len(role.UserUuid) == 0 {
//	//	return echo.NewHTTPError(http.StatusForbidden, handler.Forbidden("019"))
//	//}
//	//
//	//if role.Type != "user" && role.Type != "admin" {
//	//	return echo.NewHTTPError(http.StatusForbidden, handler.Forbidden("020"))
//	//}
//}
//
//func (t TokenService) IssueToken(user *entity.User) (issueToken *entity.Token, errRes *handler.ErrorResponse) {
//	userData := t.UserRepository.FindByEmail(user.Email)
//	if userData == nil {
//		return nil, handler.InternalServerError("")
//	}
//
//	if len(userData.Email) == 0 {
//		return nil, handler.NotFound("")
//	}
//
//	if !t.ComparePw(userData.Password, user.Password) {
//		return nil, handler.UnProcessableEntity("")
//	}
//
//	members := t.MemberRepository.FindByUserUuid(userData.Uuid)
//	if members == nil {
//		return nil, handler.InternalServerError("")
//	}
//
//	tokenStr := t.GenerateJwt(userData.Username, userData.Uuid, members)
//	refreshTokenStr := t.GenerateJwt(userData.Username, userData.Uuid, members)
//
//	if tokenStr == "" || refreshTokenStr == ""{
//		return nil, handler.InternalServerError("")
//	}
//
//	token := t.InsertToken(userData.Uuid, tokenStr, refreshTokenStr)
//	if token == nil {
//		return nil, handler.InternalServerError("")
//	}
//
//	return token, nil
//}