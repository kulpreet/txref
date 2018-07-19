package util

type ExtendedTestVector struct {
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

var ExtendedTestVectors = []ExtendedTestVector{
	{
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
		"tx1:rqqq-qqqq-qmhu-qk",
        0,
        0,
		0,
        0,0,false,
    },

	{
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
		"tx1:rqqq-qqqq-pqq7-l4fw8",
        0,
        0,
		1,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qqqq-rrr8-l4fw8", /* error correct test >rrr8< instead of >qqq8<*/
        466793,
        2205,
		0,
        1,0,false,
    },
		
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rzqq-qqqq-pqqa-hdvnq",
        1,
        0,
		1,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5ng-4jsf-mcsfu", /* complete invalid */
        0,
        0,
		0,
        1,1,false, /* enc & dec must fail */
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-lrar-pqqv-sjqy7",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        1000,
		1,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "", /* encoding must fail, no txref to chain against */
        2097152, /* invalid height */
        1000,
		0,
        2,1,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-llll-pqq3-j2rel",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8191, /* last valid tx pos is 0x1FFF */
		1,
        0,0,false,
    },

    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-llll-ll8y-5yj2q",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8191, /* last valid tx pos is 0x1FFF */
		8191,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "", /* encoding must fail, no txref to chain against */
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8192, /* invalid tx pos */
		0,
        2,1,false,
    },

    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "", /* encoding must fail, no txref to chain against */
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8191, 
		8192, /* invalid vout */
        2,1,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:r7ll-lrqq-pqqe-eamt5",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        0,
		1,
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qull-pqqk-5z3uv",
        0,
        8191, /* last valid tx pos is 0x1FFF */
		1,
        0,0,false,
    },

    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rqqq-qqqq-ll8t-emcac",
        0,
        0, 
		8191, /* last valid vout is 0x1FFF */
        0,0,false,
    },
	
    {
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rjk0-u5ng-gghq-fkg7", /* valid Bech32, but 10x5bit packages instead of 8 */
        0,
        0,
		0,
        3,2,false, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcMainnet,
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
        "tx1:t7ll-llll-pqqf-hcx5c",
        2097151,
        8191,
		1,
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
        TxrefMagicBtcMainnet,
        TxrefBech32HrpMainnet,
        "tx1:rk63-uvxf-pqql-9855u",
        467883,
        2355,
		1,
        0,0,false, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        TxrefBech32HrpTestnet,
        "txtest1:xk63-uqvx-fqpq-qkrh-0e7",
        467883,
        2355,
		1,
        0,0,true, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        TxrefBech32HrpTestnet,
        "txtest1:xqqq-qqqq-qqll-8962-2gr",
        0,
        0,
		8191,
        0,0,true, /* ignore encoding */
    },
	
    {
        TxrefMagicBtcTestnet,
        TxrefBech32HrpTestnet,
        "txtest1:x7ll-llll-llll-8435-rg5",
        0x3FFFFFF,
        0x3FFFF,
		0x1FFF,
        0,0,true, /* ignore encoding */
    },
}
