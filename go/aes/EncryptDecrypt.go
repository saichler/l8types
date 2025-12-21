/*
© 2025 Sharon Aicler (saichler@gmail.com)

Layer 8 Ecosystem is licensed under the Apache License, Version 2.0.
You may obtain a copy of the License at:

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var keyChars = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateAES256Key() string {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic("failed to generate random key: " + err.Error())
	}
	
	keyRunes := make([]rune, 32)
	for i, b := range key {
		keyRunes[i] = keyChars[int(b)%len(keyChars)]
	}
	return string(keyRunes)
}

func Encrypt(dataToEncode []byte, key string) (string, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	dataLen := len(dataToEncode)
	cipherdata := make([]byte, aes.BlockSize+dataLen)

	iv := cipherdata[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherdata[aes.BlockSize:], dataToEncode)
	return base64.StdEncoding.EncodeToString(cipherdata), nil
}

func Decrypt(stringToDecode, key string) ([]byte, error) {
	encData, err := base64.StdEncoding.DecodeString(stringToDecode)
	if err != nil {
		return nil, err
	}
	if len(encData) < aes.BlockSize {
		return nil, errors.New("encrypted data does not have an iv spec")
	}
	
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	
	iv := encData[:aes.BlockSize]
	encData = encData[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	data := make([]byte, len(encData))
	cfb.XORKeyStream(data, encData)
	return data, nil
}
