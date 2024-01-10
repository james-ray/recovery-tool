package main

import (
	"encoding/hex"
	"fmt"
	"github.com/james-ray/recovery-tool/common"
	"os"
	"strings"
)

//export hello
func hello() string {
	return "Greetings"
}

//export generateChildExtendedPrivateKey
func generateChildExtendedPrivateKey(metadataFilePath string, walletType int, vaultIndex int, chainInt int, subIndex int) string {
	// 实现生成子扩展私钥的逻辑
	hdPath := fmt.Sprintf("81/%d/%d/%d/%d", walletType, vaultIndex, chainInt, subIndex)
	metadataMap, err := common.ReadMetadataFile(metadataFilePath)
	if err != nil {
		panic(err)
	}
	privBytes, _, err := common.DeriveChildPrivateKey(metadataMap, hdPath)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(privBytes)
}

func main() {
	//privBytes, _ := hex.DecodeString("2d2d2d2d2d424547494e205253412050524956415445204b45592d2d2d2d2d0a4d49494a4b41494241414b4341674541744133336d692b534c2b2f6761466655414e6248776d796d556f345455565031643850507a53755042374c774e3444540a5a2b61412b432b70635579546f4a3554624b6c786d334d355230346e636b74627649536b725132445633494642343537716a43506165346d654837764f5873630a4a4f3671704f342b3972745362773661364e4f6a2f31436e345056612f6c7259686678486362304d594e726d6d4853463377715854316b395151306f7a564e4e0a67656b69483262316d386359506a4745536b366153796443685568374c72513958416a6c6a50524a4a6267745458755a76486539327245654768796d7277506f0a4b316970637436695763694c665a636d786a525a4b464b524c7745686d63514353536e4c7047362b697834773543342f7a61534f6e7975455a765168665371560a36663532755876666a6d654d356f726474714d7058775433683453416f6e5435464450346f577a744977457a766e5130346a484253494c7a6c45734635337a570a7252337237503250385273542b45646867454c30385a5a734f4474436d3973766d6a50394247515866636441435274437131754576652f5971516331315156330a33796159345966614e42386b526752745a41327958774769464a782b465358302f6f67614b716f3758513152697748736a43476f5a56544a43612f66305a68760a36346c72704859426a4c6e69704f555449483452503632626847324d49387763526635643059316b5264524c33464f4175344a7a7962502f354b364b595738590a6d78586e6b33587a57384f79735553566a655775443465593068735967552f6b767a365a417532625142646a3669594d42626d6857683730547933394b3768430a306f4938432f696f6f4e68744f6966633739633359524a326574636d3331586a577377776971714c4e356d716953454547744c72396c736a6c336343417745410a41514b43416742356b365449753635654a2f374339657230316848304f69446b52594c4e7533395836795035535a772b6570697849592b6437497252436e33440a5077322f757168694374666e4f78704742344f6a4c546334522b685a345848364c6363766e776e706c506942636f6b5065385a55626368484e39434a7055566d0a6e48334139332b71447034423235672f306577724b335267704551706a4a6b37485250766856365069447337484f713762674e4c4153436f773536437a6161790a4650464e72576e434d756b454c664a4178507849515050646e545243785142504d396d444a4235656f4176574378706b5430784e694c6e5036774966796979640a7130686b2b5262355a6f64394c714e527464585757662b6848314f654b426b4a48733273685834586754473855456f4b716e5769616f726e675362786636386f0a484d754c4c2b45645858616c376c696c624e62653231554553304e7075685647414c72654363644e674a514b6c6a4a35625638514f525164354f39442f456c6a0a44445a46704653334f34537170382f445430786a5069714b2b6e76692f564776775564596f477656484a4c384c774368426176696f715446464f4843693338740a4469424f4c4a50584150643563442f5547347964636a58485270584c756c416e6e674d76304544414b6873756c586d482f4147344f776a424c5138636268376c0a65534b727554324b33372f737656796377716e332b5454306349704c47783562777068316630714872733756392b2b5a2f756a6d4a4c394f61436535483861430a744a6e41323969432b7069674b524d2f6b663961576b474a7146482f3653345841624768664f622f366d4433493741613638304f4f3044463146345455786c390a617547447935543845316c4243304c412f436866494248644f35757555557133735441434b34686e3848596774446a6641514b4341514541343170364a4d704b0a663577385557547443396d7732456c482b69385573574264387a4d4e68714c375847373037704e424c38555268665141757a6233567a3264505443447a62724c0a6165484c4754302b4a635a44765361342f312f636d42656476764a6a335338527043682f7a70634b3231696b7a336c52323169452f305673726e39383253472b0a4173464452736d76707970584a3267652f7a6e50316232344c32635a7337636c624e75576c6e5630342b5069696e78484870506d3875525462316131475041360a3852375135465359525352796a6d6759493263472b5353396f3533554d6b39562b7151387435304f4b4d31373567453247474b464d32587332584d6f786e51670a4f3330696b47325057524a756c3377336b55364d4e5537456e325a47494a363230676231484e354338627454335069616245654655444167666467734e777a4b0a2b49654f493743624f50716631774b4341514541797233517968527861414d48454f42344d5549454238325a534f72385547785162507036516e4b77696b454a0a304155765a45466f71636a30354f4a4d6e4f4d6c5957742b675378712b7362477a57584a436970655a30473348692b704c66357539515966764a6650484152620a497479444a683458303037397147636d734d504b5253694e32385a692b4c756269702b67495946657a5933704c6f31774a7a47504655573256726859367957670a704f45534d50416139653233434a57425946425a39746854685546414177574743756d596a59677a422b4f766e4c4c347a47667534715a70304a54535a384d750a72743333472f33777a4d4952416f71616e73512b5452665369335264474b78366d775961773833596a77312b326b766d6547686f30324e78485a5664535658730a624f506a486778487866494b466d7a78774e43324c7a306a744d502f544f726370756d654f665a5259514b43415141613634757137346b43737930784e6849620a5a527462674e48552f64346c59704f39534435427775716764304c70504f5a72455a71526b654c4553433368567070587448626d3155646773697571515759710a63743979646a4e52696268464367625470542f4e344e546c36795733414974346a585a325770636d7363534e7456713544723970746d555a546d6a34364d697a0a736e2b2f577354513037655952323658726b324d79684c55594f766a784f536956306c72764774765933506c4c73507957774832676674347358317169396d370a3169543656376b442b42384c5152357a55537a66434358574637785977572f37784e6a323077336b74555777594645374c6a65567941704a4150676d77644a480a6162644f6d4a61377a6545734c72643561464c4661675762754e6374492b517057315543785330447854517a326e372b5a3670556d6c38754c6c37574f774d6d0a4c416268416f494241466e417747374673424c677a44372f634a593136487a66327730353469746b57424a32724b7441424d476632303035446d7275766676300a46393541424c5372706a2f4469385235756e35386730516f3347426773317734376d30786f6b3758525a3235635646585433765376646e6f447a4e5076544a780a314c2b3573706f7367783473767568646f6a45464e4f32676a77356d4c47634a66514a373450756f352f503934615077686f544c4b70767a6538386f624863560a48714e784e4535454a422b77557745364372474c7633535452314c597965616a4f4870426a50314977617562436c73706c7941536531633073704730304f61450a763230466d542f5346746167526a6777636f706179516861496e30755973477073732f484c43642f3658417a704c373541637857656f30574d6452366b55656a0a53433333577573474f7245632b764b716c686f69477a3756533550545a4f454367674542414a3775624a62776c39396261696d4b4b79337358643530554752450a586f7267526e3068375079626a457a706763356d4b71754c61634d5734386b343679454a56704135664c4c52355a356c4336794e7972707a5947762b586734720a664e556152686b435268764862617353796d694965754c344d4c4742685043326e345a3555514c7643524657794764766c6e5078587347696f4b672b5437764a0a53766b6f6e30417653762f70624b6468496d383744526a584c4f344e34337375647266706b44763659616b63544e482f41307373686365626f55786c544131440a6576655a6b2b72476e797a5167647a4a795872475148416e65326d4f765a7639755977555141304d3047774b6d63727473744f706464554f52504d48392b49450a373842576356364c757242502b6871696f793876765238434e315a4e73675a45474d6663536133466b2b494a666d31626e7377562f70323344736b3d0a2d2d2d2d2d454e44205253412050524956415445204b45592d2d2d2d2d0a")
	//fmt.Println(string(privBytes))
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	if os.Args[1] == "makeZipFile" {
		if len(os.Args) < 10 {
			printUsage()
			os.Exit(1)
		}
		//"2. recovery-tool makeZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [pubkeyFilePath] [userPrivatekeySlice] [hbcPrivatekeySlice1|hbcPrivatekeySlice2] [chaincode1|chaincode2|chaincode3] [pubkeySlice1|pubkeySlice2|pubkeySlice3]")
		//arg[2]: zipFilePath
		//arg[3]: userPassphrase
		//arg[4]: hbcPassphrase
		//arg[5]: pubkeyFilePath
		//arg[6]: userPrivatekeySlice
		//arg[7]: hbcPrivatekeySlice1|hbcPrivatekeySlice2
		//arg[8]: chaincode1|chaincode2|chaincode3
		//arg[9]: pubkeySlice1|pubkeySlice2|pubkeySlice3
		pubkeyBytes, err := os.ReadFile(os.Args[5])
		if err != nil {
			panic(err)
		}
		//pubkeyBytes = pubkeyBytes[:len(pubkeyBytes)-1]
		/*fmt.Println(string(pubkeyBytes))
		fmt.Println(len(string(pubkeyBytes)))
		fmt.Println(len(pubkeyBytes))
		fmt.Printf("%x \n", pubkeyBytes)*/
		hbcPasswdBytes, err := hex.DecodeString(os.Args[4])
		if err != nil {
			panic(err)
		}
		_, err = common.MakeZipFile([]byte(os.Args[3]), hbcPasswdBytes, pubkeyBytes, os.Args[6], strings.Split(os.Args[7], "|"), strings.Split(os.Args[8], "|"), strings.Split(os.Args[9], "|"), os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println("the zip file is successfully created")
	} else if os.Args[1] == "parseZipFile" {
		if len(os.Args) < 6 {
			printUsage()
			os.Exit(1)
		}
		privkeyBytes, err := os.ReadFile(os.Args[5])
		if err != nil {
			panic(err)
		}
		//privkeyBytes = privkeyBytes[:len(privkeyBytes)-1]
		/*fmt.Println(string(privkeyBytes))
		fmt.Println(len(string(privkeyBytes)))
		fmt.Println(len(privkeyBytes))
		fmt.Printf("%x \n", privkeyBytes)*/
		hbcPasswdBytes, err := hex.DecodeString(os.Args[4])
		if err != nil {
			panic(err)
		}
		d, err := common.ParseFile(os.Args[2], privkeyBytes, []byte(os.Args[3]), hbcPasswdBytes)
		if err != nil {
			panic(err)
		}
		fmt.Println(d)
	} else if os.Args[1] == "deriveChildPrivateKey" {
		if len(os.Args) < 4 {
			printUsage()
			os.Exit(1)
		}
		metadataMap, err := common.ReadMetadataFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		var addr string
		_, addr, err = common.DeriveChildPrivateKey(metadataMap, os.Args[3])
		if err != nil {
			panic(err)
		}
		fmt.Printf("derived addr: %s \n", addr)
	} else if os.Args[1] == "generateKeyPair" {
		err := common.GenerateRSAKeyPair()
		if err != nil {
			panic(err)
		}
	} else if os.Args[1] == "deriveCsvFile" {
		if len(os.Args) < 4 {
			printUsage()
			os.Exit(1)
		}
		records, err := common.ParseCsv(os.Args[3])
		if err != nil {
			panic(err)
		}
		metadataMap, err := common.ReadMetadataFile(os.Args[2])
		if err != nil {
			panic(err)
		}
		for _, r := range records {
			_, addr, err := common.DeriveChildPrivateKey(metadataMap, r["Path"])
			if err != nil {
				panic(err)
			}
			fmt.Printf("derived addr: %s \n", addr)
		}
	} else {
		printUsage()
	}
	os.Exit(0)
}

func printUsage() {
	fmt.Println("USAGE:")
	fmt.Println("1. recovery-tool generateKeyPair")
	fmt.Println("description: will generate two files: ./private_key.pem and ./public_key.pem")
	fmt.Println("2. recovery-tool makeZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [pubkeyFilePath] [userPrivatekeySlice] [hbcPrivatekeySlice1|hbcPrivatekeySlice2] [chaincode1|chaincode2|chaincode3] [pubkeySlice1|pubkeySlice2|pubkeySlice3]")
	fmt.Println("eg: recovery-tool makeZipFile './zipTest.zip' '123123' '456456' './public_key.pem' '5ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442c9' '7ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442c9|8ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442c9' '4ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442c9|4ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442ca|4ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442cb' '033669d206967b384d588b366b6400733987befc6604fec764f9fc2d42a3bf7a86|021b491468a9c042e6d4e994c3979df14454cd99e4fc207161a929e719f644540b|02f75cebd23a9ac7e1364d0462be378f09aaf26474eb46cc43bdef5de2817932e5'")
	fmt.Println("3. recovery-tool parseZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [privkeyFilePath]")
	fmt.Println("eg: recovery-tool parseZipFile './zipTest.zip' '123123' '456456' './private_key.pem'")
	fmt.Println("4. recovery-tool deriveChildPrivateKey [metadataFilePath] [derivePath]")
	fmt.Println("eg: recovery-tool deriveChildPrivateKey './metadata.json' '81/0/1/60/2'")
	fmt.Println("4. recovery-tool deriveCsvFile [metadataFilePath] [csvFilePath]")
	fmt.Println("eg: recovery-tool deriveCsvFile './metadata.json' './walletPaths.csv'")
}
