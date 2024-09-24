package services_test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/edipermadi/softhsm/pkg/token/messages"
	"github.com/edipermadi/softhsm/pkg/token/services"
	"github.com/edipermadi/softhsm/pkg/token/users"
	"github.com/stretchr/testify/require"
)

func TestService_Digest(t *testing.T) {
	type testVector struct {
		Message string
		Digest  string
	}
	type testCase struct {
		Title       string
		Mechanism   messages.Mechanism
		TestVectors []testVector
	}

	testCases := []testCase{
		{
			Title: "SHA1",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA_1,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "a9993e36 4706816a ba3e2571 7850c26c 9cd0d89d",
				},
				{
					Message: "",
					Digest:  "da39a3ee 5e6b4b0d 3255bfef 95601890 afd80709",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "84983e44 1c3bd26e baae4aa1 f95129e5 e54670f1",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "a49b2446 a02c645b f419f995 b6709125 3a04a259",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "34aa973c d4c4daa4 f61eeb2b dbad2731 6534016f",
				},
			},
		},
		{
			Title: "SHA224",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA224,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "23097d22 3405d822 8642a477 bda255b3 2aadbce4 bda0b3f7 e36c9da7",
				},
				{
					Message: "",
					Digest:  "d14a028c 2a3a2bc9 476102bb 288234c4 15a2b01f 828ea62a c5b3e42f",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "75388b16 512776cc 5dba5da1 fd890150 b0c6455c b4f58b19 52522525",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "c97ca9a5 59850ce9 7a04a96d ef6d99a9 e0e0e2ab 14e6b8df 265fc0b3",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "20794655 980c91d8 bbb4c1ea 97618a4b f03f4258 1948b2ee 4ee7ad67",
				},
			},
		},
		{
			Title: "SHA256",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA256,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "ba7816bf 8f01cfea 414140de 5dae2223 b00361a3 96177a9c b410ff61 f20015ad",
				},
				{
					Message: "",
					Digest:  "e3b0c442 98fc1c14 9afbf4c8 996fb924 27ae41e4 649b934c a495991b 7852b855",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "248d6a61 d20638b8 e5c02693 0c3e6039 a33ce459 64ff2167 f6ecedd4 19db06c1",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "cf5b16a7 78af8380 036ce59e 7b049237 0b249b11 e8f07a51 afac4503 7afee9d1",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "cdc76e5c 9914fb92 81a1c7e2 84d73e67 f1809a48 a497200e 046d39cc c7112cd0",
				},
			},
		},
		{
			Title: "SHA384",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA384,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "cb00753f45a35e8b b5a03d699ac65007 272c32ab0eded163 1a8b605a43ff5bed 8086072ba1e7cc23 58baeca134c825a7",
				},
				{
					Message: "",
					Digest:  "38b060a751ac9638 4cd9327eb1b1e36a 21fdb71114be0743 4c0cc7bf63f6e1da 274edebfe76f65fb d51ad2f14898b95b",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "3391fdddfc8dc739 3707a65b1b470939 7cf8b1d162af05ab fe8f450de5f36bc6 b0455a8520bc4e6f 5fe95b1fe3c8452b",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "09330c33f71147e8 3d192fc782cd1b47 53111b173b3b05d2 2fa08086e3b0f712 fcc7c71a557e2db9 66c3e9fa91746039",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "9d0e1809716474cb 086e834e310a4a1c ed149e9c00f24852 7972cec5704c2a5b 07b8b3dc38ecc4eb ae97ddd87f3d8985",
				},
			},
		},
		{
			Title: "SHA512",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA512,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "ddaf35a193617aba cc417349ae204131 12e6fa4e89a97ea2 0a9eeee64b55d39a 2192992a274fc1a8 36ba3c23a3feebbd 454d4423643ce80e 2a9ac94fa54ca49f",
				},
				{
					Message: "",
					Digest:  "cf83e1357eefb8bd f1542850d66d8007 d620e4050b5715dc 83f4a921d36ce9ce 47d0d13c5d85f2b0 ff8318d2877eec2f 63b931bd47417a81 a538327af927da3e",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "204a8fc6dda82f0a 0ced7beb8e08a416 57c16ef468b228a8 279be331a703c335 96fd15c13b1b07f9 aa1d3bea57789ca0 31ad85c7a71dd703 54ec631238ca3445",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "8e959b75dae313da 8cf4f72814fc143f 8f7779c6eb9f7fa1 7299aeadb6889018 501d289e4900f7e4 331b99dec4b5433a c7d329eeb6dd2654 5e96e55b874be909",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "e718483d0ce76964 4e2e42c7bc15b463 8e1f98b13b204428 5632a803afa973eb de0ff244877ea60a 4cb0432ce577c31b eb009c5c2c49aa2e 4eadb217ad8cc09b",
				},
			},
		},
		{
			Title: "SHA3_224",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA3_224,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "e642824c3f8cf24a d09234ee7d3c766f c9a3a5168d0c94ad 73b46fdf",
				},
				{
					Message: "",
					Digest:  "6b4e03423667dbb7 3b6e15454f0eb1ab d4597f9a1b078e3f 5b5a6bc7",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "8a24108b154ada21 c9fd5574494479ba 5c7e7ab76ef264ea d0fcce33",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "543e6868e1666c1a 643630df77367ae5 a62a85070a51c14c bf665cbc",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "d69335b93325192e 516a912e6d19a15c b51c6ed5c15243e7 a7fd653c",
				},
			},
		},
		{
			Title: "SHA3_256",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA3_256,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "3a985da74fe225b2 045c172d6bd390bd 855f086e3e9d525b 46bfe24511431532",
				},
				{
					Message: "",
					Digest:  "a7ffc6f8bf1ed766 51c14756a061d662 f580ff4de43b49fa 82d80a4b80f8434a",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "41c0dba2a9d62408 49100376a8235e2c 82e1b9998a999e21 db32dd97496d3376",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "916f6061fe879741 ca6469b43971dfdb 28b1a32dc36cb325 4e812be27aad1d18",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "5c8875ae474a3634 ba4fd55ec85bffd6 61f32aca75c6d699 d0cdcb6c115891c1",
				},
			},
		},
		{
			Title: "SHA3_384",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA3_384,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "ec01498288516fc9 26459f58e2c6ad8d f9b473cb0fc08c25 96da7cf0e49be4b2 98d88cea927ac7f5 39f1edf228376d25",
				},
				{
					Message: "",
					Digest:  "0c63a75b845e4f7d 01107d852e4c2485 c51a50aaaa94fc61 995e71bbee983a2a c3713831264adb47 fb6bd1e058d5f004",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "991c665755eb3a4b 6bbdfb75c78a492e 8c56a22c5c4d7e42 9bfdbc32b9d4ad5a a04a1f076e62fea1 9eef51acd0657c22",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "79407d3b5916b59c 3e30b09822974791 c313fb9ecc849e40 6f23592d04f625dc 8c709b98b43b3852 b337216179aa7fc7",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "eee9e24d78c18553 37983451df97c8ad 9eedf256c6334f8e 948d252d5e0e7684 7aa0774ddb90a842 190d2c558b4b8340",
				},
			},
		},
		{
			Title: "SHA3_512",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA3_512,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "b751850b1a57168a 5693cd924b6b096e 08f621827444f70d 884f5d0240d2712e 10e116e9192af3c9 1a7ec57647e39340 57340b4cf408d5a5 6592f8274eec53f0",
				},
				{
					Message: "",
					Digest:  "a69f73cca23a9ac5 c8b567dc185a756e 97c982164fe25859 e0d1dcc1475c80a6 15b2123af1f5f94c 11e3e9402c3ac558 f500199d95b6d3e3 01758586281dcd26",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "04a371e84ecfb5b8 b77cb48610fca818 2dd457ce6f326a0f d3d7ec2f1e91636d ee691fbe0c985302 ba1b0d8dc78c0863 46b533b49c030d99 a27daf1139d6e75e",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "afebb2ef542e6579 c50cad06d2e578f9 f8dd6881d7dc824d 26360feebf18a4fa 73e3261122948efc fd492e74e82e2189 ed0fb440d187f382 270cb455f21dd185",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "3c3a876da14034ab 60627c077bb98f7e 120a2a5370212dff b3385a18d4f38859 ed311d0a9d5141ce 9cc5c66ee689b266 a8aa18ace8282a0e 0db596c90b0a7b87",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			for _, tv := range tc.TestVectors {
				userService, err := users.NewService()
				require.NoError(t, err)

				svc := services.NewService(userService)
				openSessionResponse, err := svc.OpenSession(&messages.OpenSessionRequest{Flags: 0})
				require.NoError(t, err)
				sessionID := openSessionResponse.SessionID

				loginResponse, err := svc.Login(&messages.LoginRequest{
					SessionID: sessionID,
					UserType:  users.TypeUser,
					Pin:       []byte("password"),
				})
				require.NoError(t, err)
				require.Equal(t, messages.ReturnValue_OK, loginResponse.ReturnValue)

				digestInitResponse, err := svc.DigestInit(&messages.DigestInitRequest{
					SessionID: sessionID,
					Mechanism: tc.Mechanism,
				})
				require.NoError(t, err)
				require.Equal(t, messages.ReturnValue_OK, digestInitResponse.ReturnValue)

				digestResponse, err := svc.Digest(&messages.DigestRequest{
					SessionID: sessionID,
					Data:      []byte(tv.Message),
				})
				require.NoError(t, err)
				require.Equal(t, messages.ReturnValue_OK, digestResponse.ReturnValue)
				require.Equal(t, strings.ReplaceAll(tv.Digest, " ", ""), hex.EncodeToString(digestResponse.Digest))
			}
		})
	}
}

func TestService_DigestUpdate(t *testing.T) {
	type testVector struct {
		Message string
		Digest  string
	}
	type testCase struct {
		Title       string
		Mechanism   messages.Mechanism
		TestVectors []testVector
	}

	testCases := []testCase{
		{
			Title: "SHA1",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA_1,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "a9993e36 4706816a ba3e2571 7850c26c 9cd0d89d",
				},
				{
					Message: "",
					Digest:  "da39a3ee 5e6b4b0d 3255bfef 95601890 afd80709",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "84983e44 1c3bd26e baae4aa1 f95129e5 e54670f1",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "a49b2446 a02c645b f419f995 b6709125 3a04a259",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "34aa973c d4c4daa4 f61eeb2b dbad2731 6534016f",
				},
			},
		},
		{
			Title: "SHA224",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA224,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "23097d22 3405d822 8642a477 bda255b3 2aadbce4 bda0b3f7 e36c9da7",
				},
				{
					Message: "",
					Digest:  "d14a028c 2a3a2bc9 476102bb 288234c4 15a2b01f 828ea62a c5b3e42f",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "75388b16 512776cc 5dba5da1 fd890150 b0c6455c b4f58b19 52522525",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "c97ca9a5 59850ce9 7a04a96d ef6d99a9 e0e0e2ab 14e6b8df 265fc0b3",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "20794655 980c91d8 bbb4c1ea 97618a4b f03f4258 1948b2ee 4ee7ad67",
				},
			},
		},
		{
			Title: "SHA256",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA256,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "ba7816bf 8f01cfea 414140de 5dae2223 b00361a3 96177a9c b410ff61 f20015ad",
				},
				{
					Message: "",
					Digest:  "e3b0c442 98fc1c14 9afbf4c8 996fb924 27ae41e4 649b934c a495991b 7852b855",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "248d6a61 d20638b8 e5c02693 0c3e6039 a33ce459 64ff2167 f6ecedd4 19db06c1",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "cf5b16a7 78af8380 036ce59e 7b049237 0b249b11 e8f07a51 afac4503 7afee9d1",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "cdc76e5c 9914fb92 81a1c7e2 84d73e67 f1809a48 a497200e 046d39cc c7112cd0",
				},
			},
		},
		{
			Title: "SHA384",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA384,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "cb00753f45a35e8b b5a03d699ac65007 272c32ab0eded163 1a8b605a43ff5bed 8086072ba1e7cc23 58baeca134c825a7",
				},
				{
					Message: "",
					Digest:  "38b060a751ac9638 4cd9327eb1b1e36a 21fdb71114be0743 4c0cc7bf63f6e1da 274edebfe76f65fb d51ad2f14898b95b",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "3391fdddfc8dc739 3707a65b1b470939 7cf8b1d162af05ab fe8f450de5f36bc6 b0455a8520bc4e6f 5fe95b1fe3c8452b",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "09330c33f71147e8 3d192fc782cd1b47 53111b173b3b05d2 2fa08086e3b0f712 fcc7c71a557e2db9 66c3e9fa91746039",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "9d0e1809716474cb 086e834e310a4a1c ed149e9c00f24852 7972cec5704c2a5b 07b8b3dc38ecc4eb ae97ddd87f3d8985",
				},
			},
		},
		{
			Title: "SHA512",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA512,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "ddaf35a193617aba cc417349ae204131 12e6fa4e89a97ea2 0a9eeee64b55d39a 2192992a274fc1a8 36ba3c23a3feebbd 454d4423643ce80e 2a9ac94fa54ca49f",
				},
				{
					Message: "",
					Digest:  "cf83e1357eefb8bd f1542850d66d8007 d620e4050b5715dc 83f4a921d36ce9ce 47d0d13c5d85f2b0 ff8318d2877eec2f 63b931bd47417a81 a538327af927da3e",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "204a8fc6dda82f0a 0ced7beb8e08a416 57c16ef468b228a8 279be331a703c335 96fd15c13b1b07f9 aa1d3bea57789ca0 31ad85c7a71dd703 54ec631238ca3445",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "8e959b75dae313da 8cf4f72814fc143f 8f7779c6eb9f7fa1 7299aeadb6889018 501d289e4900f7e4 331b99dec4b5433a c7d329eeb6dd2654 5e96e55b874be909",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "e718483d0ce76964 4e2e42c7bc15b463 8e1f98b13b204428 5632a803afa973eb de0ff244877ea60a 4cb0432ce577c31b eb009c5c2c49aa2e 4eadb217ad8cc09b",
				},
			},
		},
		{
			Title: "SHA3_224",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA3_224,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "e642824c3f8cf24a d09234ee7d3c766f c9a3a5168d0c94ad 73b46fdf",
				},
				{
					Message: "",
					Digest:  "6b4e03423667dbb7 3b6e15454f0eb1ab d4597f9a1b078e3f 5b5a6bc7",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "8a24108b154ada21 c9fd5574494479ba 5c7e7ab76ef264ea d0fcce33",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "543e6868e1666c1a 643630df77367ae5 a62a85070a51c14c bf665cbc",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "d69335b93325192e 516a912e6d19a15c b51c6ed5c15243e7 a7fd653c",
				},
			},
		},
		{
			Title: "SHA3_256",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA3_256,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "3a985da74fe225b2 045c172d6bd390bd 855f086e3e9d525b 46bfe24511431532",
				},
				{
					Message: "",
					Digest:  "a7ffc6f8bf1ed766 51c14756a061d662 f580ff4de43b49fa 82d80a4b80f8434a",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "41c0dba2a9d62408 49100376a8235e2c 82e1b9998a999e21 db32dd97496d3376",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "916f6061fe879741 ca6469b43971dfdb 28b1a32dc36cb325 4e812be27aad1d18",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "5c8875ae474a3634 ba4fd55ec85bffd6 61f32aca75c6d699 d0cdcb6c115891c1",
				},
			},
		},
		{
			Title: "SHA3_384",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA3_384,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "ec01498288516fc9 26459f58e2c6ad8d f9b473cb0fc08c25 96da7cf0e49be4b2 98d88cea927ac7f5 39f1edf228376d25",
				},
				{
					Message: "",
					Digest:  "0c63a75b845e4f7d 01107d852e4c2485 c51a50aaaa94fc61 995e71bbee983a2a c3713831264adb47 fb6bd1e058d5f004",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "991c665755eb3a4b 6bbdfb75c78a492e 8c56a22c5c4d7e42 9bfdbc32b9d4ad5a a04a1f076e62fea1 9eef51acd0657c22",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "79407d3b5916b59c 3e30b09822974791 c313fb9ecc849e40 6f23592d04f625dc 8c709b98b43b3852 b337216179aa7fc7",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "eee9e24d78c18553 37983451df97c8ad 9eedf256c6334f8e 948d252d5e0e7684 7aa0774ddb90a842 190d2c558b4b8340",
				},
			},
		},
		{
			Title: "SHA3_512",
			Mechanism: messages.Mechanism{
				Mechanism: messages.MechanismType_SHA3_512,
				Parameter: messages.MechanismParameter{},
			},
			TestVectors: []testVector{
				{
					Message: "abc",
					Digest:  "b751850b1a57168a 5693cd924b6b096e 08f621827444f70d 884f5d0240d2712e 10e116e9192af3c9 1a7ec57647e39340 57340b4cf408d5a5 6592f8274eec53f0",
				},
				{
					Message: "",
					Digest:  "a69f73cca23a9ac5 c8b567dc185a756e 97c982164fe25859 e0d1dcc1475c80a6 15b2123af1f5f94c 11e3e9402c3ac558 f500199d95b6d3e3 01758586281dcd26",
				},
				{
					Message: "abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq",
					Digest:  "04a371e84ecfb5b8 b77cb48610fca818 2dd457ce6f326a0f d3d7ec2f1e91636d ee691fbe0c985302 ba1b0d8dc78c0863 46b533b49c030d99 a27daf1139d6e75e",
				},
				{
					Message: "abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmnoijklmnopjklmnopqklmnopqrlmnopqrsmnopqrstnopqrstu",
					Digest:  "afebb2ef542e6579 c50cad06d2e578f9 f8dd6881d7dc824d 26360feebf18a4fa 73e3261122948efc fd492e74e82e2189 ed0fb440d187f382 270cb455f21dd185",
				},
				{
					Message: strings.Repeat("a", 1000_000),
					Digest:  "3c3a876da14034ab 60627c077bb98f7e 120a2a5370212dff b3385a18d4f38859 ed311d0a9d5141ce 9cc5c66ee689b266 a8aa18ace8282a0e 0db596c90b0a7b87",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			for tvId, tv := range tc.TestVectors {
				userService, err := users.NewService()
				require.NoError(t, err)

				svc := services.NewService(userService)
				openSessionResponse, err := svc.OpenSession(&messages.OpenSessionRequest{Flags: 0})
				require.NoError(t, err)
				require.Equal(t, messages.ReturnValue_OK, openSessionResponse.ReturnValue)
				sessionID := openSessionResponse.SessionID

				loginResponse, err := svc.Login(&messages.LoginRequest{
					SessionID: sessionID,
					UserType:  users.TypeUser,
					Pin:       []byte("password"),
				})
				require.NoError(t, err)
				require.Equal(t, messages.ReturnValue_OK, loginResponse.ReturnValue)

				digestInitResponse, err := svc.DigestInit(&messages.DigestInitRequest{
					SessionID: sessionID,
					Mechanism: tc.Mechanism,
				})
				require.NoError(t, err)
				require.Equal(t, messages.ReturnValue_OK, digestInitResponse.ReturnValue)

				chunks := chunkBy(tv.Message, 64)
				for _, chunk := range chunks {
					digestUpdateResponse, err := svc.DigestUpdate(&messages.DigestUpdateRequest{
						SessionID: sessionID,
						Data:      []byte(chunk),
					})
					require.NoError(t, err)
					require.Equal(t, messages.ReturnValue_OK, digestUpdateResponse.ReturnValue)
				}

				digestFinalResponse, err := svc.DigestFinal(&messages.DigestFinalRequest{
					SessionID: sessionID,
				})

				require.NoError(t, err)
				require.Equal(t, messages.ReturnValue_OK, digestFinalResponse.ReturnValue)
				require.Equal(t, strings.ReplaceAll(tv.Digest, " ", ""), hex.EncodeToString(digestFinalResponse.Digest), fmt.Sprintf("failed on test vector #%d", tvId+1))
			}
		})
	}
}

func chunkBy(str string, length int) []string {
	parts := make([]string, 0)
	part := make([]rune, length)
	idx := 0
	for _, v := range []rune(str) {
		part[idx] = v
		idx++
		if idx == length {
			parts = append(parts, string(part))
			idx = 0
		}
	}

	if idx > 0 {
		parts = append(parts, string(part[0:idx]))
	}

	return parts
}
