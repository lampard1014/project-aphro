package Session
import (
	"time"
	"math/rand"
	"strconv"
	"github.com/lampard1014/aphro/PersistentStore/Redis"
	"github.com/lampard1014/aphro/Encryption"
	"strings"
	"github.com/lampard1014/aphro/CommonBiz/Error"
)

const (
	tokenDuration = 24 * 3600 * time.Second //1 day
)

const (
	//  common biz error //
	SessionServiceNoError   = 0
	//session过期
	SessionServiceError_SessionExpired   = iota + 700
)


var ErrorMap map[int]string = map[int]string{
	SessionServiceNoError : "",
	SessionServiceError_SessionExpired:"session expired",
}

func FetchSessionTokenValue(sessionToken string) (uid string, merchantID string, err error) {
	var returnErr error = nil

	if sessionToken == "" {
		returnErr = Error.NewCustomError(Error.BizError, "session 过期 请重新登录")

	} else {
		token,_,err := QuerySessionToken(sessionToken)

		if err == nil && token != "" {
			sessionTokenValue := token
			splitValue := strings.Split(sessionTokenValue, "#")
			uidAndMerchantID := strings.Split(splitValue[0],"@")
			var _uid uint64
			_uid,returnErr = strconv.ParseUint(uidAndMerchantID[0],36,0)
			uid = strconv.FormatUint(_uid,10)

			var _merchantID uint64
			_merchantID,returnErr = strconv.ParseUint(uidAndMerchantID[1],36,0)
			merchantID = strconv.FormatUint(_merchantID,10)

			_, returnErr = RenewSessionToken(sessionToken)
		} else {
			returnErr = Error.NewCustomError(Error.BizError, "session 过期 请重新登录")
		}
	}

	return uid,merchantID,returnErr
}

func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func QuerySessionToken(sessionToken string) (token string, ttl int64,  err error) {
	redis ,err := Redis.NewAPSRedis(nil)
	if err != nil {
		return "",0,err
	} else {
		redis.Connect()
		defer redis.Close()
		var returnErr error = nil
		qr,err1 := redis.Query(sessionToken)
		qtr, err2 := redis.QueryTTL(sessionToken)
		var hasErr bool = (err1 !=nil || err2 != nil)
		hasErr = hasErr && qtr > 0

		if hasErr {
			if err1 != nil {
				returnErr = err1
			} else {
				returnErr = err2
			}
		}
		return qr,qtr,returnErr
	}
}

func DecodeSessionTokenRequestStr(sessionTokenRequestStr string)(string ,error) {
	data , err := Encryption.Base64Decode(sessionTokenRequestStr)
	if err == nil {
		rsadata,err := Encryption.RsaDecryption(data)
		return string(rsadata),err
	} else {
		return "",err
	}
}

func CreateSessionToken(sessionTokenRequestStr string,uid uint32, merchantID uint32) (token string, ttl int64,  err error) {
	_uid := strconv.FormatUint(uint64(uid),36)
	_merchantID := strconv.FormatUint(uint64(merchantID),36)

	//general key
	t :=  time.Now().Unix()
	var tokenKey string
	tokenKeyPrefix := _uid
	tokenKey = strconv.FormatUint(uint64(t),36)
	tokenKeySuffix := _merchantID
	tokenKey = tokenKeyPrefix + tokenKey + tokenKeySuffix + sessionTokenRequestStr //uid + timestamp + mid + sessionTokenRequestStr
	tokenKey = Encryption.PswEncryption(tokenKey) //sha256 encryption
	//xxteaTokenKey,err := Encryption.XxteaEncryption(_encryptionkey,tokenKey)
	if err != nil {
		return
	}
	//general value
	tokenValue := _uid + "@" + _merchantID + "#" + sessionTokenRequestStr

	redis ,err := Redis.NewAPSRedis(nil)
	if err != nil{
		return "",0,err
	}  else {
		redis.Connect()
		defer redis.Close()
		_ ,err := redis.Set(tokenKey,tokenValue,int64(tokenDuration))
		if err != nil {
			return "",0,err
		} else {
			return tokenKey,int64(tokenDuration), err
		}
	}
}

func DeleteSessionToken(sessionToken string) (error) {
	var returnErr error = nil
	redis ,err := Redis.NewAPSRedis(nil)
	returnErr = err
	if err == nil{
		redis.Connect()
		defer redis.Close()
		_,err := redis.Delete(sessionToken)
		returnErr = err
	}
	return returnErr
}

func RenewSessionToken(sessionToken string) (ttl int64,err error) {
	redis ,err := Redis.NewAPSRedis(nil)
	if err != nil{
		return 0,err
	} else {
		redis.Connect()
		defer redis.Close()
		ttl := int64(time.Now().Add(tokenDuration).Unix())
		_,err := redis.ExpireAt(sessionToken,ttl)
		return ttl, err
	}
}

func IsSessionTokenVailate(sessionToken string) (bool,error) {

	redis ,err := Redis.NewAPSRedis(nil)
	if err != nil{
		return false,err
	} else {
		redis.Connect()
		defer redis.Close()
		isExists,err := redis.IsExists(sessionToken)
		//res, err := c.IsExists(ctx, &redisPb.IsExistsRequest{Key:token})
		if err == nil &&  !isExists {
			err = Error.NewCustomError(SessionServiceError_SessionExpired,ErrorMap[SessionServiceError_SessionExpired])
		}
		return isExists,err
	}
}

