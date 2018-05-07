package  main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
    "time"
    "reflect"
    "fmt"
    "strings"
    "database/sql"
    "crypto/sha256"
    "encoding/base64"
    _ "github.com/go-sql-driver/mysql"
    merchantServicePB "github.com/lampard1014/aphro/merchant/pb"
    encryptionServicePB "github.com/lampard1014/aphro/encryption/encryption-pb"
    sessionServicePB "github.com/lampard1014/aphro/session/pb"
    redisPb "github.com/lampard1014/aphro/redis/pb"
    "github.com/lampard1014/aphro/gateway/error"

    // "google.golang.org/grpc/status"
    // "google.golang.org/grpc/codes"

)

const (
	port  = ":10089"
    redisRPCAddress = "127.0.0.1:10101"
    sessionRPCAddress = "127.0.0.1:10088"
    encyptRPCAddress = "127.0.0.1:10087"
    mysqlDSN = "root:@tcp(127.0.0.1:3306)/iris_db"
)

type merchantService struct{}


func pswEncryption(psw string) (encryptionPsw string) {
    h := sha256.New()
    h.Write([]byte(psw))
    encryptionPsw = base64.URLEncoding.EncodeToString(h.Sum(nil))
    return encryptionPsw
}

func fetchSessionTokenValue(sessionToken string) (uid string, merchantID string, err error) {
    var returnErr error = nil
    sessionRPCConn,sessionRPCErr := grpc.Dial(sessionRPCAddress,grpc.WithInsecure())
    if sessionRPCErr == nil {
        sessionRPCConnClient := sessionServicePB.NewSessionServiceClient(sessionRPCConn)

        sessionRPCConnClientCtx, sessionRPCConnClientCancel := context.WithTimeout(context.Background(), time.Second)
        defer sessionRPCConnClientCancel()
        sessionTokenQueryResponse, sessionTokenQueryResponseErr := sessionRPCConnClient.QuerySessionToken(sessionRPCConnClientCtx, &sessionServicePB.SessionTokenQueryRequest{SessionToken:sessionToken})
        if sessionTokenQueryResponseErr == nil && sessionTokenQueryResponse != nil {
            sessionTokenValue := sessionTokenQueryResponse.SessionToken
            splitValue := strings.Split(sessionTokenValue, "#")
            uidAndMerchantID := strings.Split(splitValue[0],"@")
            uid = uidAndMerchantID[0]
            merchantID = uidAndMerchantID[1]
        } else {
            returnErr = AphroError.New(AphroError.BizError,"session 过期 请重新登录")
        }
    } else {
        returnErr = sessionRPCErr
    }
    defer sessionRPCConn.Close()
    return uid,merchantID,returnErr
}

func parseUsernameAndPsw(key string)(username string ,psw string, err error) {

    fmt.Println("parseUsernameAndPsw:",key)

    conn,encyptRPCErr := grpc.Dial(encyptRPCAddress,grpc.WithInsecure())

    var returnErr error = nil
    if encyptRPCErr == nil {
        c := encryptionServicePB.NewEncryptionServiceClient(conn)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        base64DecodeRes, base64DecodeErr := c.Base64Decode(ctx,&encryptionServicePB.EncryptionBase64DecodeRequest{DecodedStr:key})
        if base64DecodeErr == nil {
            rawData, RSADecryptionErr := c.RsaDecryption(ctx, &encryptionServicePB.DecryptionRSARequest{EncryptedStr:base64DecodeRes.RawValue})
            if RSADecryptionErr == nil {
                usernameAndPsw := string(rawData.RawValue)
                tmpSplit := strings.Split(usernameAndPsw,"@|@")
                if 2 == len(tmpSplit) {
                    username = tmpSplit[0]
                    psw = tmpSplit[1]
                } else {
                    returnErr = AphroError.New(AphroError.BizError,"拆分用户名密码错误")
                }
            } else {
               returnErr = RSADecryptionErr
            }
        } else {
            returnErr = base64DecodeErr
        }
    } else {
        returnErr = encyptRPCErr
    }
    defer conn.Close()
    return username,psw,returnErr
} 

func (s *merchantService) MerchantOpen(ctx context.Context, in *merchantServicePB.MerchantOpenRequest) (*merchantServicePB.MerchantOpenResponse, error) {

    merchantName := in.Name //select 1 from ... where 
    cellphone := in.Cellphone
    address := in.Address
    paymentBit := in.PaymentBit

    var returnErr error = nil

    db, dbOpenErr := sql.Open("mysql", mysqlDSN)
    defer db.Close()
    // Open doesn't open a connection. Validate DSN data:
    dbOpenErr = db.Ping()
    if (dbOpenErr == nil) {
        stmtIns, stmtInsErr := db.Prepare("INSERT INTO merchant (`merchant_name`,`merchant_address`,`payment_type`,`cellphone`) VALUES( ?, ?, ?, ?)") // ? = placeholder
        if stmtInsErr == nil {
            insertResult, insertErr := stmtIns.Exec(merchantName, address, paymentBit, cellphone) 
            if insertErr == nil {
                afftectedRow, afftectedRowErr := insertResult.RowsAffected()
                if afftectedRow != 1 || afftectedRowErr == nil {
                    returnErr = afftectedRowErr
                } else {

                }
            } else {
                returnErr = insertErr
            }
        } else {
            returnErr = stmtInsErr
        }
        defer stmtIns.Close()
    } else {
        returnErr = dbOpenErr        
    }

    return &merchantServicePB.MerchantOpenResponse{Successed:returnErr == nil},returnErr
}

func (s *merchantService) MerchantRegister(ctx context.Context, in *merchantServicePB.MerchantRegisterRequest) (*merchantServicePB.MerchantRegisterResponse, error) {
    cellphone, psw, err := parseUsernameAndPsw(in.Key)

    fmt.Println("user and psw err ",cellphone,psw, err)

    name := in.Name
    role := in.Role
    verifyCode := in.VerifyCode
    merchantID := in.MerchantID

    var operatorSuccess bool = false
    var returnErr error = err

    //step1 验证码检查
    conn,redisRPCError := grpc.Dial(redisRPCAddress,grpc.WithInsecure())
    if redisRPCError == nil {
        c := redisPb.NewRedisServiceClient(conn)
        ctxRPC, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel() 
        checkKey := "_verify_sms_"+ cellphone + "@1" 
        queryRes, queryRedisErr := c.Query(ctxRPC, &redisPb.QueryRedisRequest{Key: checkKey})

        hasError := queryRedisErr != nil

        if !hasError && queryRes!= nil && queryRes.Value ==  verifyCode {
            //插入数据库
            db, dbOpenErr := sql.Open("mysql", mysqlDSN)
            defer db.Close()
            dbOpenErr = db.Ping()
            if dbOpenErr == nil {
                //当注册是店长的时候 需要先自动注册一个店铺
                if role == 1 {
                    stmtMerhcantIns, stmtMerhcantInsErr := db.Prepare("INSERT INTO `merchant` (`merchant_name`,`merchant_address`,`payment_type`,`cellphone`)VALUES(?,?,?,?)") // ? = placeholder
                    if stmtMerhcantInsErr == nil {
                        stmtMerhcantInsertResult, stmtMerhcantInsertErr := stmtMerhcantIns.Exec("@", "@", 0, cellphone) 
                        if stmtMerhcantInsertErr == nil {
                            stmtMerhcantAfftectedRow, stmtMerhcantAfftectedRowErr := stmtMerhcantInsertResult.RowsAffected()
                            if stmtMerhcantAfftectedRow != 1 || stmtMerhcantAfftectedRowErr != nil {
                                returnErr = stmtMerhcantAfftectedRowErr
                                fmt.Println("[if]",stmtMerhcantAfftectedRowErr,"stmtMerhcantAfftectedRow",stmtMerhcantAfftectedRow)
                            } else {
                                fmt.Println("[insert merchant sucess!]")
                            }
                        } else {
                            returnErr = stmtMerhcantInsertErr
                        }
                    } else {
                        returnErr = stmtMerhcantInsErr
                    }
                    defer stmtMerhcantIns.Close()
                } 
                //正常的注册流程
                if returnErr == nil {
                    //成功新建商户 继续 新建操作员，todo 事务回滚
                    stmtIns, stmtInsErr := db.Prepare("INSERT INTO `merchant_account` (`name`,`cellphone`,`psw`,`role`,`merchant_id`)VALUES(?,?,?,?,?)") // ? = placeholder
                    if stmtInsErr == nil {
                        insertResult, insertErr := stmtIns.Exec(name, cellphone, pswEncryption(psw), role,merchantID) 
                        if insertErr == nil {
                            afftectedRow, afftectedRowErr := insertResult.RowsAffected()
                            if afftectedRow != 1 || afftectedRowErr != nil {
                                returnErr = afftectedRowErr
                            } else {
                                //成功
                                operatorSuccess = true
                            }
                        } else {
                            returnErr = insertErr
                        }
                    } else {
                        returnErr = stmtInsErr
                    }
                    defer stmtIns.Close()
                }
            } else {
                returnErr = dbOpenErr
            }
        } else {
            returnErr = AphroError.New(AphroError.BizError,"验证码验证错误")
        }
    } else {
        returnErr = redisRPCError
    }
    defer conn.Close()
    fmt.Println("success:",operatorSuccess,"returnErr:",returnErr ,"returnErrType:",reflect.TypeOf(returnErr))
    return &merchantServicePB.MerchantRegisterResponse{Successed:operatorSuccess},returnErr
}

func (s *merchantService) MerchantChangePsw(ctx context.Context, in *merchantServicePB.MerchantChangePswRequest) (*merchantServicePB.MerchantChangePswResponse, error) {

    newPsw := in.NewPsw
    sessionToken := in.SessionToken
    verifyCode := in.VerifyCode
    cellphone := in.Cellphone
    scene := uint32(2)
    var operatorSuccess bool = false
    var returnErr error = nil
    //检查session 是否合法

    sessionRPCConn,sessionRPCErr := grpc.Dial(sessionRPCAddress,grpc.WithInsecure())
    if sessionRPCErr == nil {
        sessionRPCConnClient := sessionServicePB.NewSessionServiceClient(sessionRPCConn)

        sessionRPCConnClientCtx, sessionRPCConnClientCancel := context.WithTimeout(context.Background(), time.Second)

        defer sessionRPCConnClientCancel()
        var isSessionTokenOK bool = false
        isSessionTokenVailate, isSessionTokenVailateErr := sessionRPCConnClient.IsSessionTokenVailate(sessionRPCConnClientCtx, &sessionServicePB.IsSessionTokenVailateRequest{SessionToken:sessionToken})
        if isSessionTokenVailateErr == nil {
            isSessionTokenOK = isSessionTokenVailate.Successed

            if isSessionTokenOK {
                encyptRPCConn,encyptRPCErr := grpc.Dial(encyptRPCAddress,grpc.WithInsecure())
                if encyptRPCErr == nil {
                    encyptRPCClient := encryptionServicePB.NewEncryptionServiceClient(encyptRPCConn)
                    encyptRPCClientCtx, encyptRPCClientCancel := context.WithTimeout(context.Background(), time.Second)
                    defer encyptRPCClientCancel()
                    base64DecodeRes, base64DecodeErr := encyptRPCClient.Base64Decode(encyptRPCClientCtx,&encryptionServicePB.EncryptionBase64DecodeRequest{DecodedStr:newPsw})
                    if base64DecodeErr == nil {
                        rawData, RSADecryptionErr := encyptRPCClient.RsaDecryption(encyptRPCClientCtx, &encryptionServicePB.DecryptionRSARequest{EncryptedStr:base64DecodeRes.RawValue})
                        if RSADecryptionErr == nil {
                            newPsw := string(rawData.RawValue)
                            if newPsw != "" {
                                    //获取到最新的密码 验证 验证码是否正常
                                    verifyCodeRes ,verifyCodeResErr := sessionRPCConnClient.MerchantVerifyCode(sessionRPCConnClientCtx,&sessionServicePB.MerchantVerifyCodeRequest{Cellphone:cellphone,Scene:scene,SmsCode:verifyCode})
                                    if verifyCodeRes.Successed {
                                        db, dbOpenErr := sql.Open("mysql", mysqlDSN)
                                        defer db.Close()
                                        dbOpenErr = db.Ping()
                                        if dbOpenErr == nil {
                                            stmtIns, stmtInsErr := db.Prepare("update `merchant_account` set psw = ? where cellphone = ? limit 1") // ? = placeholder
                                            if stmtInsErr == nil {
                                                updateResult, updateErr := stmtIns.Exec(pswEncryption(newPsw),cellphone) 
                                                if updateErr == nil {
                                                    afftectedRow, afftectedRowErr := updateResult.RowsAffected()
                                                    if afftectedRow != 1 || afftectedRowErr != nil {
                                                        returnErr = afftectedRowErr
                                                    } else {
                                                        //成功
                                                        operatorSuccess = true
                                                    }
                                                } else {
                                                    returnErr = updateErr
                                                }
                                            } else {

                                                returnErr = stmtInsErr
                                            }
                                            defer stmtIns.Close()
                                        } else {
                                            returnErr = dbOpenErr
                                        }   
                                    } else {
                                        returnErr = verifyCodeResErr
                                    }
                            } else {
                                returnErr = AphroError.New(AphroError.BizError,"解析密码错误")
                            }
                        } else {
                           returnErr = RSADecryptionErr
                        }
                    } else {
                        returnErr = base64DecodeErr
                    }
                } else {
                    returnErr = encyptRPCErr
                }
                defer encyptRPCConn.Close()

            } else {
                returnErr = AphroError.New(AphroError.BizError,"session 过期 请重新登录")
            }
        } else {
            returnErr = isSessionTokenVailateErr
        }
    } else {
        returnErr = sessionRPCErr
    }
    defer sessionRPCConn.Close()

    return &merchantServicePB.MerchantChangePswResponse{Successed:operatorSuccess},returnErr
}

func (s *merchantService) MerchantLogin(ctx context.Context, in *merchantServicePB.MerchantLoginRequest) (*merchantServicePB.MerchantLoginResponse, error) {
    cellphone, psw, err:= parseUsernameAndPsw(in.Key)
    tokenRequest := in.TokenRequest
    merchantID := in.MerchantID
    var returnErr error = err
    var sessionToken string = "";
    //查询数据库
    var (
        uid uint32
        name string
        role int
    )
    db, dbOpenErr := sql.Open("mysql", mysqlDSN)
    defer db.Close()
    dbOpenErr = db.Ping()
    if dbOpenErr == nil {
        queryRowErr := db.QueryRow("SELECT id as uid, name ,role,merchant_id FROM merchant_account WHERE cellphone = ? AND psw = ? AND merchantID = ? LIMIT 1", cellphone, psw,merchantID).Scan(&uid,&name,&role,&merchantID)
        if queryRowErr == nil {
            //制作 令牌
            sessionRPCConn,sessionRPCErr := grpc.Dial(sessionRPCAddress,grpc.WithInsecure())
            if sessionRPCErr == nil {
                c := sessionServicePB.NewSessionServiceClient(sessionRPCConn)

                ctx, cancel := context.WithTimeout(context.Background(), time.Second)

                defer cancel()

                createSessionTokenRes, err := c.CreateSessionToken(ctx, &sessionServicePB.SessionTokenCreateRequest{SessionTokenRequestStr:tokenRequest,Uid:uid,MerchantID:merchantID})
                if err == nil {
                    sessionToken = createSessionTokenRes.SessionToken

                } else {
                    returnErr = err
                }
            } else {
                returnErr = sessionRPCErr
            }
            defer sessionRPCConn.Close()
        } else if queryRowErr == sql.ErrNoRows {
            //没有记录
            returnErr = AphroError.New(AphroError.BizError,"密码或者用户名错误")
        } else {
            returnErr = queryRowErr
        }

    } else {
        returnErr = dbOpenErr
    }
    return &merchantServicePB.MerchantLoginResponse{SessionToken:sessionToken,Successed:returnErr == nil},returnErr
}

func (s *merchantService) MerchantInfo(ctx context.Context, in *merchantServicePB.MerchantInfoRequest) (*merchantServicePB.MerchantInfoResponse, error) {

    //验证token合法性
    uid, merchantID, sessionTokenError := fetchSessionTokenValue(in.Token)
    var returnErr error = nil

    var (
        merchantName string
        name string
        role int
    )
    if sessionTokenError == nil {
  
        db, dbOpenErr := sql.Open("mysql", mysqlDSN)
        defer db.Close()
        dbOpenErr = db.Ping()
        if dbOpenErr == nil {
            queryRowErr := db.QueryRow("SELECT m.merchant_name,ma.role,ma.name FROM merchant_account AS ma, merchant AS m WHERE ma.merchant_id = ? and ma.ID = ? limit 1;", merchantID, uid).Scan(&merchantName,&role,&name)
            if queryRowErr == nil {
                //制作 令牌
            } else if queryRowErr == sql.ErrNoRows {
                //没有记录
               returnErr = AphroError.New(AphroError.BizError,"没用商户信息")
            } else {
                returnErr = queryRowErr
            }

        } else {
            returnErr = dbOpenErr
        }
    } else {
        returnErr = sessionTokenError
    }

    return &merchantServicePB.MerchantInfoResponse{MerchantName:merchantName,MerchantAccount:&merchantServicePB.InnerMerchantAccount{Role:uint32(role),Name:name}},returnErr
}


func (s *merchantService) MerchantRoomInfo(ctx context.Context, in *merchantServicePB.MerchantRoomInfoRequest) (*merchantServicePB.MerchantRoomInfoResponse, error) {
    return nil,nil
}

func deferFunc() {
    if err := recover(); err != nil {
        fmt.Println("error happend:")
        fmt.Println(err)
    }
}

// // auth 验证Token
// func auth(ctx context.Context) error {
//     md, ok := metadata.FromContext(ctx)
//     if !ok {
//         return grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
//     }

//     var (
//         appid  string
//         appkey string
//     )

//     if val, ok := md["appid"]; ok {
//         appid = val[0]
//     }

//     if val, ok := md["appkey"]; ok {
//         appkey = val[0]
//     }

//     if appid != "101010" || appkey != "i am key" {
//         return grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
//     }

//     return nil
// }


func main() {
    defer deferFunc() 
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatal(err)
    }

    // var opts []grpc.ServerOption //签名 和 验签

    // // 注册interceptor
    // var interceptor grpc.UnaryServerInterceptor
    // interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
    //     err = auth(ctx)
    //     if err != nil {
    //         return
    //     }
    //     // 继续处理请求
    //     return handler(ctx, req)
    // }
    // opts = append(opts, grpc.UnaryInterceptor(interceptor))


    s := grpc.NewServer()//opts...)
    merchantServicePB.RegisterMerchantServiceServer(s, new(merchantService))
    err = s.Serve(lis)
    if err != nil {
        log.Fatal(err)
    }
}
