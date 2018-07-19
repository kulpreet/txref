package util

type TestVector struct {
	Magic int
    Hrp string
    EncodedTxref string
    Height int
    Position int
    Vout int
    EncFail int //0 == must not fail, 1 == can fail, 2 == can fail and continue with next test, 3 == skip
    DecFail int
    NonStd bool
}

var ExtendedTestVectors = []TestVector{
	{
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
		"tx1:rqqq-qqqq-qqqu-au7hl",
        0,
        0,
		0,
        0,0,false,
    },

	{
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5ng-qqq8-lsnk3",
        466793,
        2205,
		0,
        0,0,false,
    },

    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5n1-rrr0-lsnk3", /* error correct test >rrr0< instead of >qqq8<*/
        466793,
        2205,
		0,
        1,0,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qqqq-qqqu-au7hl",
        0,
        0,
		0,
        0,0,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rzqq-qqqq-qqql-4ym2c",
        1,
        0,
		0,
        0,0,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5ng-4jsf-mcsfu", /* complete invalid */
        0,
        0,
		0,
        1,1,false, /* enc & dec must fail */
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-lrar-qqqw-jmhax",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        1000,
		0,
        0,0,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "", /* encoding must fail, no txref to chain against */
        2097152, /* invalid height */
        1000,
		0,
        2,1,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-llll-qqqn-sr5q8",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8191, /* last valid tx pos is 0x1FFF */
		0,
        0,0,false,
    },

    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-llll-ll8y-5yj2q",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8191, /* last valid tx pos is 0x1FFF */
		8191,
        0,0,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "", /* encoding must fail, no txref to chain against */
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8192, /* invalid tx pos */
		0,
        2,1,false,
    },

    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "", /* encoding must fail, no txref to chain against */
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8191, 
		8192, /* invalid vout */
        2,1,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-lrqq-qqqm-m5vjv",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        0,
		0,
        0,0,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qull-qqq5-ktx95",
        0,
        8191, /* last valid tx pos is 0x1FFF */
		0,
        0,0,false,
    },

    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qqqq-ll8t-emcac",
        0,
        0, 
		8191, /* last valid vout is 0x1FFF */
        0,0,false,
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5ng-gghq-fkg7", /* valid Bech32, but 10x5bit packages instead of 8 */
        0,
        0,
		0,
        3,2,false, /* ignore encoding */
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5qd-s43z", /* valid Bech32, but 6x5bit packages instead of 8 */
        0,
        0,
		0,
        3,2,false, /* ignore encoding */
    },
	
    {
        0xB,
        TxrefBech32HrpMainnet,
        "tx1:t7ll-llll-qqqt-433dq",
        2097151,
        8191,
		0,
        0,0,false, /* ignore encoding */
    },

    {
        0xB,
        TxrefBech32HrpMainnet,
        "tx1:t7ll-llll-ll8u-3kh88",
        2097151,
        8191,
		8191,
        0,0,false, /* ignore encoding */
    },
	
    {
        TxrefMagiBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rk63-uvxf-qqqa-8wrdy",
        467883,
        2355,
		0,
        0,0,false, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        Txref_bech32_hrp_testnet,
        "txtest1:xk63-uqvx-fqqq-q5p7-cqx",
        467883,
        2355,
		0,
        0,0,true, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        Txref_bech32_hrp_testnet,
        "txtest1:xqqq-qqqq-qqqq-qj7d-vzy",
        0,
        0,
		0,
        0,0,true, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        Txref_bech32_hrp_testnet,
        "txtest1:x7ll-llll-llqq-qz4n-9zn",
        0x3FFFFFF,
        0x3FFFF,
		0,
        0,0,true, /* ignore encoding */
    },
}
