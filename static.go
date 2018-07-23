//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a simple array containing one entry for each embedded
// resource.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES []EmbeddedResource

//
// Populate our resources
//
func init() {

	var tmp EmbeddedResource

	tmp.Filename = "data/compose.html"
	tmp.Contents = "1f8b08000000000004ffbc565973db38127ef7afe841a676ec9aa5e86be22320abbc72e21c4e625b4eb2993790685190418006404a9aa9f9ef5b20481d5e6fd6c9c35855261a68747ffdf541d29fce3f0e6fbf5ebd84892b65ba45fd03245345425091740b804e9071bf00a04e3889e9509795b608efd15a56208dc3f6167895121d03c54a4c4823705669e308e45a39542e2133c1dd24e1d8881ca356f82708259c6032b2399398ecb54e0182a189735584f7b56812320c46a2db45856b261dce5dec61bf807cc28c45977cba7d151df7767e8a22b8640ead835c979590c881290ea550622c90c370348228eae04ba1eec0a04c88750b897682e8084c0c8e13e2c1d8d3382ed93ce76a9069edac33acf242aecb78b9111f0c0e0647716eed6a6f500a35c8ad252094c3c208b748889db083e3c3e85f9fbf0a317af30adfedf18bf2edcdd9dd22af5f9fbdbe290ef63f969ff2d9ec48ab839bafbc38fccc7ebd2a47b7f68ff8ddf3e326e32fa793c39a406eb4b5da8842a88430a5d5a2d4b52569c889a7607a5da35980b0603c9f0639640b58428625033637a272604dbe8a38d71c07d37b6fa18d342ca3bdc1defee0b08d6cfa6860aa3963bbbf3ebf2e4e7239fb3abcb8b8aaceaf2f3f57971fd4eef886dd7c9cbebeb77b07eef0a2981ebcfce3b7dd7fab8bebfb3c36c7ef27eae6fc7702b9d1d66a231e0646e380b34b9b8ff09b497ecb1a360a917d33d2a7e676fa30b58f13709bfff6e65a64bbfb47f7cd623a7a3f7e3dfdf89e5dde8deb2f9fe7bfcf3f5da9e1dbb323b95f0ebf7c78535d9c9417c3f3e3d9c58737f9d5f9d1ed9c9127101072dc962bb845854968095f6b819db6b7e1cf560f003296df1546d78a47b996da9c3e1b9ff8df8ba0f057780cacc306a3b1d60ecdea72c53817aa3885c3dd6a0ebbdd1d8060099e9d9cf4761e71049b9e00326d389ac8e9ea14f6ab39582d05876778e27f9ded47f10c146b223f98d6b1751832c9f2bbee32c0582b17cd501413770a9996bc3b096669dcf21678ea4b7f8dc4296b5828b49ecb9fb781ebbc2e5139d81918647cb13dae55ee8456db3bf027fcbcfd8b5055ed4e1dcedd6923acc8249e8e85b1ee979dc158e7b5ddde79017fedb43cafd7318d7d441e0bcd345ff8050055ac815c326b13a258933103e111711cb35a2e8101e562a9e9072e130a4d3496b5e03d786f6f4dab33e4bda259d301a059ed9c561d1541209b3022a78b42a2675db2ca2227c09963dd7642721df6fb6d660a740979169c126046b008e715531c7942c64c5aec763d7aa3e532e20d6800d4564cf560ac89b4920b92de06388a35a2603e1b34f67a81c5fe6fe3aac8b58a326648fa77a9d23850b90e8ab23e94c04c9419a678ffce8949fa05b392094963b6ba46632e9a35d12755f0be421e66aa4fc532571b84d25aae21f00466cc448a351b5a00548a943208efc2d8bf4bb5c598f49f031e1f8da558a102a0712d37e4ffe5aaf76a7c9ffeb75f58732c75a16b47d2cbf6f9ffbd6e50b526d058b18e423ad95b4531d95b21a68e65bec443fb05a1fd1f59674485bc9326ba41d3adc34c5b1ee55a7154be3d7ab3d499943a9ede6a1a3bde2ea9f013a3eb368773e74bd29fc5ce3cbcf6cae8f2872e8eea6c8ab9fb91bbbe707c3325649fa4d403640619183db30939dcf55f63be5b8f775bdc386f8fd3e0a90d81c62d6d21165aa5ff5099ad5ed0b84a7f440893b37b35f9b20fcb65d90711d6df60048c96d83648371e56f9e08fcccde5e9c3b92aa39247cfd78ee141036d1cb55db3d65dcb593bd2b5c911869ae3c3a669efac0abeffea2c849bd459fbf965efe6f12c4c05929e354cc8b64cb582a0f49d4df1dd21c2d3621cea6ad136f413236cf335b8c3b21a8c454cd291ff0281775856df19d046978762f03543e34cf345ba45e3892b65baf59f000000ffff4c0f3e40fd0c0000"
	tmp.Length = 3325
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/folders.html"
	tmp.Contents = "1f8b08000000000004ffbc57df739c38127ecf5fd14bf6cee3ca01e31f9bc463a02a9924761227766c27b9ecd53d08d4301a0b894802cfec94fff72b2198617cdedddc3d2c2e1748b4babfafd5fae8897e7a753ebdfe76f11a66a6e4c9a3c8de801351c41e0a2f790410cd9050fb0010196638266f24a7a8e08c6913856eea11d8d7251a028294187b0dc3db4a2ae3412685416162ef9651338b29362c43bf1dfc0398608611eeeb8c708cf7da8000ced1cc98cac7ef356b626fea9cf8d7cb0a072e0d2e4c68211f4336234aa3893f5fbff19ff77e7ef27d382306b5814c9615e34881080a25132c6748617a7505bedfc1e74cdc80421e7bda2c39ea19a2f160a6308f3d0b464fc2b0248b8c8a2095d268a3486507992cc3f54478101c04cfc24cebcd5c503211645a7bc084c14231b38c3d3d2307cf0ffd975fbe3176f5f60dbedfa327e5bbcb1737cbac3e7d717a591cec9f979fb3dbdb67521c5c7ea3c5e117f2e4a2bcbad6bf85ef9f3e6f52fa7a3e3bac3dc894d45a2a5630117b4448b12c65adbdc4ed894dc1fc538d6a094c83b2f95448215dc21a32ac33a033c52a035a651bc699a418ccbf5b0f2d53f7e8ef057bfbc161cb6cfe2031d1bc20e3274f3f154719bffd363d39b9a85e7d3afb529d7d14e3fc925c9ecf4fbfebbd03737852cc0f5efff6cbf89fe2e4d3f72c54cf3fccc4e5ab5f3dc894d45a2a769f58143a9cddb659867fb8c9ef4843ae1cb33f64faa37b3bbfbfb50f27e03afbe5ed27968ef79f7d6f96f3ab0ff9e9fcfc0339bbc9ebaf5f16bf2e3e5f88e9bb17cff87e39fdfaf16d7572549e4c5f3dbf3df9f836bb78f5ec7a41bc1f4880dbe3b65cc12c2b8cdd91b0b5e6b2d39e6b58b576009092eca650b216d4cf24976af2383fb27fc7cee0cedd026db0413f97d2a0da2cae08a54c1413381c570b18776b009c27787c74d4fb7920106c470248a5a2a87c23ab09ec570bd092330a8ff1c8fe75be1fc41308d2f8569486d83a0c2927d94db7182097c2f8b7c88a9999402a39edde38b751d8e6cde5a92ffd4112e7a421aed0fa5cfe3c022ab3ba446160375048e87294d722334c8ad12eace0e7d10e13556d26061766d230cd528e939c296d7676835c66b51eed1ec3dd6e9be7611d47a16564b144a9a44bfb001009d240c689d6b12748931205eee653cc49cdd7c020a26c6d6905973081cacf79cd680fdefa1b58758e6c5454031b8028ad8d91a24b851b78db307c238b82a3cd3a279546ea01258674d3b1974937df4f1355a089bdc72ea8074431e2e3a22282228dbd9c708dddac45af245f33de820610e98a881e8c56be147ce925d70e8e200d2b88dd8d28b4762e8bfdb5b5946552f829515ef2579946a14be51054447a2a2e337eaa88a0fd3727f492af989684f128249b655148593318da4d65b4af90fb3bd56fc57aafb6121ad57c80c0263025ca17a4d9b20288384b2202ee5b187ac9a92cd1828a42ce365000a2b0e65be3dff3df8752f670fe77301844e3b290b5f192b3f6fee751b7f2331844a1205ddeda9c75074bc9db75fce17c26b9af4bffc026ce2fa97fe0254e6eed353464f6c4fb5655abb5237b6dd7dbc6c827944ae12511eb21147c59cd6c45c2fac9d7485436b3e5c91e28d1a8f5d71d51830be3b52590336e50795d0fd68f1ac26b8c3d0f2a4e329cb5dd5bbc7edb8a60d79d4df6c6e3bf0d4874d96bd5ca5e43da9ce99eb58bddfad583d500ab9522a240085ccba8e1ce49afbb36d5bff1e5338325a446d8ff5ee6c0f9f67aac96b04f382bc48463beeed1426716ae56c1dd9d074a728c7bfd4ada495b3c6b32161e0a3a80d4d1edd00d8fd990785718477d611c0d2847b3fde4bc4165dbdf289ced0f4a66b5623904af95926a18b2b2c8bad928ac36f0562be41ab74d2f38128da09163668074790123c1068412b52605eae0bea301cb01ade16395fc5da4ba3ab62bff9f419bb4a8eb18ac1ab9c7b51ab9210c1b8b7e8736aabd4ee476ba85fb9caddfdeffdcb5c7f3e9e035dcd3b5ad57ad98f5276fd34a78c995ac5586309514efcb5abb662349fd8f818299599db65db1be5984b74eacbde445431827294790029c91adbcfb5eb7c572b01bf6da4ec19f52841fe33895d5b295dcfb587e8761bb5fc10d965590b3d04bae6c6308efb1acfe4742037651e88aa1ab19d7620da4ec81c6ab210a9ce2c5eb0e2c28d0bce6689bb197cbb774b4e3e46e67b76bf25afb408a1b5cd615c4d0376a603bb5b65c015ac76e1dc4dd8a562d03233f5715aa29d138ea3d3a73cef4c328f4cbe5d4f66a1f4989a31da7456b34b61955306a89400ce3636010595f01475198d931b0274f7637bdb68b657f4b436ccdfec5fe1d3021509d5e7f38eb18da8be530b256db800326282ecef391e3b60b490ce3dd9eb4bd386b5db6e21f50a62b4e9610c34eabc45683770641ac10fdc86221056ed67552ef6edbbd6e2ae9327914853353f2e4d17f020000ffff4c80483985100000"
	tmp.Length = 4229
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/login.html"
	tmp.Contents = "1f8b08000000000004ffcc586d73dc28f27fef4fd14bb6fe6bd7fe35f253627b2c4d9dcfceda499cf831c979df2181348c11c8803433d9daef7e0548b23499cd65efea1e70d50850d3fcfad74dd372f4c3d9d5e9fdc3f56b989a824f3622fb008e451e232ad06403209a524c6c072032cc703ab9943913704b9f2aa62889423fbb0156a2a00683c0058d51cde8bc94ca2048a530549818cd1931d398d09aa5347083ff0726986198073ac59cc63b6e4f00af686a4c19d87dea189d7a25c1fdb2a43d95862e4c68511f433ac54a53137fbcff25386cf5fc100470890dd5065259948c5302581028986019a3044eefee20081af89c89475094c7489b25a77a4aa9413055348b9105a3c76158e0454ac42891d268a3706907a92cc26e22dc1bed8d0ec254ebe7b951c1c428d51a011386e68a99658cf414ef1dee077ffdf4c0d8dd9b5fe8bb1d725ebcbd3d795ca6d5c5c9c56dbeb77b557c4ce7f30329f66e1f48beff09ff7c5ddcddeb2fe1bb578775425ecfa6fb15825449ada562391331c2428a65212b8d26de279682d94d45d512980665f9549440b2840e32740ce854b1d28056e9b3c5a92474347bb21a9ca5be1bec8c767647fbceb2d95ac3447d82b77f7e75931fa57cfe707a7e7e5d9edd5c7e2a2f3f88edec16df5ecd2e9ef4ce9ed93fcf677bafbfbcdcfe9b38bf794a4375f87e2a6ecf7e45902aa9b5546cd5b028f4381bb7590bbfe9e4b7b8c677deb26f5afabdbe9dadba763d01f7e9cb37372cd9de3d78aa97b3bbf7d9c5ecea3dbe7cccaacf9f16bf2e3e5e8bd3b727077cb738fdfce14d797e549c9f9e1dcecf3fbc49afcf0eee17187d0701dec72e5cc12c4b1afb236163cdb303ee6cc36f4e1020c1e963ae642548904a2ed5f8457664ff8e9bf7bf37cf9136b4a64126a5a1aa5b5d624298c8c7b0bf5d2e60bb5de435c18ba3a34ecf57fbc0ca46895484aac0c8720cbbe502b4e48cc00b7a64ff5ac5ebd18c04ae039b977ac81a0809c7e963bb3a93c20473caf2a919432239695f78b551e868f334b591dfe370866bece3aca1f2c74d2032ad0a2a0c6c8d14c564b9995522354c8acd2df80d7edcfc8989b232634317665c33cd124ec71953dafcb435ca645ae9cdad63f87dcb91dc8fe228b4f658285122c9d2766c8b04ae21e558eb18095c2758817f048466b8e22d32db22c23a599b723113540519af1869f0fbd6976b94d9bda91a48014449658c140d237e8086600223f39c534b3ec7a5a60401c10637d3314aa59f6fa7b1caa989d10bbf2d02ac180ee8a2c4825012a30c734d9b596b8192bcb37b051c40a44b2c5a385a0552f0259adc7b4002d72cc7d62f5168e55a3edb3658cc522982042b34f9cf0a47a12775082ec2ad519ea520515890f6160ad1e4334d0acc7814e2fec22824ac1e4cd87060a48d9b55cfb5aee97cb7427054f11e0e4b6882552070bd22071071368930f85b3244930b59500b2e0a39eb030288c28af76756300f865128703df1e90d06316b23c3c5760f49672cb7c549221708dcd98e518155ce844b332fb7cbc571c743819915b4f60705095eb51d99659a9a60cf8d75111cb69de6c56e6fdb21b0120bcac1fd064c641241dfd835b2eedc31910f34dae6cc698ebdd7e78a2ce46baf014fbead9bb2fe6f4868d2b663616fbb5c742478e536e3ac826de0360a08d325c7cbb190822217558ee80073ea2a3c9fa1dc08dc6f40b0c8a96ab9dbd9b5876b25426d8b32a98a678576d481b383602a15fb2285c11c81929cc66e1a0176593746a10382a0a0662a498caeafeeeebfe273684c13128934461663d87dd92384d9f41dd80bb25ca36525ebf484034c88146812b1d66f395f96539b5aa0eb05399709b54cb03f4835b6454eade3c456d2a849c0862e6c2ded9976d4d883a024474dc13d95da20a831af688cda5595e2084a8e533a959c501523566057c7dae7882e705172eacaba35f6ae092bdbfaa1f5df23b3d2f6ca8afe2d5c7ae23b2e0704da7d4756e02f8ef1e119b7ed7f9a352ed3c73f118125d67a2e1569e3e979fc8d48b4422b51f7202b0576de29fb5738b305e34e3f87b9b3e04e2c9415e781b2f5de9a2d9a24d0e04e25b7d97d67d77d9eda4a63ed12688fa3af0b759514ccf80cd8f66d28c4dda88919f791dc65b2c408488cf0b7c2d7c6db9b714d6efc83e928b4160fb57c15722b1383613370352840544efe4f24ba3c8ec272f2cf0c9c9ea8f950b00587ef76c6fb21f4bf27da54fe5ca675acb8dcd2f948acbfe9bbf7feeaee16db362c5c06af5cadd22e7efe8640933b59a994c2a924744dd5d2af6fdaff02e4cc4cabc4e54dfdb808e7be264393931a338e134e410af042eb6aa16125d438c4a2b76d48c13f3411becfc653592eddd1f84e0b9dbf468fb42847190bd1e4ce7e0fc23b5a947fd2a09e753674edf7a48d99284c24594e36a2706a0a3ed9f87b000000ffffac1983648c120000"
	tmp.Length = 4748
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/message.html"
	tmp.Contents = "1f8b08000000000004ffbc59fb72dbb6d2ff3f4fb165fc7d92262669c76913dba4661c3bb1dd3a37db494edae99c818895041b041800ba55a3773f0380a448c5ceedcc9cf5b4226e8bdf5eb0586c925f4ede1c5f7f7afb02c626e7fd0789fd014ec4280d5004fd0700c91809b51f0089618663ff156a4d4608cb65547e462f25a7a860b582bf964ba8bbdf9f9fc06af57712fb850fc032c9d1101024c73498329c1552990032290c0a93063346cd38a538651986aeb10d4c30c3080f754638a6bb0e1680673436a608f1f3844dd3e0d83309af170536581a9c9bd80a7608d998288d267d7ffd327c56f1f9250ce18218d4063299178c23052228e44cb021430ac757571086257ccec42d28e469a0cd82a31e239a00c60a876960c1e88338cec93ca3221a4869b451a4b08d4ce671dd11ef457bd1d338d37add17e54c4499d601306170a49859a4811e93bd674fc2e71f3e317675fe12ffd8a5a7f9ef9747b78b6c7276747639da7bfc267f9fcd664fa5d8bbfc44474f3e90476ff3ab6bfd4ffcc76fcfa603fae266fc641240a6a4d652b11113694084148b5c4e74d0f736b12ab8793741b500a641597d2aa43058400d196a0de84cb1c28056d95ae24c528c6e3e5b0e4e52ff19ee46bb8fa3274eb29b3b0513d323b2f3e8b777a3fd8ccf3e1d9f9ebe2d4ede5d7c282e5e8b9de125b97c7373f659efee9927a7a39bbd17fffcbaf32f71faee7316ab67afc6e2f2e4cf003225b5968a6d0a96c41e6769362be1578dfc3b99922b2fd95725fd5edbde6c9af66e055c67bf9ebf63839dc74f3f4f173757af8667376f5e918bdbe1e4e387f99ff3f76fc5f1ef474ff9e3fcf8e3ebf3e2743f3f3d3e79363b7d7d9ebd3d797a3d27c17728c0dbd8b92b984581a93f12d6d7bc76dce987a59b07000392dd8e949c081a66924b75f070b86fff0efd8495ff89b4c1298643290daaf5e28250cac4e8009eec1473d829d700784ef0707fbfe273c746d0de0960201545151a591cc0e3620e5a7246e121eedbbf92f79d782241a6a10d5d4d6c25860127d96db9186028850967c84663730003c96939e2d926b1d39bd753e5fa0d25de9029f18e56e972ab0b5466931c85815ea490d045773811996152747bb084ad6e878962620e0ccecdc1946936e07830644a9b4e2f1aca6ca2bbbd4358f59c9e9b7e9cc456228b251948bab01f00892053c838d13a0d04990e8802ff13521c9209af814142593dd3065cc204aa70c8278c56e02dbfc6ac9291dd1555630e403298182345a90adf08da30422347238e56eb9c141a690094185276a741267d7fd54dd4084d1a3cf49b0640142321ce0b2228d2341812aeb1ecb5e895e4b5c42d6800892e88a8c068154ac11741ffdac31164ca46c45a2389ed3cafc58a5a4b59264538202ae8ffafa626b157651354422a51bc66c281228256774e1cf43fe220278c2731592f4b62caa68da6352aa395876c5aaa32456dab964293096f20b00a1c10150a326dcd024838eb2704fc5d1807fd3399a30595c49cada10024f184b7daf7f1afb652f6707eb9193476e372242726e85fb8df6fefdad24fa391c482947a733a2b0f9692b37aff667f2679a8f370cf2a2ecc69b857cf6a1f37660f7c68836ad198b1e1a98d4921a1548aa09fb00ac1882f8ab17548a8bf428d446563eb9dec0e0f4d1cbff2841a9c9bc079c09071832a2853b0aa35257c82691040c149866397cfa5f5a88b81657276b0bbb3f37f0d211acab3d4d40e67ba12da6fedd8eac66280e5521131422873480d2b1f783dad7d7fcd2b640673181861ffab821c78de4105d5ca1b12ce46e280e3b0ced0623f2d5e2ea3d52a002539a655f4eabb4eeb3a2ef07a5a2e51d006a496b4ed863d62a5bb946eb15fb9c57e43e4e5920d217aa194544dbe85ddbeec4de2628d61b944aeb139d5e6326748282a5de72a9612430636defabbc037dcff436d142b9096adb19ca22abffd055b0f6552501436563701780b6ddde2621bb69ca7c041ba4ef32b2c0d8cf6b1a0fa89a1fde5d22d84d52a890d755d60fb3c9baa37366abda1758a4dad3b39d6539c0eae0dc9c6f69e5daba1698419720eb97fac84c4d4936bd12ad7db62dbb0d516e9683dbd96aac9bb190fda27bd34fd6ff52e965ab74a7d7e1b27b92005aa8cb3a2bce0c68c52146960d404edf9debca896cb2d8c5e328eaf49def40d80ae1b2a9f42f625b45af56a7f6e39ecf7e0aeaf9eeab8158ae5442ddcd19beba0e6eca8759cda43d5edfeef0127e27663b0bc31bc8de2e572ab7e41fa87e56ad5ea7c7f7ee27b9a2ad860e992074e06f6a1762267824bd24c722c7dd328b45ae7b97dcb2600d5460dd15be1a4a5fe56a31966ca019fb65b72cefe5cd245fbb4d71795a72f6fd1d090c1dadb3d2d9760c34ff5703f23faecfad545fbe8ba9bbc0a65f72e209961532c81b7bcdd529d073cb4996a689fdfa51ad7b95b63c47b8e21832a1974b95aea7afa16614b8f9636930a1b340005fd1951aa10fb4322d91be66e91fcc85745bac6b94b521a9e62f3a2769ee4ecd09e029b8a5564f6d3202ec9ec9b6a6de76a8decc85233221a3270a645d1ced4dc9c5a170511c8ebecd3aeb13d3024190213f77b676599d2c23e9d58fb8fbd65d63eed1dba3c4616a7a7ef4442bf8ec46244aed116bbee47e55ca09f140a2d34ebfd6b7cd6f6f6deb383ff15c88616ac1bdcb3dde5d14777cddeb1ddc6eee5592e55d6182c3fdd4052f4ff5f0c74716893949f69782e65d9c03e49fc67ed14be09cdea4295a8ad9f6e75bc71462d331e1b59dc9bb61efdf28acbe9e6d5dc7a7c34565a4a781d07d7f584a07f25272a43389614bf0c43ad575055111c31339e0c5c694cdfcee3997fb105fda32961dc656d52809f648fe426d7cd53f8e5dbae56817b866c641f2d11e1fb643c96c5c2bdbb36b1f8a854c7f74a4267afe816f3221ab238e85fd9ea10fc8179f1830295cee6d41f7b67287dc6d7591a0f9a3bab2feb5a4b5dd0daea7622e3b46cd45fae0661338ebf3bbd084936be6b05c056d78c99eed98a6cb7934d9496aab3dd29a42dceaa4e2f724974d7c12ce98e8d3d55ac08a5c73637ef767ce0e8f4ca5293a7d5f60f715398cb297e95612f92a2dbc9e544e3a4e86cd708a18bbd75c5ce939e31938da18bd16cccb2f13a677454abb2a238860b1c1a38e62cbb8d9ac02d654423ec1e6c765775b188cbcc5560205d2bc718d5edd4c6d9d08da58142727be8a35083e2185e314addc3e75e2c8fbfc032b575329c7d6482cad90fe198b925912c50746b0edbd0f1a9ed1dc0312a144e5198135f90ebb6ed6ea994cd7ed654056247eb247ed53bac7243fb69479b35425b53f065ecaf5629adf8be3e90aecd3242f382a37d513d5f9cd36ec717076a89dcfc488a5b5c4c0a481beed4f026cbd8af83b45ce19e7a9191ef8b02d531d1d850809dce99be1b857ebe70fe6d1f39dd8e7fbad7686ce55641d7326090c2ce2130482caf88a31899f121b0478f1ac0c0ed65ffe109523bed2ff677c4844065f385c6416443e8da596dc0111314e76f865d2f5b0ffa29ecb40e09678ea52b954494e9829305a4d071850b5bb2e83436b11984335d49f72d1652e07a5de911fea769f4241e48bae83f48e2b1c979ffc17f020000ffffb22c5fe3d81b0000"
	tmp.Length = 7128
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/messages.html"
	tmp.Contents = "1f8b08000000000004ffbc586d73db3612feee5fb16572276912927e6b13dba4665239b1ddda79b19de4d24ee7062456226c106000502fd5e8bfdf8020294a7172c97d38681212e062f1ecb38bc5c2d14fa76f46b79fdebe84cce47cb813d907702226b187c21bee004419126a5f0022c30cc7e17209c12bc9292a58ada2d00dee8015c8d1101024c7d89b329c1552190f52290c0a137b33464d16539cb214fdaaf31498608611eeeb94708cf7aa25019ca2cc98c2c7cf259bc6dec829f16f170576541a9c9bd0823e8134234aa389dfdfbef29f377a7ef27db82406b58154e605e34881080a39136ccc90c2e8e6067cbf86cf99b807853cf6b45970d419a2f12053388e3d0b461f87614ee629154122a5d14691c276529987ed4078101c04cfc254ebf558903311a45a7bc084c1896266117b3a2307cf0ffd5f3f7c62ece6e215febe47cff2dfae5fdc2fd2f2fcc5f9f5e460ff4dfe3e9dcd9e497170fd894e0e3f90276ff39b5bfd77f8fb2fcfa7097d79971d961ea44a6a2d159b30117b4448b1c865a9bda1f389a5e0ee5d896a014c83b27c2aa4902ca0850c2d033a55ac30a055bab638951483bbcf564365a97bf5f782bdfde0b0b2ecee41c3c4f405d97df2cbbbc951ca679f4667676f8bd377971f8acbd762777c4daedfdc9d7fd67b07e6f06c7277f0f2ef9f77ff25cede7d4e43f5fc2a13d7a77f78902aa9b5546cdbb02874386bb7590bbfe9e4dfc894dc38cbbe69e9f7faf66edbb50f13709bfe7cf18e25bbfbcf3e4f17773757e3f3bb3757e4f27e5c7efc30ff63fefead18fdf6e219dfcf471f5f5f146747f9d9e8f4f9ececf545faf6f4d9ed9c78df4180f37115ae601605c66e4bd85873ec543b1b96951c002424bd9f28590aeaa7924b75fc687c647f274e60e51e813638457f2ca541b59e5c104a99981cc3e16e3187dd7a0e80d3048f8e8e1a3d0f2c049b2b0124525154be91c531ec1773d092330a8ff0c8fe6add0fe2090499fa362d75b1d518124ed2fb7a32c0580ae3cf904d32730c89e4b4fee2d44661c59be3a909fd0e8977644a5ca0355c3eee03956999a33030081412bae88f4b911a26457f004b78dcef315194e6d8e0dc1c4f996609c7e33153daf406c158a6a5ee0f4e6035a878eec671145a8b2c96289174615f002241a69072a275ec09324d8802f7f0298e49c95b601051d64ada844b9840e58f79c96803deeaeb48d58aecaaa83a320051521a23454d85eb789b307c2327138e96754e0a8dd4034a0ca987632f956ebc19266a8226f61eb9453d208a111fe705111469ec8d09d7588f5af44af2d6e20d6800912e8868c068e54bc117def0d6c11164ca26c47a230aad9c63b1691b53592a859f10e50dff5fa251e8a8ec828a48638a63c64f1411b43973426ff811939c301e85643d2d0a299b76bad6a98c3611b2eda9c615adaf36088d4ade4160094c88f205996e4801449c0d2302ee2c0cbde1b9ccd1828a42ced65000a2b0e41bfdafe96f965276737eb9187456e372224be30d2fabe77f5f75839f4e270a05a979ab38ab379692b376fdee782ab9af73ffc012e7e7d43f68a536b71bb31bdeb749b5e8486c456a47c827944ae10d23d62098f04591d98084f6cdd748549ad9e8640f446854e9ab77a8c1b9f1aa0818336e50797509d6f4a68497187b1e149ca49855a55bdc7ead72605d9c1defedeefea363444d5e95ac6cebd2c3996eac766b577a756736c072a9889860532f6a58b9ccebda3af8d7ba7c663087c408fbafc972e0747b0d566bb04f389b88638ee3b6440b9d58b85c06ab95074a728c9bf435ac066decb4c65878286807526d6e8daebbcbba86d77171d4c4c551c7e4e5928d2178a994545dbd855dbe1e8dc2628d61b944aeb12b6a4862f3aacbf9ae53fdef6ba35881b4ee65728aaa7e770769fb299582a2b03979bd8e2dded53032d9f094188c4293559d574ae66de7a64cee3035ae1f1ad59ddd3af20ab52613ac3c5913655b64943b00ea04913b31eb8a8bd3d52a5c2e1fd77706eb99e5122c4daf7106ab551d7f5f1ed59e65c7fa6718193ab4b70e8bdddd39683b664dd81eab2d6987bf30a652dbc28fc28ae1b5bdcb0d814e20d4afd5cca818fe5324ba38b10efd5f3a4e4b5d63d9fced5edbfcedbad02dc59aa05e9f73ad8f372354b802a0fdba99b16c04e7d4ffa5f319b64e828d4f55fa6f92d5baf8f28637b25429c24852dc3e08aa39eb24de5c9f26cc646552dd23f4fd3c9cb9e3cd1bbe9812c6abd097029c90ddacdb5a378f97da1d36006ddba4a0cad9df3211becfc6912c16d521b58de52b1656fe0aee312f82310bbde18d2da5e177cc8b1f34a8635d14ba60a863c615a59decff60a9ba2e4cdbeaff71bf1754b10e46fd59156c76bffed51b0448d2eca119008ffb26637a60afaffd5e5a2a2d55ef69af90f626ab7a83a0ca44fdd609b6ea2e9b8ab85dd8b54615a17464135cbf4752c3a6d81bd475b96baba73fa44d612ea7f84d8583408a7e2f97a5c6b2e83d850621f471b0bedeb8a667cca419f43198652ccd065d2cb02d0c10867089630323ced2fb60531820251a61ef787bb8b944045ca655b90af19a1c6354bfd73a678b1bdb1285e4fec465a14e0b43b8629456a7c757b1ec7f81656a2f1538fbc80495b31fc231aba604b240d16f353c85debf134ec4fd03c03128144e519853777be96ffaddb6da36fbdab675a60668ae4ed5dbc94e4dc2aa56d4bd50d902ccddf9bf79a5b3e6bb622a5ebb6582e625477bcdfb757141fb3d5749b51655f28114f7b8280b883be1d48926abd8cd83b89e5115628191ef8b02d58868ec1060c539d30fa3d0bf2eaaf87e4d72ecf75c99d3a2b1d75c057dab80410cbb27c020b2ba028e6262b213604f9e748041b596fd2b1dc456ec4ff657c08440757e7b75d9d9886c0c7d2bb509386082e2fccdb8ef6c1bc03086dd8d4dc259a5b23ad703ca74c1c90262e855459e2def7a9d45906bec7afb6b938514b89e5747847b749d1e8589a48be14e146626e7c39dff040000ffff7cb9c313e1140000"
	tmp.Length = 5345
	RESOURCES = append(RESOURCES, tmp)

}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	for _, entry := range RESOURCES {
		//
		// We found the file contents.
		//
		if entry.Filename == path {
			var raw bytes.Buffer

			// Decode the data.
			in, err := hex.DecodeString(entry.Contents)
			if err != nil {
				return nil, err
			}

			// Gunzip the data to the client
			gr, err := gzip.NewReader(bytes.NewBuffer(in))
			defer gr.Close()
			data, err := ioutil.ReadAll(gr)
			if err != nil {
				return nil, err
			}
			_, err = raw.Write(data)
			if err != nil {
				return nil, err
			}

			// Return it.
			return raw.Bytes(), nil
		}
	}
	return nil, errors.New("Failed to find resource")
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
