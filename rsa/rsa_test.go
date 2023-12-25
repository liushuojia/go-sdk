package rsa

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// 目录下 go test
func TestRedis(t *testing.T) {
	t.Log("rsa test")
	var (
		s   []byte
		sr  []byte
		err error
		str string
	)

	r := New()
	err = r.GenerateKey(4096)
	fmt.Println("===============================")
	fmt.Println(err)
	fmt.Println(string(r.GetPublicKey()))
	fmt.Println(string(r.GetPrivateKey()))

	data := []byte(
		`Go is expressive, concise, clean, and efficient. 
		Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, 
		while its novel type system enables flexible and modular program construction. 
		Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. 
		It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.
		Developing modules
		Developing and publishing modules
		You can collect related packages into modules, then publish the modules for other developers to use. This topic gives an overview of developing and publishing modules.
		
		Module release and versioning workflow
		When you develop modules for use by other developers, you can follow a workflow that helps ensure a reliable, consistent experience for developers using the module. This topic describes the high-level steps in that workflow.
		
		Managing module source
		When you're developing modules to publish for others to use, you can help ensure that your modules are easier for other developers to use by following the repository conventions described in this topic.
		
		Organizing a Go module
		What is the right way to organize the files and directories in a typical Go project? This topic discusses some common layouts depending on the kind of module you have.
		
		Developing a major version update
		A major version update can be very disruptive to your module's users because it includes breaking changes and represents a new module. Learn more in this topic.
		
		Publishing a module
		When you want to make a module available for other developers, you publish it so that it's visible to Go tools. Once you've published the module, developers importing its packages will be able to resolve a dependency on the module by running commands such as go get.
		
		Module version numbering
		A module's developer uses each part of a module's version number to signal the version’s stability and backward compatibility. For each new release, a module's release version number specifically reflects the nature of the module's changes since the preceding release.
		`,
	)

	fmt.Println("==============================================================")
	s, err = r.SignWithSha256(data)
	str = hex.EncodeToString(s)
	fmt.Println("私钥签名", err, str)

	s, _ = hex.DecodeString(str)
	fmt.Println("公钥验证", r.VerifySignWithSha256(data, s))

	fmt.Println("")
	fmt.Println("==============================================================")
	s, err = r.PublicEncrypt(data)
	str = hex.EncodeToString(s)
	fmt.Println("公钥加密", err, str)

	sr, err = r.PrivateDecrypt(s)
	fmt.Println("私钥解密", err, string(sr))

	fmt.Println("")
	fmt.Println("==============================================================")
	s, err = r.PrivateEncrypt(data)
	str = hex.EncodeToString(s)
	fmt.Println("私钥加密", err, str)

	sr, err = r.PublicDecrypt(s)
	fmt.Println("公钥解密", err, string(sr))

}
