package main

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"hash"
	"io"
	"log"
	"math/big"
	"os"
)

func main(){
    //HASH算法
	hex.EncodeToString(array)  //将数组转换为十六进制字符串
	hex.DecodeString(str)  //将十六进制字符串转换为数组

	var text = "12313213123"
	var hashInstance hash.Hash
	hashInstance = md5.New()  //md5
	hashInstance = sha1.New()  //sha1
	hashInstance = sha256.New()  //sha256
	hashInstance = sha512.New()  //sha512

	arr, _ := hex.DecodeString(text)
	hashInstance.Write(arr)
	hashInstance.Write([]byte(arr))

	cipherBytes := hashInstance.Sum(nil)
	fmt.Sprintf("%x",cipherBytes)

	//BASE64
	base64.StdEncoding.EncodeToString([]byte(str))
	base64.StdEncoding.DecodeString(str)


	str := []byte("锄禾日当午，汗滴禾下土")
	key := []byte("12345678")
	//对称加密算法
	//DES算法
	//三个参数：KEY：8个字节共64位的工作密钥   Data:加密或者解密数据   Mode： 加密或者解密的工作方式
	encryptDES(str,key)
	decryptDES(str,key)

	//AES算法
    //四个模式：ECB电子密码本  CBC密码分组链接  CFB加密反馈  OFB输出反馈
	encryptAES(str,key)
	decryptAES(str,key)

    //ECDSA椭圆算法
	encryptECDSA()

	//ECC椭圆算法
	encryptECC()

	//RSA算法
	RSA1()
	RSA2()
	encryptRSA(plainText,path)
	decryptRSA(cipherText,path)
}



//DES加密
func encryptDES(src, key []byte) []byte {
	//1 创建并返回一个使用DES算法的cipher.Block接口
	block, e := des.NewCipher(key)
	if e != nil {
		panic(e)
	}
	//2 对最后一个明文分组进行数据填充
	src = paddingText(src, block.BlockSize())
	//3 创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	iv := []byte("aaaabbbb")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//4 加密连续的数据块
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

//DES解密
func decryptDES(src, key []byte) []byte {
	//1 创建并返回一个使用DES算法的cipher.Block接口
	block, e := des.NewCipher(key)
	if e != nil {
		panic(e)
	}
	//2 创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	iv := []byte("aaaabbbb")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//3 数据块解密
	blockMode.CryptBlocks(src, src)
	//4 去掉最后一组的填充数据
	newText := unPaddingText(src)
	return newText
}

//AES加密
func encryptAES(src, key []byte) []byte {
	//1 创建并返回一个使用AES算法的cipher.Block接口
	block, e := aes.NewCipher(key)
	if e!=nil{
		panic(e)
	}
	//2 填充数据
	src = paddingText(src, block.BlockSize())
	//3 创建一个密码分组为链接模式的，底层使用AES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, key)
	//4 数据加密
	blockMode.CryptBlocks(src,src)
	return src
}

//AES解密
func decryptAES(src, key []byte) []byte {
	//1 创建并返回一个使用AES算法的cipher.Block接口
	block, e := aes.NewCipher(key)
	if e!=nil{
		panic(e)
	}
	//2 创建一个密码分组为链接模式的，底层使用AES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, key)
	//3 数据块解密
	blockMode.CryptBlocks(src,src)
	//4 去掉最后一组的填充数据
	unPaddingText(src)
	return src
}

//ECC椭圆算法
func encryptECC(){
	var Aprivate, Apublic [32]byte
	//产生随机数
	if _, err := io.ReadFull(rand.Reader, Aprivate[:]); err != nil {
		os.Exit(0)
	}
	curve25519.ScalarBaseMult(&Apublic, &Aprivate)
	fmt.Println("A私钥", base64.StdEncoding.EncodeToString(Aprivate[:]))
	fmt.Println("A公钥", base64.StdEncoding.EncodeToString(Apublic[:])) //作为椭圆起点

	var Bprivate, Bpublic [32]byte
	//产生随机数
	if _, err := io.ReadFull(rand.Reader, Bprivate[:]); err != nil {
		os.Exit(0)
	}
	curve25519.ScalarBaseMult(&Bpublic, &Bprivate)
	fmt.Println("B私钥",  base64.StdEncoding.EncodeToString(Bprivate[:]))
	fmt.Println("B公钥",  base64.StdEncoding.EncodeToString(Bpublic[:])) //作为椭圆起点

	var Akey, Bkey [32]byte
	//A的私钥加上Ｂ的公钥计算A的key
	curve25519.ScalarMult(&Akey, &Aprivate, &Bpublic)

	//B的私钥加上A的公钥计算B的key
	curve25519.ScalarMult(&Bkey, &Bprivate, &Apublic)

	fmt.Println("A交互的KEY",  base64.StdEncoding.EncodeToString(Akey[:]))
	fmt.Println("B交互的KEY",  base64.StdEncoding.EncodeToString(Bkey[:]))
}

//ECDSA椭圆算法
func encryptECDSA() {
	// 1、对需要签名的文件进行hash运算
	data := "from xiaoxiao to maomao 100 btc"
	hashInstance := sha256.New()
	hashInstance.Write([]byte(data))
	hashed := hashInstance.Sum(nil)
	// 2、生成私钥和公钥字节
	privateKey, publicKeyBytes := NewKeyPair()
	// 3、生成签名的der编码格式
	derSignString := ECDSASign(hashed, privateKey)
	fmt.Printf("签名信息为：%s\n", derSignString)
	// 4、验证签名
	flag := ECDSAVerify(publicKeyBytes, hashed, derSignString)
	fmt.Println("签名验证结果：", flag)
}

//RSA算法
func RSA1(){
	//生成私钥
	priv, e := rsa.GenerateKey(rand.Reader, 1024)
	if e != nil {
		fmt.Println(e)
	}

	//根据私钥产生公钥
	pub := &priv.PublicKey

	//明文
	plaintext := []byte("Hello world")

	//加密生成密文
	fmt.Printf("%q\n加密:\n", plaintext)
	ciphertext, e := rsa.EncryptOAEP(md5.New(), rand.Reader, pub, plaintext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%x\n", ciphertext)

	//解密得到明文
	fmt.Printf("解密:\n")
	plaintext, e = rsa.DecryptOAEP(md5.New(), rand.Reader, priv, ciphertext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%q\n", plaintext)

	//消息先进行Hash处理
	h := md5.New()
	h.Write(plaintext)
	hashed := h.Sum(nil)
	fmt.Printf("%q MD5 Hashed:\n\t%x\n", plaintext, hashed)

	//签名
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	sig, e := rsa.SignPSS(rand.Reader, priv, crypto.MD5, hashed, opts)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("签名:\n\t%x\n", sig)

	//认证
	fmt.Printf("验证结果:")
	if e := rsa.VerifyPSS(pub, crypto.MD5, hashed, sig, opts); e != nil {
		fmt.Println("失败:", e)
	} else {
		fmt.Println("成功.")
	}
}
func RSA2(){
	//生成密钥对，保存到文件
	GenerateRSAKey(2048)
	message:=[]byte("hello world")
	//加密
	cipherText:=encryptRSA(message,"public.pem")
	fmt.Println("加密后为：",string(cipherText))
	//解密
	plainText := decryptRSA(cipherText, "private.pem")
	fmt.Println("解密后为：",string(plainText))
}

//生成RSA私钥和公钥，保存到文件中
func GenerateRSAKey(bits int){
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err!=nil{
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err!=nil{
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock:= pem.Block{Type: "RSA Private Key",Bytes:X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile,&privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey:=privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey,err:=x509.MarshalPKIXPublicKey(&publicKey)
	if err!=nil{
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err!=nil{
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock:= pem.Block{Type: "RSA Public Key",Bytes:X509PublicKey}
	//保存到文件
	pem.Encode(publicFile,&publicBlock)
}

//RSA加密
func encryptRSA(plainText []byte,path string)[]byte{
	//打开文件
	file,err:=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	//类型断言
	publicKey:=publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err!=nil{
		panic(err)
	}
	//返回密文
	return cipherText
}

//RSA解密
func decryptRSA(cipherText []byte,path string) []byte{
	//打开文件
	file,err:=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	//对密文进行解密
	plainText,_:=rsa.DecryptPKCS1v15(rand.Reader,privateKey,cipherText)
	//返回明文
	return plainText
}

//填充最后一个分组的函数
func paddingText(src []byte, blockSize int) []byte {
	//src 原始数据
	//blockSize 每个分组的数据长度
	//1 求出最后一个分组要填充多少个字节
	padding := blockSize - len(src)%blockSize
	//2 创建新的切片，切片的字节数为padding，并初始化，每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	//3 将创建出的新切片和原始数据进行连接
	newText := append(src, padText...)
	//4 返回新的字符串
	return newText

}

//删除末尾填充的字符
func unPaddingText(src []byte) []byte {
	//1 求出需要处理的切片的长度
	len := len(src)
	//2  取出最后一个字符，得到其整型值
	number := int(src[len-1])
	//3 将切片末尾的number个字节删除
	newText := src[:len-number]
	return newText
}

// NewKeyPair 生成私钥和公钥，生成的私钥为结构体ecdsa.PrivateKey的指针
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	// 1、生成椭圆曲线对象
	curve := elliptic.P256()
	// 2、生成秘钥对，返回私钥对象（ecdsa.PrivateKey指针）
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	// 3、编码生成公钥字节数组，参数是椭圆曲线对象、x坐标、y坐标
	publicKeyBytes := elliptic.Marshal(curve, privateKey.PublicKey.X, privateKey.Y)
	fmt.Printf("私钥:%x\n", *privateKey)
	fmt.Printf("公钥:%x\n", publicKeyBytes)
	return *privateKey, publicKeyBytes
}

// ECDSASign ECDSA数字签名
func ECDSASign(hashed []byte, privateKey ecdsa.PrivateKey) string {
	// 1、数字签名生成r、s的big.Int对象，参数是随机数、私钥、签名文件的哈希串
	r, s, err := ecdsa.Sign(rand.Reader, &privateKey, hashed)
	if err != nil {
		return ""
	}
	fmt.Println("r结果：", r)
	fmt.Println("s结果：", s)
	// 2、将r、s转成r/s字符串
	strSignR := fmt.Sprintf("%x", r)
	strSignS := fmt.Sprintf("%x", s)
	if len(strSignR) == 63 {
		strSignR = "0" + strSignR
	}
	if len(strSignS) == 63 {
		strSignS = "0" + strSignS
	}
	fmt.Printf("r的16进制为：%s，长度为：%d\n", strSignR, len(strSignR))
	fmt.Printf("s的16进制为：%s，长度为：%d\n", strSignS, len(strSignS))
	// 3、r和s字符串拼接，形成数字签名的der格式
	derString := MakeDERSignString(strSignR, strSignS)
	return derString
}

// MakeDERSignString 生成数字签名的DER编码格式
func MakeDERSignString(strR, strS string) string {
	// 1、获取R和S的长度
	lenSignR := len(strR) / 2
	lenSignS := len(strS) / 2
	// 2、计算DER序列的总长度
	len := lenSignR + lenSignS + 4
	fmt.Printf("lenSignR为：%d，lenSignS为：%d，len为：%d\n", lenSignR, lenSignS, len)
	// 3、将10进制长度转16进制字符串
	strLenSignR := fmt.Sprintf("%x", int64(lenSignR))
	strLenSignS := fmt.Sprintf("%x", int64(lenSignS))
	strLen := fmt.Sprintf("%x", int64(len))
	fmt.Printf("strLenSignR为：%s，strLenSignS为：%s，strLen为：%s\n", strLenSignR, strLenSignS, strLen)
	// 4、拼接DER编码格式
	derString := "30" + strLen
	derString += "02" + strLenSignR + strR
	derString += "02" + strLenSignS + strS
	derString += "01"
	return derString
}

// ECDSAVerify ECDSA验证签名 （比特币系统中公钥具有0x04前缀）
func ECDSAVerify(publicKeyBytes, hashed []byte, derSignString string) bool {
	// 公钥长度
	keyLen := len(publicKeyBytes)
	if keyLen != 65 {
		return false
	}
	// 1、生成椭圆曲线对象
	curve := elliptic.P256()
	// 2、根据公钥字节数字，获取公钥中的x和y
	// 公钥字节中的前一半为x轴坐标，再将字节数组转成big.Int类型
	publicKeyBytes = publicKeyBytes[1:]
	// x := big.NewInt(0).SetBytes(publicKeyBytes[:32])
	x := new(big.Int).SetBytes(publicKeyBytes[:32])
	y := new(big.Int).SetBytes(publicKeyBytes[32:])
	// 3、生成公钥对象
	publicKey := ecdsa.PublicKey{Curve: curve, X: x, Y: y}
	// 4、对der格式的签名进行解析，获取r/s字节数组后转成big.Int类型
	rBytes, sBytes := ParseDERSignString(derSignString)
	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)
	return ecdsa.Verify(&publicKey, hashed, r, s)
}

// ParseDERSignString 对der格式的签名进行解析
func ParseDERSignString(derString string) (rBytes, sBytes []byte) {
	fmt.Println("derString：", derString)
	derBytes, _ := hex.DecodeString(derString)
	fmt.Println("derBytes", derBytes)
	rBytes = derBytes[4:36]
	sBytes = derBytes[len(derBytes)-33 : len(derBytes)-1]
	return
}

