package jwtauth

import ( // นำเข้าฟังก์ชันและแพ็กเกจที่จำเป็น
	"context" // ใช้สำหรับจัดการ Context ในการเรียกใช้งาน API
	"errors"  // ใช้จัดการข้อผิดพลาด
	"math"    // ใช้ฟังก์ชันทางคณิตศาสตร์ เช่น pow
	"sync"    // ใช้สำหรับควบคุมการทำงานแบบ thread-safe
	"time"    // ใช้จัดการเวลาและวันที่

	"github.com/golang-jwt/jwt/v5"    // แพ็กเกจสำหรับจัดการ JWT
	"google.golang.org/grpc/metadata" // ใช้ metadata สำหรับ gRPC
)

type ( // กำหนดประเภทข้อมูล (struct) และ Interface ที่จะใช้
	AuthFactory interface { // Interface ที่บอกว่า Token ต้องสามารถเซ็นได้
		SignToken() string // ฟังก์ชันเซ็น Token
	}
	Claims struct { // โครงสร้างข้อมูล Claims สำหรับ Token
		Id       string `json:"id"`        // ID ผู้ใช้
		Rolecode string `json:"role_code"` // สิทธิ์หรือบทบาทผู้ใช้
	}
	AuthMapClaims struct { // ข้อมูล Claims ที่ผสมระหว่างข้อมูลส่วนตัวและข้อมูลทั่วไปของ JWT
		*Claims              // ข้อมูลผู้ใช้
		jwt.RegisteredClaims // ข้อมูลมาตรฐานของ JWT เช่น เวลาออก, วันหมดอายุ
	}
	authConcrete struct { // โครงสร้างพื้นฐานสำหรับ Token แต่ละประเภท
		SecretKey []byte         // SecretKey สำหรับเซ็น Token
		Claims    *AuthMapClaims `json:"claims"` // Claims ที่เกี่ยวข้อง
	}
	accessToken  struct{ *authConcrete } // ประเภท Access Token
	refreshToken struct{ *authConcrete } // ประเภท Refresh Token
	apiKey       struct{ *authConcrete } // ประเภท API Key
)

func (a *authConcrete) SignToken() string { // ฟังก์ชันสำหรับเซ็น Token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, a.Claims) // สร้าง Token ด้วยวิธีเซ็น ES256 และข้อมูล Claims
	ss, _ := token.SignedString(a.SecretKey)                     // เซ็น Token ด้วย SecretKey
	return ss                                                    // คืนค่า Token เป็น String
}

func now() time.Time { // ฟังก์ชันคืนค่าปัจจุบันในเขตเวลา Asia/Bangkok
	loc, _ := time.LoadLocation("Asia/Bangkok") // โหลดเขตเวลา Asia/Bangkok
	return time.Now().In(loc)                   // คืนค่าเวลาปัจจุบันในเขตที่กำหนด
}

func jwtTimeDurationCal(t int64) *jwt.NumericDate { // คำนวณระยะเวลาในหน่วยนาโนวินาที
	return jwt.NewNumericDate(now().Add(time.Duration(t * int64(math.Pow10(9))))) // คืนค่าเวลาที่เพิ่มขึ้นตาม `t`
}

func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate { // แปลง Unix Timestamp เป็น NumericDate
	return jwt.NewNumericDate(time.Unix(t, 0)) // คืนค่าจาก Unix Timestamp
}

func NewAccessToken(secret string, expriedAt int64, claims *Claims) AuthFactory { // สร้าง Access Token
	return &accessToken{
		authConcrete: &authConcrete{
			SecretKey: []byte(secret), // เก็บ SecretKey
			Claims: &AuthMapClaims{
				Claims: claims, // ข้อมูลผู้ใช้
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "Shop-Game",                   // ใครเป็นคนออก Token
					Subject:   "access-token",                // ประเภทของ Token
					ExpiresAt: jwtTimeDurationCal(expriedAt), // วันหมดอายุ
					NotBefore: jwt.NewNumericDate(now()),     // วันเริ่มใช้ Token
					IssuedAt:  jwt.NewNumericDate(now()),     // วันออก Token
				},
			},
		},
	}
}

func NewRefreshToken(secret string, expriedAt int64, claims *Claims) AuthFactory { // สร้าง Refresh Token
	return &accessToken{ // โครงสร้างเหมือน Access Token แต่ Subject ต่างกัน
		authConcrete: &authConcrete{
			SecretKey: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "Shop-Game",
					Subject:   "refresh-token",
					ExpiresAt: jwtTimeDurationCal(expriedAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ReloadToken(secret string, expiredAt int64, claims *Claims) string { // สร้าง Refresh Token ใหม่
	obj := &refreshToken{ // กำหนดข้อมูล Refresh Token ใหม่
		authConcrete: &authConcrete{
			SecretKey: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "Shop-Game",
					Subject:   "refresh-token",
					ExpiresAt: jwtTimeRepeatAdapter(expiredAt), // ใช้วันหมดอายุที่ระบุ
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
	return obj.SignToken() // คืนค่า Token ใหม่ที่เซ็นแล้ว
}

func NewApiKey(secret string) AuthFactory { // สร้าง API Key
	return &apiKey{
		authConcrete: &authConcrete{
			SecretKey: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: &Claims{}, // ไม่ต้องมีข้อมูลผู้ใช้ใน API Key
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "Shop-Game",
					Subject:   "refresh-token",
					ExpiresAt: jwtTimeDurationCal(31560000), // หมดอายุใน 1 ปี
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ParseToken(secret string, tokenString string) (*AuthMapClaims, error) { // ตรวจสอบและแปลง Token
	token, err := jwt.ParseWithClaims(tokenString, &AuthMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok { // ตรวจสอบ Signing Method
			return nil, errors.New("error: unexpected signing method")
		}
		return []byte(secret), nil // คืนค่า SecretKey สำหรับตรวจสอบ
	})

	if err != nil { // จัดการข้อผิดพลาด
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("error: token is expired")
		} else {
			return nil, errors.New("error: token is invalid")
		}
	}

	if claims, ok := token.Claims.(*AuthMapClaims); ok { // แปลง Claims หากสำเร็จ
		return claims, nil
	}

	return nil, errors.New("error: claims type is invalid") // คืนข้อผิดพลาดหากแปลงไม่สำเร็จ
}

var apiKeyInstant string // เก็บ API Key ที่สร้างแล้ว
var once sync.Once       // ใช้เพื่อให้ฟังก์ชันทำงานเพียงครั้งเดียว (Singleton)

func SetApiKey(secret string) { // สร้าง API Key แบบครั้งเดียว
	once.Do(func() {
		apiKeyInstant = NewApiKey(secret).SignToken() // สร้าง API Key และเก็บในตัวแปร global
	})
}

func SetApiKeyInContext(pctx *context.Context) { // เพิ่ม API Key ลงใน metadata ของ Context
	*pctx = metadata.NewOutgoingContext(*pctx, metadata.Pairs("auth", apiKeyInstant))
}
