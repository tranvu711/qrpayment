package qrpayment

type BankList struct {
	VietinBank        Bank
	Vietcombank       Bank
	BIDV              Bank
	Agribank          Bank
	OCB               Bank
	MBBank            Bank
	Techcombank       Bank
	ACB               Bank
	VPBank            Bank
	TPBank            Bank
	Sacombank         Bank
	HDBank            Bank
	BVBank            Bank
	SCB               Bank
	VIB               Bank
	SHB               Bank
	Eximbank          Bank
	MSB               Bank
	CAKE              Bank
	Ubank             Bank
	Timo              Bank
	ViettelMoney      Bank
	VNPTMoney         Bank
	SaigonBank        Bank
	BacABank          Bank
	PVcomBank         Bank
	Oceanbank         Bank
	NCB               Bank
	ShinhanBank       Bank
	ABBANK            Bank
	VietABank         Bank
	NamABank          Bank
	PGBank            Bank
	VietBank          Bank
	BaoVietBank       Bank
	SeABank           Bank
	COOPBANK          Bank
	LienVietPostBank  Bank
	KienLongBank      Bank
	KBank             Bank
	KookminHN         Bank
	KEBHanaHCM        Bank
	KEBHANAHN         Bank
	MAFC              Bank
	Citibank          Bank
	KookminHCM        Bank
	VBSP              Bank
	Woori             Bank
	VRB               Bank
	UnitedOverseas    Bank
	StandardChartered Bank
	PublicBank        Bank
	Nonghyup          Bank
	IndovinaBank      Bank
	IBKHCM            Bank
	IBKHN             Bank
	HSBC              Bank
	HongLeong         Bank
	GPBank            Bank
	DongABank         Bank
	DBSBank           Bank
	CIMB              Bank
	CBBank            Bank
}

var BANKS = BankList{
	VietinBank: Bank{
		BIN:       "970415",
		Name:      "Ngân hàng TMCP Công thương Việt Nam",
		ShortName: "VietinBank",
		Code:      "ICB",
	},
	Vietcombank: Bank{
		BIN:       "970436",
		Name:      "Ngân hàng TMCP Ngoại Thương Việt Nam",
		ShortName: "Vietcombank",
		Code:      "VCB",
	},
	BIDV: Bank{
		BIN:       "970418",
		Name:      "Ngân hàng TMCP Đầu tư và Phát triển Việt Nam",
		ShortName: "BIDV",
		Code:      "BIDV",
	},
	Agribank: Bank{
		BIN:       "970405",
		Name:      "Ngân hàng Nông nghiệp và Phát triển Nông thôn Việt Nam",
		ShortName: "Agribank",
		Code:      "VBA",
	},
	OCB: Bank{
		BIN:       "970448",
		Name:      "Ngân hàng TMCP Phương Đông",
		ShortName: "OCB",
		Code:      "OCB",
	},
	MBBank: Bank{
		BIN:       "970422",
		Name:      "Ngân hàng TMCP Quân đội",
		ShortName: "MBBank",
		Code:      "MB",
	},
	Techcombank: Bank{
		BIN:       "970407",
		Name:      "Ngân hàng TMCP Kỹ thương Việt Nam",
		ShortName: "Techcombank",
		Code:      "TCB",
	},
	ACB: Bank{
		BIN:       "970416",
		Name:      "Ngân hàng TMCP Á Châu",
		ShortName: "ACB",
		Code:      "ACB",
	},
	VPBank: Bank{
		BIN:       "970432",
		Name:      "Ngân hàng TMCP Việt Nam Thịnh Vượng",
		ShortName: "VPBank",
		Code:      "VPB",
	},
	TPBank: Bank{
		BIN:       "970423",
		Name:      "Ngân hàng TMCP Tiên Phong",
		ShortName: "TPBank",
		Code:      "TPB",
		SwiftCode: "",
	},
	Sacombank: Bank{
		BIN:       "970403",
		Name:      "Ngân hàng TMCP Sài Gòn Thương Tín",
		ShortName: "Sacombank",
		Code:      "STB",
	},
	HDBank: Bank{
		BIN:       "970437",
		Name:      "Ngân hàng TMCP Phát triển Thành phố Hồ Chí Minh",
		ShortName: "HDBank",
		Code:      "HDB",
	},
	BVBank: Bank{
		BIN:       "970454",
		Name:      "Ngân hàng TMCP Bản Việt",
		ShortName: "VietCapitalBank - BVBank",
		Code:      "VCCB",
	},
	SCB: Bank{
		BIN:       "970429",
		Name:      "Ngân hàng TMCP Sài Gòn",
		ShortName: "SCB",
		Code:      "SCB",
	},
	VIB: Bank{
		BIN:       "970441",
		Name:      "Ngân hàng TMCP Quốc tế Việt Nam",
		ShortName: "VIB",
		Code:      "VIB",
	},
	SHB: Bank{
		BIN:       "970443",
		Name:      "Ngân hàng TMCP Sài Gòn - Hà Nội",
		ShortName: "SHB",
		Code:      "SHB",
	},
	Eximbank: Bank{
		BIN:       "970431",
		Name:      "Ngân hàng TMCP Xuất Nhập khẩu Việt Nam",
		ShortName: "Eximbank",
		Code:      "EIB",
	},
	MSB: Bank{
		BIN:       "970426",
		Name:      "Ngân hàng TMCP Hàng Hải",
		ShortName: "MSB",
		Code:      "MSB",
	},
	CAKE: Bank{
		BIN:       "546034",
		Name:      "TMCP Việt Nam Thịnh Vượng - Ngân hàng số CAKE by VPBank",
		ShortName: "CAKE",
		Code:      "CAKE",
	},
	Ubank: Bank{
		BIN:       "546035",
		Name:      "TMCP Việt Nam Thịnh Vượng - Ngân hàng số Ubank by VPBank",
		ShortName: "Ubank",
		Code:      "Ubank",
	},
	Timo: Bank{
		BIN:       "963388",
		Name:      "Ngân hàng số Timo by Ban Viet Bank (Timo by Ban Viet Bank)",
		ShortName: "Timo",
		Code:      "TIMO",
	},
	ViettelMoney: Bank{
		BIN:       "971005",
		Name:      "Tổng Công ty Dịch vụ số Viettel - Chi nhánh tập đoàn công nghiệp viễn thông Quân Đội",
		ShortName: "ViettelMoney",
		Code:      "VTLMONEY",
	},
	VNPTMoney: Bank{
		BIN:       "971011",
		Name:      "VNPT Money",
		ShortName: "VNPTMoney",
		Code:      "VNPTMONEY",
	},
	SaigonBank: Bank{
		BIN:       "970400",
		Name:      "Ngân hàng TMCP Sài Gòn Công Thương",
		ShortName: "SaigonBank",
		Code:      "SGICB",
	},
	BacABank: Bank{
		BIN:       "970409",
		Name:      "Ngân hàng TMCP Bắc Á",
		ShortName: "BacABank",
		Code:      "BAB",
	},
	PVcomBank: Bank{
		BIN:       "970412",
		Name:      "Ngân hàng TMCP Đại Chúng Việt Nam",
		ShortName: "PVcomBank",
		Code:      "PVCB",
	},
	Oceanbank: Bank{
		BIN:       "970414",
		Name:      "Ngân hàng Thương mại TNHH MTV Đại Dương",
		ShortName: "Oceanbank",
		Code:      "Oceanbank",
	},
	NCB: Bank{
		BIN:       "970419",
		Name:      "Ngân hàng TMCP Quốc Dân",
		ShortName: "NCB",
		Code:      "NCB",
	},
	ShinhanBank: Bank{
		BIN:       "970424",
		Name:      "Ngân hàng TNHH MTV Shinhan Việt Nam",
		ShortName: "ShinhanBank",
		Code:      "SHBVN",
	},
	ABBANK: Bank{
		BIN:       "970425",
		Name:      "Ngân hàng TMCP An Bình",
		ShortName: "ABBANK",
		Code:      "ABB",
	},
	VietABank: Bank{
		BIN:       "970427",
		Name:      "Ngân hàng TMCP Việt Á",
		ShortName: "VietABank",
		Code:      "VAB",
	},
	NamABank: Bank{
		BIN:       "970428",
		Name:      "Ngân hàng TMCP Nam Á",
		ShortName: "NamABank",
		Code:      "NAB",
	},
	PGBank: Bank{
		BIN:       "970430",
		Name:      "Ngân hàng TMCP Xăng dầu Petrolimex",
		ShortName: "PGBank",
		Code:      "PGB",
	},
	VietBank: Bank{
		BIN:       "970433",
		Name:      "Ngân hàng TMCP Việt Nam Thương Tín",
		ShortName: "VietBank",
		Code:      "VIETBANK",
	},
	BaoVietBank: Bank{
		BIN:       "970438",
		Name:      "Ngân hàng TMCP Bảo Việt",
		ShortName: "BaoVietBank",
		Code:      "BVB",
	},
	SeABank: Bank{
		BIN:       "970440",
		Name:      "Ngân hàng TMCP Đông Nam Á",
		ShortName: "SeABank",
		Code:      "SEAB",
	},
	COOPBANK: Bank{
		BIN:       "970446",
		Name:      "Ngân hàng Hợp tác xã Việt Nam",
		ShortName: "COOPBANK",
		Code:      "COOPBANK",
	},
	LienVietPostBank: Bank{
		BIN:       "970449",
		Name:      "Ngân hàng TMCP Bưu Điện Liên Việt",
		ShortName: "LienVietPostBank",
		Code:      "LPB",
	},
	KienLongBank: Bank{
		BIN:       "970452",
		Name:      "Ngân hàng TMCP Kiên Long",
		ShortName: "KienLongBank",
		Code:      "KLB",
	},
	KBank: Bank{
		BIN:       "668888",
		Name:      "Ngân hàng Đại chúng TNHH Kasikornbank",
		ShortName: "KBank",
		Code:      "KBank",
	},
	KookminHN: Bank{
		BIN:       "970462",
		Name:      "Ngân hàng Kookmin - Chi nhánh Hà Nội",
		ShortName: "KookminHN",
		Code:      "KBHN",
	},
	KEBHanaHCM: Bank{
		BIN:       "970466",
		Name:      "Ngân hàng KEB Hana – Chi nhánh Thành phố Hồ Chí Minh",
		ShortName: "KEBHanaHCM",
		Code:      "KEBHANAHCM",
	},
	KEBHANAHN: Bank{
		BIN:       "970467",
		Name:      "Ngân hàng KEB Hana – Chi nhánh Hà Nội",
		ShortName: "KEBHANAHN",
		Code:      "KEBHANAHN",
	},
	MAFC: Bank{
		BIN:       "977777",
		Name:      "Công ty Tài chính TNHH MTV Mirae Asset (Việt Nam) ",
		ShortName: "MAFC",
		Code:      "MAFC",
	},
	Citibank: Bank{
		BIN:       "533948",
		Name:      "Ngân hàng Citibank, N.A. - Chi nhánh Hà Nội",
		ShortName: "Citibank",
		Code:      "CITIBANK",
	},
	KookminHCM: Bank{
		BIN:       "970463",
		Name:      "Ngân hàng Kookmin - Chi nhánh Thành phố Hồ Chí Minh",
		ShortName: "KookminHCM",
		Code:      "KBHCM",
	},
	VBSP: Bank{
		BIN:       "999888",
		Name:      "Ngân hàng Chính sách Xã hội",
		ShortName: "VBSP",
		Code:      "VBSP",
	},
	Woori: Bank{
		BIN:       "970457",
		Name:      "Ngân hàng TNHH MTV Woori Việt Nam",
		ShortName: "Woori",
		Code:      "WVN",
	},
	VRB: Bank{
		BIN:       "970421",
		Name:      "Ngân hàng Liên doanh Việt - Nga",
		ShortName: "VRB",
		Code:      "VRB",
	},
	UnitedOverseas: Bank{
		BIN:       "970458",
		Name:      "Ngân hàng United Overseas - Chi nhánh TP. Hồ Chí Minh",
		ShortName: "UnitedOverseas",
		Code:      "UOB",
	},
	StandardChartered: Bank{
		BIN:       "970410",
		Name:      "Ngân hàng TNHH MTV Standard Chartered Bank Việt Nam",
		ShortName: "StandardChartered",
		Code:      "SCVN",
	},
	PublicBank: Bank{
		BIN:       "970439",
		Name:      "Ngân hàng TNHH MTV Public Việt Nam",
		ShortName: "PublicBank",
		Code:      "PBVN",
	},
	Nonghyup: Bank{
		BIN:       "801011",
		Name:      "Ngân hàng Nonghyup - Chi nhánh Hà Nội",
		ShortName: "Nonghyup",
		Code:      "NHB HN",
	},
	IndovinaBank: Bank{
		BIN:       "970434",
		Name:      "Ngân hàng TNHH Indovina",
		ShortName: "IndovinaBank",
		Code:      "IVB",
	},
	IBKHCM: Bank{
		BIN:       "970456",
		Name:      "Ngân hàng Công nghiệp Hàn Quốc - Chi nhánh TP. Hồ Chí Minh",
		ShortName: "IBKHCM",
		Code:      "IBK - HCM",
	},
	IBKHN: Bank{
		BIN:       "970455",
		Name:      "Ngân hàng Công nghiệp Hàn Quốc - Chi nhánh Hà Nội",
		ShortName: "IBKHN",
		Code:      "IBK - HN",
	},
	HSBC: Bank{
		BIN:       "458761",
		Name:      "Ngân hàng TNHH MTV HSBC (Việt Nam)",
		ShortName: "HSBC",
		Code:      "HSBC",
	},
	HongLeong: Bank{
		BIN:       "970442",
		Name:      "Ngân hàng TNHH MTV Hong Leong Việt Nam",
		ShortName: "HongLeong",
		Code:      "HLBVN",
	},
	GPBank: Bank{
		BIN:       "970408",
		Name:      "Ngân hàng Thương mại TNHH MTV Dầu Khí Toàn Cầu",
		ShortName: "GPBank",
		Code:      "GPB",
	},
	DongABank: Bank{
		BIN:       "970406",
		Name:      "Ngân hàng TMCP Đông Á",
		ShortName: "DongABank",
		Code:      "DOB",
	},
	DBSBank: Bank{
		BIN:       "796500",
		Name:      "DBS Bank Ltd - Chi nhánh Thành phố Hồ Chí Minh",
		ShortName: "DBSBank",
		Code:      "DBS",
	},
	CIMB: Bank{
		BIN:       "422589",
		Name:      "Ngân hàng TNHH MTV CIMB Việt Nam",
		ShortName: "CIMB",
		Code:      "CIMB",
	},
	CBBank: Bank{
		BIN:       "970444",
		Name:      "Ngân hàng Thương mại TNHH MTV Xây dựng Việt Nam",
		ShortName: "CBBank",
		Code:      "CBB",
	},
}
